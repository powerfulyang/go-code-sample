package interfaces

import (
	"math"
	"sort"
	"testing"
)

func TestBasicInterfaces(t *testing.T) {
	t.Run("ShapeInterface", func(t *testing.T) {
		rect := Rectangle{Width: 5, Height: 3}
		circle := Circle{Radius: 4}

		// 测试矩形
		rectArea := rect.Area()
		rectPerimeter := rect.Perimeter()
		expectedRectArea := 15.0
		expectedRectPerimeter := 16.0

		if rectArea != expectedRectArea {
			t.Errorf("矩形面积错误: 期望 %.2f, 实际 %.2f", expectedRectArea, rectArea)
		}
		if rectPerimeter != expectedRectPerimeter {
			t.Errorf("矩形周长错误: 期望 %.2f, 实际 %.2f", expectedRectPerimeter, rectPerimeter)
		}

		// 测试圆形
		circleArea := circle.Area()
		circlePerimeter := circle.Perimeter()
		expectedCircleArea := math.Pi * 16         // π * r²
		expectedCirclePerimeter := 2 * math.Pi * 4 // 2πr

		if math.Abs(circleArea-expectedCircleArea) > 0.01 {
			t.Errorf("圆形面积错误: 期望 %.2f, 实际 %.2f", expectedCircleArea, circleArea)
		}
		if math.Abs(circlePerimeter-expectedCirclePerimeter) > 0.01 {
			t.Errorf("圆形周长错误: 期望 %.2f, 实际 %.2f", expectedCirclePerimeter, circlePerimeter)
		}

		t.Logf("矩形 - 面积: %.2f, 周长: %.2f", rectArea, rectPerimeter)
		t.Logf("圆形 - 面积: %.2f, 周长: %.2f", circleArea, circlePerimeter)
	})

	t.Run("InterfaceAsParameter", func(t *testing.T) {
		shapes := []Shape{
			Rectangle{Width: 4, Height: 6},
			Circle{Radius: 3},
			Rectangle{Width: 2, Height: 8},
		}

		totalArea := CalculateTotalArea(shapes)
		expectedTotal := 24.0 + (math.Pi * 9) + 16.0 // 4*6 + π*3² + 2*8

		if math.Abs(totalArea-expectedTotal) > 0.01 {
			t.Errorf("总面积计算错误: 期望 %.2f, 实际 %.2f", expectedTotal, totalArea)
		}

		t.Logf("形状总面积: %.2f", totalArea)
	})
}

func TestEmptyInterface(t *testing.T) {
	t.Run("TypeAssertion", func(t *testing.T) {
		var value interface{} = "Hello, Go!"

		// 安全的类型断言
		if str, ok := value.(string); ok {
			if str != "Hello, Go!" {
				t.Errorf("类型断言结果错误: 期望 'Hello, Go!', 实际 '%s'", str)
			}
			t.Logf("成功断言为字符串: %s", str)
		} else {
			t.Error("应该能够断言为字符串")
		}

		// 错误的类型断言
		if _, ok := value.(int); ok {
			t.Error("不应该能够断言为整数")
		} else {
			t.Log("正确地识别出不是整数类型")
		}
	})

	t.Run("TypeSwitch", func(t *testing.T) {
		testCases := []struct {
			value    interface{}
			expected string
		}{
			{42, "int"},
			{"hello", "string"},
			{3.14, "float64"},
			{true, "bool"},
			{[]int{1, 2, 3}, "[]int"},
		}

		for _, tc := range testCases {
			var result string
			switch tc.value.(type) {
			case int:
				result = "int"
			case string:
				result = "string"
			case float64:
				result = "float64"
			case bool:
				result = "bool"
			case []int:
				result = "[]int"
			default:
				result = "unknown"
			}

			if result != tc.expected {
				t.Errorf("类型开关错误: 值 %v, 期望 %s, 实际 %s", tc.value, tc.expected, result)
			}
			t.Logf("值 %v 的类型: %s", tc.value, result)
		}
	})
}

