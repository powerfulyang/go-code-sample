package datatypes

import (
	"fmt"
	"strconv"
	"unsafe"
)

// BooleansExample 展示Go中的布尔类型
func BooleansExample() {
	fmt.Println("=== 布尔类型 ===")

	// 布尔变量声明和初始化
	var isTrue bool = true
	var isFalse bool = false

	shortTrue := true
	shortFalse := false

	// 布尔值的大小
	fmt.Printf("布尔值大小: %d 字节\n", unsafe.Sizeof(isTrue))

	// 布尔值输出
	fmt.Printf("isTrue = %t\n", isTrue)
	fmt.Printf("isFalse = %t\n", isFalse)
	fmt.Printf("shortTrue = %t\n", shortTrue)
	fmt.Printf("shortFalse = %t\n", shortFalse)

	// 逻辑运算符: AND(&&)
	fmt.Println("\n=== 逻辑与 (AND) ===")
	fmt.Printf("true && true = %t\n", true && true)
	fmt.Printf("true && false = %t\n", true && false)
	fmt.Printf("false && true = %t\n", false && true)
	fmt.Printf("false && false = %t\n", false && false)

	// 逻辑运算符: OR(||)
	fmt.Println("\n=== 逻辑或 (OR) ===")
	fmt.Printf("true || true = %t\n", true || true)
	fmt.Printf("true || false = %t\n", true || false)
	fmt.Printf("false || true = %t\n", false || true)
	fmt.Printf("false || false = %t\n", false || false)

	// 逻辑运算符: NOT(!)
	fmt.Println("\n=== 逻辑非 (NOT) ===")
	fmt.Printf("!true = %t\n", !true)
	fmt.Printf("!false = %t\n", !false)

	// 使用短路求值的例子
	fmt.Println("\n=== 短路求值 ===")

	x := 10

	// 在第一个条件为false时，第二个条件不会被求值
	result1 := x > 20 && evaluateAndPrint("第二个条件被求值")
	fmt.Printf("x > 20 && ... = %t (第二个条件%s被求值)\n", result1, map[bool]string{true: "", false: "不"}[result1])

	// 在第一个条件为true时，第二个条件不会被求值
	result2 := x > 5 || evaluateAndPrint("第二个条件被求值")
	fmt.Printf("x > 5 || ... = %t (第二个条件%s被求值)\n", result2, map[bool]string{true: "不", false: ""}[result2])

	// 比较运算符
	fmt.Println("\n=== 比较运算符 ===")

	a, b := 5, 10

	fmt.Printf("%d == %d: %t\n", a, a, a == a)
	fmt.Printf("%d == %d: %t\n", a, b, a == b)
	fmt.Printf("%d != %d: %t\n", a, b, a != b)
	fmt.Printf("%d < %d: %t\n", a, b, a < b)
	fmt.Printf("%d <= %d: %t\n", a, b, a <= b)
	fmt.Printf("%d > %d: %t\n", a, b, a > b)
	fmt.Printf("%d >= %d: %t\n", a, b, a >= b)

	// 条件结构中的布尔值
	fmt.Println("\n=== 在条件结构中使用布尔值 ===")

	condition := true

	if condition {
		fmt.Println("条件为真")
	} else {
		fmt.Println("条件为假")
	}

	// 布尔值转换为字符串
	fmt.Println("\n=== 布尔值与字符串转换 ===")

	trueString := strconv.FormatBool(true)
	falseString := strconv.FormatBool(false)

	fmt.Printf("true转为字符串: %s (类型: %T)\n", trueString, trueString)
	fmt.Printf("false转为字符串: %s (类型: %T)\n", falseString, falseString)

	// 字符串解析为布尔值
	str1 := "true"
	str2 := "false"
	str3 := "True" // 大小写不敏感
	str4 := "1"    // 1可以被解析为true
	str5 := "0"    // 0可以被解析为false
	str6 := "invalid"

	parseBoolValue(str1)
	parseBoolValue(str2)
	parseBoolValue(str3)
	parseBoolValue(str4)
	parseBoolValue(str5)
	parseBoolValue(str6)
}

// 辅助函数：打印布尔值并返回自身
func evaluateAndPrint(message string) bool {
	fmt.Println(message)
	return true
}

// 辅助函数：将字符串解析为布尔值并处理可能的错误
func parseBoolValue(s string) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Printf("解析 %q 出错: %v\n", s, err)
	} else {
		fmt.Printf("字符串 %q 解析为: %t\n", s, b)
	}
}
