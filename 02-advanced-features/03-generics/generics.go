package generics

import (
	"fmt"
	"sort"
	"strings"
)

// 排序约束类型
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// 基本泛型函数示例

// Max 返回两个可比较值中的最大值
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min 返回两个可比较值中的最小值
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Swap 交换两个值
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// 泛型切片操作

// Contains 检查切片是否包含指定元素
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Filter 过滤切片元素
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map 映射切片元素
func Map[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

// Reduce 归约切片元素
func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// Find 查找第一个满足条件的元素
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	var zero T
	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}
	return zero, false
}

// Reverse 反转切片
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// Unique 去重切片元素
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// 泛型数据结构

// Stack 泛型栈
type Stack[T any] struct {
	items []T
}

// NewStack 创建新栈
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push 入栈
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop 出栈
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Peek 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty 检查栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 获取栈大小
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Queue 泛型队列
type Queue[T any] struct {
	items []T
}

// NewQueue 创建新队列
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue 入队
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue 出队
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front 查看队首元素
func (q *Queue[T]) Front() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	return q.items[0], true
}

// IsEmpty 检查队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size 获取队列大小
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// 泛型约束示例

// Numeric 数值类型约束
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sum 计算数值切片的和
func Sum[T Numeric](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Average 计算数值切片的平均值
func Average[T Numeric](numbers []T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(Sum(numbers)) / float64(len(numbers))
}

// Stringer 字符串化接口约束
type Stringer interface {
	String() string
}

// Join 连接实现了String()方法的元素
func Join[T Stringer](items []T, separator string) string {
	if len(items) == 0 {
		return ""
	}

	var parts []string
	for _, item := range items {
		parts = append(parts, item.String())
	}
	return strings.Join(parts, separator)
}

// 泛型映射操作

// SafeMap 线程安全的泛型映射
type SafeMap[K comparable, V any] struct {
	data map[K]V
	// 在实际应用中这里应该有mutex，为了简化示例省略
}

// NewSafeMap 创建新的安全映射
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Set 设置键值对
func (sm *SafeMap[K, V]) Set(key K, value V) {
	sm.data[key] = value
}

// Get 获取值
func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	value, exists := sm.data[key]
	return value, exists
}

// Delete 删除键值对
func (sm *SafeMap[K, V]) Delete(key K) {
	delete(sm.data, key)
}

// Keys 获取所有键
func (sm *SafeMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(sm.data))
	for k := range sm.data {
		keys = append(keys, k)
	}
	return keys
}

// Values 获取所有值
func (sm *SafeMap[K, V]) Values() []V {
	values := make([]V, 0, len(sm.data))
	for _, v := range sm.data {
		values = append(values, v)
	}
	return values
}

// 实际应用示例

// Result 泛型结果类型
type Result[T any] struct {
	Value T
	Error error
}

// NewResult 创建成功结果
func NewResult[T any](value T) Result[T] {
	return Result[T]{Value: value, Error: nil}
}

// NewError 创建错误结果
func NewError[T any](err error) Result[T] {
	var zero T
	return Result[T]{Value: zero, Error: err}
}

// IsOk 检查是否成功
func (r Result[T]) IsOk() bool {
	return r.Error == nil
}

// IsErr 检查是否有错误
func (r Result[T]) IsErr() bool {
	return r.Error != nil
}

// Unwrap 解包结果
func (r Result[T]) Unwrap() T {
	if r.Error != nil {
		panic(r.Error)
	}
	return r.Value
}

// UnwrapOr 解包结果或返回默认值
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.Error != nil {
		return defaultValue
	}
	return r.Value
}

// 示例函数

