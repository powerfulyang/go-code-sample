package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// 测试辅助函数
func setupTestTaskManager(t testing.TB) (*TaskManager, func()) {
	// 创建临时文件
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test_tasks.json")

	tm := NewTaskManager(filename)

	cleanup := func() {
		os.Remove(filename)
	}

	return tm, cleanup
}

func TestTaskManager_AddTask(t *testing.T) {
	tm, cleanup := setupTestTaskManager(t)
	defer cleanup()

	t.Run("AddValidTask", func(t *testing.T) {
		task, err := tm.AddTask("测试任务", "这是一个测试任务", "high", nil)
		if err != nil {
			t.Errorf("添加任务失败: %v", err)
		}

		if task.ID == 0 {
			t.Error("任务ID应该被设置")
		}

		if task.Title != "测试任务" {
			t.Errorf("任务标题不匹配: 期望 %s, 实际 %s", "测试任务", task.Title)
		}

		if task.Priority != "high" {
			t.Errorf("任务优先级不匹配: 期望 %s, 实际 %s", "high", task.Priority)
		}

		if task.Completed {
			t.Error("新任务不应该是已完成状态")
		}

		if task.CreatedAt.IsZero() {
			t.Error("CreatedAt应该被设置")
		}

		t.Log("添加有效任务测试通过")
	})

	t.Run("AddTaskWithDueDate", func(t *testing.T) {
		dueDate := time.Now().AddDate(0, 0, 7) // 一周后
		task, err := tm.AddTask("有截止日期的任务", "描述", "medium", &dueDate)
		if err != nil {
			t.Errorf("添加任务失败: %v", err)
		}

		if task.DueDate == nil {
			t.Error("截止日期应该被设置")
		}

		if !task.DueDate.Equal(dueDate) {
			t.Errorf("截止日期不匹配: 期望 %v, 实际 %v", dueDate, *task.DueDate)
		}

		t.Log("添加带截止日期任务测试通过")
	})

	t.Run("AddTaskEmptyTitle", func(t *testing.T) {
		_, err := tm.AddTask("", "描述", "medium", nil)
		if err == nil {
			t.Error("空标题应该返回错误")
		}

		t.Log("空标题测试通过")
	})

	t.Run("AddTaskInvalidPriority", func(t *testing.T) {
		_, err := tm.AddTask("任务", "描述", "invalid", nil)
		if err == nil {
			t.Error("无效优先级应该返回错误")
		}

		t.Log("无效优先级测试通过")
	})

	t.Run("AddTaskDefaultPriority", func(t *testing.T) {
		task, err := tm.AddTask("默认优先级任务", "描述", "", nil)
		if err != nil {
			t.Errorf("添加任务失败: %v", err)
		}

		if task.Priority != "medium" {
			t.Errorf("默认优先级应该是medium, 实际 %s", task.Priority)
		}

		t.Log("默认优先级测试通过")
	})
}

func TestTaskManager_GetTask(t *testing.T) {
	tm, cleanup := setupTestTaskManager(t)
	defer cleanup()

	t.Run("GetExistingTask", func(t *testing.T) {
		// 添加任务
		originalTask, err := tm.AddTask("查找任务", "描述", "high", nil)
		if err != nil {
			t.Fatalf("添加任务失败: %v", err)
		}

		// 获取任务
		foundTask, err := tm.GetTask(originalTask.ID)
		if err != nil {
			t.Errorf("获取任务失败: %v", err)
		}

		if foundTask.ID != originalTask.ID {
			t.Errorf("任务ID不匹配: 期望 %d, 实际 %d", originalTask.ID, foundTask.ID)
		}

		if foundTask.Title != originalTask.Title {
			t.Errorf("任务标题不匹配: 期望 %s, 实际 %s", originalTask.Title, foundTask.Title)
		}

		t.Log("获取存在任务测试通过")
	})

	t.Run("GetNonExistentTask", func(t *testing.T) {
		_, err := tm.GetTask(9999)
		if err == nil {
			t.Error("获取不存在的任务应该返回错误")
		}

		t.Log("获取不存在任务测试通过")
	})
}

