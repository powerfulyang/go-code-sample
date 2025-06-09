package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

// 🎓 学习导向的测试 - 通过测试学习Go接口

// TestLearnInterfaceBasics 学习接口基础概念
func TestLearnInterfaceBasics(t *testing.T) {
	t.Log("🎯 学习目标: 理解Go接口的基本概念和设计哲学")
	t.Log("📚 本测试将教您: 接口定义、隐式实现、多态性")

	t.Run("学习接口的隐式实现", func(t *testing.T) {
		t.Log("📖 知识点: Go接口是隐式实现的，无需显式声明")

		// 🔍 探索: 不同类型实现相同接口
		var shapes []Shape

		// 创建不同的形状
		rect := Rectangle{Width: 5, Height: 3}
		circle := Circle{Radius: 4}

		// 添加到接口切片中
		shapes = append(shapes, rect)
		shapes = append(shapes, circle)

		t.Log("🔍 多态性演示:")
		for i, shape := range shapes {
			area := shape.Area()
			perimeter := shape.Perimeter()

			// 使用类型断言获取具体类型信息
			switch s := shape.(type) {
			case Rectangle:
				t.Logf("   形状%d: 矩形 %.1f×%.1f, 面积=%.2f, 周长=%.2f",
					i+1, s.Width, s.Height, area, perimeter)
			case Circle:
				t.Logf("   形状%d: 圆形 半径=%.1f, 面积=%.2f, 周长=%.2f",
					i+1, s.Radius, area, perimeter)
			}
		}

		// ✅ 验证多态性
		if len(shapes) != 2 {
			t.Errorf("❌ 应该有2个形状，实际有%d个", len(shapes))
		}

		rectArea := shapes[0].Area()
		if rectArea != 15.0 {
			t.Errorf("❌ 矩形面积计算错误: 期望15.0，得到%.1f", rectArea)
		}

		t.Log("✅ 很好！您理解了接口的隐式实现和多态性")

		// 💡 学习提示
		t.Log("💡 核心概念: 接口定义行为，结构体提供实现")
		t.Log("💡 设计哲学: '接受接口，返回结构体'")
		t.Log("💡 多态性: 同一接口可以有多种实现")
	})

	t.Run("学习空接口interface{}", func(t *testing.T) {
		t.Log("📖 知识点: interface{}可以持有任何类型的值")

		// 🔍 探索: 空接口的使用
		var anything interface{}

		values := []interface{}{
			42,
			"Hello",
			3.14,
			true,
			[]int{1, 2, 3},
			map[string]int{"a": 1},
			Rectangle{Width: 2, Height: 3},
		}

		t.Log("🔍 空接口可以存储任何类型:")
		for i, value := range values {
			anything = value
			t.Logf("   值%d: %v (类型: %T)", i+1, anything, anything)

			// 演示类型断言
			switch v := anything.(type) {
			case int:
				t.Logf("     → 这是一个整数: %d", v)
			case string:
				t.Logf("     → 这是一个字符串: %s", v)
			case Rectangle:
				t.Logf("     → 这是一个矩形: %.1f×%.1f", v.Width, v.Height)
			default:
				t.Logf("     → 其他类型: %T", v)
			}
		}

		// ✅ 验证空接口
		anything = "test"
		if str, ok := anything.(string); !ok || str != "test" {
			t.Error("❌ 空接口类型断言失败")
		} else {
			t.Log("✅ 很好！您理解了空接口和类型断言")
		}

		// 💡 学习提示
		t.Log("💡 使用场景: JSON解析、通用容器、反射")
		t.Log("💡 性能考虑: 空接口会有装箱开销")
		t.Log("💡 类型安全: 使用时需要类型断言")
	})
}

