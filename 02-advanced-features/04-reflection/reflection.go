package reflection

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 基本反射示例

// TypeInfo 获取类型信息
func TypeInfo(value interface{}) {
	fmt.Printf("=== 类型信息分析 ===\n")

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	fmt.Printf("值: %v\n", value)
	fmt.Printf("类型: %s\n", t)
	fmt.Printf("类型种类: %s\n", t.Kind())
	fmt.Printf("类型名称: %s\n", t.Name())
	fmt.Printf("包路径: %s\n", t.PkgPath())
	fmt.Printf("是否可设置: %t\n", v.CanSet())
	fmt.Printf("是否为零值: %t\n", v.IsZero())

	if t.Kind() == reflect.Ptr {
		fmt.Printf("指针指向类型: %s\n", t.Elem())
		if !v.IsNil() {
			fmt.Printf("指针指向值: %v\n", v.Elem().Interface())
		}
	}
}

// 结构体反射示例

// User 示例结构体
type User struct {
	ID       int       `json:"id" db:"user_id" validate:"required"`
	Name     string    `json:"name" db:"username" validate:"required,min=2"`
	Email    string    `json:"email" db:"email" validate:"required,email"`
	Age      int       `json:"age" db:"age" validate:"min=0,max=120"`
	IsActive bool      `json:"is_active" db:"is_active"`
	Tags     []string  `json:"tags" db:"tags"`
	Profile  *Profile  `json:"profile,omitempty" db:"profile"`
	Created  time.Time `json:"created" db:"created_at"`
}

// Profile 用户资料
type Profile struct {
	Bio     string `json:"bio" db:"bio"`
	Website string `json:"website" db:"website"`
}

// StructInfo 分析结构体信息
func StructInfo(s interface{}) {
	fmt.Printf("\n=== 结构体信息分析 ===\n")

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// 如果是指针，获取指向的类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		fmt.Printf("不是结构体类型: %s\n", t.Kind())
		return
	}

	fmt.Printf("结构体名称: %s\n", t.Name())
	fmt.Printf("字段数量: %d\n", t.NumField())
	fmt.Printf("方法数量: %d\n", t.NumMethod())

	// 遍历字段
	fmt.Println("\n字段信息:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		fmt.Printf("  字段 %d:\n", i+1)
		fmt.Printf("    名称: %s\n", field.Name)
		fmt.Printf("    类型: %s\n", field.Type)
		fmt.Printf("    标签: %s\n", field.Tag)
		fmt.Printf("    是否导出: %t\n", field.IsExported())
		fmt.Printf("    是否可设置: %t\n", fieldValue.CanSet())

		if fieldValue.IsValid() && fieldValue.CanInterface() {
			fmt.Printf("    值: %v\n", fieldValue.Interface())
		}

		// 解析标签
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			fmt.Printf("    JSON标签: %s\n", jsonTag)
		}
		if dbTag := field.Tag.Get("db"); dbTag != "" {
			fmt.Printf("    DB标签: %s\n", dbTag)
		}
		if validateTag := field.Tag.Get("validate"); validateTag != "" {
			fmt.Printf("    验证标签: %s\n", validateTag)
		}
		fmt.Println()
	}
}

// 动态调用方法

// Calculator 计算器示例
type Calculator struct {
	Result float64
}

// Add 加法
func (c *Calculator) Add(a, b float64) float64 {
	c.Result = a + b
	return c.Result
}

// Multiply 乘法
func (c *Calculator) Multiply(a, b float64) float64 {
	c.Result = a * b
	return c.Result
}

// GetResult 获取结果
func (c *Calculator) GetResult() float64 {
	return c.Result
}

// CallMethod 动态调用方法
func CallMethod(obj interface{}, methodName string, args ...interface{}) ([]reflect.Value, error) {
	v := reflect.ValueOf(obj)
	method := v.MethodByName(methodName)

	if !method.IsValid() {
		return nil, fmt.Errorf("方法 %s 不存在", methodName)
	}

	// 准备参数
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// 调用方法
	results := method.Call(in)
	return results, nil
}

// MethodInfo 分析方法信息
func MethodInfo(obj interface{}) {
	fmt.Printf("\n=== 方法信息分析 ===\n")

	t := reflect.TypeOf(obj)

	fmt.Printf("类型: %s\n", t)
	fmt.Printf("方法数量: %d\n", t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("\n方法 %d:\n", i+1)
		fmt.Printf("  名称: %s\n", method.Name)
		fmt.Printf("  类型: %s\n", method.Type)
		fmt.Printf("  输入参数数量: %d\n", method.Type.NumIn())
		fmt.Printf("  输出参数数量: %d\n", method.Type.NumOut())

		// 参数类型
		for j := 0; j < method.Type.NumIn(); j++ {
			fmt.Printf("    输入参数 %d: %s\n", j, method.Type.In(j))
		}

		for j := 0; j < method.Type.NumOut(); j++ {
			fmt.Printf("    输出参数 %d: %s\n", j, method.Type.Out(j))
		}
	}

	// 演示方法调用
	fmt.Println("\n=== 动态方法调用演示 ===")
	if calc, ok := obj.(*Calculator); ok {
		// 调用Add方法
		if results, err := CallMethod(calc, "Add", 10.5, 20.3); err == nil {
			fmt.Printf("Add(10.5, 20.3) = %v\n", results[0].Interface())
		}

		// 调用Multiply方法
		if results, err := CallMethod(calc, "Multiply", 5.0, 6.0); err == nil {
			fmt.Printf("Multiply(5.0, 6.0) = %v\n", results[0].Interface())
		}

		// 调用GetResult方法
		if results, err := CallMethod(calc, "GetResult"); err == nil {
			fmt.Printf("GetResult() = %v\n", results[0].Interface())
		}
	}
}

