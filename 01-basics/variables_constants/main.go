package main

import (
	"fmt"
)

func main() {
	// 变量声明与初始化的不同方式

	// 1. 使用 var 关键字声明变量，明确指定类型
	var name string = "张三"
	var age int = 30
	var isActive bool = true

	// 2. 使用 var 关键字声明变量，类型推断
	var country = "中国"
	var population = 1400000000
	var hasOcean = true

	// 3. 短变量声明 (仅在函数内部可用)
	city := "北京"
	cityPopulation := 21500000
	isSafe := true

	// 4. 声明多个变量
	var (
		province    string  = "广东"
		area        float64 = 179800.0
		isPopulated bool    = true
	)

	// 5. 多变量同时赋值
	var a, b, c int = 1, 2, 3
	x, y, z := 10, "hello", true

	// 6. 变量值交换
	i, j := 5, 10
	fmt.Printf("交换前: i = %d, j = %d\n", i, j)
	i, j = j, i
	fmt.Printf("交换后: i = %d, j = %d\n", i, j)

	// 常量声明

	// 1. 使用 const 关键字声明单个常量
	const Pi = 3.14159
	const MaxConnections = 100
	const AppName = "Golang例子"

	// 2. 声明多个常量
	const (
		StatusOK       = 200
		StatusCreated  = 201
		StatusAccepted = 202
	)

	// 3. iota 常量生成器
	const (
		Monday    = iota + 1 // 1
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)

	const (
		_           = iota             // 忽略第一个值
		KB ByteSize = 1 << (10 * iota) // 1 << 10 = 1024
		MB                             // 1 << 20
		GB                             // 1 << 30
		TB                             // 1 << 40
		PB                             // 1 << 50
	)

	// 打印变量值
	fmt.Println("---- 变量示例 ----")
	fmt.Printf("姓名: %s, 年龄: %d, 是否活跃: %t\n", name, age, isActive)
	fmt.Printf("国家: %s, 人口: %d, 有海洋: %t\n", country, population, hasOcean)
	fmt.Printf("城市: %s, 城市人口: %d, 是否安全: %t\n", city, cityPopulation, isSafe)
	fmt.Printf("省份: %s, 面积: %.2f 平方公里, 是否有人口: %t\n", province, area, isPopulated)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	fmt.Printf("x = %d, y = %s, z = %t\n", x, y, z)

	// 打印常量值
	fmt.Println("\n---- 常量示例 ----")
	fmt.Printf("Pi = %.5f\n", Pi)
	fmt.Printf("最大连接数 = %d\n", MaxConnections)
	fmt.Printf("应用名称 = %s\n", AppName)

	fmt.Printf("HTTP 状态码 - OK: %d, Created: %d, Accepted: %d\n",
		StatusOK, StatusCreated, StatusAccepted)

	fmt.Printf("星期几 - 周一: %d, 周二: %d, 周三: %d, 周四: %d, 周五: %d, 周六: %d, 周日: %d\n",
		Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	fmt.Printf("存储单位 - KB: %d bytes, MB: %d bytes, GB: %d bytes\n", KB, MB, GB)
}

// ByteSize 是一个自定义类型，用于表示字节大小
type ByteSize int
