package structs

import (
	"fmt"
	"math"
	"time"
)

// 基本结构体示例
type Person struct {
	Name string
	Age  int
	City string
}

// 带标签的结构体
type User struct {
	ID       int    `json:"id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

// 嵌套结构体
type Address struct {
	Street   string
	City     string
	Province string
	ZipCode  string
}

type Employee struct {
	Person  // 匿名字段（嵌入）
	ID      int
	Salary  float64
	Address Address // 命名字段
}

// 方法示例
type Rectangle struct {
	Width  float64
	Height float64
}

// 值接收者方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 值接收者方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 指针接收者方法
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 指针接收者方法
func (r *Rectangle) SetDimensions(width, height float64) {
	r.Width = width
	r.Height = height
}

// 银行账户示例
type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

// 存款方法
func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("存款金额必须大于0")
	}
	ba.Balance += amount
	return nil
}

// 取款方法
func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("取款金额必须大于0")
	}
	if amount > ba.Balance {
		return fmt.Errorf("余额不足")
	}
	ba.Balance -= amount
	return nil
}

// 查询余额方法
func (ba BankAccount) GetBalance() float64 {
	return ba.Balance
}

// 获取账户信息方法
func (ba BankAccount) GetAccountInfo() string {
	return fmt.Sprintf("账户: %s, 持有人: %s, 余额: %.2f",
		ba.AccountNumber, ba.HolderName, ba.Balance)
}

// 圆形结构体
type Circle struct {
	Radius float64
}

// 圆形方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// 学生结构体
type Student struct {
	ID     int
	Name   string
	Scores []int
}

// 计算平均分
func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}
	total := 0
	for _, score := range s.Scores {
		total += score
	}
	return float64(total) / float64(len(s.Scores))
}

// 添加分数
func (s *Student) AddScore(score int) {
	s.Scores = append(s.Scores, score)
}

// 获取最高分
func (s Student) HighestScore() int {
	if len(s.Scores) == 0 {
		return 0
	}
	highest := s.Scores[0]
	for _, score := range s.Scores {
		if score > highest {
			highest = score
		}
	}
	return highest
}

// 获取最低分
func (s Student) LowestScore() int {
	if len(s.Scores) == 0 {
		return 0
	}
	lowest := s.Scores[0]
	for _, score := range s.Scores {
		if score < lowest {
			lowest = score
		}
	}
	return lowest
}

// 时间相关结构体
type Event struct {
	Name      string
	StartTime time.Time
	Duration  time.Duration
}

// 获取结束时间
func (e Event) EndTime() time.Time {
	return e.StartTime.Add(e.Duration)
}

// 检查是否正在进行
func (e Event) IsOngoing(currentTime time.Time) bool {
	return currentTime.After(e.StartTime) && currentTime.Before(e.EndTime())
}

// 获取剩余时间
func (e Event) TimeRemaining(currentTime time.Time) time.Duration {
	endTime := e.EndTime()
	if currentTime.After(endTime) {
		return 0
	}
	return endTime.Sub(currentTime)
}

// 构造函数示例
func NewPerson(name string, age int, city string) *Person {
	return &Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

func NewBankAccount(accountNumber, holderName string, initialBalance float64) *BankAccount {
	return &BankAccount{
		AccountNumber: accountNumber,
		HolderName:    holderName,
		Balance:       initialBalance,
	}
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

func NewStudent(id int, name string) *Student {
	return &Student{
		ID:     id,
		Name:   name,
		Scores: make([]int, 0),
	}
}

// 结构体示例函数
func StructExamples() {
	fmt.Println("=== 结构体示例 ===")

	// 创建结构体实例
	person1 := Person{
		Name: "张三",
		Age:  25,
		City: "北京",
	}

	person2 := Person{"李四", 30, "上海"} // 按字段顺序初始化

	var person3 Person // 零值结构体
	person3.Name = "王五"
	person3.Age = 28
	person3.City = "广州"

	fmt.Printf("person1: %+v\n", person1)
	fmt.Printf("person2: %+v\n", person2)
	fmt.Printf("person3: %+v\n", person3)

	// 使用构造函数
	person4 := NewPerson("赵六", 35, "深圳")
	fmt.Printf("person4: %+v\n", *person4)

	// 嵌套结构体
	employee := Employee{
		Person: Person{
			Name: "钱七",
			Age:  32,
			City: "杭州",
		},
		ID:     1001,
		Salary: 8500.0,
		Address: Address{
			Street:   "西湖路123号",
			City:     "杭州",
			Province: "浙江",
			ZipCode:  "310000",
		},
	}

	fmt.Printf("员工信息: %+v\n", employee)
	fmt.Printf("员工姓名: %s\n", employee.Name) // 可以直接访问嵌入字段
	fmt.Printf("员工地址: %s, %s\n", employee.Address.Street, employee.Address.City)
}

// 方法示例函数
func MethodExamples() {
	fmt.Println("\n=== 方法示例 ===")

	// 矩形示例
	rect := NewRectangle(5.0, 3.0)
	fmt.Printf("矩形: 宽=%.1f, 高=%.1f\n", rect.Width, rect.Height)
	fmt.Printf("面积: %.2f\n", rect.Area())
	fmt.Printf("周长: %.2f\n", rect.Perimeter())

	// 修改矩形
	rect.Scale(2.0)
	fmt.Printf("放大2倍后: 宽=%.1f, 高=%.1f\n", rect.Width, rect.Height)
	fmt.Printf("新面积: %.2f\n", rect.Area())

	rect.SetDimensions(10.0, 8.0)
	fmt.Printf("设置新尺寸后: 宽=%.1f, 高=%.1f\n", rect.Width, rect.Height)

	// 圆形示例
	circle := Circle{Radius: 5.0}
	fmt.Printf("圆形: 半径=%.1f\n", circle.Radius)
	fmt.Printf("面积: %.2f\n", circle.Area())
	fmt.Printf("周长: %.2f\n", circle.Circumference())
	fmt.Printf("直径: %.2f\n", circle.Diameter())

	// 银行账户示例
	account := NewBankAccount("123456789", "张三", 1000.0)
	fmt.Printf("初始状态: %s\n", account.GetAccountInfo())

	// 存款
	if err := account.Deposit(500.0); err != nil {
		fmt.Printf("存款失败: %v\n", err)
	} else {
		fmt.Printf("存款后: %s\n", account.GetAccountInfo())
	}

	// 取款
	if err := account.Withdraw(200.0); err != nil {
		fmt.Printf("取款失败: %v\n", err)
	} else {
		fmt.Printf("取款后: %s\n", account.GetAccountInfo())
	}

	// 尝试取款超过余额
	if err := account.Withdraw(2000.0); err != nil {
		fmt.Printf("取款失败: %v\n", err)
	}
}

// 学生管理示例
func StudentManagementExample() {
	fmt.Println("\n=== 学生管理示例 ===")

	student := NewStudent(1001, "李明")
	fmt.Printf("学生: %s (ID: %d)\n", student.Name, student.ID)

	// 添加分数
	scores := []int{85, 92, 78, 95, 88}
	for _, score := range scores {
		student.AddScore(score)
	}

	fmt.Printf("分数: %v\n", student.Scores)
	fmt.Printf("平均分: %.2f\n", student.Average())
	fmt.Printf("最高分: %d\n", student.HighestScore())
	fmt.Printf("最低分: %d\n", student.LowestScore())

	// 再添加一些分数
	student.AddScore(90)
	student.AddScore(87)

	fmt.Printf("更新后分数: %v\n", student.Scores)
	fmt.Printf("更新后平均分: %.2f\n", student.Average())
}

// 事件管理示例
func EventManagementExample() {
	fmt.Println("\n=== 事件管理示例 ===")

	// 创建事件
	startTime := time.Now().Add(time.Hour) // 1小时后开始
	event := Event{
		Name:      "Go语言培训",
		StartTime: startTime,
		Duration:  2 * time.Hour, // 持续2小时
	}

	fmt.Printf("事件: %s\n", event.Name)
	fmt.Printf("开始时间: %s\n", event.StartTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("结束时间: %s\n", event.EndTime().Format("2006-01-02 15:04:05"))
	fmt.Printf("持续时间: %v\n", event.Duration)

	// 检查当前状态
	now := time.Now()
	fmt.Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("是否正在进行: %t\n", event.IsOngoing(now))

	if !event.IsOngoing(now) {
		remaining := event.TimeRemaining(now)
		if remaining > 0 {
			fmt.Printf("距离结束还有: %v\n", remaining)
		} else {
			fmt.Println("事件已结束")
		}
	}
}

// 比较结构体
func CompareStructs() {
	fmt.Println("\n=== 结构体比较 ===")

	person1 := Person{Name: "张三", Age: 25, City: "北京"}
	person2 := Person{Name: "张三", Age: 25, City: "北京"}
	person3 := Person{Name: "李四", Age: 30, City: "上海"}

	fmt.Printf("person1: %+v\n", person1)
	fmt.Printf("person2: %+v\n", person2)
	fmt.Printf("person3: %+v\n", person3)

	fmt.Printf("person1 == person2: %t\n", person1 == person2)
	fmt.Printf("person1 == person3: %t\n", person1 == person3)

	// 注意：包含切片、映射或函数的结构体不能直接比较
	student1 := Student{ID: 1, Name: "张三", Scores: []int{85, 90}}
	student2 := Student{ID: 1, Name: "张三", Scores: []int{85, 90}}

	// 这会编译错误：student1 == student2
	// 需要手动比较各个字段
	fmt.Printf("学生ID相同: %t\n", student1.ID == student2.ID)
	fmt.Printf("学生姓名相同: %t\n", student1.Name == student2.Name)
}
