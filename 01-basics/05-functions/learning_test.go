package functions

import (
	"errors"
	"fmt"
	"testing"
)

// 🎓 学习导向的测试 - 通过测试学习Go函数

// TestLearnBasicFunctions 学习基础函数概念
func TestLearnBasicFunctions(t *testing.T) {
	t.Log("🎯 学习目标: 掌握Go函数的基本语法和特性")
	t.Log("📚 本测试将教您: 函数定义、参数传递、返回值")

	t.Run("学习函数的基本结构", func(t *testing.T) {
		t.Log("📖 知识点: Go函数的基本语法 func name(params) returnType { body }")

		// 🔍 探索: 简单函数
		add := func(a, b int) int {
			return a + b
		}

		result := add(3, 5)
		t.Logf("🔍 函数调用: add(3, 5) = %d", result)

		// ✅ 验证函数行为
		if result != 8 {
			t.Errorf("❌ 函数计算错误: 期望8，得到%d", result)
		} else {
			t.Log("✅ 很好！您理解了基本函数的定义和调用")
		}

		// 💡 学习提示
		t.Log("💡 语法要点: 参数类型在参数名后面")
		t.Log("💡 语法要点: 返回类型在参数列表后面")
	})

	t.Run("学习多返回值函数", func(t *testing.T) {
		t.Log("📖 知识点: Go函数可以返回多个值，常用于返回结果和错误")

		// 🔍 探索: 多返回值函数
		divide := func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, errors.New("除数不能为零")
			}
			return a / b, nil
		}

		// 测试正常情况
		result, err := divide(10, 2)
		t.Logf("🔍 正常除法: divide(10, 2) = %.1f, error = %v", result, err)

		// 测试错误情况
		_, err2 := divide(10, 0)
		t.Logf("🔍 除零错误: divide(10, 0) error = %v", err2)

		// ✅ 验证多返回值
		if err != nil {
			t.Errorf("❌ 正常除法不应该有错误: %v", err)
		} else if result != 5.0 {
			t.Errorf("❌ 除法结果错误: 期望5.0，得到%.1f", result)
		} else {
			t.Log("✅ 很好！您理解了多返回值函数")
		}

		if err2 == nil {
			t.Error("❌ 除零应该返回错误")
		} else {
			t.Log("✅ 很好！您理解了错误处理")
		}

		// 💡 学习提示
		t.Log("💡 Go惯例: 最后一个返回值通常是error类型")
		t.Log("💡 最佳实践: 总是检查错误返回值")
	})

	t.Run("学习命名返回值", func(t *testing.T) {
		t.Log("📖 知识点: Go支持命名返回值，可以提高代码可读性")

		// 🔍 探索: 命名返回值
		calculateStats := func(numbers []int) (sum, count int, average float64) {
			for _, num := range numbers {
				sum += num
				count++
			}
			if count > 0 {
				average = float64(sum) / float64(count)
			}
			return // 裸返回，自动返回命名的变量
		}

		numbers := []int{1, 2, 3, 4, 5}
		sum, count, avg := calculateStats(numbers)

		t.Logf("🔍 统计计算: numbers = %v", numbers)
		t.Logf("   总和: %d", sum)
		t.Logf("   数量: %d", count)
		t.Logf("   平均值: %.1f", avg)

		// ✅ 验证命名返回值
		if sum != 15 || count != 5 || avg != 3.0 {
			t.Errorf("❌ 统计计算错误: sum=%d, count=%d, avg=%.1f", sum, count, avg)
		} else {
			t.Log("✅ 很好！您理解了命名返回值")
		}

		// 💡 学习提示
		t.Log("💡 优势: 命名返回值可以作为文档，说明返回值的含义")
		t.Log("💡 注意: 裸返回在长函数中可能降低可读性")
	})
}