// 切片和映射反射

// SliceInfo 分析切片信息
func SliceInfo(slice interface{}) {
	fmt.Printf("\n=== 切片信息分析 ===\n")

	v := reflect.ValueOf(slice)
	t := reflect.TypeOf(slice)

	if t.Kind() != reflect.Slice {
		fmt.Printf("不是切片类型: %s\n", t.Kind())
		return
	}

	fmt.Printf("切片类型: %s\n", t)
	fmt.Printf("元素类型: %s\n", t.Elem())
	fmt.Printf("长度: %d\n", v.Len())
	fmt.Printf("容量: %d\n", v.Cap())

	// 遍历元素
	fmt.Println("元素:")
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("  [%d]: %v (类型: %s)\n", i, elem.Interface(), elem.Type())
	}
}

// MapInfo 分析映射信息
func MapInfo(m interface{}) {
	fmt.Printf("\n=== 映射信息分析 ===\n")

	v := reflect.ValueOf(m)
	t := reflect.TypeOf(m)

	if t.Kind() != reflect.Map {
		fmt.Printf("不是映射类型: %s\n", t.Kind())
		return
	}

	fmt.Printf("映射类型: %s\n", t)
	fmt.Printf("键类型: %s\n", t.Key())
	fmt.Printf("值类型: %s\n", t.Elem())
	fmt.Printf("长度: %d\n", v.Len())

	// 遍历键值对
	fmt.Println("键值对:")
	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		fmt.Printf("  %v: %v\n", key.Interface(), value.Interface())
	}
}

// 动态创建和修改

