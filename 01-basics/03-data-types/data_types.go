package datatypes

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 数值类型示例
func NumericTypes() {
	fmt.Println("=== 数值类型示例 ===")

	// 整数类型
	var int8Val int8 = 127     // -128 到 127
	var int16Val int16 = 32767 // -32768 到 32767
	var int32Val int32 = 2147483647
	var int64Val int64 = 9223372036854775807

	// 无符号整数类型
	var uint8Val uint8 = 255     // 0 到 255
	var uint16Val uint16 = 65535 // 0 到 65535
	var uint32Val uint32 = 4294967295
	var uint64Val uint64 = 18446744073709551615

	// 平台相关类型
	var intVal int = 2147483647
	var uintVal uint = 4294967295

	fmt.Printf("int8: %d (大小: %d 字节)\n", int8Val, 1)
	fmt.Printf("int16: %d (大小: %d 字节)\n", int16Val, 2)
	fmt.Printf("int32: %d (大小: %d 字节)\n", int32Val, 4)
	fmt.Printf("int64: %d (大小: %d 字节)\n", int64Val, 8)
	fmt.Printf("uint8: %d (大小: %d 字节)\n", uint8Val, 1)
	fmt.Printf("uint16: %d (大小: %d 字节)\n", uint16Val, 2)
	fmt.Printf("uint32: %d (大小: %d 字节)\n", uint32Val, 4)
	fmt.Printf("uint64: %d (大小: %d 字节)\n", uint64Val, 8)
	fmt.Printf("int: %d\n", intVal)
	fmt.Printf("uint: %d\n", uintVal)
}

// 浮点数类型示例
func FloatingPointTypes() {
	fmt.Println("\n=== 浮点数类型示例 ===")

	var float32Val float32 = 3.14159
	var float64Val float64 = 3.141592653589793

	fmt.Printf("float32: %.6f (精度: ~7位)\n", float32Val)
	fmt.Printf("float64: %.15f (精度: ~15位)\n", float64Val)

	// 特殊浮点数值
	fmt.Printf("正无穷: %f\n", math.Inf(1))
	fmt.Printf("负无穷: %f\n", math.Inf(-1))
	fmt.Printf("NaN: %f\n", math.NaN())

	// 浮点数运算
	result := float64Val * 2
	fmt.Printf("π × 2 = %.15f\n", result)
}

// 布尔类型示例
func BooleanTypes() {
	fmt.Println("\n=== 布尔类型示例 ===")

	var isTrue bool = true
	var isFalse bool = false
	var defaultBool bool // 零值为 false

	fmt.Printf("true: %t\n", isTrue)
	fmt.Printf("false: %t\n", isFalse)
	fmt.Printf("默认值: %t\n", defaultBool)

	// 布尔运算
	fmt.Printf("true && false = %t\n", isTrue && isFalse)
	fmt.Printf("true || false = %t\n", isTrue || isFalse)
	fmt.Printf("!true = %t\n", !isTrue)
	fmt.Printf("!false = %t\n", !isFalse)
}

// 字符串类型示例
func StringTypes() {
	fmt.Println("\n=== 字符串类型示例 ===")

	var greeting string = "你好，世界！"
	var empty string // 零值为空字符串
	var multiline string = `这是一个
多行字符串
可以包含换行符`

	fmt.Printf("问候语: %s\n", greeting)
	fmt.Printf("空字符串: '%s'\n", empty)
	fmt.Printf("多行字符串:\n%s\n", multiline)

	// 字符串长度（字节数和字符数）
	fmt.Printf("字节长度: %d\n", len(greeting))
	fmt.Printf("字符数量: %d\n", utf8.RuneCountInString(greeting))

	// 字符串操作
	fmt.Printf("转大写: %s\n", strings.ToUpper("hello"))
	fmt.Printf("转小写: %s\n", strings.ToLower("WORLD"))
	fmt.Printf("包含检查: %t\n", strings.Contains(greeting, "世界"))
	fmt.Printf("前缀检查: %t\n", strings.HasPrefix(greeting, "你好"))
	fmt.Printf("后缀检查: %t\n", strings.HasSuffix(greeting, "！"))
}