func TestTaskManager_UpdateTask(t *testing.T) {
	tm, cleanup := setupTestTaskManager(t)
	defer cleanup()

	t.Run("UpdateExistingTask", func(t *testing.T) {
		// 添加任务
		task, err := tm.AddTask("原始任务", "原始描述", "low", nil)
		if err != nil {
			t.Fatalf("添加任务失败: %v", err)
		}

		originalUpdatedAt := task.UpdatedAt
		time.Sleep(time.Millisecond) // 确保时间不同

		// 更新任务
		newDueDate := time.Now().AddDate(0, 0, 3)
		err = tm.UpdateTask(task.ID, "更新任务", "更新描述", "high", &newDueDate)
		if err != nil {
			t.Errorf("更新任务失败: %v", err)
		}

		// 验证更新
		updatedTask, err := tm.GetTask(task.ID)
		if err != nil {
			t.Errorf("获取更新后任务失败: %v", err)
		}

		if updatedTask.Title != "更新任务" {
			t.Errorf("标题未更新: 期望 %s, 实际 %s", "更新任务", updatedTask.Title)
		}

		if updatedTask.Description != "更新描述" {
			t.Errorf("描述未更新: 期望 %s, 实际 %s", "更新描述", updatedTask.Description)
		}

		if updatedTask.Priority != "high" {
			t.Errorf("优先级未更新: 期望 %s, 实际 %s", "high", updatedTask.Priority)
		}

		if updatedTask.DueDate == nil || !updatedTask.DueDate.Equal(newDueDate) {
			t.Error("截止日期未正确更新")
		}

		if updatedTask.UpdatedAt.Equal(originalUpdatedAt) {
			t.Error("UpdatedAt应该改变")
		}

		t.Log("更新任务测试通过")
	})

	t.Run("UpdateNonExistentTask", func(t *testing.T) {
		err := tm.UpdateTask(9999, "标题", "描述", "high", nil)
		if err == nil {
			t.Error("更新不存在的任务应该返回错误")
		}

		t.Log("更新不存在任务测试通过")
	})

	t.Run("UpdateTaskPartial", func(t *testing.T) {
		// 添加任务
		task, err := tm.AddTask("部分更新任务", "原始描述", "medium", nil)
		if err != nil {
			t.Fatalf("添加任务失败: %v", err)
		}

		// 只更新标题
		err = tm.UpdateTask(task.ID, "新标题", "", "", nil)
		if err != nil {
			t.Errorf("部分更新失败: %v", err)
		}

		// 验证更新
		updatedTask, err := tm.GetTask(task.ID)
		if err != nil {
			t.Errorf("获取更新后任务失败: %v", err)
		}

		if updatedTask.Title != "新标题" {
			t.Errorf("标题未更新: 期望 %s, 实际 %s", "新标题", updatedTask.Title)
		}

		if updatedTask.Description != "原始描述" {
			t.Errorf("描述不应该改变: 期望 %s, 实际 %s", "原始描述", updatedTask.Description)
		}

		if updatedTask.Priority != "medium" {
			t.Errorf("优先级不应该改变: 期望 %s, 实际 %s", "medium", updatedTask.Priority)
		}

		t.Log("部分更新任务测试通过")
	})
}

func TestTaskManager_CompleteTask(t *testing.T) {
	tm, cleanup := setupTestTaskManager(t)
	defer cleanup()

	t.Run("CompleteExistingTask", func(t *testing.T) {
		// 添加任务
		task, err := tm.AddTask("待完成任务", "描述", "high", nil)
		if err != nil {
			t.Fatalf("添加任务失败: %v", err)
		}

		if task.Completed {
			t.Error("新任务不应该是已完成状态")
		}

		// 完成任务
		err = tm.CompleteTask(task.ID)
		if err != nil {
			t.Errorf("完成任务失败: %v", err)
		}

		// 验证完成状态
		completedTask, err := tm.GetTask(task.ID)
		if err != nil {
			t.Errorf("获取完成后任务失败: %v", err)
		}

		if !completedTask.Completed {
			t.Error("任务应该是已完成状态")
		}

		t.Log("完成任务测试通过")
	})

	t.Run("CompleteNonExistentTask", func(t *testing.T) {
		err := tm.CompleteTask(9999)
		if err == nil {
			t.Error("完成不存在的任务应该返回错误")
		}

		t.Log("完成不存在任务测试通过")
	})
}

