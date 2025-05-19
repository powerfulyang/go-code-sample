package main

import (
	"fmt"
	"runtime"
	"time"
)

func SwitchCaseExample() {
	fmt.Println("=== switch-case 条件语句示例 ===")

	// 1. 基本switch语句
	day := "星期三"

	fmt.Printf("今天是%s，", day)
	switch day {
	case "星期一":
		fmt.Println("开始新的一周")
	case "星期三":
		fmt.Println("已经一周过半了")
	case "星期五":
		fmt.Println("周末快到了")
	case "星期六", "星期日": // 多个匹配值
		fmt.Println("终于到周末了")
	default: // 默认情况
		fmt.Println("普通的工作日")
	}

	// 2. switch没有表达式时，case中可使用布尔表达式
	fmt.Println("\n没有表达式的switch:")

	hour := time.Now().Hour()

	switch {
	case hour < 6:
		fmt.Println("凌晨时分")
	case hour < 12:
		fmt.Println("上午好")
	case hour < 18:
		fmt.Println("下午好")
	default:
		fmt.Println("晚上好")
	}

	// 3. switch初始化语句
	fmt.Println("\nswitch带初始化语句:")

	switch dayOfWeek := time.Now().Weekday(); dayOfWeek {
	case time.Saturday, time.Sunday:
		fmt.Println("周末愉快!")
	default:
		fmt.Printf("还有%d天到周末\n", time.Saturday-dayOfWeek)
	}

	// 4. fallthrough (继续执行下一个case)
	fmt.Println("\nfallthrough示例:")

	num := 75
	switch {
	case num >= 90:
		fmt.Println("成绩优秀")
	case num >= 70:
		fmt.Println("成绩良好")
		fallthrough
	case num >= 60:
		fmt.Println("及格了")
	default:
		fmt.Println("需要继续努力")
	}

	// 5. 类型switch (用于接口类型判断)
	fmt.Println("\n类型switch示例:")

	var x interface{} = 3.14

	switch v := x.(type) {
	case int:
		fmt.Println("x是整数:", v)
	case float64:
		fmt.Println("x是浮点数:", v)
	case string:
		fmt.Println("x是字符串:", v)
	case bool:
		fmt.Println("x是布尔值:", v)
	default:
		fmt.Printf("x的类型未知: %T\n", v)
	}

	// 6. 使用switch判断操作系统
	fmt.Println("\n当前操作系统:")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("运行在macOS上")
	case "linux":
		fmt.Println("运行在Linux上")
	case "windows":
		fmt.Println("运行在Windows上")
	default:
		fmt.Printf("运行在%s上\n", os)
	}

	// 7. 比较if-else和switch性能
	fmt.Println("\nswitch和if-else对比:")
	score := 85

	// if-else版本
	fmt.Print("使用if-else: ")
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	// switch版本
	fmt.Print("使用switch: ")
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}

	// 8. 复杂条件
	fmt.Println("\n复杂条件判断:")

	age := 25
	switch {
	case age < 13:
		fmt.Println("儿童")
	case age >= 13 && age < 18:
		fmt.Println("青少年")
	case age >= 18 && age < 60:
		fmt.Println("成年人")
	default:
		fmt.Println("老年人")
	}

	// 9. switch中使用break
	fmt.Println("\nswitch中使用break:")

	switch letter := "B"; letter {
	case "A":
		fmt.Println("选择了A")
	case "B":
		fmt.Println("选择了B")
		if true {
			fmt.Println("但我们提前结束了")
			break
		}
		fmt.Println("这行不会被执行")
	case "C":
		fmt.Println("选择了C")
	}
}
