package controlflow

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestConditionalStatements(t *testing.T) {
	t.Run("BasicIfStatement", func(t *testing.T) {
		age := 20
		var result string

		if age >= 18 {
			result = "成年"
		} else {
			result = "未成年"
		}

		t.Logf("年龄 %d，状态: %s", age, result)

		if result != "成年" {
			t.Errorf("期望 '成年', 实际 '%s'", result)
		}
	})

	t.Run("IfElseChain", func(t *testing.T) {
		scores := []int{95, 85, 75, 65, 55}
		expected := []string{"优秀", "良好", "中等", "及格", "不及格"}

		for i, score := range scores {
			var grade string
			if score >= 90 {
				grade = "优秀"
			} else if score >= 80 {
				grade = "良好"
			} else if score >= 70 {
				grade = "中等"
			} else if score >= 60 {
				grade = "及格"
			} else {
				grade = "不及格"
			}

			t.Logf("分数 %d，等级: %s", score, grade)

			if grade != expected[i] {
				t.Errorf("分数 %d: 期望 '%s', 实际 '%s'", score, expected[i], grade)
			}
		}
	})

	t.Run("IfWithInitialization", func(t *testing.T) {
		// 带初始化的 if 语句
		if num := 42; num > 40 {
			t.Logf("数字 %d 大于 40", num)
		} else {
			t.Logf("数字 %d 不大于 40", num)
		}

		// 验证作用域
		// num 在这里不可访问
	})

	t.Run("NestedConditions", func(t *testing.T) {
		age := 25
		hasLicense := true
		var canDrive bool

		if age >= 18 {
			if hasLicense {
				canDrive = true
			} else {
				canDrive = false
			}
		} else {
			canDrive = false
		}

		t.Logf("年龄: %d, 有驾照: %t, 可以开车: %t", age, hasLicense, canDrive)

		if !canDrive {
			t.Error("应该可以开车")
		}
	})
}

func TestSwitchStatements(t *testing.T) {
	t.Run("BasicSwitch", func(t *testing.T) {
		day := time.Monday
		var dayType string

		switch day {
		case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
			dayType = "工作日"
		case time.Saturday, time.Sunday:
			dayType = "周末"
		default:
			dayType = "未知"
		}

		t.Logf("星期: %v, 类型: %s", day, dayType)

		if dayType != "工作日" {
			t.Errorf("期望 '工作日', 实际 '%s'", dayType)
		}
	})

	t.Run("SwitchWithoutExpression", func(t *testing.T) {
		hour := 14
		var period string

		switch {
		case hour < 6:
			period = "凌晨"
		case hour < 12:
			period = "上午"
		case hour < 18:
			period = "下午"
		default:
			period = "晚上"
		}

		t.Logf("时间 %d 点，时段: %s", hour, period)

		if period != "下午" {
			t.Errorf("期望 '下午', 实际 '%s'", period)
		}
	})

	t.Run("SwitchWithFallthrough", func(t *testing.T) {
		grade := 'B'
		var description []string

		switch grade {
		case 'A':
			description = append(description, "优秀")
			fallthrough
		case 'B':
			description = append(description, "良好")
			fallthrough
		case 'C':
			description = append(description, "中等")
			fallthrough
		default:
			description = append(description, "需要努力")
		}

		t.Logf("成绩 %c，描述: %v", grade, description)

		expected := []string{"良好", "中等", "需要努力"}
		if len(description) != len(expected) {
			t.Errorf("描述长度不匹配: 期望 %d, 实际 %d", len(expected), len(description))
		}
	})

	t.Run("TypeSwitch", func(t *testing.T) {
		var values []interface{} = []interface{}{42, "hello", 3.14, true}

		for i, v := range values {
			var typeDesc string
			switch v.(type) {
			case int:
				typeDesc = "整数"
			case string:
				typeDesc = "字符串"
			case float64:
				typeDesc = "浮点数"
			case bool:
				typeDesc = "布尔值"
			default:
				typeDesc = "未知类型"
			}

			t.Logf("值 %d: %v (类型: %s)", i, v, typeDesc)
		}
	})
}

