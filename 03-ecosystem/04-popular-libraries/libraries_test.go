package libraries

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// MockTestingT 模拟测试接口
type MockTestingT struct {
	errors  []string
	failNow bool
}

func (m *MockTestingT) Errorf(format string, args ...interface{}) {
	m.errors = append(m.errors, fmt.Sprintf(format, args...))
}

func (m *MockTestingT) FailNow() {
	m.failNow = true
}

func TestLogrusLikeLogger(t *testing.T) {
	logger := NewLogger()

	t.Run("BasicLogging", func(t *testing.T) {
		// 测试基本日志记录（这里只是确保不会panic）
		logger.Info("test info message")
		logger.Error("test error message")
		logger.Warn("test warning message")
		logger.Debug("test debug message")

		t.Log("基本日志记录测试通过")
	})

	t.Run("WithField", func(t *testing.T) {
		fieldLogger := logger.WithField("user_id", 12345)

		// 验证原logger没有被修改
		if len(logger.fields) != 0 {
			t.Error("Original logger should not be modified")
		}

		// 验证新logger有字段
		if len(fieldLogger.fields) != 1 {
			t.Errorf("Field logger should have 1 field, got %d", len(fieldLogger.fields))
		}

		if fieldLogger.fields["user_id"] != 12345 {
			t.Errorf("Field value should be 12345, got %v", fieldLogger.fields["user_id"])
		}

		fieldLogger.Info("test message with field")

		t.Log("WithField测试通过")
	})

	t.Run("WithFields", func(t *testing.T) {
		fields := map[string]interface{}{
			"module": "test",
			"action": "testing",
			"count":  42,
		}

		fieldsLogger := logger.WithFields(fields)

		if len(fieldsLogger.fields) != 3 {
			t.Errorf("Fields logger should have 3 fields, got %d", len(fieldsLogger.fields))
		}

		for key, value := range fields {
			if fieldsLogger.fields[key] != value {
				t.Errorf("Field %s should be %v, got %v", key, value, fieldsLogger.fields[key])
			}
		}

		fieldsLogger.Info("test message with multiple fields")

		t.Log("WithFields测试通过")
	})

	t.Run("ChainedFields", func(t *testing.T) {
		chainedLogger := logger.WithField("step", 1).
			WithField("module", "test").
			WithFields(map[string]interface{}{
				"action": "chain",
				"count":  3,
			})

		if len(chainedLogger.fields) != 4 {
			t.Errorf("Chained logger should have 4 fields, got %d", len(chainedLogger.fields))
		}

		chainedLogger.Info("test chained fields")

		t.Log("链式字段测试通过")
	})
}

