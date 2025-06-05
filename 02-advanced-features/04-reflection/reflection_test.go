package reflection

import (
	"reflect"
	"testing"
	"time"
)

func TestTypeInfo(t *testing.T) {
	t.Run("BasicTypes", func(t *testing.T) {
		// 测试基本类型
		testCases := []struct {
			value    interface{}
			typeName string
			kind     reflect.Kind
		}{
			{42, "int", reflect.Int},
			{"hello", "string", reflect.String},
			{3.14, "float64", reflect.Float64},
			{true, "bool", reflect.Bool},
		}

		for _, tc := range testCases {
			typ := reflect.TypeOf(tc.value)
			if typ.Name() != tc.typeName {
				t.Errorf("类型名称错误: 期望 %s, 实际 %s", tc.typeName, typ.Name())
			}
			if typ.Kind() != tc.kind {
				t.Errorf("类型种类错误: 期望 %s, 实际 %s", tc.kind, typ.Kind())
			}
		}

		t.Log("基本类型反射测试通过")
	})

	t.Run("PointerTypes", func(t *testing.T) {
		var ptr *int = new(int)
		*ptr = 100

		typ := reflect.TypeOf(ptr)
		val := reflect.ValueOf(ptr)

		if typ.Kind() != reflect.Ptr {
			t.Errorf("应该是指针类型, 实际: %s", typ.Kind())
		}

		if typ.Elem().Kind() != reflect.Int {
			t.Errorf("指针指向类型应该是int, 实际: %s", typ.Elem().Kind())
		}

		if val.IsNil() {
			t.Error("指针不应该为nil")
		}

		if val.Elem().Interface().(int) != 100 {
			t.Errorf("指针指向的值应该是100, 实际: %v", val.Elem().Interface())
		}

		t.Log("指针类型反射测试通过")
	})
}

func TestStructInfo(t *testing.T) {
	t.Run("UserStruct", func(t *testing.T) {
		user := User{
			ID:       1,
			Name:     "测试用户",
			Email:    "test@example.com",
			Age:      25,
			IsActive: true,
			Tags:     []string{"test", "user"},
			Profile: &Profile{
				Bio:     "测试简介",
				Website: "https://test.com",
			},
			Created: time.Now(),
		}

		typ := reflect.TypeOf(user)
		val := reflect.ValueOf(user)

		if typ.Kind() != reflect.Struct {
			t.Errorf("应该是结构体类型, 实际: %s", typ.Kind())
		}

		if typ.Name() != "User" {
			t.Errorf("结构体名称应该是User, 实际: %s", typ.Name())
		}

		expectedFields := []string{"ID", "Name", "Email", "Age", "IsActive", "Tags", "Profile", "Created"}
		if typ.NumField() != len(expectedFields) {
			t.Errorf("字段数量应该是%d, 实际: %d", len(expectedFields), typ.NumField())
		}

		// 检查特定字段
		nameField, found := typ.FieldByName("Name")
		if !found {
			t.Error("应该找到Name字段")
		}

		if nameField.Type.Kind() != reflect.String {
			t.Errorf("Name字段应该是string类型, 实际: %s", nameField.Type.Kind())
		}

		// 检查标签
		if jsonTag := nameField.Tag.Get("json"); jsonTag != "name" {
			t.Errorf("Name字段的json标签应该是'name', 实际: %s", jsonTag)
		}

		// 检查字段值
		nameValue := val.FieldByName("Name")
		if nameValue.String() != "测试用户" {
			t.Errorf("Name字段值应该是'测试用户', 实际: %s", nameValue.String())
		}

		t.Log("结构体反射测试通过")
	})
}

func TestMethodInfo(t *testing.T) {
	t.Run("CalculatorMethods", func(t *testing.T) {
		calc := &Calculator{}
		typ := reflect.TypeOf(calc)

		expectedMethods := []string{"Add", "GetResult", "Multiply"}
		if typ.NumMethod() != len(expectedMethods) {
			t.Errorf("方法数量应该是%d, 实际: %d", len(expectedMethods), typ.NumMethod())
		}

		// 检查Add方法
		addMethod, found := typ.MethodByName("Add")
		if !found {
			t.Error("应该找到Add方法")
		}

		// Add方法应该有3个输入参数（receiver + 2个float64）
		if addMethod.Type.NumIn() != 3 {
			t.Errorf("Add方法输入参数应该是3个, 实际: %d", addMethod.Type.NumIn())
		}

		// Add方法应该有1个输出参数
		if addMethod.Type.NumOut() != 1 {
			t.Errorf("Add方法输出参数应该是1个, 实际: %d", addMethod.Type.NumOut())
		}

		t.Log("方法反射测试通过")
	})
}

