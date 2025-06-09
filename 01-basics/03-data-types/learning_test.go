package datatypes

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

// 🎓 学习导向的测试 - 通过测试学习Go数据类型

// TestLearnBasicTypes 学习基础数据类型
func TestLearnBasicTypes(t *testing.T) {
	t.Log("🎯 学习目标: 理解Go语言的基础数据类型")
	t.Log("📚 本测试将教您: 整数、浮点数、布尔值、字符串的特性")

	t.Run("学习整数类型的范围和特性", func(t *testing.T) {
		t.Log("📖 知识点: Go有多种整数类型，每种都有不同的取值范围")

		// 🔍 探索: 不同整数类型的最大值
		t.Log("🔍 探索不同整数类型的最大值:")

		var i8 int8 = math.MaxInt8
		var i16 int16 = math.MaxInt16
		var i32 int32 = math.MaxInt32
		var i64 int64 = math.MaxInt64

		t.Logf("   int8  最大值: %d (占用 %d 字节)", i8, 1)
		t.Logf("   int16 最大值: %d (占用 %d 字节)", i16, 2)
		t.Logf("   int32 最大值: %d (占用 %d 字节)", i32, 4)
		t.Logf("   int64 最大值: %d (占用 %d 字节)", i64, 8)

		// ✅ 验证学习成果
		if i8 != 127 {
			t.Errorf("❌ 学习检查失败: int8最大值应该是127，您得到了%d", i8)
		} else {
			t.Log("✅ 很好！您正确理解了int8的范围")
		}

		// 💡 学习提示
		t.Log("💡 学习提示: 选择合适的整数类型可以节省内存")
		t.Log("💡 实践建议: 一般情况下使用int类型，除非有特殊需求")
	})

	t.Run("学习浮点数的精度问题", func(t *testing.T) {
		t.Log("📖 知识点: 浮点数存在精度问题，需要特别注意")

		// 🔍 探索: 浮点数精度问题
		a := 0.1
		b := 0.2
		sum := a + b
		expected := 0.3

		t.Logf("🔍 探索浮点数精度: %.1f + %.1f = %.17f", a, b, sum)
		t.Logf("   期望结果: %.1f", expected)
		t.Logf("   直接比较: %t", sum == expected)

		// ✅ 正确的浮点数比较方法
		epsilon := 1e-9
		isEqual := math.Abs(sum-expected) < epsilon
		t.Logf("   使用epsilon比较: %t (epsilon = %e)", isEqual, epsilon)

		if !isEqual {
			t.Error("❌ 浮点数比较方法需要改进")
		} else {
			t.Log("✅ 很好！您掌握了正确的浮点数比较方法")
		}

		// 💡 学习提示
		t.Log("💡 重要提示: 永远不要直接比较浮点数是否相等")
		t.Log("💡 最佳实践: 使用epsilon进行近似比较")
	})

	t.Run("学习字符串的内部结构", func(t *testing.T) {
		t.Log("📖 知识点: Go字符串是UTF-8编码的字节序列")

		// 🔍 探索: 字符串的字节和字符
		text := "Hello世界"
		byteLen := len(text)
		runeCount := utf8.RuneCountInString(text)

		t.Logf("🔍 探索字符串 '%s':", text)
		t.Logf("   字节长度: %d", byteLen)
		t.Logf("   字符数量: %d", runeCount)

		// 🔍 详细分析每个字符
		t.Log("   字符详细分析:")
		for i, r := range text {
			t.Logf("     位置%d: '%c' (Unicode: U+%04X, 值: %d)", i, r, r, r)
		}

		// ✅ 验证理解
		if byteLen == runeCount {
			t.Log("✅ 这个字符串只包含ASCII字符")
		} else {
			t.Log("✅ 这个字符串包含多字节字符（如中文）")
		}

		// 💡 学习提示
		t.Log("💡 关键概念: len()返回字节数，range遍历字符")
		t.Log("💡 实践建议: 处理国际化文本时要考虑字符和字节的区别")
	})
}

