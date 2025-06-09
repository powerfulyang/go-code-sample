package thirdparty

import (
	"fmt"
	"sort"
	"strings"
)

// 实现类似 samber/lo 的函数式编程工具库
// 这里我们自己实现这些功能来学习函数式编程概念

// Map 映射函数 - 将切片中的每个元素通过函数转换
func Map[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

// Filter 过滤函数 - 保留满足条件的元素
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce 归约函数 - 将切片归约为单个值
func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// Find 查找函数 - 查找第一个满足条件的元素
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	var zero T
	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}
	return zero, false
}

// FindIndex 查找索引 - 查找第一个满足条件的元素的索引
func FindIndex[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Contains 检查是否包含元素
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// ContainsBy 通过函数检查是否包含元素
func ContainsBy[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Unique 去重 - 移除重复元素
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// UniqueBy 通过函数去重
func UniqueBy[T any, U comparable](slice []T, keyFunc func(T) U) []T {
	seen := make(map[U]bool)
	var result []T
	for _, v := range slice {
		key := keyFunc(v)
		if !seen[key] {
			seen[key] = true
			result = append(result, v)
		}
	}
	return result
}

// Reverse 反转切片
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// Chunk 分块 - 将切片分成指定大小的块
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	var result [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

// Flatten 扁平化 - 将二维切片扁平化为一维
func Flatten[T any](slices [][]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// GroupBy 分组 - 按照函数结果分组
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := keyFunc(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Partition 分区 - 将切片分为满足和不满足条件的两部分
func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	var trueSlice, falseSlice []T
	for _, v := range slice {
		if predicate(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return trueSlice, falseSlice
}

// Ordered 可排序类型约束
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// SortBy 排序 - 按照函数结果排序
func SortBy[T any, U Ordered](slice []T, keyFunc func(T) U) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	sort.Slice(result, func(i, j int) bool {
		return keyFunc(result[i]) < keyFunc(result[j])
	})

	return result
}

// MinBy 最小值 - 按照函数结果找最小值
func MinBy[T any, U Ordered](slice []T, keyFunc func(T) U) (T, bool) {
	var zero T
	if len(slice) == 0 {
		return zero, false
	}

	min := slice[0]
	minKey := keyFunc(min)

	for _, v := range slice[1:] {
		key := keyFunc(v)
		if key < minKey {
			min = v
			minKey = key
		}
	}

	return min, true
}

// MaxBy 最大值 - 按照函数结果找最大值
func MaxBy[T any, U Ordered](slice []T, keyFunc func(T) U) (T, bool) {
	var zero T
	if len(slice) == 0 {
		return zero, false
	}

	max := slice[0]
	maxKey := keyFunc(max)

	for _, v := range slice[1:] {
		key := keyFunc(v)
		if key > maxKey {
			max = v
			maxKey = key
		}
	}

	return max, true
}

// Sum 求和 - 数值类型求和
func Sum[T Numeric](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumBy 按函数求和
func SumBy[T any, U Numeric](slice []T, keyFunc func(T) U) U {
	var sum U
	for _, v := range slice {
		sum += keyFunc(v)
	}
	return sum
}

// Numeric 数值类型约束
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Every 检查是否所有元素都满足条件
func Every[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Some 检查是否有元素满足条件
func Some[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Count 计算满足条件的元素数量
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count
}

// CountBy 按函数计算元素数量
func CountBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K]int {
	result := make(map[K]int)
	for _, v := range slice {
		key := keyFunc(v)
		result[key]++
	}
	return result
}

// Drop 丢弃前n个元素
func Drop[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return []T{}
	}
	if n <= 0 {
		return slice
	}
	return slice[n:]
}

// DropRight 丢弃后n个元素
func DropRight[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return []T{}
	}
	if n <= 0 {
		return slice
	}
	return slice[:len(slice)-n]
}

// Take 取前n个元素
func Take[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return slice
	}
	if n <= 0 {
		return []T{}
	}
	return slice[:n]
}

// TakeRight 取后n个元素
func TakeRight[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return slice
	}
	if n <= 0 {
		return []T{}
	}
	return slice[len(slice)-n:]
}

// 实际应用示例

// User 用户结构体
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	City     string `json:"city"`
	IsActive bool   `json:"is_active"`
	Salary   int    `json:"salary"`
}

// Product 产品结构体
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Rating   float64 `json:"rating"`
}

