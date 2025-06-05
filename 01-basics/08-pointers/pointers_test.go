package pointers

import (
	"testing"
	"unsafe"
)

func TestBasicPointers(t *testing.T) {
	t.Run("PointerBasics", func(t *testing.T) {
		// 基本指针操作
		num := 42
		str := "Hello"

		numPtr := &num
		strPtr := &str

		t.Logf("变量 num: %d, 地址: %p", num, &num)
		t.Logf("变量 str: %s, 地址: %p", str, &str)
		t.Logf("指针 numPtr: %p, 指向的值: %d", numPtr, *numPtr)
		t.Logf("指针 strPtr: %p, 指向的值: %s", strPtr, *strPtr)

		// 验证指针指向正确的值
		if *numPtr != 42 {
			t.Errorf("numPtr 应该指向 42, 实际指向 %d", *numPtr)
		}
		if *strPtr != "Hello" {
			t.Errorf("strPtr 应该指向 'Hello', 实际指向 '%s'", *strPtr)
		}

		// 通过指针修改值
		*numPtr = 100
		*strPtr = "World"

		t.Logf("修改后 num: %d", num)
		t.Logf("修改后 str: %s", str)

		// 验证修改结果
		if num != 100 {
			t.Errorf("通过指针修改后 num 应该是 100, 实际是 %d", num)
		}
		if str != "World" {
			t.Errorf("通过指针修改后 str 应该是 'World', 实际是 '%s'", str)
		}
	})

	t.Run("PointerAddresses", func(t *testing.T) {
		// 测试指针地址
		var a, b int = 10, 20
		ptrA := &a
		ptrB := &b

		t.Logf("变量 a 地址: %p", &a)
		t.Logf("变量 b 地址: %p", &b)
		t.Logf("指针 ptrA: %p", ptrA)
		t.Logf("指针 ptrB: %p", ptrB)

		// 验证指针指向正确的地址
		if ptrA != &a {
			t.Error("ptrA 应该指向变量 a 的地址")
		}
		if ptrB != &b {
			t.Error("ptrB 应该指向变量 b 的地址")
		}

		// 验证不同变量有不同地址
		if ptrA == ptrB {
			t.Error("不同变量应该有不同的地址")
		}
	})
}

func TestNilPointers(t *testing.T) {
	t.Run("NilPointerCheck", func(t *testing.T) {
		var ptr *int

		t.Logf("零值指针: %v", ptr)

		// 验证零值指针
		if ptr != nil {
			t.Error("未初始化的指针应该是 nil")
		}

		// 使用 new 分配内存
		ptr = new(int)
		t.Logf("new 分配后的指针: %p", ptr)
		t.Logf("new 分配后的值: %d", *ptr)

		// 验证 new 分配的指针不是 nil
		if ptr == nil {
			t.Error("new 分配的指针不应该是 nil")
		}

		// 验证 new 分配的值是零值
		if *ptr != 0 {
			t.Errorf("new 分配的 int 应该是 0, 实际是 %d", *ptr)
		}

		// 赋值并验证
		*ptr = 123
		if *ptr != 123 {
			t.Errorf("赋值后应该是 123, 实际是 %d", *ptr)
		}
	})

	t.Run("SafePointerAccess", func(t *testing.T) {
		var ptr *int

		// 安全访问指针
		if ptr != nil {
			t.Logf("指针值: %d", *ptr)
		} else {
			t.Log("指针为 nil，跳过解引用")
		}

		// 分配内存后访问
		ptr = new(int)
		*ptr = 456

		if ptr != nil {
			t.Logf("指针值: %d", *ptr)
			if *ptr != 456 {
				t.Errorf("期望 456, 实际 %d", *ptr)
			}
		}
	})
}