// TestLearnInterfaceComposition 学习接口组合
func TestLearnInterfaceComposition(t *testing.T) {
	t.Log("🎯 学习目标: 掌握接口组合的设计模式")
	t.Log("📚 本测试将教您: 接口嵌入、组合设计、依赖注入")

	t.Run("学习接口嵌入", func(t *testing.T) {
		t.Log("📖 知识点: 接口可以嵌入其他接口，形成更大的接口")

		// 🔍 探索: 组合接口的使用
		player := &Player{
			Name: "勇敢的冒险者",
			X:    10,
			Y:    20,
		}

		// 测试作为Movable接口
		var movable Movable = player
		t.Log("🔍 作为Movable接口使用:")
		x, y := player.GetPosition()
		t.Logf("   当前位置: (%.1f, %.1f)", x, y)

		movable.Move(5, -3)
		x, y = player.GetPosition()
		t.Logf("   移动后位置: (%.1f, %.1f)", x, y)

		// 测试作为Drawable接口
		var drawable Drawable = player
		t.Log("🔍 作为Drawable接口使用:")
		drawResult := drawable.Draw()
		t.Logf("   绘制结果: %s", drawResult)

		// 测试作为GameObject接口（组合接口）
		var gameObject GameObject = player
		t.Log("🔍 作为GameObject接口使用:")
		gameObject.Move(0, 0)
		result := gameObject.Draw()
		t.Logf("   游戏对象: %s", result)

		// ✅ 验证接口组合
		if x != 15 || y != 17 {
			t.Errorf("❌ 位置计算错误: 期望(15, 17)，得到(%.1f, %.1f)", x, y)
		}

		if !strings.Contains(drawResult, "勇敢的冒险者") {
			t.Error("❌ 绘制结果应该包含玩家名称")
		}

		t.Log("✅ 很好！您理解了接口组合")

		// 💡 学习提示
		t.Log("💡 组合优势: 小接口组合成大接口，更灵活")
		t.Log("💡 设计原则: 接口隔离原则 - 客户端不应依赖不需要的方法")
		t.Log("💡 实际应用: io.ReadWriter = io.Reader + io.Writer")
	})

	t.Run("学习依赖注入", func(t *testing.T) {
		t.Log("📖 知识点: 通过接口实现依赖注入，提高代码可测试性")

		// 🔍 探索: 依赖注入模式
		// 创建不同的日志实现
		memoryLogger := &MemoryLogger{}
		consoleLogger := &ConsoleLoggerImpl{}

		// 使用内存日志器
		service1 := NewService(memoryLogger)
		service1.DoWork()

		t.Log("🔍 使用内存日志器:")
		t.Logf("   记录的日志: %v", memoryLogger.GetLogs())

		// 使用控制台日志器
		service2 := NewService(consoleLogger)
		service2.DoWork()

		t.Log("🔍 使用控制台日志器:")
		t.Log("   日志已输出到控制台")

		// ✅ 验证依赖注入
		logs := memoryLogger.GetLogs()
		expectedLogs := []string{"开始工作", "工作完成"}

		if len(logs) != len(expectedLogs) {
			t.Errorf("❌ 日志数量错误: 期望%d，得到%d", len(expectedLogs), len(logs))
		}

		for i, expected := range expectedLogs {
			if i < len(logs) && logs[i] != expected {
				t.Errorf("❌ 日志%d错误: 期望'%s'，得到'%s'", i, expected, logs[i])
			}
		}

		t.Log("✅ 很好！您理解了依赖注入模式")

		// 💡 学习提示
		t.Log("💡 测试优势: 可以注入模拟对象进行单元测试")
		t.Log("💡 解耦效果: 业务逻辑与具体实现解耦")
		t.Log("💡 配置灵活: 运行时可以切换不同实现")
	})
}

