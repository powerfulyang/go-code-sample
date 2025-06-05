package concurrency

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBasicGoroutines(t *testing.T) {
	t.Run("GoroutineExecution", func(t *testing.T) {
		var counter int64
		var wg sync.WaitGroup

		numGoroutines := 10
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				atomic.AddInt64(&counter, 1)
			}()
		}

		wg.Wait()

		if counter != int64(numGoroutines) {
			t.Errorf("期望计数器为 %d, 实际为 %d", numGoroutines, counter)
		}

		t.Logf("成功执行了 %d 个goroutine", counter)
	})

	t.Run("GoroutineWithParameters", func(t *testing.T) {
		results := make(chan int, 5)
		var wg sync.WaitGroup

		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				results <- n * n
			}(i)
		}

		wg.Wait()
		close(results)

		var sum int
		for result := range results {
			sum += result
		}

		expected := 1 + 4 + 9 + 16 + 25 // 1² + 2² + 3² + 4² + 5²
		if sum != expected {
			t.Errorf("期望和为 %d, 实际为 %d", expected, sum)
		}

		t.Logf("平方和: %d", sum)
	})
}

func TestWaitGroup(t *testing.T) {
	t.Run("WaitGroupBasic", func(t *testing.T) {
		var wg sync.WaitGroup
		var completed int64

		numWorkers := 5
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				time.Sleep(100 * time.Millisecond) // 模拟工作
				atomic.AddInt64(&completed, 1)
				t.Logf("Worker %d 完成", id)
			}(i)
		}

		wg.Wait()

		if completed != int64(numWorkers) {
			t.Errorf("期望 %d 个worker完成, 实际 %d 个", numWorkers, completed)
		}
	})

	t.Run("WaitGroupWithError", func(t *testing.T) {
		var wg sync.WaitGroup
		errors := make(chan error, 3)

		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				if id == 1 {
					errors <- fmt.Errorf("worker %d 失败", id)
				} else {
					errors <- nil
				}
			}(i)
		}

		wg.Wait()
		close(errors)

		var errorCount int
		for err := range errors {
			if err != nil {
				errorCount++
				t.Logf("错误: %v", err)
			}
		}

		if errorCount != 1 {
			t.Errorf("期望 1 个错误, 实际 %d 个", errorCount)
		}
	})
}

func TestMutex(t *testing.T) {
	t.Run("MutexProtection", func(t *testing.T) {
		var mu sync.Mutex
		var counter int
		var wg sync.WaitGroup

		numGoroutines := 1000
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				mu.Lock()
				counter++
				mu.Unlock()
			}()
		}

		wg.Wait()

		if counter != numGoroutines {
			t.Errorf("期望计数器为 %d, 实际为 %d", numGoroutines, counter)
		}

		t.Logf("安全计数器最终值: %d", counter)
	})

	t.Run("RWMutexPerformance", func(t *testing.T) {
		var rwmu sync.RWMutex
		var data = make(map[string]int)
		var wg sync.WaitGroup

		// 初始化数据
		rwmu.Lock()
		for i := 0; i < 10; i++ {
			data[fmt.Sprintf("key%d", i)] = i
		}
		rwmu.Unlock()

		// 启动多个读goroutine
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for j := 0; j < 100; j++ {
					rwmu.RLock()
					_ = data[fmt.Sprintf("key%d", j%10)]
					rwmu.RUnlock()
				}
			}(i)
		}

		// 启动少量写goroutine
		for i := 0; i < 2; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for j := 0; j < 10; j++ {
					rwmu.Lock()
					data[fmt.Sprintf("key%d", j)] = j * 10
					rwmu.Unlock()
					time.Sleep(10 * time.Millisecond)
				}
			}(i)
		}

		wg.Wait()
		t.Logf("读写锁测试完成，数据长度: %d", len(data))
	})
}

