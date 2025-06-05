package functions

import (
	"fmt"
	"math"
	"strings"
)

// 基本函数示例
func BasicFunctions() {
	fmt.Println("=== 基本函数示例 ===")

	// 调用各种函数
	greet("张三")
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	area := calculateArea(5.0)
	fmt.Printf("半径为 5 的圆的面积: %.2f\n", area)
}

// 简单的问候函数
func greet(name string) {
	fmt.Printf("你好, %s!\n", name)
}

// 带返回值的函数
func add(a, b int) int {
	return a + b
}

// 计算圆的面积
func calculateArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// 多返回值函数示例
func MultipleReturnValues() {
	fmt.Println("\n=== 多返回值函数示例 ===")

	// 调用多返回值函数
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", quotient, remainder)

	// 带错误处理的函数
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	// 除零错误
	_, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	// 忽略某个返回值
	name, _ := getNameAndAge()
	fmt.Printf("姓名: %s\n", name)
}

// 返回商和余数
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// 安全除法，返回结果和错误
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// 返回姓名和年龄
func getNameAndAge() (string, int) {
	return "李四", 25
}

// 命名返回值示例
func NamedReturnValues() {
	fmt.Println("\n=== 命名返回值示例 ===")

	// 调用命名返回值函数
	min, max := findMinMax([]int{3, 7, 1, 9, 2, 8})
	fmt.Printf("最小值: %d, 最大值: %d\n", min, max)

	// 计算矩形信息
	area, perimeter := rectangleInfo(5.0, 3.0)
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", area, perimeter)
}

// 使用命名返回值查找最小值和最大值
func findMinMax(numbers []int) (min, max int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return // 裸返回
}