// TestLearnStandardInterfaces 学习标准库接口
func TestLearnStandardInterfaces(t *testing.T) {
	t.Log("🎯 学习目标: 掌握Go标准库中的重要接口")
	t.Log("📚 本测试将教您: io.Reader/Writer、fmt.Stringer、error接口")

	t.Run("学习io.Reader和io.Writer", func(t *testing.T) {
		t.Log("📖 知识点: io.Reader和io.Writer是Go中最重要的接口")

		// 🔍 探索: 使用strings.Reader实现io.Reader
		data := "Hello, Go interfaces!"
		reader := strings.NewReader(data)

		t.Log("🔍 使用io.Reader读取数据:")
		buffer := make([]byte, 10)
		totalRead := 0

		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				chunk := string(buffer[:n])
				t.Logf("   读取了%d字节: '%s'", n, chunk)
				totalRead += n
			}
			if err == io.EOF {
				t.Log("   读取完成")
				break
			}
			if err != nil {
				t.Fatalf("   读取错误: %v", err)
			}
		}

		// 🔍 探索: 使用bytes.Buffer实现io.Writer
		var writer bytes.Buffer

		t.Log("🔍 使用io.Writer写入数据:")
		messages := []string{"Hello", " ", "World", "!"}

		for _, msg := range messages {
			n, err := writer.Write([]byte(msg))
			if err != nil {
				t.Fatalf("   写入错误: %v", err)
			}
			t.Logf("   写入了%d字节: '%s'", n, msg)
		}

		result := writer.String()
		t.Logf("   最终结果: '%s'", result)

		// ✅ 验证io接口
		if totalRead != len(data) {
			t.Errorf("❌ 读取字节数错误: 期望%d，得到%d", len(data), totalRead)
		}

		if result != "Hello World!" {
			t.Errorf("❌ 写入结果错误: 期望'Hello World!'，得到'%s'", result)
		}

		t.Log("✅ 很好！您理解了io.Reader和io.Writer接口")

		// 💡 学习提示
		t.Log("💡 设计精髓: 简单的接口，强大的组合能力")
		t.Log("💡 实际应用: 文件操作、网络通信、数据处理")
		t.Log("💡 组合示例: io.Copy(writer, reader)")
	})

	t.Run("学习fmt.Stringer接口", func(t *testing.T) {
		t.Log("📖 知识点: fmt.Stringer接口定义了字符串表示方法")

		// 🔍 探索: 实现Stringer接口
		person := PersonWithString{Name: "张三", Age: 25}

		// 直接调用String方法
		str1 := person.String()
		t.Logf("🔍 直接调用String(): %s", str1)

		// 通过fmt包调用（会自动调用String方法）
		str2 := fmt.Sprintf("%s", person)
		t.Logf("🔍 通过fmt.Sprintf: %s", str2)

		// 通过fmt.Println（也会调用String方法）
		t.Logf("🔍 通过fmt包: %v", person)

		// ✅ 验证Stringer接口
		expected := "张三 (25岁)"
		if str1 != expected || str2 != expected {
			t.Errorf("❌ String方法结果错误: 期望'%s'", expected)
		} else {
			t.Log("✅ 很好！您理解了fmt.Stringer接口")
		}

		// 💡 学习提示
		t.Log("💡 自动调用: fmt包会自动检查并调用String()方法")
		t.Log("💡 调试友好: 实现Stringer让对象更易于调试")
		t.Log("💡 性能考虑: String()方法应该高效且无副作用")
	})

	t.Run("学习error接口", func(t *testing.T) {
		t.Log("📖 知识点: error是Go中错误处理的核心接口")

		// 🔍 探索: 自定义错误类型
		customErr := &CustomError{
			Code:    404,
			Message: "资源未找到",
		}

		// 作为error接口使用
		var err error = customErr
		t.Logf("🔍 自定义错误: %v", err)

		// 类型断言获取详细信息
		if ce, ok := err.(*CustomError); ok {
			t.Logf("🔍 错误详情: 代码=%d, 消息=%s", ce.Code, ce.Message)
		}

		// 🔍 探索: 错误包装
		wrappedErr := fmt.Errorf("操作失败: %w", customErr)
		t.Logf("🔍 包装错误: %v", wrappedErr)

		// ✅ 验证error接口
		if err.Error() != "错误 404: 资源未找到" {
			t.Errorf("❌ 错误消息格式错误: %s", err.Error())
		} else {
			t.Log("✅ 很好！您理解了error接口")
		}

		// 💡 学习提示
		t.Log("💡 简单设计: error接口只有一个Error()方法")
		t.Log("💡 错误处理: Go推荐显式错误处理")
		t.Log("💡 错误包装: 使用fmt.Errorf和%w动词包装错误")
	})
}

// 辅助类型和实现

// MemoryLogger 内存日志器（用于测试）
type MemoryLogger struct {
	logs []string
}

func (m *MemoryLogger) Log(message string) {
	m.logs = append(m.logs, message)
}

func (m *MemoryLogger) GetLogs() []string {
	return m.logs
}

// ConsoleLoggerImpl 控制台日志器实现
type ConsoleLoggerImpl struct{}

func (c *ConsoleLoggerImpl) Log(message string) {
	fmt.Printf("[LOG] %s\n", message)
}

// CustomError 自定义错误类型
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("错误 %d: %s", e.Code, e.Message)
}

// PersonWithString 实现了String方法的Person类型
type PersonWithString struct {
	Name string
	Age  int
}

func (p PersonWithString) String() string {
	return fmt.Sprintf("%s (%d岁)", p.Name, p.Age)
}

