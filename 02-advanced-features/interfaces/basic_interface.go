package main

import "fmt"

// Shape 形状接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Stringer 字符串接口（类似fmt.Stringer）
type Stringer interface {
	String() string
}

// Rectangle 矩形
type Rectangle struct {
	Width, Height float64
}

// Circle 圆形
type Circle struct {
	Radius float64
}

// 实现Shape接口 - Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

// 实现Shape接口 - Circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius: %.2f)", c.Radius)
}

func main() {
	fmt.Println("=== Go 接口示例 ===")

	// 接口变量
	fmt.Println("\n--- 接口变量 ---")
	var shape Shape

	// 将Rectangle赋值给接口
	rect := Rectangle{Width: 10, Height: 5}
	shape = rect
	fmt.Printf("矩形面积: %.2f\n", shape.Area())
	fmt.Printf("矩形周长: %.2f\n", shape.Perimeter())

	// 将Circle赋值给接口
	circle := Circle{Radius: 3}
	shape = circle
	fmt.Printf("圆形面积: %.2f\n", shape.Area())
	fmt.Printf("圆形周长: %.2f\n", shape.Perimeter())

	// 接口切片
	fmt.Println("\n--- 接口切片 ---")
	shapes := []Shape{
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 2.5},
		Rectangle{Width: 3, Height: 3},
	}

	totalArea := 0.0
	for i, s := range shapes {
		area := s.Area()
		totalArea += area
		fmt.Printf("形状%d 面积: %.2f\n", i+1, area)
	}
	fmt.Printf("总面积: %.2f\n", totalArea)

	// 类型断言
	fmt.Println("\n--- 类型断言 ---")
	var s Shape = Rectangle{Width: 8, Height: 4}

	// 安全的类型断言
	if rect, ok := s.(Rectangle); ok {
		fmt.Printf("这是一个矩形: %+v\n", rect)
		fmt.Printf("宽度: %.2f, 高度: %.2f\n", rect.Width, rect.Height)
	} else {
		fmt.Println("不是矩形")
	}

	// 类型断言失败的例子
	if circle, ok := s.(Circle); ok {
		fmt.Printf("这是一个圆形: %+v\n", circle)
	} else {
		fmt.Println("不是圆形")
	}

	// type switch
	fmt.Println("\n--- Type Switch ---")
	shapes2 := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		Rectangle{Width: 2, Height: 8},
	}

	for i, shape := range shapes2 {
		switch s := shape.(type) {
		case Rectangle:
			fmt.Printf("形状%d: 矩形 %.2f x %.2f\n", i+1, s.Width, s.Height)
		case Circle:
			fmt.Printf("形状%d: 圆形 半径 %.2f\n", i+1, s.Radius)
		default:
			fmt.Printf("形状%d: 未知类型\n", i+1)
		}
	}

	// 多接口实现
	fmt.Println("\n--- 多接口实现 ---")
	var stringer Stringer = Rectangle{Width: 7, Height: 2}
	fmt.Printf("字符串表示: %s\n", stringer.String())

	// 接口组合
	fmt.Println("\n--- 接口组合 ---")
	type ShapeStringer interface {
		Shape
		Stringer
	}

	var ss ShapeStringer = Rectangle{Width: 6, Height: 4}
	fmt.Printf("面积: %.2f\n", ss.Area())
	fmt.Printf("描述: %s\n", ss.String())

	// 空接口
	fmt.Println("\n--- 空接口 ---")
	var anything interface{}

	anything = 42
	fmt.Printf("整数: %v (类型: %T)\n", anything, anything)

	anything = "Hello"
	fmt.Printf("字符串: %v (类型: %T)\n", anything, anything)

	anything = Rectangle{Width: 1, Height: 2}
	fmt.Printf("结构体: %v (类型: %T)\n", anything, anything)

	// 接口值的比较
	fmt.Println("\n--- 接口值比较 ---")
	var shape1 Shape = Rectangle{Width: 3, Height: 4}
	var shape2 Shape = Rectangle{Width: 3, Height: 4}
	var shape3 Shape = Circle{Radius: 2}

	fmt.Printf("shape1 == shape2: %t\n", shape1 == shape2)
	fmt.Printf("shape1 == shape3: %t\n", shape1 == shape3)

	// nil接口
	fmt.Println("\n--- nil接口 ---")
	var nilShape Shape
	fmt.Printf("nil接口: %v\n", nilShape)
	fmt.Printf("是否为nil: %t\n", nilShape == nil)

	if nilShape == nil {
		fmt.Println("接口为nil，不能调用方法")
	}
}
