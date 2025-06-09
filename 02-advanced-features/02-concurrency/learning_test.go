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

// 🎓 学习导向的测试 - 通过测试学习Go并发编程

// TestLearnGoroutines 学习Goroutines基础
func TestLearnGoroutines(t *testing.T) {
	t.Log("🎯 学习目标: 理解Goroutines的基本概念和使用")
	t.Log("📚 本测试将教您: 创建goroutine、并发执行、等待完成")

	t.Run("学习创建和启动Goroutines", func(t *testing.T) {
		t.Log("📖 知识点: 使用go关键字启动goroutine")

		// 🔍 探索: 基本goroutine使用
		var wg sync.WaitGroup
		results := make([]string, 3)

		// 启动多个goroutines
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				// 模拟工作
				time.Sleep(time.Duration(id*10) * time.Millisecond)
				results[id] = fmt.Sprintf("Goroutine %d 完成", id)
				t.Logf("   🔄 Goroutine %d 执行完成", id)
			}(i)
		}

		t.Log("🔍 等待所有goroutines完成...")
		wg.Wait()

		t.Log("🔍 所有goroutines的结果:")
		for i, result := range results {
			t.Logf("   结果[%d]: %s", i, result)
		}

		// ✅ 验证goroutines执行
		for i, result := range results {
			expected := fmt.Sprintf("Goroutine %d 完成", i)
			if result != expected {
				t.Errorf("❌ Goroutine %d 结果错误: 期望 %s, 得到 %s", i, expected, result)
			}
		}
		t.Log("✅ 很好！您理解了如何创建和等待goroutines")

		// 💡 学习提示
		t.Log("💡 重要概念: go关键字创建新的goroutine")
		t.Log("💡 同步工具: sync.WaitGroup用于等待goroutines完成")
		t.Log("💡 闭包陷阱: 注意循环变量的捕获问题")
	})

	t.Run("学习Goroutines的并发特性", func(t *testing.T) {
		t.Log("📖 知识点: Goroutines是轻量级线程，可以创建大量实例")

		// 🔍 探索: 大量goroutines的性能
		numGoroutines := 1000
		var counter int64
		var wg sync.WaitGroup

		start := time.Now()

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				atomic.AddInt64(&counter, 1)
			}()
		}

		wg.Wait()
		duration := time.Since(start)

		t.Logf("🔍 并发性能测试:")
		t.Logf("   创建了 %d 个goroutines", numGoroutines)
		t.Logf("   总耗时: %v", duration)
		t.Logf("   最终计数: %d", counter)
		t.Logf("   当前活跃goroutines: %d", runtime.NumGoroutine())

		// ✅ 验证并发执行
		if counter != int64(numGoroutines) {
			t.Errorf("❌ 并发计数错误: 期望 %d, 得到 %d", numGoroutines, counter)
		} else {
			t.Log("✅ 很好！您理解了goroutines的并发特性")
		}

		// 💡 学习提示
		t.Log("💡 性能优势: Goroutines比系统线程更轻量")
		t.Log("💡 原子操作: 使用atomic包保证并发安全")
		t.Log("💡 资源管理: 注意goroutine的生命周期管理")
	})
}

