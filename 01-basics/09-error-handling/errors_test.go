package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestBasicErrorHandling(t *testing.T) {
	t.Run("DivideFunction", func(t *testing.T) {
		// 正常除法
		result, err := divide(10, 2)
		if err != nil {
			t.Errorf("正常除法不应该有错误: %v", err)
		}
		if result != 5.0 {
			t.Errorf("10 ÷ 2 应该等于 5.0, 实际 %.2f", result)
		}
		t.Logf("10 ÷ 2 = %.2f", result)

		// 除零错误
		result, err = divide(10, 0)
		if err == nil {
			t.Error("除零应该返回错误")
		}
		if result != 0 {
			t.Errorf("除零时结果应该是 0, 实际 %.2f", result)
		}
		t.Logf("除零错误: %v", err)
	})

	t.Run("ParseNumber", func(t *testing.T) {
		// 正常解析
		num, err := parseNumber("123")
		if err != nil {
			t.Errorf("解析 '123' 不应该有错误: %v", err)
		}
		if num != 123 {
			t.Errorf("解析 '123' 应该得到 123, 实际 %d", num)
		}
		t.Logf("解析 '123' = %d", num)

		// 解析错误
		num, err = parseNumber("abc")
		if err == nil {
			t.Error("解析 'abc' 应该返回错误")
		}
		if num != 0 {
			t.Errorf("解析失败时应该返回 0, 实际 %d", num)
		}
		t.Logf("解析错误: %v", err)

		// 检查错误包装
		if !errors.Is(err, strconv.ErrSyntax) {
			t.Error("解析错误应该包装 strconv.ErrSyntax")
		}
	})
}

func TestCustomErrors(t *testing.T) {
	t.Run("ValidationError", func(t *testing.T) {
		err := ValidationError{
			Field:   "Name",
			Value:   "",
			Message: "姓名不能为空",
		}

		expectedMsg := "验证失败 - 字段: Name, 值: , 消息: 姓名不能为空"
		if err.Error() != expectedMsg {
			t.Errorf("错误消息不匹配:\n期望: %s\n实际: %s", expectedMsg, err.Error())
		}
		t.Logf("验证错误: %v", err)
	})

	t.Run("UserValidation", func(t *testing.T) {
		testCases := []struct {
			user        User
			shouldError bool
			errorField  string
		}{
			{
				user:        User{Name: "张三", Email: "zhangsan@example.com", Age: 25},
				shouldError: false,
			},
			{
				user:        User{Name: "", Email: "test@example.com", Age: 25},
				shouldError: true,
				errorField:  "Name",
			},
			{
				user:        User{Name: "李", Email: "test@example.com", Age: 25},
				shouldError: false,
			},
			{
				user:        User{Name: "王五", Email: "invalid-email", Age: 25},
				shouldError: true,
				errorField:  "Email",
			},
			{
				user:        User{Name: "赵六", Email: "test@example.com", Age: -5},
				shouldError: true,
				errorField:  "Age",
			},
			{
				user:        User{Name: "钱七", Email: "test@example.com", Age: 200},
				shouldError: true,
				errorField:  "Age",
			},
		}

		for i, tc := range testCases {
			t.Logf("测试用例 %d: %+v", i+1, tc.user)
			err := validateUser(tc.user)

			if tc.shouldError {
				if err == nil {
					t.Errorf("用例 %d 应该返回错误", i+1)
					continue
				}

				var validationErr ValidationError
				if !errors.As(err, &validationErr) {
					t.Errorf("用例 %d 应该返回 ValidationError, 实际: %T", i+1, err)
					continue
				}

				if validationErr.Field != tc.errorField {
					t.Errorf("用例 %d 错误字段应该是 %s, 实际: %s", i+1, tc.errorField, validationErr.Field)
				}

				t.Logf("  验证错误: %v", err)
			} else {
				if err != nil {
					t.Errorf("用例 %d 不应该返回错误: %v", i+1, err)
				} else {
					t.Log("  验证通过")
				}
			}
		}
	})
}