// TestLearnInterfaceDesignPrinciples 学习接口设计原则
func TestLearnInterfaceDesignPrinciples(t *testing.T) {
	t.Log("🎯 学习目标: 掌握接口设计的最佳实践")
	t.Log("📚 本测试将教您: 接口隔离、依赖倒置、里氏替换")

	t.Run("学习接口隔离原则", func(t *testing.T) {
		t.Log("📖 知识点: 接口应该小而专一，客户端不应依赖不需要的方法")

		// 🔍 探索: 好的接口设计 vs 坏的接口设计

		// 好的设计：小接口
		var reader io.Reader = strings.NewReader("test data")
		var writer io.Writer = &bytes.Buffer{}

		// 客户端只需要读取功能
		data := make([]byte, 4)
		readBytes, err := reader.Read(data)

		t.Logf("🔍 好的设计 - 只使用需要的接口:")
		t.Logf("   读取了%d字节: %s", readBytes, string(data[:readBytes]))

		// 客户端只需要写入功能
		writeBytes, err := writer.Write([]byte("hello"))
		if err == nil {
			t.Logf("   写入了%d字节", writeBytes)
		}

		// ✅ 验证接口隔离
		if readBytes != 4 {
			t.Errorf("❌ 读取字节数错误: 期望4，得到%d", readBytes)
		} else if writeBytes != 5 {
			t.Errorf("❌ 写入字节数错误: 期望5，得到%d", writeBytes)
		} else {
			t.Log("✅ 很好！您理解了接口隔离原则")
		}

		// 💡 学习提示
		t.Log("💡 设计原则: 接口应该内聚，职责单一")
		t.Log("💡 实际好处: 更容易测试、更容易理解、更容易维护")
		t.Log("💡 Go哲学: 'The bigger the interface, the weaker the abstraction'")
	})

	t.Run("学习依赖倒置原则", func(t *testing.T) {
		t.Log("📖 知识点: 高层模块不应依赖低层模块，都应依赖抽象")

		// 🔍 探索: 通过接口实现依赖倒置

		// 高层模块（业务逻辑）依赖抽象（接口）
		emailService := &EmailService{}
		smsService := &SMSService{}

		// 通知管理器依赖抽象，不依赖具体实现
		notifier1 := &NotificationManager{service: emailService}
		notifier2 := &NotificationManager{service: smsService}

		// 发送通知
		result1 := notifier1.SendNotification("测试消息1")
		result2 := notifier2.SendNotification("测试消息2")

		t.Logf("🔍 依赖倒置演示:")
		t.Logf("   邮件通知: %s", result1)
		t.Logf("   短信通知: %s", result2)

		// ✅ 验证依赖倒置
		if !strings.Contains(result1, "邮件") || !strings.Contains(result2, "短信") {
			t.Error("❌ 依赖倒置实现错误")
		} else {
			t.Log("✅ 很好！您理解了依赖倒置原则")
		}

		// 💡 学习提示
		t.Log("💡 解耦效果: 高层模块不依赖具体实现")
		t.Log("💡 扩展性: 可以轻松添加新的通知方式")
		t.Log("💡 测试性: 可以注入模拟对象进行测试")
	})
}

// 通知服务接口
type NotificationService interface {
	Send(message string) string
}

// 邮件服务实现
type EmailService struct{}

func (e *EmailService) Send(message string) string {
	return fmt.Sprintf("邮件发送: %s", message)
}

// 短信服务实现
type SMSService struct{}

func (s *SMSService) Send(message string) string {
	return fmt.Sprintf("短信发送: %s", message)
}

// 通知管理器（高层模块）
type NotificationManager struct {
	service NotificationService
}

func (n *NotificationManager) SendNotification(message string) string {
	return n.service.Send(message)
}

// BenchmarkLearnInterfacePerformance 学习接口性能
func BenchmarkLearnInterfacePerformance(b *testing.B) {
	b.Log("🎯 学习目标: 了解接口调用的性能特征")

	rect := Rectangle{Width: 5, Height: 3}

	b.Run("直接方法调用", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = rect.Area()
		}
	})

	b.Run("通过接口调用", func(b *testing.B) {
		var shape Shape = rect
		for i := 0; i < b.N; i++ {
			_ = shape.Area()
		}
	})

	b.Run("空接口类型断言", func(b *testing.B) {
		var anything interface{} = rect
		for i := 0; i < b.N; i++ {
			if shape, ok := anything.(Shape); ok {
				_ = shape.Area()
			}
		}
	})
}

// SimpleGreeter 简单的问候接口
type SimpleGreeter interface {
	Greet(name string) string
}

// EnglishGreeter 英文问候实现
type EnglishGreeter struct{}

func (e EnglishGreeter) Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Example_learnInterfaceBasics 接口基础示例
func Example_learnInterfaceBasics() {
	// 使用接口
	var greeter SimpleGreeter = EnglishGreeter{}
	message := greeter.Greet("Go")
	fmt.Println(message)

	// Output:
	// Hello, Go!
}