// TestLearnChannels 学习Channels基础
func TestLearnChannels(t *testing.T) {
	t.Log("🎯 学习目标: 掌握Channels的基本使用")
	t.Log("📚 本测试将教您: 创建channel、发送接收数据、关闭channel")

	t.Run("学习无缓冲Channel", func(t *testing.T) {
		t.Log("📖 知识点: 无缓冲channel是同步的，发送和接收必须同时准备好")

		// 🔍 探索: 无缓冲channel的同步特性
		ch := make(chan string)

		go func() {
			t.Log("   🔄 Goroutine: 准备发送数据...")
			ch <- "Hello from goroutine"
			t.Log("   🔄 Goroutine: 数据发送完成")
		}()

		t.Log("🔍 主goroutine: 等待接收数据...")
		message := <-ch
		t.Logf("🔍 主goroutine: 接收到数据: %s", message)

		// ✅ 验证channel通信
		if message != "Hello from goroutine" {
			t.Errorf("❌ Channel通信错误: 期望 'Hello from goroutine', 得到 '%s'", message)
		} else {
			t.Log("✅ 很好！您理解了无缓冲channel的同步特性")
		}

		// 💡 学习提示
		t.Log("💡 同步特性: 无缓冲channel的发送和接收是同步的")
		t.Log("💡 阻塞行为: 发送方会阻塞直到有接收方准备好")
	})

	t.Run("学习有缓冲Channel", func(t *testing.T) {
		t.Log("📖 知识点: 有缓冲channel可以存储一定数量的值")

		// 🔍 探索: 有缓冲channel的异步特性
		ch := make(chan int, 3) // 缓冲区大小为3

		// 发送数据（不会阻塞，因为有缓冲区）
		ch <- 1
		ch <- 2
		ch <- 3

		t.Logf("🔍 缓冲channel状态:")
		t.Logf("   缓冲区长度: %d", len(ch))
		t.Logf("   缓冲区容量: %d", cap(ch))

		// 接收数据
		values := make([]int, 0, 3)
		for i := 0; i < 3; i++ {
			val := <-ch
			values = append(values, val)
			t.Logf("   接收到: %d, 剩余缓冲: %d", val, len(ch))
		}

		// ✅ 验证缓冲channel
		expected := []int{1, 2, 3}
		for i, val := range values {
			if val != expected[i] {
				t.Errorf("❌ 缓冲channel数据错误: 位置%d期望%d, 得到%d", i, expected[i], val)
			}
		}
		t.Log("✅ 很好！您理解了有缓冲channel的特性")

		// 💡 学习提示
		t.Log("💡 异步特性: 有缓冲channel在缓冲区未满时不会阻塞发送")
		t.Log("💡 容量管理: 合理设置缓冲区大小可以提高性能")
	})

	t.Run("学习Channel的关闭", func(t *testing.T) {
		t.Log("📖 知识点: 关闭channel表示不再发送数据")

		// 🔍 探索: channel关闭和range遍历
		ch := make(chan int, 5)

		// 发送数据并关闭channel
		go func() {
			for i := 1; i <= 5; i++ {
				ch <- i
				t.Logf("   🔄 发送: %d", i)
			}
			close(ch)
			t.Log("   🔄 Channel已关闭")
		}()

		// 使用range遍历channel
		t.Log("🔍 使用range遍历channel:")
		var received []int
		for val := range ch {
			received = append(received, val)
			t.Logf("   接收: %d", val)
		}

		// 检查channel是否已关闭
		val, ok := <-ch
		t.Logf("🔍 从已关闭channel接收: 值=%d, 是否有效=%t", val, ok)

		// ✅ 验证channel关闭
		if len(received) != 5 {
			t.Errorf("❌ 接收数据数量错误: 期望5, 得到%d", len(received))
		}
		if ok {
			t.Error("❌ 已关闭的channel应该返回ok=false")
		}
		if val != 0 {
			t.Errorf("❌ 已关闭channel的零值应该是0, 得到%d", val)
		}
		t.Log("✅ 很好！您理解了channel的关闭机制")

		// 💡 学习提示
		t.Log("💡 关闭语义: close(ch)表示不再发送数据")
		t.Log("💡 接收检查: val, ok := <-ch 可以检查channel是否关闭")
		t.Log("💡 range遍历: range会在channel关闭时自动退出")
	})
}

