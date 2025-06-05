package controlflow

import (
	"fmt"
	"math/rand"
	"time"
)

// 条件语句示例
func ConditionalStatements() {
	fmt.Println("=== 条件语句示例 ===")

	// 基本 if 语句
	age := 20
	if age >= 18 {
		fmt.Printf("年龄 %d，已成年\n", age)
	}

	// if-else 语句
	score := 85
	if score >= 90 {
		fmt.Printf("分数 %d，等级：优秀\n", score)
	} else if score >= 80 {
		fmt.Printf("分数 %d，等级：良好\n", score)
	} else if score >= 70 {
		fmt.Printf("分数 %d，等级：中等\n", score)
	} else if score >= 60 {
		fmt.Printf("分数 %d，等级：及格\n", score)
	} else {
		fmt.Printf("分数 %d，等级：不及格\n", score)
	}

	// 带初始化的 if 语句
	if num := rand.Intn(100); num > 50 {
		fmt.Printf("随机数 %d 大于 50\n", num)
	} else {
		fmt.Printf("随机数 %d 小于等于 50\n", num)
	}
}

// switch 语句示例
func SwitchStatements() {
	fmt.Println("\n=== switch 语句示例 ===")

	// 基本 switch 语句
	day := time.Now().Weekday()
	switch day {
	case time.Monday:
		fmt.Println("今天是星期一")
	case time.Tuesday:
		fmt.Println("今天是星期二")
	case time.Wednesday:
		fmt.Println("今天是星期三")
	case time.Thursday:
		fmt.Println("今天是星期四")
	case time.Friday:
		fmt.Println("今天是星期五")
	case time.Saturday, time.Sunday:
		fmt.Println("今天是周末")
	default:
		fmt.Println("未知的星期")
	}

	// 无表达式的 switch
	hour := time.Now().Hour()
	switch {
	case hour < 6:
		fmt.Println("凌晨时分")
	case hour < 12:
		fmt.Println("上午时光")
	case hour < 18:
		fmt.Println("下午时光")
	default:
		fmt.Println("晚上时光")
	}

	// 带 fallthrough 的 switch
	grade := 'B'
	fmt.Printf("成绩 %c: ", grade)
	switch grade {
	case 'A':
		fmt.Print("优秀")
		fallthrough
	case 'B':
		fmt.Print("良好")
		fallthrough
	case 'C':
		fmt.Print("中等")
		fallthrough
	default:
		fmt.Println(" - 需要继续努力")
	}
}

// for 循环示例
func ForLoops() {
	fmt.Println("\n=== for 循环示例 ===")

	// 基本 for 循环
	fmt.Println("基本 for 循环:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  第 %d 次循环\n", i)
	}

	// 类似 while 的 for 循环
	fmt.Println("类似 while 的循环:")
	count := 0
	for count < 3 {
		fmt.Printf("  计数: %d\n", count)
		count++
	}

	// 无限循环（带 break）
	fmt.Println("无限循环（带 break）:")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("  无限循环第 %d 次\n", i)
		i++
	}

	// 带 continue 的循环
	fmt.Println("跳过偶数:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("  奇数: %d\n", i)
	}
}

// range 循环示例
func RangeLoops() {
	fmt.Println("\n=== range 循环示例 ===")

	// 遍历数组
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println("遍历数组:")
	for index, value := range numbers {
		fmt.Printf("  索引 %d: 值 %d\n", index, value)
	}

	// 只要索引
	fmt.Println("只要索引:")
	for index := range numbers {
		fmt.Printf("  索引: %d\n", index)
	}

	// 只要值
	fmt.Println("只要值:")
	for _, value := range numbers {
		fmt.Printf("  值: %d\n", value)
	}

	// 遍历切片
	fruits := []string{"苹果", "香蕉", "橙子"}
	fmt.Println("遍历切片:")
	for i, fruit := range fruits {
		fmt.Printf("  %d: %s\n", i, fruit)
	}

	// 遍历映射
	scores := map[string]int{
		"张三": 95,
		"李四": 87,
		"王五": 92,
	}
	fmt.Println("遍历映射:")
	for name, score := range scores {
		fmt.Printf("  %s: %d 分\n", name, score)
	}

	// 遍历字符串
	text := "Go语言"
	fmt.Println("遍历字符串:")
	for i, char := range text {
		fmt.Printf("  位置 %d: %c\n", i, char)
	}
}

