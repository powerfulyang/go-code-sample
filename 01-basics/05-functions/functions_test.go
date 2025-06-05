package functions

import (
	"strings"
	"testing"
)

func TestBasicFunctions(t *testing.T) {
	t.Run("SimpleFunction", func(t *testing.T) {
		// 测试简单的加法函数
		result := add(10, 20)
		expected := 30

		t.Logf("add(10, 20) = %d", result)

		if result != expected {
			t.Errorf("期望 %d, 实际 %d", expected, result)
		}
	})

	t.Run("CalculateArea", func(t *testing.T) {
		// 测试圆面积计算
		radius := 5.0
		area := calculateArea(radius)

		t.Logf("半径 %.1f 的圆面积: %.2f", radius, area)

		// 验证面积计算（π * r²）
		if area < 78 || area > 79 { // 大约 78.54
			t.Errorf("圆面积计算可能有误: %.2f", area)
		}
	})
}

func TestMultipleReturnValues(t *testing.T) {
	t.Run("DivideFunction", func(t *testing.T) {
		// 测试除法函数
		quotient, remainder := divide(17, 5)

		t.Logf("17 ÷ 5 = %d 余 %d", quotient, remainder)

		if quotient != 3 || remainder != 2 {
			t.Errorf("除法计算错误: 期望 3 余 2, 实际 %d 余 %d", quotient, remainder)
		}
	})

	t.Run("SafeDivideSuccess", func(t *testing.T) {
		// 测试安全除法 - 成功情况
		result, err := safeDivide(10, 2)

		t.Logf("safeDivide(10, 2) = %.2f, error = %v", result, err)

		if err != nil {
			t.Errorf("不应该有错误: %v", err)
		}
		if result != 5.0 {
			t.Errorf("期望 5.0, 实际 %.2f", result)
		}
	})

	t.Run("SafeDivideError", func(t *testing.T) {
		// 测试安全除法 - 错误情况
		result, err := safeDivide(10, 0)

		t.Logf("safeDivide(10, 0) = %.2f, error = %v", result, err)

		if err == nil {
			t.Error("应该有除零错误")
		}
		if result != 0 {
			t.Errorf("错误时应该返回 0, 实际 %.2f", result)
		}
	})

	t.Run("IgnoreReturnValue", func(t *testing.T) {
		// 测试忽略返回值
		name, _ := getNameAndAge()

		t.Logf("只获取姓名: %s", name)

		if name != "李四" {
			t.Errorf("期望 '李四', 实际 '%s'", name)
		}
	})
}

func TestNamedReturnValues(t *testing.T) {
	t.Run("FindMinMax", func(t *testing.T) {
		numbers := []int{3, 7, 1, 9, 2, 8}
		min, max := findMinMax(numbers)

		t.Logf("数组 %v 的最小值: %d, 最大值: %d", numbers, min, max)

		if min != 1 || max != 9 {
			t.Errorf("期望最小值 1, 最大值 9, 实际最小值 %d, 最大值 %d", min, max)
		}
	})

	t.Run("FindMinMaxEmpty", func(t *testing.T) {
		// 测试空数组
		var numbers []int
		min, max := findMinMax(numbers)

		t.Logf("空数组的最小值: %d, 最大值: %d", min, max)

		if min != 0 || max != 0 {
			t.Errorf("空数组应该返回 0, 0, 实际 %d, %d", min, max)
		}
	})

	t.Run("RectangleInfo", func(t *testing.T) {
		width, height := 5.0, 3.0
		area, perimeter := rectangleInfo(width, height)

		t.Logf("矩形 %.1f×%.1f 的面积: %.2f, 周长: %.2f", width, height, area, perimeter)

		expectedArea := 15.0
		expectedPerimeter := 16.0

		if area != expectedArea {
			t.Errorf("面积错误: 期望 %.2f, 实际 %.2f", expectedArea, area)
		}
		if perimeter != expectedPerimeter {
			t.Errorf("周长错误: 期望 %.2f, 实际 %.2f", expectedPerimeter, perimeter)
		}
	})
}