// 计算矩形的面积和周长
func rectangleInfo(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

// 可变参数函数示例
func VariadicFunctions() {
	fmt.Println("\n=== 可变参数函数示例 ===")

	// 调用可变参数函数
	sum1 := sum(1, 2, 3, 4, 5)
	fmt.Printf("1+2+3+4+5 = %d\n", sum1)

	sum2 := sum(10, 20)
	fmt.Printf("10+20 = %d\n", sum2)

	// 传递切片
	numbers := []int{1, 2, 3, 4, 5}
	sum3 := sum(numbers...)
	fmt.Printf("切片求和 = %d\n", sum3)

	// 格式化输出
	formatPrint("用户信息", "姓名", "张三", "年龄", 25, "城市", "北京")
}

// 可变参数求和函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 格式化打印函数
func formatPrint(title string, pairs ...interface{}) {
	fmt.Printf("=== %s ===\n", title)
	for i := 0; i < len(pairs); i += 2 {
		if i+1 < len(pairs) {
			fmt.Printf("%v: %v\n", pairs[i], pairs[i+1])
		}
	}
}

// 高阶函数示例
func HigherOrderFunctions() {
	fmt.Println("\n=== 高阶函数示例 ===")

	// 函数作为参数
	numbers := []int{1, 2, 3, 4, 5}

	doubled := mapInts(numbers, func(x int) int { return x * 2 })
	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("翻倍后: %v\n", doubled)

	evens := filterInts(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("偶数: %v\n", evens)

	total := reduceInts(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("求和: %d\n", total)

	// 函数作为返回值
	multiplier := createMultiplier(3)
	result := multiplier(10)
	fmt.Printf("3 × 10 = %d\n", result)
}

// 映射函数
func mapInts(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 过滤函数
func filterInts(slice []int, fn func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// 归约函数
func reduceInts(slice []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// 创建乘法器函数
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 闭包示例
func Closures() {
	fmt.Println("\n=== 闭包示例 ===")

	// 计数器闭包
	counter := createCounter()
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())

	// 另一个计数器实例
	counter2 := createCounter()
	fmt.Printf("计数器2: %d\n", counter2())

	// 累加器闭包
	adder := createAdder(10)
	fmt.Printf("累加器 (初始值 10): %d\n", adder(5))
	fmt.Printf("累加器: %d\n", adder(3))

	// 字符串处理器
	processor := createStringProcessor("前缀-", "-后缀")
	result := processor("内容")
	fmt.Printf("处理结果: %s\n", result)
}

// 创建计数器闭包
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 创建累加器闭包
func createAdder(initial int) func(int) int {
	sum := initial
	return func(x int) int {
		sum += x
		return sum
	}
}

// 创建字符串处理器闭包
func createStringProcessor(prefix, suffix string) func(string) string {
	return func(content string) string {
		return prefix + content + suffix
	}
}

// 递归函数示例
func RecursiveFunctions() {
	fmt.Println("\n=== 递归函数示例 ===")

	// 阶乘
	n := 5
	fact := factorial(n)
	fmt.Printf("%d! = %d\n", n, fact)

	// 斐波那契数列
	fmt.Printf("斐波那契数列前 10 项: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()

	// 二分查找
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	index := binarySearch(sortedArray, target, 0, len(sortedArray)-1)
	if index != -1 {
		fmt.Printf("在数组 %v 中找到 %d，索引为 %d\n", sortedArray, target, index)
	} else {
		fmt.Printf("在数组中未找到 %d\n", target)
	}
}

// 递归计算阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 递归计算斐波那契数
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 递归二分查找
func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1)
	} else {
		return binarySearch(arr, target, mid+1, right)
	}
}

// 匿名函数示例
func AnonymousFunctions() {
	fmt.Println("\n=== 匿名函数示例 ===")

	// 立即执行的匿名函数
	result := func(a, b int) int {
		return a * b
	}(5, 6)
	fmt.Printf("匿名函数结果: %d\n", result)

	// 将匿名函数赋值给变量
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("5 的平方: %d\n", square(5))

	// 在切片中使用匿名函数
	operations := []func(int, int) int{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a - b },
		func(a, b int) int { return a * b },
		func(a, b int) int { return a / b },
	}

	a, b := 20, 4
	opNames := []string{"加法", "减法", "乘法", "除法"}
	for i, op := range operations {
		result := op(a, b)
		fmt.Printf("%s: %d %s %d = %d\n", opNames[i], a,
			[]string{"+", "-", "×", "÷"}[i], b, result)
	}
}

// 实际应用示例
func PracticalExamples() {
	fmt.Println("\n=== 实际应用示例 ===")

	// 字符串处理工具
	text := "  Hello, World!  "
	processed := processString(text,
		strings.TrimSpace,
		strings.ToLower,
		func(s string) string { return strings.ReplaceAll(s, "world", "go") },
	)
	fmt.Printf("原文本: '%s'\n", text)
	fmt.Printf("处理后: '%s'\n", processed)

	// 数学计算器
	calculator := createCalculator()
	fmt.Printf("计算器: 5 + 3 = %.2f\n", calculator("add", 5, 3))
	fmt.Printf("计算器: 10 - 4 = %.2f\n", calculator("subtract", 10, 4))
	fmt.Printf("计算器: 6 × 7 = %.2f\n", calculator("multiply", 6, 7))
	fmt.Printf("计算器: 15 ÷ 3 = %.2f\n", calculator("divide", 15, 3))

	// 数据验证
	users := []map[string]interface{}{
		{"name": "张三", "age": 25, "email": "zhangsan@example.com"},
		{"name": "", "age": 17, "email": "invalid-email"},
		{"name": "李四", "age": 30, "email": "lisi@example.com"},
	}

	for i, user := range users {
		if err := validateUser(user); err != nil {
			fmt.Printf("用户 %d 验证失败: %v\n", i+1, err)
		} else {
			fmt.Printf("用户 %d 验证通过\n", i+1)
		}
	}
}

// 字符串处理管道
func processString(input string, processors ...func(string) string) string {
	result := input
	for _, processor := range processors {
		result = processor(result)
	}
	return result
}

// 创建计算器
func createCalculator() func(string, float64, float64) float64 {
	operations := map[string]func(float64, float64) float64{
		"add":      func(a, b float64) float64 { return a + b },
		"subtract": func(a, b float64) float64 { return a - b },
		"multiply": func(a, b float64) float64 { return a * b },
		"divide":   func(a, b float64) float64 { return a / b },
	}

	return func(op string, a, b float64) float64 {
		if fn, exists := operations[op]; exists {
			return fn(a, b)
		}
		return 0
	}
}

// 用户数据验证
func validateUser(user map[string]interface{}) error {
	validators := []func(map[string]interface{}) error{
		validateName,
		validateAge,
		validateEmail,
	}

	for _, validator := range validators {
		if err := validator(user); err != nil {
			return err
		}
	}
	return nil
}

func validateName(user map[string]interface{}) error {
	name, ok := user["name"].(string)
	if !ok || strings.TrimSpace(name) == "" {
		return fmt.Errorf("姓名不能为空")
	}
	return nil
}

func validateAge(user map[string]interface{}) error {
	age, ok := user["age"].(int)
	if !ok || age < 18 {
		return fmt.Errorf("年龄必须大于等于 18")
	}
	return nil
}

func validateEmail(user map[string]interface{}) error {
	email, ok := user["email"].(string)
	if !ok || !strings.Contains(email, "@") {
		return fmt.Errorf("邮箱格式不正确")
	}
	return nil
}