func TestInterfaceComposition(t *testing.T) {
	t.Run("GameObjectInterface", func(t *testing.T) {
		player := &Player{Name: "测试玩家", X: 0, Y: 0}

		// 测试初始位置
		x, y := player.GetPosition()
		if x != 0 || y != 0 {
			t.Errorf("初始位置错误: 期望 (0, 0), 实际 (%.1f, %.1f)", x, y)
		}

		// 测试移动
		player.Move(10, 5)
		x, y = player.GetPosition()
		if x != 10 || y != 5 {
			t.Errorf("移动后位置错误: 期望 (10, 5), 实际 (%.1f, %.1f)", x, y)
		}

		// 测试绘制
		drawResult := player.Draw()
		expected := "玩家 测试玩家 在 (10.0, 5.0)"
		if drawResult != expected {
			t.Errorf("绘制结果错误: 期望 '%s', 实际 '%s'", expected, drawResult)
		}

		t.Logf("玩家状态: %s", drawResult)
	})
}

func TestFormatterInterface(t *testing.T) {
	t.Run("DifferentFormatters", func(t *testing.T) {
		data := "测试数据"
		jsonFormatter := JSONFormatter{}
		xmlFormatter := XMLFormatter{}

		jsonResult := FormatData(data, jsonFormatter)
		xmlResult := FormatData(data, xmlFormatter)

		t.Logf("JSON格式: %s", jsonResult)
		t.Logf("XML格式: %s", xmlResult)

		// 验证格式包含预期内容
		if !contains(jsonResult, "测试数据") {
			t.Error("JSON格式应该包含原始数据")
		}
		if !contains(xmlResult, "测试数据") {
			t.Error("XML格式应该包含原始数据")
		}
	})
}

func TestLoggerInterface(t *testing.T) {
	t.Run("DependencyInjection", func(t *testing.T) {
		// 创建模拟日志记录器
		mockLogger := &MockLogger{}
		service := NewService(mockLogger)

		service.DoWork()

		// 验证日志记录
		expectedLogs := []string{"开始工作", "工作完成"}
		if len(mockLogger.logs) != len(expectedLogs) {
			t.Errorf("日志数量错误: 期望 %d, 实际 %d", len(expectedLogs), len(mockLogger.logs))
		}

		for i, expectedLog := range expectedLogs {
			if i < len(mockLogger.logs) && mockLogger.logs[i] != expectedLog {
				t.Errorf("日志 %d 错误: 期望 '%s', 实际 '%s'", i, expectedLog, mockLogger.logs[i])
			}
		}

		t.Logf("记录的日志: %v", mockLogger.logs)
	})
}

// 模拟日志记录器用于测试
type MockLogger struct {
	logs []string
}

func (m *MockLogger) Log(message string) {
	m.logs = append(m.logs, message)
}

func TestSortInterface(t *testing.T) {
	t.Run("SortByAge", func(t *testing.T) {
		people := []Person{
			{"张三", 25},
			{"李四", 30},
			{"王五", 20},
			{"赵六", 35},
		}

		sort.Sort(ByAge(people))

		// 验证按年龄排序
		expectedAges := []int{20, 25, 30, 35}
		for i, person := range people {
			if person.Age != expectedAges[i] {
				t.Errorf("排序错误: 位置 %d, 期望年龄 %d, 实际年龄 %d", i, expectedAges[i], person.Age)
			}
		}

		t.Logf("按年龄排序后: %v", people)
	})

	t.Run("SortByName", func(t *testing.T) {
		people := []Person{
			{"张三", 25},
			{"李四", 30},
			{"王五", 20},
			{"赵六", 35},
		}

		sort.Sort(ByName(people))

		// 验证按姓名排序（中文按Unicode排序）
		for i := 1; i < len(people); i++ {
			if people[i-1].Name > people[i].Name {
				t.Errorf("姓名排序错误: %s 应该在 %s 之前", people[i].Name, people[i-1].Name)
			}
		}

		t.Logf("按姓名排序后: %v", people)
	})
}

func TestPaymentStrategy(t *testing.T) {
	t.Run("StrategyPattern", func(t *testing.T) {
		processor := &PaymentProcessor{}
		amount := 100.50

		strategies := []struct {
			name     string
			strategy PaymentStrategy
		}{
			{"信用卡", CreditCardPayment{CardNumber: "**** 1234"}},
			{"支付宝", AlipayPayment{Account: "test@example.com"}},
			{"微信", WeChatPayment{Phone: "138****8888"}},
		}

		for _, s := range strategies {
			processor.SetStrategy(s.strategy)
			result := processor.ProcessPayment(amount)

			if result == "" {
				t.Errorf("%s 支付策略返回空结果", s.name)
			}

			t.Logf("%s 支付结果: %s", s.name, result)
		}

		// 测试未设置策略的情况
		processor.SetStrategy(nil)
		result := processor.ProcessPayment(amount)
		expected := "未设置支付方式"
		if result != expected {
			t.Errorf("未设置策略时应该返回 '%s', 实际返回 '%s'", expected, result)
		}
	})
}

