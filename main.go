package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	// 基础模块

	// 高级特性模块

	// 生态系统模块
	packages "golang-examples/01-basics/10-packages-modules"
	performance "golang-examples/02-advanced-features/06-performance"
	thirdparty "golang-examples/03-ecosystem/02-third-party-libraries"
	tools "golang-examples/03-ecosystem/03-go-tools"
	libraries "golang-examples/03-ecosystem/04-popular-libraries"

	// 实际应用模块
	webapi "golang-examples/04-practical-applications/01-web-api"
	database "golang-examples/04-practical-applications/02-database"
	cli "golang-examples/04-practical-applications/03-cli-tool"
	network "golang-examples/04-practical-applications/04-network"
	security "golang-examples/04-practical-applications/07-security"
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
	case "functional":
		runFunctionalDemo()
	case "webapi":
		runWebAPIDemo()
	case "database":
		runDatabaseDemo()
	case "cli":
		runCLIDemo()
	case "network":
		runNetworkDemo()
	case "security":
		runSecurityDemo()
	case "packages":
		runPackagesDemo()
	case "performance":
		runPerformanceDemo()
	case "tools":
		runToolsDemo()
	case "popular":
		runPopularLibrariesDemo()
	case "all":
		runAllDemos()
	default:
		fmt.Printf("❌ 抱歉，找不到模块 '%s'\n", module)
		fmt.Println("💡 提示: 请检查模块名称是否正确，或查看下面的可用模块列表")
		fmt.Println()
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
	fmt.Println("🎓 欢迎使用 Go 语言学习示例项目！")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("📖 这是一个完整的 Go 语言学习资源，包含从基础到高级的所有内容")
	fmt.Println()
	fmt.Println("🔧 使用方法:")
	fmt.Println("   go run main.go <模块名>")
	fmt.Println()
	fmt.Println("📚 学习模块 (建议按顺序学习):")
	fmt.Println()
	fmt.Println("🌱 基础入门:")
	fmt.Println("   demo        - 🎯 Go语言基础演示 (推荐新手从这里开始)")
	fmt.Println("   packages    - 📦 包和模块系统 (理解Go的组织方式)")
	fmt.Println()
	fmt.Println("🚀 核心特性:")
	fmt.Println("   interfaces  - 🔌 接口编程 (Go的核心设计理念)")
	fmt.Println("   concurrency - ⚡ 并发编程 (Go的杀手级特性)")
	fmt.Println("   generics    - 🧬 泛型编程 (Go 1.18+ 新特性)")
	fmt.Println("   reflection  - 🪞 反射机制 (动态编程技巧)")
	fmt.Println()
	fmt.Println("🛠️ 开发技能:")
	fmt.Println("   testing     - 🧪 测试框架 (保证代码质量)")
	fmt.Println("   performance - 🏃 性能优化 (编写高效代码)")
	fmt.Println("   tools       - 🔨 Go工具链 (提高开发效率)")
	fmt.Println()
	fmt.Println("🌐 生态系统:")
	fmt.Println("   stdlib      - 📚 标准库使用 (Go内置功能)")
	fmt.Println("   functional  - 🔄 函数式编程 (现代编程范式)")
	fmt.Println("   popular     - ⭐ 流行库使用 (社区最佳实践)")
	fmt.Println()
	fmt.Println("💼 实战项目:")
	fmt.Println("   webapi      - 🌍 Web API开发 (构建REST服务)")
	fmt.Println("   database    - 🗄️  数据库操作 (数据持久化)")
	fmt.Println("   cli         - 💻 CLI工具开发 (命令行应用)")
	fmt.Println("   network     - 🔗 网络编程 (TCP/UDP/WebSocket)")
	fmt.Println("   security    - 🔐 安全认证 (JWT/加密技术)")
	fmt.Println()
	fmt.Println("🎯 特殊选项:")
	fmt.Println("   test        - 📋 测试说明 (如何运行测试)")
	fmt.Println("   all         - 🎪 运行所有示例 (完整演示)")
	fmt.Println()
	fmt.Println("💡 使用建议:")
	fmt.Println("   • 新手推荐: demo → packages → interfaces → concurrency")
	fmt.Println("   • 进阶学习: generics → reflection → testing → performance")
	fmt.Println("   • 实战练习: webapi → database → cli → network → security")
	fmt.Println("   • 生态了解: stdlib → functional → popular → tools")
	fmt.Println()
	fmt.Println("🚀 快速开始:")
	fmt.Println("   go run main.go demo      # 🎯 从基础演示开始")
	fmt.Println("   go run main.go all       # 🎪 查看所有功能")
	fmt.Println("   go test ./...            # 🧪 运行所有测试")
	fmt.Println()
	fmt.Println("📖 更多信息请查看 README.md 文件")
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

