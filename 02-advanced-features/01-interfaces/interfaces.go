package interfaces

import (
	"fmt"
	"math"
	"sort"
)

// 基本接口示例
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 实现Shape接口的结构体
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 使用接口的函数
func PrintShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func CalculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// 空接口示例
func PrintAnything(value interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", value, value)
}

// 类型断言示例
func DescribeValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s (长度: %d)\n", v, len(v))
	case bool:
		fmt.Printf("布尔值: %t\n", v)
	case Shape:
		fmt.Printf("形状 - 面积: %.2f\n", v.Area())
	case []int:
		fmt.Printf("整数切片: %v (长度: %d)\n", v, len(v))
	default:
		fmt.Printf("未知类型: %T, 值: %v\n", v, v)
	}
}

// 多接口组合
type Drawable interface {
	Draw() string
}

type Movable interface {
	Move(x, y float64)
}

type GameObject interface {
	Drawable
	Movable
	GetPosition() (float64, float64)
}

// 实现组合接口
type Player struct {
	Name string
	X, Y float64
}

func (p *Player) Draw() string {
	return fmt.Sprintf("玩家 %s 在 (%.1f, %.1f)", p.Name, p.X, p.Y)
}

func (p *Player) Move(x, y float64) {
	p.X += x
	p.Y += y
}

func (p Player) GetPosition() (float64, float64) {
	return p.X, p.Y
}

// 接口作为参数和返回值
type Formatter interface {
	Format(value interface{}) string
}

type JSONFormatter struct{}

func (j JSONFormatter) Format(value interface{}) string {
	return fmt.Sprintf(`{"value": "%v", "type": "%T"}`, value, value)
}

type XMLFormatter struct{}

func (x XMLFormatter) Format(value interface{}) string {
	return fmt.Sprintf(`<data type="%T">%v</data>`, value, value)
}

func FormatData(data interface{}, formatter Formatter) string {
	return formatter.Format(data)
}

// 接口的实际应用示例
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[CONSOLE] %s\n", message)
}

type FileLogger struct {
	Filename string
}

func (f FileLogger) Log(message string) {
	fmt.Printf("[FILE:%s] %s\n", f.Filename, message)
}

type Service struct {
	logger Logger
}

func NewService(logger Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) DoWork() {
	s.logger.Log("开始工作")
	// 模拟工作
	s.logger.Log("工作完成")
}

// 排序接口示例
type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

// 接口示例函数
func InterfaceExamples() {
	fmt.Println("=== 接口示例 ===")

	// 基本接口使用
	fmt.Println("\n🔹 基本接口使用")
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	fmt.Printf("矩形: ")
	PrintShapeInfo(rect)
	fmt.Printf("圆形: ")
	PrintShapeInfo(circle)

	shapes := []Shape{rect, circle}
	totalArea := CalculateTotalArea(shapes)
	fmt.Printf("总面积: %.2f\n", totalArea)

	// 空接口
	fmt.Println("\n🔹 空接口示例")
	values := []interface{}{
		42,
		"Hello",
		true,
		[]int{1, 2, 3},
		rect,
	}

	for _, value := range values {
		PrintAnything(value)
		DescribeValue(value)
		fmt.Println()
	}

	// 接口组合
	fmt.Println("\n🔹 接口组合示例")
	player := &Player{Name: "英雄", X: 0, Y: 0}
	fmt.Println(player.Draw())

	player.Move(10, 5)
	fmt.Println(player.Draw())

	x, y := player.GetPosition()
	fmt.Printf("当前位置: (%.1f, %.1f)\n", x, y)

	// 格式化器示例
	fmt.Println("\n🔹 格式化器示例")
	data := "测试数据"
	jsonFormatter := JSONFormatter{}
	xmlFormatter := XMLFormatter{}

	fmt.Println("JSON格式:", FormatData(data, jsonFormatter))
	fmt.Println("XML格式:", FormatData(data, xmlFormatter))

	// 依赖注入示例
	fmt.Println("\n🔹 依赖注入示例")
	consoleService := NewService(ConsoleLogger{})
	consoleService.DoWork()

	fileService := NewService(FileLogger{Filename: "app.log"})
	fileService.DoWork()

	// 排序示例
	fmt.Println("\n🔹 排序示例")
	people := []Person{
		{"张三", 25},
		{"李四", 30},
		{"王五", 20},
		{"赵六", 35},
	}

	fmt.Printf("原始顺序: %v\n", people)

	sort.Sort(ByAge(people))
	fmt.Printf("按年龄排序: %v\n", people)

	sort.Sort(ByName(people))
	fmt.Printf("按姓名排序: %v\n", people)
}

