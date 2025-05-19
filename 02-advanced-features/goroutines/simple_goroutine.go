package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go Goroutine 示例 ===")

	// 基本goroutine
	fmt.Println("\n--- 基本 Goroutine ---")

	// 普通函数调用（同步）
	fmt.Println("开始执行...")
	sayHello("同步调用")
	fmt.Println("同步调用完成")

	// goroutine调用（异步）
	fmt.Println("\n启动goroutine...")
	go sayHello("异步调用1")
	go sayHello("异步调用2")
	go sayHello("异步调用3")

	// 等待goroutine执行
	time.Sleep(2 * time.Second)
	fmt.Println("主函数继续执行")

	// 匿名函数goroutine
	fmt.Println("\n--- 匿名函数 Goroutine ---")
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("匿名goroutine: %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 带参数的匿名函数goroutine
	message := "Hello from goroutine"
	go func(msg string) {
		fmt.Printf("参数化匿名goroutine: %s\n", msg)
	}(message)

	// 多个goroutine并发执行
	fmt.Println("\n--- 多个 Goroutine 并发 ---")
	for i := 1; i <= 3; i++ {
		go worker(i)
	}

	// 等待所有goroutine完成
	time.Sleep(3 * time.Second)

	// goroutine与闭包
	fmt.Println("\n--- Goroutine 与闭包 ---")
	for i := 1; i <= 3; i++ {
		// 错误的方式 - 闭包捕获循环变量
		go func() {
			fmt.Printf("错误方式 - goroutine %d\n", i) // 可能打印相同的值
		}()
	}

	time.Sleep(100 * time.Millisecond)

	// 正确的方式 - 传递参数
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("正确方式 - goroutine %d\n", id)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// goroutine的生命周期
	fmt.Println("\n--- Goroutine 生命周期 ---")
	done := make(chan bool)

	go func() {
		fmt.Println("Goroutine 开始执行")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 执行完成")
		done <- true
	}()

	fmt.Println("等待goroutine完成...")
	<-done
	fmt.Println("主函数收到完成信号")

	fmt.Println("\n程序结束")
}

// sayHello 打印问候消息
func sayHello(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: Hello %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

// worker 模拟工作任务
func worker(id int) {
	fmt.Printf("Worker %d 开始工作\n", id)

	// 模拟工作
	for i := 1; i <= 3; i++ {
		fmt.Printf("Worker %d 正在处理任务 %d\n", id, i)
		time.Sleep(time.Duration(id*200) * time.Millisecond)
	}

	fmt.Printf("Worker %d 完成工作\n", id)
}