func TestViperLikeConfig(t *testing.T) {
	config := NewConfig()

	t.Run("SetAndGet", func(t *testing.T) {
		config.Set("test.key", "test value")

		value := config.Get("test.key")
		if value != "test value" {
			t.Errorf("Get should return 'test value', got %v", value)
		}

		t.Log("设置和获取配置测试通过")
	})

	t.Run("TypedGetters", func(t *testing.T) {
		config.Set("string.key", "hello")
		config.Set("int.key", 42)
		config.Set("bool.key", true)

		// 测试字符串获取
		strValue := config.GetString("string.key")
		if strValue != "hello" {
			t.Errorf("GetString should return 'hello', got %s", strValue)
		}

		// 测试整数获取
		intValue := config.GetInt("int.key")
		if intValue != 42 {
			t.Errorf("GetInt should return 42, got %d", intValue)
		}

		// 测试布尔获取
		boolValue := config.GetBool("bool.key")
		if !boolValue {
			t.Error("GetBool should return true")
		}

		t.Log("类型化获取器测试通过")
	})

	t.Run("DefaultValues", func(t *testing.T) {
		newConfig := NewConfig()

		newConfig.SetDefault("default.key", "default value")
		newConfig.SetDefault("default.int", 100)

		// 获取默认值
		if newConfig.GetString("default.key") != "default value" {
			t.Error("Should return default value")
		}

		// 设置新值应该覆盖默认值
		newConfig.Set("default.key", "new value")
		if newConfig.GetString("default.key") != "new value" {
			t.Error("Should return new value, not default")
		}

		t.Log("默认值测试通过")
	})

	t.Run("FileOperations", func(t *testing.T) {
		tempFile := filepath.Join(os.TempDir(), "test_config.json")
		defer os.Remove(tempFile)

		// 设置一些配置
		config.Set("app.name", "Test App")
		config.Set("app.version", "1.0.0")
		config.Set("debug", true)

		// 写入文件
		err := config.WriteConfig(tempFile)
		if err != nil {
			t.Fatalf("WriteConfig failed: %v", err)
		}

		// 从文件读取
		newConfig := NewConfig()
		err = newConfig.ReadInConfig(tempFile)
		if err != nil {
			t.Fatalf("ReadInConfig failed: %v", err)
		}

		// 验证读取的配置
		if newConfig.GetString("app.name") != "Test App" {
			t.Error("Config not properly read from file")
		}

		if newConfig.GetString("app.version") != "1.0.0" {
			t.Error("Config version not properly read from file")
		}

		if !newConfig.GetBool("debug") {
			t.Error("Config debug flag not properly read from file")
		}

		t.Log("文件操作测试通过")
	})
}

func TestCobraLikeCommand(t *testing.T) {
	t.Run("CommandCreation", func(t *testing.T) {
		cmd := NewCommand("test", "Test command")

		if cmd.Name != "test" {
			t.Errorf("Command name should be 'test', got %s", cmd.Name)
		}

		if cmd.Description != "Test command" {
			t.Errorf("Command description should be 'Test command', got %s", cmd.Description)
		}

		t.Log("命令创建测试通过")
	})

	t.Run("Flags", func(t *testing.T) {
		cmd := NewCommand("test", "Test command")

		cmd.StringFlag("name", "n", "default", "Name flag")
		cmd.IntFlag("count", "c", 10, "Count flag")
		cmd.BoolFlag("verbose", "v", false, "Verbose flag")

		if len(cmd.Flags) != 3 {
			t.Errorf("Should have 3 flags, got %d", len(cmd.Flags))
		}

		// 测试标志获取
		if cmd.GetStringFlag("name") != "default" {
			t.Error("String flag default value incorrect")
		}

		if cmd.GetIntFlag("count") != 10 {
			t.Error("Int flag default value incorrect")
		}

		if cmd.GetBoolFlag("verbose") != false {
			t.Error("Bool flag default value incorrect")
		}

		t.Log("标志测试通过")
	})

	t.Run("SubCommands", func(t *testing.T) {
		rootCmd := NewCommand("root", "Root command")
		subCmd := NewCommand("sub", "Sub command")

		rootCmd.AddCommand(subCmd)

		if len(rootCmd.SubCommands) != 1 {
			t.Errorf("Should have 1 subcommand, got %d", len(rootCmd.SubCommands))
		}

		if rootCmd.SubCommands["sub"] != subCmd {
			t.Error("Subcommand not properly added")
		}

		t.Log("子命令测试通过")
	})

	t.Run("CommandExecution", func(t *testing.T) {
		executed := false

		cmd := NewCommand("test", "Test command")
		cmd.RunFunc = func(cmd *CobraLikeCommand, args []string) error {
			executed = true
			return nil
		}

		err := cmd.Execute([]string{})
		if err != nil {
			t.Fatalf("Command execution failed: %v", err)
		}

		if !executed {
			t.Error("Command RunFunc should have been executed")
		}

		t.Log("命令执行测试通过")
	})

	t.Run("FlagParsing", func(t *testing.T) {
		cmd := NewCommand("test", "Test command")
		cmd.StringFlag("name", "n", "default", "Name flag")
		cmd.IntFlag("count", "c", 10, "Count flag")

		cmd.RunFunc = func(cmd *CobraLikeCommand, args []string) error {
			name := cmd.GetStringFlag("name")
			count := cmd.GetIntFlag("count")

			if name != "testname" {
				t.Errorf("Name flag should be 'testname', got %s", name)
			}

			if count != 42 {
				t.Errorf("Count flag should be 42, got %d", count)
			}

			return nil
		}

		err := cmd.Execute([]string{"--name", "testname", "--count", "42"})
		if err != nil {
			t.Fatalf("Command execution with flags failed: %v", err)
		}

		t.Log("标志解析测试通过")
	})
}

