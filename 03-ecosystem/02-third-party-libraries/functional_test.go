package thirdparty

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("IntToString", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		strings := Map(numbers, func(n int) string {
			return fmt.Sprintf("num_%d", n)
		})

		expected := []string{"num_1", "num_2", "num_3", "num_4", "num_5"}
		if !reflect.DeepEqual(strings, expected) {
			t.Errorf("Map failed: got %v, want %v", strings, expected)
		}

		t.Log("Map整数到字符串测试通过")
	})

	t.Run("Square", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		squares := Map(numbers, func(n int) int { return n * n })

		expected := []int{1, 4, 9, 16}
		if !reflect.DeepEqual(squares, expected) {
			t.Errorf("Map squares failed: got %v, want %v", squares, expected)
		}

		t.Log("Map平方计算测试通过")
	})
}

func TestFilter(t *testing.T) {
	t.Run("FilterEvens", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evens := Filter(numbers, func(n int) bool { return n%2 == 0 })

		expected := []int{2, 4, 6, 8, 10}
		if !reflect.DeepEqual(evens, expected) {
			t.Errorf("Filter evens failed: got %v, want %v", evens, expected)
		}

		t.Log("Filter偶数测试通过")
	})

	t.Run("FilterStrings", func(t *testing.T) {
		words := []string{"go", "rust", "python", "java", "c++"}
		shortWords := Filter(words, func(s string) bool { return len(s) <= 3 })

		expected := []string{"go", "c++"}
		if !reflect.DeepEqual(shortWords, expected) {
			t.Errorf("Filter short words failed: got %v, want %v", shortWords, expected)
		}

		t.Log("Filter短单词测试通过")
	})
}

func TestReduce(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })

		expected := 15
		if sum != expected {
			t.Errorf("Reduce sum failed: got %d, want %d", sum, expected)
		}

		t.Log("Reduce求和测试通过")
	})

	t.Run("Product", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		product := Reduce(numbers, 1, func(acc, n int) int { return acc * n })

		expected := 24
		if product != expected {
			t.Errorf("Reduce product failed: got %d, want %d", product, expected)
		}

		t.Log("Reduce乘积测试通过")
	})

	t.Run("StringConcat", func(t *testing.T) {
		words := []string{"Hello", "World", "Go"}
		sentence := Reduce(words, "", func(acc, word string) string {
			if acc == "" {
				return word
			}
			return acc + " " + word
		})

		expected := "Hello World Go"
		if sentence != expected {
			t.Errorf("Reduce concat failed: got %q, want %q", sentence, expected)
		}

		t.Log("Reduce字符串连接测试通过")
	})
}

func TestFind(t *testing.T) {
	t.Run("FindExisting", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		found, ok := Find(numbers, func(n int) bool { return n > 3 })

		if !ok {
			t.Error("Find should return true for existing element")
		}
		if found != 4 {
			t.Errorf("Find failed: got %d, want 4", found)
		}

		t.Log("Find存在元素测试通过")
	})

	t.Run("FindNonExisting", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		_, ok := Find(numbers, func(n int) bool { return n > 10 })

		if ok {
			t.Error("Find should return false for non-existing element")
		}

		t.Log("Find不存在元素测试通过")
	})
}