// 类型断言和类型开关
func TypeAssertionExamples() {
	fmt.Println("\n=== 类型断言示例 ===")

	var value interface{} = "Hello, World!"

	// 安全的类型断言
	if str, ok := value.(string); ok {
		fmt.Printf("字符串值: %s\n", str)
	} else {
		fmt.Println("不是字符串类型")
	}

	// 不安全的类型断言（可能panic）
	str := value.(string)
	fmt.Printf("直接断言: %s\n", str)

	// 类型开关
	fmt.Println("\n🔹 类型开关示例")
	values := []interface{}{
		42,
		"Go语言",
		3.14,
		true,
		[]string{"a", "b", "c"},
		map[string]int{"key": 100},
	}

	for i, v := range values {
		fmt.Printf("值 %d: ", i+1)
		switch val := v.(type) {
		case int:
			fmt.Printf("整数 %d\n", val)
		case string:
			fmt.Printf("字符串 '%s' (长度: %d)\n", val, len(val))
		case float64:
			fmt.Printf("浮点数 %.2f\n", val)
		case bool:
			fmt.Printf("布尔值 %t\n", val)
		case []string:
			fmt.Printf("字符串切片 %v\n", val)
		case map[string]int:
			fmt.Printf("映射 %v\n", val)
		default:
			fmt.Printf("未知类型 %T: %v\n", val, val)
		}
	}
}

// 接口嵌套相关类型定义
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type FileHandler struct {
	content string
}

func (f *FileHandler) Read() string {
	return f.content
}

func (f *FileHandler) Write(data string) {
	f.content = data
}

// 接口的高级用法
func AdvancedInterfaceUsage() {
	fmt.Println("\n=== 接口高级用法 ===")

	// 接口嵌套
	fmt.Println("\n🔹 接口嵌套")
	var rw ReadWriter = &FileHandler{}
	rw.Write("Hello, Interface!")
	fmt.Printf("读取内容: %s\n", rw.Read())

	// 接口值的比较
	fmt.Println("\n🔹 接口值比较")
	var shape1 Shape = Rectangle{Width: 5, Height: 3}
	var shape2 Shape = Rectangle{Width: 5, Height: 3}
	var shape3 Shape = Circle{Radius: 2}

	fmt.Printf("shape1 == shape2: %t\n", shape1 == shape2)
	fmt.Printf("shape1 == shape3: %t\n", shape1 == shape3)

	// nil接口
	fmt.Println("\n🔹 nil接口")
	var nilShape Shape
	fmt.Printf("nil接口: %v\n", nilShape)
	fmt.Printf("nil接口是否为nil: %t\n", nilShape == nil)

	// 接口包含nil指针
	var nilRect *Rectangle
	var shapeWithNilPointer Shape = nilRect
	fmt.Printf("包含nil指针的接口: %v\n", shapeWithNilPointer)
	fmt.Printf("包含nil指针的接口是否为nil: %t\n", shapeWithNilPointer == nil)
}

// 实际应用：策略模式
type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct {
	CardNumber string
}

func (c CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用信用卡 %s 支付 %.2f 元", c.CardNumber, amount)
}

type AlipayPayment struct {
	Account string
}

func (a AlipayPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用支付宝账户 %s 支付 %.2f 元", a.Account, amount)
}

type WeChatPayment struct {
	Phone string
}

func (w WeChatPayment) Pay(amount float64) string {
	return fmt.Sprintf("使用微信 %s 支付 %.2f 元", w.Phone, amount)
}

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentProcessor) ProcessPayment(amount float64) string {
	if p.strategy == nil {
		return "未设置支付方式"
	}
	return p.strategy.Pay(amount)
}

func StrategyPatternExample() {
	fmt.Println("\n=== 策略模式示例 ===")

	processor := &PaymentProcessor{}
	amount := 100.50

	// 信用卡支付
	processor.SetStrategy(CreditCardPayment{CardNumber: "**** **** **** 1234"})
	fmt.Println(processor.ProcessPayment(amount))

	// 支付宝支付
	processor.SetStrategy(AlipayPayment{Account: "user@example.com"})
	fmt.Println(processor.ProcessPayment(amount))

	// 微信支付
	processor.SetStrategy(WeChatPayment{Phone: "138****8888"})
	fmt.Println(processor.ProcessPayment(amount))
}

// 实际应用：观察者模式
type Observer interface {
	Update(message string)
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(message string)
}

type NewsAgency struct {
	observers []Observer
}

func (n *NewsAgency) Attach(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NewsAgency) Detach(observer Observer) {
	for i, obs := range n.observers {
		if obs == observer {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NewsAgency) Notify(message string) {
	for _, observer := range n.observers {
		observer.Update(message)
	}
}

type NewsChannel struct {
	Name string
}

func (nc NewsChannel) Update(message string) {
	fmt.Printf("[%s] 收到新闻: %s\n", nc.Name, message)
}

func ObserverPatternExample() {
	fmt.Println("\n=== 观察者模式示例 ===")

	agency := &NewsAgency{}

	cctv := NewsChannel{Name: "CCTV"}
	xinhua := NewsChannel{Name: "新华社"}
	peoples := NewsChannel{Name: "人民日报"}

	agency.Attach(cctv)
	agency.Attach(xinhua)
	agency.Attach(peoples)

	agency.Notify("重要新闻：Go语言发布新版本")

	// 移除一个观察者
	agency.Detach(xinhua)
	fmt.Println("\n新华社取消订阅后:")
	agency.Notify("科技新闻：人工智能新突破")
}
