package testingexamples

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// 被测试的示例代码

// Calculator 计算器
type Calculator struct {
	history []Operation
}

// Operation 操作记录
type Operation struct {
	Type   string
	A, B   float64
	Result float64
	Time   time.Time
}

// NewCalculator 创建新计算器
func NewCalculator() *Calculator {
	return &Calculator{
		history: make([]Operation, 0),
	}
}

// Add 加法
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.addToHistory("ADD", a, b, result)
	return result
}

// Subtract 减法
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.addToHistory("SUBTRACT", a, b, result)
	return result
}

// Multiply 乘法
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.addToHistory("MULTIPLY", a, b, result)
	return result
}

// Divide 除法
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	result := a / b
	c.addToHistory("DIVIDE", a, b, result)
	return result, nil
}

// Sqrt 平方根
func (c *Calculator) Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("负数不能开平方根")
	}
	result := math.Sqrt(x)
	c.addToHistory("SQRT", x, 0, result)
	return result, nil
}

// GetHistory 获取历史记录
func (c *Calculator) GetHistory() []Operation {
	return c.history
}

// ClearHistory 清空历史记录
func (c *Calculator) ClearHistory() {
	c.history = c.history[:0]
}

func (c *Calculator) addToHistory(opType string, a, b, result float64) {
	op := Operation{
		Type:   opType,
		A:      a,
		B:      b,
		Result: result,
		Time:   time.Now(),
	}
	c.history = append(c.history, op)
}

// StringUtils 字符串工具
type StringUtils struct{}

// Reverse 反转字符串
func (su *StringUtils) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome 检查是否为回文
func (su *StringUtils) IsPalindrome(s string) bool {
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return cleaned == su.Reverse(cleaned)
}

// WordCount 统计单词数
func (su *StringUtils) WordCount(s string) int {
	if strings.TrimSpace(s) == "" {
		return 0
	}
	return len(strings.Fields(s))
}

// Capitalize 首字母大写
func (su *StringUtils) Capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// SortUtils 排序工具
type SortUtils struct{}

// BubbleSort 冒泡排序
func (su *SortUtils) BubbleSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// QuickSort 快速排序
func (su *SortUtils) QuickSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	su.quickSortHelper(result, 0, len(result)-1)
	return result
}

func (su *SortUtils) quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := su.partition(arr, low, high)
		su.quickSortHelper(arr, low, pi-1)
		su.quickSortHelper(arr, pi+1, high)
	}
}

func (su *SortUtils) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// IsSorted 检查是否已排序
func (su *SortUtils) IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// BinarySearch 二分查找
func (su *SortUtils) BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// UserService 用户服务（用于演示依赖注入测试）
type UserService struct {
	repo UserRepository
}

// UserRepository 用户仓库接口
type UserRepository interface {
	GetUser(id int) (*User, error)
	SaveUser(user *User) error
	DeleteUser(id int) error
}

// User 用户模型
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// NewUserService 创建用户服务
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser 获取用户
func (us *UserService) GetUser(id int) (*User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("无效的用户ID: %d", id)
	}
	return us.repo.GetUser(id)
}

// CreateUser 创建用户
func (us *UserService) CreateUser(name, email string, age int) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("用户名不能为空")
	}
	if email == "" {
		return nil, fmt.Errorf("邮箱不能为空")
	}
	if age < 0 || age > 150 {
		return nil, fmt.Errorf("年龄必须在0-150之间")
	}

	user := &User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	return user, us.repo.SaveUser(user)
}

// UpdateUser 更新用户
func (us *UserService) UpdateUser(user *User) error {
	if user == nil {
		return fmt.Errorf("用户不能为nil")
	}
	if user.ID <= 0 {
		return fmt.Errorf("无效的用户ID")
	}

	// 检查用户是否存在
	existing, err := us.repo.GetUser(user.ID)
	if err != nil {
		return fmt.Errorf("获取用户失败: %v", err)
	}
	if existing == nil {
		return fmt.Errorf("用户不存在: %d", user.ID)
	}

	return us.repo.SaveUser(user)
}

