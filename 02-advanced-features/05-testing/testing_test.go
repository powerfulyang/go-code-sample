package testingexamples

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 基本单元测试示例

func TestCalculator(t *testing.T) {
	calc := NewCalculator()

	t.Run("Addition", func(t *testing.T) {
		result := calc.Add(10, 5)
		expected := 15.0
		if result != expected {
			t.Errorf("Add(10, 5) = %f; want %f", result, expected)
		}
	})

	t.Run("Subtraction", func(t *testing.T) {
		result := calc.Subtract(10, 5)
		expected := 5.0
		if result != expected {
			t.Errorf("Subtract(10, 5) = %f; want %f", result, expected)
		}
	})

	t.Run("Multiplication", func(t *testing.T) {
		result := calc.Multiply(10, 5)
		expected := 50.0
		if result != expected {
			t.Errorf("Multiply(10, 5) = %f; want %f", result, expected)
		}
	})

	t.Run("Division", func(t *testing.T) {
		result, err := calc.Divide(10, 5)
		if err != nil {
			t.Errorf("Divide(10, 5) returned error: %v", err)
		}
		expected := 2.0
		if result != expected {
			t.Errorf("Divide(10, 5) = %f; want %f", result, expected)
		}
	})

	t.Run("DivisionByZero", func(t *testing.T) {
		_, err := calc.Divide(10, 0)
		if err == nil {
			t.Error("Divide(10, 0) should return error")
		}
	})

	t.Run("SquareRoot", func(t *testing.T) {
		result, err := calc.Sqrt(16)
		if err != nil {
			t.Errorf("Sqrt(16) returned error: %v", err)
		}
		expected := 4.0
		if result != expected {
			t.Errorf("Sqrt(16) = %f; want %f", result, expected)
		}
	})

	t.Run("SquareRootNegative", func(t *testing.T) {
		_, err := calc.Sqrt(-1)
		if err == nil {
			t.Error("Sqrt(-1) should return error")
		}
	})

	t.Run("History", func(t *testing.T) {
		calc.ClearHistory()
		calc.Add(1, 2)
		calc.Multiply(3, 4)

		history := calc.GetHistory()
		if len(history) != 2 {
			t.Errorf("History length should be 2, got %d", len(history))
		}

		if history[0].Type != "ADD" {
			t.Errorf("First operation should be ADD, got %s", history[0].Type)
		}

		if history[1].Type != "MULTIPLY" {
			t.Errorf("Second operation should be MULTIPLY, got %s", history[1].Type)
		}
	})
}

// 表格驱动测试示例

func TestStringUtils(t *testing.T) {
	strUtils := &StringUtils{}

	t.Run("Reverse", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"Empty string", "", ""},
			{"Single character", "a", "a"},
			{"Simple word", "hello", "olleh"},
			{"With spaces", "hello world", "dlrow olleh"},
			{"Unicode", "你好", "好你"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := strUtils.Reverse(tt.input)
				if result != tt.expected {
					t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("IsPalindrome", func(t *testing.T) {
		tests := []struct {
			input    string
			expected bool
		}{
			{"", true},
			{"a", true},
			{"aa", true},
			{"aba", true},
			{"level", true},
			{"A man a plan a canal Panama", true},
			{"hello", false},
			{"world", false},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("IsPalindrome(%q)", tt.input), func(t *testing.T) {
				result := strUtils.IsPalindrome(tt.input)
				if result != tt.expected {
					t.Errorf("IsPalindrome(%q) = %t; want %t", tt.input, result, tt.expected)
				}
			})
		}
	})

	t.Run("WordCount", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int
		}{
			{"", 0},
			{"   ", 0},
			{"hello", 1},
			{"hello world", 2},
			{"  hello   world  ", 2},
			{"one two three four five", 5},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("WordCount(%q)", tt.input), func(t *testing.T) {
				result := strUtils.WordCount(tt.input)
				if result != tt.expected {
					t.Errorf("WordCount(%q) = %d; want %d", tt.input, result, tt.expected)
				}
			})
		}
	})
}

// 错误处理测试示例

func TestMathFunctions(t *testing.T) {
	t.Run("Factorial", func(t *testing.T) {
		tests := []struct {
			input       int
			expected    int64
			expectError bool
		}{
			{0, 1, false},
			{1, 1, false},
			{5, 120, false},
			{10, 3628800, false},
			{-1, 0, true},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("Factorial(%d)", tt.input), func(t *testing.T) {
				result, err := Factorial(tt.input)

				if tt.expectError {
					if err == nil {
						t.Errorf("Factorial(%d) should return error", tt.input)
					}
				} else {
					if err != nil {
						t.Errorf("Factorial(%d) returned unexpected error: %v", tt.input, err)
					}
					if result != tt.expected {
						t.Errorf("Factorial(%d) = %d; want %d", tt.input, result, tt.expected)
					}
				}
			})
		}
	})

	t.Run("IsPrime", func(t *testing.T) {
		tests := []struct {
			input    int
			expected bool
		}{
			{-1, false},
			{0, false},
			{1, false},
			{2, true},
			{3, true},
			{4, false},
			{17, true},
			{25, false},
			{97, true},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("IsPrime(%d)", tt.input), func(t *testing.T) {
				result := IsPrime(tt.input)
				if result != tt.expected {
					t.Errorf("IsPrime(%d) = %t; want %t", tt.input, result, tt.expected)
				}
			})
		}
	})
}