func TestPointerParameters(t *testing.T) {
	t.Run("ValueVsPointerPassing", func(t *testing.T) {
		x, y := 10, 20

		t.Logf("交换前: x=%d, y=%d", x, y)

		// 值传递不会改变原变量
		swapByValue(x, y)
		t.Logf("值传递后: x=%d, y=%d", x, y)

		if x != 10 || y != 20 {
			t.Error("值传递不应该改变原变量")
		}

		// 指针传递会改变原变量
		swapByPointer(&x, &y)
		t.Logf("指针传递后: x=%d, y=%d", x, y)

		if x != 20 || y != 10 {
			t.Error("指针传递应该交换原变量的值")
		}
	})

	t.Run("ModifyThroughPointer", func(t *testing.T) {
		num := 5
		t.Logf("修改前: %d", num)

		doubleValue(&num)
		t.Logf("修改后: %d", num)

		if num != 10 {
			t.Errorf("doubleValue 后应该是 10, 实际是 %d", num)
		}
	})
}

func TestPointersAndArrays(t *testing.T) {
	t.Run("ArrayPointer", func(t *testing.T) {
		arr := [5]int{1, 2, 3, 4, 5}
		t.Logf("原数组: %v", arr)

		// 数组指针
		arrPtr := &arr
		t.Logf("数组指针: %p", arrPtr)

		// 通过指针修改数组
		(*arrPtr)[0] = 100
		t.Logf("修改后数组: %v", arr)

		if arr[0] != 100 {
			t.Errorf("通过指针修改后第一个元素应该是 100, 实际是 %d", arr[0])
		}
	})

	t.Run("ElementPointers", func(t *testing.T) {
		arr := [3]int{10, 20, 30}

		// 获取元素指针
		firstPtr := &arr[0]
		lastPtr := &arr[2]

		t.Logf("第一个元素地址: %p, 值: %d", firstPtr, *firstPtr)
		t.Logf("最后一个元素地址: %p, 值: %d", lastPtr, *lastPtr)

		// 通过元素指针修改值
		*firstPtr = 100
		*lastPtr = 300

		t.Logf("修改后数组: %v", arr)

		if arr[0] != 100 || arr[2] != 300 {
			t.Error("通过元素指针修改失败")
		}

		// 验证地址差
		addressDiff := uintptr(unsafe.Pointer(lastPtr)) - uintptr(unsafe.Pointer(firstPtr))
		expectedDiff := uintptr(2 * 8) // 2个int，每个8字节（64位系统）
		t.Logf("地址差: %d 字节", addressDiff)

		if addressDiff != expectedDiff {
			t.Logf("地址差可能因系统而异，期望 %d, 实际 %d", expectedDiff, addressDiff)
		}
	})
}

func TestPointersAndSlices(t *testing.T) {
	t.Run("SliceModification", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		original := make([]int, len(slice))
		copy(original, slice)

		t.Logf("原切片: %v", slice)

		// 切片本身包含指向底层数组的指针
		modifySlice(slice)
		t.Logf("修改后切片: %v", slice)

		// 验证修改结果
		for i, v := range slice {
			expected := original[i] * 2
			if v != expected {
				t.Errorf("索引 %d: 期望 %d, 实际 %d", i, expected, v)
			}
		}
	})

	t.Run("SlicePointer", func(t *testing.T) {
		slice := []int{1, 2, 3}
		t.Logf("原切片: %v", slice)

		slicePtr := &slice

		// 通过指针修改切片
		*slicePtr = append(*slicePtr, 4, 5)
		t.Logf("通过指针添加元素后: %v", slice)

		if len(slice) != 5 {
			t.Errorf("添加元素后长度应该是 5, 实际是 %d", len(slice))
		}
	})

	t.Run("ElementPointers", func(t *testing.T) {
		slice := []int{10, 20, 30}

		// 获取元素指针并修改
		for i := range slice {
			elemPtr := &slice[i]
			*elemPtr *= 10
		}

		t.Logf("通过元素指针修改后: %v", slice)

		expected := []int{100, 200, 300}
		for i, v := range slice {
			if v != expected[i] {
				t.Errorf("索引 %d: 期望 %d, 实际 %d", i, expected[i], v)
			}
		}
	})
}

