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
	case "array", "arrays":
		fmt.Println("运行数组示例...")
		runArrayExample()
	case "slice", "slices":
		fmt.Println("运行切片示例...")
		runSliceExample()
	case "map", "maps":
		fmt.Println("运行映射示例...")
		runMapExample()
	case "all":
		fmt.Println("运行所有数组、切片和映射示例...")
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("数组示例")
		fmt.Println(strings.Repeat("=", 50))
		runArrayExample()

		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("切片示例")
		fmt.Println(strings.Repeat("=", 50))
		runSliceExample()

		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("映射示例")
		fmt.Println(strings.Repeat("=", 50))
		runMapExample()
	default:
		fmt.Println("未知示例类型。请使用以下参数之一:")
		fmt.Println("  array  - 运行数组示例")
		fmt.Println("  slice  - 运行切片示例")
		fmt.Println("  map    - 运行映射示例")
		fmt.Println("  all    - 运行所有示例 (默认)")
	}
}