func TestCallMethod(t *testing.T) {
	t.Run("DynamicMethodCall", func(t *testing.T) {
		calc := &Calculator{}

		// 测试Add方法
		results, err := CallMethod(calc, "Add", 10.5, 20.3)
		if err != nil {
			t.Errorf("调用Add方法失败: %v", err)
		}

		if len(results) != 1 {
			t.Errorf("Add方法应该返回1个值, 实际: %d", len(results))
		}

		result := results[0].Float()
		expected := 30.8
		if result != expected {
			t.Errorf("Add(10.5, 20.3)应该返回%f, 实际: %f", expected, result)
		}

		// 测试不存在的方法
		_, err = CallMethod(calc, "NonExistentMethod")
		if err == nil {
			t.Error("调用不存在的方法应该返回错误")
		}

		t.Log("动态方法调用测试通过")
	})
}

func TestSliceInfo(t *testing.T) {
	t.Run("IntSlice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		typ := reflect.TypeOf(slice)
		val := reflect.ValueOf(slice)

		if typ.Kind() != reflect.Slice {
			t.Errorf("应该是切片类型, 实际: %s", typ.Kind())
		}

		if typ.Elem().Kind() != reflect.Int {
			t.Errorf("切片元素应该是int类型, 实际: %s", typ.Elem().Kind())
		}

		if val.Len() != 5 {
			t.Errorf("切片长度应该是5, 实际: %d", val.Len())
		}

		// 检查元素
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			expected := i + 1
			if int(elem.Int()) != expected {
				t.Errorf("元素[%d]应该是%d, 实际: %d", i, expected, elem.Int())
			}
		}

		t.Log("切片反射测试通过")
	})
}

func TestMapInfo(t *testing.T) {
	t.Run("StringIntMap", func(t *testing.T) {
		m := map[string]int{
			"apple":  5,
			"banana": 3,
			"orange": 8,
		}

		typ := reflect.TypeOf(m)
		val := reflect.ValueOf(m)

		if typ.Kind() != reflect.Map {
			t.Errorf("应该是映射类型, 实际: %s", typ.Kind())
		}

		if typ.Key().Kind() != reflect.String {
			t.Errorf("映射键应该是string类型, 实际: %s", typ.Key().Kind())
		}

		if typ.Elem().Kind() != reflect.Int {
			t.Errorf("映射值应该是int类型, 实际: %s", typ.Elem().Kind())
		}

		if val.Len() != 3 {
			t.Errorf("映射长度应该是3, 实际: %d", val.Len())
		}

		// 检查特定键值对
		appleValue := val.MapIndex(reflect.ValueOf("apple"))
		if !appleValue.IsValid() {
			t.Error("应该找到apple键")
		}

		if int(appleValue.Int()) != 5 {
			t.Errorf("apple的值应该是5, 实际: %d", appleValue.Int())
		}

		t.Log("映射反射测试通过")
	})
}

func TestCreateStruct(t *testing.T) {
	t.Run("DynamicStructCreation", func(t *testing.T) {
		instance := CreateStruct()

		val := reflect.ValueOf(instance)
		typ := reflect.TypeOf(instance)

		if typ.Kind() != reflect.Struct {
			t.Errorf("应该是结构体类型, 实际: %s", typ.Kind())
		}

		// 检查字段
		nameField := val.FieldByName("Name")
		if !nameField.IsValid() {
			t.Error("应该有Name字段")
		}

		if nameField.String() != "动态用户" {
			t.Errorf("Name字段值应该是'动态用户', 实际: %s", nameField.String())
		}

		ageField := val.FieldByName("Age")
		if !ageField.IsValid() {
			t.Error("应该有Age字段")
		}

		if int(ageField.Int()) != 25 {
			t.Errorf("Age字段值应该是25, 实际: %d", ageField.Int())
		}

		t.Log("动态结构体创建测试通过")
	})
}

