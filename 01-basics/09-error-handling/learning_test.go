package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 🎓 学习导向的测试 - 通过测试学习Go错误处理

// LearningValidationError 学习用的验证错误类型
type LearningValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *LearningValidationError) Error() string {
	return fmt.Sprintf("验证失败 [%s]: %s (值: %v)", e.Field, e.Message, e.Value)
}

// LearningNetworkError 学习用的网络错误类型
type LearningNetworkError struct {
	Op   string
	Addr string
	Err  error
}

func (e *LearningNetworkError) Error() string {
	return fmt.Sprintf("网络错误 [%s %s]: %v", e.Op, e.Addr, e.Err)
}

func (e *LearningNetworkError) Unwrap() error {
	return e.Err
}

// LearningMultiError 学习用的多错误聚合类型
type LearningMultiError struct {
	Errors []error
}

func (me *LearningMultiError) Error() string {
	if len(me.Errors) == 0 {
		return "无错误"
	}

	var messages []string
	for i, err := range me.Errors {
		messages = append(messages, fmt.Sprintf("错误%d: %v", i+1, err))
	}
	return fmt.Sprintf("多个错误: %s", strings.Join(messages, "; "))
}

func (me *LearningMultiError) Add(err error) {
	if err != nil {
		me.Errors = append(me.Errors, err)
	}
}

func (me *LearningMultiError) HasErrors() bool {
	return len(me.Errors) > 0
}

// TestLearnBasicErrorHandling 学习基础错误处理
func TestLearnBasicErrorHandling(t *testing.T) {
	t.Log("🎯 学习目标: 掌握Go语言的错误处理机制")
	t.Log("📚 本测试将教您: error接口、错误创建、错误检查")

	t.Run("学习error接口的基本使用", func(t *testing.T) {
		t.Log("📖 知识点: error是Go中的内置接口，只有一个Error()方法")

		// 🔍 探索: 使用errors.New创建错误
		err1 := errors.New("这是一个简单的错误")
		t.Logf("🔍 使用errors.New: %v", err1)
		t.Logf("   错误类型: %T", err1)
		t.Logf("   错误消息: %s", err1.Error())

		// 🔍 探索: 使用fmt.Errorf创建格式化错误
		userID := 12345
		err2 := fmt.Errorf("用户 %d 不存在", userID)
		t.Logf("🔍 使用fmt.Errorf: %v", err2)

		// 🔍 探索: nil错误表示成功
		var err3 error = nil
		t.Logf("🔍 nil错误: %v (是否为nil: %t)", err3, err3 == nil)

		// ✅ 验证错误基础
		if err1 == nil {
			t.Error("❌ errors.New应该返回非nil错误")
		}

		if err1.Error() != "这是一个简单的错误" {
			t.Errorf("❌ 错误消息不匹配: 期望'这是一个简单的错误'，得到'%s'", err1.Error())
		}

		if err3 != nil {
			t.Error("❌ nil错误检查失败")
		}

		t.Log("✅ 很好！您理解了error接口的基本使用")

		// 💡 学习提示
		t.Log("💡 核心概念: error是接口，nil表示无错误")
		t.Log("💡 惯例: 函数的最后一个返回值通常是error")
		t.Log("💡 检查: 总是检查错误，不要忽略")
	})

	t.Run("学习错误处理的最佳实践", func(t *testing.T) {
		t.Log("📖 知识点: Go推荐显式错误处理，而不是异常机制")

		// 🔍 探索: 函数返回错误的模式
		divide := func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, errors.New("除数不能为零")
			}
			return a / b, nil
		}

		// 正确的错误处理方式
		t.Log("🔍 正确的错误处理模式:")

		// 测试正常情况
		result, err := divide(10, 2)
		if err != nil {
			t.Errorf("正常除法不应该有错误: %v", err)
		} else {
			t.Logf("   10 ÷ 2 = %.1f", result)
		}

		// 测试错误情况
		_, err = divide(10, 0)
		if err == nil {
			t.Error("除零应该返回错误")
		} else {
			t.Logf("   除零错误: %v", err)
		}

		// 🔍 探索: 错误处理的不同策略
		strategies := []struct {
			name     string
			strategy func(error) string
		}{
			{
				"忽略错误（不推荐）",
				func(err error) string {
					return "继续执行，忽略错误"
				},
			},
			{
				"记录并返回",
				func(err error) string {
					if err != nil {
						return fmt.Sprintf("记录错误: %v", err)
					}
					return "操作成功"
				},
			},
			{
				"包装错误",
				func(err error) string {
					if err != nil {
						return fmt.Sprintf("操作失败: %v", err)
					}
					return "操作成功"
				},
			},
		}

		t.Log("🔍 不同的错误处理策略:")
		for _, strategy := range strategies {
			result := strategy.strategy(err)
			t.Logf("   %s: %s", strategy.name, result)
		}

		// ✅ 验证错误处理
		if result != 5.0 {
			t.Errorf("❌ 正常除法结果错误: 期望5.0，得到%.1f", result)
		}

		if err == nil {
			t.Error("❌ 除零应该产生错误")
		}

		t.Log("✅ 很好！您理解了错误处理的最佳实践")

		// 💡 学习提示
		t.Log("💡 显式处理: Go要求显式处理每个错误")
		t.Log("💡 早期返回: 遇到错误立即返回，避免深层嵌套")
		t.Log("💡 错误传播: 将错误向上传播给调用者")
	})
}

