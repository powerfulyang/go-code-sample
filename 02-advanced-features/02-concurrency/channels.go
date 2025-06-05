package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 基本通道示例
func BasicChannels() {
	fmt.Println("=== 基本通道示例 ===")

	// 创建通道
	ch := make(chan string)

	// 发送数据到通道（在goroutine中）
	go func() {
		ch <- "Hello"
		ch <- "World"
		ch <- "from"
		ch <- "Channel"
		close(ch) // 关闭通道
	}()

	// 从通道接收数据
	for message := range ch {
		fmt.Printf("接收到: %s\n", message)
	}
}

// 缓冲通道示例
func BufferedChannels() {
	fmt.Println("\n=== 缓冲通道示例 ===")

	// 创建缓冲通道
	ch := make(chan int, 3)

	// 发送数据（不会阻塞，因为有缓冲）
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Printf("通道长度: %d, 容量: %d\n", len(ch), cap(ch))

	// 接收数据
	fmt.Printf("接收: %d\n", <-ch)
	fmt.Printf("接收: %d\n", <-ch)
	fmt.Printf("接收: %d\n", <-ch)

	// 演示缓冲区满的情况
	fmt.Println("\n🔸 缓冲区满的情况:")
	buffer := make(chan string, 2)

	go func() {
		buffer <- "消息1"
		fmt.Println("发送了消息1")
		buffer <- "消息2"
		fmt.Println("发送了消息2")
		buffer <- "消息3" // 这会阻塞，直到有空间
		fmt.Println("发送了消息3")
		close(buffer)
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("接收: %s\n", <-buffer) // 释放空间
	time.Sleep(1 * time.Second)

	for msg := range buffer {
		fmt.Printf("接收: %s\n", msg)
	}
}

// 通道方向示例
func ChannelDirections() {
	fmt.Println("\n=== 通道方向示例 ===")

	ch := make(chan string, 1)

	// 只发送通道
	go sender(ch)

	// 只接收通道
	receiver(ch)
}

// 只能发送的通道参数
func sender(ch chan<- string) {
	ch <- "来自sender的消息"
	close(ch)
}

// 只能接收的通道参数
func receiver(ch <-chan string) {
	msg := <-ch
	fmt.Printf("receiver收到: %s\n", msg)
}

// select语句示例
func SelectStatement() {
	fmt.Println("\n=== select语句示例 ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// 发送数据到不同通道
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自ch2"
	}()

	// 使用select等待多个通道
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("收到ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("收到ch2: %s\n", msg2)
		}
	}
}

// 非阻塞通道操作
func NonBlockingChannels() {
	fmt.Println("\n=== 非阻塞通道操作 ===")

	messages := make(chan string)
	signals := make(chan bool)

	// 非阻塞接收
	select {
	case msg := <-messages:
		fmt.Printf("收到消息: %s\n", msg)
	default:
		fmt.Println("没有消息可接收")
	}

	// 非阻塞发送
	msg := "Hello"
	select {
	case messages <- msg:
		fmt.Printf("发送了消息: %s\n", msg)
	default:
		fmt.Println("无法发送消息")
	}

	// 多路非阻塞select
	select {
	case msg := <-messages:
		fmt.Printf("收到消息: %s\n", msg)
	case sig := <-signals:
		fmt.Printf("收到信号: %t\n", sig)
	default:
		fmt.Println("没有活动")
	}
}

// 超时处理
func TimeoutHandling() {
	fmt.Println("\n=== 超时处理示例 ===")

	ch := make(chan string, 1)

	// 模拟耗时操作
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "操作完成"
	}()

	// 1秒超时
	select {
	case result := <-ch:
		fmt.Printf("收到结果: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("操作超时")
	}

	// 等待实际完成
	time.Sleep(2 * time.Second)
	select {
	case result := <-ch:
		fmt.Printf("延迟收到结果: %s\n", result)
	default:
		fmt.Println("没有结果")
	}
}

// 定时器和ticker
func TimersAndTickers() {
	fmt.Println("\n=== 定时器和Ticker示例 ===")

	// Timer示例
	fmt.Println("🔸 Timer示例:")
	timer := time.NewTimer(1 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer触发")
	}()

	time.Sleep(1500 * time.Millisecond)

	// Ticker示例
	fmt.Println("\n🔸 Ticker示例:")
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		count := 0
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				count++
				fmt.Printf("Tick %d at %v\n", count, t.Format("15:04:05"))
				if count >= 3 {
					done <- true
				}
			}
		}
	}()

	time.Sleep(2 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker停止")
}

