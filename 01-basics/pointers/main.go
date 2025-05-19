package main

import "fmt"

func main() {
	fmt.Println("=== Go 指针示例 ===")

	// 基本指针操作
	fmt.Println("\n--- 基本指针操作 ---")

	// 声明变量
	var x int = 42
	fmt.Printf("变量x的值: %d\n", x)
	fmt.Printf("变量x的地址: %p\n", &x)

	// 声明指针
	var p *int
	fmt.Printf("未初始化的指针p: %v\n", p) // nil

	// 指针赋值
	p = &x
	fmt.Printf("指针p指向x的地址: %p\n", p)
	fmt.Printf("通过指针p访问x的值: %d\n", *p)

	// 通过指针修改值
	*p = 100
	fmt.Printf("通过指针修改后，x的值: %d\n", x)

	// 指针的指针
	fmt.Println("\n--- 指针的指针 ---")
	var pp **int = &p
	fmt.Printf("指针的指针pp: %p\n", pp)
	fmt.Printf("pp指向的指针的地址: %p\n", *pp)
	fmt.Printf("通过pp访问x的值: %d\n", **pp)

	// 修改通过指针的指针
	**pp = 200
	fmt.Printf("通过指针的指针修改后，x的值: %d\n", x)

	// 不同类型的指针
	fmt.Println("\n--- 不同类型的指针 ---")

	var str string = "Hello, Go!"
	var strPtr *string = &str
	fmt.Printf("字符串: %s\n", str)
	fmt.Printf("字符串指针: %p\n", strPtr)
	fmt.Printf("通过指针访问字符串: %s\n", *strPtr)

	var f float64 = 3.14159
	var fPtr *float64 = &f
	fmt.Printf("浮点数: %.5f\n", f)
	fmt.Printf("通过指针访问浮点数: %.5f\n", *fPtr)

	var b bool = true
	var bPtr *bool = &b
	fmt.Printf("布尔值: %t\n", b)
	fmt.Printf("通过指针访问布尔值: %t\n", *bPtr)

	// 指针和数组
	fmt.Println("\n--- 指针和数组 ---")
	arr := [5]int{1, 2, 3, 4, 5}
	var arrPtr *[5]int = &arr

	fmt.Printf("数组: %v\n", arr)
	fmt.Printf("通过指针访问数组: %v\n", *arrPtr)
	fmt.Printf("通过指针访问数组第一个元素: %d\n", (*arrPtr)[0])

	// 修改数组元素
	(*arrPtr)[0] = 10
	fmt.Printf("修改后的数组: %v\n", arr)

	// 指针和切片
	fmt.Println("\n--- 指针和切片 ---")
	slice := []int{10, 20, 30, 40, 50}
	var slicePtr *[]int = &slice

	fmt.Printf("切片: %v\n", slice)
	fmt.Printf("通过指针访问切片: %v\n", *slicePtr)

	// 修改切片
	*slicePtr = append(*slicePtr, 60)
	fmt.Printf("通过指针追加元素后: %v\n", slice)

	// 指针和结构体
	fmt.Println("\n--- 指针和结构体 ---")
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 25}
	var personPtr *Person = &person

	fmt.Printf("结构体: %+v\n", person)
	fmt.Printf("通过指针访问结构体: %+v\n", *personPtr)
	fmt.Printf("通过指针访问姓名: %s\n", personPtr.Name)   // Go自动解引用
	fmt.Printf("通过指针访问年龄: %d\n", (*personPtr).Age) // 显式解引用

	// 修改结构体字段
	personPtr.Age = 26
	fmt.Printf("修改后的结构体: %+v\n", person)

	// 指针作为函数参数
	fmt.Println("\n--- 指针作为函数参数 ---")
	num := 10
	fmt.Printf("调用前num的值: %d\n", num)

	// 值传递
	incrementValue(num)
	fmt.Printf("值传递后num的值: %d\n", num)

	// 指针传递
	incrementPointer(&num)
	fmt.Printf("指针传递后num的值: %d\n", num)

	// 函数返回指针
	fmt.Println("\n--- 函数返回指针 ---")
	newPersonPtr := createPerson("Bob", 30)
	fmt.Printf("新创建的Person: %+v\n", *newPersonPtr)

	// 指针比较
	fmt.Println("\n--- 指针比较 ---")
	var1 := 10
	var2 := 10
	ptrVar1 := &var1
	ptrVar2 := &var2
	ptrVar1_2 := &var1

	fmt.Printf("ptrVar1 == ptrVar2: %t (不同变量的地址)\n", ptrVar1 == ptrVar2)
	fmt.Printf("ptrVar1 == ptrVar1_2: %t (同一变量的地址)\n", ptrVar1 == ptrVar1_2)
	fmt.Printf("*ptrVar1 == *ptrVar2: %t (值相等)\n", *ptrVar1 == *ptrVar2)

	// nil指针检查
	fmt.Println("\n--- nil指针检查 ---")
	var nilPtr *int
	fmt.Printf("nil指针: %v\n", nilPtr)
	fmt.Printf("指针是否为nil: %t\n", nilPtr == nil)

	// 安全的指针操作
	if nilPtr != nil {
		fmt.Printf("指针值: %d\n", *nilPtr)
	} else {
		fmt.Println("指针为nil，不能解引用")
	}

	// new函数创建指针
	fmt.Println("\n--- new函数创建指针 ---")
	intPtr := new(int)
	fmt.Printf("new创建的int指针: %p\n", intPtr)
	fmt.Printf("new创建的int值: %d\n", *intPtr) // 零值

	*intPtr = 42
	fmt.Printf("赋值后的值: %d\n", *intPtr)

	personPtr2 := new(Person)
	fmt.Printf("new创建的Person指针: %+v\n", *personPtr2) // 零值结构体

	personPtr2.Name = "Charlie"
	personPtr2.Age = 35
	fmt.Printf("赋值后的Person: %+v\n", *personPtr2)

	// 指针和映射
	fmt.Println("\n--- 指针和映射 ---")
	m := make(map[string]int)
	m["key1"] = 100

	var mapPtr *map[string]int = &m
	fmt.Printf("映射: %v\n", m)
	fmt.Printf("通过指针访问映射: %v\n", *mapPtr)

	(*mapPtr)["key2"] = 200
	fmt.Printf("通过指针修改映射后: %v\n", m)

	// 指针算术（Go不支持，但可以展示为什么）
	fmt.Println("\n--- 指针特性 ---")
	fmt.Println("Go不支持指针算术，这使得Go更安全")
	fmt.Println("Go的指针不能进行加减运算")
	fmt.Println("Go的指针不能转换为整数")

	// 展示指针的内存地址
	var demo1 int = 1
	var demo2 int = 2
	fmt.Printf("demo1地址: %p\n", &demo1)
	fmt.Printf("demo2地址: %p\n", &demo2)
	fmt.Printf("地址差异展示了内存布局\n")
}

// incrementValue 值传递，不会修改原变量
func incrementValue(n int) {
	n++
	fmt.Printf("函数内n的值: %d\n", n)
}

// incrementPointer 指针传递，会修改原变量
func incrementPointer(n *int) {
	*n++
	fmt.Printf("函数内通过指针修改后的值: %d\n", *n)
}

// createPerson 返回指向新创建Person的指针
func createPerson(name string, age int) *Person {
	// 局部变量，但返回其指针是安全的
	// Go会自动将其分配到堆上
	p := Person{Name: name, Age: age}
	return &p
}

type Person struct {
	Name string
	Age  int
}
