package performance

import (
	"testing"
	"time"
)

func TestMemoryOptimization(t *testing.T) {
	mo := NewMemoryOptimization()
	
	t.Run("ObjectPool", func(t *testing.T) {
		testData := []byte("test data for object pool")
		
		// 测试对象池处理
		result := mo.ProcessDataWithPool(testData)
		if len(result) != len(testData) {
			t.Errorf("ProcessDataWithPool result length: got %d, want %d", len(result), len(testData))
		}
		
		t.Log("对象池处理测试通过")
	})
	
	t.Run("WithoutPool", func(t *testing.T) {
		testData := []byte("test data without pool")
		
		// 测试不使用对象池处理
		result := mo.ProcessDataWithoutPool(testData)
		if len(result) != len(testData) {
			t.Errorf("ProcessDataWithoutPool result length: got %d, want %d", len(result), len(testData))
		}
		
		t.Log("非对象池处理测试通过")
	})
}

func TestStringOptimization(t *testing.T) {
	so := &StringOptimization{}
	
	t.Run("StringsBuilder", func(t *testing.T) {
		testStrings := []string{"hello", " ", "world", "!"}
		
		result := so.ConcatenateStringsBuilder(testStrings)
		expected := "hello world!"
		
		if result != expected {
			t.Errorf("ConcatenateStringsBuilder: got %s, want %s", result, expected)
		}
		
		t.Log("strings.Builder连接测试通过")
	})
	
	t.Run("NaiveConcatenation", func(t *testing.T) {
		testStrings := []string{"hello", " ", "world", "!"}
		
		result := so.ConcatenateStringsNaive(testStrings)
		expected := "hello world!"
		
		if result != expected {
			t.Errorf("ConcatenateStringsNaive: got %s, want %s", result, expected)
		}
		
		t.Log("朴素连接测试通过")
	})
	
	t.Run("UnsafeConversion", func(t *testing.T) {
		testBytes := []byte("test string")
		
		// 测试unsafe转换
		str := so.BytesToStringUnsafe(testBytes)
		if str != "test string" {
			t.Errorf("BytesToStringUnsafe: got %s, want test string", str)
		}
		
		// 测试反向转换
		bytes := so.StringToBytesUnsafe(str)
		if string(bytes) != str {
			t.Errorf("StringToBytesUnsafe: got %s, want %s", string(bytes), str)
		}
		
		t.Log("unsafe转换测试通过")
	})
}

func TestSliceOptimization(t *testing.T) {
	so := &SliceOptimization{}
	
	t.Run("PreallocationVsNoPreallocation", func(t *testing.T) {
		n := 1000
		
		// 测试预分配
		slice1 := so.AppendWithPreallocation(n)
		if len(slice1) != n {
			t.Errorf("AppendWithPreallocation length: got %d, want %d", len(slice1), n)
		}
		
		// 测试不预分配
		slice2 := so.AppendWithoutPreallocation(n)
		if len(slice2) != n {
			t.Errorf("AppendWithoutPreallocation length: got %d, want %d", len(slice2), n)
		}
		
		// 验证内容相同
		for i := 0; i < n; i++ {
			if slice1[i] != slice2[i] {
				t.Errorf("Slice content differs at index %d: %d vs %d", i, slice1[i], slice2[i])
			}
		}
		
		t.Log("切片预分配测试通过")
	})
	
	t.Run("SliceCopy", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}
		
		// 测试优化复制
		dst1 := so.CopySliceOptimized(src)
		if len(dst1) != len(src) {
			t.Errorf("CopySliceOptimized length: got %d, want %d", len(dst1), len(src))
		}
		
		// 测试朴素复制
		dst2 := so.CopySliceNaive(src)
		if len(dst2) != len(src) {
			t.Errorf("CopySliceNaive length: got %d, want %d", len(dst2), len(src))
		}
		
		// 验证内容相同
		for i := 0; i < len(src); i++ {
			if dst1[i] != src[i] || dst2[i] != src[i] {
				t.Errorf("Copy content differs at index %d", i)
			}
		}
		
		t.Log("切片复制测试通过")
	})
}