// GenericExamples 泛型示例
func GenericExamples() {
	fmt.Println("=== 泛型示例 ===")

	// 基本泛型函数
	fmt.Println("\n🔹 基本泛型函数")
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\"): %s\n", Max("apple", "banana"))

	a, b := 100, 200
	fmt.Printf("交换前: a=%d, b=%d\n", a, b)
	Swap(&a, &b)
	fmt.Printf("交换后: a=%d, b=%d\n", a, b)

	// 泛型切片操作
	fmt.Println("\n🔹 泛型切片操作")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("原始数组: %v\n", numbers)

	// 过滤偶数
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("偶数: %v\n", evens)

	// 映射为平方
	squares := Map(numbers, func(n int) int { return n * n })
	fmt.Printf("平方: %v\n", squares)

	// 求和
	sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	fmt.Printf("求和: %d\n", sum)

	// 查找
	if found, ok := Find(numbers, func(n int) bool { return n > 5 }); ok {
		fmt.Printf("第一个大于5的数: %d\n", found)
	}

	// 去重
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	unique := Unique(duplicates)
	fmt.Printf("去重前: %v\n", duplicates)
	fmt.Printf("去重后: %v\n", unique)

	// 泛型数据结构
	fmt.Println("\n🔹 泛型数据结构")

	// 栈示例
	stack := NewStack[string]()
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	fmt.Printf("栈大小: %d\n", stack.Size())
	if top, ok := stack.Peek(); ok {
		fmt.Printf("栈顶元素: %s\n", top)
	}

	for !stack.IsEmpty() {
		if item, ok := stack.Pop(); ok {
			fmt.Printf("出栈: %s\n", item)
		}
	}

	// 队列示例
	queue := NewQueue[int]()
	for i := 1; i <= 5; i++ {
		queue.Enqueue(i)
	}

	fmt.Printf("队列大小: %d\n", queue.Size())
	for !queue.IsEmpty() {
		if item, ok := queue.Dequeue(); ok {
			fmt.Printf("出队: %d\n", item)
		}
	}

	// 数值约束示例
	fmt.Println("\n🔹 数值约束示例")
	intNumbers := []int{1, 2, 3, 4, 5}
	floatNumbers := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Printf("整数和: %d\n", Sum(intNumbers))
	fmt.Printf("浮点数和: %.2f\n", Sum(floatNumbers))
	fmt.Printf("整数平均值: %.2f\n", Average(intNumbers))
	fmt.Printf("浮点数平均值: %.2f\n", Average(floatNumbers))

	// 泛型映射示例
	fmt.Println("\n🔹 泛型映射示例")
	safeMap := NewSafeMap[string, int]()
	safeMap.Set("apple", 5)
	safeMap.Set("banana", 3)
	safeMap.Set("orange", 8)

	if value, exists := safeMap.Get("apple"); exists {
		fmt.Printf("apple的数量: %d\n", value)
	}

	fmt.Printf("所有键: %v\n", safeMap.Keys())
	fmt.Printf("所有值: %v\n", safeMap.Values())

	// Result类型示例
	fmt.Println("\n🔹 Result类型示例")
	successResult := NewResult(42)
	errorResult := NewError[int](fmt.Errorf("something went wrong"))

	fmt.Printf("成功结果: 值=%d, 是否成功=%t\n", successResult.Value, successResult.IsOk())
	fmt.Printf("错误结果: 错误=%v, 是否有错=%t\n", errorResult.Error, errorResult.IsErr())

	fmt.Printf("成功结果解包: %d\n", successResult.Unwrap())
	fmt.Printf("错误结果解包或默认值: %d\n", errorResult.UnwrapOr(0))
}

// 泛型排序函数
func SortSlice[T Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// 排序示例
func SortExamples() {
	fmt.Println("\n=== 泛型排序示例 ===")

	// 整数排序
	ints := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("排序前: %v\n", ints)
	SortSlice(ints)
	fmt.Printf("排序后: %v\n", ints)

	// 字符串排序
	strings := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("排序前: %v\n", strings)
	SortSlice(strings)
	fmt.Printf("排序后: %v\n", strings)

	// 浮点数排序
	floats := []float64{3.14, 2.71, 1.41, 1.73}
	fmt.Printf("排序前: %v\n", floats)
	SortSlice(floats)
	fmt.Printf("排序后: %v\n", floats)
}
