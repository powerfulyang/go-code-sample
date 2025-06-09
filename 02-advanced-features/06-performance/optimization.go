package performance

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// MemoryOptimization 内存优化示例
type MemoryOptimization struct {
	data []byte
	pool *sync.Pool
}

// NewMemoryOptimization 创建内存优化实例
func NewMemoryOptimization() *MemoryOptimization {
	return &MemoryOptimization{
		pool: &sync.Pool{
			New: func() interface{} {
				return make([]byte, 1024) // 1KB缓冲区
			},
		},
	}
}

// GetBuffer 从对象池获取缓冲区
func (mo *MemoryOptimization) GetBuffer() []byte {
	return mo.pool.Get().([]byte)
}

// PutBuffer 将缓冲区放回对象池
func (mo *MemoryOptimization) PutBuffer(buf []byte) {
	// 重置缓冲区
	buf = buf[:0]
	mo.pool.Put(buf)
}

// ProcessDataWithPool 使用对象池处理数据
func (mo *MemoryOptimization) ProcessDataWithPool(data []byte) []byte {
	buf := mo.GetBuffer()
	defer mo.PutBuffer(buf)

	// 确保缓冲区足够大
	if cap(buf) < len(data) {
		buf = make([]byte, len(data))
	}

	buf = buf[:len(data)]
	copy(buf, data)

	// 模拟数据处理
	for i := range buf {
		buf[i] = buf[i] ^ 0xFF // 简单的位运算
	}

	// 返回处理后的数据副本
	result := make([]byte, len(buf))
	copy(result, buf)
	return result
}

// ProcessDataWithoutPool 不使用对象池处理数据
func (mo *MemoryOptimization) ProcessDataWithoutPool(data []byte) []byte {
	buf := make([]byte, len(data))
	copy(buf, data)

	// 模拟数据处理
	for i := range buf {
		buf[i] = buf[i] ^ 0xFF
	}

	return buf
}

// StringOptimization 字符串优化示例
type StringOptimization struct{}

// ConcatenateStringsBuilder 使用strings.Builder连接字符串
func (so *StringOptimization) ConcatenateStringsBuilder(strs []string) string {
	var builder strings.Builder

	// 预分配容量
	totalLen := 0
	for _, s := range strs {
		totalLen += len(s)
	}
	builder.Grow(totalLen)

	for _, s := range strs {
		builder.WriteString(s)
	}

	return builder.String()
}

// ConcatenateStringsNaive 朴素的字符串连接
func (so *StringOptimization) ConcatenateStringsNaive(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

// BytesToStringUnsafe 使用unsafe包进行零拷贝转换
func (so *StringOptimization) BytesToStringUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytesUnsafe 使用unsafe包进行零拷贝转换
func (so *StringOptimization) StringToBytesUnsafe(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// SliceOptimization 切片优化示例
type SliceOptimization struct{}

// AppendWithPreallocation 预分配切片容量
func (so *SliceOptimization) AppendWithPreallocation(n int) []int {
	// 预分配足够的容量
	slice := make([]int, 0, n)

	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}

	return slice
}

// AppendWithoutPreallocation 不预分配切片容量
func (so *SliceOptimization) AppendWithoutPreallocation(n int) []int {
	var slice []int

	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}

	return slice
}

// CopySliceOptimized 优化的切片复制
func (so *SliceOptimization) CopySliceOptimized(src []int) []int {
	// 一次性分配目标切片
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

// CopySliceNaive 朴素的切片复制
func (so *SliceOptimization) CopySliceNaive(src []int) []int {
	var dst []int
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}

// ConcurrencyOptimization 并发优化示例
type ConcurrencyOptimization struct{}

// ProcessDataConcurrent 并发处理数据
func (co *ConcurrencyOptimization) ProcessDataConcurrent(data []int, workers int) []int {
	if len(data) == 0 {
		return data
	}

	result := make([]int, len(data))
	chunkSize := len(data) / workers
	if chunkSize == 0 {
		chunkSize = 1
		workers = len(data)
	}

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == workers-1 {
			end = len(data) // 最后一个worker处理剩余数据
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			for j := start; j < end; j++ {
				// 模拟CPU密集型操作
				result[j] = data[j] * data[j]
			}
		}(start, end)
	}

	wg.Wait()
	return result
}

