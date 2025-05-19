package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 自定义错误类型
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

// 实现error接口
func (e ValidationError) Error() string {
	return fmt.Sprintf("验证错误 - 字段: %s, 值: %v, 消息: %s", e.Field, e.Value, e.Message)
}

// 另一个自定义错误类型
type MathError struct {
	Op     string
	Values []float64
	Err    string
}

func (e MathError) Error() string {
	return fmt.Sprintf("数学运算错误 - 操作: %s, 值: %v, 错误: %s", e.Op, e.Values, e.Err)
}

func main() {
	fmt.Println("=== Go 自定义错误示例 ===")

	// 使用errors.New创建简单错误
	fmt.Println("\n--- errors.New 创建错误 ---")
	err1 := errors.New("这是一个简单的错误")
	fmt.Printf("错误: %v\n", err1)
	fmt.Printf("错误类型: %T\n", err1)

	// 使用fmt.Errorf创建格式化错误
	fmt.Println("\n--- fmt.Errorf 创建格式化错误 ---")
	username := "admin"
	err2 := fmt.Errorf("用户 %s 不存在", username)
	fmt.Printf("错误: %v\n", err2)

	age := -5
	err3 := fmt.Errorf("无效的年龄: %d，年龄必须大于0", age)
	fmt.Printf("错误: %v\n", err3)

	// 函数返回错误
	fmt.Println("\n--- 函数返回错误 ---")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// 自定义错误类型
	fmt.Println("\n--- 自定义错误类型 ---")
	err = validateAge(-5)
	if err != nil {
		fmt.Printf("验证错误: %v\n", err)
		// 类型断言获取具体错误信息
		if validationErr, ok := err.(ValidationError); ok {
			fmt.Printf("错误字段: %s\n", validationErr.Field)
			fmt.Printf("错误值: %v\n", validationErr.Value)
			fmt.Printf("错误消息: %s\n", validationErr.Message)
		}
	}

	err = validateEmail("invalid-email")
	if err != nil {
		fmt.Printf("验证错误: %v\n", err)
	}

	// 数学运算错误
	fmt.Println("\n--- 数学运算错误 ---")
	result, err = sqrt(-4)
	if err != nil {
		fmt.Printf("数学错误: %v\n", err)
		if mathErr, ok := err.(MathError); ok {
			fmt.Printf("操作: %s\n", mathErr.Op)
			fmt.Printf("值: %v\n", mathErr.Values)
		}
	}

	// 错误处理模式
	fmt.Println("\n--- 错误处理模式 ---")

	// 模式1：立即返回错误
	data, err := processData("valid_data")
	if err != nil {
		fmt.Printf("处理数据失败: %v\n", err)
		return
	}
	fmt.Printf("处理结果: %s\n", data)

	// 模式2：记录错误但继续执行
	err = saveToFile("data.txt", "some data")
	if err != nil {
		fmt.Printf("保存文件失败: %v，使用默认存储\n", err)
		// 继续执行其他逻辑
	}

	// 模式3：重试机制
	fmt.Println("\n--- 重试机制 ---")
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err = unreliableOperation()
		if err == nil {
			fmt.Println("操作成功")
			break
		}
		fmt.Printf("第%d次尝试失败: %v\n", i+1, err)
		if i == maxRetries-1 {
			fmt.Println("所有重试都失败了")
		}
	}

	// 错误聚合
	fmt.Println("\n--- 错误聚合 ---")
	errs := validateUser("", "invalid-email", -5)
	if len(errs) > 0 {
		fmt.Println("用户验证失败:")
		for i, err := range errs {
			fmt.Printf("  %d. %v\n", i+1, err)
		}
	}

	// 预定义错误
	fmt.Println("\n--- 预定义错误 ---")
	err = parseNumber("abc")
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
		// 检查是否是特定的预定义错误
		if errors.Is(err, ErrInvalidFormat) {
			fmt.Println("这是一个格式错误")
		}
	}

	// 错误链
	fmt.Println("\n--- 错误链 ---")
	err = processFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("处理文件错误: %v\n", err)
	}
}

// divide 除法函数，演示基本错误返回
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// validateAge 验证年龄，返回自定义错误
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "年龄不能为负数",
		}
	}
	if age > 150 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "年龄不能超过150岁",
		}
	}
	return nil
}

// validateEmail 验证邮箱
func validateEmail(email string) error {
	if email == "" {
		return ValidationError{
			Field:   "email",
			Value:   email,
			Message: "邮箱不能为空",
		}
	}
	if len(email) < 5 || !contains(email, "@") {
		return ValidationError{
			Field:   "email",
			Value:   email,
			Message: "邮箱格式无效",
		}
	}
	return nil
}

// sqrt 计算平方根，演示数学错误
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, MathError{
			Op:     "sqrt",
			Values: []float64{x},
			Err:    "不能计算负数的平方根",
		}
	}
	// 简单的平方根计算（实际应该使用math.Sqrt）
	return x * 0.5, nil // 这只是示例，不是真正的平方根
}

// processData 处理数据
func processData(data string) (string, error) {
	if data == "" {
		return "", fmt.Errorf("数据不能为空")
	}
	if data == "invalid_data" {
		return "", fmt.Errorf("无效的数据格式: %s", data)
	}
	return "processed_" + data, nil
}

// saveToFile 保存文件（模拟）
func saveToFile(filename, data string) error {
	if filename == "" {
		return fmt.Errorf("文件名不能为空")
	}
	// 模拟文件保存失败
	if filename == "readonly.txt" {
		return fmt.Errorf("文件 %s 是只读的", filename)
	}
	return nil
}

// unreliableOperation 不可靠的操作（模拟）
var attemptCount = 0

func unreliableOperation() error {
	attemptCount++
	if attemptCount < 3 {
		return fmt.Errorf("操作失败，尝试次数: %d", attemptCount)
	}
	attemptCount = 0 // 重置计数器
	return nil
}

// validateUser 验证用户信息，返回多个错误
func validateUser(name, email string, age int) []error {
	var errs []error

	if name == "" {
		errs = append(errs, ValidationError{
			Field:   "name",
			Value:   name,
			Message: "姓名不能为空",
		})
	}

	if err := validateEmail(email); err != nil {
		errs = append(errs, err)
	}

	if err := validateAge(age); err != nil {
		errs = append(errs, err)
	}

	return errs
}

// 预定义错误
var (
	ErrInvalidFormat = errors.New("无效的格式")
	ErrNotFound      = errors.New("未找到")
	ErrPermission    = errors.New("权限不足")
)

// parseNumber 解析数字
func parseNumber(s string) error {
	_, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidFormat, s)
	}
	return nil
}

// processFile 处理文件，演示错误链
func processFile(filename string) error {
	err := openFile(filename)
	if err != nil {
		return fmt.Errorf("处理文件 %s 失败: %w", filename, err)
	}
	return nil
}

// openFile 打开文件（模拟）
func openFile(filename string) error {
	if filename == "nonexistent.txt" {
		return fmt.Errorf("文件不存在: %s", filename)
	}
	return nil
}

// contains 检查字符串是否包含子字符串（简单实现）
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s != substr
}