func TestErrorWrapping(t *testing.T) {
	t.Run("FileProcessing", func(t *testing.T) {
		testCases := []struct {
			filename    string
			shouldError bool
			description string
		}{
			{
				filename:    "document.txt",
				shouldError: false,
				description: "正常文件",
			},
			{
				filename:    "",
				shouldError: true,
				description: "空文件名",
			},
			{
				filename:    "data.csv",
				shouldError: true,
				description: "不支持的文件类型",
			},
		}

		for _, tc := range testCases {
			t.Logf("测试 %s: %s", tc.description, tc.filename)
			err := processFile(tc.filename)

			if tc.shouldError {
				if err == nil {
					t.Errorf("%s 应该返回错误", tc.description)
					continue
				}

				t.Logf("  错误: %v", err)

				// 测试错误解包
				t.Log("  错误链:")
				currentErr := err
				level := 0
				for currentErr != nil {
					t.Logf("    %d. %v", level, currentErr)
					currentErr = errors.Unwrap(currentErr)
					level++
					if level > 10 { // 防止无限循环
						break
					}
				}
			} else {
				if err != nil {
					t.Errorf("%s 不应该返回错误: %v", tc.description, err)
				} else {
					t.Log("  处理成功")
				}
			}
		}
	})
}

func TestMultiError(t *testing.T) {
	t.Run("MultipleErrors", func(t *testing.T) {
		var multiErr MultiError

		// 添加多个错误
		multiErr.Add(errors.New("错误1"))
		multiErr.Add(errors.New("错误2"))
		multiErr.Add(nil) // 应该被忽略
		multiErr.Add(errors.New("错误3"))

		if !multiErr.HasErrors() {
			t.Error("应该有错误")
		}

		if len(multiErr.Errors) != 3 {
			t.Errorf("应该有 3 个错误, 实际 %d 个", len(multiErr.Errors))
		}

		errorMsg := multiErr.Error()
		t.Logf("多重错误消息: %s", errorMsg)

		// 检查错误消息包含所有子错误
		if !strings.Contains(errorMsg, "错误1") ||
			!strings.Contains(errorMsg, "错误2") ||
			!strings.Contains(errorMsg, "错误3") {
			t.Error("错误消息应该包含所有子错误")
		}
	})

	t.Run("BatchUserValidation", func(t *testing.T) {
		users := []User{
			{Name: "张三", Email: "zhangsan@example.com", Age: 25}, // 正确
			{Name: "", Email: "invalid-email", Age: -5},          // 多个错误
			{Name: "李四", Email: "lisi@example.com", Age: 30},     // 正确
			{Name: "王", Email: "wangwu@example.com", Age: 200},   // 多个错误
		}

		err := validateUsers(users)
		if err == nil {
			t.Error("批量验证应该返回错误")
			return
		}

		var multiErr MultiError
		if !errors.As(err, &multiErr) {
			t.Errorf("应该返回 MultiError, 实际: %T", err)
			return
		}

		t.Logf("批量验证错误: %v", err)
		t.Logf("错误数量: %d", len(multiErr.Errors))

		// 应该有2个用户的验证错误
		expectedErrorCount := 2
		if len(multiErr.Errors) != expectedErrorCount {
			t.Errorf("应该有 %d 个用户验证错误, 实际 %d 个", expectedErrorCount, len(multiErr.Errors))
		}
	})
}

func TestErrorRecovery(t *testing.T) {
	t.Run("SafeOperation", func(t *testing.T) {
		result, err := safeOperation()
		if err == nil {
			t.Error("安全操作应该捕获 panic 并返回错误")
		}

		if result != "" {
			t.Errorf("发生错误时结果应该为空, 实际: %s", result)
		}

		t.Logf("捕获的错误: %v", err)

		// 检查错误消息包含 panic 信息
		if !strings.Contains(err.Error(), "panic") {
			t.Error("错误消息应该包含 panic 信息")
		}
	})

	t.Run("RetryMechanism", func(t *testing.T) {
		// 由于 unreliableOperation 的实现是确定性的，
		// 我们知道它会失败，所以测试重试机制
		err := retryOperation(3)
		if err == nil {
			t.Log("重试成功")
		} else {
			t.Logf("重试失败: %v", err)
			// 检查错误消息包含重试次数
			if !strings.Contains(err.Error(), "3 次") {
				t.Error("错误消息应该包含重试次数")
			}
		}
	})
}