// ProcessDataSequential 顺序处理数据
func (co *ConcurrencyOptimization) ProcessDataSequential(data []int) []int {
	result := make([]int, len(data))

	for i, v := range data {
		// 模拟CPU密集型操作
		result[i] = v * v
	}

	return result
}

// WorkerPool 工作池实现
type WorkerPool struct {
	workers   int
	taskQueue chan func()
	wg        sync.WaitGroup
}

// NewWorkerPool 创建工作池
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers:   workers,
		taskQueue: make(chan func(), workers*2), // 缓冲队列
	}
}

// Start 启动工作池
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go wp.worker()
	}
}

// worker 工作协程
func (wp *WorkerPool) worker() {
	for task := range wp.taskQueue {
		task()
		wp.wg.Done()
	}
}

// Submit 提交任务
func (wp *WorkerPool) Submit(task func()) {
	wp.wg.Add(1)
	wp.taskQueue <- task
}

// Wait 等待所有任务完成
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

// Stop 停止工作池
func (wp *WorkerPool) Stop() {
	close(wp.taskQueue)
}

// MemoryProfiler 内存分析器
type MemoryProfiler struct {
	startStats runtime.MemStats
	endStats   runtime.MemStats
}

// NewMemoryProfiler 创建内存分析器
func NewMemoryProfiler() *MemoryProfiler {
	return &MemoryProfiler{}
}

// Start 开始内存分析
func (mp *MemoryProfiler) Start() {
	runtime.GC() // 强制垃圾回收
	runtime.ReadMemStats(&mp.startStats)
}

// Stop 停止内存分析并返回统计信息
func (mp *MemoryProfiler) Stop() map[string]interface{} {
	runtime.GC() // 强制垃圾回收
	runtime.ReadMemStats(&mp.endStats)

	return map[string]interface{}{
		"alloc_delta":    mp.endStats.Alloc - mp.startStats.Alloc,
		"total_alloc":    mp.endStats.TotalAlloc - mp.startStats.TotalAlloc,
		"mallocs":        mp.endStats.Mallocs - mp.startStats.Mallocs,
		"frees":          mp.endStats.Frees - mp.startStats.Frees,
		"heap_alloc":     mp.endStats.HeapAlloc - mp.startStats.HeapAlloc,
		"heap_objects":   mp.endStats.HeapObjects - mp.startStats.HeapObjects,
		"gc_cycles":      mp.endStats.NumGC - mp.startStats.NumGC,
		"pause_total_ns": mp.endStats.PauseTotalNs - mp.startStats.PauseTotalNs,
	}
}

// TimeProfiler 时间分析器
type TimeProfiler struct {
	startTime time.Time
	endTime   time.Time
}

// NewTimeProfiler 创建时间分析器
func NewTimeProfiler() *TimeProfiler {
	return &TimeProfiler{}
}

// Start 开始时间分析
func (tp *TimeProfiler) Start() {
	tp.startTime = time.Now()
}

// Stop 停止时间分析并返回耗时
func (tp *TimeProfiler) Stop() time.Duration {
	tp.endTime = time.Now()
	return tp.endTime.Sub(tp.startTime)
}