func TestPointersAndStructs(t *testing.T) {
	t.Run("StructPointer", func(t *testing.T) {
		person := Person{Name: "张三", Age: 25}
		t.Logf("原结构体: %+v", person)

		personPtr := &person

		// 通过指针访问字段
		name1 := (*personPtr).Name
		name2 := personPtr.Name // Go的语法糖

		if name1 != name2 {
			t.Error("两种访问方式应该得到相同结果")
		}

		// 通过指针修改字段
		personPtr.Name = "李四"
		personPtr.Age = 30

		t.Logf("修改后结构体: %+v", person)

		if person.Name != "李四" || person.Age != 30 {
			t.Error("通过指针修改结构体字段失败")
		}
	})

	t.Run("FieldPointers", func(t *testing.T) {
		person := Person{Name: "王五", Age: 35}

		// 获取字段指针
		namePtr := &person.Name
		agePtr := &person.Age

		t.Logf("姓名字段地址: %p, 值: %s", namePtr, *namePtr)
		t.Logf("年龄字段地址: %p, 值: %d", agePtr, *agePtr)

		// 通过字段指针修改
		*namePtr = "赵六"
		*agePtr = 40

		t.Logf("通过字段指针修改后: %+v", person)

		if person.Name != "赵六" || person.Age != 40 {
			t.Error("通过字段指针修改失败")
		}
	})
}

func TestPointersAndMethods(t *testing.T) {
	t.Run("PointerReceiverMethods", func(t *testing.T) {
		person := Person{Name: "钱七", Age: 28}
		t.Logf("初始状态: %s", person.GetInfo())

		// 指针接收者方法
		person.UpdateAge(32)
		t.Logf("更新年龄后: %s", person.GetInfo())

		if person.Age != 32 {
			t.Errorf("更新年龄后应该是 32, 实际是 %d", person.Age)
		}

		// 通过指针调用方法
		personPtr := &person
		personPtr.UpdateAge(35)
		t.Logf("通过指针更新后: %s", personPtr.GetInfo())

		if person.Age != 35 {
			t.Errorf("通过指针更新年龄后应该是 35, 实际是 %d", person.Age)
		}
	})

	t.Run("ValueReceiverMethods", func(t *testing.T) {
		person := Person{Name: "孙八", Age: 45}

		// 值接收者方法
		info := person.GetInfoByValue()
		t.Logf("值接收者方法: %s", info)

		expected := "姓名: 孙八, 年龄: 45"
		if info != expected {
			t.Errorf("期望 '%s', 实际 '%s'", expected, info)
		}
	})
}

func TestMemoryAllocation(t *testing.T) {
	t.Run("NewAllocation", func(t *testing.T) {
		// 使用 new 分配不同类型
		intPtr := new(int)
		stringPtr := new(string)
		personPtr := new(Person)

		t.Logf("new int: %d, 地址: %p", *intPtr, intPtr)
		t.Logf("new string: '%s', 地址: %p", *stringPtr, stringPtr)
		t.Logf("new Person: %+v, 地址: %p", *personPtr, personPtr)

		// 验证零值
		if *intPtr != 0 {
			t.Errorf("new int 应该是 0, 实际是 %d", *intPtr)
		}
		if *stringPtr != "" {
			t.Errorf("new string 应该是空字符串, 实际是 '%s'", *stringPtr)
		}
		if personPtr.Name != "" || personPtr.Age != 0 {
			t.Error("new Person 应该是零值结构体")
		}

		// 赋值
		*intPtr = 42
		*stringPtr = "Hello"
		personPtr.Name = "测试"
		personPtr.Age = 25

		t.Logf("赋值后 int: %d", *intPtr)
		t.Logf("赋值后 string: '%s'", *stringPtr)
		t.Logf("赋值后 Person: %+v", *personPtr)
	})

	t.Run("LiteralAllocation", func(t *testing.T) {
		// 使用字面量创建指针
		intPtr := &[]int{42}[0]
		personPtr := &Person{Name: "字面量", Age: 30}

		t.Logf("字面量 int 指针: %d, 地址: %p", *intPtr, intPtr)
		t.Logf("字面量 Person 指针: %+v, 地址: %p", *personPtr, personPtr)

		if *intPtr != 42 {
			t.Errorf("字面量 int 应该是 42, 实际是 %d", *intPtr)
		}
		if personPtr.Name != "字面量" || personPtr.Age != 30 {
			t.Error("字面量 Person 字段值不正确")
		}
	})
}

