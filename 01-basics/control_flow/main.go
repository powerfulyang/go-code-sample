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
	case "if", "if-else":
		fmt.Println("运行if-else条件语句示例...")
		IfElseExample()
	case "for", "loops":
		fmt.Println("运行for循环示例...")
		ForLoopsExample()
	case "switch", "switch-case":
		fmt.Println("运行switch-case条件语句示例...")
		SwitchCaseExample()
	case "all":
		fmt.Println("运行所有控制流示例...")

		fmt.Println("\n---------------------------------------------")
		fmt.Println("IF-ELSE条件语句")
		fmt.Println("---------------------------------------------")
		IfElseExample()

		fmt.Println("\n---------------------------------------------")
		fmt.Println("FOR循环")
		fmt.Println("---------------------------------------------")
		ForLoopsExample()

		fmt.Println("\n---------------------------------------------")
		fmt.Println("SWITCH-CASE条件语句")
		fmt.Println("---------------------------------------------")
		SwitchCaseExample()
	default:
		fmt.Println("未知示例类型。请使用以下参数之一:")
		fmt.Println("  if-else   - 运行if-else条件语句示例")
		fmt.Println("  for       - 运行for循环示例")
		fmt.Println("  switch    - 运行switch-case条件语句示例")
		fmt.Println("  all       - 运行所有示例 (默认)")
	}
}
