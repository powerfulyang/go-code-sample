package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go 通道示例 ===")

	// 无缓冲通道
	fmt.Println("\n--- 无缓冲通道 ---")
	unbufferedDemo()

	// 有缓冲通道
	fmt.Println("\n--- 有缓冲通道 ---")
	bufferedDemo()

	// 通道方向
	fmt.Println("\n--- 通道方向 ---")
	channelDirectionDemo()

	// 通道关闭
	fmt.Println("\n--- 通道关闭 ---")
	channelCloseDemo()

	// select语句
	fmt.Println("\n--- Select 语句 ---")
	selectDemo()
}

// 无缓冲通道示例
func unbufferedDemo() {
	// 创建无缓冲通道
	ch := make(chan string)

	// 启动goroutine发送数据
	go func() {
		fmt.Println("Goroutine: 准备发送数据")
		ch <- "Hello from goroutine"
		fmt.Println("Goroutine: 数据已发送")
	}()

	// 主goroutine接收数据
	fmt.Println("Main: 准备接收数据")
	message := <-ch
	fmt.Printf("Main: 收到数据: %s\n", message)

	// 演示阻塞特性
	fmt.Println("\n无缓冲通道的阻塞特性:")
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine: 2秒后发送数据")
		ch2 <- 42
	}()

	fmt.Println("Main: 等待数据...")
	value := <-ch2
	fmt.Printf("Main: 收到数据: %d\n", value)
}

// 有缓冲通道示例
func bufferedDemo() {
	// 创建有缓冲通道
	ch := make(chan int, 3)

	// 发送数据（不会阻塞，直到缓冲区满）
	fmt.Println("发送数据到缓冲通道:")
	ch <- 1
	fmt.Println("发送: 1")
	ch <- 2
	fmt.Println("发送: 2")
	ch <- 3
	fmt.Println("发送: 3")

	// 查看通道长度和容量
	fmt.Printf("通道长度: %d, 容量: %d\n", len(ch), cap(ch))

	// 接收数据
	fmt.Println("\n从缓冲通道接收数据:")
	for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Printf("接收: %d\n", value)
		fmt.Printf("通道长度: %d, 容量: %d\n", len(ch), cap(ch))
	}

	// 演示缓冲区满时的阻塞
	fmt.Println("\n缓冲区满时的阻塞:")
	ch2 := make(chan string, 2)

	go func() {
		ch2 <- "first"
		fmt.Println("发送: first")
		ch2 <- "second"
		fmt.Println("发送: second")
		ch2 <- "third" // 这里会阻塞，直到有空间
		fmt.Println("发送: third")
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("接收: %s\n", <-ch2)
	time.Sleep(1 * time.Second)
	fmt.Printf("接收: %s\n", <-ch2)
	time.Sleep(1 * time.Second)
	fmt.Printf("接收: %s\n", <-ch2)
}

// 通道方向示例
func channelDirectionDemo() {
	ch := make(chan int, 1)

	// 只发送通道
	go sender(ch)

	// 只接收通道
	go receiver(ch)

	time.Sleep(1 * time.Second)
}

// 只能发送的通道
func sender(ch chan<- int) {
	fmt.Println("Sender: 发送数据")
	ch <- 100
}

// 只能接收的通道
func receiver(ch <-chan int) {
	value := <-ch
	fmt.Printf("Receiver: 接收到数据 %d\n", value)
}

// 通道关闭示例
func channelCloseDemo() {
	ch := make(chan int, 3)

	// 发送数据
	ch <- 1
	ch <- 2
	ch <- 3

	// 关闭通道
	close(ch)

	// 从关闭的通道接收数据
	fmt.Println("从关闭的通道接收数据:")
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("接收: %d\n", value)
	}

	// 使用range遍历通道
	fmt.Println("\n使用range遍历通道:")
	ch2 := make(chan string, 2)
	ch2 <- "hello"
	ch2 <- "world"
	close(ch2)

	for value := range ch2 {
		fmt.Printf("Range接收: %s\n", value)
	}
}

// select语句示例
func selectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个goroutine
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自通道1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自通道2"
	}()

	// 使用select等待多个通道
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("收到: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("收到: %s\n", msg2)
		}
	}

	// select with default
	fmt.Println("\nSelect with default:")
	ch3 := make(chan int)

	select {
	case value := <-ch3:
		fmt.Printf("收到值: %d\n", value)
	default:
		fmt.Println("没有数据可接收，执行默认分支")
	}

	// 超时处理
	fmt.Println("\n超时处理:")
	ch4 := make(chan string)

	select {
	case msg := <-ch4:
		fmt.Printf("收到消息: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("超时：1秒内没有收到消息")
	}
}
