package pointers

import (
	"fmt"
	"unsafe"
)

// 基本指针示例
func BasicPointers() {
	fmt.Println("=== 基本指针示例 ===")

	// 声明变量
	var num int = 42
	var str string = "Hello"

	// 获取变量的地址（指针）
	var numPtr *int = &num
	var strPtr *string = &str

	fmt.Printf("变量 num: %d, 地址: %p\n", num, &num)
	fmt.Printf("变量 str: %s, 地址: %p\n", str, &str)
	fmt.Printf("指针 numPtr: %p, 指向的值: %d\n", numPtr, *numPtr)
	fmt.Printf("指针 strPtr: %p, 指向的值: %s\n", strPtr, *strPtr)

	// 通过指针修改值
	*numPtr = 100
	*strPtr = "World"

	fmt.Printf("修改后 num: %d\n", num)
	fmt.Printf("修改后 str: %s\n", str)
}

// 零值指针示例
func NilPointers() {
	fmt.Println("\n=== 零值指针示例 ===")

	var ptr *int
	fmt.Printf("零值指针: %v\n", ptr)
	fmt.Printf("指针是否为 nil: %t\n", ptr == nil)

	// 安全检查指针
	if ptr != nil {
		fmt.Printf("指针指向的值: %d\n", *ptr)
	} else {
		fmt.Println("指针为 nil，不能解引用")
	}

	// 分配内存
	ptr = new(int)
	fmt.Printf("new 分配后的指针: %p\n", ptr)
	fmt.Printf("new 分配后的值: %d\n", *ptr)

	*ptr = 123
	fmt.Printf("赋值后的值: %d\n", *ptr)
}

// 指针作为函数参数
func PointerParameters() {
	fmt.Println("\n=== 指针作为函数参数 ===")

	x := 10
	y := 20

	fmt.Printf("交换前: x=%d, y=%d\n", x, y)

	// 值传递（不会交换）
	swapByValue(x, y)
	fmt.Printf("值传递后: x=%d, y=%d\n", x, y)

	// 指针传递（会交换）
	swapByPointer(&x, &y)
	fmt.Printf("指针传递后: x=%d, y=%d\n", x, y)

	// 修改值的示例
	num := 5
	fmt.Printf("修改前: %d\n", num)
	doubleValue(&num)
	fmt.Printf("修改后: %d\n", num)
}

// 值传递函数
func swapByValue(a, b int) {
	a, b = b, a
	fmt.Printf("函数内交换后: a=%d, b=%d\n", a, b)
}

// 指针传递函数
func swapByPointer(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("函数内交换后: *a=%d, *b=%d\n", *a, *b)
}

// 修改值的函数
func doubleValue(num *int) {
	*num *= 2
}

// 指针与数组
func PointersAndArrays() {
	fmt.Println("\n=== 指针与数组 ===")

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("原数组: %v\n", arr)

	// 数组指针
	arrPtr := &arr
	fmt.Printf("数组指针: %p\n", arrPtr)
	fmt.Printf("通过指针访问数组: %v\n", *arrPtr)

	// 修改数组元素
	(*arrPtr)[0] = 100
	fmt.Printf("修改后数组: %v\n", arr)

	// 元素指针
	firstPtr := &arr[0]
	lastPtr := &arr[4]

	fmt.Printf("第一个元素地址: %p, 值: %d\n", firstPtr, *firstPtr)
	fmt.Printf("最后一个元素地址: %p, 值: %d\n", lastPtr, *lastPtr)

	// 指针算术（在Go中不直接支持，但可以通过unsafe包实现）
	fmt.Printf("数组元素地址差: %d 字节\n", uintptr(unsafe.Pointer(lastPtr))-uintptr(unsafe.Pointer(firstPtr)))
}

// 指针与切片
func PointersAndSlices() {
	fmt.Println("\n=== 指针与切片 ===")

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("原切片: %v\n", slice)

	// 切片本身就包含指向底层数组的指针
	modifySlice(slice)
	fmt.Printf("修改后切片: %v\n", slice)

	// 切片指针
	slicePtr := &slice
	fmt.Printf("切片指针: %p\n", slicePtr)

	// 通过指针修改切片
	*slicePtr = append(*slicePtr, 6, 7)
	fmt.Printf("通过指针添加元素后: %v\n", slice)

	// 元素指针
	for i := range slice {
		elemPtr := &slice[i]
		fmt.Printf("元素 %d 地址: %p, 值: %d\n", i, elemPtr, *elemPtr)
	}
}

