package datatypes

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestNumericTypes(t *testing.T) {
	t.Run("IntegerTypes", func(t *testing.T) {
		// 有符号整数类型
		var int8Val int8 = 127
		var int16Val int16 = 32767
		var int32Val int32 = 2147483647
		var int64Val int64 = 9223372036854775807

		t.Logf("int8 最大值: %d", int8Val)
		t.Logf("int16 最大值: %d", int16Val)
		t.Logf("int32 最大值: %d", int32Val)
		t.Logf("int64 最大值: %d", int64Val)

		// 验证范围
		if int8Val != 127 {
			t.Errorf("int8 最大值应该是 127, 实际是 %d", int8Val)
		}
	})

	t.Run("UnsignedIntegerTypes", func(t *testing.T) {
		// 无符号整数类型
		var uint8Val uint8 = 255
		var uint16Val uint16 = 65535
		var uint32Val uint32 = 4294967295
		var uint64Val uint64 = 18446744073709551615

		t.Logf("uint8 最大值: %d", uint8Val)
		t.Logf("uint16 最大值: %d", uint16Val)
		t.Logf("uint32 最大值: %d", uint32Val)
		t.Logf("uint64 最大值: %d", uint64Val)

		// 验证范围
		if uint8Val != 255 {
			t.Errorf("uint8 最大值应该是 255, 实际是 %d", uint8Val)
		}
	})

	t.Run("ArithmeticOperations", func(t *testing.T) {
		a, b := 10, 3

		t.Logf("a = %d, b = %d", a, b)
		t.Logf("加法: %d + %d = %d", a, b, a+b)
		t.Logf("减法: %d - %d = %d", a, b, a-b)
		t.Logf("乘法: %d × %d = %d", a, b, a*b)
		t.Logf("除法: %d ÷ %d = %d", a, b, a/b)
		t.Logf("取余: %d %% %d = %d", a, b, a%b)

		// 验证运算结果
		if a+b != 13 {
			t.Errorf("加法错误: 期望 13, 实际 %d", a+b)
		}
	})
}

func TestFloatingPointTypes(t *testing.T) {
	t.Run("BasicFloatTypes", func(t *testing.T) {
		var float32Val float32 = 3.14159
		var float64Val float64 = 3.141592653589793

		t.Logf("float32: %.6f", float32Val)
		t.Logf("float64: %.15f", float64Val)

		// 精度比较
		diff := math.Abs(float64(float32Val) - 3.14159)
		if diff > 1e-5 {
			t.Errorf("float32 精度不足")
		}
	})

	t.Run("SpecialFloatValues", func(t *testing.T) {
		posInf := math.Inf(1)
		negInf := math.Inf(-1)
		notANumber := math.NaN()

		t.Logf("正无穷: %f", posInf)
		t.Logf("负无穷: %f", negInf)
		t.Logf("NaN: %f", notANumber)

		// 验证特殊值
		if !math.IsInf(posInf, 1) {
			t.Error("正无穷检测失败")
		}
		if !math.IsInf(negInf, -1) {
			t.Error("负无穷检测失败")
		}
		if !math.IsNaN(notANumber) {
			t.Error("NaN 检测失败")
		}
	})

	t.Run("FloatArithmetic", func(t *testing.T) {
		a, b := 3.14, 2.71
		t.Logf("a = %.2f, b = %.2f", a, b)
		t.Logf("加法: %.2f + %.2f = %.2f", a, b, a+b)
		t.Logf("减法: %.2f - %.2f = %.2f", a, b, a-b)
		t.Logf("乘法: %.2f × %.2f = %.2f", a, b, a*b)
		t.Logf("除法: %.2f ÷ %.2f = %.2f", a, b, a/b)

		v, m := 0.1, 0.2
		t.Logf("v = %.1f, m = %.1f", v, m)
		t.Logf("加法: %.1f + %.1f = %.1f", v, m, v+m)
		t.Logf("减法: %.1f - %.1f = %.1f", v, m, v-m)
		t.Logf("乘法: %.1f × %.1f = %.1f", v, m, v*m)
		t.Logf("除法: %.1f ÷ %.1f = %.1f", v, m, v/m)
		r := 0.3

		t.Log("v + m == r: ", v+m == r)
		epsilon := math.Nextafter(1, 2) - 1
		t.Log("v + m == r: ", math.Abs(v+m-r) < epsilon)
	})
}

