package main

import (
	"fmt"
	"log"
	"strings"

	"golang-examples/03-practical-examples/01-package-management/math"
	"golang-examples/03-practical-examples/01-package-management/stringutils"
)

func main() {
	fmt.Println("🚀 Go包管理和模块示例")
	fmt.Println("=" + strings.Repeat("=", 49))

	// 数学计算示例
	mathExamples()

	// 字符串处理示例
	stringExamples()

	// 计算器示例
	calculatorExample()
}

func mathExamples() {
	fmt.Println("\n🔹 数学计算包示例")
	fmt.Println(strings.Repeat("-", 30))

	// 基本运算
	fmt.Printf("加法: %.2f + %.2f = %.2f\n", 10.5, 5.3, math.Add(10.5, 5.3))
	fmt.Printf("减法: %.2f - %.2f = %.2f\n", 10.5, 5.3, math.Subtract(10.5, 5.3))
	fmt.Printf("乘法: %.2f × %.2f = %.2f\n", 10.5, 5.3, math.Multiply(10.5, 5.3))

	if result, err := math.Divide(10.5, 5.3); err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("除法: %.2f ÷ %.2f = %.2f\n", 10.5, 5.3, result)
	}

	// 高级数学函数
	fmt.Printf("最大值: max(%.2f, %.2f) = %.2f\n", 10.5, 5.3, math.Max(10.5, 5.3))
	fmt.Printf("最小值: min(%.2f, %.2f) = %.2f\n", 10.5, 5.3, math.Min(10.5, 5.3))
	fmt.Printf("绝对值: abs(%.2f) = %.2f\n", -10.5, math.Abs(-10.5))
	fmt.Printf("四舍五入: round(%.3f, 2) = %.2f\n", 10.567, math.Round(10.567, 2))

	// 数论函数
	fmt.Printf("是否为偶数: isEven(10) = %t\n", math.IsEven(10))
	fmt.Printf("是否为奇数: isOdd(10) = %t\n", math.IsOdd(10))

	if factorial, err := math.Factorial(5); err != nil {
		fmt.Printf("阶乘错误: %v\n", err)
	} else {
		fmt.Printf("阶乘: 5! = %d\n", factorial)
	}

	fmt.Printf("最大公约数: gcd(48, 18) = %d\n", math.GCD(48, 18))
	fmt.Printf("最小公倍数: lcm(48, 18) = %d\n", math.LCM(48, 18))
	fmt.Printf("是否为质数: isPrime(17) = %t\n", math.IsPrime(17))

	if fib, err := math.Fibonacci(10); err != nil {
		fmt.Printf("斐波那契错误: %v\n", err)
	} else {
		fmt.Printf("斐波那契数列第10项: %d\n", fib)
	}

	// 统计函数
	numbers := []float64{1.5, 2.8, 3.2, 4.1, 5.9}
	fmt.Printf("数组: %v\n", numbers)
	fmt.Printf("求和: %.2f\n", math.Sum(numbers))

	if avg, err := math.Average(numbers); err != nil {
		fmt.Printf("平均值错误: %v\n", err)
	} else {
		fmt.Printf("平均值: %.2f\n", avg)
	}

	if median, err := math.Median(numbers); err != nil {
		fmt.Printf("中位数错误: %v\n", err)
	} else {
		fmt.Printf("中位数: %.2f\n", median)
	}
}

