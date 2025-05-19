package main

import (
	"fmt"
	"strings"
)

// 1. 基本可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 2. 带普通参数的可变参数函数
// 可变参数必须是函数的最后一个参数
func joinWithSeparator(separator string, parts ...string) string {
	return strings.Join(parts, separator)
}

// 3. 可变参数类型可以是任何类型
func printAny(args ...interface{}) {
	for i, arg := range args {
		fmt.Printf("第%d个参数: %v (类型: %T)\n", i+1, arg, arg)
	}
}

// 4. 传递可变参数切片
func calculateAverage(numbers ...int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	return float64(sum(numbers...)) / float64(len(numbers))
}

// 5. 封装日志函数
func logInfo(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

func logError(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}

// 6. 参数处理函数
func processStrings(operation func(string) string, values ...string) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = operation(v)
	}
	return result
}

func VariadicFunctionsExample() {
	fmt.Println("=== 可变参数函数示例 ===")

	// 1. 基本可变参数函数
	fmt.Println("\n基本可变参数函数:")
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("1 + 2 + 3 + 4 + 5 = %d\n", total)

	// 不传递参数
	total = sum()
	fmt.Printf("不传参数: %d\n", total)

	// 2. 带普通参数的可变参数函数
	fmt.Println("\n带普通参数的可变参数函数:")
	result := joinWithSeparator(", ", "苹果", "香蕉", "橙子")
	fmt.Printf("连接结果: %s\n", result)

	// 3. 任意类型的可变参数
	fmt.Println("\n任意类型的可变参数:")
	printAny(123, "hello", true, 3.14)

	// 4. 传递可变参数切片
	fmt.Println("\n传递可变参数切片:")
	scores := []int{85, 90, 75, 95, 80}

	// 使用...将切片展开为可变参数
	avg := calculateAverage(scores...)
	fmt.Printf("分数: %v, 平均分: %.2f\n", scores, avg)

	// 也可以部分使用切片，部分使用单独的值
	newScores := []int{65, 70}
	avg = calculateAverage(100, 95, newScores...)
	fmt.Printf("新分数集: [100, 95, %v], 平均分: %.2f\n", newScores, avg)

	// 5. 日志函数示例
	fmt.Println("\n日志函数示例:")
	username := "admin"
	loginCount := 5
	logInfo("用户 %s 已登录，这是第 %d 次登录", username, loginCount)

	err := "数据库连接错误"
	logError("无法完成操作: %s", err)

	// 6. 参数处理函数
	fmt.Println("\n参数处理函数:")
	fruits := []string{"apple", "banana", "orange"}

	// 转大写
	upperFruits := processStrings(strings.ToUpper, fruits...)
	fmt.Printf("转大写: %v\n", upperFruits)

	// 自定义处理函数
	addPrefix := func(s string) string {
		return "fruit-" + s
	}
	prefixedFruits := processStrings(addPrefix, fruits...)
	fmt.Printf("添加前缀: %v\n", prefixedFruits)
}