func TestChannels(t *testing.T) {
	t.Run("BasicChannel", func(t *testing.T) {
		ch := make(chan string, 1)

		// 发送
		ch <- "test message"

		// 接收
		msg := <-ch

		if msg != "test message" {
			t.Errorf("期望 'test message', 实际 '%s'", msg)
		}

		t.Logf("通道消息: %s", msg)
	})

	t.Run("BufferedChannel", func(t *testing.T) {
		ch := make(chan int, 3)

		// 发送多个值
		ch <- 1
		ch <- 2
		ch <- 3

		if len(ch) != 3 {
			t.Errorf("期望通道长度为 3, 实际为 %d", len(ch))
		}

		if cap(ch) != 3 {
			t.Errorf("期望通道容量为 3, 实际为 %d", cap(ch))
		}

		// 接收所有值
		var sum int
		for i := 0; i < 3; i++ {
			sum += <-ch
		}

		if sum != 6 {
			t.Errorf("期望和为 6, 实际为 %d", sum)
		}

		t.Logf("缓冲通道测试完成，和: %d", sum)
	})

	t.Run("ChannelClose", func(t *testing.T) {
		ch := make(chan int, 2)

		ch <- 1
		ch <- 2
		close(ch)

		var values []int
		for value := range ch {
			values = append(values, value)
		}

		if len(values) != 2 {
			t.Errorf("期望接收 2 个值, 实际接收 %d 个", len(values))
		}

		// 测试从关闭的通道接收
		value, ok := <-ch
		if ok {
			t.Error("从关闭的通道接收应该返回 ok=false")
		}
		if value != 0 {
			t.Errorf("从关闭的通道接收应该返回零值, 实际 %d", value)
		}

		t.Logf("通道关闭测试完成，接收到的值: %v", values)
	})
}

func TestSelect(t *testing.T) {
	t.Run("SelectBasic", func(t *testing.T) {
		ch1 := make(chan string, 1)
		ch2 := make(chan string, 1)

		ch1 <- "from ch1"
		ch2 <- "from ch2"

		var received []string
		for i := 0; i < 2; i++ {
			select {
			case msg1 := <-ch1:
				received = append(received, msg1)
			case msg2 := <-ch2:
				received = append(received, msg2)
			}
		}

		if len(received) != 2 {
			t.Errorf("期望接收 2 个消息, 实际接收 %d 个", len(received))
		}

		t.Logf("Select接收到: %v", received)
	})

	t.Run("SelectTimeout", func(t *testing.T) {
		ch := make(chan string)

		start := time.Now()
		select {
		case msg := <-ch:
			t.Errorf("不应该接收到消息: %s", msg)
		case <-time.After(100 * time.Millisecond):
			// 超时是预期的
		}

		elapsed := time.Since(start)
		if elapsed < 100*time.Millisecond {
			t.Errorf("超时时间太短: %v", elapsed)
		}

		t.Logf("超时测试完成，耗时: %v", elapsed)
	})

	t.Run("SelectDefault", func(t *testing.T) {
		ch := make(chan string)

		select {
		case msg := <-ch:
			t.Errorf("不应该接收到消息: %s", msg)
		default:
			t.Log("执行了default分支")
		}
	})
}

func TestContext(t *testing.T) {
	t.Run("ContextTimeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		start := time.Now()
		<-ctx.Done()
		elapsed := time.Since(start)

		if elapsed < 100*time.Millisecond {
			t.Errorf("超时时间太短: %v", elapsed)
		}

		if ctx.Err() != context.DeadlineExceeded {
			t.Errorf("期望 DeadlineExceeded 错误, 实际: %v", ctx.Err())
		}

		t.Logf("Context超时测试完成，耗时: %v", elapsed)
	})

	t.Run("ContextCancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		done := make(chan bool)
		go func() {
			<-ctx.Done()
			done <- true
		}()

		// 取消context
		cancel()

		select {
		case <-done:
			t.Log("Context取消成功")
		case <-time.After(1 * time.Second):
			t.Error("Context取消超时")
		}

		if ctx.Err() != context.Canceled {
			t.Errorf("期望 Canceled 错误, 实际: %v", ctx.Err())
		}
	})

	t.Run("ContextValue", func(t *testing.T) {
		type key string
		const userKey key = "user"

		ctx := context.WithValue(context.Background(), userKey, "testuser")

		value := ctx.Value(userKey)
		if value != "testuser" {
			t.Errorf("期望 'testuser', 实际 '%v'", value)
		}

		// 测试不存在的key
		nonExistent := ctx.Value("nonexistent")
		if nonExistent != nil {
			t.Errorf("不存在的key应该返回nil, 实际: %v", nonExistent)
		}

		t.Logf("Context值: %v", value)
	})
}

