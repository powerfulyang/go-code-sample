package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 基本错误处理示例
func BasicErrorHandling() {
	fmt.Println("=== 基本错误处理示例 ===")

	// 调用可能出错的函数
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	// 除零错误
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 0 = %.2f\n", result)
	}

	// 字符串转换
	num, err := parseNumber("123")
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
	} else {
		fmt.Printf("解析结果: %d\n", num)
	}

	num, err = parseNumber("abc")
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
	} else {
		fmt.Printf("解析结果: %d\n", num)
	}
}

// 除法函数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 解析数字函数
func parseNumber(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("无法解析 '%s' 为数字: %w", s, err)
	}
	return num, nil
}

// 自定义错误类型
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("验证失败 - 字段: %s, 值: %v, 消息: %s", e.Field, e.Value, e.Message)
}

// 用户验证示例
type User struct {
	Name  string
	Email string
	Age   int
}

func validateUser(user User) error {
	trimmedName := strings.TrimSpace(user.Name)
	if trimmedName == "" {
		return ValidationError{
			Field:   "Name",
			Value:   user.Name,
			Message: "姓名不能为空",
		}
	}

	if utf8.RuneCountInString(trimmedName) < 1 {
		return ValidationError{
			Field:   "Name",
			Value:   user.Name,
			Message: "姓名至少需要1个字符",
		}
	}

	if !strings.Contains(user.Email, "@") {
		return ValidationError{
			Field:   "Email",
			Value:   user.Email,
			Message: "邮箱格式不正确",
		}
	}

	if user.Age < 0 || user.Age > 150 {
		return ValidationError{
			Field:   "Age",
			Value:   user.Age,
			Message: "年龄必须在0-150之间",
		}
	}

	return nil
}

// 错误包装示例
func processFile(filename string) error {
	content, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("处理文件失败: %w", err)
	}

	err = validateContent(content)
	if err != nil {
		return fmt.Errorf("验证文件内容失败: %w", err)
	}

	err = saveProcessedContent(content)
	if err != nil {
		return fmt.Errorf("保存处理后的内容失败: %w", err)
	}

	return nil
}

// 模拟文件读取
func readFile(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("文件名不能为空")
	}
	if !strings.HasSuffix(filename, ".txt") {
		return "", errors.New("只支持 .txt 文件")
	}
	// 模拟读取成功
	return "文件内容", nil
}

// 模拟内容验证
func validateContent(content string) error {
	if len(content) == 0 {
		return errors.New("文件内容为空")
	}
	if len(content) > 1000 {
		return errors.New("文件内容过长")
	}
	return nil
}

// 模拟保存内容
func saveProcessedContent(content string) error {
	if strings.Contains(content, "error") {
		return errors.New("内容包含错误关键字")
	}
	return nil
}

// 多个错误处理
type MultiError struct {
	Errors []error
}

func (me MultiError) Error() string {
	var messages []string
	for _, err := range me.Errors {
		messages = append(messages, err.Error())
	}
	return fmt.Sprintf("多个错误: [%s]", strings.Join(messages, "; "))
}

func (me *MultiError) Add(err error) {
	if err != nil {
		me.Errors = append(me.Errors, err)
	}
}

func (me MultiError) HasErrors() bool {
	return len(me.Errors) > 0
}

// 批量验证用户
func validateUsers(users []User) error {
	var multiErr MultiError

	for i, user := range users {
		if err := validateUser(user); err != nil {
			wrappedErr := fmt.Errorf("用户 %d 验证失败: %w", i+1, err)
			multiErr.Add(wrappedErr)
		}
	}

	if multiErr.HasErrors() {
		return multiErr
	}
	return nil
}

// 错误恢复示例
func safeOperation() (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("操作发生panic: %v", r)
		}
	}()

	// 模拟可能panic的操作
	riskyOperation()
	return "操作成功", nil
}

func riskyOperation() {
	// 模拟panic
	panic("模拟的panic")
}

// 重试机制示例
func retryOperation(maxRetries int) error {
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		err := unreliableOperation()
		if err == nil {
			fmt.Printf("操作在第 %d 次尝试时成功\n", i+1)
			return nil
		}

		lastErr = err
		fmt.Printf("第 %d 次尝试失败: %v\n", i+1, err)
	}

	return fmt.Errorf("操作在 %d 次尝试后仍然失败，最后错误: %w", maxRetries, lastErr)
}

// 模拟不可靠的操作
func unreliableOperation() error {
	// 模拟50%的失败率
	if len(fmt.Sprintf("%d", 42))%2 == 0 {
		return errors.New("模拟的操作失败")
	}
	return nil
}