func TestErrorTypeChecking(t *testing.T) {
	t.Run("ErrorAsChecking", func(t *testing.T) {
		// 测试 ValidationError
		validationErr := ValidationError{
			Field:   "TestField",
			Value:   "TestValue",
			Message: "测试消息",
		}

		var extractedErr ValidationError
		if !errors.As(validationErr, &extractedErr) {
			t.Error("应该能够提取 ValidationError")
		}

		if extractedErr.Field != "TestField" {
			t.Errorf("提取的字段应该是 'TestField', 实际: %s", extractedErr.Field)
		}

		// 测试包装的错误
		wrappedErr := fmt.Errorf("包装错误: %w", validationErr)
		if !errors.As(wrappedErr, &extractedErr) {
			t.Error("应该能够从包装错误中提取 ValidationError")
		}
	})

	t.Run("ErrorIsChecking", func(t *testing.T) {
		// 测试哨兵错误
		err := ErrPermissionDenied
		if !errors.Is(err, ErrPermissionDenied) {
			t.Error("应该能够识别哨兵错误")
		}

		// 测试包装的哨兵错误
		wrappedErr := fmt.Errorf("操作失败: %w", ErrPermissionDenied)
		if !errors.Is(wrappedErr, ErrPermissionDenied) {
			t.Error("应该能够从包装错误中识别哨兵错误")
		}

		// 测试不同的错误
		if errors.Is(ErrPermissionDenied, ErrResourceNotFound) {
			t.Error("不同的哨兵错误不应该相等")
		}
	})

	t.Run("MultiErrorChecking", func(t *testing.T) {
		var multiErr MultiError
		multiErr.Add(errors.New("错误1"))
		multiErr.Add(ValidationError{Field: "test", Message: "测试"})

		var extractedMultiErr MultiError
		if !errors.As(multiErr, &extractedMultiErr) {
			t.Error("应该能够提取 MultiError")
		}

		if len(extractedMultiErr.Errors) != 2 {
			t.Errorf("提取的 MultiError 应该有 2 个错误, 实际 %d 个", len(extractedMultiErr.Errors))
		}
	})
}

func TestSentinelErrors(t *testing.T) {
	t.Run("SentinelErrorUsage", func(t *testing.T) {
		testCases := []struct {
			input       string
			expectedErr error
		}{
			{"", ErrInvalidInput},
			{"valid", nil},
		}

		for _, tc := range testCases {
			err := validateInput(tc.input)
			if tc.expectedErr == nil {
				if err != nil {
					t.Errorf("输入 '%s' 不应该有错误: %v", tc.input, err)
				}
			} else {
				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("输入 '%s' 应该返回 %v, 实际: %v", tc.input, tc.expectedErr, err)
				}
			}
		}
	})

	t.Run("PermissionCheck", func(t *testing.T) {
		err := checkPermission("guest")
		if !errors.Is(err, ErrPermissionDenied) {
			t.Errorf("guest 用户应该被拒绝访问, 实际错误: %v", err)
		}

		err = checkPermission("admin")
		if err != nil {
			t.Errorf("admin 用户应该有权限, 实际错误: %v", err)
		}
	})
}

func TestErrorClassification(t *testing.T) {
	t.Run("TemporaryError", func(t *testing.T) {
		tempErr := TemporaryError{Msg: "临时错误"}

		if !isTemporaryError(tempErr) {
			t.Error("应该识别为临时错误")
		}

		if isPermanentError(tempErr) {
			t.Error("不应该识别为永久错误")
		}

		t.Logf("临时错误: %v", tempErr)
	})

	t.Run("PermanentError", func(t *testing.T) {
		permErr := PermanentError{Msg: "永久错误"}

		if isTemporaryError(permErr) {
			t.Error("不应该识别为临时错误")
		}

		if !isPermanentError(permErr) {
			t.Error("应该识别为永久错误")
		}

		t.Logf("永久错误: %v", permErr)
	})

	t.Run("UnknownError", func(t *testing.T) {
		unknownErr := errors.New("未知错误")

		if isTemporaryError(unknownErr) {
			t.Error("未知错误不应该识别为临时错误")
		}

		if isPermanentError(unknownErr) {
			t.Error("未知错误不应该识别为永久错误")
		}

		t.Logf("未知错误: %v", unknownErr)
	})
}

// 基准测试
func BenchmarkErrorCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errors.New("测试错误")
	}
}

func BenchmarkErrorWrapping(b *testing.B) {
	baseErr := errors.New("基础错误")
	for i := 0; i < b.N; i++ {
		_ = fmt.Errorf("包装错误: %w", baseErr)
	}
}

func BenchmarkValidationError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ValidationError{
			Field:   "TestField",
			Value:   "TestValue",
			Message: "测试消息",
		}
	}
}