// TestLearnTypeConversions 学习类型转换
func TestLearnTypeConversions(t *testing.T) {
	t.Log("🎯 学习目标: 掌握Go语言的类型转换规则")
	t.Log("📚 本测试将教您: 显式转换、字符串转换、类型安全")

	t.Run("学习数值类型转换", func(t *testing.T) {
		t.Log("📖 知识点: Go要求显式类型转换，不允许隐式转换")

		// 🔍 探索: 不同数值类型之间的转换
		var i int = 42
		var f float64 = float64(i) // 必须显式转换
		var i32 int32 = int32(i)

		t.Logf("🔍 类型转换示例:")
		t.Logf("   int(%d) → float64(%.1f)", i, f)
		t.Logf("   int(%d) → int32(%d)", i, i32)

		// 🔍 探索: 精度丢失的情况
		var bigFloat float64 = 3.14159
		var intFromFloat int = int(bigFloat) // 小数部分会丢失

		t.Logf("   float64(%.5f) → int(%d) [小数部分丢失]", bigFloat, intFromFloat)

		// ✅ 验证转换结果
		if intFromFloat != 3 {
			t.Errorf("❌ 类型转换理解有误: 期望3，得到%d", intFromFloat)
		} else {
			t.Log("✅ 很好！您理解了浮点数转整数会截断小数部分")
		}

		// 💡 学习提示
		t.Log("💡 安全提示: 类型转换可能导致数据丢失")
		t.Log("💡 最佳实践: 转换前检查数值范围")
	})

	t.Run("学习字符串和数值的转换", func(t *testing.T) {
		t.Log("📖 知识点: 使用strconv包进行字符串和数值的转换")

		// 🔍 探索: 数值转字符串
		num := 123
		str := strconv.Itoa(num)
		t.Logf("🔍 数值转字符串: %d → \"%s\"", num, str)

		// 🔍 探索: 字符串转数值（可能失败）
		validStr := "456"
		invalidStr := "abc"

		if result, err := strconv.Atoi(validStr); err == nil {
			t.Logf("   成功转换: \"%s\" → %d", validStr, result)
		} else {
			t.Logf("   转换失败: \"%s\" → 错误: %v", validStr, err)
		}

		if result, err := strconv.Atoi(invalidStr); err == nil {
			t.Logf("   成功转换: \"%s\" → %d", invalidStr, result)
		} else {
			t.Logf("   转换失败: \"%s\" → 错误: %v", invalidStr, err)
		}

		// ✅ 验证错误处理
		if _, err := strconv.Atoi(invalidStr); err == nil {
			t.Error("❌ 应该检测到无效字符串转换错误")
		} else {
			t.Log("✅ 很好！您理解了字符串转换可能失败")
		}

		// 💡 学习提示
		t.Log("💡 重要概念: 字符串转数值可能失败，必须检查错误")
		t.Log("💡 最佳实践: 总是检查strconv函数的错误返回值")
	})
}

// TestLearnZeroValues 学习零值概念
func TestLearnZeroValues(t *testing.T) {
	t.Log("🎯 学习目标: 理解Go语言的零值概念")
	t.Log("📚 本测试将教您: 每种类型的零值是什么")

	t.Run("探索各种类型的零值", func(t *testing.T) {
		t.Log("📖 知识点: Go中每种类型都有一个零值（默认值）")

		// 🔍 探索: 声明变量但不初始化
		var intVal int
		var floatVal float64
		var boolVal bool
		var stringVal string
		var sliceVal []int
		var mapVal map[string]int
		var ptrVal *int

		t.Log("🔍 各种类型的零值:")
		t.Logf("   int零值: %d", intVal)
		t.Logf("   float64零值: %f", floatVal)
		t.Logf("   bool零值: %t", boolVal)
		t.Logf("   string零值: \"%s\" (长度: %d)", stringVal, len(stringVal))
		t.Logf("   slice零值: %v (是否为nil: %t)", sliceVal, sliceVal == nil)
		t.Logf("   map零值: %v (是否为nil: %t)", mapVal, mapVal == nil)
		t.Logf("   pointer零值: %v (是否为nil: %t)", ptrVal, ptrVal == nil)

		// ✅ 验证零值知识
		tests := []struct {
			name     string
			actual   interface{}
			expected interface{}
		}{
			{"int零值", intVal, 0},
			{"float64零值", floatVal, 0.0},
			{"bool零值", boolVal, false},
			{"string零值", stringVal, ""},
		}

		for _, test := range tests {
			if test.actual != test.expected {
				t.Errorf("❌ %s错误: 期望%v，得到%v", test.name, test.expected, test.actual)
			} else {
				t.Logf("✅ %s正确", test.name)
			}
		}

		// 💡 学习提示
		t.Log("💡 重要概念: 零值让Go程序更安全，避免了未初始化变量的问题")
		t.Log("💡 实践建议: 利用零值特性简化代码逻辑")
	})
}