// 错误类型检查示例
func handleError(err error) {
	if err == nil {
		fmt.Println("没有错误")
		return
	}

	// 检查特定错误类型
	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("验证错误 - 字段: %s, 消息: %s\n", validationErr.Field, validationErr.Message)
		return
	}

	// 检查错误是否包含特定错误
	if errors.Is(err, strconv.ErrSyntax) {
		fmt.Println("数字解析语法错误")
		return
	}

	// 检查多重错误
	var multiErr MultiError
	if errors.As(err, &multiErr) {
		fmt.Printf("多重错误，共 %d 个错误:\n", len(multiErr.Errors))
		for i, e := range multiErr.Errors {
			fmt.Printf("  %d. %v\n", i+1, e)
		}
		return
	}

	// 默认处理
	fmt.Printf("未知错误: %v\n", err)
}

// 实际应用示例
func UserManagementExample() {
	fmt.Println("\n=== 用户管理示例 ===")

	users := []User{
		{Name: "张三", Email: "zhangsan@example.com", Age: 25},
		{Name: "", Email: "invalid-email", Age: -5},
		{Name: "李四", Email: "lisi@example.com", Age: 30},
		{Name: "王", Email: "wangwu@example.com", Age: 200},
	}

	fmt.Println("批量验证用户:")
	err := validateUsers(users)
	handleError(err)

	fmt.Println("\n单个用户验证:")
	for i, user := range users {
		fmt.Printf("验证用户 %d: %+v\n", i+1, user)
		err := validateUser(user)
		if err != nil {
			handleError(err)
		} else {
			fmt.Println("  验证通过")
		}
	}
}

// 文件处理示例
func FileProcessingExample() {
	fmt.Println("\n=== 文件处理示例 ===")

	files := []string{"document.txt", "", "data.csv", "content.txt"}

	for _, filename := range files {
		fmt.Printf("处理文件: '%s'\n", filename)
		err := processFile(filename)
		if err != nil {
			fmt.Printf("  处理失败: %v\n", err)

			// 解包错误链
			fmt.Println("  错误链:")
			for err != nil {
				fmt.Printf("    - %v\n", err)
				err = errors.Unwrap(err)
			}
		} else {
			fmt.Println("  处理成功")
		}
		fmt.Println()
	}
}

// 错误恢复示例
func ErrorRecoveryExample() {
	fmt.Println("\n=== 错误恢复示例 ===")

	result, err := safeOperation()
	if err != nil {
		fmt.Printf("安全操作失败: %v\n", err)
	} else {
		fmt.Printf("安全操作成功: %s\n", result)
	}

	// 重试示例
	fmt.Println("\n重试机制示例:")
	err = retryOperation(3)
	if err != nil {
		fmt.Printf("重试失败: %v\n", err)
	}
}

// 最佳实践示例
func BestPracticesExample() {
	fmt.Println("\n=== 错误处理最佳实践 ===")

	// 1. 及早返回错误
	fmt.Println("1. 及早返回错误:")
	if err := validateInput(""); err != nil {
		fmt.Printf("   输入验证失败: %v\n", err)
		return
	}

	// 2. 提供上下文信息
	fmt.Println("2. 提供上下文信息:")
	if err := processData("invalid"); err != nil {
		fmt.Printf("   数据处理失败: %v\n", err)
	}

	// 3. 使用哨兵错误
	fmt.Println("3. 使用哨兵错误:")
	if err := checkPermission("guest"); err != nil {
		if errors.Is(err, ErrPermissionDenied) {
			fmt.Println("   权限被拒绝")
		} else {
			fmt.Printf("   其他错误: %v\n", err)
		}
	}

	// 4. 错误分类
	fmt.Println("4. 错误分类:")
	err := performOperation()
	switch {
	case isTemporaryError(err):
		fmt.Println("   临时错误，可以重试")
	case isPermanentError(err):
		fmt.Println("   永久错误，需要人工干预")
	default:
		fmt.Println("   未知错误类型")
	}
}

// 哨兵错误
var (
	ErrPermissionDenied = errors.New("权限被拒绝")
	ErrResourceNotFound = errors.New("资源未找到")
	ErrInvalidInput     = errors.New("输入无效")
)

func validateInput(input string) error {
	if input == "" {
		return ErrInvalidInput
	}
	return nil
}

func processData(data string) error {
	if data == "invalid" {
		return fmt.Errorf("处理数据 '%s' 时发生错误: %w", data, ErrInvalidInput)
	}
	return nil
}

func checkPermission(role string) error {
	if role == "guest" {
		return ErrPermissionDenied
	}
	return nil
}

// 错误类型分类
type TemporaryError struct {
	Msg string
}

func (e TemporaryError) Error() string {
	return e.Msg
}

func (e TemporaryError) Temporary() bool {
	return true
}

type PermanentError struct {
	Msg string
}

func (e PermanentError) Error() string {
	return e.Msg
}

func performOperation() error {
	// 模拟返回不同类型的错误
	return TemporaryError{Msg: "网络连接超时"}
}

func isTemporaryError(err error) bool {
	type temporary interface {
		Temporary() bool
	}
	if t, ok := err.(temporary); ok {
		return t.Temporary()
	}
	return false
}

func isPermanentError(err error) bool {
	_, ok := err.(PermanentError)
	return ok
}
