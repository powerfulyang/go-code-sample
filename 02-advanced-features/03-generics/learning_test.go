package generics

import (
	"fmt"
	"testing"
)

// 🎓 学习导向的测试 - 通过测试学习Go泛型

// MaxSlice 泛型函数：找到切片中的最大值
func MaxSlice[T Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	maxVal := slice[0]
	for _, v := range slice[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// MinValue 泛型函数：找到两个值中的较小值
func MinValue[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// LearningPerson 学习用的Person类型
type LearningPerson struct {
	Name string
	Age  int
}

func (p LearningPerson) String() string {
	return fmt.Sprintf("%s (%d岁)", p.Name, p.Age)
}

// FormatAll 格式化所有实现了String方法的元素
func FormatAll[T Stringer](items []T) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = item.String()
	}
	return result
}

// Container 泛型容器
type Container[T any] struct {
	items []T
}

// NewContainer 创建新的容器
func NewContainer[T any]() *Container[T] {
	return &Container[T]{
		items: make([]T, 0),
	}
}

// Add 添加元素
func (c *Container[T]) Add(item T) {
	c.items = append(c.items, item)
}

// Size 获取容器大小
func (c *Container[T]) Size() int {
	return len(c.items)
}

// Contains 检查是否包含元素
func (c *Container[T]) Contains(item T) bool {
	for _, existing := range c.items {
		// 这里需要使用反射或者要求T实现comparable
		// 为了简化，我们假设T是comparable
		if any(existing) == any(item) {
			return true
		}
	}
	return false
}

// ForEach 遍历元素
func (c *Container[T]) ForEach(fn func(T)) {
	for _, item := range c.items {
		fn(item)
	}
}

// TestLearnBasicGenerics 学习泛型基础
func TestLearnBasicGenerics(t *testing.T) {
	t.Log("🎯 学习目标: 理解Go 1.18引入的泛型特性")
	t.Log("📚 本测试将教您: 类型参数、约束、泛型函数和类型")

	t.Run("学习泛型函数基础", func(t *testing.T) {
		t.Log("📖 知识点: 泛型函数可以处理多种类型，提高代码复用性")

		// 🔍 探索: 泛型函数的定义和使用
		// 测试不同类型
		intSlice := []int{3, 1, 4, 1, 5, 9}
		floatSlice := []float64{3.14, 2.71, 1.41, 1.73}
		stringSlice := []string{"apple", "banana", "cherry", "date"}

		maxInt := MaxSlice(intSlice)
		maxFloat := MaxSlice(floatSlice)
		maxString := MaxSlice(stringSlice)

		t.Logf("🔍 泛型函数测试:")
		t.Logf("   整数切片 %v 的最大值: %d", intSlice, maxInt)
		t.Logf("   浮点切片 %v 的最大值: %.2f", floatSlice, maxFloat)
		t.Logf("   字符串切片 %v 的最大值: %s", stringSlice, maxString)

		// ✅ 验证泛型函数
		if maxInt != 9 {
			t.Errorf("❌ 整数最大值错误: 期望9，得到%d", maxInt)
		}
		if maxFloat != 3.14 {
			t.Errorf("❌ 浮点最大值错误: 期望3.14，得到%.2f", maxFloat)
		}
		if maxString != "date" {
			t.Errorf("❌ 字符串最大值错误: 期望'date'，得到'%s'", maxString)
		}

		t.Log("✅ 很好！您理解了泛型函数的基本使用")

		// 💡 学习提示
		t.Log("💡 语法要点: func[T constraint](params) returnType")
		t.Log("💡 类型推断: Go可以自动推断类型参数")
		t.Log("💡 约束: comparable约束允许使用==和!=操作符")
	})

	t.Run("学习泛型类型", func(t *testing.T) {
		t.Log("📖 知识点: 泛型类型可以创建类型安全的数据结构")

		// 🔍 探索: 泛型栈的实现和使用
		// 使用已定义的泛型栈
		intStack := NewStack[int]()
		stringStack := NewStack[string]()

		// 测试整数栈
		intStack.Push(1)
		intStack.Push(2)
		intStack.Push(3)

		t.Logf("🔍 整数栈操作:")
		t.Logf("   栈大小: %d", intStack.Size())

		val, ok := intStack.Pop()
		if ok {
			t.Logf("   弹出: %d", val)
		}

		val, ok = intStack.Peek()
		if ok {
			t.Logf("   栈顶: %d", val)
		}

		// 测试字符串栈
		stringStack.Push("hello")
		stringStack.Push("world")

		t.Logf("🔍 字符串栈操作:")
		t.Logf("   栈大小: %d", stringStack.Size())

		str, ok := stringStack.Pop()
		if ok {
			t.Logf("   弹出: %s", str)
		}

		// ✅ 验证泛型类型
		if intStack.Size() != 2 {
			t.Errorf("❌ 整数栈大小错误: 期望2，得到%d", intStack.Size())
		}
		if stringStack.Size() != 1 {
			t.Errorf("❌ 字符串栈大小错误: 期望1，得到%d", stringStack.Size())
		}
		if str != "world" {
			t.Errorf("❌ 弹出的字符串错误: 期望'world'，得到'%s'", str)
		}

		t.Log("✅ 很好！您理解了泛型类型的使用")

		// 💡 学习提示
		t.Log("💡 类型安全: 泛型提供编译时类型检查")
		t.Log("💡 代码复用: 一个实现可以处理多种类型")
		t.Log("💡 性能: 泛型避免了interface{}的装箱开销")
	})
}

// TestLearnGenericConstraints 学习泛型约束
func TestLearnGenericConstraints(t *testing.T) {
	t.Log("🎯 学习目标: 掌握泛型约束的使用")
	t.Log("📚 本测试将教您: 内置约束、自定义约束、类型集合")

	t.Run("学习内置约束", func(t *testing.T) {
		t.Log("📖 知识点: Go提供了一些内置的约束类型")

		// 🔍 探索: 使用内置约束
		// 使用已定义的数值求和函数
		intSum := Sum([]int{1, 2, 3, 4, 5})
		floatSum := Sum([]float64{1.1, 2.2, 3.3})

		t.Logf("🔍 内置约束测试:")
		t.Logf("   整数求和: %d", intSum)
		t.Logf("   浮点求和: %.1f", floatSum)

		// ✅ 验证内置约束
		if intSum != 15 {
			t.Errorf("❌ 整数求和错误: 期望15，得到%d", intSum)
		}
		if floatSum != 6.6 {
			t.Errorf("❌ 浮点求和错误: 期望6.6，得到%.1f", floatSum)
		}

		t.Log("✅ 很好！您理解了内置约束的使用")

		// 💡 学习提示
		t.Log("💡 常用约束: any, comparable, Ordered")
		t.Log("💡 golang.org/x/exp/constraints包提供更多约束")
		t.Log("💡 约束组合: 可以使用|组合多个约束")
	})

	t.Run("学习自定义约束", func(t *testing.T) {
		t.Log("📖 知识点: 可以定义自己的约束来限制类型参数")

		// 🔍 探索: 自定义约束的使用
		// 测试自定义约束
		people := []LearningPerson{
			{"张三", 25},
			{"李四", 30},
			{"王五", 28},
		}

		result := FormatAll(people)

		t.Logf("🔍 自定义约束测试:")
		for i, formatted := range result {
			t.Logf("   格式化结果[%d]: %s", i, formatted)
		}

		// ✅ 验证自定义约束
		if len(result) != 3 {
			t.Errorf("❌ 格式化结果数量错误: 期望3，得到%d", len(result))
		}
		if result[0] != "张三 (25岁)" {
			t.Errorf("❌ 格式化结果错误: 期望'张三 (25岁)'，得到'%s'", result[0])
		}

		t.Log("✅ 很好！您理解了自定义约束的使用")

		// 💡 学习提示
		t.Log("💡 接口约束: 约束可以是接口类型")
		t.Log("💡 方法约束: 约束可以要求特定的方法")
		t.Log("💡 灵活性: 自定义约束提供更精确的类型控制")
	})
}

// TestLearnGenericPatterns 学习泛型模式
func TestLearnGenericPatterns(t *testing.T) {
	t.Log("🎯 学习目标: 掌握常用的泛型编程模式")
	t.Log("📚 本测试将教您: 泛型容器、函数式编程、类型转换")

	t.Run("学习泛型容器模式", func(t *testing.T) {
		t.Log("📖 知识点: 泛型非常适合实现类型安全的容器")

		// 🔍 探索: 泛型容器的使用
		// 使用已定义的泛型容器
		container := NewContainer[string]()

		// 添加元素
		container.Add("apple")
		container.Add("banana")
		container.Add("cherry")

		t.Logf("🔍 泛型容器测试:")
		t.Logf("   容器大小: %d", container.Size())

		// 遍历元素
		t.Log("   容器内容:")
		container.ForEach(func(item string) {
			t.Logf("     - %s", item)
		})

		// 查找元素
		found := container.Contains("banana")
		t.Logf("   包含'banana': %t", found)

		// ✅ 验证泛型容器
		if container.Size() != 3 {
			t.Errorf("❌ 容器大小错误: 期望3，得到%d", container.Size())
		}
		if !found {
			t.Error("❌ 应该找到'banana'")
		}

		t.Log("✅ 很好！您理解了泛型容器模式")

		// 💡 学习提示
		t.Log("💡 类型安全: 编译时就能发现类型错误")
		t.Log("💡 性能优势: 避免了类型断言的运行时开销")
		t.Log("💡 API设计: 泛型让API更加清晰和易用")
	})

	t.Run("学习函数式编程模式", func(t *testing.T) {
		t.Log("📖 知识点: 泛型使函数式编程模式更加类型安全")

		// 🔍 探索: 泛型函数式编程
		numbers := []int{1, 2, 3, 4, 5}

		// Map操作：将数字转换为字符串
		strings := Map(numbers, func(n int) string {
			return fmt.Sprintf("数字%d", n)
		})

		// Filter操作：过滤偶数
		evens := Filter(numbers, func(n int) bool {
			return n%2 == 0
		})

		// Reduce操作：计算乘积
		product := Reduce(numbers, 1, func(acc, n int) int {
			return acc * n
		})

		t.Logf("🔍 函数式编程测试:")
		t.Logf("   原始数据: %v", numbers)
		t.Logf("   Map结果: %v", strings)
		t.Logf("   Filter结果: %v", evens)
		t.Logf("   Reduce结果: %d", product)

		// ✅ 验证函数式编程
		if len(strings) != 5 {
			t.Errorf("❌ Map结果数量错误: 期望5，得到%d", len(strings))
		}
		if len(evens) != 2 {
			t.Errorf("❌ Filter结果数量错误: 期望2，得到%d", len(evens))
		}
		if product != 120 {
			t.Errorf("❌ Reduce结果错误: 期望120，得到%d", product)
		}

		t.Log("✅ 很好！您理解了泛型函数式编程")

		// 💡 学习提示
		t.Log("💡 高阶函数: 泛型让高阶函数更加类型安全")
		t.Log("💡 链式调用: 可以组合多个操作形成数据处理管道")
		t.Log("💡 不可变性: 函数式编程鼓励不可变数据结构")
	})
}

// BenchmarkLearnGenericsPerformance 学习泛型性能
func BenchmarkLearnGenericsPerformance(b *testing.B) {
	b.Log("🎯 学习目标: 了解泛型的性能特征")

	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	b.Run("泛型函数", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Sum(data)
		}
	})

	b.Run("interface{}版本", func(b *testing.B) {
		interfaceSum := func(slice []interface{}) int {
			sum := 0
			for _, v := range slice {
				sum += v.(int)
			}
			return sum
		}

		interfaceData := make([]interface{}, len(data))
		for i, v := range data {
			interfaceData[i] = v
		}

		for i := 0; i < b.N; i++ {
			_ = interfaceSum(interfaceData)
		}
	})
}

// Example_learnBasicGenerics 泛型基础示例
func Example_learnBasicGenerics() {
	// 使用泛型函数
	fmt.Println("整数最小值:", MinValue(3, 5))
	fmt.Println("字符串最小值:", MinValue("apple", "banana"))

	// Output:
	// 整数最小值: 3
	// 字符串最小值: apple
}