// FunctionalExamples 函数式编程示例
func FunctionalExamples() {
	fmt.Println("=== 函数式编程工具库示例 ===")

	// 基本数据操作
	fmt.Println("\n🔹 基本数据操作")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("原始数据: %v\n", numbers)

	// Map - 映射操作
	squares := Map(numbers, func(n int) int { return n * n })
	fmt.Printf("平方: %v\n", squares)

	// Filter - 过滤操作
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("偶数: %v\n", evens)

	// Reduce - 归约操作
	sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	fmt.Printf("求和: %d\n", sum)

	// 字符串操作示例
	fmt.Println("\n🔹 字符串操作")
	words := []string{"hello", "world", "go", "programming", "functional"}
	fmt.Printf("原始单词: %v\n", words)

	// 转换为大写
	upperWords := Map(words, strings.ToUpper)
	fmt.Printf("大写: %v\n", upperWords)

	// 过滤长度大于4的单词
	longWords := Filter(words, func(s string) bool { return len(s) > 4 })
	fmt.Printf("长单词: %v\n", longWords)

	// 连接所有单词
	sentence := Reduce(words, "", func(acc, word string) string {
		if acc == "" {
			return word
		}
		return acc + " " + word
	})
	fmt.Printf("连接: %s\n", sentence)

	// 用户数据示例
	fmt.Println("\n🔹 用户数据处理")
	users := []User{
		{ID: 1, Name: "张三", Email: "zhangsan@example.com", Age: 25, City: "北京", IsActive: true, Salary: 8000},
		{ID: 2, Name: "李四", Email: "lisi@example.com", Age: 30, City: "上海", IsActive: true, Salary: 12000},
		{ID: 3, Name: "王五", Email: "wangwu@example.com", Age: 28, City: "北京", IsActive: false, Salary: 9000},
		{ID: 4, Name: "赵六", Email: "zhaoliu@example.com", Age: 35, City: "深圳", IsActive: true, Salary: 15000},
		{ID: 5, Name: "钱七", Email: "qianqi@example.com", Age: 22, City: "上海", IsActive: true, Salary: 6000},
	}

	// 获取所有用户名
	names := Map(users, func(u User) string { return u.Name })
	fmt.Printf("用户名: %v\n", names)

	// 过滤活跃用户
	activeUsers := Filter(users, func(u User) bool { return u.IsActive })
	fmt.Printf("活跃用户数: %d\n", len(activeUsers))

	// 按城市分组
	usersByCity := GroupBy(users, func(u User) string { return u.City })
	fmt.Println("按城市分组:")
	for city, cityUsers := range usersByCity {
		userNames := Map(cityUsers, func(u User) string { return u.Name })
		fmt.Printf("  %s: %v\n", city, userNames)
	}

	// 计算平均薪资
	totalSalary := SumBy(activeUsers, func(u User) int { return u.Salary })
	avgSalary := float64(totalSalary) / float64(len(activeUsers))
	fmt.Printf("活跃用户平均薪资: %.2f\n", avgSalary)

	// 找到薪资最高的用户
	if highestPaid, found := MaxBy(users, func(u User) int { return u.Salary }); found {
		fmt.Printf("薪资最高: %s (%.0f元)\n", highestPaid.Name, float64(highestPaid.Salary))
	}

	// 产品数据示例
	fmt.Println("\n🔹 产品数据处理")
	products := []Product{
		{ID: 1, Name: "iPhone 15", Category: "手机", Price: 5999, Stock: 100, Rating: 4.8},
		{ID: 2, Name: "MacBook Pro", Category: "电脑", Price: 12999, Stock: 50, Rating: 4.9},
		{ID: 3, Name: "iPad Air", Category: "平板", Price: 3999, Stock: 80, Rating: 4.7},
		{ID: 4, Name: "AirPods Pro", Category: "耳机", Price: 1999, Stock: 200, Rating: 4.6},
		{ID: 5, Name: "Apple Watch", Category: "手表", Price: 2999, Stock: 120, Rating: 4.5},
	}

	// 按类别分组
	productsByCategory := GroupBy(products, func(p Product) string { return p.Category })
	fmt.Println("按类别分组:")
	for category, categoryProducts := range productsByCategory {
		productNames := Map(categoryProducts, func(p Product) string { return p.Name })
		fmt.Printf("  %s: %v\n", category, productNames)
	}

	// 高评分产品
	highRatedProducts := Filter(products, func(p Product) bool { return p.Rating >= 4.7 })
	highRatedNames := Map(highRatedProducts, func(p Product) string { return p.Name })
	fmt.Printf("高评分产品: %v\n", highRatedNames)

	// 按价格排序
	sortedByPrice := SortBy(products, func(p Product) float64 { return p.Price })
	fmt.Println("按价格排序:")
	for _, p := range sortedByPrice {
		fmt.Printf("  %s: %.0f元\n", p.Name, p.Price)
	}

	// 数组操作示例
	fmt.Println("\n🔹 数组操作")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	fmt.Printf("原始数据: %v\n", data)

	// 去重
	uniqueData := Unique(data)
	fmt.Printf("去重后: %v\n", uniqueData)

	// 分块
	chunks := Chunk(data, 3)
	fmt.Printf("分块(3): %v\n", chunks)

	// 反转
	reversed := Reverse(data)
	fmt.Printf("反转: %v\n", reversed)

	// 分区
	evens2, odds := Partition(data, func(n int) bool { return n%2 == 0 })
	fmt.Printf("偶数: %v\n", evens2)
	fmt.Printf("奇数: %v\n", odds)

	// 取前5个
	first5 := Take(data, 5)
	fmt.Printf("前5个: %v\n", first5)

	// 丢弃前3个
	dropped := Drop(data, 3)
	fmt.Printf("丢弃前3个: %v\n", dropped)

	// 统计示例
	fmt.Println("\n🔹 统计操作")

	// 检查是否所有数字都是正数
	allPositive := Every(numbers, func(n int) bool { return n > 0 })
	fmt.Printf("所有数字都是正数: %t\n", allPositive)

	// 检查是否有大于5的数字
	hasLargeNumber := Some(numbers, func(n int) bool { return n > 5 })
	fmt.Printf("有大于5的数字: %t\n", hasLargeNumber)

	// 计算偶数数量
	evenCount := Count(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("偶数数量: %d\n", evenCount)

	// 按奇偶性计数
	parityCount := CountBy(numbers, func(n int) string {
		if n%2 == 0 {
			return "偶数"
		}
		return "奇数"
	})
	fmt.Printf("奇偶性统计: %v\n", parityCount)
}