// 字符类型示例
func CharacterTypes() {
	fmt.Println("\n=== 字符类型示例 ===")

	var char1 rune = 'A' // rune 是 int32 的别名
	var char2 rune = '中'
	var char3 byte = 65 // byte 是 uint8 的别名

	fmt.Printf("字符 'A': %c (Unicode: %d)\n", char1, char1)
	fmt.Printf("字符 '中': %c (Unicode: %d)\n", char2, char2)
	fmt.Printf("字节 65: %c (ASCII: %d)\n", char3, char3)

	// 字符串遍历
	text := "Go语言"
	fmt.Printf("遍历字符串 '%s':\n", text)
	for i, r := range text {
		fmt.Printf("  位置 %d: %c (Unicode: %d)\n", i, r, r)
	}
}

// 类型转换示例
func TypeConversions() {
	fmt.Println("\n=== 类型转换示例 ===")

	// 数值类型转换
	var intVal int = 42
	var floatVal float64 = float64(intVal)
	var stringVal string = strconv.Itoa(intVal)

	fmt.Printf("整数: %d\n", intVal)
	fmt.Printf("转换为浮点数: %f\n", floatVal)
	fmt.Printf("转换为字符串: %s\n", stringVal)

	// 字符串转数值
	numStr := "123"
	if num, err := strconv.Atoi(numStr); err == nil {
		fmt.Printf("字符串 '%s' 转换为整数: %d\n", numStr, num)
	}

	floatStr := "3.14"
	if num, err := strconv.ParseFloat(floatStr, 64); err == nil {
		fmt.Printf("字符串 '%s' 转换为浮点数: %f\n", floatStr, num)
	}

	// 布尔值转换
	boolStr := "true"
	if b, err := strconv.ParseBool(boolStr); err == nil {
		fmt.Printf("字符串 '%s' 转换为布尔值: %t\n", boolStr, b)
	}
}

// 复数类型示例
func ComplexTypes() {
	fmt.Println("\n=== 复数类型示例 ===")

	var complex64Val complex64 = 1 + 2i
	var complex128Val complex128 = 3.14 + 2.71i

	fmt.Printf("complex64: %v\n", complex64Val)
	fmt.Printf("complex128: %v\n", complex128Val)

	// 复数运算
	result := complex128Val * 2
	fmt.Printf("复数乘法: %v × 2 = %v\n", complex128Val, result)

	// 获取实部和虚部
	fmt.Printf("实部: %f, 虚部: %f\n", real(complex128Val), imag(complex128Val))
}

// 指针类型示例
func PointerTypes() {
	fmt.Println("\n=== 指针类型示例 ===")

	var num int = 42
	var ptr *int = &num

	fmt.Printf("变量值: %d\n", num)
	fmt.Printf("变量地址: %p\n", &num)
	fmt.Printf("指针值: %p\n", ptr)
	fmt.Printf("指针指向的值: %d\n", *ptr)

	// 修改指针指向的值
	*ptr = 100
	fmt.Printf("通过指针修改后的值: %d\n", num)

	// 零值指针
	var nilPtr *int
	fmt.Printf("零值指针: %v\n", nilPtr)
}

// 类型别名示例
func TypeAliases() {
	fmt.Println("\n=== 类型别名示例 ===")

	// 定义类型别名
	type UserID int
	type UserName string
	type IsActive bool

	var id UserID = 12345
	var name UserName = "张三"
	var active IsActive = true

	fmt.Printf("用户ID: %d (类型: %T)\n", id, id)
	fmt.Printf("用户名: %s (类型: %T)\n", name, name)
	fmt.Printf("是否活跃: %t (类型: %T)\n", active, active)

	// 类型转换
	var normalInt int = int(id)
	fmt.Printf("转换为普通整数: %d (类型: %T)\n", normalInt, normalInt)
}

// 常量和枚举示例
func ConstantsAndEnums() {
	fmt.Println("\n=== 常量和枚举示例 ===")

	// 类型化常量
	const (
		Pi      float64 = 3.14159
		E       float64 = 2.71828
		MaxSize int     = 1000
	)

	fmt.Printf("π: %f\n", Pi)
	fmt.Printf("e: %f\n", E)
	fmt.Printf("最大大小: %d\n", MaxSize)

	// 枚举示例
	type Status int
	const (
		StatusPending Status = iota
		StatusApproved
		StatusRejected
		StatusCancelled
	)

	var currentStatus Status = StatusApproved
	fmt.Printf("当前状态: %d\n", currentStatus)

	// 字符串枚举
	const (
		ColorRed   = "red"
		ColorGreen = "green"
		ColorBlue  = "blue"
	)

	fmt.Printf("颜色: %s\n", ColorRed)
}
