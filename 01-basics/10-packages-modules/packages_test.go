package packages

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPackageVisibility(t *testing.T) {
	t.Run("PublicFields", func(t *testing.T) {
		pv := NewPackageVisibility("test", 42)

		// 测试公开字段访问
		if pv.PublicField != "test" {
			t.Errorf("PublicField: got %s, want test", pv.PublicField)
		}

		if pv.ExportedValue != 42 {
			t.Errorf("ExportedValue: got %d, want 42", pv.ExportedValue)
		}

		t.Log("公开字段访问测试通过")
	})

	t.Run("PrivateFieldAccess", func(t *testing.T) {
		pv := NewPackageVisibility("test", 42)

		// 测试通过方法访问私有字段
		privateValue := pv.GetPrivateField()
		if privateValue != "internal" {
			t.Errorf("GetPrivateField: got %s, want internal", privateValue)
		}

		t.Log("私有字段访问测试通过")
	})

	t.Run("InternalMethodCall", func(t *testing.T) {
		pv := NewPackageVisibility("test", 42)

		// 测试通过公开方法调用私有方法
		pv.UpdateInternal(100)
		// 注意：我们无法直接验证私有字段的值，但可以确保方法调用不出错

		t.Log("内部方法调用测试通过")
	})
}

func TestPackageConstants(t *testing.T) {
	t.Run("ExportedConstants", func(t *testing.T) {
		if Version == "" {
			t.Error("Version should not be empty")
		}

		if Author == "" {
			t.Error("Author should not be empty")
		}

		if MaxFileSize <= 0 {
			t.Error("MaxFileSize should be positive")
		}

		t.Logf("Version: %s, Author: %s, MaxFileSize: %d", Version, Author, MaxFileSize)
		t.Log("导出常量测试通过")
	})

	t.Run("ExportedVariables", func(t *testing.T) {
		if DefaultTimeout <= 0 {
			t.Error("DefaultTimeout should be positive")
		}

		if MaxRetries <= 0 {
			t.Error("MaxRetries should be positive")
		}

		t.Logf("DefaultTimeout: %d, MaxRetries: %d", DefaultTimeout, MaxRetries)
		t.Log("导出变量测试通过")
	})
}

func TestPackageFunctions(t *testing.T) {
	t.Run("GetVersion", func(t *testing.T) {
		version := GetVersion()
		if version != Version {
			t.Errorf("GetVersion: got %s, want %s", version, Version)
		}

		t.Log("GetVersion函数测试通过")
	})

	t.Run("IsDebugMode", func(t *testing.T) {
		debugMode := IsDebugMode()
		// debugMode的值取决于环境变量，我们只测试函数能正常调用
		t.Logf("Debug mode: %t", debugMode)
		t.Log("IsDebugMode函数测试通过")
	})

	t.Run("ProcessInput", func(t *testing.T) {
		// 测试有效输入
		validInput := "test input"
		result, err := ProcessInput(validInput)
		if err != nil {
			t.Errorf("ProcessInput with valid input failed: %v", err)
		}

		if string(result) != validInput {
			t.Errorf("ProcessInput result: got %s, want %s", result, validInput)
		}

		// 测试无效输入
		invalidInput := ""
		_, err = ProcessInput(invalidInput)
		if err == nil {
			t.Error("ProcessInput should fail with empty input")
		}

		t.Log("ProcessInput函数测试通过")
	})
}

func TestTypeAliases(t *testing.T) {
	t.Run("UserIDType", func(t *testing.T) {
		var userID UserID = 12345

		// 类型别名应该与底层类型兼容
		var normalInt int64 = int64(userID)
		if normalInt != 12345 {
			t.Errorf("UserID conversion: got %d, want 12345", normalInt)
		}

		t.Log("UserID类型别名测试通过")
	})

	t.Run("UsernameType", func(t *testing.T) {
		var username Username = "gopher"

		// 类型别名应该与底层类型兼容
		var normalString string = string(username)
		if normalString != "gopher" {
			t.Errorf("Username conversion: got %s, want gopher", normalString)
		}

		t.Log("Username类型别名测试通过")
	})
}