func TestUnique(t *testing.T) {
	t.Run("UniqueInts", func(t *testing.T) {
		numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
		unique := Unique(numbers)

		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(unique, expected) {
			t.Errorf("Unique failed: got %v, want %v", unique, expected)
		}

		t.Log("Unique整数测试通过")
	})

	t.Run("UniqueStrings", func(t *testing.T) {
		words := []string{"apple", "banana", "apple", "cherry", "banana"}
		unique := Unique(words)

		expected := []string{"apple", "banana", "cherry"}
		if !reflect.DeepEqual(unique, expected) {
			t.Errorf("Unique strings failed: got %v, want %v", unique, expected)
		}

		t.Log("Unique字符串测试通过")
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("GroupByLength", func(t *testing.T) {
		words := []string{"go", "rust", "python", "java", "c", "javascript"}
		grouped := GroupBy(words, func(s string) int { return len(s) })

		// 检查分组结果
		if len(grouped[1]) != 1 || grouped[1][0] != "c" {
			t.Errorf("Group by length failed for length 1")
		}
		if len(grouped[2]) != 1 || grouped[2][0] != "go" {
			t.Errorf("Group by length failed for length 2")
		}
		if len(grouped[4]) != 2 {
			t.Errorf("Group by length failed for length 4")
		}

		t.Log("GroupBy长度测试通过")
	})
}

func TestPartition(t *testing.T) {
	t.Run("PartitionEvensOdds", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evens, odds := Partition(numbers, func(n int) bool { return n%2 == 0 })

		expectedEvens := []int{2, 4, 6, 8, 10}
		expectedOdds := []int{1, 3, 5, 7, 9}

		if !reflect.DeepEqual(evens, expectedEvens) {
			t.Errorf("Partition evens failed: got %v, want %v", evens, expectedEvens)
		}
		if !reflect.DeepEqual(odds, expectedOdds) {
			t.Errorf("Partition odds failed: got %v, want %v", odds, expectedOdds)
		}

		t.Log("Partition奇偶数测试通过")
	})
}

func TestChunk(t *testing.T) {
	t.Run("ChunkNormal", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		chunks := Chunk(numbers, 3)

		expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		if !reflect.DeepEqual(chunks, expected) {
			t.Errorf("Chunk failed: got %v, want %v", chunks, expected)
		}

		t.Log("Chunk正常分块测试通过")
	})

	t.Run("ChunkIncomplete", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		chunks := Chunk(numbers, 3)

		expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
		if !reflect.DeepEqual(chunks, expected) {
			t.Errorf("Chunk incomplete failed: got %v, want %v", chunks, expected)
		}

		t.Log("Chunk不完整分块测试通过")
	})
}

func TestSortBy(t *testing.T) {
	t.Run("SortByLength", func(t *testing.T) {
		words := []string{"python", "go", "rust", "javascript", "c"}
		sorted := SortBy(words, func(s string) int { return len(s) })

		expected := []string{"c", "go", "rust", "python", "javascript"}
		if !reflect.DeepEqual(sorted, expected) {
			t.Errorf("SortBy length failed: got %v, want %v", sorted, expected)
		}

		t.Log("SortBy长度测试通过")
	})
}

func TestMinMaxBy(t *testing.T) {
	users := []User{
		{ID: 1, Name: "张三", Age: 25, Salary: 8000},
		{ID: 2, Name: "李四", Age: 30, Salary: 12000},
		{ID: 3, Name: "王五", Age: 28, Salary: 9000},
	}

	t.Run("MinByAge", func(t *testing.T) {
		youngest, found := MinBy(users, func(u User) int { return u.Age })

		if !found {
			t.Error("MinBy should find youngest user")
		}
		if youngest.Name != "张三" {
			t.Errorf("MinBy age failed: got %s, want 张三", youngest.Name)
		}

		t.Log("MinBy年龄测试通过")
	})

	t.Run("MaxBySalary", func(t *testing.T) {
		richest, found := MaxBy(users, func(u User) int { return u.Salary })

		if !found {
			t.Error("MaxBy should find richest user")
		}
		if richest.Name != "李四" {
			t.Errorf("MaxBy salary failed: got %s, want 李四", richest.Name)
		}

		t.Log("MaxBy薪资测试通过")
	})
}

func TestSum(t *testing.T) {
	t.Run("SumInts", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)

		expected := 15
		if sum != expected {
			t.Errorf("Sum failed: got %d, want %d", sum, expected)
		}

		t.Log("Sum整数测试通过")
	})

	t.Run("SumFloats", func(t *testing.T) {
		numbers := []float64{1.1, 2.2, 3.3}
		sum := Sum(numbers)

		expected := 6.6
		if sum < expected-0.01 || sum > expected+0.01 {
			t.Errorf("Sum floats failed: got %f, want %f", sum, expected)
		}

		t.Log("Sum浮点数测试通过")
	})
}