// CreateStruct 动态创建结构体
func CreateStruct() interface{} {
	fmt.Printf("\n=== 动态创建结构体 ===\n")

	// 定义字段
	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
			Tag:  `json:"name"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
			Tag:  `json:"age"`,
		},
		{
			Name: "Email",
			Type: reflect.TypeOf(""),
			Tag:  `json:"email"`,
		},
	}

	// 创建结构体类型
	structType := reflect.StructOf(fields)
	fmt.Printf("创建的结构体类型: %s\n", structType)

	// 创建实例
	structValue := reflect.New(structType).Elem()

	// 设置字段值
	structValue.FieldByName("Name").SetString("动态用户")
	structValue.FieldByName("Age").SetInt(25)
	structValue.FieldByName("Email").SetString("dynamic@example.com")

	instance := structValue.Interface()
	fmt.Printf("创建的实例: %+v\n", instance)

	return instance
}

// ModifyStruct 修改结构体字段
func ModifyStruct(s interface{}) {
	fmt.Printf("\n=== 修改结构体字段 ===\n")

	v := reflect.ValueOf(s)

	// 如果是指针，获取指向的值
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Printf("不是结构体类型: %s\n", v.Kind())
		return
	}

	fmt.Printf("修改前: %+v\n", v.Interface())

	// 修改字段
	if field := v.FieldByName("Name"); field.IsValid() && field.CanSet() {
		if field.Kind() == reflect.String {
			field.SetString("修改后的名称")
		}
	}

	if field := v.FieldByName("Age"); field.IsValid() && field.CanSet() {
		if field.Kind() == reflect.Int {
			field.SetInt(30)
		}
	}

	fmt.Printf("修改后: %+v\n", v.Interface())
}

// 实用工具函数

// DeepEqual 深度比较两个值
func DeepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// IsNil 检查值是否为nil
func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

// Clone 深度克隆值
func Clone(src interface{}) interface{} {
	srcValue := reflect.ValueOf(src)
	return cloneValue(srcValue).Interface()
}

func cloneValue(src reflect.Value) reflect.Value {
	switch src.Kind() {
	case reflect.Ptr:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		clone := reflect.New(src.Type().Elem())
		clone.Elem().Set(cloneValue(src.Elem()))
		return clone

	case reflect.Slice:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		clone := reflect.MakeSlice(src.Type(), src.Len(), src.Cap())
		for i := 0; i < src.Len(); i++ {
			clone.Index(i).Set(cloneValue(src.Index(i)))
		}
		return clone

	case reflect.Map:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		clone := reflect.MakeMap(src.Type())
		for _, key := range src.MapKeys() {
			clone.SetMapIndex(key, cloneValue(src.MapIndex(key)))
		}
		return clone

	case reflect.Struct:
		clone := reflect.New(src.Type()).Elem()
		for i := 0; i < src.NumField(); i++ {
			if clone.Field(i).CanSet() {
				clone.Field(i).Set(cloneValue(src.Field(i)))
			}
		}
		return clone

	default:
		return src
	}
}

// ToMap 将结构体转换为映射
func ToMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	// 如果是指针，获取指向的值
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if field.IsExported() && fieldValue.CanInterface() {
			// 使用json标签作为键名，如果没有则使用字段名
			key := field.Name
			if jsonTag := field.Tag.Get("json"); jsonTag != "" {
				if parts := strings.Split(jsonTag, ","); len(parts) > 0 && parts[0] != "" {
					key = parts[0]
				}
			}

			result[key] = fieldValue.Interface()
		}
	}

	return result
}

// FromMap 从映射创建结构体
func FromMap(m map[string]interface{}, target interface{}) error {
	v := reflect.ValueOf(target)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target必须是结构体指针")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if !field.IsExported() || !fieldValue.CanSet() {
			continue
		}

		// 获取映射中的键名
		key := field.Name
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			if parts := strings.Split(jsonTag, ","); len(parts) > 0 && parts[0] != "" {
				key = parts[0]
			}
		}

		if value, exists := m[key]; exists {
			if err := setFieldValue(fieldValue, value); err != nil {
				return fmt.Errorf("设置字段 %s 失败: %v", field.Name, err)
			}
		}
	}

	return nil
}

func setFieldValue(field reflect.Value, value interface{}) error {
	valueReflect := reflect.ValueOf(value)

	// 类型匹配直接设置
	if valueReflect.Type().AssignableTo(field.Type()) {
		field.Set(valueReflect)
		return nil
	}

	// 类型转换
	switch field.Kind() {
	case reflect.String:
		field.SetString(fmt.Sprintf("%v", value))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if i, err := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64); err == nil {
			field.SetInt(i)
		} else {
			return err
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if u, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 64); err == nil {
			field.SetUint(u)
		} else {
			return err
		}
	case reflect.Float32, reflect.Float64:
		if f, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64); err == nil {
			field.SetFloat(f)
		} else {
			return err
		}
	case reflect.Bool:
		if b, err := strconv.ParseBool(fmt.Sprintf("%v", value)); err == nil {
			field.SetBool(b)
		} else {
			return err
		}
	default:
		return fmt.Errorf("不支持的类型转换: %s", field.Kind())
	}

	return nil
}

// 反射示例函数
func ReflectionExamples() {
	fmt.Println("=== Go反射示例 ===")

	// 基本类型反射
	fmt.Println("\n🔹 基本类型反射")
	TypeInfo(42)
	TypeInfo("Hello, Reflection!")
	TypeInfo(3.14)

	var ptr *int = new(int)
	*ptr = 100
	TypeInfo(ptr)

	// 结构体反射
	user := &User{
		ID:       1,
		Name:     "张三",
		Email:    "zhangsan@example.com",
		Age:      25,
		IsActive: true,
		Tags:     []string{"developer", "golang"},
		Profile: &Profile{
			Bio:     "Go语言开发者",
			Website: "https://example.com",
		},
		Created: time.Now(),
	}

	StructInfo(user)

	// 方法反射
	calc := &Calculator{}
	MethodInfo(calc)

	// 切片反射
	numbers := []int{1, 2, 3, 4, 5}
	SliceInfo(numbers)

	// 映射反射
	userMap := map[string]interface{}{
		"name": "李四",
		"age":  30,
		"city": "北京",
	}
	MapInfo(userMap)

	// 动态创建和修改
	dynamicStruct := CreateStruct()
	ModifyStruct(dynamicStruct)

	// 实用工具演示
	fmt.Println("\n🔹 实用工具演示")

	// 深度比较
	user1 := User{ID: 1, Name: "张三"}
	user2 := User{ID: 1, Name: "张三"}
	user3 := User{ID: 2, Name: "李四"}

	fmt.Printf("user1 == user2: %t\n", DeepEqual(user1, user2))
	fmt.Printf("user1 == user3: %t\n", DeepEqual(user1, user3))

	// nil检查
	var nilPtr *User
	var validPtr = &user1
	fmt.Printf("nilPtr是否为nil: %t\n", IsNil(nilPtr))
	fmt.Printf("validPtr是否为nil: %t\n", IsNil(validPtr))

	// 结构体转映射
	userMap2 := ToMap(user1)
	fmt.Printf("结构体转映射: %+v\n", userMap2)

	// 映射转结构体
	var newUser User
	data := map[string]interface{}{
		"id":   2,
		"name": "王五",
		"age":  28,
	}
	if err := FromMap(data, &newUser); err == nil {
		fmt.Printf("映射转结构体: %+v\n", newUser)
	}

	// 克隆
	clonedUser := Clone(user1).(User)
	fmt.Printf("原始用户: %+v\n", user1)
	fmt.Printf("克隆用户: %+v\n", clonedUser)
}
