package main

import (
	"fmt"
	"math/rand"
	"time"
)

func IfElseExample() {
	fmt.Println("=== if-else 条件语句示例 ===")

	// 简单条件判断
	x := 10

	if x > 5 {
		fmt.Println("x 大于 5")
	}

	// if-else 结构
	y := 3

	if y > 5 {
		fmt.Println("y 大于 5")
	} else {
		fmt.Println("y 不大于 5")
	}

	// if-else if-else 结构
	z := 7

	if z > 10 {
		fmt.Println("z 大于 10")
	} else if z > 5 {
		fmt.Println("z 大于 5 但不大于 10")
	} else {
		fmt.Println("z 不大于 5")
	}

	// 初始化语句和条件判断
	if num := rand.Intn(10); num < 5 {
		fmt.Printf("生成的数字 %d 小于 5\n", num)
	} else {
		fmt.Printf("生成的数字 %d 大于或等于 5\n", num)
	}
	// 注意：num 变量只在 if-else 语句块内可见

	// 嵌套的条件语句
	hour := time.Now().Hour()

	if hour < 12 {
		fmt.Println("现在是上午")
		if hour < 6 {
			fmt.Println("天还没亮呢")
		} else {
			fmt.Println("早上好！")
		}
	} else {
		fmt.Println("现在是下午或晚上")
		if hour < 18 {
			fmt.Println("下午好！")
		} else {
			fmt.Println("晚上好！")
		}
	}

	// 多条件组合
	age := 25
	hasID := true

	if age >= 18 && hasID {
		fmt.Println("可以购买酒精饮料")
	} else {
		fmt.Println("不能购买酒精饮料")
	}

	temperature := 15
	raining := false

	if temperature < 10 || raining {
		fmt.Println("穿外套")
	} else {
		fmt.Println("不需要穿外套")
	}

	// 条件的优化技巧

	// 1. 提前返回，减少嵌套
	printWeatherAdvice(25, true)
	printWeatherAdvice(5, false)

	// 2. 简化布尔判断
	loggedIn := true
	fmt.Println("用户登录状态:", loggedIn)

	// 3. 避免代码重复
	score := 85
	printGrade(score)
}

// 提前返回示例
func printWeatherAdvice(temperature int, raining bool) {
	fmt.Printf("\n天气情况: %d°C, 下雨: %t\n", temperature, raining)

	// 提前处理特殊情况
	if raining {
		fmt.Println("带伞")
		// 即使下雨，我们还是需要根据温度来判断穿什么
	}

	// 处理温度建议
	if temperature < 0 {
		fmt.Println("非常冷，穿厚外套")
		return
	}

	if temperature < 10 {
		fmt.Println("冷，穿外套")
		return
	}

	if temperature < 20 {
		fmt.Println("凉爽，穿轻薄外套")
		return
	}

	fmt.Println("温暖，穿轻便衣物")
}

// 避免代码重复示例
func printGrade(score int) {
	var grade string

	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("分数: %d, 等级: %s\n", score, grade)
}
