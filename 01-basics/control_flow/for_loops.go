package main

import (
	"fmt"
	"strings"
)

func ForLoopsExample() {
	// 1. 基本for循环（类似C/Java的for循环）
	fmt.Println("\n基本for循环:")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// 2. 只有条件的for循环（相当于while循环）
	fmt.Println("\n只有条件的for循环:")
	j := 0
	for j < 5 {
		fmt.Printf("j = %d\n", j)
		j++
	}

	// 3. 无限循环 + break
	fmt.Println("\n无限循环 + break:")
	counter := 0
	for {
		fmt.Printf("counter = %d\n", counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	// 4. 使用continue跳过迭代
	fmt.Println("\n使用continue跳过偶数:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("奇数: %d\n", i)
	}

	// 5. 嵌套循环
	fmt.Println("\n嵌套循环 - 乘法表:")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 6. break 标签
	fmt.Println("\nbreak 标签示例:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("在 i =", i, "j =", j, "时跳出外层循环")
				break OuterLoop
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}

	// 7. for-range 循环
	fmt.Println("\nfor-range 循环遍历切片:")
	fruits := []string{"苹果", "香蕉", "橙子", "葡萄"}
	for index, fruit := range fruits {
		fmt.Printf("索引: %d, 值: %s\n", index, fruit)
	}

	// 如果不需要索引，可以使用下划线忽略
	fmt.Println("\n忽略索引:")
	for _, fruit := range fruits {
		fmt.Printf("水果: %s\n", fruit)
	}

	// 如果只需要索引，不需要值
	fmt.Println("\n只使用索引:")
	for i := range fruits {
		fmt.Printf("索引 %d: %s\n", i, fruits[i])
	}

	// 8. for-range 循环遍历map
	fmt.Println("\nfor-range 循环遍历map:")
	scores := map[string]int{
		"张三": 90,
		"李四": 85,
		"王五": 78,
	}

	for name, score := range scores {
		fmt.Printf("学生: %s, 分数: %d\n", name, score)
	}

	// 9. for-range 循环遍历字符串（Unicode码点）
	fmt.Println("\nfor-range 循环遍历字符串:")
	text := "Hello, 世界!"
	for i, char := range text {
		fmt.Printf("位置 %d: %c [%d]\n", i, char, char)
	}

	// 10. 使用for循环模拟do-while循环
	fmt.Println("\n模拟do-while循环:")
	k := 0
	for {
		fmt.Printf("k = %d\n", k)
		k++
		if k >= 5 {
			break
		}
	}

	// 11. 循环控制技巧

	// 使用函数分解复杂循环
	fmt.Println("\n使用函数分解复杂循环:")
	printPyramid(5)

	// 更复杂的示例：使用循环生成九宫格
	fmt.Println("\n生成九宫格:")
	gridSize := 3
	generateGrid(gridSize)
}

// 打印一个金字塔图案
func printPyramid(height int) {
	for i := 1; i <= height; i++ {
		// 打印前置空格
		fmt.Print(strings.Repeat(" ", height-i))

		// 打印星号
		fmt.Println(strings.Repeat("* ", i))
	}
}

// 生成九宫格
func generateGrid(size int) {
	for i := 0; i < size; i++ {
		// 打印水平分隔线
		if i > 0 {
			fmt.Println(strings.Repeat("-----", size))
		}

		for j := 0; j < size; j++ {
			fmt.Printf(" %d,%d ", i, j)

			// 打印垂直分隔线
			if j < size-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("=== for 循环示例 ===")
	ForLoopsExample()
	fmt.Println("=== for 循环示例结束 ===")
}
