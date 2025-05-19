package main

import "fmt"

func runSliceExample() {
	fmt.Println("=== Go 切片示例 ===")

	// 切片声明和初始化
	fmt.Println("\n--- 切片声明和初始化 ---")

	// 方式1：声明空切片
	var slice1 []int
	fmt.Printf("空切片: %v, 长度: %d, 容量: %d\n", slice1, len(slice1), cap(slice1))

	// 方式2：使用make创建切片
	slice2 := make([]int, 5)     // 长度为5，容量为5
	slice3 := make([]int, 3, 10) // 长度为3，容量为10
	fmt.Printf("make切片1: %v, 长度: %d, 容量: %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("make切片2: %v, 长度: %d, 容量: %d\n", slice3, len(slice3), cap(slice3))

	// 方式3：切片字面量
	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("字面量切片: %v, 长度: %d, 容量: %d\n", slice4, len(slice4), cap(slice4))

	// 方式4：从数组创建切片
	arr := [5]int{10, 20, 30, 40, 50}
	slice5 := arr[1:4] // 从索引1到3（不包括4）
	fmt.Printf("从数组创建: %v, 长度: %d, 容量: %d\n", slice5, len(slice5), cap(slice5))

	// 切片操作
	fmt.Println("\n--- 切片操作 ---")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("原始切片: %v\n", numbers)

	// 切片截取
	fmt.Printf("numbers[2:5]: %v\n", numbers[2:5]) // 索引2到4
	fmt.Printf("numbers[:3]: %v\n", numbers[:3])   // 从开始到索引2
	fmt.Printf("numbers[3:]: %v\n", numbers[3:])   // 从索引3到结束
	fmt.Printf("numbers[:]: %v\n", numbers[:])     // 完整切片

	// append操作
	fmt.Println("\n--- append操作 ---")
	fruits := []string{"apple", "banana"}
	fmt.Printf("原始: %v, 长度: %d, 容量: %d\n", fruits, len(fruits), cap(fruits))

	// 添加单个元素
	fruits = append(fruits, "orange")
	fmt.Printf("添加orange: %v, 长度: %d, 容量: %d\n", fruits, len(fruits), cap(fruits))

	// 添加多个元素
	fruits = append(fruits, "grape", "kiwi")
	fmt.Printf("添加多个: %v, 长度: %d, 容量: %d\n", fruits, len(fruits), cap(fruits))

	// 添加另一个切片
	moreFruits := []string{"mango", "pineapple"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("添加切片: %v, 长度: %d, 容量: %d\n", fruits, len(fruits), cap(fruits))

	// copy操作
	fmt.Println("\n--- copy操作 ---")
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(source))

	copied := copy(dest, source)
	fmt.Printf("源切片: %v\n", source)
	fmt.Printf("目标切片: %v\n", dest)
	fmt.Printf("复制了%d个元素\n", copied)

	// 修改目标切片不会影响源切片
	dest[0] = 99
	fmt.Printf("修改后源切片: %v\n", source)
	fmt.Printf("修改后目标切片: %v\n", dest)

	// 切片的零值
	fmt.Println("\n--- 切片的零值 ---")
	var nilSlice []int
	fmt.Printf("nil切片: %v, 长度: %d, 容量: %d, 是否为nil: %t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)

	// 切片遍历
	fmt.Println("\n--- 切片遍历 ---")
	colors := []string{"red", "green", "blue", "yellow"}

	// 使用range
	fmt.Println("使用range:")
	for index, color := range colors {
		fmt.Printf("  索引%d: %s\n", index, color)
	}

	// 只要值
	fmt.Println("只要值:")
	for _, color := range colors {
		fmt.Printf("  %s\n", color)
	}

	// 只要索引
	fmt.Println("只要索引:")
	for index := range colors {
		fmt.Printf("  索引%d\n", index)
	}

	// 切片作为函数参数
	fmt.Println("\n--- 切片作为函数参数 ---")
	scores := []int{85, 92, 78, 96, 88}
	fmt.Printf("原始成绩: %v\n", scores)

	avg := calculateSliceAverage(scores)
	fmt.Printf("平均分: %.2f\n", avg)

	// 修改切片
	addBonus(scores, 5)
	fmt.Printf("加分后: %v\n", scores)

	// 多维切片
	fmt.Println("\n--- 多维切片 ---")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("矩阵:")
	for i, row := range matrix {
		fmt.Printf("行%d: %v\n", i, row)
	}

	// 动态创建多维切片
	rows, cols := 3, 4
	dynamicMatrix := make([][]int, rows)
	for i := range dynamicMatrix {
		dynamicMatrix[i] = make([]int, cols)
		for j := range dynamicMatrix[i] {
			dynamicMatrix[i][j] = i*cols + j + 1
		}
	}

	fmt.Println("动态矩阵:")
	for i, row := range dynamicMatrix {
		fmt.Printf("行%d: %v\n", i, row)
	}

	// 切片容量扩展示例
	fmt.Println("\n--- 切片容量扩展 ---")
	s := make([]int, 0, 1)
	fmt.Printf("初始: 长度=%d, 容量=%d\n", len(s), cap(s))

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("添加%d后: 长度=%d, 容量=%d\n", i, len(s), cap(s))
	}
}

// calculateSliceAverage 计算切片的平均值
func calculateSliceAverage(scores []int) float64 {
	if len(scores) == 0 {
		return 0
	}

	sum := 0
	for _, score := range scores {
		sum += score
	}
	return float64(sum) / float64(len(scores))
}

// addBonus 给切片中的每个分数加上奖励分
// 注意：切片作为参数时是引用传递，修改会影响原切片
func addBonus(scores []int, bonus int) {
	for i := range scores {
		scores[i] += bonus
	}
}
