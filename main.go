package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	// 高级特性模块
	// 生态系统模块
)

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Greet 方法
func (p Person) Greet() string {
	return fmt.Sprintf("你好，我是 %s，今年 %d 岁", p.Name, p.Age)
}

// 简单的演示程序
func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	module := os.Args[1]

	switch module {
	case "demo":
		runBasicDemo()
	case "test":
		runTestDemo()
	case "interfaces":
		runInterfacesDemo()
	case "concurrency":
		runConcurrencyDemo()
	case "generics":
		runGenericsDemo()
	case "reflection":
		runReflectionDemo()
	case "testing":
		runTestingDemo()
	case "stdlib":
		runStandardLibraryDemo()
	case "all":
		runAllDemos()
	default:
		fmt.Printf("❌ 未知模块: %s\n", module)
		showUsage()
	}
}

// 加法函数
func add(a, b int) int {
	return a + b
}

// 除法和取余函数
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// 安全除法函数
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 基本演示
func runBasicDemo() {
	fmt.Println("🚀 Go语言基础演示")
	fmt.Println(strings.Repeat("=", 50))

	// 变量和常量
	fmt.Println("\n🔹 1. 变量和常量")
	name := "Go语言"
	version := 1.21
	const maxUsers = 1000

	fmt.Printf("语言: %s, 版本: %.2f, 最大用户数: %d\n", name, version, maxUsers)

	// 数据类型
	fmt.Println("\n🔹 2. 数据类型")
	var numbers []int = []int{1, 2, 3, 4, 5}
	userInfo := map[string]interface{}{
		"name": "张三",
		"age":  25,
		"city": "北京",
	}

	fmt.Printf("数组: %v\n", numbers)
	fmt.Printf("用户信息: %v\n", userInfo)

	// 控制流程
	fmt.Println("\n🔹 3. 控制流程")
	for i, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("索引 %d: %d 是偶数\n", i, num)
		} else {
			fmt.Printf("索引 %d: %d 是奇数\n", i, num)
		}
	}

	// 函数
	fmt.Println("\n🔹 4. 函数")
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	quotient, remainder := divmod(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", quotient, remainder)

	// 结构体
	fmt.Println("\n🔹 5. 结构体")
	person := Person{Name: "李四", Age: 30}
	fmt.Printf("人员信息: %+v\n", person)
	fmt.Printf("问候: %s\n", person.Greet())

	// 错误处理
	fmt.Println("\n🔹 6. 错误处理")
	result2, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %.2f\n", result2)
	}
}

// 测试演示
func runTestDemo() {
	fmt.Println("🧪 运行测试演示")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("请运行以下命令来执行测试:")
	fmt.Println()
	fmt.Println("# 运行所有测试")
	fmt.Println("go test ./...")
	fmt.Println()
	fmt.Println("# 运行特定模块测试")
	fmt.Println("go test -v ./01-basics/03-data-types/...")
	fmt.Println("go test -v ./01-basics/05-functions/...")
	fmt.Println("go test -v ./01-basics/07-structs-methods/...")
	fmt.Println()
	fmt.Println("# 运行基准测试")
	fmt.Println("go test -bench=. ./01-basics/05-functions/...")
	fmt.Println()
	fmt.Println("# 查看测试覆盖率")
	fmt.Println("go test -cover ./01-basics/03-data-types/...")
}

func showUsage() {
	fmt.Println("🚀 Go语言学习示例")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("用法: go run main.go <模块名>")
	fmt.Println()
	fmt.Println("可用模块:")
	fmt.Println("  demo        - 基础演示")
	fmt.Println("  test        - 测试说明")
	fmt.Println("  interfaces  - 接口示例")
	fmt.Println("  concurrency - 并发编程示例")
	fmt.Println("  generics    - 泛型示例")
	fmt.Println("  reflection  - 反射示例")
	fmt.Println("  testing     - 测试框架示例")
	fmt.Println("  stdlib      - 标准库示例")
	fmt.Println("  all         - 运行所有示例")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run main.go demo")
	fmt.Println("  go run main.go interfaces")
	fmt.Println("  go run main.go all")
}

// 新增的演示函数
func runInterfacesDemo() {
	fmt.Println("🔹 接口示例演示")
	fmt.Println("请先运行: go run ./02-advanced-features/01-interfaces/")
	fmt.Println("或查看接口相关代码和测试")
}

func runConcurrencyDemo() {
	fmt.Println("🔹 并发编程示例演示")
	fmt.Println("请先运行: go run ./02-advanced-features/02-concurrency/")
	fmt.Println("或查看并发相关代码和测试")
}

func runGenericsDemo() {
	fmt.Println("🔹 泛型示例演示")
	fmt.Println("请先运行: go test -v ./02-advanced-features/03-generics/")
	fmt.Println("或查看泛型相关代码和测试")
}

func runReflectionDemo() {
	fmt.Println("🔹 反射示例演示")
	fmt.Println("请先运行: go test -v ./02-advanced-features/04-reflection/")
	fmt.Println("或查看反射相关代码和测试")
}

func runTestingDemo() {
	fmt.Println("🔹 测试框架示例演示")
	fmt.Println("请先运行: go test -v ./02-advanced-features/05-testing/")
	fmt.Println("或查看测试框架相关代码和测试")
}

func runStandardLibraryDemo() {
	fmt.Println("🔹 标准库示例演示")
	fmt.Println("请先运行: go test -v ./03-ecosystem/01-standard-library/")
	fmt.Println("或查看标准库相关代码和测试")
}

func runAllDemos() {
	fmt.Println("🚀 运行所有示例演示")
	fmt.Println(strings.Repeat("=", 50))

	runBasicDemo()
	fmt.Println()
	runInterfacesDemo()
	fmt.Println()
	runConcurrencyDemo()
	fmt.Println()
	runGenericsDemo()
	fmt.Println()
	runReflectionDemo()
	fmt.Println()
	runTestingDemo()
	fmt.Println()
	runStandardLibraryDemo()

	fmt.Println("\n✅ 所有示例演示完成!")
	fmt.Println("💡 提示: 运行 'go test ./...' 来执行所有测试")
}
