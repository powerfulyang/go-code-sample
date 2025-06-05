package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// 基本goroutine示例
func BasicGoroutines() {
	fmt.Println("=== 基本Goroutine示例 ===")

	// 普通函数调用
	fmt.Println("开始执行...")
	sayHello("同步调用")

	// goroutine调用
	go sayHello("异步调用1")
	go sayHello("异步调用2")
	go sayHello("异步调用3")

	// 等待goroutines完成
	time.Sleep(2 * time.Second)
	fmt.Println("主程序结束")
}

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("[%s] Hello %d\n", name, i+1)
		time.Sleep(500 * time.Millisecond)
	}
}

// 匿名goroutine示例
func AnonymousGoroutines() {
	fmt.Println("\n=== 匿名Goroutine示例 ===")

	// 匿名函数goroutine
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("匿名goroutine: %d\n", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// 带参数的匿名函数goroutine
	message := "Hello from closure"
	go func(msg string) {
		fmt.Printf("参数goroutine: %s\n", msg)
	}(message)

	// 闭包goroutine
	counter := 0
	for i := 0; i < 3; i++ {
		go func() {
			counter++ // 注意：这里有竞态条件
			fmt.Printf("闭包goroutine: counter = %d\n", counter)
		}()
	}

	time.Sleep(1 * time.Second)
}

// WaitGroup示例
func WaitGroupExample() {
	fmt.Println("\n=== WaitGroup示例 ===")

	var wg sync.WaitGroup
	workers := 5

	for i := 0; i < workers; i++ {
		wg.Add(1) // 增加等待计数
		go worker(i, &wg)
	}

	fmt.Println("等待所有worker完成...")
	wg.Wait() // 等待所有goroutine完成
	fmt.Println("所有worker已完成")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 完成时减少计数

	fmt.Printf("Worker %d 开始工作\n", id)

	// 模拟工作
	workTime := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(workTime)

	fmt.Printf("Worker %d 完成工作 (耗时: %v)\n", id, workTime)
}

// 竞态条件示例
func RaceConditionExample() {
	fmt.Println("\n=== 竞态条件示例 ===")

	// 不安全的计数器
	fmt.Println("🔸 不安全的计数器:")
	unsafeCounter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeCounter++ // 竞态条件
		}()
	}

	wg.Wait()
	fmt.Printf("不安全计数器结果: %d (期望: 1000)\n", unsafeCounter)

	// 使用互斥锁的安全计数器
	fmt.Println("\n🔸 使用互斥锁的安全计数器:")
	safeCounter := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			safeCounter++
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("安全计数器结果: %d (期望: 1000)\n", safeCounter)
}

// 互斥锁示例
func MutexExample() {
	fmt.Println("\n=== 互斥锁示例 ===")

	type SafeCounter struct {
		mu    sync.Mutex
		value int
	}

	counter := &SafeCounter{}

	// 增加方法
	increment := func() {
		counter.mu.Lock()
		defer counter.mu.Unlock()
		counter.value++
	}

	// 获取值方法
	getValue := func() int {
		counter.mu.Lock()
		defer counter.mu.Unlock()
		return counter.value
	}

	var wg sync.WaitGroup

	// 启动多个goroutine进行增加操作
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", getValue())
}

// 读写锁示例
func RWMutexExample() {
	fmt.Println("\n=== 读写锁示例 ===")

	type SafeMap struct {
		mu   sync.RWMutex
		data map[string]int
	}

	safeMap := &SafeMap{
		data: make(map[string]int),
	}

	// 写操作
	set := func(key string, value int) {
		safeMap.mu.Lock()
		defer safeMap.mu.Unlock()
		safeMap.data[key] = value
		fmt.Printf("设置 %s = %d\n", key, value)
	}

	// 读操作
	get := func(key string) (int, bool) {
		safeMap.mu.RLock()
		defer safeMap.mu.RUnlock()
		value, exists := safeMap.data[key]
		return value, exists
	}

	var wg sync.WaitGroup

	// 启动写goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			set(key, i*10)
		}(i)
	}

	// 启动读goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i%5)
			if value, exists := get(key); exists {
				fmt.Printf("读取 %s = %d\n", key, value)
			} else {
				fmt.Printf("键 %s 不存在\n", key)
			}
		}(i)
	}

	wg.Wait()
}

