package generics

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestBasicGenericFunctions(t *testing.T) {
	t.Run("MaxFunction", func(t *testing.T) {
		// 测试整数
		if result := Max(10, 20); result != 20 {
			t.Errorf("Max(10, 20) = %d; want 20", result)
		}

		// 测试浮点数
		if result := Max(3.14, 2.71); result != 3.14 {
			t.Errorf("Max(3.14, 2.71) = %f; want 3.14", result)
		}

		// 测试字符串
		if result := Max("apple", "banana"); result != "banana" {
			t.Errorf("Max(\"apple\", \"banana\") = %s; want \"banana\"", result)
		}

		t.Logf("Max函数测试通过")
	})

	t.Run("MinFunction", func(t *testing.T) {
		// 测试整数
		if result := Min(10, 20); result != 10 {
			t.Errorf("Min(10, 20) = %d; want 10", result)
		}

		// 测试浮点数
		if result := Min(3.14, 2.71); result != 2.71 {
			t.Errorf("Min(3.14, 2.71) = %f; want 2.71", result)
		}

		t.Logf("Min函数测试通过")
	})

	t.Run("SwapFunction", func(t *testing.T) {
		a, b := 100, 200
		originalA, originalB := a, b

		Swap(&a, &b)

		if a != originalB || b != originalA {
			t.Errorf("Swap failed: a=%d, b=%d; want a=%d, b=%d", a, b, originalB, originalA)
		}

		t.Logf("交换前: a=%d, b=%d", originalA, originalB)
		t.Logf("交换后: a=%d, b=%d", a, b)
	})
}

func TestGenericSliceOperations(t *testing.T) {
	t.Run("ContainsFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		if !Contains(numbers, 3) {
			t.Error("Contains should return true for existing element")
		}

		if Contains(numbers, 10) {
			t.Error("Contains should return false for non-existing element")
		}

		t.Logf("Contains函数测试通过")
	})

	t.Run("FilterFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evens := Filter(numbers, func(n int) bool { return n%2 == 0 })

		expected := []int{2, 4, 6, 8, 10}
		if !reflect.DeepEqual(evens, expected) {
			t.Errorf("Filter evens = %v; want %v", evens, expected)
		}

		t.Logf("过滤结果: %v", evens)
	})

	t.Run("MapFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		squares := Map(numbers, func(n int) int { return n * n })

		expected := []int{1, 4, 9, 16, 25}
		if !reflect.DeepEqual(squares, expected) {
			t.Errorf("Map squares = %v; want %v", squares, expected)
		}

		t.Logf("映射结果: %v", squares)
	})

	t.Run("ReduceFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })

		expected := 15
		if sum != expected {
			t.Errorf("Reduce sum = %d; want %d", sum, expected)
		}

		t.Logf("归约结果: %d", sum)
	})

	t.Run("FindFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		// 查找存在的元素
		if found, ok := Find(numbers, func(n int) bool { return n > 3 }); !ok || found != 4 {
			t.Errorf("Find should return (4, true), got (%d, %t)", found, ok)
		}

		// 查找不存在的元素
		if _, ok := Find(numbers, func(n int) bool { return n > 10 }); ok {
			t.Error("Find should return (0, false) for non-existing element")
		}

		t.Logf("Find函数测试通过")
	})

	t.Run("ReverseFunction", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		reversed := Reverse(numbers)

		expected := []int{5, 4, 3, 2, 1}
		if !reflect.DeepEqual(reversed, expected) {
			t.Errorf("Reverse = %v; want %v", reversed, expected)
		}

		t.Logf("反转结果: %v", reversed)
	})

	t.Run("UniqueFunction", func(t *testing.T) {
		duplicates := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
		unique := Unique(duplicates)

		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(unique, expected) {
			t.Errorf("Unique = %v; want %v", unique, expected)
		}

		t.Logf("去重前: %v", duplicates)
		t.Logf("去重后: %v", unique)
	})
}

func TestGenericDataStructures(t *testing.T) {
	t.Run("Stack", func(t *testing.T) {
		stack := NewStack[string]()

		// 测试空栈
		if !stack.IsEmpty() {
			t.Error("New stack should be empty")
		}

		if stack.Size() != 0 {
			t.Errorf("Empty stack size should be 0, got %d", stack.Size())
		}

		// 测试入栈
		items := []string{"first", "second", "third"}
		for _, item := range items {
			stack.Push(item)
		}

		if stack.Size() != len(items) {
			t.Errorf("Stack size should be %d, got %d", len(items), stack.Size())
		}

		// 测试查看栈顶
		if top, ok := stack.Peek(); !ok || top != "third" {
			t.Errorf("Peek should return (\"third\", true), got (%s, %t)", top, ok)
		}

		// 测试出栈
		for i := len(items) - 1; i >= 0; i-- {
			if item, ok := stack.Pop(); !ok || item != items[i] {
				t.Errorf("Pop should return (%s, true), got (%s, %t)", items[i], item, ok)
			}
		}

		// 测试空栈出栈
		if _, ok := stack.Pop(); ok {
			t.Error("Pop from empty stack should return (zero, false)")
		}

		t.Logf("Stack测试通过")
	})

	t.Run("Queue", func(t *testing.T) {
		queue := NewQueue[int]()

		// 测试空队列
		if !queue.IsEmpty() {
			t.Error("New queue should be empty")
		}

		// 测试入队
		items := []int{1, 2, 3, 4, 5}
		for _, item := range items {
			queue.Enqueue(item)
		}

		if queue.Size() != len(items) {
			t.Errorf("Queue size should be %d, got %d", len(items), queue.Size())
		}

		// 测试查看队首
		if front, ok := queue.Front(); !ok || front != 1 {
			t.Errorf("Front should return (1, true), got (%d, %t)", front, ok)
		}

		// 测试出队
		for i, expected := range items {
			if item, ok := queue.Dequeue(); !ok || item != expected {
				t.Errorf("Dequeue[%d] should return (%d, true), got (%d, %t)", i, expected, item, ok)
			}
		}

		t.Logf("Queue测试通过")
	})
}