// 嵌套循环示例
func NestedLoops() {
	fmt.Println("\n=== 嵌套循环示例 ===")

	// 打印乘法表
	fmt.Println("九九乘法表:")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%d  ", j, i, i*j)
		}
		fmt.Println()
	}

	// 二维数组遍历
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("二维数组:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

// 循环控制示例
func LoopControl() {
	fmt.Println("\n=== 循环控制示例 ===")

	// 标签和 break
	fmt.Println("使用标签的 break:")
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				fmt.Printf("  在 i=%d, j=%d 时跳出外层循环\n", i, j)
				break outer
			}
			fmt.Printf("  i=%d, j=%d, 乘积=%d\n", i, j, i*j)
		}
	}

	// 标签和 continue
	fmt.Println("使用标签的 continue:")
outer2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Printf("  在 i=%d, j=%d 时跳过外层循环的当前迭代\n", i, j)
				continue outer2
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}
}

// 条件表达式示例
func ConditionalExpressions() {
	fmt.Println("\n=== 条件表达式示例 ===")

	// 使用函数模拟三元运算符
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	a, b := 10, 20
	fmt.Printf("max(%d, %d) = %d\n", a, b, max(a, b))
	fmt.Printf("min(%d, %d) = %d\n", a, b, min(a, b))

	// 条件赋值
	status := "active"
	message := ""
	if status == "active" {
		message = "用户活跃"
	} else {
		message = "用户非活跃"
	}
	fmt.Printf("状态: %s, 消息: %s\n", status, message)
}

// 实际应用示例
func PracticalExamples() {
	fmt.Println("\n=== 实际应用示例 ===")

	// 成绩统计
	scores := []int{95, 87, 92, 78, 85, 90, 88}
	total := 0
	highest := scores[0]
	lowest := scores[0]
	passCount := 0

	for _, score := range scores {
		total += score
		if score > highest {
			highest = score
		}
		if score < lowest {
			lowest = score
		}
		if score >= 60 {
			passCount++
		}
	}

	average := float64(total) / float64(len(scores))
	passRate := float64(passCount) / float64(len(scores)) * 100

	fmt.Printf("成绩统计:\n")
	fmt.Printf("  总分: %d\n", total)
	fmt.Printf("  平均分: %.2f\n", average)
	fmt.Printf("  最高分: %d\n", highest)
	fmt.Printf("  最低分: %d\n", lowest)
	fmt.Printf("  及格人数: %d\n", passCount)
	fmt.Printf("  及格率: %.1f%%\n", passRate)

	// 数字分类
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evens, odds []int

	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}

	fmt.Printf("数字分类:\n")
	fmt.Printf("  偶数: %v\n", evens)
	fmt.Printf("  奇数: %v\n", odds)
}

// 错误处理中的控制流
func ErrorHandlingFlow() {
	fmt.Println("\n=== 错误处理中的控制流 ===")

	// 模拟可能出错的操作
	divide := func(a, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("除数不能为零")
		}
		return a / b, nil
	}

	// 处理多个操作
	operations := [][2]int{{10, 2}, {15, 3}, {20, 0}, {25, 5}}

	for i, op := range operations {
		if result, err := divide(op[0], op[1]); err != nil {
			fmt.Printf("  操作 %d: %d ÷ %d 失败: %v\n", i+1, op[0], op[1], err)
		} else {
			fmt.Printf("  操作 %d: %d ÷ %d = %d\n", i+1, op[0], op[1], result)
		}
	}
}