// 新增的演示函数
func runFunctionalDemo() {
	fmt.Println("🔹 函数式编程示例演示")
	fmt.Println(strings.Repeat("=", 50))
	thirdparty.FunctionalExamples()
}

func runWebAPIDemo() {
	fmt.Println("🔹 Web API示例演示")
	fmt.Println(strings.Repeat("=", 50))
	webapi.WebAPIExamples()
}

func runDatabaseDemo() {
	fmt.Println("🔹 数据库操作示例演示")
	fmt.Println(strings.Repeat("=", 50))
	database.DatabaseExamples()
}

func runCLIDemo() {
	fmt.Println("🔹 CLI工具示例演示")
	fmt.Println(strings.Repeat("=", 50))
	cli.CLIExamples()
}

func runNetworkDemo() {
	fmt.Println("🔹 网络编程示例演示")
	fmt.Println(strings.Repeat("=", 50))
	network.TCPExamples()
	fmt.Println()
	network.UDPExamples()
	fmt.Println()
	network.WebSocketExamples()
}

func runSecurityDemo() {
	fmt.Println("🔹 安全和认证示例演示")
	fmt.Println(strings.Repeat("=", 50))
	security.JWTExamples()
	fmt.Println()
	security.EncryptionExamples()
}

func runPackagesDemo() {
	fmt.Println("📦 Go语言包和模块系统")
	fmt.Println("🎯 理解Go的代码组织方式")
	fmt.Println(strings.Repeat("=", 60))
	packages.PackageExamples()
}

func runPerformanceDemo() {
	fmt.Println("🏃 Go语言性能优化")
	fmt.Println("🎯 让您的代码飞起来")
	fmt.Println(strings.Repeat("=", 60))
	performance.PerformanceExamples()
}

func runToolsDemo() {
	fmt.Println("🔨 Go开发工具链")
	fmt.Println("🎯 提高开发效率的利器")
	fmt.Println(strings.Repeat("=", 60))
	tools.GoToolsExamples()
}

func runPopularLibrariesDemo() {
	fmt.Println("⭐ Go生态系统流行库")
	fmt.Println("🎯 站在巨人的肩膀上")
	fmt.Println(strings.Repeat("=", 60))
	libraries.PopularLibrariesExamples()
}

func runAllDemos() {
	fmt.Println("🎪 Go语言学习项目 - 完整演示")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🚀 即将为您展示所有模块的精彩内容！")
	fmt.Println("⏱️  预计需要几分钟时间，请耐心等待...")
	fmt.Println("💡 您可以随时按 Ctrl+C 中断演示")
	fmt.Println()
	fmt.Println("📋 演示内容包括:")
	fmt.Println("   • 基础语法和数据类型")
	fmt.Println("   • 高级特性和并发编程")
	fmt.Println("   • 生态系统和工具链")
	fmt.Println("   • 实战项目和最佳实践")
	fmt.Println()
	fmt.Println("🎬 演示开始...")
	fmt.Println()

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
	fmt.Println()
	runFunctionalDemo()
	fmt.Println()
	runWebAPIDemo()
	fmt.Println()
	runDatabaseDemo()
	fmt.Println()
	runCLIDemo()
	fmt.Println()
	runNetworkDemo()
	fmt.Println()
	runSecurityDemo()
	fmt.Println()
	runPackagesDemo()
	fmt.Println()
	runPerformanceDemo()
	fmt.Println()
	runToolsDemo()
	fmt.Println()
	runPopularLibrariesDemo()

	fmt.Println("\n🎉 恭喜！所有示例演示完成！")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🎓 您已经完整体验了Go语言学习项目的所有内容")
	fmt.Println()
	fmt.Println("🚀 接下来您可以:")
	fmt.Println("   📚 深入学习感兴趣的模块: go run main.go <模块名>")
	fmt.Println("   🧪 运行测试验证理解: go test ./...")
	fmt.Println("   🔨 查看基准测试: go test -bench=. ./...")
	fmt.Println("   📖 阅读源码了解实现细节")
	fmt.Println("   💼 开始您的Go项目实践")
	fmt.Println()
	fmt.Println("💡 学习建议:")
	fmt.Println("   • 多动手实践，修改代码看效果")
	fmt.Println("   • 阅读Go官方文档和最佳实践")
	fmt.Println("   • 参与Go社区，分享学习心得")
	fmt.Println("   • 构建实际项目应用所学知识")
	fmt.Println()
	fmt.Println("🌟 祝您Go语言学习愉快！")
}