func TestTestifyLikeAssert(t *testing.T) {
	mockT := &MockTestingT{}
	assert := NewAssert(mockT)

	t.Run("EqualAssertion", func(t *testing.T) {
		// 测试成功断言
		result := assert.Equal(42, 42)
		if !result {
			t.Error("Equal assertion should succeed")
		}

		if len(mockT.errors) != 0 {
			t.Error("Should have no errors for successful assertion")
		}

		// 测试失败断言
		mockT.errors = nil // 重置错误
		result = assert.Equal(42, 43)
		if result {
			t.Error("Equal assertion should fail")
		}

		if len(mockT.errors) != 1 {
			t.Errorf("Should have 1 error for failed assertion, got %d", len(mockT.errors))
		}

		t.Log("相等断言测试通过")
	})

	t.Run("NotEqualAssertion", func(t *testing.T) {
		mockT.errors = nil

		// 测试成功断言
		result := assert.NotEqual(42, 43)
		if !result {
			t.Error("NotEqual assertion should succeed")
		}

		// 测试失败断言
		result = assert.NotEqual(42, 42)
		if result {
			t.Error("NotEqual assertion should fail")
		}

		t.Log("不相等断言测试通过")
	})

	t.Run("NilAssertion", func(t *testing.T) {
		mockT.errors = nil

		// 测试成功断言
		result := assert.Nil(nil)
		if !result {
			t.Error("Nil assertion should succeed")
		}

		// 测试失败断言
		result = assert.Nil("not nil")
		if result {
			t.Error("Nil assertion should fail")
		}

		t.Log("Nil断言测试通过")
	})

	t.Run("NotNilAssertion", func(t *testing.T) {
		mockT.errors = nil

		// 测试成功断言
		result := assert.NotNil("not nil")
		if !result {
			t.Error("NotNil assertion should succeed")
		}

		// 测试失败断言
		result = assert.NotNil(nil)
		if result {
			t.Error("NotNil assertion should fail")
		}

		t.Log("NotNil断言测试通过")
	})

	t.Run("BooleanAssertions", func(t *testing.T) {
		mockT.errors = nil

		// True断言
		if !assert.True(true) {
			t.Error("True assertion should succeed")
		}

		if assert.True(false) {
			t.Error("True assertion should fail")
		}

		// False断言
		if !assert.False(false) {
			t.Error("False assertion should succeed")
		}

		if assert.False(true) {
			t.Error("False assertion should fail")
		}

		t.Log("布尔断言测试通过")
	})
}

// 基准测试
func BenchmarkLoggerWithField(b *testing.B) {
	logger := NewLogger()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.WithField("iteration", i).Info("benchmark message")
	}
}

func BenchmarkConfigGet(b *testing.B) {
	config := NewConfig()
	config.Set("test.key", "test value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config.GetString("test.key")
	}
}

func BenchmarkCommandExecution(b *testing.B) {
	cmd := NewCommand("bench", "Benchmark command")
	cmd.RunFunc = func(cmd *CobraLikeCommand, args []string) error {
		return nil
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmd.Execute([]string{})
	}
}

func BenchmarkAssertEqual(b *testing.B) {
	mockT := &MockTestingT{}
	assert := NewAssert(mockT)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		assert.Equal(42, 42)
	}
}
