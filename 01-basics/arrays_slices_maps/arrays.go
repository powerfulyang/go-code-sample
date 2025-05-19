package main

import "fmt"

func runArrayExample() {
	fmt.Println("=== Go 数组示例 ===")

	// 数组声明和初始化
	fmt.Println("\n--- 数组声明和初始化 ---")

	// 方式1：声明后赋值
	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Printf("arr1: %v\n", arr1)

	// 方式2：声明时初始化
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr2: %v\n", arr2)

	// 方式3：简短声明
	arr3 := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("arr3: %v\n", arr3)

	// 方式4：让编译器推断长度
	arr4 := [...]int{100, 200, 300}
	fmt.Printf("arr4: %v, 长度: %d\n", arr4, len(arr4))

	// 方式5：指定索引初始化
	arr5 := [5]int{1: 10, 3: 30, 4: 40}
	fmt.Printf("arr5: %v\n", arr5)

	// 数组访问和修改
	fmt.Println("\n--- 数组访问和修改 ---")
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("原始数组: %v\n", numbers)
	fmt.Printf("第一个元素: %d\n", numbers[0])
	fmt.Printf("最后一个元素: %d\n", numbers[len(numbers)-1])

	// 修改元素
	numbers[2] = 99
	fmt.Printf("修改后: %v\n", numbers)

	// 数组遍历
	fmt.Println("\n--- 数组遍历 ---")
	fruits := [4]string{"apple", "banana", "orange", "grape"}

	// 方式1：传统for循环
	fmt.Println("传统for循环:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("  索引%d: %s\n", i, fruits[i])
	}

	// 方式2：range循环
	fmt.Println("range循环:")
	for index, value := range fruits {
		fmt.Printf("  索引%d: %s\n", index, value)
	}

	// 方式3：只要值，不要索引
	fmt.Println("只获取值:")
	for _, value := range fruits {
		fmt.Printf("  %s\n", value)
	}

	// 数组比较
	fmt.Println("\n--- 数组比较 ---")
	arr6 := [3]int{1, 2, 3}
	arr7 := [3]int{1, 2, 3}
	arr8 := [3]int{1, 2, 4}

	fmt.Printf("arr6 == arr7: %t\n", arr6 == arr7)
	fmt.Printf("arr6 == arr8: %t\n", arr6 == arr8)

	// 多维数组
	fmt.Println("\n--- 多维数组 ---")
	var matrix [3][3]int

	// 初始化矩阵
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}

	// 打印矩阵
	fmt.Println("3x3矩阵:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}

	// 直接初始化多维数组
	matrix2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("matrix2: %v\n", matrix2)

	// 数组作为函数参数
	fmt.Println("\n--- 数组作为函数参数 ---")
	scores := [5]int{85, 92, 78, 96, 88}
	avg := calculateAverage(scores)
	fmt.Printf("成绩: %v\n", scores)
	fmt.Printf("平均分: %.2f\n", avg)

	// 数组的零值
	fmt.Println("\n--- 数组的零值 ---")
	var zeroArray [3]int
	var zeroStrings [3]string
	var zeroBools [3]bool

	fmt.Printf("int数组零值: %v\n", zeroArray)
	fmt.Printf("string数组零值: %v\n", zeroStrings)
	fmt.Printf("bool数组零值: %v\n", zeroBools)
}

// calculateAverage 计算数组的平均值
// 注意：数组作为参数时是值传递，会复制整个数组
func calculateAverage(scores [5]int) float64 {
	sum := 0
	for _, score := range scores {
		sum += score
	}
	return float64(sum) / float64(len(scores))
}