func TestEveryAndSome(t *testing.T) {
	numbers := []int{2, 4, 6, 8, 10}

	t.Run("EveryEven", func(t *testing.T) {
		allEven := Every(numbers, func(n int) bool { return n%2 == 0 })

		if !allEven {
			t.Error("Every should return true for all even numbers")
		}

		t.Log("Every偶数测试通过")
	})

	t.Run("SomeGreaterThan5", func(t *testing.T) {
		hasLarge := Some(numbers, func(n int) bool { return n > 5 })

		if !hasLarge {
			t.Error("Some should return true for numbers greater than 5")
		}

		t.Log("Some大于5测试通过")
	})
}

func TestTakeAndDrop(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("Take", func(t *testing.T) {
		first5 := Take(numbers, 5)
		expected := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(first5, expected) {
			t.Errorf("Take failed: got %v, want %v", first5, expected)
		}

		t.Log("Take前5个测试通过")
	})

	t.Run("Drop", func(t *testing.T) {
		dropped := Drop(numbers, 3)
		expected := []int{4, 5, 6, 7, 8, 9, 10}

		if !reflect.DeepEqual(dropped, expected) {
			t.Errorf("Drop failed: got %v, want %v", dropped, expected)
		}

		t.Log("Drop前3个测试通过")
	})

	t.Run("TakeRight", func(t *testing.T) {
		last3 := TakeRight(numbers, 3)
		expected := []int{8, 9, 10}

		if !reflect.DeepEqual(last3, expected) {
			t.Errorf("TakeRight failed: got %v, want %v", last3, expected)
		}

		t.Log("TakeRight后3个测试通过")
	})

	t.Run("DropRight", func(t *testing.T) {
		droppedRight := DropRight(numbers, 2)
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

		if !reflect.DeepEqual(droppedRight, expected) {
			t.Errorf("DropRight failed: got %v, want %v", droppedRight, expected)
		}

		t.Log("DropRight后2个测试通过")
	})
}

func TestUserDataProcessing(t *testing.T) {
	users := []User{
		{ID: 1, Name: "张三", Age: 25, City: "北京", IsActive: true, Salary: 8000},
		{ID: 2, Name: "李四", Age: 30, City: "上海", IsActive: true, Salary: 12000},
		{ID: 3, Name: "王五", Age: 28, City: "北京", IsActive: false, Salary: 9000},
		{ID: 4, Name: "赵六", Age: 35, City: "深圳", IsActive: true, Salary: 15000},
	}

	t.Run("FilterActiveUsers", func(t *testing.T) {
		activeUsers := Filter(users, func(u User) bool { return u.IsActive })

		if len(activeUsers) != 3 {
			t.Errorf("Active users count: got %d, want 3", len(activeUsers))
		}

		t.Log("过滤活跃用户测试通过")
	})

	t.Run("GroupByCity", func(t *testing.T) {
		usersByCity := GroupBy(users, func(u User) string { return u.City })

		if len(usersByCity["北京"]) != 2 {
			t.Errorf("Beijing users count: got %d, want 2", len(usersByCity["北京"]))
		}
		if len(usersByCity["上海"]) != 1 {
			t.Errorf("Shanghai users count: got %d, want 1", len(usersByCity["上海"]))
		}

		t.Log("按城市分组测试通过")
	})

	t.Run("CalculateAverageSalary", func(t *testing.T) {
		activeUsers := Filter(users, func(u User) bool { return u.IsActive })
		totalSalary := SumBy(activeUsers, func(u User) int { return u.Salary })
		avgSalary := float64(totalSalary) / float64(len(activeUsers))

		expectedAvg := float64(8000+12000+15000) / 3.0
		if avgSalary != expectedAvg {
			t.Errorf("Average salary: got %.2f, want %.2f", avgSalary, expectedAvg)
		}

		t.Log("计算平均薪资测试通过")
	})
}

// 基准测试
func BenchmarkMap(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(numbers, func(n int) int { return n * n })
	}
}

func BenchmarkFilter(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(numbers, func(n int) bool { return n%2 == 0 })
	}
}

func BenchmarkReduce(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	}
}
