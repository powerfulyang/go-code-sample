package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 获取命令行参数，确定要运行的示例
	args := os.Args[1:]
	example := "all"

	if len(args) > 0 {
		example = strings.ToLower(args[0])
	}

	switch example {
	case "struct", "structs":
		fmt.Println("运行结构体示例...")
		runStructExample()
	case "method", "methods":
		fmt.Println("运行方法示例...")
		runMethodExample()
	case "all":
		fmt.Println("运行所有结构体和方法示例...")
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("结构体示例")
		fmt.Println(strings.Repeat("=", 50))
		runStructExample()

		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("方法示例")
		fmt.Println(strings.Repeat("=", 50))
		runMethodExample()
	default:
		fmt.Println("未知示例类型。请使用以下参数之一:")
		fmt.Println("  struct  - 运行结构体示例")
		fmt.Println("  method  - 运行方法示例")
		fmt.Println("  all     - 运行所有示例 (默认)")
	}
}