// PerformanceExamples 性能优化示例
func PerformanceExamples() {
	fmt.Println("=== 性能优化示例 ===")

	// 内存优化示例
	fmt.Println("\n🔹 内存优化示例:")

	mo := NewMemoryOptimization()
	testData := make([]byte, 1000)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// 使用内存分析器比较
	profiler := NewMemoryProfiler()

	// 测试对象池
	profiler.Start()
	for i := 0; i < 1000; i++ {
		mo.ProcessDataWithPool(testData)
	}
	poolStats := profiler.Stop()

	// 测试不使用对象池
	profiler.Start()
	for i := 0; i < 1000; i++ {
		mo.ProcessDataWithoutPool(testData)
	}
	noPoolStats := profiler.Stop()

	fmt.Printf("  对象池内存分配: %d bytes\n", poolStats["total_alloc"])
	fmt.Printf("  普通方式内存分配: %d bytes\n", noPoolStats["total_alloc"])

	// 字符串优化示例
	fmt.Println("\n🔹 字符串优化示例:")

	so := &StringOptimization{}
	testStrings := make([]string, 1000)
	for i := range testStrings {
		testStrings[i] = fmt.Sprintf("string_%d", i)
	}

	timeProfiler := NewTimeProfiler()

	// 测试strings.Builder
	timeProfiler.Start()
	result1 := so.ConcatenateStringsBuilder(testStrings)
	builderTime := timeProfiler.Stop()

	// 测试朴素连接
	timeProfiler.Start()
	result2 := so.ConcatenateStringsNaive(testStrings)
	naiveTime := timeProfiler.Stop()

	fmt.Printf("  strings.Builder: %v (长度: %d)\n", builderTime, len(result1))
	fmt.Printf("  朴素连接: %v (长度: %d)\n", naiveTime, len(result2))
	fmt.Printf("  性能提升: %.2fx\n", float64(naiveTime)/float64(builderTime))

	// 切片优化示例
	fmt.Println("\n🔹 切片优化示例:")

	sliceOpt := &SliceOptimization{}
	n := 10000

	// 测试预分配
	timeProfiler.Start()
	slice1 := sliceOpt.AppendWithPreallocation(n)
	preallocTime := timeProfiler.Stop()

	// 测试不预分配
	timeProfiler.Start()
	slice2 := sliceOpt.AppendWithoutPreallocation(n)
	noPreallocTime := timeProfiler.Stop()

	fmt.Printf("  预分配: %v (长度: %d)\n", preallocTime, len(slice1))
	fmt.Printf("  不预分配: %v (长度: %d)\n", noPreallocTime, len(slice2))
	fmt.Printf("  性能提升: %.2fx\n", float64(noPreallocTime)/float64(preallocTime))

	// 并发优化示例
	fmt.Println("\n🔹 并发优化示例:")

	co := &ConcurrencyOptimization{}
	testData2 := make([]int, 100000)
	for i := range testData2 {
		testData2[i] = i
	}

	workers := runtime.NumCPU()

	// 测试并发处理
	timeProfiler.Start()
	result3 := co.ProcessDataConcurrent(testData2, workers)
	concurrentTime := timeProfiler.Stop()

	// 测试顺序处理
	timeProfiler.Start()
	result4 := co.ProcessDataSequential(testData2)
	sequentialTime := timeProfiler.Stop()

	fmt.Printf("  并发处理(%d workers): %v (长度: %d)\n", workers, concurrentTime, len(result3))
	fmt.Printf("  顺序处理: %v (长度: %d)\n", sequentialTime, len(result4))
	fmt.Printf("  性能提升: %.2fx\n", float64(sequentialTime)/float64(concurrentTime))

	// 工作池示例
	fmt.Println("\n🔹 工作池示例:")

	pool := NewWorkerPool(4)
	pool.Start()
	defer pool.Stop()

	timeProfiler.Start()
	for i := 0; i < 100; i++ {
		i := i // 捕获循环变量
		pool.Submit(func() {
			// 模拟工作
			time.Sleep(time.Millisecond)
			_ = i * i
		})
	}
	pool.Wait()
	poolTime := timeProfiler.Stop()

	fmt.Printf("  工作池处理100个任务: %v\n", poolTime)

	fmt.Println("\n✅ 性能优化示例演示完成!")
	fmt.Println("💡 提示: 使用 go test -bench=. 运行基准测试")
	fmt.Println("💡 提示: 使用 go tool pprof 进行性能分析")
	fmt.Println("💡 提示: 预分配内存可以显著提升性能")
}
