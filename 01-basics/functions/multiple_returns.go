package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// 1. 返回多个值的函数
func getPlayerInfo() (string, int, bool) {
	return "张三", 25, true
}

// 2. 带命名返回值的多返回值函数
func getCircleInfo(radius float64) (area, circumference float64) {
	area = math.Pi * radius * radius
	circumference = 2 * math.Pi * radius
	return // 裸返回，会返回area和circumference的值
}

// 3. 常见场景：返回结果和错误
func divideWithError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 4. 返回可忽略的多个值
func parseFullName(fullName string) (firstName, lastName, middleName string) {
	parts := strings.Split(fullName, " ")

	if len(parts) == 1 {
		return parts[0], "", ""
	} else if len(parts) == 2 {
		return parts[0], parts[1], ""
	} else {
		return parts[0], parts[len(parts)-1], strings.Join(parts[1:len(parts)-1], " ")
	}
}

// 5. 使用多返回值解析复杂数据
func analyzeText(text string) (wordCount, lineCount int, containsNumber bool) {
	lines := strings.Split(text, "\n")
	lineCount = len(lines)

	for _, line := range lines {
		words := strings.Fields(line)
		wordCount += len(words)

		for _, word := range words {
			if strings.ContainsAny(word, "0123456789") {
				containsNumber = true
			}
		}
	}

	return
}

// 6. 多返回值可以简化代码
func findMinMax(numbers []int) (min, max int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return
}

func MultipleReturnsExample() {
	fmt.Println("=== 多返回值函数示例 ===")

	// 1. 基本多返回值示例
	fmt.Println("\n基本多返回值:")
	name, age, active := getPlayerInfo()
	fmt.Printf("姓名: %s, 年龄: %d, 状态: %t\n", name, age, active)

	// 2. 命名返回值示例
	fmt.Println("\n命名返回值:")
	area, circumference := getCircleInfo(5)
	fmt.Printf("半径为5的圆 - 面积: %.2f, 周长: %.2f\n", area, circumference)

	// 3. 返回结果和错误
	fmt.Println("\n返回结果和错误:")

	// 正常情况
	result, err := divideWithError(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// 错误情况
	result, err = divideWithError(10, 0)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// 4. 忽略某些返回值
	fmt.Println("\n忽略某些返回值:")
	firstName, lastName, _ := parseFullName("John William Smith")
	fmt.Printf("名: %s, 姓: %s\n", firstName, lastName)

	// 使用全部返回值
	firstName, lastName, middleName := parseFullName("John William Smith")
	fmt.Printf("名: %s, 中间名: %s, 姓: %s\n", firstName, middleName, lastName)

	// 5. 多返回值解析复杂数据
	fmt.Println("\n解析复杂数据:")
	text := "这是第一行文本。\n这是第二行，包含数字123。\n这是最后一行。"
	words, lines, hasNumbers := analyzeText(text)
	fmt.Printf("文本包含 %d 个单词, %d 行, 包含数字: %t\n", words, lines, hasNumbers)

	// 6. 找出最小值和最大值
	fmt.Println("\n找出最小值和最大值:")
	numbers := []int{7, 2, 9, 4, 6}
	min, max := findMinMax(numbers)
	fmt.Printf("数组: %v\n", numbers)
	fmt.Printf("最小值: %d, 最大值: %d\n", min, max)
}