// 修改切片的函数
func modifySlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

// 指针与结构体
type Person struct {
	Name string
	Age  int
}

func PointersAndStructs() {
	fmt.Println("\n=== 指针与结构体 ===")

	person := Person{Name: "张三", Age: 25}
	fmt.Printf("原结构体: %+v\n", person)

	// 结构体指针
	personPtr := &person
	fmt.Printf("结构体指针: %p\n", personPtr)

	// 通过指针访问字段
	fmt.Printf("通过指针访问姓名: %s\n", (*personPtr).Name)
	fmt.Printf("通过指针访问年龄: %d\n", personPtr.Age) // Go的语法糖

	// 通过指针修改字段
	personPtr.Name = "李四"
	personPtr.Age = 30

	fmt.Printf("修改后结构体: %+v\n", person)

	// 字段指针
	namePtr := &person.Name
	agePtr := &person.Age

	fmt.Printf("姓名字段地址: %p, 值: %s\n", namePtr, *namePtr)
	fmt.Printf("年龄字段地址: %p, 值: %d\n", agePtr, *agePtr)

	// 通过字段指针修改
	*namePtr = "王五"
	*agePtr = 35

	fmt.Printf("通过字段指针修改后: %+v\n", person)
}

// 指针与方法
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

func (p *Person) GetInfo() string {
	return fmt.Sprintf("姓名: %s, 年龄: %d", p.Name, p.Age)
}

func (p Person) GetInfoByValue() string {
	return fmt.Sprintf("姓名: %s, 年龄: %d", p.Name, p.Age)
}

func PointersAndMethods() {
	fmt.Println("\n=== 指针与方法 ===")

	person := Person{Name: "赵六", Age: 28}
	fmt.Printf("初始状态: %s\n", person.GetInfo())

	// 指针接收者方法
	person.UpdateAge(32)
	fmt.Printf("更新年龄后: %s\n", person.GetInfo())

	// 通过指针调用方法
	personPtr := &person
	personPtr.UpdateAge(35)
	fmt.Printf("通过指针更新后: %s\n", personPtr.GetInfo())

	// 值接收者方法
	info := person.GetInfoByValue()
	fmt.Printf("值接收者方法: %s\n", info)
}

// 指针与内存分配
func MemoryAllocation() {
	fmt.Println("\n=== 指针与内存分配 ===")

	// 使用 new 分配内存
	intPtr := new(int)
	*intPtr = 42
	fmt.Printf("new 分配的整数: %d, 地址: %p\n", *intPtr, intPtr)

	stringPtr := new(string)
	*stringPtr = "Hello"
	fmt.Printf("new 分配的字符串: %s, 地址: %p\n", *stringPtr, stringPtr)

	// 使用取地址操作符
	value := 100
	valuePtr := &value
	fmt.Printf("取地址的值: %d, 地址: %p\n", *valuePtr, valuePtr)

	// 结构体指针
	personPtr := new(Person)
	personPtr.Name = "钱七"
	personPtr.Age = 40
	fmt.Printf("new 分配的结构体: %+v, 地址: %p\n", *personPtr, personPtr)

	// 使用字面量创建指针
	person2Ptr := &Person{Name: "孙八", Age: 45}
	fmt.Printf("字面量创建的结构体指针: %+v, 地址: %p\n", *person2Ptr, person2Ptr)
}