func TestBooleanTypes(t *testing.T) {
	t.Run("BasicBooleans", func(t *testing.T) {
		var isTrue bool = true
		var isFalse bool = false
		var defaultBool bool // 零值

		t.Logf("true: %t", isTrue)
		t.Logf("false: %t", isFalse)
		t.Logf("默认值: %t", defaultBool)

		// 验证零值
		if defaultBool != false {
			t.Error("bool 零值应该是 false")
		}
	})

	t.Run("BooleanOperations", func(t *testing.T) {
		a, b := true, false

		t.Logf("a = %t, b = %t", a, b)
		t.Logf("AND: %t && %t = %t", a, b, a && b)
		t.Logf("OR: %t || %t = %t", a, b, a || b)
		t.Logf("NOT a: !%t = %t", a, !a)
		t.Logf("NOT b: !%t = %t", b, !b)

		// 验证逻辑运算
		if (a && b) != false {
			t.Error("true && false 应该是 false")
		}
		if (a || b) != true {
			t.Error("true || false 应该是 true")
		}
	})

	t.Run("ComparisonOperations", func(t *testing.T) {
		x, y := 10, 20

		t.Logf("x = %d, y = %d", x, y)
		t.Logf("相等: %d == %d = %t", x, y, x == y)
		t.Logf("不等: %d != %d = %t", x, y, x != y)
		t.Logf("小于: %d < %d = %t", x, y, x < y)
		t.Logf("小于等于: %d <= %d = %t", x, y, x <= y)
		t.Logf("大于: %d > %d = %t", x, y, x > y)
		t.Logf("大于等于: %d >= %d = %t", x, y, x >= y)
	})
}

func TestStringTypes(t *testing.T) {
	t.Run("BasicStrings", func(t *testing.T) {
		var greeting string = "你好，世界！"
		var empty string
		var rawString string = `这是一个
原始字符串
可以包含换行符和"引号"`

		t.Logf("问候语: %s", greeting)
		t.Logf("空字符串: '%s'", empty)
		t.Logf("空字符串带引号格式化: %q", empty)
		t.Logf("原始字符串:\n%s", rawString)

		// 验证零值
		if empty != "" {
			t.Error("string 零值应该是空字符串")
		}
	})

	t.Run("StringLength", func(t *testing.T) {
		text := "Go语言"
		byteLen := len(text)
		runeLen := utf8.RuneCountInString(text)

		t.Logf("字符串: %s", text)
		t.Logf("字节长度: %d", byteLen)
		t.Logf("字符数量: %d", runeLen)

		// 验证长度
		if byteLen != 8 { // "Go" (2字节) + "语言" (6字节)
			t.Errorf("字节长度错误: 期望 8, 实际 %d", byteLen)
		}
		if runeLen != 4 { // 4个字符
			t.Errorf("字符数量错误: 期望 4, 实际 %d", runeLen)
		}
	})

	t.Run("StringOperations", func(t *testing.T) {
		text := "Hello World"

		t.Logf("原字符串: %s", text)
		t.Logf("转大写: %s", strings.ToUpper(text))
		t.Logf("转小写: %s", strings.ToLower(text))
		t.Logf("包含 'World': %t", strings.Contains(text, "World"))
		t.Logf("以 'Hello' 开头: %t", strings.HasPrefix(text, "Hello"))
		t.Logf("以 'World' 结尾: %t", strings.HasSuffix(text, "World"))
		t.Logf("替换 'World' 为 'Go': %s", strings.Replace(text, "World", "Go", 1))

		// 字符串分割和连接
		parts := strings.Split(text, " ")
		t.Logf("分割结果: %v", parts)
		joined := strings.Join(parts, "-")
		t.Logf("连接结果: %s", joined)
	})

	t.Run("StringIteration", func(t *testing.T) {
		text := "Go语言"
		t.Logf("遍历字符串 '%s':", text)

		// 按字节遍历
		t.Log("按字节遍历:")
		for i := 0; i < len(text); i++ {
			unicode := fmt.Sprintf("%02X", text[i])
			t.Logf("  [%d]: %d (%c) (Unicode: %s)", i, text[i], text[i], unicode)
		}

		// 按字符遍历
		t.Log("按字符遍历:")
		for i, r := range text {
			t.Logf("  [%d]: %c (Unicode: %d) （Unicode HEX: %X)", i, r, r, r)
		}
	})
}

