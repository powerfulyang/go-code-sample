package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/example/golang-examples/01-basics/data_types/datatypes"
)

func main() {
	// 获取命令行参数，确定要运行的示例
	args := os.Args[1:]
	example := "all"

	if len(args) > 0 {
		example = strings.ToLower(args[0])
	}

	switch example {
	case "numeric", "numbers", "num":
		fmt.Println("运行数值类型示例...")
		datatypes.NumericExample()
	case "string", "strings", "str":
		fmt.Println("运行字符串类型示例...")
		datatypes.StringsExample()
	case "boolean", "bool":
		fmt.Println("运行布尔类型示例...")
		datatypes.BooleansExample()
	case "all":
		fmt.Println("运行所有数据类型示例...")
		fmt.Println("\n---------------------------------------------")
		fmt.Println("数值类型示例")
		fmt.Println("---------------------------------------------")
		datatypes.NumericExample()

		fmt.Println("\n---------------------------------------------")
		fmt.Println("字符串类型示例")
		fmt.Println("---------------------------------------------")
		datatypes.StringsExample()

		fmt.Println("\n---------------------------------------------")
		fmt.Println("布尔类型示例")
		fmt.Println("---------------------------------------------")
		datatypes.BooleansExample()
	default:
		fmt.Println("未知示例类型。请使用以下参数之一:")
		fmt.Println("  numeric - 运行数值类型示例")
		fmt.Println("  string  - 运行字符串类型示例")
		fmt.Println("  boolean - 运行布尔类型示例")
		fmt.Println("  all     - 运行所有示例 (默认)")
	}
}

// 以下是三个示例函数，由于main包与datatypes包在同一个目录下
// 我们需要重新实现这些函数，以便在main包中使用

func NumericExample() {
	datatypes.NumericExample()
}

func StringsExample() {
	datatypes.StringsExample()
}

func BooleansExample() {
	datatypes.BooleansExample()
}