func TestForLoops(t *testing.T) {
	t.Run("BasicForLoop", func(t *testing.T) {
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}

		t.Logf("1 到 10 的和: %d", sum)

		expected := 55 // 1+2+...+10 = 55
		if sum != expected {
			t.Errorf("期望 %d, 实际 %d", expected, sum)
		}
	})

	t.Run("WhileStyleLoop", func(t *testing.T) {
		count := 0
		i := 1
		for i <= 5 {
			count++
			i++
		}

		t.Logf("循环次数: %d", count)

		if count != 5 {
			t.Errorf("期望循环 5 次, 实际 %d 次", count)
		}
	})

	t.Run("InfiniteLoopWithBreak", func(t *testing.T) {
		count := 0
		for {
			count++
			if count >= 3 {
				break
			}
		}

		t.Logf("无限循环计数: %d", count)

		if count != 3 {
			t.Errorf("期望计数 3, 实际 %d", count)
		}
	})

	t.Run("LoopWithContinue", func(t *testing.T) {
		var odds []int
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				continue
			}
			odds = append(odds, i)
		}

		t.Logf("奇数: %v", odds)

		expected := []int{1, 3, 5, 7, 9}
		if len(odds) != len(expected) {
			t.Errorf("奇数数量不匹配: 期望 %d, 实际 %d", len(expected), len(odds))
		}
	})
}

func TestRangeLoops(t *testing.T) {
	t.Run("RangeOverArray", func(t *testing.T) {
		numbers := [5]int{10, 20, 30, 40, 50}
		sum := 0

		for _, value := range numbers {
			sum += value
		}

		t.Logf("数组元素和: %d", sum)

		expected := 150
		if sum != expected {
			t.Errorf("期望 %d, 实际 %d", expected, sum)
		}
	})

	t.Run("RangeOverSlice", func(t *testing.T) {
		fruits := []string{"苹果", "香蕉", "橙子"}
		var result []string

		for i, fruit := range fruits {
			result = append(result, fmt.Sprintf("%d:%s", i, fruit))
		}

		t.Logf("水果列表: %v", result)

		if len(result) != 3 {
			t.Errorf("期望 3 个元素, 实际 %d 个", len(result))
		}
	})

	t.Run("RangeOverMap", func(t *testing.T) {
		scores := map[string]int{
			"张三": 95,
			"李四": 87,
			"王五": 92,
		}

		total := 0
		count := 0

		for name, score := range scores {
			t.Logf("学生: %s, 分数: %d", name, score)
			total += score
			count++
		}

		average := float64(total) / float64(count)
		t.Logf("平均分: %.2f", average)

		if count != 3 {
			t.Errorf("期望 3 个学生, 实际 %d 个", count)
		}
	})

	t.Run("RangeOverString", func(t *testing.T) {
		text := "Go语言"
		var chars []rune

		for _, char := range text {
			chars = append(chars, char)
		}

		t.Logf("字符串 '%s' 的字符: %v", text, chars)

		if len(chars) != 4 { // G, o, 语, 言
			t.Errorf("期望 4 个字符, 实际 %d 个", len(chars))
		}
	})
}

func TestNestedLoops(t *testing.T) {
	t.Run("MultiplicationTable", func(t *testing.T) {
		// 生成部分乘法表
		var table [][]int
		for i := 1; i <= 3; i++ {
			var row []int
			for j := 1; j <= 3; j++ {
				row = append(row, i*j)
			}
			table = append(table, row)
		}

		t.Logf("3x3 乘法表:")
		for i, row := range table {
			t.Logf("  第 %d 行: %v", i+1, row)
		}

		// 验证结果
		if table[0][0] != 1 || table[2][2] != 9 {
			t.Error("乘法表计算错误")
		}
	})

	t.Run("MatrixTraversal", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		sum := 0
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				sum += matrix[i][j]
			}
		}

		t.Logf("矩阵元素和: %d", sum)

		expected := 45 // 1+2+...+9 = 45
		if sum != expected {
			t.Errorf("期望 %d, 实际 %d", expected, sum)
		}
	})
}