// TestLearnCustomErrors 学习自定义错误
func TestLearnCustomErrors(t *testing.T) {
	t.Log("🎯 学习目标: 创建和使用自定义错误类型")
	t.Log("📚 本测试将教您: 自定义错误、错误包装、错误类型判断")

	t.Run("学习创建自定义错误类型", func(t *testing.T) {
		t.Log("📖 知识点: 可以创建自定义错误类型来携带更多信息")

		// 创建自定义错误
		validateAge := func(age int) error {
			if age < 0 {
				return &LearningValidationError{
					Field:   "age",
					Value:   age,
					Message: "年龄不能为负数",
				}
			}
			if age > 150 {
				return &LearningValidationError{
					Field:   "age",
					Value:   age,
					Message: "年龄不能超过150岁",
				}
			}
			return nil
		}

		// 测试自定义错误
		testCases := []struct {
			age      int
			hasError bool
		}{
			{25, false},
			{-5, true},
			{200, true},
			{0, false},
		}

		t.Log("🔍 自定义错误测试:")
		for _, tc := range testCases {
			err := validateAge(tc.age)

			if tc.hasError {
				if err == nil {
					t.Errorf("❌ 年龄%d应该产生错误", tc.age)
				} else {
					t.Logf("   年龄%d: %v", tc.age, err)

					// 类型断言获取详细信息
					if ve, ok := err.(*LearningValidationError); ok {
						t.Logf("     字段: %s, 值: %v", ve.Field, ve.Value)
					}
				}
			} else {
				if err != nil {
					t.Errorf("❌ 年龄%d不应该产生错误: %v", tc.age, err)
				} else {
					t.Logf("   年龄%d: 验证通过", tc.age)
				}
			}
		}

		t.Log("✅ 很好！您理解了自定义错误类型")

		// 💡 学习提示
		t.Log("💡 丰富信息: 自定义错误可以携带更多上下文信息")
		t.Log("💡 类型断言: 使用类型断言获取错误的详细信息")
		t.Log("💡 结构化: 自定义错误让错误处理更加结构化")
	})

	t.Run("学习错误包装和解包", func(t *testing.T) {
		t.Log("📖 知识点: Go 1.13引入了错误包装机制")

		// 🔍 探索: 错误包装
		originalErr := errors.New("原始错误")
		wrappedErr := fmt.Errorf("包装错误: %w", originalErr)
		doubleWrappedErr := fmt.Errorf("二次包装: %w", wrappedErr)

		t.Log("🔍 错误包装链:")
		t.Logf("   原始错误: %v", originalErr)
		t.Logf("   包装错误: %v", wrappedErr)
		t.Logf("   二次包装: %v", doubleWrappedErr)

		// 🔍 探索: 错误解包
		t.Log("🔍 错误解包测试:")

		// 使用errors.Is检查错误链
		if errors.Is(doubleWrappedErr, originalErr) {
			t.Log("   ✅ errors.Is: 在错误链中找到了原始错误")
		} else {
			t.Error("   ❌ errors.Is: 应该能在错误链中找到原始错误")
		}

		// 使用errors.Unwrap逐层解包
		t.Log("   逐层解包:")
		current := doubleWrappedErr
		level := 0
		for current != nil {
			t.Logf("     层级%d: %v", level, current)
			current = errors.Unwrap(current)
			level++
			if level > 5 { // 防止无限循环
				break
			}
		}

		// ✅ 验证错误包装
		if !errors.Is(wrappedErr, originalErr) {
			t.Error("❌ 包装错误应该包含原始错误")
		}

		if !errors.Is(doubleWrappedErr, originalErr) {
			t.Error("❌ 二次包装错误应该包含原始错误")
		}

		t.Log("✅ 很好！您理解了错误包装机制")

		// 💡 学习提示
		t.Log("💡 错误链: %w动词创建错误链，保留原始错误")
		t.Log("💡 errors.Is: 检查错误链中是否包含特定错误")
		t.Log("💡 errors.Unwrap: 逐层解包错误")
	})
}

