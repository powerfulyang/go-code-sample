package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

// 🎓 学习导向的测试 - 通过测试学习Go反射

// DeepCopy 深拷贝函数（简化版本）
func DeepCopy(src interface{}) interface{} {
	srcVal := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)

	// 创建新的值
	dstVal := reflect.New(srcType).Elem()

	// 递归拷贝
	copyValue(srcVal, dstVal)

	return dstVal.Interface()
}

// copyValue 递归拷贝值
func copyValue(src, dst reflect.Value) {
	switch src.Kind() {
	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			copyValue(src.Field(i), dst.Field(i))
		}
	case reflect.Slice:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			copyValue(src.Index(i), dst.Index(i))
		}
	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {
			dstKey := reflect.New(key.Type()).Elem()
			copyValue(key, dstKey)
			dstVal := reflect.New(src.MapIndex(key).Type()).Elem()
			copyValue(src.MapIndex(key), dstVal)
			dst.SetMapIndex(dstKey, dstVal)
		}
	default:
		if dst.CanSet() {
			dst.Set(src)
		}
	}
}

// TestLearnBasicReflection 学习反射基础
func TestLearnBasicReflection(t *testing.T) {
	t.Log("🎯 学习目标: 理解Go反射的基本概念和使用")
	t.Log("📚 本测试将教您: Type、Value、Kind的区别和使用")

	t.Run("学习Type和Value", func(t *testing.T) {
		t.Log("📖 知识点: reflect.Type描述类型信息，reflect.Value描述值信息")

		// 🔍 探索: 不同类型的反射信息
		values := []interface{}{
			42,
			"hello",
			3.14,
			true,
			[]int{1, 2, 3},
			map[string]int{"a": 1},
			struct{ Name string }{"Go"},
		}

		t.Log("🔍 反射信息探索:")
		for i, v := range values {
			typ := reflect.TypeOf(v)
			val := reflect.ValueOf(v)

			t.Logf("   值%d: %v", i+1, v)
			t.Logf("     Type: %v", typ)
			t.Logf("     Kind: %v", typ.Kind())
			t.Logf("     Value: %v", val)
			t.Logf("     CanSet: %t", val.CanSet())
			t.Log("")
		}

		// ✅ 验证反射基础
		intType := reflect.TypeOf(42)
		if intType.Kind() != reflect.Int {
			t.Errorf("❌ int类型的Kind应该是reflect.Int，得到%v", intType.Kind())
		}

		stringVal := reflect.ValueOf("hello")
		if stringVal.String() != "hello" {
			t.Errorf("❌ 字符串值错误: 期望'hello'，得到'%s'", stringVal.String())
		}

		t.Log("✅ 很好！您理解了Type和Value的基本概念")

		// 💡 学习提示
		t.Log("💡 Type vs Kind: Type是具体类型，Kind是基础分类")
		t.Log("💡 Value操作: Value提供了访问和修改值的方法")
		t.Log("💡 性能考虑: 反射比直接操作慢，谨慎使用")
	})

	t.Run("学习结构体反射", func(t *testing.T) {
		t.Log("📖 知识点: 反射可以检查结构体的字段和方法")

		// 🔍 探索: 结构体反射
		type Person struct {
			Name    string `json:"name" validate:"required"`
			Age     int    `json:"age" validate:"min=0,max=150"`
			Email   string `json:"email" validate:"email"`
			private string // 私有字段
		}

		person := Person{
			Name:    "张三",
			Age:     25,
			Email:   "zhangsan@example.com",
			private: "secret",
		}

		typ := reflect.TypeOf(person)
		val := reflect.ValueOf(person)

		t.Logf("🔍 结构体反射分析:")
		t.Logf("   类型名称: %s", typ.Name())
		t.Logf("   字段数量: %d", typ.NumField())

		// 遍历字段
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			fieldVal := val.Field(i)

			t.Logf("   字段%d: %s", i+1, field.Name)
			t.Logf("     类型: %v", field.Type)
			t.Logf("     标签: %s", field.Tag)
			t.Logf("     JSON标签: %s", field.Tag.Get("json"))
			t.Logf("     验证标签: %s", field.Tag.Get("validate"))

			if fieldVal.CanInterface() {
				t.Logf("     值: %v", fieldVal.Interface())
			} else {
				t.Logf("     值: <不可访问>")
			}
			t.Log("")
		}

		// ✅ 验证结构体反射
		if typ.NumField() != 4 {
			t.Errorf("❌ 字段数量错误: 期望4，得到%d", typ.NumField())
		}

		nameField, found := typ.FieldByName("Name")
		if !found {
			t.Error("❌ 应该找到Name字段")
		} else if nameField.Tag.Get("json") != "name" {
			t.Errorf("❌ Name字段的json标签错误: 期望'name'，得到'%s'", nameField.Tag.Get("json"))
		}

		t.Log("✅ 很好！您理解了结构体反射")

		// 💡 学习提示
		t.Log("💡 字段访问: 使用Field()和FieldByName()访问字段")
		t.Log("💡 标签解析: Tag.Get()可以获取结构体标签")
		t.Log("💡 可见性: 私有字段可以通过反射访问但不能Interface()")
	})
}