func TestPointerArrays(t *testing.T) {
	t.Run("ArrayOfPointers", func(t *testing.T) {
		// 指针数组
		var ptrArray [3]*int
		a, b, c := 1, 2, 3
		ptrArray[0] = &a
		ptrArray[1] = &b
		ptrArray[2] = &c

		t.Log("指针数组:")
		for i, ptr := range ptrArray {
			t.Logf("  索引 %d: 地址 %p, 值 %d", i, ptr, *ptr)
		}

		// 修改原变量
		a, b, c = 10, 20, 30

		t.Log("修改原变量后:")
		for i, ptr := range ptrArray {
			t.Logf("  索引 %d: 地址 %p, 值 %d", i, ptr, *ptr)
		}

		// 验证指针数组反映了原变量的变化
		if *ptrArray[0] != 10 || *ptrArray[1] != 20 || *ptrArray[2] != 30 {
			t.Error("指针数组应该反映原变量的变化")
		}
	})

	t.Run("PointerToArray", func(t *testing.T) {
		// 数组指针
		arr := [3]int{100, 200, 300}
		arrPtr := &arr

		t.Logf("数组指针指向的数组: %v", *arrPtr)

		// 通过数组指针修改
		(*arrPtr)[1] = 250

		t.Logf("修改后的数组: %v", arr)

		if arr[1] != 250 {
			t.Errorf("通过数组指针修改后应该是 250, 实际是 %d", arr[1])
		}
	})
}

func TestMultiLevelPointers(t *testing.T) {
	t.Run("MultipleIndirection", func(t *testing.T) {
		value := 42
		ptr1 := &value
		ptr2 := &ptr1
		ptr3 := &ptr2

		t.Logf("原值: %d", value)
		t.Logf("一级指针: %p, 值: %d", ptr1, *ptr1)
		t.Logf("二级指针: %p, 值: %d", ptr2, **ptr2)
		t.Logf("三级指针: %p, 值: %d", ptr3, ***ptr3)

		// 验证多级指针指向同一个值
		if *ptr1 != value || **ptr2 != value || ***ptr3 != value {
			t.Error("多级指针应该指向同一个值")
		}

		// 通过三级指针修改值
		***ptr3 = 100

		t.Logf("通过三级指针修改后: %d", value)

		if value != 100 {
			t.Errorf("通过三级指针修改后应该是 100, 实际是 %d", value)
		}
	})
}

func TestPointerComparison(t *testing.T) {
	t.Run("PointerEquality", func(t *testing.T) {
		a := 10
		b := 10
		ptrA1 := &a
		ptrA2 := &a
		ptrB := &b

		t.Logf("ptrA1 == ptrA2: %t (同一变量)", ptrA1 == ptrA2)
		t.Logf("ptrA1 == ptrB: %t (不同变量)", ptrA1 == ptrB)

		// 验证指针比较
		if ptrA1 != ptrA2 {
			t.Error("指向同一变量的指针应该相等")
		}
		if ptrA1 == ptrB {
			t.Error("指向不同变量的指针不应该相等")
		}

		// nil 指针比较
		var nilPtr1, nilPtr2 *int
		if nilPtr1 != nilPtr2 {
			t.Error("nil 指针应该相等")
		}
		if nilPtr1 == ptrA1 {
			t.Error("nil 指针不应该等于非 nil 指针")
		}
	})
}

// 基准测试
func BenchmarkPointerAccess(b *testing.B) {
	value := 42
	ptr := &value

	for i := 0; i < b.N; i++ {
		_ = *ptr
	}
}

func BenchmarkDirectAccess(b *testing.B) {
	value := 42

	for i := 0; i < b.N; i++ {
		_ = value
	}
}