// Mock测试示例

// MockUserRepository 模拟用户仓库
type MockUserRepository struct {
	users  map[int]*User
	nextID int
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
	if id == 999 {
		return nil, errors.New("database error")
	}
	user, exists := m.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

func (m *MockUserRepository) SaveUser(user *User) error {
	if user.Name == "error" {
		return errors.New("save error")
	}
	if user.ID == 0 {
		user.ID = m.nextID
		m.nextID++
	}
	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) DeleteUser(id int) error {
	if id == 999 {
		return errors.New("delete error")
	}
	delete(m.users, id)
	return nil
}

func TestUserService(t *testing.T) {
	mockRepo := NewMockUserRepository()
	service := NewUserService(mockRepo)

	t.Run("GetUser", func(t *testing.T) {
		// 准备测试数据
		testUser := &User{ID: 1, Name: "张三", Email: "zhangsan@example.com", Age: 25}
		mockRepo.SaveUser(testUser)

		// 测试获取存在的用户
		user, err := service.GetUser(1)
		if err != nil {
			t.Errorf("GetUser(1) returned error: %v", err)
		}
		if user == nil {
			t.Error("GetUser(1) returned nil user")
		}
		if user.Name != "张三" {
			t.Errorf("User name = %s; want 张三", user.Name)
		}

		// 测试获取不存在的用户
		user, err = service.GetUser(2)
		if err != nil {
			t.Errorf("GetUser(2) returned error: %v", err)
		}
		if user != nil {
			t.Error("GetUser(2) should return nil for non-existent user")
		}

		// 测试无效ID
		_, err = service.GetUser(0)
		if err == nil {
			t.Error("GetUser(0) should return error for invalid ID")
		}

		// 测试数据库错误
		_, err = service.GetUser(999)
		if err == nil {
			t.Error("GetUser(999) should return database error")
		}
	})

	t.Run("CreateUser", func(t *testing.T) {
		// 测试正常创建
		user, err := service.CreateUser("李四", "lisi@example.com", 30)
		if err != nil {
			t.Errorf("CreateUser returned error: %v", err)
		}
		if user == nil {
			t.Error("CreateUser returned nil user")
		}
		if user.ID == 0 {
			t.Error("Created user should have non-zero ID")
		}

		// 测试验证错误
		testCases := []struct {
			name  string
			email string
			age   int
		}{
			{"", "valid@example.com", 25},       // 空名称
			{"Valid", "", 25},                   // 空邮箱
			{"Valid", "valid@example.com", -1},  // 无效年龄
			{"Valid", "valid@example.com", 200}, // 无效年龄
		}

		for _, tc := range testCases {
			_, err := service.CreateUser(tc.name, tc.email, tc.age)
			if err == nil {
				t.Errorf("CreateUser(%s, %s, %d) should return validation error", tc.name, tc.email, tc.age)
			}
		}
	})
}

// 基准测试示例

func BenchmarkSortUtils(b *testing.B) {
	sortUtils := &SortUtils{}

	// 准备测试数据
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = size - i // 逆序数组
		}

		b.Run(fmt.Sprintf("BubbleSort-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sortUtils.BubbleSort(arr)
			}
		})

		b.Run(fmt.Sprintf("QuickSort-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sortUtils.QuickSort(arr)
			}
		})
	}
}

func BenchmarkStringUtils(b *testing.B) {
	strUtils := &StringUtils{}

	strings := []string{
		"hello",
		"hello world",
		"The quick brown fox jumps over the lazy dog",
		"A man a plan a canal Panama",
	}

	for _, s := range strings {
		b.Run(fmt.Sprintf("Reverse-%d", len(s)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				strUtils.Reverse(s)
			}
		})

		b.Run(fmt.Sprintf("IsPalindrome-%d", len(s)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				strUtils.IsPalindrome(s)
			}
		})
	}
}

func BenchmarkMathFunctions(b *testing.B) {
	b.Run("Factorial", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Factorial(10)
		}
	})

	b.Run("Fibonacci", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Fibonacci(20)
		}
	})

	b.Run("IsPrime", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsPrime(97)
		}
	})
}

// 示例测试