// TestLearnReflectionModification 学习反射修改
func TestLearnReflectionModification(t *testing.T) {
	t.Log("🎯 学习目标: 掌握通过反射修改值")
	t.Log("📚 本测试将教您: 可设置性、指针反射、切片和映射操作")

	t.Run("学习值的可设置性", func(t *testing.T) {
		t.Log("📖 知识点: 只有可寻址的值才能通过反射修改")

		// 🔍 探索: 不同情况下的可设置性
		x := 42

		// 直接值 - 不可设置
		val1 := reflect.ValueOf(x)
		t.Logf("🔍 可设置性测试:")
		t.Logf("   直接值 CanSet: %t", val1.CanSet())

		// 指针的元素 - 可设置
		val2 := reflect.ValueOf(&x).Elem()
		t.Logf("   指针元素 CanSet: %t", val2.CanSet())

		// 修改值
		if val2.CanSet() {
			oldVal := val2.Int()
			val2.SetInt(100)
			t.Logf("   修改前: %d, 修改后: %d", oldVal, x)
		}

		// ✅ 验证值修改
		if x != 100 {
			t.Errorf("❌ 值修改失败: 期望100，得到%d", x)
		} else {
			t.Log("✅ 很好！您理解了值的可设置性")
		}

		// 💡 学习提示
		t.Log("💡 可寻址性: 只有可寻址的值才能修改")
		t.Log("💡 指针操作: 使用Elem()获取指针指向的值")
		t.Log("💡 类型匹配: SetXxx方法必须与值的类型匹配")
	})

	t.Run("学习结构体字段修改", func(t *testing.T) {
		t.Log("📖 知识点: 可以通过反射修改结构体的可导出字段")

		// 🔍 探索: 结构体字段修改
		type Config struct {
			Host    string
			Port    int
			Enabled bool
			private string
		}

		config := &Config{
			Host:    "localhost",
			Port:    8080,
			Enabled: false,
			private: "secret",
		}

		val := reflect.ValueOf(config).Elem()

		t.Logf("🔍 结构体字段修改:")
		t.Logf("   修改前: %+v", config)

		// 修改字段
		hostField := val.FieldByName("Host")
		if hostField.CanSet() {
			hostField.SetString("example.com")
		}

		portField := val.FieldByName("Port")
		if portField.CanSet() {
			portField.SetInt(9090)
		}

		enabledField := val.FieldByName("Enabled")
		if enabledField.CanSet() {
			enabledField.SetBool(true)
		}

		// 尝试修改私有字段
		privateField := val.FieldByName("private")
		t.Logf("   私有字段 CanSet: %t", privateField.CanSet())

		t.Logf("   修改后: %+v", config)

		// ✅ 验证字段修改
		if config.Host != "example.com" {
			t.Errorf("❌ Host修改失败: 期望'example.com'，得到'%s'", config.Host)
		}
		if config.Port != 9090 {
			t.Errorf("❌ Port修改失败: 期望9090，得到%d", config.Port)
		}
		if !config.Enabled {
			t.Error("❌ Enabled修改失败: 期望true，得到false")
		}

		t.Log("✅ 很好！您理解了结构体字段修改")

		// 💡 学习提示
		t.Log("💡 字段可见性: 只能修改可导出的字段")
		t.Log("💡 类型安全: SetXxx方法会检查类型匹配")
		t.Log("💡 实际应用: 配置注入、ORM映射等")
	})
}