func TestVariadicFunctions(t *testing.T) {
	t.Run("SumFunction", func(t *testing.T) {
		// 测试可变参数求和
		result1 := sum(1, 2, 3, 4, 5)
		result2 := sum(10, 20)
		result3 := sum()

		t.Logf("sum(1,2,3,4,5) = %d", result1)
		t.Logf("sum(10,20) = %d", result2)
		t.Logf("sum() = %d", result3)

		if result1 != 15 {
			t.Errorf("期望 15, 实际 %d", result1)
		}
		if result2 != 30 {
			t.Errorf("期望 30, 实际 %d", result2)
		}
		if result3 != 0 {
			t.Errorf("期望 0, 实际 %d", result3)
		}
	})

	t.Run("SumWithSlice", func(t *testing.T) {
		// 测试传递切片
		numbers := []int{1, 2, 3, 4, 5}
		result := sum(numbers...)

		t.Logf("sum(%v...) = %d", numbers, result)

		if result != 15 {
			t.Errorf("期望 15, 实际 %d", result)
		}
	})
}

func TestHigherOrderFunctions(t *testing.T) {
	t.Run("MapFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		doubled := mapInts(numbers, func(x int) int { return x * 2 })

		t.Logf("原数组: %v", numbers)
		t.Logf("翻倍后: %v", doubled)

		expected := []int{2, 4, 6, 8, 10}
		for i, v := range doubled {
			if v != expected[i] {
				t.Errorf("索引 %d: 期望 %d, 实际 %d", i, expected[i], v)
			}
		}
	})

	t.Run("FilterFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		evens := filterInts(numbers, func(x int) bool { return x%2 == 0 })

		t.Logf("原数组: %v", numbers)
		t.Logf("偶数: %v", evens)

		expected := []int{2, 4, 6}
		if len(evens) != len(expected) {
			t.Errorf("长度不匹配: 期望 %d, 实际 %d", len(expected), len(evens))
		}
	})

	t.Run("ReduceFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		total := reduceInts(numbers, 0, func(acc, x int) int { return acc + x })

		t.Logf("数组 %v 的和: %d", numbers, total)

		if total != 15 {
			t.Errorf("期望 15, 实际 %d", total)
		}
	})

	t.Run("CreateMultiplier", func(t *testing.T) {
		multiplier := createMultiplier(3)
		result := multiplier(10)

		t.Logf("3 倍乘法器: 3 × 10 = %d", result)

		if result != 30 {
			t.Errorf("期望 30, 实际 %d", result)
		}
	})
}

func TestClosures(t *testing.T) {
	t.Run("Counter", func(t *testing.T) {
		counter := createCounter()

		// 测试计数器
		results := []int{counter(), counter(), counter()}
		expected := []int{1, 2, 3}

		t.Logf("计数器结果: %v", results)

		for i, result := range results {
			if result != expected[i] {
				t.Errorf("计数 %d: 期望 %d, 实际 %d", i+1, expected[i], result)
			}
		}
	})

	t.Run("IndependentCounters", func(t *testing.T) {
		counter1 := createCounter()
		counter2 := createCounter()

		// 测试独立的计数器
		result1 := counter1()
		result2 := counter2()
		result3 := counter1()

		t.Logf("计数器1第1次: %d", result1)
		t.Logf("计数器2第1次: %d", result2)
		t.Logf("计数器1第2次: %d", result3)

		if result1 != 1 || result2 != 1 || result3 != 2 {
			t.Error("计数器应该是独立的")
		}
	})

	t.Run("Adder", func(t *testing.T) {
		adder := createAdder(10)

		result1 := adder(5)
		result2 := adder(3)

		t.Logf("累加器 (初始值 10): +5 = %d, +3 = %d", result1, result2)

		if result1 != 15 || result2 != 18 {
			t.Errorf("累加器错误: 期望 15, 18, 实际 %d, %d", result1, result2)
		}
	})

	t.Run("StringProcessor", func(t *testing.T) {
		processor := createStringProcessor("【", "】")
		result := processor("重要")

		t.Logf("字符串处理器: %s", result)

		expected := "【重要】"
		if result != expected {
			t.Errorf("期望 '%s', 实际 '%s'", expected, result)
		}
	})
}