// TestLearnChannelPatterns 学习Channel模式
func TestLearnChannelPatterns(t *testing.T) {
	t.Log("🎯 学习目标: 掌握常用的Channel编程模式")
	t.Log("📚 本测试将教您: 生产者-消费者、扇入扇出、超时控制")

	t.Run("学习生产者-消费者模式", func(t *testing.T) {
		t.Log("📖 知识点: 使用channel实现生产者-消费者模式")

		// 🔍 探索: 生产者-消费者模式
		jobs := make(chan int, 10)
		results := make(chan int, 10)

		// 启动消费者
		var wg sync.WaitGroup
		numWorkers := 3

		for w := 1; w <= numWorkers; w++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for job := range jobs {
					result := job * job // 简单的工作：计算平方
					t.Logf("   🔄 Worker %d: %d² = %d", id, job, result)
					results <- result
				}
			}(w)
		}

		// 生产者：发送任务
		go func() {
			for j := 1; j <= 9; j++ {
				jobs <- j
			}
			close(jobs)
		}()

		// 等待所有worker完成并关闭结果channel
		go func() {
			wg.Wait()
			close(results)
		}()

		// 收集结果
		var allResults []int
		for result := range results {
			allResults = append(allResults, result)
		}

		t.Logf("🔍 生产者-消费者结果: %v", allResults)

		// ✅ 验证结果
		if len(allResults) != 9 {
			t.Errorf("❌ 结果数量错误: 期望9, 得到%d", len(allResults))
		} else {
			t.Log("✅ 很好！您理解了生产者-消费者模式")
		}

		// 💡 学习提示
		t.Log("💡 设计模式: 生产者-消费者模式解耦了数据生产和处理")
		t.Log("💡 并发控制: 多个worker可以并行处理任务")
	})

	t.Run("学习Select语句", func(t *testing.T) {
		t.Log("📖 知识点: select语句用于处理多个channel操作")

		// 🔍 探索: select的多路复用
		ch1 := make(chan string)
		ch2 := make(chan string)

		// 启动两个goroutines发送数据
		go func() {
			time.Sleep(100 * time.Millisecond)
			ch1 <- "来自channel 1"
		}()

		go func() {
			time.Sleep(200 * time.Millisecond)
			ch2 <- "来自channel 2"
		}()

		// 使用select接收数据
		received := make([]string, 0, 2)
		for i := 0; i < 2; i++ {
			select {
			case msg1 := <-ch1:
				t.Logf("   📨 从ch1接收: %s", msg1)
				received = append(received, msg1)
			case msg2 := <-ch2:
				t.Logf("   📨 从ch2接收: %s", msg2)
				received = append(received, msg2)
			case <-time.After(300 * time.Millisecond):
				t.Log("   ⏰ 超时")
			}
		}

		// ✅ 验证select
		if len(received) != 2 {
			t.Errorf("❌ Select接收数量错误: 期望2, 得到%d", len(received))
		} else {
			t.Log("✅ 很好！您理解了select语句的使用")
		}

		// 💡 学习提示
		t.Log("💡 多路复用: select可以同时等待多个channel操作")
		t.Log("💡 非阻塞: default case可以实现非阻塞操作")
		t.Log("💡 超时控制: time.After可以实现超时机制")
	})
}

// TestLearnSynchronization 学习同步原语
func TestLearnSynchronization(t *testing.T) {
	t.Log("🎯 学习目标: 掌握Go的同步原语")
	t.Log("📚 本测试将教您: Mutex、RWMutex、Once、Cond")

	t.Run("学习Mutex互斥锁", func(t *testing.T) {
		t.Log("📖 知识点: Mutex用于保护共享资源，防止竞态条件")

		// 🔍 探索: 使用Mutex保护共享变量
		var mu sync.Mutex
		var counter int
		var wg sync.WaitGroup

		numGoroutines := 100
		incrementsPerGoroutine := 100

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < incrementsPerGoroutine; j++ {
					mu.Lock()
					counter++
					mu.Unlock()
				}
			}()
		}

		wg.Wait()

		expected := numGoroutines * incrementsPerGoroutine
		t.Logf("🔍 Mutex保护的计数器:")
		t.Logf("   期望值: %d", expected)
		t.Logf("   实际值: %d", counter)

		// ✅ 验证Mutex
		if counter != expected {
			t.Errorf("❌ Mutex保护失败: 期望%d, 得到%d", expected, counter)
		} else {
			t.Log("✅ 很好！您理解了Mutex的作用")
		}

		// 💡 学习提示
		t.Log("💡 竞态条件: 多个goroutine同时访问共享资源会导致数据竞争")
		t.Log("💡 临界区: Lock()和Unlock()之间的代码是临界区")
	})

	t.Run("学习Context上下文", func(t *testing.T) {
		t.Log("📖 知识点: Context用于传递取消信号和超时控制")

		// 🔍 探索: 使用Context控制goroutine
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		result := make(chan string, 1)

		go func() {
			select {
			case <-time.After(300 * time.Millisecond):
				result <- "工作完成"
			case <-ctx.Done():
				result <- "工作被取消: " + ctx.Err().Error()
			}
		}()

		msg := <-result
		t.Logf("🔍 Context控制结果: %s", msg)

		// ✅ 验证Context
		if !contains(msg, "取消") && !contains(msg, "timeout") {
			t.Errorf("❌ Context超时控制失败: %s", msg)
		} else {
			t.Log("✅ 很好！您理解了Context的超时控制")
		}

		// 💡 学习提示
		t.Log("💡 优雅取消: Context提供了优雅取消goroutine的机制")
		t.Log("💡 传递性: Context可以在调用链中传递取消信号")
	})
}