// TestLearnTypeReflection 学习类型反射
func TestLearnTypeReflection(t *testing.T) {
	t.Log("🎯 学习目标: 了解如何在运行时获取类型信息")
	t.Log("📚 本测试将教您: 使用reflect包检查类型")

	t.Run("探索类型信息", func(t *testing.T) {
		t.Log("📖 知识点: 使用reflect包可以在运行时检查类型")

		// 🔍 探索: 不同值的类型信息
		values := []interface{}{
			42,
			3.14,
			"hello",
			true,
			[]int{1, 2, 3},
			map[string]int{"a": 1},
		}

		t.Log("🔍 各种值的类型信息:")
		for i, v := range values {
			t.Logf("   值%d: %v", i+1, v)
			t.Logf("     类型: %T", v)
			t.Logf("     reflect.Type: %v", reflect.TypeOf(v))
			t.Logf("     reflect.Kind: %v", reflect.TypeOf(v).Kind())
			t.Logf("     reflect.Value: %v", reflect.ValueOf(v))
			t.Log("")
		}

		// ✅ 验证类型检查
		intVal := 42
		if reflect.TypeOf(intVal).Kind() != reflect.Int {
			t.Error("❌ 类型检查失败")
		} else {
			t.Log("✅ 很好！您学会了使用reflect检查类型")
		}

		// 💡 学习提示
		t.Log("💡 高级概念: 反射是强大的工具，但会影响性能")
		t.Log("💡 使用场景: JSON序列化、ORM、通用函数等")
	})
}

// TestLearnCommonMistakes 学习常见错误
func TestLearnCommonMistakes(t *testing.T) {
	t.Log("🎯 学习目标: 了解数据类型使用中的常见错误")
	t.Log("📚 本测试将教您: 避免常见的类型使用陷阱")

	t.Run("避免浮点数比较陷阱", func(t *testing.T) {
		t.Log("📖 常见错误: 直接比较浮点数")

		// ❌ 错误的做法
		a := 0.1 + 0.2
		b := 0.3
		directCompare := (a == b)

		t.Logf("❌ 错误做法: (0.1 + 0.2) == 0.3 = %t", directCompare)
		t.Logf("   实际值: %.17f", a)
		t.Logf("   期望值: %.17f", b)

		// ✅ 正确的做法
		epsilon := 1e-9
		correctCompare := math.Abs(a-b) < epsilon
		t.Logf("✅ 正确做法: |a - b| < epsilon = %t", correctCompare)

		if directCompare {
			t.Log("⚠️  警告: 您的环境中浮点数比较恰好成功，但这不可靠")
		}

		// 💡 学习提示
		t.Log("💡 避免陷阱: 永远不要直接比较浮点数")
		t.Log("💡 正确方法: 使用epsilon进行近似比较")
	})

	t.Run("避免字符串索引陷阱", func(t *testing.T) {
		t.Log("📖 常见错误: 按字节索引多字节字符")

		text := "Hello世界"

		// ❌ 可能有问题的做法（按字节索引）
		t.Log("❌ 按字节索引可能有问题:")
		for i := 0; i < len(text); i++ {
			char := text[i]
			if char < 128 { // ASCII字符
				t.Logf("   [%d]: %c (ASCII)", i, char)
			} else {
				t.Logf("   [%d]: %d (非ASCII字节)", i, char)
			}
		}

		// ✅ 正确的做法（按字符遍历）
		t.Log("✅ 正确的字符遍历:")
		for i, r := range text {
			t.Logf("   [%d]: %c (Unicode: %d)", i, r, r)
		}

		// 💡 学习提示
		t.Log("💡 避免陷阱: 不要假设一个字符等于一个字节")
		t.Log("💡 正确方法: 使用range遍历字符，或使用utf8包")
	})
}

// BenchmarkLearnPerformance 学习性能基准测试
func BenchmarkLearnPerformance(b *testing.B) {
	b.Log("🎯 学习目标: 了解不同操作的性能差异")

	// 字符串连接性能比较
	b.Run("字符串连接性能", func(b *testing.B) {
		b.Run("使用+操作符", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result := ""
				for j := 0; j < 100; j++ {
					result += "a"
				}
				_ = result
			}
		})

		b.Run("使用strings.Builder", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var builder strings.Builder
				for j := 0; j < 100; j++ {
					builder.WriteString("a")
				}
				_ = builder.String()
			}
		})
	})
}

// Example_learnBasicUsage 学习示例
func Example_learnBasicUsage() {
	// 这是一个可执行的示例，展示基本用法
	var age int = 25
	var name string = "张三"
	var isStudent bool = true

	fmt.Printf("姓名: %s, 年龄: %d, 是否学生: %t\n", name, age, isStudent)

	// Output:
	// 姓名: 张三, 年龄: 25, 是否学生: true
}