func TestRecursiveFunctions(t *testing.T) {
	t.Run("Factorial", func(t *testing.T) {
		testCases := []struct {
			input    int
			expected int
		}{
			{0, 1},
			{1, 1},
			{5, 120},
			{6, 720},
		}

		for _, tc := range testCases {
			result := factorial(tc.input)
			t.Logf("%d! = %d", tc.input, result)

			if result != tc.expected {
				t.Errorf("%d!: 期望 %d, 实际 %d", tc.input, tc.expected, result)
			}
		}
	})

	t.Run("Fibonacci", func(t *testing.T) {
		// 测试斐波那契数列前几项
		expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

		for i, exp := range expected {
			result := fibonacci(i)
			t.Logf("fibonacci(%d) = %d", i, result)

			if result != exp {
				t.Errorf("fibonacci(%d): 期望 %d, 实际 %d", i, exp, result)
			}
		}
	})

	t.Run("BinarySearch", func(t *testing.T) {
		sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15}

		testCases := []struct {
			target   int
			expected int
		}{
			{7, 3},   // 找到
			{1, 0},   // 第一个元素
			{15, 7},  // 最后一个元素
			{4, -1},  // 不存在
			{16, -1}, // 超出范围
		}

		for _, tc := range testCases {
			result := binarySearch(sortedArray, tc.target, 0, len(sortedArray)-1)
			t.Logf("在 %v 中查找 %d: 索引 %d", sortedArray, tc.target, result)

			if result != tc.expected {
				t.Errorf("查找 %d: 期望索引 %d, 实际 %d", tc.target, tc.expected, result)
			}
		}
	})
}

func TestAnonymousFunctions(t *testing.T) {
	t.Run("ImmediateExecution", func(t *testing.T) {
		// 立即执行的匿名函数
		result := func(a, b int) int {
			return a * b
		}(5, 6)

		t.Logf("匿名函数 5 × 6 = %d", result)

		if result != 30 {
			t.Errorf("期望 30, 实际 %d", result)
		}
	})

	t.Run("AssignToVariable", func(t *testing.T) {
		// 赋值给变量的匿名函数
		square := func(x int) int {
			return x * x
		}

		result := square(7)
		t.Logf("7 的平方 = %d", result)

		if result != 49 {
			t.Errorf("期望 49, 实际 %d", result)
		}
	})

	t.Run("FunctionSlice", func(t *testing.T) {
		// 函数切片
		operations := []func(int, int) int{
			func(a, b int) int { return a + b },
			func(a, b int) int { return a - b },
			func(a, b int) int { return a * b },
			func(a, b int) int { return a / b },
		}

		a, b := 20, 4
		expected := []int{24, 16, 80, 5}

		for i, op := range operations {
			result := op(a, b)
			t.Logf("操作 %d: %d op %d = %d", i, a, b, result)

			if result != expected[i] {
				t.Errorf("操作 %d: 期望 %d, 实际 %d", i, expected[i], result)
			}
		}
	})
}

func TestPracticalExamples(t *testing.T) {
	t.Run("StringProcessing", func(t *testing.T) {
		text := "  Hello, World!  "
		processed := processString(text,
			strings.TrimSpace,
			strings.ToLower,
			func(s string) string { return strings.ReplaceAll(s, "world", "go") },
		)

		t.Logf("原文本: '%s'", text)
		t.Logf("处理后: '%s'", processed)

		expected := "hello, go!"
		if processed != expected {
			t.Errorf("期望 '%s', 实际 '%s'", expected, processed)
		}
	})

	t.Run("Calculator", func(t *testing.T) {
		calculator := createCalculator()

		testCases := []struct {
			op       string
			a, b     float64
			expected float64
		}{
			{"add", 5, 3, 8},
			{"subtract", 10, 4, 6},
			{"multiply", 6, 7, 42},
			{"divide", 15, 3, 5},
		}

		for _, tc := range testCases {
			result := calculator(tc.op, tc.a, tc.b)
			t.Logf("计算器: %.0f %s %.0f = %.2f", tc.a, tc.op, tc.b, result)

			if result != tc.expected {
				t.Errorf("%s 操作: 期望 %.2f, 实际 %.2f", tc.op, tc.expected, result)
			}
		}
	})

	t.Run("UserValidation", func(t *testing.T) {
		users := []struct {
			user       map[string]interface{}
			shouldPass bool
		}{
			{
				map[string]interface{}{"name": "张三", "age": 25, "email": "zhangsan@example.com"},
				true,
			},
			{
				map[string]interface{}{"name": "", "age": 17, "email": "invalid-email"},
				false,
			},
			{
				map[string]interface{}{"name": "李四", "age": 30, "email": "lisi@example.com"},
				true,
			},
		}

		for i, test := range users {
			err := validateUser(test.user)
			t.Logf("用户 %d 验证: %v", i+1, err)

			if test.shouldPass && err != nil {
				t.Errorf("用户 %d 应该通过验证, 但失败了: %v", i+1, err)
			}
			if !test.shouldPass && err == nil {
				t.Errorf("用户 %d 应该验证失败, 但通过了", i+1)
			}
		}
	})
}

// 基准测试示例
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(10)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(20)
	}
}

func BenchmarkSum(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum(numbers...)
	}
}