// TestLearnErrorPatterns 学习错误处理模式
func TestLearnErrorPatterns(t *testing.T) {
	t.Log("🎯 学习目标: 掌握常用的错误处理模式")
	t.Log("📚 本测试将教您: 哨兵错误、错误类型、错误行为")

	t.Run("学习哨兵错误模式", func(t *testing.T) {
		t.Log("📖 知识点: 哨兵错误是预定义的错误值，用于表示特定条件")

		// 🔍 探索: 定义哨兵错误
		var (
			ErrNotFound     = errors.New("资源未找到")
			ErrUnauthorized = errors.New("未授权访问")
			ErrInvalidInput = errors.New("输入无效")
		)

		// 模拟函数返回哨兵错误
		getUserByID := func(id int) (string, error) {
			switch id {
			case 1:
				return "张三", nil
			case 2:
				return "", ErrUnauthorized
			case 3:
				return "", ErrNotFound
			default:
				return "", ErrInvalidInput
			}
		}

		// 测试哨兵错误
		testCases := []struct {
			id          int
			expectedErr error
		}{
			{1, nil},
			{2, ErrUnauthorized},
			{3, ErrNotFound},
			{999, ErrInvalidInput},
		}

		t.Log("🔍 哨兵错误测试:")
		for _, tc := range testCases {
			user, err := getUserByID(tc.id)

			if tc.expectedErr == nil {
				if err != nil {
					t.Errorf("❌ ID %d 不应该有错误: %v", tc.id, err)
				} else {
					t.Logf("   ID %d: 用户 %s", tc.id, user)
				}
			} else {
				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("❌ ID %d 错误不匹配: 期望 %v, 得到 %v", tc.id, tc.expectedErr, err)
				} else {
					t.Logf("   ID %d: 错误 %v", tc.id, err)
				}
			}
		}

		t.Log("✅ 很好！您理解了哨兵错误模式")

		// 💡 学习提示
		t.Log("💡 预定义: 哨兵错误是包级别的预定义错误变量")
		t.Log("💡 比较: 使用errors.Is进行错误比较")
		t.Log("💡 语义: 哨兵错误表达特定的业务语义")
	})

	t.Run("学习错误类型模式", func(t *testing.T) {
		t.Log("📖 知识点: 错误类型模式通过类型断言获取错误详情")

		// 模拟网络操作
		connectToServer := func(addr string) error {
			if addr == "" {
				return &LearningNetworkError{
					Op:   "connect",
					Addr: addr,
					Err:  errors.New("地址为空"),
				}
			}
			if addr == "invalid" {
				return &LearningNetworkError{
					Op:   "connect",
					Addr: addr,
					Err:  errors.New("无效地址"),
				}
			}
			return nil
		}

		// 测试错误类型
		testAddrs := []string{"", "invalid", "valid.com"}

		t.Log("🔍 错误类型测试:")
		for _, addr := range testAddrs {
			err := connectToServer(addr)

			if err != nil {
				t.Logf("   地址 '%s': %v", addr, err)

				// 类型断言获取详细信息
				var netErr *LearningNetworkError
				if errors.As(err, &netErr) {
					t.Logf("     操作: %s, 地址: %s, 原因: %v",
						netErr.Op, netErr.Addr, netErr.Err)
				}
			} else {
				t.Logf("   地址 '%s': 连接成功", addr)
			}
		}

		t.Log("✅ 很好！您理解了错误类型模式")

		// 💡 学习提示
		t.Log("💡 类型断言: 使用errors.As进行类型断言")
		t.Log("💡 结构化: 错误类型可以携带结构化信息")
		t.Log("💡 解包: 实现Unwrap()方法支持错误链")
	})
}

