package main

import (
	"testing"
)

func TestVariableDeclaration(t *testing.T) {
	t.Run("BasicDeclaration", func(t *testing.T) {
		// 基本变量声明
		var name string = "张三"
		var age int = 25
		var isStudent bool = true
		var height float64 = 175.5

		t.Logf("姓名: %s", name)
		t.Logf("年龄: %d", age)
		t.Logf("是否学生: %t", isStudent)
		t.Logf("身高: %.1f cm", height)

		// 验证值
		if name != "张三" {
			t.Errorf("期望姓名为 '张三', 实际为 '%s'", name)
		}
		if age != 25 {
			t.Errorf("期望年龄为 25, 实际为 %d", age)
		}
	})

	t.Run("ShortDeclaration", func(t *testing.T) {
		// 短变量声明
		name := "李四"
		age := 30
		isWorking := true
		salary := 8500.50

		t.Logf("姓名: %s", name)
		t.Logf("年龄: %d", age)
		t.Logf("是否工作: %t", isWorking)
		t.Logf("薪资: %.2f", salary)
	})

	t.Run("MultipleDeclaration", func(t *testing.T) {
		// 多变量声明
		var (
			firstName string = "王"
			lastName  string = "五"
			fullName  string
		)

		fullName = firstName + lastName
		t.Logf("姓: %s", firstName)
		t.Logf("名: %s", lastName)
		t.Logf("全名: %s", fullName)

		// 多变量短声明
		x, y, z := 10, 20, 30
		t.Logf("x=%d, y=%d, z=%d", x, y, z)
	})

	t.Run("ZeroValues", func(t *testing.T) {
		// 零值测试
		var (
			intVar    int
			floatVar  float64
			boolVar   bool
			stringVar string
		)

		t.Logf("int 零值: %d", intVar)
		t.Logf("float64 零值: %f", floatVar)
		t.Logf("bool 零值: %t", boolVar)
		t.Logf("string 零值: '%s'", stringVar)

		// 验证零值
		if intVar != 0 {
			t.Errorf("int 零值应该为 0, 实际为 %d", intVar)
		}
		if stringVar != "" {
			t.Errorf("string 零值应该为空字符串, 实际为 '%s'", stringVar)
		}
	})
}

func TestConstants(t *testing.T) {
	t.Run("BasicConstants", func(t *testing.T) {
		// 基本常量
		const pi = 3.14159
		const greeting = "你好"
		const maxUsers = 1000
		const isDebug = true

		t.Logf("π 的值: %f", pi)
		t.Logf("问候语: %s", greeting)
		t.Logf("最大用户数: %d", maxUsers)
		t.Logf("调试模式: %t", isDebug)
	})

	t.Run("ConstantBlock", func(t *testing.T) {
		// 常量块
		const (
			StatusPending  = "pending"
			StatusApproved = "approved"
			StatusRejected = "rejected"
			MaxRetries     = 3
			TimeoutSeconds = 30
		)

		t.Logf("状态 - 待处理: %s", StatusPending)
		t.Logf("状态 - 已批准: %s", StatusApproved)
		t.Logf("状态 - 已拒绝: %s", StatusRejected)
		t.Logf("最大重试次数: %d", MaxRetries)
		t.Logf("超时时间: %d 秒", TimeoutSeconds)
	})

	t.Run("IotaConstants", func(t *testing.T) {
		// iota 常量
		const (
			Sunday    = iota // 0
			Monday           // 1
			Tuesday          // 2
			Wednesday        // 3
			Thursday         // 4
			Friday           // 5
			Saturday         // 6
		)

		t.Logf("星期日: %d", Sunday)
		t.Logf("星期一: %d", Monday)
		t.Logf("星期二: %d", Tuesday)
		t.Logf("星期三: %d", Wednesday)
		t.Logf("星期四: %d", Thursday)
		t.Logf("星期五: %d", Friday)
		t.Logf("星期六: %d", Saturday)

		// 验证 iota 值
		if Sunday != 0 || Monday != 1 || Saturday != 6 {
			t.Error("iota 常量值不正确")
		}
	})

	t.Run("IotaExpressions", func(t *testing.T) {
		// iota 表达式
		const (
			_  = iota             // 0, 忽略
			KB = 1 << (10 * iota) // 1 << 10 = 1024
			MB                    // 1 << 20 = 1048576
			GB                    // 1 << 30 = 1073741824
		)

		t.Logf("1 KB = %d 字节", KB)
		t.Logf("1 MB = %d 字节", MB)
		t.Logf("1 GB = %d 字节", GB)

		// 验证计算结果
		if KB != 1024 {
			t.Errorf("KB 应该为 1024, 实际为 %d", KB)
		}
		if MB != 1048576 {
			t.Errorf("MB 应该为 1048576, 实际为 %d", MB)
		}
	})
}