func TestTaskManager_DeleteTask(t *testing.T) {
	tm, cleanup := setupTestTaskManager(t)
	defer cleanup()

	t.Run("DeleteExistingTask", func(t *testing.T) {
		// 添加任务
		task, err := tm.AddTask("待删除任务", "描述", "medium", nil)
		if err != nil {
			t.Fatalf("添加任务失败: %v", err)
		}

		// 删除任务
		err = tm.DeleteTask(task.ID)
		if err != nil {
			t.Errorf("删除任务失败: %v", err)
		}

		// 验证删除
		_, err = tm.GetTask(task.ID)
		if err == nil {
			t.Error("删除后任务仍然存在")
		}

		t.Log("删除任务测试通过")
	})

	t.Run("DeleteNonExistentTask", func(t *testing.T) {
		err := tm.DeleteTask(9999)
		if err == nil {
			t.Error("删除不存在的任务应该返回错误")
		}

		t.Log("删除不存在任务测试通过")
	})
}

func TestTaskManager_ListTasks(t *testing.T) {
	tm, cleanup := setupTestTaskManager(t)
	defer cleanup()

	// 添加测试数据
	_, _ = tm.AddTask("高优先级任务", "描述1", "high", nil)
	task2, _ := tm.AddTask("中优先级任务", "描述2", "medium", nil)
	_, _ = tm.AddTask("低优先级任务", "描述3", "low", nil)

	// 完成一个任务
	tm.CompleteTask(task2.ID)

	// 添加过期任务
	pastDate := time.Now().AddDate(0, 0, -1) // 昨天
	_, _ = tm.AddTask("过期任务", "描述4", "medium", &pastDate)

	t.Run("ListAllTasks", func(t *testing.T) {
		tasks := tm.ListTasks("all")
		if len(tasks) != 4 {
			t.Errorf("任务总数不正确: 期望 4, 实际 %d", len(tasks))
		}

		// 验证排序（高优先级在前）
		if tasks[0].Priority != "high" {
			t.Errorf("第一个任务应该是高优先级, 实际 %s", tasks[0].Priority)
		}

		t.Log("列出所有任务测试通过")
	})

	t.Run("ListPendingTasks", func(t *testing.T) {
		tasks := tm.ListTasks("pending")

		pendingCount := 0
		for _, task := range tasks {
			if !task.Completed {
				pendingCount++
			}
		}

		if pendingCount != len(tasks) {
			t.Error("待完成任务列表包含已完成任务")
		}

		t.Log("列出待完成任务测试通过")
	})

	t.Run("ListCompletedTasks", func(t *testing.T) {
		tasks := tm.ListTasks("completed")

		for _, task := range tasks {
			if !task.Completed {
				t.Error("已完成任务列表包含未完成任务")
			}
		}

		t.Log("列出已完成任务测试通过")
	})

	t.Run("ListHighPriorityTasks", func(t *testing.T) {
		tasks := tm.ListTasks("high")

		for _, task := range tasks {
			if task.Priority != "high" || task.Completed {
				t.Error("高优先级任务列表包含非高优先级或已完成任务")
			}
		}

		t.Log("列出高优先级任务测试通过")
	})

	t.Run("ListOverdueTasks", func(t *testing.T) {
		tasks := tm.ListTasks("overdue")

		now := time.Now()
		for _, task := range tasks {
			if task.DueDate == nil || task.DueDate.After(now) || task.Completed {
				t.Error("过期任务列表包含非过期或已完成任务")
			}
		}

		if len(tasks) == 0 {
			t.Error("应该有过期任务")
		}

		t.Log("列出过期任务测试通过")
	})
}