// 通道关闭和检测
func ChannelClosing() {
	fmt.Println("\n=== 通道关闭示例 ===")

	ch := make(chan int, 3)

	// 发送数据
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	// 检测通道是否关闭
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("收到值: %d\n", value)
	}

	// 使用range自动检测关闭
	fmt.Println("\n🔸 使用range:")
	ch2 := make(chan string, 2)
	ch2 <- "第一个"
	ch2 <- "第二个"
	close(ch2)

	for value := range ch2 {
		fmt.Printf("Range收到: %s\n", value)
	}
}

// 工作分发示例
func WorkDistribution() {
	fmt.Println("\n=== 工作分发示例 ===")

	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// 启动workers
	for w := 1; w <= numWorkers; w++ {
		go distributionWorker(w, jobs, results)
	}

	// 发送工作
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Data: fmt.Sprintf("任务-%d", j)}
	}
	close(jobs)

	// 收集结果
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("结果: ID=%d, Output=%s\n", result.JobID, result.Output)
	}
}

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
}

func distributionWorker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d 处理 %s\n", id, job.Data)

		// 模拟工作
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		result := Result{
			JobID:  job.ID,
			Output: fmt.Sprintf("Worker%d处理了%s", id, job.Data),
		}
		results <- result
	}
}

// 速率限制
func RateLimiting() {
	fmt.Println("\n=== 速率限制示例 ===")

	// 基本速率限制
	fmt.Println("🔸 基本速率限制:")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter // 等待限制器
		fmt.Printf("请求 %d 在 %v 处理\n", req, time.Now().Format("15:04:05.000"))
	}

	// 突发限制
	fmt.Println("\n🔸 突发限制:")
	burstyLimiter := make(chan time.Time, 3)

	// 填充突发限制器
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每200ms添加一个新令牌
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			select {
			case burstyLimiter <- t:
			default:
			}
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Printf("突发请求 %d 在 %v 处理\n", req, time.Now().Format("15:04:05.000"))
	}
}

// 通道同步
func ChannelSynchronization() {
	fmt.Println("\n=== 通道同步示例 ===")

	done := make(chan bool, 1)

	go func() {
		fmt.Println("工作开始...")
		time.Sleep(1 * time.Second)
		fmt.Println("工作完成")
		done <- true
	}()

	// 等待工作完成
	<-done
	fmt.Println("主程序继续执行")
}

// 信号量模式
func SemaphorePattern() {
	fmt.Println("\n=== 信号量模式示例 ===")

	// 创建信号量，限制并发数为3
	semaphore := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 获取信号量
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // 释放信号量

			fmt.Printf("任务 %d 开始执行\n", id)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("任务 %d 执行完成\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("所有任务完成")
}

// 通道链
func ChannelChaining() {
	fmt.Println("\n=== 通道链示例 ===")

	// 创建通道链
	input := make(chan int)
	output := chainProcessor(input, 3) // 3个处理阶段

	// 发送数据
	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()

	// 接收处理结果
	for result := range output {
		fmt.Printf("最终结果: %d\n", result)
	}
}

func chainProcessor(input <-chan int, stages int) <-chan int {
	current := input

	for i := 0; i < stages; i++ {
		next := make(chan int)
		go func(in <-chan int, out chan<- int, stage int) {
			defer close(out)
			for value := range in {
				processed := value * (stage + 1) // 简单的处理逻辑
				fmt.Printf("阶段 %d: %d -> %d\n", stage+1, value, processed)
				out <- processed
			}
		}(current, next, i)
		current = next
	}

	return current
}

// 扇入模式
func FanInPattern() {
	fmt.Println("\n=== 扇入模式示例 ===")

	// 创建多个输入通道
	input1 := make(chan string)
	input2 := make(chan string)
	input3 := make(chan string)

	// 合并通道
	output := fanInChannels(input1, input2, input3)

	// 发送数据到不同通道
	go func() {
		defer close(input1)
		for i := 1; i <= 3; i++ {
			input1 <- fmt.Sprintf("通道1-消息%d", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		defer close(input2)
		for i := 1; i <= 3; i++ {
			input2 <- fmt.Sprintf("通道2-消息%d", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		defer close(input3)
		for i := 1; i <= 3; i++ {
			input3 <- fmt.Sprintf("通道3-消息%d", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// 接收合并的消息
	for msg := range output {
		fmt.Printf("扇入收到: %s\n", msg)
	}
}

func fanInChannels(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup

	for i, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string, id int) {
			defer wg.Done()
			for msg := range ch {
				output <- msg
			}
		}(input, i)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