func TestVariableScope(t *testing.T) {
	t.Run("LocalScope", func(t *testing.T) {
		// 局部作用域
		name := "外部变量"
		t.Logf("外部 name: %s", name)

		{
			// 内部作用域
			name := "内部变量"
			t.Logf("内部 name: %s", name)
		}

		t.Logf("外部 name (内部作用域后): %s", name)
	})

	t.Run("VariableShadowing", func(t *testing.T) {
		// 变量遮蔽
		x := 10
		t.Logf("外部 x: %d", x)

		if true {
			x := 20 // 遮蔽外部变量
			t.Logf("if 块内 x: %d", x)
		}

		t.Logf("外部 x (if 块后): %d", x)

		for i := 0; i < 1; i++ {
			x := 30 // 遮蔽外部变量
			t.Logf("for 循环内 x: %d", x)
		}

		t.Logf("外部 x (for 循环后): %d", x)
	})
}

func TestTypeInference(t *testing.T) {
	t.Run("AutoTypeInference", func(t *testing.T) {
		// 自动类型推断
		name := "Go语言"       // string
		version := 1.21      // float64
		isStable := true     // bool
		userCount := 1000000 // int

		t.Logf("语言名称: %s (类型: %T)", name, name)
		t.Logf("版本: %g (类型: %T)", version, version)
		t.Logf("是否稳定: %t (类型: %T)", isStable, isStable)
		t.Logf("用户数量: %d (类型: %T)", userCount, userCount)
	})

	t.Run("ExplicitTypeConversion", func(t *testing.T) {
		// 显式类型转换
		var intVal int = 42
		var floatVal float64 = float64(intVal)
		var stringVal string = string(rune(65)) // ASCII 65 = 'A'

		t.Logf("整数: %d", intVal)
		t.Logf("转换为浮点数: %f", floatVal)
		t.Logf("ASCII 65 转换为字符: %s", stringVal)

		// 验证转换结果
		if floatVal != 42.0 {
			t.Errorf("类型转换失败: 期望 42.0, 实际 %f", floatVal)
		}
		if stringVal != "A" {
			t.Errorf("字符转换失败: 期望 'A', 实际 '%s'", stringVal)
		}
	})
}

func TestPracticalExamples(t *testing.T) {
	t.Run("UserProfile", func(t *testing.T) {
		// 用户档案示例
		const (
			MinAge = 18
			MaxAge = 65
		)

		userName := "张三"
		userAge := 28
		isAdult := userAge >= MinAge
		canRetire := userAge >= MaxAge

		t.Logf("用户信息:")
		t.Logf("  姓名: %s", userName)
		t.Logf("  年龄: %d", userAge)
		t.Logf("  是否成年: %t", isAdult)
		t.Logf("  可以退休: %t", canRetire)
		t.Logf("常量:")
		t.Logf("  最小年龄: %d", MinAge)
		t.Logf("  退休年龄: %d", MaxAge)
	})

	t.Run("ConfigurationSettings", func(t *testing.T) {
		// 配置设置示例
		const (
			AppName        = "学习管理系统"
			Version        = "1.0.0"
			DefaultPort    = 8080
			MaxConnections = 1000
			EnableLogging  = true
		)

		var (
			currentPort   = DefaultPort
			activeUsers   = 0
			serverRunning = false
		)

		t.Logf("应用配置:")
		t.Logf("  应用名称: %s", AppName)
		t.Logf("  版本: %s", Version)
		t.Logf("  默认端口: %d", DefaultPort)
		t.Logf("  最大连接数: %d", MaxConnections)
		t.Logf("  启用日志: %t", EnableLogging)

		t.Logf("运行时状态:")
		t.Logf("  当前端口: %d", currentPort)
		t.Logf("  活跃用户: %d", activeUsers)
		t.Logf("  服务器运行: %t", serverRunning)
	})
}