func TestCharacterTypes(t *testing.T) {
	t.Run("RuneAndByte", func(t *testing.T) {
		var char1 rune = 'A'
		var char2 rune = '中'
		var char3 byte = 65

		t.Logf("rune 'A': %c (值: %d)", char1, char1)
		t.Logf("rune '中': %c (值: %d)", char2, char2)
		t.Logf("byte 65: %c (值: %d)", char3, char3)

		// 验证字符值
		if char1 != 65 {
			t.Errorf("字符 'A' 的值应该是 65, 实际是 %d", char1)
		}
		if char3 != 65 {
			t.Errorf("字节 65 应该是 65, 实际是 %d", char3)
		}
	})

	t.Run("UnicodeHandling", func(t *testing.T) {
		// Unicode 字符处理
		emoji := '😀'
		chinese := '中'
		english := 'A'

		t.Logf("表情符号: %c (Unicode: %d)", emoji, emoji)
		t.Logf("中文字符: %c (Unicode: %d)", chinese, chinese)
		t.Logf("英文字符: %c (Unicode: %d)", english, english)

		// 验证 Unicode 范围
		if chinese < 0x4E00 || chinese > 0x9FFF {
			t.Error("中文字符不在预期的 Unicode 范围内")
		}
	})
}

func TestTypeConversions(t *testing.T) {
	t.Run("NumericConversions", func(t *testing.T) {
		var intVal int = 42
		var floatVal float64 = float64(intVal)
		var int32Val int32 = int32(intVal)

		t.Logf("原整数: %d", intVal)
		t.Logf("转换为 float64: %f", floatVal)
		t.Logf("转换为 int32: %d", int32Val)

		// 验证转换
		if floatVal != 42.0 {
			t.Errorf("转换错误: 期望 42.0, 实际 %f", floatVal)
		}
	})

	t.Run("StringConversions", func(t *testing.T) {
		// 数值转字符串
		intVal := 123
		floatVal := 3.14
		boolVal := true

		intStr := strconv.Itoa(intVal)
		floatStr := strconv.FormatFloat(floatVal, 'f', 2, 64)
		boolStr := strconv.FormatBool(boolVal)

		t.Logf("整数 %d 转字符串: %s", intVal, intStr)
		t.Logf("浮点数 %f 转字符串: %s", floatVal, floatStr)
		t.Logf("布尔值 %t 转字符串: %s", boolVal, boolStr)

		// 字符串转数值
		if num, err := strconv.Atoi("456"); err == nil {
			t.Logf("字符串 '456' 转整数: %d", num)
		}

		if num, err := strconv.ParseFloat("2.71", 64); err == nil {
			t.Logf("字符串 '2.71' 转浮点数: %f", num)
		}

		if b, err := strconv.ParseBool("false"); err == nil {
			t.Logf("字符串 'false' 转布尔值: %t", b)
		}
	})
}

func TestComplexTypes(t *testing.T) {
	t.Run("ComplexNumbers", func(t *testing.T) {
		var c1 complex64 = 1 + 2i
		var c2 complex128 = 3.14 + 2.71i

		t.Logf("complex64: %v", c1)
		t.Logf("complex128: %v", c2)

		// 复数运算
		sum := c2 + complex128(c1)
		product := c2 * 2

		t.Logf("复数加法: %v + %v = %v", c2, c1, sum)
		t.Logf("复数乘法: %v × 2 = %v", c2, product)

		// 获取实部和虚部
		realPart := real(c2)
		imagPart := imag(c2)

		t.Logf("实部: %f", realPart)
		t.Logf("虚部: %f", imagPart)

		// 验证实部和虚部
		if realPart != 3.14 {
			t.Errorf("实部错误: 期望 3.14, 实际 %f", realPart)
		}
		if imagPart != 2.71 {
			t.Errorf("虚部错误: 期望 2.71, 实际 %f", imagPart)
		}
	})
}

func TestTypeAliases(t *testing.T) {
	t.Run("CustomTypes", func(t *testing.T) {
		// 定义自定义类型
		type UserID int
		type UserName string
		type Score float64

		var id UserID = 12345
		var name UserName = "张三"
		var score Score = 95.5

		t.Logf("用户ID: %d (类型: %T)", id, id)
		t.Logf("用户名: %s (类型: %T)", name, name)
		t.Logf("分数: %.1f (类型: %T)", score, score)

		// 类型转换
		normalInt := int(id)
		normalString := string(name)
		normalFloat := float64(score)

		t.Logf("转换后的ID: %d (类型: %T)", normalInt, normalInt)
		t.Logf("转换后的名称: %s (类型: %T)", normalString, normalString)
		t.Logf("转换后的分数: %.1f (类型: %T)", normalFloat, normalFloat)
	})
}