func TestNumericConstraints(t *testing.T) {
	t.Run("SumFunction", func(t *testing.T) {
		// 测试整数
		intNumbers := []int{1, 2, 3, 4, 5}
		intSum := Sum(intNumbers)
		if intSum != 15 {
			t.Errorf("Sum(ints) = %d; want 15", intSum)
		}

		// 测试浮点数
		floatNumbers := []float64{1.1, 2.2, 3.3}
		floatSum := Sum(floatNumbers)
		expected := 6.6
		if floatSum < expected-0.01 || floatSum > expected+0.01 {
			t.Errorf("Sum(floats) = %f; want %f", floatSum, expected)
		}

		t.Logf("整数和: %d", intSum)
		t.Logf("浮点数和: %.2f", floatSum)
	})

	t.Run("AverageFunction", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8, 10}
		avg := Average(numbers)
		expected := 6.0

		if avg != expected {
			t.Errorf("Average = %f; want %f", avg, expected)
		}

		// 测试空切片
		emptyAvg := Average([]int{})
		if emptyAvg != 0 {
			t.Errorf("Average of empty slice should be 0, got %f", emptyAvg)
		}

		t.Logf("平均值: %.2f", avg)
	})
}

func TestSafeMap(t *testing.T) {
	t.Run("BasicOperations", func(t *testing.T) {
		safeMap := NewSafeMap[string, int]()

		// 测试设置和获取
		safeMap.Set("apple", 5)
		safeMap.Set("banana", 3)

		if value, exists := safeMap.Get("apple"); !exists || value != 5 {
			t.Errorf("Get(\"apple\") should return (5, true), got (%d, %t)", value, exists)
		}

		if _, exists := safeMap.Get("orange"); exists {
			t.Error("Get(\"orange\") should return (0, false)")
		}

		// 测试删除
		safeMap.Delete("banana")
		if _, exists := safeMap.Get("banana"); exists {
			t.Error("After delete, banana should not exist")
		}

		// 测试键和值
		keys := safeMap.Keys()
		values := safeMap.Values()

		if len(keys) != 1 || keys[0] != "apple" {
			t.Errorf("Keys should be [\"apple\"], got %v", keys)
		}

		if len(values) != 1 || values[0] != 5 {
			t.Errorf("Values should be [5], got %v", values)
		}

		t.Logf("SafeMap测试通过")
	})
}

func TestResultType(t *testing.T) {
	t.Run("SuccessResult", func(t *testing.T) {
		result := NewResult(42)

		if !result.IsOk() {
			t.Error("Success result should be Ok")
		}

		if result.IsErr() {
			t.Error("Success result should not be Err")
		}

		if value := result.Unwrap(); value != 42 {
			t.Errorf("Unwrap should return 42, got %d", value)
		}

		if value := result.UnwrapOr(0); value != 42 {
			t.Errorf("UnwrapOr should return 42, got %d", value)
		}

		t.Logf("成功结果测试通过")
	})

	t.Run("ErrorResult", func(t *testing.T) {
		err := errors.New("test error")
		result := NewError[int](err)

		if result.IsOk() {
			t.Error("Error result should not be Ok")
		}

		if !result.IsErr() {
			t.Error("Error result should be Err")
		}

		if value := result.UnwrapOr(100); value != 100 {
			t.Errorf("UnwrapOr should return 100, got %d", value)
		}

		t.Logf("错误结果测试通过")
	})
}

// 自定义类型用于测试Stringer约束
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

func TestStringerConstraint(t *testing.T) {
	t.Run("JoinFunction", func(t *testing.T) {
		people := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 35},
		}

		result := Join(people, ", ")
		expected := "Alice(25), Bob(30), Charlie(35)"

		if result != expected {
			t.Errorf("Join = %s; want %s", result, expected)
		}

		t.Logf("Join结果: %s", result)
	})
}

// 基准测试
func BenchmarkGenericMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Max(i, i+1)
	}
}

func BenchmarkGenericFilter(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(numbers, func(n int) bool { return n%2 == 0 })
	}
}

func BenchmarkGenericMap(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(numbers, func(n int) int { return n * n })
	}
}

func BenchmarkStack(b *testing.B) {
	stack := NewStack[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
		if i%2 == 0 {
			stack.Pop()
		}
	}
}

func BenchmarkQueue(b *testing.B) {
	queue := NewQueue[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
		if i%2 == 0 {
			queue.Dequeue()
		}
	}
}