func stringExamples() {
	fmt.Println("\n🔹 字符串处理包示例")
	fmt.Println(strings.Repeat("-", 30))

	text := "Hello World Go Programming"
	fmt.Printf("原始字符串: %s\n", text)

	// 基本字符串操作
	fmt.Printf("反转: %s\n", stringutils.Reverse(text))
	fmt.Printf("字符数: %d\n", stringutils.CharCount(text))
	fmt.Printf("字节数: %d\n", stringutils.ByteCount(text))
	fmt.Printf("单词数: %d\n", stringutils.WordCount(text))

	// 大小写转换
	fmt.Printf("首字母大写: %s\n", stringutils.Capitalize("hello world"))
	fmt.Printf("标题格式: %s\n", stringutils.Title("hello world go programming"))
	fmt.Printf("驼峰命名: %s\n", stringutils.CamelCase("hello world go programming"))
	fmt.Printf("帕斯卡命名: %s\n", stringutils.PascalCase("hello world go programming"))
	fmt.Printf("蛇形命名: %s\n", stringutils.SnakeCase("hello world go programming"))
	fmt.Printf("短横线命名: %s\n", stringutils.KebabCase("hello world go programming"))

	// 字符串处理
	longText := "这是一个很长的字符串，用来演示截断功能"
	fmt.Printf("截断(10): %s\n", stringutils.Truncate(longText, 10))
	fmt.Printf("截断带省略号(10): %s\n", stringutils.TruncateWithEllipsis(longText, 10))

	// 填充
	fmt.Printf("左填充: '%s'\n", stringutils.PadLeft("Go", 10, '*'))
	fmt.Printf("右填充: '%s'\n", stringutils.PadRight("Go", 10, '*'))
	fmt.Printf("居中填充: '%s'\n", stringutils.Pad("Go", 10, '*'))

	// 回文检测
	palindromes := []string{"level", "A man a plan a canal Panama", "hello"}
	for _, p := range palindromes {
		fmt.Printf("'%s' 是回文: %t\n", p, stringutils.IsPalindrome(p))
	}

	// 格式验证
	emails := []string{"test@example.com", "invalid-email", "user@domain.org"}
	for _, email := range emails {
		fmt.Printf("'%s' 是有效邮箱: %t\n", email, stringutils.IsEmail(email))
	}

	phones := []string{"13812345678", "1234567890", "15987654321"}
	for _, phone := range phones {
		fmt.Printf("'%s' 是有效手机号: %t\n", phone, stringutils.IsPhone(phone))
	}

	// 遮蔽敏感信息
	fmt.Printf("遮蔽邮箱: %s -> %s\n", "user@example.com", stringutils.MaskEmail("user@example.com"))
	fmt.Printf("遮蔽手机: %s -> %s\n", "13812345678", stringutils.MaskPhone("13812345678"))

	// 字符串相似度
	s1, s2 := "hello", "hallo"
	fmt.Printf("'%s' 和 '%s' 的相似度: %.2f\n", s1, s2, stringutils.Similarity(s1, s2))
	fmt.Printf("'%s' 和 '%s' 的编辑距离: %d\n", s1, s2, stringutils.LevenshteinDistance(s1, s2))

	// 提取信息
	mixedText := "联系我们: email@example.com 或拨打 13812345678，访问 https://example.com"
	fmt.Printf("提取的邮箱: %v\n", stringutils.ExtractEmails(mixedText))
	fmt.Printf("提取的数字: %v\n", stringutils.ExtractNumbers(mixedText))

	// 字符串包含检查
	keywords := []string{"Go", "Programming", "Language"}
	fmt.Printf("'%s' 包含任意关键词 %v: %t\n", text, keywords, stringutils.ContainsAny(text, keywords))
	fmt.Printf("'%s' 包含所有关键词 %v: %t\n", text, keywords, stringutils.ContainsAll(text, keywords))
}

func calculatorExample() {
	fmt.Println("\n🔹 计算器示例")
	fmt.Println(strings.Repeat("-", 30))

	// 创建计算器实例
	calc := math.New()

	// 执行一系列计算
	fmt.Printf("加法: %.2f\n", calc.Add(10, 5))
	fmt.Printf("减法: %.2f\n", calc.Subtract(20, 8))
	fmt.Printf("乘法: %.2f\n", calc.Multiply(6, 7))

	if result, err := calc.Divide(15, 3); err != nil {
		log.Printf("除法错误: %v", err)
	} else {
		fmt.Printf("除法: %.2f\n", result)
	}

	fmt.Printf("幂运算: %.2f\n", calc.Power(2, 8))

	if result, err := calc.Sqrt(16); err != nil {
		log.Printf("平方根错误: %v", err)
	} else {
		fmt.Printf("平方根: %.2f\n", result)
	}

	// 三角函数（使用弧度）
	angle := 1.5708 // π/2 弧度，即90度
	fmt.Printf("sin(π/2): %.4f\n", calc.Sin(angle))
	fmt.Printf("cos(π/2): %.4f\n", calc.Cos(angle))

	// 获取最后一次计算结果
	if lastResult, err := calc.GetLastResult(); err != nil {
		log.Printf("获取最后结果错误: %v", err)
	} else {
		fmt.Printf("最后一次计算结果: %.4f\n", lastResult)
	}

	// 打印计算历史
	fmt.Println("\n计算历史:")
	calc.PrintHistory()

	// 清空历史
	fmt.Println("\n清空历史后:")
	calc.ClearHistory()
	calc.PrintHistory()
}