// DeleteUser 删除用户
func (us *UserService) DeleteUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("无效的用户ID: %d", id)
	}

	// 检查用户是否存在
	existing, err := us.repo.GetUser(id)
	if err != nil {
		return fmt.Errorf("获取用户失败: %v", err)
	}
	if existing == nil {
		return fmt.Errorf("用户不存在: %d", id)
	}

	return us.repo.DeleteUser(id)
}

// 数学工具函数

// Factorial 计算阶乘
func Factorial(n int) (int64, error) {
	if n < 0 {
		return 0, fmt.Errorf("负数没有阶乘")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}

	var result int64 = 1
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result, nil
}

// Fibonacci 计算斐波那契数列
func Fibonacci(n int) (int64, error) {
	if n < 0 {
		return 0, fmt.Errorf("n必须为非负数")
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}

	var a, b int64 = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

// IsPrime 判断是否为质数
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GCD 计算最大公约数
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM 计算最小公倍数
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// 集合操作

// Set 简单的集合实现
type Set struct {
	items map[interface{}]bool
}

// NewSet 创建新集合
func NewSet() *Set {
	return &Set{
		items: make(map[interface{}]bool),
	}
}

// Add 添加元素
func (s *Set) Add(item interface{}) {
	s.items[item] = true
}

// Remove 移除元素
func (s *Set) Remove(item interface{}) {
	delete(s.items, item)
}

// Contains 检查是否包含元素
func (s *Set) Contains(item interface{}) bool {
	return s.items[item]
}

// Size 获取集合大小
func (s *Set) Size() int {
	return len(s.items)
}

// ToSlice 转换为切片
func (s *Set) ToSlice() []interface{} {
	result := make([]interface{}, 0, len(s.items))
	for item := range s.items {
		result = append(result, item)
	}
	return result
}

// Union 并集
func (s *Set) Union(other *Set) *Set {
	result := NewSet()

	for item := range s.items {
		result.Add(item)
	}

	for item := range other.items {
		result.Add(item)
	}

	return result
}

// Intersection 交集
func (s *Set) Intersection(other *Set) *Set {
	result := NewSet()

	for item := range s.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// Difference 差集
func (s *Set) Difference(other *Set) *Set {
	result := NewSet()

	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// 缓存实现

// Cache 简单的内存缓存
type Cache struct {
	items map[string]CacheItem
}

// CacheItem 缓存项
type CacheItem struct {
	Value  interface{}
	Expiry time.Time
}

// NewCache 创建新缓存
func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

// Set 设置缓存项
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.items[key] = CacheItem{
		Value:  value,
		Expiry: time.Now().Add(duration),
	}
}

// Get 获取缓存项
func (c *Cache) Get(key string) (interface{}, bool) {
	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(item.Expiry) {
		delete(c.items, key)
		return nil, false
	}

	return item.Value, true
}

// Delete 删除缓存项
func (c *Cache) Delete(key string) {
	delete(c.items, key)
}

// Clear 清空缓存
func (c *Cache) Clear() {
	c.items = make(map[string]CacheItem)
}

// Size 获取缓存大小
func (c *Cache) Size() int {
	return len(c.items)
}

// CleanExpired 清理过期项
func (c *Cache) CleanExpired() int {
	now := time.Now()
	count := 0

	for key, item := range c.items {
		if now.After(item.Expiry) {
			delete(c.items, key)
			count++
		}
	}

	return count
}

// 示例函数
func TestingExamples() {
	fmt.Println("=== 测试框架示例 ===")

	// 计算器示例
	fmt.Println("\n🔹 计算器测试示例")
	calc := NewCalculator()

	fmt.Printf("10 + 5 = %.2f\n", calc.Add(10, 5))
	fmt.Printf("20 - 8 = %.2f\n", calc.Subtract(20, 8))
	fmt.Printf("6 * 7 = %.2f\n", calc.Multiply(6, 7))

	if result, err := calc.Divide(15, 3); err == nil {
		fmt.Printf("15 / 3 = %.2f\n", result)
	}

	if result, err := calc.Sqrt(16); err == nil {
		fmt.Printf("√16 = %.2f\n", result)
	}

	fmt.Printf("历史记录数量: %d\n", len(calc.GetHistory()))

	// 字符串工具示例
	fmt.Println("\n🔹 字符串工具测试示例")
	strUtils := &StringUtils{}

	text := "Hello World"
	fmt.Printf("原文: %s\n", text)
	fmt.Printf("反转: %s\n", strUtils.Reverse(text))
	fmt.Printf("单词数: %d\n", strUtils.WordCount(text))
	fmt.Printf("首字母大写: %s\n", strUtils.Capitalize("hello"))

	palindromes := []string{"level", "A man a plan a canal Panama", "hello"}
	for _, p := range palindromes {
		fmt.Printf("'%s' 是回文: %t\n", p, strUtils.IsPalindrome(p))
	}

	// 排序工具示例
	fmt.Println("\n🔹 排序工具测试示例")
	sortUtils := &SortUtils{}

	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("原数组: %v\n", arr)

	bubbleSorted := sortUtils.BubbleSort(arr)
	fmt.Printf("冒泡排序: %v\n", bubbleSorted)

	quickSorted := sortUtils.QuickSort(arr)
	fmt.Printf("快速排序: %v\n", quickSorted)

	fmt.Printf("是否已排序: %t\n", sortUtils.IsSorted(quickSorted))

	target := 25
	index := sortUtils.BinarySearch(quickSorted, target)
	fmt.Printf("二分查找 %d: 索引 %d\n", target, index)

	// 数学函数示例
	fmt.Println("\n🔹 数学函数测试示例")

	if fact, err := Factorial(5); err == nil {
		fmt.Printf("5! = %d\n", fact)
	}

	if fib, err := Fibonacci(10); err == nil {
		fmt.Printf("Fibonacci(10) = %d\n", fib)
	}

	primes := []int{2, 3, 4, 5, 17, 25}
	for _, n := range primes {
		fmt.Printf("%d 是质数: %t\n", n, IsPrime(n))
	}

	fmt.Printf("GCD(48, 18) = %d\n", GCD(48, 18))
	fmt.Printf("LCM(48, 18) = %d\n", LCM(48, 18))

	// 集合操作示例
	fmt.Println("\n🔹 集合操作测试示例")
	set1 := NewSet()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2 := NewSet()
	set2.Add(3)
	set2.Add(4)
	set2.Add(5)

	fmt.Printf("集合1: %v\n", set1.ToSlice())
	fmt.Printf("集合2: %v\n", set2.ToSlice())

	union := set1.Union(set2)
	fmt.Printf("并集: %v\n", union.ToSlice())

	intersection := set1.Intersection(set2)
	fmt.Printf("交集: %v\n", intersection.ToSlice())

	difference := set1.Difference(set2)
	fmt.Printf("差集: %v\n", difference.ToSlice())

	// 缓存示例
	fmt.Println("\n🔹 缓存测试示例")
	cache := NewCache()

	cache.Set("user:1", "张三", 5*time.Second)
	cache.Set("user:2", "李四", 10*time.Second)

	if value, found := cache.Get("user:1"); found {
		fmt.Printf("缓存命中: %v\n", value)
	}

	fmt.Printf("缓存大小: %d\n", cache.Size())

	// 等待过期
	time.Sleep(6 * time.Second)

	if _, found := cache.Get("user:1"); !found {
		fmt.Println("user:1 已过期")
	}

	expired := cache.CleanExpired()
	fmt.Printf("清理了 %d 个过期项\n", expired)
}