// Once示例
func OnceExample() {
	fmt.Println("\n=== sync.Once示例 ===")

	var once sync.Once
	var config string

	loadConfig := func() {
		fmt.Println("加载配置...")
		time.Sleep(100 * time.Millisecond) // 模拟耗时操作
		config = "配置已加载"
		fmt.Println("配置加载完成")
	}

	getConfig := func() string {
		once.Do(loadConfig) // 只会执行一次
		return config
	}

	var wg sync.WaitGroup

	// 多个goroutine同时获取配置
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cfg := getConfig()
			fmt.Printf("Goroutine %d 获取配置: %s\n", id, cfg)
		}(i)
	}

	wg.Wait()
}

// 工作池示例
func WorkerPoolExample() {
	fmt.Println("\n=== 工作池示例 ===")

	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动workers
	for w := 1; w <= numWorkers; w++ {
		go workerFunc(w, jobs, results)
	}

	// 发送工作
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("结果: %d\n", result)
	}
}

func workerFunc(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d 开始处理任务 %d\n", id, j)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		result := j * 2
		fmt.Printf("Worker %d 完成任务 %d, 结果: %d\n", id, j, result)
		results <- result
	}
}

// Context示例
func ContextExample() {
	fmt.Println("\n=== Context示例 ===")

	// 带超时的context
	fmt.Println("🔸 超时Context:")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("任务完成")
		case <-ctx.Done():
			fmt.Printf("任务被取消: %v\n", ctx.Err())
		}
	}()

	time.Sleep(3 * time.Second)

	// 手动取消的context
	fmt.Println("\n🔸 手动取消Context:")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx2.Done():
				fmt.Printf("任务被手动取消: %v\n", ctx2.Err())
				return
			default:
				fmt.Println("任务运行中...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1500 * time.Millisecond)
	cancel2() // 手动取消
	time.Sleep(500 * time.Millisecond)
}

// 生产者消费者示例
func ProducerConsumerExample() {
	fmt.Println("\n=== 生产者消费者示例 ===")

	buffer := make(chan string, 5) // 缓冲通道
	var wg sync.WaitGroup

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(buffer)

		for i := 1; i <= 10; i++ {
			item := fmt.Sprintf("商品-%d", i)
			buffer <- item
			fmt.Printf("生产者: 生产了 %s\n", item)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("生产者: 生产完成")
	}()

	// 消费者
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		go func(consumerID int) {
			defer wg.Done()
			for item := range buffer {
				fmt.Printf("消费者%d: 消费了 %s\n", consumerID, item)
				time.Sleep(300 * time.Millisecond)
			}
			fmt.Printf("消费者%d: 消费完成\n", consumerID)
		}(i)
	}

	wg.Wait()
}

// 扇入扇出模式
func FanInFanOutExample() {
	fmt.Println("\n=== 扇入扇出模式示例 ===")

	// 输入通道
	input := make(chan int)

	// 扇出：创建多个worker处理输入
	const numWorkers = 3
	workers := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workers[i] = fanOutWorker(i, input)
	}

	// 扇入：合并所有worker的输出
	output := fanIn(workers...)

	// 发送数据
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()

	// 接收结果
	for result := range output {
		fmt.Printf("最终结果: %d\n", result)
	}
}

func fanOutWorker(id int, input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range input {
			result := n * n // 计算平方
			fmt.Printf("Worker %d: %d² = %d\n", id, n, result)
			output <- result
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return output
}

func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				output <- n
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// 运行时信息
func RuntimeInfo() {
	fmt.Println("\n=== 运行时信息 ===")

	fmt.Printf("CPU核心数: %d\n", runtime.NumCPU())
	fmt.Printf("当前Goroutine数: %d\n", runtime.NumGoroutine())
	fmt.Printf("Go版本: %s\n", runtime.Version())
	fmt.Printf("操作系统: %s\n", runtime.GOOS)
	fmt.Printf("架构: %s\n", runtime.GOARCH)

	// 启动一些goroutines
	for i := 0; i < 10; i++ {
		go func(id int) {
			time.Sleep(1 * time.Second)
		}(i)
	}

	fmt.Printf("启动10个goroutine后的数量: %d\n", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	fmt.Printf("2秒后的Goroutine数: %d\n", runtime.NumGoroutine())
}

// 管道模式示例
func PipelineExample() {
	fmt.Println("\n=== 管道模式示例 ===")

	// 阶段1：生成数字
	numbers := generate(1, 2, 3, 4, 5)

	// 阶段2：计算平方
	squares := square(numbers)

	// 阶段3：过滤偶数
	evens := filter(squares)

	// 消费结果
	for result := range evens {
		fmt.Printf("管道结果: %d\n", result)
	}
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}