// TestLearnConcurrencyPatterns 学习并发模式
func TestLearnConcurrencyPatterns(t *testing.T) {
	t.Log("🎯 学习目标: 掌握常用的并发编程模式")
	t.Log("📚 本测试将教您: Pipeline、Fan-in/Fan-out、Worker Pool")

	t.Run("学习Pipeline模式", func(t *testing.T) {
		t.Log("📖 知识点: Pipeline将数据处理分解为多个阶段")

		// 🔍 探索: 三阶段Pipeline
		// 阶段1: 生成数字
		numbers := func() <-chan int {
			out := make(chan int)
			go func() {
				defer close(out)
				for i := 1; i <= 5; i++ {
					out <- i
				}
			}()
			return out
		}

		// 阶段2: 计算平方
		square := func(in <-chan int) <-chan int {
			out := make(chan int)
			go func() {
				defer close(out)
				for n := range in {
					out <- n * n
				}
			}()
			return out
		}

		// 阶段3: 转换为字符串
		toString := func(in <-chan int) <-chan string {
			out := make(chan string)
			go func() {
				defer close(out)
				for n := range in {
					out <- fmt.Sprintf("数字: %d", n)
				}
			}()
			return out
		}

		// 构建pipeline
		pipeline := toString(square(numbers()))

		// 收集结果
		var results []string
		for result := range pipeline {
			results = append(results, result)
			t.Logf("   📊 Pipeline输出: %s", result)
		}

		// ✅ 验证Pipeline
		if len(results) != 5 {
			t.Errorf("❌ Pipeline结果数量错误: 期望5, 得到%d", len(results))
		} else {
			t.Log("✅ 很好！您理解了Pipeline模式")
		}

		// 💡 学习提示
		t.Log("💡 模块化: Pipeline将复杂处理分解为简单阶段")
		t.Log("💡 并发性: 各阶段可以并发执行，提高吞吐量")
	})
}

// 辅助函数
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			findSubstring(s, substr))))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// BenchmarkLearnConcurrencyPerformance 学习并发性能
func BenchmarkLearnConcurrencyPerformance(b *testing.B) {
	b.Log("🎯 学习目标: 了解并发编程的性能特征")

	b.Run("顺序vs并发处理", func(b *testing.B) {
		work := func() {
			// 模拟CPU密集型工作
			sum := 0
			for i := 0; i < 1000; i++ {
				sum += i
			}
		}

		b.Run("顺序处理", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 100; j++ {
					work()
				}
			}
		})

		b.Run("并发处理", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var wg sync.WaitGroup
				for j := 0; j < 100; j++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						work()
					}()
				}
				wg.Wait()
			}
		})
	})
}

// Example_learnBasicGoroutine 基础Goroutine示例
func Example_learnBasicGoroutine() {
	// 创建一个channel用于通信
	done := make(chan bool)

	// 启动一个goroutine
	go func() {
		fmt.Println("Goroutine正在运行")
		done <- true
	}()

	// 等待goroutine完成
	<-done
	fmt.Println("程序结束")

	// Output:
	// Goroutine正在运行
	// 程序结束
}