// TestLearnReflectionMethods 学习方法反射
func TestLearnReflectionMethods(t *testing.T) {
	t.Log("🎯 学习目标: 掌握方法的反射调用")
	t.Log("📚 本测试将教您: 方法查找、参数准备、动态调用")

	t.Run("学习方法反射", func(t *testing.T) {
		t.Log("📖 知识点: 可以通过反射动态调用方法")

		// 🔍 探索: 方法反射
		calc := &Calculator{}
		val := reflect.ValueOf(calc)
		typ := reflect.TypeOf(calc)

		t.Logf("🔍 方法反射分析:")
		t.Logf("   方法数量: %d", typ.NumMethod())

		// 遍历方法
		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)
			t.Logf("   方法%d: %s", i+1, method.Name)
			t.Logf("     类型: %v", method.Type)
			t.Logf("     输入参数数量: %d", method.Type.NumIn())
			t.Logf("     输出参数数量: %d", method.Type.NumOut())
		}

		// 调用Add方法
		addMethod := val.MethodByName("Add")
		if addMethod.IsValid() {
			args := []reflect.Value{
				reflect.ValueOf(3.14),
				reflect.ValueOf(2.86),
			}
			results := addMethod.Call(args)

			if len(results) > 0 {
				result := results[0].Float()
				t.Logf("   Add(3.14, 2.86) = %.2f", result)
			}
		}

		// 调用GetResult方法
		getResultMethod := val.MethodByName("GetResult")
		if getResultMethod.IsValid() {
			results := getResultMethod.Call(nil)
			if len(results) > 0 {
				result := results[0].Float()
				t.Logf("   GetResult() = %.2f", result)
			}
		}

		// ✅ 验证方法调用
		if calc.Result != 6.0 {
			t.Errorf("❌ 方法调用结果错误: 期望6.0，得到%.2f", calc.Result)
		} else {
			t.Log("✅ 很好！您理解了方法反射")
		}

		// 💡 学习提示
		t.Log("💡 方法查找: 使用MethodByName()查找方法")
		t.Log("💡 参数准备: 参数必须包装为reflect.Value")
		t.Log("💡 返回值: Call()返回reflect.Value切片")
	})
}

// TestLearnReflectionPatterns 学习反射模式
func TestLearnReflectionPatterns(t *testing.T) {
	t.Log("🎯 学习目标: 掌握常用的反射编程模式")
	t.Log("📚 本测试将教您: 深拷贝、序列化、验证器")

	t.Run("学习深拷贝模式", func(t *testing.T) {
		t.Log("📖 知识点: 使用反射实现通用的深拷贝")

		// 🔍 探索: 反射深拷贝
		type Address struct {
			Street string
			City   string
		}

		type Person struct {
			Name    string
			Age     int
			Address Address
			Hobbies []string
		}

		original := Person{
			Name: "张三",
			Age:  25,
			Address: Address{
				Street: "中山路123号",
				City:   "北京",
			},
			Hobbies: []string{"读书", "游泳", "编程"},
		}

		// 使用已定义的深拷贝函数
		copied := DeepCopy(original).(Person)

		t.Logf("🔍 深拷贝测试:")
		t.Logf("   原始对象: %+v", original)
		t.Logf("   拷贝对象: %+v", copied)

		// 修改拷贝对象
		copied.Name = "李四"
		copied.Address.City = "上海"
		copied.Hobbies[0] = "电影"

		t.Logf("   修改拷贝后:")
		t.Logf("   原始对象: %+v", original)
		t.Logf("   拷贝对象: %+v", copied)

		// ✅ 验证深拷贝
		if original.Name == copied.Name {
			t.Error("❌ 深拷贝失败: 修改拷贝对象影响了原始对象")
		}
		if original.Address.City == copied.Address.City {
			t.Error("❌ 深拷贝失败: 嵌套结构体没有正确拷贝")
		}
		if original.Hobbies[0] == copied.Hobbies[0] {
			t.Error("❌ 深拷贝失败: 切片没有正确拷贝")
		}

		t.Log("✅ 很好！您理解了反射深拷贝模式")

		// 💡 学习提示
		t.Log("💡 递归处理: 深拷贝需要递归处理嵌套结构")
		t.Log("💡 类型判断: 不同类型需要不同的拷贝策略")
		t.Log("💡 性能考虑: 反射拷贝比手写拷贝慢")
	})
}

// BenchmarkLearnReflectionPerformance 学习反射性能
func BenchmarkLearnReflectionPerformance(b *testing.B) {
	b.Log("🎯 学习目标: 了解反射的性能开销")

	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "张三", Age: 25}

	b.Run("直接访问", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = person.Name
			_ = person.Age
		}
	})

	b.Run("反射访问", func(b *testing.B) {
		val := reflect.ValueOf(person)
		for i := 0; i < b.N; i++ {
			_ = val.FieldByName("Name").String()
			_ = val.FieldByName("Age").Int()
		}
	})

	b.Run("缓存反射", func(b *testing.B) {
		val := reflect.ValueOf(person)
		nameField := val.FieldByName("Name")
		ageField := val.FieldByName("Age")

		for i := 0; i < b.N; i++ {
			_ = nameField.String()
			_ = ageField.Int()
		}
	})
}

// Example_learnBasicReflection 反射基础示例
func Example_learnBasicReflection() {
	// 获取类型和值信息
	x := 42
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("类型: %v\n", t)
	fmt.Printf("种类: %v\n", t.Kind())
	fmt.Printf("值: %v\n", v)
	fmt.Printf("整数值: %d\n", v.Int())

	// Output:
	// 类型: int
	// 种类: int
	// 值: 42
	// 整数值: 42
}