func ExampleCalculator_Add() {
	calc := NewCalculator()
	result := calc.Add(2, 3)
	fmt.Printf("%.0f", result)
	// Output: 5
}

func ExampleStringUtils_Reverse() {
	strUtils := &StringUtils{}
	result := strUtils.Reverse("hello")
	fmt.Println(result)
	// Output: olleh
}

func ExampleStringUtils_IsPalindrome() {
	strUtils := &StringUtils{}
	result := strUtils.IsPalindrome("level")
	fmt.Println(result)
	// Output: true
}

func ExampleFactorial() {
	result, _ := Factorial(5)
	fmt.Printf("%d", result)
	// Output: 120
}

// 模糊测试示例 (Go 1.18+)

func FuzzStringReverse(f *testing.F) {
	strUtils := &StringUtils{}

	// 添加种子语料
	testcases := []string{"", "a", "hello", "世界"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, s string) {
		// 反转两次应该得到原字符串
		reversed := strUtils.Reverse(s)
		doubleReversed := strUtils.Reverse(reversed)

		if s != doubleReversed {
			t.Errorf("Double reverse failed: original=%q, double_reversed=%q", s, doubleReversed)
		}

		// 反转后长度应该相同
		if len(s) != len(reversed) {
			t.Errorf("Length mismatch: original=%d, reversed=%d", len(s), len(reversed))
		}
	})
}

// 辅助函数测试

func TestSet(t *testing.T) {
	t.Run("BasicOperations", func(t *testing.T) {
		set := NewSet()

		// 测试添加
		set.Add(1)
		set.Add(2)
		set.Add(3)

		if set.Size() != 3 {
			t.Errorf("Set size should be 3, got %d", set.Size())
		}

		// 测试包含
		if !set.Contains(2) {
			t.Error("Set should contain 2")
		}

		if set.Contains(4) {
			t.Error("Set should not contain 4")
		}

		// 测试移除
		set.Remove(2)
		if set.Contains(2) {
			t.Error("Set should not contain 2 after removal")
		}

		if set.Size() != 2 {
			t.Errorf("Set size should be 2 after removal, got %d", set.Size())
		}
	})

	t.Run("SetOperations", func(t *testing.T) {
		set1 := NewSet()
		set1.Add(1)
		set1.Add(2)
		set1.Add(3)

		set2 := NewSet()
		set2.Add(3)
		set2.Add(4)
		set2.Add(5)

		// 测试并集
		union := set1.Union(set2)
		if union.Size() != 5 {
			t.Errorf("Union size should be 5, got %d", union.Size())
		}

		// 测试交集
		intersection := set1.Intersection(set2)
		if intersection.Size() != 1 {
			t.Errorf("Intersection size should be 1, got %d", intersection.Size())
		}
		if !intersection.Contains(3) {
			t.Error("Intersection should contain 3")
		}

		// 测试差集
		difference := set1.Difference(set2)
		if difference.Size() != 2 {
			t.Errorf("Difference size should be 2, got %d", difference.Size())
		}
		if !difference.Contains(1) || !difference.Contains(2) {
			t.Error("Difference should contain 1 and 2")
		}
	})
}

func TestCache(t *testing.T) {
	t.Run("BasicOperations", func(t *testing.T) {
		cache := NewCache()

		// 测试设置和获取
		cache.Set("key1", "value1", time.Minute)

		if value, found := cache.Get("key1"); !found || value != "value1" {
			t.Errorf("Cache should return 'value1', got %v, found=%t", value, found)
		}

		// 测试不存在的键
		if _, found := cache.Get("nonexistent"); found {
			t.Error("Cache should not find non-existent key")
		}

		// 测试删除
		cache.Delete("key1")
		if _, found := cache.Get("key1"); found {
			t.Error("Cache should not find deleted key")
		}
	})

	t.Run("Expiration", func(t *testing.T) {
		cache := NewCache()

		// 设置短期过期的项
		cache.Set("temp", "value", 100*time.Millisecond)

		// 立即获取应该成功
		if _, found := cache.Get("temp"); !found {
			t.Error("Cache should find recently set key")
		}

		// 等待过期
		time.Sleep(150 * time.Millisecond)

		// 过期后获取应该失败
		if _, found := cache.Get("temp"); found {
			t.Error("Cache should not find expired key")
		}
	})
}

// 测试辅助函数

func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func assertError(t *testing.T, err error, expectError bool) {
	t.Helper()
	if expectError && err == nil {
		t.Error("Expected error but got nil")
	}
	if !expectError && err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
}

func assertContains(t *testing.T, slice []interface{}, item interface{}) {
	t.Helper()
	for _, v := range slice {
		if reflect.DeepEqual(v, item) {
			return
		}
	}
	t.Errorf("Slice %v should contain %v", slice, item)
}