func TestTaskManager_GetStats(t *testing.T) {
	tm, cleanup := setupTestTaskManager(t)
	defer cleanup()

	// 添加测试数据
	task1, _ := tm.AddTask("任务1", "描述1", "high", nil)
	_, _ = tm.AddTask("任务2", "描述2", "medium", nil)
	_, _ = tm.AddTask("任务3", "描述3", "low", nil)

	// 完成一个任务
	tm.CompleteTask(task1.ID)

	// 添加过期任务
	pastDate := time.Now().AddDate(0, 0, -1)
	tm.AddTask("过期任务", "描述", "medium", &pastDate)

	t.Run("GetStats", func(t *testing.T) {
		stats := tm.GetStats()

		if stats["total"] != 4 {
			t.Errorf("总任务数不正确: 期望 4, 实际 %d", stats["total"])
		}

		if stats["completed"] != 1 {
			t.Errorf("已完成任务数不正确: 期望 1, 实际 %d", stats["completed"])
		}

		if stats["pending"] != 3 {
			t.Errorf("待完成任务数不正确: 期望 3, 实际 %d", stats["pending"])
		}

		if stats["high"] != 1 {
			t.Errorf("高优先级任务数不正确: 期望 1, 实际 %d", stats["high"])
		}

		if stats["medium"] != 2 {
			t.Errorf("中优先级任务数不正确: 期望 2, 实际 %d", stats["medium"])
		}

		if stats["low"] != 1 {
			t.Errorf("低优先级任务数不正确: 期望 1, 实际 %d", stats["low"])
		}

		if stats["overdue"] != 1 {
			t.Errorf("过期任务数不正确: 期望 1, 实际 %d", stats["overdue"])
		}

		t.Log("获取统计信息测试通过")
	})
}

func TestTaskManager_Persistence(t *testing.T) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "persistence_test.json")

	t.Run("SaveAndLoadTasks", func(t *testing.T) {
		// 创建第一个任务管理器并添加任务
		tm1 := NewTaskManager(filename)
		task1, err := tm1.AddTask("持久化任务1", "描述1", "high", nil)
		if err != nil {
			t.Fatalf("添加任务失败: %v", err)
		}

		task2, err := tm1.AddTask("持久化任务2", "描述2", "medium", nil)
		if err != nil {
			t.Fatalf("添加任务失败: %v", err)
		}

		// 完成一个任务
		tm1.CompleteTask(task1.ID)

		// 创建第二个任务管理器，应该加载之前的任务
		tm2 := NewTaskManager(filename)

		// 验证任务被正确加载
		loadedTask1, err := tm2.GetTask(task1.ID)
		if err != nil {
			t.Errorf("加载任务1失败: %v", err)
		}

		if loadedTask1.Title != task1.Title {
			t.Errorf("任务1标题不匹配: 期望 %s, 实际 %s", task1.Title, loadedTask1.Title)
		}

		if !loadedTask1.Completed {
			t.Error("任务1应该是已完成状态")
		}

		loadedTask2, err := tm2.GetTask(task2.ID)
		if err != nil {
			t.Errorf("加载任务2失败: %v", err)
		}

		if loadedTask2.Title != task2.Title {
			t.Errorf("任务2标题不匹配: 期望 %s, 实际 %s", task2.Title, loadedTask2.Title)
		}

		if loadedTask2.Completed {
			t.Error("任务2不应该是已完成状态")
		}

		// 验证nextID正确设置
		newTask, err := tm2.AddTask("新任务", "描述", "low", nil)
		if err != nil {
			t.Errorf("添加新任务失败: %v", err)
		}

		expectedID := 3 // 应该是下一个ID
		if newTask.ID != expectedID {
			t.Errorf("新任务ID不正确: 期望 %d, 实际 %d", expectedID, newTask.ID)
		}

		t.Log("持久化测试通过")
	})
}

// 基准测试
func BenchmarkTaskManager_AddTask(b *testing.B) {
	tm, cleanup := setupTestTaskManager(b)
	defer cleanup()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tm.AddTask(fmt.Sprintf("基准任务%d", i), "描述", "medium", nil)
	}
}

func BenchmarkTaskManager_ListTasks(b *testing.B) {
	tm, cleanup := setupTestTaskManager(b)
	defer cleanup()

	// 添加一些任务
	for i := 0; i < 100; i++ {
		tm.AddTask(fmt.Sprintf("任务%d", i), "描述", "medium", nil)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tm.ListTasks("all")
	}
}

func BenchmarkTaskManager_GetStats(b *testing.B) {
	tm, cleanup := setupTestTaskManager(b)
	defer cleanup()

	// 添加一些任务
	for i := 0; i < 100; i++ {
		priority := []string{"high", "medium", "low"}[i%3]
		tm.AddTask(fmt.Sprintf("任务%d", i), "描述", priority, nil)
		if i%3 == 0 {
			tm.CompleteTask(i + 1) // 完成一些任务
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tm.GetStats()
	}
}