func TestProcessor(t *testing.T) {
	t.Run("ProcessorCreation", func(t *testing.T) {
		processor := NewProcessor("test-config")
		if processor == nil {
			t.Error("NewProcessor should return non-nil processor")
		}

		t.Log("处理器创建测试通过")
	})

	t.Run("ProcessorValidation", func(t *testing.T) {
		processor := NewProcessor("test-config")

		// 测试有效输入
		if !processor.Validate("valid input") {
			t.Error("Processor should validate non-empty input")
		}

		// 测试无效输入
		if processor.Validate("") {
			t.Error("Processor should not validate empty input")
		}

		t.Log("处理器验证测试通过")
	})

	t.Run("ProcessorProcessing", func(t *testing.T) {
		processor := NewProcessor("test-config")

		testData := []byte("test data")
		result, err := processor.Process(testData)
		if err != nil {
			t.Errorf("Processor.Process failed: %v", err)
		}

		if string(result) != string(testData) {
			t.Errorf("Process result: got %s, want %s", result, testData)
		}

		t.Log("处理器处理测试通过")
	})
}

func TestPackageManager(t *testing.T) {
	t.Run("PackageManagerCreation", func(t *testing.T) {
		pm := NewPackageManager("/test/path")
		if pm == nil {
			t.Error("NewPackageManager should return non-nil manager")
		}

		if pm.rootPath != "/test/path" {
			t.Errorf("PackageManager rootPath: got %s, want /test/path", pm.rootPath)
		}

		t.Log("包管理器创建测试通过")
	})

	t.Run("AnalyzeNonExistentPackage", func(t *testing.T) {
		pm := NewPackageManager("/test/path")

		_, err := pm.AnalyzePackage("/non/existent/path")
		if err == nil {
			t.Error("AnalyzePackage should fail for non-existent path")
		}

		if !strings.Contains(err.Error(), "包路径不存在") {
			t.Errorf("Error message should contain '包路径不存在': %v", err)
		}

		t.Log("不存在包分析测试通过")
	})
}

func TestInitFunction(t *testing.T) {
	t.Run("InitFunctionExecution", func(t *testing.T) {
		// init函数在包导入时已经执行
		// 我们可以测试它的副作用

		// 测试环境变量影响
		originalDebug := os.Getenv("DEBUG")
		defer os.Setenv("DEBUG", originalDebug)

		// 这个测试只能验证当前状态，因为init已经执行过了
		debugMode := IsDebugMode()
		t.Logf("当前调试模式: %t", debugMode)

		t.Log("init函数执行测试通过")
	})
}

// 测试包级别的私有函数（通过公开函数间接测试）
func TestPrivateFunctions(t *testing.T) {
	t.Run("ValidateInputThroughProcessInput", func(t *testing.T) {
		// 通过ProcessInput间接测试validateInput

		// 有效输入
		_, err := ProcessInput("valid")
		if err != nil {
			t.Errorf("ProcessInput with valid input should succeed: %v", err)
		}

		// 无效输入
		_, err = ProcessInput("")
		if err == nil {
			t.Error("ProcessInput with empty input should fail")
		}

		t.Log("私有验证函数测试通过")
	})

	t.Run("ProcessDataThroughProcessInput", func(t *testing.T) {
		// 通过ProcessInput间接测试processData

		testInput := "test data"
		result, err := ProcessInput(testInput)
		if err != nil {
			t.Errorf("ProcessInput failed: %v", err)
		}

		if string(result) != testInput {
			t.Errorf("processData result: got %s, want %s", result, testInput)
		}

		t.Log("私有处理函数测试通过")
	})
}

// 基准测试
func BenchmarkProcessInput(b *testing.B) {
	testInput := "benchmark test data"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ProcessInput(testInput)
	}
}

func BenchmarkProcessorProcess(b *testing.B) {
	processor := NewProcessor("benchmark-config")
	testData := []byte("benchmark test data")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = processor.Process(testData)
	}
}

func BenchmarkProcessorValidate(b *testing.B) {
	processor := NewProcessor("benchmark-config")
	testInput := "benchmark test input"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processor.Validate(testInput)
	}
}

// 示例测试
func ExampleNewPackageVisibility() {
	pv := NewPackageVisibility("example", 42)
	fmt.Printf("Public: %s, Exported: %d", pv.PublicField, pv.ExportedValue)
	// Output: Public: example, Exported: 42
}

func ExampleProcessInput() {
	result, err := ProcessInput("hello")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Result: %s", result)
	// Output: Result: hello
}