// TestLearnErrorRecovery 学习错误恢复
func TestLearnErrorRecovery(t *testing.T) {
	t.Log("🎯 学习目标: 学习错误恢复和重试机制")
	t.Log("📚 本测试将教您: 重试逻辑、降级处理、错误聚合")

	t.Run("学习重试机制", func(t *testing.T) {
		t.Log("📖 知识点: 对于临时性错误，可以实现重试机制")

		// 🔍 探索: 实现重试逻辑
		attempt := 0
		unreliableOperation := func() error {
			attempt++
			if attempt < 3 {
				return fmt.Errorf("临时错误 (尝试 %d)", attempt)
			}
			return nil
		}

		// 重试函数
		retry := func(operation func() error, maxAttempts int) error {
			var lastErr error
			for i := 0; i < maxAttempts; i++ {
				err := operation()
				if err == nil {
					return nil
				}
				lastErr = err
				t.Logf("   尝试 %d 失败: %v", i+1, err)
			}
			return fmt.Errorf("重试 %d 次后仍然失败: %w", maxAttempts, lastErr)
		}

		t.Log("🔍 重试机制测试:")
		err := retry(unreliableOperation, 5)

		if err != nil {
			t.Errorf("❌ 重试应该成功: %v", err)
		} else {
			t.Logf("   重试成功，总共尝试了 %d 次", attempt)
		}

		t.Log("✅ 很好！您理解了重试机制")

		// 💡 学习提示
		t.Log("💡 适用场景: 网络请求、数据库连接等临时性错误")
		t.Log("💡 退避策略: 可以加入指数退避或随机延迟")
		t.Log("💡 错误分类: 区分可重试和不可重试的错误")
	})

	t.Run("学习错误聚合", func(t *testing.T) {
		t.Log("📖 知识点: 当有多个操作时，可以聚合所有错误")

		// 模拟多个操作
		operations := []func() error{
			func() error { return nil },
			func() error { return errors.New("操作2失败") },
			func() error { return nil },
			func() error { return errors.New("操作4失败") },
		}

		// 执行所有操作并聚合错误
		var multiErr LearningMultiError

		t.Log("🔍 错误聚合测试:")
		for i, op := range operations {
			err := op()
			if err != nil {
				t.Logf("   操作%d失败: %v", i+1, err)
				multiErr.Add(err)
			} else {
				t.Logf("   操作%d成功", i+1)
			}
		}

		if multiErr.HasErrors() {
			t.Logf("   聚合错误: %v", &multiErr)
			t.Logf("   错误数量: %d", len(multiErr.Errors))
		}

		// ✅ 验证错误聚合
		if len(multiErr.Errors) != 2 {
			t.Errorf("❌ 应该有2个错误，实际有%d个", len(multiErr.Errors))
		}

		t.Log("✅ 很好！您理解了错误聚合")

		// 💡 学习提示
		t.Log("💡 批量操作: 聚合错误适用于批量操作场景")
		t.Log("💡 完整信息: 保留所有错误信息，便于调试")
		t.Log("💡 结构化: 可以按类型或来源分类错误")
	})
}

// TestLearnPanicAndRecover 学习panic和recover
func TestLearnPanicAndRecover(t *testing.T) {
	t.Log("🎯 学习目标: 理解panic和recover机制")
	t.Log("📚 本测试将教您: panic触发、recover恢复、使用场景")

	t.Run("学习panic的基本使用", func(t *testing.T) {
		t.Log("📖 知识点: panic用于不可恢复的错误，会终止程序")

		// 🔍 探索: 使用recover捕获panic
		safeDivide := func(a, b int) (result int, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic恢复: %v", r)
				}
			}()

			if b == 0 {
				panic("除数不能为零")
			}

			result = a / b
			return
		}

		// 测试正常情况
		result, err := safeDivide(10, 2)
		if err != nil {
			t.Errorf("❌ 正常除法不应该有错误: %v", err)
		} else {
			t.Logf("🔍 正常除法: 10 ÷ 2 = %d", result)
		}

		// 测试panic情况
		_, err = safeDivide(10, 0)
		if err == nil {
			t.Error("❌ 除零应该产生错误")
		} else {
			t.Logf("🔍 panic恢复: %v", err)
		}

		t.Log("✅ 很好！您理解了panic和recover")

		// 💡 学习提示
		t.Log("💡 使用场景: panic用于程序无法继续的严重错误")
		t.Log("💡 恢复机制: recover只能在defer函数中使用")
		t.Log("💡 最佳实践: 优先使用error，谨慎使用panic")
	})
}

// BenchmarkLearnErrorPerformance 学习错误处理性能
func BenchmarkLearnErrorPerformance(b *testing.B) {
	b.Log("🎯 学习目标: 了解不同错误处理方式的性能")

	b.Run("errors.New", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = errors.New("test error")
		}
	})

	b.Run("fmt.Errorf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Errorf("test error %d", i)
		}
	})

	b.Run("自定义错误", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = &LearningValidationError{
				Field:   "test",
				Value:   i,
				Message: "test error",
			}
		}
	})
}

// Example_learnBasicErrorHandling 基础错误处理示例
func Example_learnBasicErrorHandling() {
	// 定义一个可能返回错误的函数
	parseNumber := func(s string) (int, error) {
		num, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("解析数字失败: %w", err)
		}
		return num, nil
	}

	// 正确的错误处理
	if num, err := parseNumber("123"); err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("数字: %d\n", num)
	}

	// 错误情况
	if _, err := parseNumber("abc"); err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	// Output:
	// 数字: 123
	// 错误: 解析数字失败: strconv.Atoi: parsing "abc": invalid syntax
}