func TestUtilityFunctions(t *testing.T) {
	t.Run("DeepEqual", func(t *testing.T) {
		user1 := User{ID: 1, Name: "张三", Age: 25}
		user2 := User{ID: 1, Name: "张三", Age: 25}
		user3 := User{ID: 2, Name: "李四", Age: 30}

		if !DeepEqual(user1, user2) {
			t.Error("相同的结构体应该深度相等")
		}

		if DeepEqual(user1, user3) {
			t.Error("不同的结构体不应该深度相等")
		}

		t.Log("深度比较测试通过")
	})

	t.Run("IsNil", func(t *testing.T) {
		var nilPtr *User
		var validPtr = &User{}

		if !IsNil(nilPtr) {
			t.Error("nil指针应该被识别为nil")
		}

		if IsNil(validPtr) {
			t.Error("有效指针不应该被识别为nil")
		}

		if IsNil(42) {
			t.Error("基本类型不应该被识别为nil")
		}

		t.Log("nil检查测试通过")
	})

	t.Run("ToMap", func(t *testing.T) {
		user := User{
			ID:   1,
			Name: "张三",
			Age:  25,
		}

		m := ToMap(user)

		if m["id"] != 1 {
			t.Errorf("映射中id应该是1, 实际: %v", m["id"])
		}

		if m["name"] != "张三" {
			t.Errorf("映射中name应该是'张三', 实际: %v", m["name"])
		}

		if m["age"] != 25 {
			t.Errorf("映射中age应该是25, 实际: %v", m["age"])
		}

		t.Log("结构体转映射测试通过")
	})

	t.Run("FromMap", func(t *testing.T) {
		data := map[string]interface{}{
			"id":   2,
			"name": "李四",
			"age":  30,
		}

		var user User
		err := FromMap(data, &user)
		if err != nil {
			t.Errorf("从映射创建结构体失败: %v", err)
		}

		if user.ID != 2 {
			t.Errorf("ID应该是2, 实际: %d", user.ID)
		}

		if user.Name != "李四" {
			t.Errorf("Name应该是'李四', 实际: %s", user.Name)
		}

		if user.Age != 30 {
			t.Errorf("Age应该是30, 实际: %d", user.Age)
		}

		t.Log("映射转结构体测试通过")
	})

	t.Run("Clone", func(t *testing.T) {
		original := User{
			ID:   1,
			Name: "原始用户",
			Tags: []string{"tag1", "tag2"},
		}

		cloned := Clone(original).(User)

		// 检查克隆是否正确
		if !DeepEqual(original, cloned) {
			t.Error("克隆的对象应该与原始对象深度相等")
		}

		// 修改原始对象，确保克隆不受影响
		original.Name = "修改后的用户"
		original.Tags[0] = "modified"

		if cloned.Name == "修改后的用户" {
			t.Error("修改原始对象不应该影响克隆对象")
		}

		if cloned.Tags[0] == "modified" {
			t.Error("修改原始对象的切片不应该影响克隆对象")
		}

		t.Log("克隆测试通过")
	})
}

func TestModifyStruct(t *testing.T) {
	t.Run("StructModification", func(t *testing.T) {
		user := User{
			ID:   1,
			Name: "原始名称",
			Age:  20,
		}

		// 修改结构体
		v := reflect.ValueOf(&user).Elem()

		// 修改Name字段
		nameField := v.FieldByName("Name")
		if nameField.IsValid() && nameField.CanSet() {
			nameField.SetString("修改后的名称")
		}

		// 修改Age字段
		ageField := v.FieldByName("Age")
		if ageField.IsValid() && ageField.CanSet() {
			ageField.SetInt(25)
		}

		// 验证修改
		if user.Name != "修改后的名称" {
			t.Errorf("Name应该被修改为'修改后的名称', 实际: %s", user.Name)
		}

		if user.Age != 25 {
			t.Errorf("Age应该被修改为25, 实际: %d", user.Age)
		}

		t.Log("结构体修改测试通过")
	})
}

// 基准测试
func BenchmarkTypeOf(b *testing.B) {
	value := "test string"
	for i := 0; i < b.N; i++ {
		_ = reflect.TypeOf(value)
	}
}

func BenchmarkValueOf(b *testing.B) {
	value := "test string"
	for i := 0; i < b.N; i++ {
		_ = reflect.ValueOf(value)
	}
}

func BenchmarkMethodCall(b *testing.B) {
	calc := &Calculator{}
	for i := 0; i < b.N; i++ {
		CallMethod(calc, "Add", 1.0, 2.0)
	}
}

func BenchmarkDirectMethodCall(b *testing.B) {
	calc := &Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Add(1.0, 2.0)
	}
}

func BenchmarkToMap(b *testing.B) {
	user := User{
		ID:   1,
		Name: "测试用户",
		Age:  25,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMap(user)
	}
}

func BenchmarkFromMap(b *testing.B) {
	data := map[string]interface{}{
		"id":   1,
		"name": "测试用户",
		"age":  25,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var user User
		FromMap(data, &user)
	}
}