func TestObserverPattern(t *testing.T) {
	t.Run("NewsAgencyObserver", func(t *testing.T) {
		agency := &NewsAgency{}

		// 创建模拟观察者
		observer1 := &MockObserver{name: "观察者1"}
		observer2 := &MockObserver{name: "观察者2"}
		observer3 := &MockObserver{name: "观察者3"}

		// 添加观察者
		agency.Attach(observer1)
		agency.Attach(observer2)
		agency.Attach(observer3)

		// 发送通知
		message := "测试新闻"
		agency.Notify(message)

		// 验证所有观察者都收到了消息
		observers := []*MockObserver{observer1, observer2, observer3}
		for i, obs := range observers {
			if len(obs.messages) != 1 {
				t.Errorf("观察者 %d 应该收到 1 条消息, 实际收到 %d 条", i+1, len(obs.messages))
			}
			if len(obs.messages) > 0 && obs.messages[0] != message {
				t.Errorf("观察者 %d 收到错误消息: 期望 '%s', 实际 '%s'", i+1, message, obs.messages[0])
			}
		}

		// 移除一个观察者
		agency.Detach(observer2)

		// 再次发送通知
		message2 := "第二条新闻"
		agency.Notify(message2)

		// 验证被移除的观察者没有收到第二条消息
		if len(observer2.messages) != 1 {
			t.Errorf("被移除的观察者不应该收到新消息, 实际收到 %d 条", len(observer2.messages))
		}

		// 验证其他观察者收到了第二条消息
		if len(observer1.messages) != 2 || len(observer3.messages) != 2 {
			t.Error("剩余观察者应该收到两条消息")
		}

		t.Logf("观察者1收到的消息: %v", observer1.messages)
		t.Logf("观察者2收到的消息: %v", observer2.messages)
		t.Logf("观察者3收到的消息: %v", observer3.messages)
	})
}

// 模拟观察者用于测试
type MockObserver struct {
	name     string
	messages []string
}

func (m *MockObserver) Update(message string) {
	m.messages = append(m.messages, message)
}

func TestInterfaceComparison(t *testing.T) {
	t.Run("InterfaceEquality", func(t *testing.T) {
		rect1 := Rectangle{Width: 5, Height: 3}
		rect2 := Rectangle{Width: 5, Height: 3}
		rect3 := Rectangle{Width: 4, Height: 3}

		var shape1 Shape = rect1
		var shape2 Shape = rect2
		var shape3 Shape = rect3

		// 相同值的接口应该相等
		if shape1 != shape2 {
			t.Error("相同值的接口应该相等")
		}

		// 不同值的接口不应该相等
		if shape1 == shape3 {
			t.Error("不同值的接口不应该相等")
		}

		t.Logf("shape1 == shape2: %t", shape1 == shape2)
		t.Logf("shape1 == shape3: %t", shape1 == shape3)
	})

	t.Run("NilInterface", func(t *testing.T) {
		var nilShape Shape
		if nilShape != nil {
			t.Error("nil接口应该等于nil")
		}

		var nilRect *Rectangle
		var shapeWithNilPointer Shape = nilRect
		if shapeWithNilPointer == nil {
			t.Error("包含nil指针的接口不应该等于nil")
		}

		t.Logf("nil接口: %v", nilShape)
		t.Logf("包含nil指针的接口: %v", shapeWithNilPointer)
	})
}

// 辅助函数
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] == substr ||
		len(s) >= len(substr) && s[:len(substr)] == substr ||
		len(s) > len(substr) && findSubstring(s, substr)
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// 基准测试
func BenchmarkInterfaceCall(b *testing.B) {
	shape := Rectangle{Width: 5, Height: 3}
	for i := 0; i < b.N; i++ {
		_ = shape.Area()
	}
}

func BenchmarkInterfaceCallThroughInterface(b *testing.B) {
	var shape Shape = Rectangle{Width: 5, Height: 3}
	for i := 0; i < b.N; i++ {
		_ = shape.Area()
	}
}

func BenchmarkTypeAssertion(b *testing.B) {
	var value interface{} = "test string"
	for i := 0; i < b.N; i++ {
		if _, ok := value.(string); ok {
			// 类型断言成功
		}
	}
}
