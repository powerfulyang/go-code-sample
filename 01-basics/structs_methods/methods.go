package main

import (
	"fmt"
	"math"
)

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// BankAccount 银行账户结构体
type BankAccount struct {
	AccountNumber string
	Balance       float64
	Owner         string
}

// Counter 计数器结构体
type Counter struct {
	Value int
}

// 值接收者方法
// Area 计算矩形面积（值接收者）
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长（值接收者）
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// String 实现fmt.Stringer接口（值接收者）
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

// IsSquare 判断是否为正方形（值接收者）
func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

// 指针接收者方法
// Scale 缩放矩形（指针接收者，会修改原结构体）
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// SetDimensions 设置矩形尺寸（指针接收者）
func (r *Rectangle) SetDimensions(width, height float64) {
	r.Width = width
	r.Height = height
}

// Circle的方法
// Area 计算圆形面积（值接收者）
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circumference 计算圆形周长（值接收者）
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// String 实现fmt.Stringer接口（值接收者）
func (c Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// SetRadius 设置圆形半径（指针接收者）
func (c *Circle) SetRadius(radius float64) {
	if radius > 0 {
		c.Radius = radius
	}
}

// BankAccount的方法
// GetBalance 获取余额（值接收者）
func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

// String 实现fmt.Stringer接口（值接收者）
func (b BankAccount) String() string {
	return fmt.Sprintf("Account: %s, Owner: %s, Balance: $%.2f",
		b.AccountNumber, b.Owner, b.Balance)
}

// Deposit 存款（指针接收者）
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("存入 $%.2f，当前余额: $%.2f\n", amount, b.Balance)
	} else {
		fmt.Println("存款金额必须大于0")
	}
}

// Withdraw 取款（指针接收者）
func (b *BankAccount) Withdraw(amount float64) bool {
	if amount <= 0 {
		fmt.Println("取款金额必须大于0")
		return false
	}
	if amount > b.Balance {
		fmt.Println("余额不足")
		return false
	}
	b.Balance -= amount
	fmt.Printf("取出 $%.2f，当前余额: $%.2f\n", amount, b.Balance)
	return true
}

// Transfer 转账（指针接收者）
func (b *BankAccount) Transfer(to *BankAccount, amount float64) bool {
	if b.Withdraw(amount) {
		to.Deposit(amount)
		fmt.Printf("从 %s 转账 $%.2f 到 %s\n", b.Owner, amount, to.Owner)
		return true
	}
	return false
}

// Counter的方法
// Get 获取计数值（值接收者）
func (c Counter) Get() int {
	return c.Value
}

// Increment 增加计数（指针接收者）
func (c *Counter) Increment() {
	c.Value++
}

// IncrementBy 增加指定数量（指针接收者）
func (c *Counter) IncrementBy(n int) {
	c.Value += n
}

// Reset 重置计数（指针接收者）
func (c *Counter) Reset() {
	c.Value = 0
}

func runMethodExample() {
	fmt.Println("=== Go 方法示例 ===")

	// 值接收者方法示例
	fmt.Println("\n--- 值接收者方法 ---")
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Printf("矩形: %s\n", rect)
	fmt.Printf("面积: %.2f\n", rect.Area())
	fmt.Printf("周长: %.2f\n", rect.Perimeter())
	fmt.Printf("是否为正方形: %t\n", rect.IsSquare())

	// 指针接收者方法示例
	fmt.Println("\n--- 指针接收者方法 ---")
	fmt.Printf("缩放前: %s\n", rect)
	rect.Scale(2.0) // Go自动取地址
	fmt.Printf("缩放2倍后: %s\n", rect)

	// 显式使用指针
	rectPtr := &Rectangle{Width: 3, Height: 4}
	fmt.Printf("指针矩形: %s\n", rectPtr)
	rectPtr.SetDimensions(6, 8)
	fmt.Printf("设置尺寸后: %s\n", rectPtr)

	// 圆形示例
	fmt.Println("\n--- 圆形方法 ---")
	circle := Circle{Radius: 5}
	fmt.Printf("圆形: %s\n", circle)
	fmt.Printf("面积: %.2f\n", circle.Area())
	fmt.Printf("周长: %.2f\n", circle.Circumference())

	circle.SetRadius(10)
	fmt.Printf("设置半径后: %s\n", circle)
	fmt.Printf("新面积: %.2f\n", circle.Area())

	// 银行账户示例
	fmt.Println("\n--- 银行账户方法 ---")
	account1 := BankAccount{
		AccountNumber: "ACC001",
		Balance:       1000.0,
		Owner:         "Alice",
	}
	account2 := BankAccount{
		AccountNumber: "ACC002",
		Balance:       500.0,
		Owner:         "Bob",
	}

	fmt.Printf("账户1: %s\n", account1)
	fmt.Printf("账户2: %s\n", account2)

	// 存款
	account1.Deposit(200)

	// 取款
	account1.Withdraw(150)

	// 转账
	account1.Transfer(&account2, 300)

	fmt.Printf("操作后账户1: %s\n", account1)
	fmt.Printf("操作后账户2: %s\n", account2)

	// 计数器示例
	fmt.Println("\n--- 计数器方法 ---")
	counter := Counter{Value: 0}

	fmt.Printf("初始计数: %d\n", counter.Get())

	counter.Increment()
	fmt.Printf("增加1后: %d\n", counter.Get())

	counter.IncrementBy(5)
	fmt.Printf("增加5后: %d\n", counter.Get())

	counter.Reset()
	fmt.Printf("重置后: %d\n", counter.Get())

	// 方法值和方法表达式
	fmt.Println("\n--- 方法值和方法表达式 ---")
	rect2 := Rectangle{Width: 4, Height: 3}

	// 方法值：绑定到特定实例
	areaMethod := rect2.Area
	fmt.Printf("通过方法值计算面积: %.2f\n", areaMethod())

	// 方法表达式：需要传入接收者
	areaFunc := Rectangle.Area
	fmt.Printf("通过方法表达式计算面积: %.2f\n", areaFunc(rect2))

	// 指针接收者的方法值
	scaleMethod := rect2.Scale
	fmt.Printf("缩放前: %s\n", rect2)
	scaleMethod(1.5)
	fmt.Printf("缩放后: %s\n", rect2)

	// 接口和方法
	fmt.Println("\n--- 接口和方法 ---")
	shapes := []fmt.Stringer{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		&Rectangle{Width: 2, Height: 2},
		&Circle{Radius: 7},
	}

	fmt.Println("形状列表:")
	for i, shape := range shapes {
		fmt.Printf("  %d: %s\n", i+1, shape)
	}

	// 值接收者 vs 指针接收者的区别
	fmt.Println("\n--- 值接收者 vs 指针接收者 ---")

	// 值接收者：不会修改原结构体
	rect3 := Rectangle{Width: 10, Height: 5}
	fmt.Printf("原始矩形: %s\n", rect3)

	// 即使通过指针调用值接收者方法，也不会修改原结构体
	rectPtr3 := &rect3
	area := rectPtr3.Area() // 值接收者方法
	fmt.Printf("面积: %.2f，矩形未变: %s\n", area, rect3)

	// 指针接收者：会修改原结构体
	rectPtr3.Scale(2) // 指针接收者方法
	fmt.Printf("缩放后矩形: %s\n", rect3)
}
