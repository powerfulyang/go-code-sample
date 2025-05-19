package main

import (
	"fmt"
)

// 1. 无参数无返回值的函数
func sayHello() {
	fmt.Println("你好，世界！")
}

// 2. 带参数的函数
func greet(name string) {
	fmt.Printf("你好，%s！\n", name)
}

// 3. 带多个参数的函数
func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

// 4. 带参数和返回值的函数
func multiply(a, b int) int {
	return a * b
}

// 5. 带命名返回值的函数
func divide(a, b float64) (result float64) {
	result = a / b
	return // 裸返回，会返回result的值
}

// 6. 带多个返回值的函数
func calculateStats(numbers []int) (min, max, sum int) {
	if len(numbers) == 0 {
		return 0, 0, 0
	}

	min = numbers[0]
	max = numbers[0]
	sum = 0

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}

	return
}

// 7. 函数变量 - 函数也是值
func getGreetFunction() func(string) {
	return greet
}

// 8. 函数作为参数
func executeFunction(f func(string), name string) {
	f(name)
}

// 9. 函数递归
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 10. 函数类型定义
type MathFunc func(int, int) int

func applyMathFunc(f MathFunc, a, b int) int {
	return f(a, b)
}

func BasicFunctionsExample() {
	fmt.Println("=== 基本函数示例 ===")

	// 1. 调用无参数无返回值的函数
	fmt.Println("\n无参数无返回值的函数:")
	sayHello()

	// 2. 调用带参数的函数
	fmt.Println("\n带参数的函数:")
	greet("张三")

	// 3. 调用带多个参数的函数
	fmt.Println("\n带多个参数的函数:")
	add(5, 3)

	// 4. 调用带返回值的函数
	fmt.Println("\n带返回值的函数:")
	product := multiply(4, 7)
	fmt.Printf("4 x 7 = %d\n", product)

	// 5. 调用带命名返回值的函数
	fmt.Println("\n带命名返回值的函数:")
	quotient := divide(10, 3)
	fmt.Printf("10 / 3 = %.2f\n", quotient)

	// 6. 调用带多个返回值的函数
	fmt.Println("\n带多个返回值的函数:")
	numbers := []int{7, 2, 9, 4, 6}
	minVal, maxVal, sumVal := calculateStats(numbers)
	fmt.Printf("数组: %v\n", numbers)
	fmt.Printf("最小值: %d, 最大值: %d, 总和: %d\n", minVal, maxVal, sumVal)

	// 7. 函数变量
	fmt.Println("\n函数变量:")
	greetFunc := getGreetFunction()
	greetFunc("李四")

	// 8. 将函数作为参数传递
	fmt.Println("\n函数作为参数:")
	executeFunction(greet, "王五")

	// 9. 递归函数
	fmt.Println("\n递归函数 - 阶乘:")
	fmt.Printf("5! = %d\n", factorial(5))

	// 10. 函数类型
	fmt.Println("\n函数类型:")
	result := applyMathFunc(multiply, 5, 6)
	fmt.Printf("5 x 6 = %d\n", result)
}
