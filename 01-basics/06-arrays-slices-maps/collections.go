package collections

import (
	"fmt"
	"sort"
	"strings"
)

// 数组示例
func ArrayExamples() {
	fmt.Println("=== 数组示例 ===")

	// 声明和初始化数组
	var numbers [5]int                      // 零值数组
	scores := [3]int{95, 87, 92}            // 初始化数组
	fruits := [...]string{"苹果", "香蕉", "橙子"} // 自动推断长度

	fmt.Printf("零值数组: %v\n", numbers)
	fmt.Printf("分数数组: %v\n", scores)
	fmt.Printf("水果数组: %v\n", fruits)

	// 访问和修改数组元素
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Printf("修改后的数组: %v\n", numbers)
	fmt.Printf("第一个元素: %d\n", numbers[0])
	fmt.Printf("数组长度: %d\n", len(numbers))

	// 遍历数组
	fmt.Println("遍历数组:")
	for i, value := range scores {
		fmt.Printf("  索引 %d: %d\n", i, value)
	}

	// 数组比较
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3)
}

// 切片示例
func SliceExamples() {
	fmt.Println("\n=== 切片示例 ===")

	// 创建切片的不同方式
	var slice1 []int               // 零值切片
	slice2 := []int{1, 2, 3, 4, 5} // 字面量创建
	slice3 := make([]int, 3)       // 使用 make 创建，长度为 3
	slice4 := make([]int, 3, 5)    // 长度为 3，容量为 5

	fmt.Printf("零值切片: %v (长度: %d, 容量: %d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("字面量切片: %v (长度: %d, 容量: %d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("make切片1: %v (长度: %d, 容量: %d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("make切片2: %v (长度: %d, 容量: %d)\n", slice4, len(slice4), cap(slice4))

	// 切片操作
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("原始切片: %v\n", original)

	// 切片的切片
	sub1 := original[2:5] // [2, 3, 4]
	sub2 := original[:3]  // [0, 1, 2]
	sub3 := original[7:]  // [7, 8, 9]
	sub4 := original[:]   // 完整副本

	fmt.Printf("original[2:5]: %v\n", sub1)
	fmt.Printf("original[:3]: %v\n", sub2)
	fmt.Printf("original[7:]: %v\n", sub3)
	fmt.Printf("original[:]: %v\n", sub4)

	// 添加元素
	numbers := []int{1, 2, 3}
	fmt.Printf("原始: %v\n", numbers)

	numbers = append(numbers, 4)
	fmt.Printf("添加一个元素: %v\n", numbers)

	numbers = append(numbers, 5, 6, 7)
	fmt.Printf("添加多个元素: %v\n", numbers)

	more := []int{8, 9, 10}
	numbers = append(numbers, more...)
	fmt.Printf("添加另一个切片: %v\n", numbers)
}

// 切片高级操作
func SliceAdvancedOperations() {
	fmt.Println("\n=== 切片高级操作 ===")

	// 复制切片
	original := []int{1, 2, 3, 4, 5}
	copied := make([]int, len(original))
	copy(copied, original)

	fmt.Printf("原始切片: %v\n", original)
	fmt.Printf("复制切片: %v\n", copied)

	// 修改复制的切片不影响原始切片
	copied[0] = 100
	fmt.Printf("修改后 - 原始: %v, 复制: %v\n", original, copied)

	// 删除元素
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("删除前: %v\n", slice)

	// 删除索引为 2 的元素
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("删除索引 %d 后: %v\n", index, slice)

	// 插入元素
	slice = []int{1, 2, 4, 5}
	fmt.Printf("插入前: %v\n", slice)

	// 在索引 2 处插入 3
	index = 2
	value := 3
	slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
	fmt.Printf("在索引 %d 插入 %d 后: %v\n", index, value, slice)

	// 切片排序
	unsorted := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("排序前: %v\n", unsorted)

	sort.Ints(unsorted)
	fmt.Printf("排序后: %v\n", unsorted)

	// 字符串切片排序
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("字符串排序前: %v\n", words)

	sort.Strings(words)
	fmt.Printf("字符串排序后: %v\n", words)
}

// 映射示例
func MapExamples() {
	fmt.Println("\n=== 映射示例 ===")

	// 创建映射的不同方式
	var map1 map[string]int      // 零值映射
	map2 := make(map[string]int) // 使用 make 创建
	map3 := map[string]int{      // 字面量创建
		"张三": 95,
		"李四": 87,
		"王五": 92,
	}

	fmt.Printf("零值映射: %v\n", map1)
	fmt.Printf("make映射: %v\n", map2)
	fmt.Printf("字面量映射: %v\n", map3)

	// 添加和修改元素
	map2["Go"] = 100
	map2["Python"] = 95
	map2["Java"] = 90

	fmt.Printf("添加元素后: %v\n", map2)

	// 访问元素
	score := map3["张三"]
	fmt.Printf("张三的分数: %d\n", score)

	// 检查键是否存在
	if score, exists := map3["赵六"]; exists {
		fmt.Printf("赵六的分数: %d\n", score)
	} else {
		fmt.Println("赵六不在映射中")
	}

	// 删除元素
	delete(map3, "李四")
	fmt.Printf("删除李四后: %v\n", map3)

	// 遍历映射
	fmt.Println("遍历映射:")
	for name, score := range map3 {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// 获取所有键
	var keys []string
	for key := range map3 {
		keys = append(keys, key)
	}
	fmt.Printf("所有键: %v\n", keys)

	// 获取所有值
	var values []int
	for _, value := range map3 {
		values = append(values, value)
	}
	fmt.Printf("所有值: %v\n", values)
}

// 复杂数据结构示例
func ComplexDataStructures() {
	fmt.Println("\n=== 复杂数据结构示例 ===")

	// 切片的切片（二维切片）
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("二维切片:")
	for i, row := range matrix {
		fmt.Printf("  行 %d: %v\n", i, row)
	}

	// 映射的切片
	students := []map[string]interface{}{
		{"name": "张三", "age": 20, "score": 95},
		{"name": "李四", "age": 21, "score": 87},
		{"name": "王五", "age": 19, "score": 92},
	}

	fmt.Println("学生信息:")
	for i, student := range students {
		fmt.Printf("  学生 %d: %v\n", i+1, student)
	}

	// 切片的映射
	grades := map[string][]int{
		"数学": {95, 87, 92, 78, 85},
		"英语": {88, 92, 85, 90, 87},
		"物理": {90, 85, 88, 92, 89},
	}

	fmt.Println("各科成绩:")
	for subject, scores := range grades {
		fmt.Printf("  %s: %v\n", subject, scores)
	}

	// 映射的映射
	company := map[string]map[string]interface{}{
		"技术部": {
			"人数":  10,
			"预算":  1000000,
			"负责人": "张经理",
		},
		"销售部": {
			"人数":  15,
			"预算":  800000,
			"负责人": "李经理",
		},
	}

	fmt.Println("公司部门信息:")
	for dept, info := range company {
		fmt.Printf("  %s:\n", dept)
		for key, value := range info {
			fmt.Printf("    %s: %v\n", key, value)
		}
	}
}

// 实际应用示例
func PracticalExamples() {
	fmt.Println("\n=== 实际应用示例 ===")

	// 单词计数
	text := "go is great go is simple go is powerful"
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Printf("文本: %s\n", text)
	fmt.Println("单词计数:")
	for word, count := range wordCount {
		fmt.Printf("  %s: %d\n", word, count)
	}

	// 学生成绩管理
	type Student struct {
		Name   string
		Scores []int
	}

	students := []Student{
		{"张三", []int{95, 87, 92}},
		{"李四", []int{88, 90, 85}},
		{"王五", []int{92, 89, 94}},
	}

	fmt.Println("学生成绩统计:")
	for _, student := range students {
		total := 0
		for _, score := range student.Scores {
			total += score
		}
		average := float64(total) / float64(len(student.Scores))
		fmt.Printf("  %s: 总分 %d, 平均分 %.2f\n", student.Name, total, average)
	}

	// 购物车示例
	type Product struct {
		Name  string
		Price float64
	}

	cart := map[Product]int{
		{"苹果", 5.0}: 3,
		{"香蕉", 3.0}: 5,
		{"橙子", 4.0}: 2,
	}

	fmt.Println("购物车:")
	total := 0.0
	for product, quantity := range cart {
		subtotal := product.Price * float64(quantity)
		total += subtotal
		fmt.Printf("  %s: %.2f × %d = %.2f\n",
			product.Name, product.Price, quantity, subtotal)
	}
	fmt.Printf("总计: %.2f\n", total)

	// 数据去重
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	seen := make(map[int]bool)
	var unique []int

	for _, num := range numbers {
		if !seen[num] {
			seen[num] = true
			unique = append(unique, num)
		}
	}

	fmt.Printf("原始数据: %v\n", numbers)
	fmt.Printf("去重后: %v\n", unique)

	// 分组操作
	people := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	groups := make(map[int][]string)

	for i, person := range people {
		groupIndex := i % 3 // 分成3组
		groups[groupIndex] = append(groups[groupIndex], person)
	}

	fmt.Println("分组结果:")
	for groupIndex, members := range groups {
		fmt.Printf("  组 %d: %v\n", groupIndex, members)
	}
}