// TestLearnParameterPassing 学习参数传递
func TestLearnParameterPassing(t *testing.T) {
	t.Log("🎯 学习目标: 理解Go的参数传递机制")
	t.Log("📚 本测试将教您: 值传递、引用类型、指针参数")

	t.Run("学习值传递", func(t *testing.T) {
		t.Log("📖 知识点: Go默认使用值传递，函数内修改不影响原变量")

		// 🔍 探索: 值传递行为
		modifyValue := func(x int) {
			x = x * 2
			t.Logf("   函数内部: x = %d", x)
		}

		original := 10
		t.Logf("🔍 值传递测试:")
		t.Logf("   调用前: original = %d", original)
		modifyValue(original)
		t.Logf("   调用后: original = %d", original)

		// ✅ 验证值传递
		if original != 10 {
			t.Errorf("❌ 值传递理解错误: 原值应该不变，但变成了%d", original)
		} else {
			t.Log("✅ 很好！您理解了值传递不会修改原变量")
		}

		// 💡 学习提示
		t.Log("💡 重要概念: 基本类型（int, float, bool等）都是值传递")
		t.Log("💡 性能考虑: 大结构体值传递可能影响性能")
	})

	t.Run("学习指针参数", func(t *testing.T) {
		t.Log("📖 知识点: 使用指针可以在函数内修改原变量")

		// 🔍 探索: 指针传递行为
		modifyByPointer := func(x *int) {
			*x = *x * 2
			t.Logf("   函数内部: *x = %d", *x)
		}

		original := 10
		t.Logf("🔍 指针传递测试:")
		t.Logf("   调用前: original = %d", original)
		modifyByPointer(&original)
		t.Logf("   调用后: original = %d", original)

		// ✅ 验证指针传递
		if original != 20 {
			t.Errorf("❌ 指针传递理解错误: 期望20，得到%d", original)
		} else {
			t.Log("✅ 很好！您理解了指针可以修改原变量")
		}

		// 💡 学习提示
		t.Log("💡 语法要点: &取地址，*解引用")
		t.Log("💡 使用场景: 需要修改原变量或避免大对象拷贝时使用指针")
	})

	t.Run("学习切片和映射的传递", func(t *testing.T) {
		t.Log("📖 知识点: 切片和映射是引用类型，修改会影响原数据")

		// 🔍 探索: 切片传递
		modifySlice := func(s []int) {
			if len(s) > 0 {
				s[0] = 999
			}
			t.Logf("   函数内修改切片: %v", s)
		}

		slice := []int{1, 2, 3, 4, 5}
		t.Logf("🔍 切片传递测试:")
		t.Logf("   调用前: slice = %v", slice)
		modifySlice(slice)
		t.Logf("   调用后: slice = %v", slice)

		// ✅ 验证切片传递
		if slice[0] != 999 {
			t.Errorf("❌ 切片传递理解错误: 期望第一个元素为999，得到%d", slice[0])
		} else {
			t.Log("✅ 很好！您理解了切片是引用类型")
		}

		// 💡 学习提示
		t.Log("💡 重要概念: 切片、映射、通道都是引用类型")
		t.Log("💡 注意事项: 修改引用类型的内容会影响原数据")
	})
}

// TestLearnVariadicFunctions 学习可变参数函数
func TestLearnVariadicFunctions(t *testing.T) {
	t.Log("🎯 学习目标: 掌握可变参数函数的使用")
	t.Log("📚 本测试将教您: ...语法、参数展开、实际应用")

	t.Run("学习可变参数基础", func(t *testing.T) {
		t.Log("📖 知识点: 使用...语法定义可变参数函数")

		// 🔍 探索: 可变参数函数
		sum := func(numbers ...int) int {
			total := 0
			t.Logf("   接收到%d个参数: %v", len(numbers), numbers)
			for _, num := range numbers {
				total += num
			}
			return total
		}

		// 测试不同数量的参数
		result1 := sum()
		result2 := sum(1)
		result3 := sum(1, 2, 3)
		result4 := sum(1, 2, 3, 4, 5)

		t.Logf("🔍 可变参数测试:")
		t.Logf("   sum() = %d", result1)
		t.Logf("   sum(1) = %d", result2)
		t.Logf("   sum(1,2,3) = %d", result3)
		t.Logf("   sum(1,2,3,4,5) = %d", result4)

		// ✅ 验证可变参数
		if result3 != 6 || result4 != 15 {
			t.Errorf("❌ 可变参数计算错误")
		} else {
			t.Log("✅ 很好！您理解了可变参数函数")
		}

		// 💡 学习提示
		t.Log("💡 语法要点: ...type表示可变参数")
		t.Log("💡 内部实现: 可变参数在函数内部是切片")
	})

	t.Run("学习参数展开", func(t *testing.T) {
		t.Log("📖 知识点: 使用...可以展开切片作为参数")

		// 🔍 探索: 参数展开
		max := func(numbers ...int) int {
			if len(numbers) == 0 {
				return 0
			}
			maxVal := numbers[0]
			for _, num := range numbers[1:] {
				if num > maxVal {
					maxVal = num
				}
			}
			return maxVal
		}

		// 直接传递参数
		result1 := max(3, 1, 4, 1, 5, 9)

		// 展开切片
		numbers := []int{3, 1, 4, 1, 5, 9}
		result2 := max(numbers...)

		t.Logf("🔍 参数展开测试:")
		t.Logf("   直接传递: max(3,1,4,1,5,9) = %d", result1)
		t.Logf("   展开切片: max(numbers...) = %d", result2)

		// ✅ 验证参数展开
		if result1 != result2 || result1 != 9 {
			t.Errorf("❌ 参数展开理解错误")
		} else {
			t.Log("✅ 很好！您理解了参数展开")
		}

		// 💡 学习提示
		t.Log("💡 语法要点: slice...可以展开切片")
		t.Log("💡 应用场景: 将切片数据传递给可变参数函数")
	})
}