func TestWorkerPool(t *testing.T) {
	t.Run("WorkerPoolPattern", func(t *testing.T) {
		const numWorkers = 3
		const numJobs = 10

		jobs := make(chan int, numJobs)
		results := make(chan int, numJobs)

		// 启动workers
		var wg sync.WaitGroup
		for w := 1; w <= numWorkers; w++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for job := range jobs {
					result := job * 2 // 简单的处理
					results <- result
				}
			}(w)
		}

		// 发送工作
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)

		// 等待workers完成
		go func() {
			wg.Wait()
			close(results)
		}()

		// 收集结果
		var sum int
		var count int
		for result := range results {
			sum += result
			count++
		}

		if count != numJobs {
			t.Errorf("期望处理 %d 个工作, 实际处理 %d 个", numJobs, count)
		}

		expectedSum := 0
		for i := 1; i <= numJobs; i++ {
			expectedSum += i * 2
		}

		if sum != expectedSum {
			t.Errorf("期望结果和为 %d, 实际为 %d", expectedSum, sum)
		}

		t.Logf("工作池测试完成，处理了 %d 个工作，结果和: %d", count, sum)
	})
}

func TestRaceCondition(t *testing.T) {
	t.Run("DetectRaceCondition", func(t *testing.T) {
		// 这个测试需要使用 -race 标志运行才能检测到竞态条件
		var counter int
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// 这里故意创建竞态条件
				temp := counter
				time.Sleep(1 * time.Microsecond)
				counter = temp + 1
			}()
		}

		wg.Wait()

		// 由于竞态条件，counter可能不等于100
		t.Logf("竞态条件测试，计数器值: %d (可能不等于100)", counter)
	})

	t.Run("FixedRaceCondition", func(t *testing.T) {
		var counter int64
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				atomic.AddInt64(&counter, 1)
			}()
		}

		wg.Wait()

		if counter != 100 {
			t.Errorf("期望计数器为 100, 实际为 %d", counter)
		}

		t.Logf("修复竞态条件后，计数器值: %d", counter)
	})
}

func TestGoroutineLeaks(t *testing.T) {
	t.Run("GoroutineLeak", func(t *testing.T) {
		initialGoroutines := runtime.NumGoroutine()

		// 创建一个会泄漏的goroutine
		ch := make(chan struct{})
		go func() {
			<-ch // 永远等待
		}()

		time.Sleep(100 * time.Millisecond)

		currentGoroutines := runtime.NumGoroutine()
		if currentGoroutines <= initialGoroutines {
			t.Error("应该检测到goroutine增加")
		}

		t.Logf("初始goroutine数: %d, 当前: %d", initialGoroutines, currentGoroutines)

		// 清理：关闭通道让goroutine退出
		close(ch)
		time.Sleep(100 * time.Millisecond)

		finalGoroutines := runtime.NumGoroutine()
		t.Logf("清理后goroutine数: %d", finalGoroutines)
	})
}

// 基准测试
func BenchmarkGoroutineCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkChannelSendReceive(b *testing.B) {
	ch := make(chan int, 1)
	for i := 0; i < b.N; i++ {
		ch <- i
		<-ch
	}
}

func BenchmarkMutexLock(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

func BenchmarkAtomicAdd(b *testing.B) {
	var counter int64
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&counter, 1)
	}
}
