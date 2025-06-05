// Package math 提供基本的数学计算功能
package math

import (
	"errors"
	"fmt"
	"math"
)

// Calculator 计算器结构体
type Calculator struct {
	history []Operation
}

// Operation 表示一个计算操作
type Operation struct {
	Type   string
	A, B   float64
	Result float64
}

// New 创建一个新的计算器实例
func New() *Calculator {
	return &Calculator{
		history: make([]Operation, 0),
	}
}

// Add 加法运算
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.addToHistory("ADD", a, b, result)
	return result
}

// Subtract 减法运算
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.addToHistory("SUBTRACT", a, b, result)
	return result
}

// Multiply 乘法运算
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.addToHistory("MULTIPLY", a, b, result)
	return result
}

// Divide 除法运算
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	result := a / b
	c.addToHistory("DIVIDE", a, b, result)
	return result, nil
}

// Power 幂运算
func (c *Calculator) Power(base, exponent float64) float64 {
	result := math.Pow(base, exponent)
	c.addToHistory("POWER", base, exponent, result)
	return result
}

// Sqrt 平方根运算
func (c *Calculator) Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("负数不能开平方根")
	}
	result := math.Sqrt(x)
	c.addToHistory("SQRT", x, 0, result)
	return result, nil
}

// Sin 正弦函数
func (c *Calculator) Sin(x float64) float64 {
	result := math.Sin(x)
	c.addToHistory("SIN", x, 0, result)
	return result
}

// Cos 余弦函数
func (c *Calculator) Cos(x float64) float64 {
	result := math.Cos(x)
	c.addToHistory("COS", x, 0, result)
	return result
}

// Tan 正切函数
func (c *Calculator) Tan(x float64) float64 {
	result := math.Tan(x)
	c.addToHistory("TAN", x, 0, result)
	return result
}

// GetHistory 获取计算历史
func (c *Calculator) GetHistory() []Operation {
	// 返回副本以防止外部修改
	history := make([]Operation, len(c.history))
	copy(history, c.history)
	return history
}

// ClearHistory 清空计算历史
func (c *Calculator) ClearHistory() {
	c.history = c.history[:0]
}

// GetLastResult 获取最后一次计算结果
func (c *Calculator) GetLastResult() (float64, error) {
	if len(c.history) == 0 {
		return 0, errors.New("没有计算历史")
	}
	return c.history[len(c.history)-1].Result, nil
}

// PrintHistory 打印计算历史
func (c *Calculator) PrintHistory() {
	if len(c.history) == 0 {
		fmt.Println("没有计算历史")
		return
	}

	fmt.Println("计算历史:")
	for i, op := range c.history {
		switch op.Type {
		case "ADD":
			fmt.Printf("%d. %.2f + %.2f = %.2f\n", i+1, op.A, op.B, op.Result)
		case "SUBTRACT":
			fmt.Printf("%d. %.2f - %.2f = %.2f\n", i+1, op.A, op.B, op.Result)
		case "MULTIPLY":
			fmt.Printf("%d. %.2f × %.2f = %.2f\n", i+1, op.A, op.B, op.Result)
		case "DIVIDE":
			fmt.Printf("%d. %.2f ÷ %.2f = %.2f\n", i+1, op.A, op.B, op.Result)
		case "POWER":
			fmt.Printf("%d. %.2f ^ %.2f = %.2f\n", i+1, op.A, op.B, op.Result)
		case "SQRT":
			fmt.Printf("%d. √%.2f = %.2f\n", i+1, op.A, op.Result)
		case "SIN":
			fmt.Printf("%d. sin(%.2f) = %.2f\n", i+1, op.A, op.Result)
		case "COS":
			fmt.Printf("%d. cos(%.2f) = %.2f\n", i+1, op.A, op.Result)
		case "TAN":
			fmt.Printf("%d. tan(%.2f) = %.2f\n", i+1, op.A, op.Result)
		}
	}
}

// addToHistory 添加操作到历史记录
func (c *Calculator) addToHistory(opType string, a, b, result float64) {
	op := Operation{
		Type:   opType,
		A:      a,
		B:      b,
		Result: result,
	}
	c.history = append(c.history, op)
}

// 包级别的便利函数

// Add 包级别的加法函数
func Add(a, b float64) float64 {
	return a + b
}

// Subtract 包级别的减法函数
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply 包级别的乘法函数
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide 包级别的除法函数
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// Max 返回两个数中的最大值
func Max(a, b float64) float64 {
	return math.Max(a, b)
}

// Min 返回两个数中的最小值
func Min(a, b float64) float64 {
	return math.Min(a, b)
}

// Abs 返回绝对值
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Round 四舍五入到指定小数位
func Round(x float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(x*multiplier) / multiplier
}

// IsEven 判断是否为偶数
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd 判断是否为奇数
func IsOdd(n int) bool {
	return n%2 != 0
}

// Factorial 计算阶乘
func Factorial(n int) (int64, error) {
	if n < 0 {
		return 0, errors.New("负数没有阶乘")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}

	var result int64 = 1
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result, nil
}

// GCD 计算最大公约数
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM 计算最小公倍数
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// IsPrime 判断是否为质数
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fibonacci 计算斐波那契数列的第n项
func Fibonacci(n int) (int64, error) {
	if n < 0 {
		return 0, errors.New("n必须为非负数")
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}

	var a, b int64 = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

// Sum 计算切片中所有数字的和
func Sum(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Average 计算平均值
func Average(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("空切片无法计算平均值")
	}
	return Sum(numbers) / float64(len(numbers)), nil
}

// Median 计算中位数
func Median(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("空切片无法计算中位数")
	}

	// 复制切片以避免修改原始数据
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)

	// 简单的冒泡排序
	for i := 0; i < len(sorted); i++ {
		for j := 0; j < len(sorted)-1-i; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	n := len(sorted)
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2, nil
	}
	return sorted[n/2], nil
}