// TestLearnHigherOrderFunctions 学习高阶函数
func TestLearnHigherOrderFunctions(t *testing.T) {
	t.Log("🎯 学习目标: 理解函数作为一等公民的概念")
	t.Log("📚 本测试将教您: 函数变量、函数参数、函数返回值")

	t.Run("学习函数作为变量", func(t *testing.T) {
		t.Log("📖 知识点: 函数可以赋值给变量，像其他值一样使用")

		// 🔍 探索: 函数变量
		var operation func(int, int) int

		add := func(a, b int) int { return a + b }
		multiply := func(a, b int) int { return a * b }

		// 动态选择函数
		operation = add
		result1 := operation(3, 4)

		operation = multiply
		result2 := operation(3, 4)

		t.Logf("🔍 函数变量测试:")
		t.Logf("   使用add函数: 3 + 4 = %d", result1)
		t.Logf("   使用multiply函数: 3 × 4 = %d", result2)

		// ✅ 验证函数变量
		if result1 != 7 || result2 != 12 {
			t.Errorf("❌ 函数变量理解错误")
		} else {
			t.Log("✅ 很好！您理解了函数可以作为变量")
		}

		// 💡 学习提示
		t.Log("💡 重要概念: 函数是一等公民，可以像其他值一样使用")
		t.Log("💡 应用场景: 策略模式、回调函数、事件处理")
	})

	t.Run("学习函数作为参数", func(t *testing.T) {
		t.Log("📖 知识点: 函数可以作为参数传递给其他函数")

		// 🔍 探索: 函数作为参数
		applyOperation := func(a, b int, op func(int, int) int) int {
			result := op(a, b)
			t.Logf("   执行操作: %d op %d = %d", a, b, result)
			return result
		}

		add := func(a, b int) int { return a + b }
		subtract := func(a, b int) int { return a - b }

		result1 := applyOperation(10, 3, add)
		result2 := applyOperation(10, 3, subtract)

		t.Logf("🔍 函数作为参数测试:")
		t.Logf("   传递add函数: %d", result1)
		t.Logf("   传递subtract函数: %d", result2)

		// ✅ 验证函数参数
		if result1 != 13 || result2 != 7 {
			t.Errorf("❌ 函数参数理解错误")
		} else {
			t.Log("✅ 很好！您理解了函数可以作为参数")
		}

		// 💡 学习提示
		t.Log("💡 设计模式: 这是策略模式的实现方式")
		t.Log("💡 实际应用: sort.Slice、http.HandleFunc等都使用了这种模式")
	})

	t.Run("学习函数作为返回值", func(t *testing.T) {
		t.Log("📖 知识点: 函数可以返回其他函数，形成闭包")

		// 🔍 探索: 函数返回函数
		makeMultiplier := func(factor int) func(int) int {
			return func(x int) int {
				return x * factor
			}
		}

		double := makeMultiplier(2)
		triple := makeMultiplier(3)

		result1 := double(5)
		result2 := triple(5)

		t.Logf("🔍 函数返回函数测试:")
		t.Logf("   double(5) = %d", result1)
		t.Logf("   triple(5) = %d", result2)

		// ✅ 验证函数返回值
		if result1 != 10 || result2 != 15 {
			t.Errorf("❌ 函数返回值理解错误")
		} else {
			t.Log("✅ 很好！您理解了函数可以返回函数")
		}

		// 💡 学习提示
		t.Log("💡 重要概念: 这创建了闭包，内部函数可以访问外部变量")
		t.Log("💡 应用场景: 工厂函数、装饰器模式、中间件")
	})
}