func TestLoopControl(t *testing.T) {
	t.Run("LabeledBreak", func(t *testing.T) {
		var results []string

	outer:
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				if i*j > 4 {
					results = append(results, fmt.Sprintf("break at i=%d,j=%d", i, j))
					break outer
				}
				results = append(results, fmt.Sprintf("i=%d,j=%d", i, j))
			}
		}

		t.Logf("标签 break 结果: %v", results)

		// 验证在正确位置跳出
		found := false
		for _, result := range results {
			if result == "break at i=2,j=3" {
				found = true
				break
			}
		}
		if !found {
			t.Error("没有在预期位置跳出循环")
		}
	})

	t.Run("LabeledContinue", func(t *testing.T) {
		var results []string

	outer:
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				if j == 2 {
					results = append(results, fmt.Sprintf("continue at i=%d,j=%d", i, j))
					continue outer
				}
				results = append(results, fmt.Sprintf("i=%d,j=%d", i, j))
			}
		}

		t.Logf("标签 continue 结果: %v", results)

		// 验证跳过了 j=2 和 j=3 的情况
		for _, result := range results {
			if strings.Contains(result, "j=3") {
				t.Error("不应该执行 j=3 的情况")
			}
		}
	})
}

func TestPracticalExamples(t *testing.T) {
	t.Run("GradeStatistics", func(t *testing.T) {
		scores := []int{95, 87, 92, 78, 85, 90, 88}

		// 计算统计信息
		total := 0
		highest := scores[0]
		lowest := scores[0]
		passCount := 0

		for _, score := range scores {
			total += score
			if score > highest {
				highest = score
			}
			if score < lowest {
				lowest = score
			}
			if score >= 60 {
				passCount++
			}
		}

		average := float64(total) / float64(len(scores))
		passRate := float64(passCount) / float64(len(scores)) * 100

		t.Logf("成绩统计:")
		t.Logf("  总分: %d", total)
		t.Logf("  平均分: %.2f", average)
		t.Logf("  最高分: %d", highest)
		t.Logf("  最低分: %d", lowest)
		t.Logf("  及格人数: %d", passCount)
		t.Logf("  及格率: %.1f%%", passRate)

		// 验证计算结果
		if highest != 95 {
			t.Errorf("最高分错误: 期望 95, 实际 %d", highest)
		}
		if lowest != 78 {
			t.Errorf("最低分错误: 期望 78, 实际 %d", lowest)
		}
		if passCount != 7 {
			t.Errorf("及格人数错误: 期望 7, 实际 %d", passCount)
		}
	})

	t.Run("NumberClassification", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var evens, odds []int

		for _, num := range numbers {
			if num%2 == 0 {
				evens = append(evens, num)
			} else {
				odds = append(odds, num)
			}
		}

		t.Logf("偶数: %v", evens)
		t.Logf("奇数: %v", odds)

		// 验证分类结果
		if len(evens) != 5 || len(odds) != 5 {
			t.Errorf("分类错误: 偶数 %d 个, 奇数 %d 个", len(evens), len(odds))
		}
	})

	t.Run("SearchAndFilter", func(t *testing.T) {
		names := []string{"张三", "李四", "王五", "赵六", "钱七"}
		target := "王五"
		var found bool
		var index int

		// 搜索
		for i, name := range names {
			if name == target {
				found = true
				index = i
				break
			}
		}

		t.Logf("搜索 '%s': 找到=%t, 索引=%d", target, found, index)

		// 过滤（包含"三"或"五"的名字）
		var filtered []string
		for _, name := range names {
			if strings.Contains(name, "三") || strings.Contains(name, "五") {
				filtered = append(filtered, name)
			}
		}

		t.Logf("过滤结果: %v", filtered)

		// 验证结果
		if !found || index != 2 {
			t.Errorf("搜索失败: 期望找到索引 2, 实际 found=%t, index=%d", found, index)
		}
		if len(filtered) != 2 {
			t.Errorf("过滤结果错误: 期望 2 个, 实际 %d 个", len(filtered))
		}
	})
}