func TestConcurrencyOptimization(t *testing.T) {
	co := &ConcurrencyOptimization{}
	
	t.Run("ConcurrentVsSequential", func(t *testing.T) {
		testData := make([]int, 1000)
		for i := range testData {
			testData[i] = i + 1
		}
		
		workers := 4
		
		// 测试并发处理
		result1 := co.ProcessDataConcurrent(testData, workers)
		if len(result1) != len(testData) {
			t.Errorf("ProcessDataConcurrent length: got %d, want %d", len(result1), len(testData))
		}
		
		// 测试顺序处理
		result2 := co.ProcessDataSequential(testData)
		if len(result2) != len(testData) {
			t.Errorf("ProcessDataSequential length: got %d, want %d", len(result2), len(testData))
		}
		
		// 验证结果相同
		for i := 0; i < len(testData); i++ {
			expected := testData[i] * testData[i]
			if result1[i] != expected || result2[i] != expected {
				t.Errorf("Processing result differs at index %d: %d vs %d (expected %d)", 
					i, result1[i], result2[i], expected)
			}
		}
		
		t.Log("并发处理测试通过")
	})
}

func TestWorkerPool(t *testing.T) {
	t.Run("WorkerPoolExecution", func(t *testing.T) {
		pool := NewWorkerPool(2)
		pool.Start()
		defer pool.Stop()
		
		// 提交任务
		taskCount := 10
		results := make([]int, taskCount)
		
		for i := 0; i < taskCount; i++ {
			i := i // 捕获循环变量
			pool.Submit(func() {
				results[i] = i * i
			})
		}
		
		// 等待所有任务完成
		pool.Wait()
		
		// 验证结果
		for i := 0; i < taskCount; i++ {
			expected := i * i
			if results[i] != expected {
				t.Errorf("Task result at index %d: got %d, want %d", i, results[i], expected)
			}
		}
		
		t.Log("工作池测试通过")
	})
}

func TestMemoryProfiler(t *testing.T) {
	t.Run("MemoryProfiling", func(t *testing.T) {
		profiler := NewMemoryProfiler()
		
		profiler.Start()
		
		// 分配一些内存
		data := make([][]byte, 100)
		for i := range data {
			data[i] = make([]byte, 1024)
		}
		
		stats := profiler.Stop()
		
		// 验证统计信息
		if stats["total_alloc"] == nil {
			t.Error("Memory profiler should return total_alloc")
		}
		
		if stats["mallocs"] == nil {
			t.Error("Memory profiler should return mallocs")
		}
		
		t.Logf("Memory stats: %+v", stats)
		t.Log("内存分析器测试通过")
	})
}

func TestTimeProfiler(t *testing.T) {
	t.Run("TimeProfiling", func(t *testing.T) {
		profiler := NewTimeProfiler()
		
		profiler.Start()
		
		// 模拟一些工作
		time.Sleep(10 * time.Millisecond)
		
		duration := profiler.Stop()
		
		// 验证时间测量
		if duration < 10*time.Millisecond {
			t.Errorf("Time profiler duration too short: %v", duration)
		}
		
		if duration > 50*time.Millisecond {
			t.Errorf("Time profiler duration too long: %v", duration)
		}
		
		t.Logf("Measured duration: %v", duration)
		t.Log("时间分析器测试通过")
	})
}

// 基准测试
func BenchmarkMemoryPoolVsNoPool(b *testing.B) {
	mo := NewMemoryOptimization()
	testData := make([]byte, 1024)
	
	b.Run("WithPool", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mo.ProcessDataWithPool(testData)
		}
	})
	
	b.Run("WithoutPool", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mo.ProcessDataWithoutPool(testData)
		}
	})
}

func BenchmarkStringConcatenation(b *testing.B) {
	so := &StringOptimization{}
	testStrings := make([]string, 100)
	for i := range testStrings {
		testStrings[i] = "test"
	}
	
	b.Run("StringBuilder", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			so.ConcatenateStringsBuilder(testStrings)
		}
	})
	
	b.Run("NaiveConcat", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			so.ConcatenateStringsNaive(testStrings)
		}
	})
}

func BenchmarkSlicePreallocation(b *testing.B) {
	so := &SliceOptimization{}
	n := 1000
	
	b.Run("WithPreallocation", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			so.AppendWithPreallocation(n)
		}
	})
	
	b.Run("WithoutPreallocation", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			so.AppendWithoutPreallocation(n)
		}
	})
}

func BenchmarkConcurrentProcessing(b *testing.B) {
	co := &ConcurrencyOptimization{}
	testData := make([]int, 10000)
	for i := range testData {
		testData[i] = i
	}
	
	b.Run("Concurrent", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			co.ProcessDataConcurrent(testData, 4)
		}
	})
	
	b.Run("Sequential", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			co.ProcessDataSequential(testData)
		}
	})
}