// TestLearnRecursion 学习递归
func TestLearnRecursion(t *testing.T) {
	t.Log("🎯 学习目标: 掌握递归函数的设计和使用")
	t.Log("📚 本测试将教您: 递归基础、终止条件、尾递归优化")

	t.Run("学习基础递归", func(t *testing.T) {
		t.Log("📖 知识点: 递归函数调用自身来解决问题")

		// 🔍 探索: 阶乘递归
		var factorial func(int) int
		factorial = func(n int) int {
			t.Logf("   计算 factorial(%d)", n)
			if n <= 1 {
				t.Logf("   基础情况: factorial(%d) = 1", n)
				return 1
			}
			result := n * factorial(n-1)
			t.Logf("   递归情况: factorial(%d) = %d × factorial(%d) = %d", n, n, n-1, result)
			return result
		}

		t.Log("🔍 递归计算阶乘:")
		result := factorial(5)
		t.Logf("   最终结果: 5! = %d", result)

		// ✅ 验证递归
		if result != 120 {
			t.Errorf("❌ 递归计算错误: 期望120，得到%d", result)
		} else {
			t.Log("✅ 很好！您理解了递归的基本原理")
		}

		// 💡 学习提示
		t.Log("💡 递归要素: 基础情况（终止条件）+ 递归情况（自我调用）")
		t.Log("💡 注意事项: 必须有明确的终止条件，否则会无限递归")
	})

	t.Run("学习斐波那契数列", func(t *testing.T) {
		t.Log("📖 知识点: 斐波那契数列是递归的经典例子")

		// 🔍 探索: 斐波那契递归（效率较低）
		var fibSlow func(int) int
		fibSlow = func(n int) int {
			if n <= 1 {
				return n
			}
			return fibSlow(n-1) + fibSlow(n-2)
		}

		// 🔍 探索: 优化的斐波那契（记忆化）
		memo := make(map[int]int)
		var fibFast func(int) int
		fibFast = func(n int) int {
			if n <= 1 {
				return n
			}
			if val, exists := memo[n]; exists {
				return val
			}
			result := fibFast(n-1) + fibFast(n-2)
			memo[n] = result
			return result
		}

		n := 10
		result1 := fibSlow(n)
		result2 := fibFast(n)

		t.Logf("🔍 斐波那契数列第%d项:", n)
		t.Logf("   普通递归: %d", result1)
		t.Logf("   记忆化递归: %d", result2)

		// ✅ 验证斐波那契
		if result1 != result2 || result1 != 55 {
			t.Errorf("❌ 斐波那契计算错误")
		} else {
			t.Log("✅ 很好！您理解了递归优化")
		}

		// 💡 学习提示
		t.Log("💡 性能优化: 记忆化可以避免重复计算")
		t.Log("💡 替代方案: 某些递归可以用迭代替代，提高效率")
	})
}

// BenchmarkLearnFunctionPerformance 学习函数性能
func BenchmarkLearnFunctionPerformance(b *testing.B) {
	b.Log("🎯 学习目标: 了解不同函数实现的性能差异")

	// 递归 vs 迭代性能比较
	b.Run("斐波那契性能比较", func(b *testing.B) {
		// 递归版本
		var fibRecursive func(int) int
		fibRecursive = func(n int) int {
			if n <= 1 {
				return n
			}
			return fibRecursive(n-1) + fibRecursive(n-2)
		}

		// 迭代版本
		fibIterative := func(n int) int {
			if n <= 1 {
				return n
			}
			a, b := 0, 1
			for i := 2; i <= n; i++ {
				a, b = b, a+b
			}
			return b
		}

		n := 20

		b.Run("递归版本", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibRecursive(n)
			}
		})

		b.Run("迭代版本", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibIterative(n)
			}
		})
	})
}

// Example_learnFunctionBasics 函数基础示例
func Example_learnFunctionBasics() {
	// 定义一个简单的函数
	greet := func(name string) string {
		return fmt.Sprintf("Hello, %s!", name)
	}

	// 调用函数
	message := greet("Go")
	fmt.Println(message)

	// Output:
	// Hello, Go!
}
