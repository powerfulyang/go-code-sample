package main

import (
	"fmt"
	"sort"
)

func runMapExample() {
	fmt.Println("=== Go 映射示例 ===")

	// 映射声明和初始化
	fmt.Println("\n--- 映射声明和初始化 ---")

	// 方式1：声明空映射
	var map1 map[string]int
	fmt.Printf("空映射: %v, 长度: %d, 是否为nil: %t\n", map1, len(map1), map1 == nil)

	// 方式2：使用make创建映射
	map2 := make(map[string]int)
	fmt.Printf("make映射: %v, 长度: %d, 是否为nil: %t\n", map2, len(map2), map2 == nil)

	// 方式3：映射字面量
	map3 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	fmt.Printf("字面量映射: %v, 长度: %d\n", map3, len(map3))

	// 映射操作
	fmt.Println("\n--- 映射操作 ---")
	scores := make(map[string]int)

	// 添加元素
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("添加元素后: %v\n", scores)

	// 读取元素
	fmt.Printf("Alice的分数: %d\n", scores["Alice"])
	fmt.Printf("不存在的键: %d\n", scores["David"]) // 返回零值

	// 检查键是否存在
	fmt.Println("\n--- 检查键是否存在 ---")
	if score, exists := scores["Alice"]; exists {
		fmt.Printf("Alice存在，分数: %d\n", score)
	} else {
		fmt.Println("Alice不存在")
	}

	if score, exists := scores["David"]; exists {
		fmt.Printf("David存在，分数: %d\n", score)
	} else {
		fmt.Println("David不存在")
	}

	// 修改元素
	fmt.Println("\n--- 修改元素 ---")
	fmt.Printf("修改前Bob的分数: %d\n", scores["Bob"])
	scores["Bob"] = 90
	fmt.Printf("修改后Bob的分数: %d\n", scores["Bob"])

	// 删除元素
	fmt.Println("\n--- 删除元素 ---")
	fmt.Printf("删除前: %v\n", scores)
	delete(scores, "Charlie")
	fmt.Printf("删除Charlie后: %v\n", scores)

	// 删除不存在的键（安全操作）
	delete(scores, "NonExistent")
	fmt.Printf("删除不存在的键后: %v\n", scores)

	// 映射遍历
	fmt.Println("\n--- 映射遍历 ---")
	fruits := map[string]int{
		"apple":      10,
		"banana":     5,
		"orange":     8,
		"grape":      12,
		"watermelon": 3,
	}

	// 遍历键值对
	fmt.Println("遍历键值对:")
	for fruit, count := range fruits {
		fmt.Printf("  %s: %d\n", fruit, count)
	}

	// 只遍历键
	fmt.Println("只遍历键:")
	for fruit := range fruits {
		fmt.Printf("  %s\n", fruit)
	}

	// 只遍历值
	fmt.Println("只遍历值:")
	for _, count := range fruits {
		fmt.Printf("  %d\n", count)
	}

	// 有序遍历（映射本身是无序的）
	fmt.Println("\n--- 有序遍历 ---")
	var keys []string
	for fruit := range fruits {
		keys = append(keys, fruit)
	}
	sort.Strings(keys)

	fmt.Println("按键排序遍历:")
	for _, fruit := range keys {
		fmt.Printf("  %s: %d\n", fruit, fruits[fruit])
	}

	// 映射作为函数参数
	fmt.Println("\n--- 映射作为函数参数 ---")
	inventory := map[string]int{
		"laptop":   10,
		"mouse":    25,
		"keyboard": 15,
	}

	fmt.Printf("原始库存: %v\n", inventory)
	total := calculateTotal(inventory)
	fmt.Printf("总数量: %d\n", total)

	// 修改映射
	updateInventory(inventory, "laptop", 5)
	fmt.Printf("更新后库存: %v\n", inventory)

	// 嵌套映射
	fmt.Println("\n--- 嵌套映射 ---")
	students := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"English": 87,
			"Science": 92,
		},
		"Bob": {
			"Math":    78,
			"English": 85,
			"Science": 88,
		},
	}

	fmt.Println("学生成绩:")
	for student, subjects := range students {
		fmt.Printf("%s:\n", student)
		for subject, score := range subjects {
			fmt.Printf("  %s: %d\n", subject, score)
		}
	}

	// 添加新学生
	students["Charlie"] = make(map[string]int)
	students["Charlie"]["Math"] = 90
	students["Charlie"]["English"] = 93
	students["Charlie"]["Science"] = 89

	fmt.Printf("添加Charlie后: %v\n", students)

	// 映射的零值和初始化
	fmt.Println("\n--- 映射的零值和初始化 ---")
	var nilMap map[string]int
	fmt.Printf("nil映射: %v, 是否为nil: %t\n", nilMap, nilMap == nil)

	// 不能向nil映射写入，会panic
	// nilMap["key"] = 1 // 这会导致panic

	// 正确的初始化方式
	nilMap = make(map[string]int)
	nilMap["key"] = 1
	fmt.Printf("初始化后: %v\n", nilMap)

	// 映射比较
	fmt.Println("\n--- 映射比较 ---")
	// 映射不能直接比较，只能与nil比较
	map4 := make(map[string]int)
	map5 := make(map[string]int)

	fmt.Printf("map4 == nil: %t\n", map4 == nil)
	fmt.Printf("map5 == nil: %t\n", map5 == nil)
	// fmt.Printf("map4 == map5: %t\n", map4 == map5) // 编译错误

	// 如果需要比较映射内容，需要手动实现
	fmt.Printf("映射内容相等: %t\n", mapsEqual(map4, map5))

	map4["a"] = 1
	map5["a"] = 1
	fmt.Printf("添加相同元素后相等: %t\n", mapsEqual(map4, map5))
}

// calculateTotal 计算映射中所有值的总和
func calculateTotal(inventory map[string]int) int {
	total := 0
	for _, count := range inventory {
		total += count
	}
	return total
}

// updateInventory 更新库存
// 注意：映射作为参数时是引用传递，修改会影响原映射
func updateInventory(inventory map[string]int, item string, newCount int) {
	inventory[item] = newCount
}

// mapsEqual 比较两个映射是否相等
func mapsEqual(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, exists := map2[key]; !exists || value1 != value2 {
			return false
		}
	}

	return true
}