// 指针数组和数组指针
func PointerArrays() {
	fmt.Println("\n=== 指针数组和数组指针 ===")

	// 指针数组（数组的元素是指针）
	var ptrArray [3]*int
	a, b, c := 1, 2, 3
	ptrArray[0] = &a
	ptrArray[1] = &b
	ptrArray[2] = &c

	fmt.Println("指针数组:")
	for i, ptr := range ptrArray {
		fmt.Printf("  索引 %d: 地址 %p, 值 %d\n", i, ptr, *ptr)
	}

	// 修改原变量
	a = 10
	b = 20
	c = 30

	fmt.Println("修改原变量后的指针数组:")
	for i, ptr := range ptrArray {
		fmt.Printf("  索引 %d: 地址 %p, 值 %d\n", i, ptr, *ptr)
	}

	// 数组指针（指向数组的指针）
	arr := [3]int{100, 200, 300}
	arrPtr := &arr

	fmt.Printf("数组指针指向的数组: %v\n", *arrPtr)
	fmt.Printf("通过数组指针访问元素: %d\n", (*arrPtr)[1])

	// 修改数组
	(*arrPtr)[1] = 250
	fmt.Printf("修改后的数组: %v\n", arr)
}

// 多级指针
func MultiLevelPointers() {
	fmt.Println("\n=== 多级指针 ===")

	value := 42
	ptr1 := &value // 一级指针
	ptr2 := &ptr1  // 二级指针
	ptr3 := &ptr2  // 三级指针

	fmt.Printf("原值: %d\n", value)
	fmt.Printf("一级指针: %p, 指向的值: %d\n", ptr1, *ptr1)
	fmt.Printf("二级指针: %p, 指向的地址: %p, 最终值: %d\n", ptr2, *ptr2, **ptr2)
	fmt.Printf("三级指针: %p, 指向的地址: %p, 最终值: %d\n", ptr3, *ptr3, ***ptr3)

	// 通过多级指针修改值
	***ptr3 = 100
	fmt.Printf("通过三级指针修改后的值: %d\n", value)
}

// 指针的实际应用示例
func PracticalExamples() {
	fmt.Println("\n=== 指针的实际应用示例 ===")

	// 链表节点
	type ListNode struct {
		Value int
		Next  *ListNode
	}

	// 创建链表
	node1 := &ListNode{Value: 1}
	node2 := &ListNode{Value: 2}
	node3 := &ListNode{Value: 3}

	node1.Next = node2
	node2.Next = node3

	fmt.Println("链表遍历:")
	current := node1
	for current != nil {
		fmt.Printf("  节点值: %d, 地址: %p\n", current.Value, current)
		current = current.Next
	}

	// 创建二叉树
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}

	fmt.Println("二叉树前序遍历:")
	preorderTraversal(root)

	// 可选值模拟
	type OptionalInt struct {
		Value    int
		HasValue bool
	}

	opt1 := OptionalInt{Value: 42, HasValue: true}
	opt2 := OptionalInt{HasValue: false}

	fmt.Printf("可选值1: 有值=%t, 值=%d\n", opt1.HasValue, opt1.Value)
	fmt.Printf("可选值2: 有值=%t, 值=%d\n", opt2.HasValue, opt2.Value)
}

// TreeNode 二叉树节点
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历二叉树
func preorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("  访问节点: %d\n", node.Value)
	preorderTraversal(node.Left)
	preorderTraversal(node.Right)
}

// 指针的注意事项
func PointerCaveats() {
	fmt.Println("\n=== 指针的注意事项 ===")

	// 1. 悬空指针（在Go中由GC处理，但要注意逻辑）
	var ptr *int
	{
		value := 42
		ptr = &value
		fmt.Printf("作用域内: %d\n", *ptr)
	}
	// value 在这里已经超出作用域，但Go的GC会处理
	fmt.Printf("作用域外: %d\n", *ptr)

	// 2. 指针比较
	a := 10
	b := 10
	ptrA := &a
	ptrB := &b
	ptrA2 := &a

	fmt.Printf("ptrA == ptrB: %t (不同变量的地址)\n", ptrA == ptrB)
	fmt.Printf("ptrA == ptrA2: %t (同一变量的地址)\n", ptrA == ptrA2)

	// 3. 指针的零值检查
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("指针为 nil，避免解引用")
	}

	// 4. 切片和映射已经是引用类型
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // 共享底层数组
	slice2[0] = 100
	fmt.Printf("切片共享底层数组: slice1=%v, slice2=%v\n", slice1, slice2)
}
