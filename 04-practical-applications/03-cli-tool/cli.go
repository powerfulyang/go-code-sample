package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Task 任务结构体
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	Priority    string     `json:"priority"` // high, medium, low
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

// TaskManager 任务管理器
type TaskManager struct {
	tasks    []Task
	nextID   int
	filename string
}

// NewTaskManager 创建任务管理器
func NewTaskManager(filename string) *TaskManager {
	tm := &TaskManager{
		tasks:    make([]Task, 0),
		nextID:   1,
		filename: filename,
	}

	// 尝试加载现有任务
	tm.loadTasks()

	return tm
}

// loadTasks 从文件加载任务
func (tm *TaskManager) loadTasks() error {
	if _, err := os.Stat(tm.filename); os.IsNotExist(err) {
		return nil // 文件不存在，返回空任务列表
	}

	file, err := os.Open(tm.filename)
	if err != nil {
		return fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tm.tasks); err != nil {
		return fmt.Errorf("解析JSON失败: %v", err)
	}

	// 更新nextID
	for _, task := range tm.tasks {
		if task.ID >= tm.nextID {
			tm.nextID = task.ID + 1
		}
	}

	return nil
}

// saveTasks 保存任务到文件
func (tm *TaskManager) saveTasks() error {
	// 确保目录存在
	dir := filepath.Dir(tm.filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	file, err := os.Create(tm.filename)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tm.tasks); err != nil {
		return fmt.Errorf("编码JSON失败: %v", err)
	}

	return nil
}

// AddTask 添加任务
func (tm *TaskManager) AddTask(title, description, priority string, dueDate *time.Time) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("任务标题不能为空")
	}

	if priority == "" {
		priority = "medium"
	}

	// 验证优先级
	validPriorities := map[string]bool{"high": true, "medium": true, "low": true}
	if !validPriorities[priority] {
		return nil, fmt.Errorf("无效的优先级: %s (可选: high, medium, low)", priority)
	}

	task := Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DueDate:     dueDate,
	}

	tm.tasks = append(tm.tasks, task)
	tm.nextID++

	if err := tm.saveTasks(); err != nil {
		return nil, err
	}

	return &task, nil
}

// GetTask 获取任务
func (tm *TaskManager) GetTask(id int) (*Task, error) {
	for i, task := range tm.tasks {
		if task.ID == id {
			return &tm.tasks[i], nil
		}
	}
	return nil, fmt.Errorf("任务不存在: %d", id)
}

// UpdateTask 更新任务
func (tm *TaskManager) UpdateTask(id int, title, description, priority string, dueDate *time.Time) error {
	task, err := tm.GetTask(id)
	if err != nil {
		return err
	}

	if title != "" {
		task.Title = title
	}
	if description != "" {
		task.Description = description
	}
	if priority != "" {
		validPriorities := map[string]bool{"high": true, "medium": true, "low": true}
		if !validPriorities[priority] {
			return fmt.Errorf("无效的优先级: %s", priority)
		}
		task.Priority = priority
	}
	if dueDate != nil {
		task.DueDate = dueDate
	}

	task.UpdatedAt = time.Now()

	return tm.saveTasks()
}

// CompleteTask 完成任务
func (tm *TaskManager) CompleteTask(id int) error {
	task, err := tm.GetTask(id)
	if err != nil {
		return err
	}

	task.Completed = true
	task.UpdatedAt = time.Now()

	return tm.saveTasks()
}

// DeleteTask 删除任务
func (tm *TaskManager) DeleteTask(id int) error {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return tm.saveTasks()
		}
	}
	return fmt.Errorf("任务不存在: %d", id)
}

// ListTasks 列出任务
func (tm *TaskManager) ListTasks(filter string) []Task {
	var filtered []Task

	for _, task := range tm.tasks {
		switch filter {
		case "completed":
			if task.Completed {
				filtered = append(filtered, task)
			}
		case "pending":
			if !task.Completed {
				filtered = append(filtered, task)
			}
		case "high":
			if task.Priority == "high" && !task.Completed {
				filtered = append(filtered, task)
			}
		case "overdue":
			if task.DueDate != nil && task.DueDate.Before(time.Now()) && !task.Completed {
				filtered = append(filtered, task)
			}
		default:
			filtered = append(filtered, task)
		}
	}

	// 按优先级和创建时间排序
	sort.Slice(filtered, func(i, j int) bool {
		priorityOrder := map[string]int{"high": 3, "medium": 2, "low": 1}

		if filtered[i].Priority != filtered[j].Priority {
			return priorityOrder[filtered[i].Priority] > priorityOrder[filtered[j].Priority]
		}

		return filtered[i].CreatedAt.Before(filtered[j].CreatedAt)
	})

	return filtered
}

// GetStats 获取统计信息
func (tm *TaskManager) GetStats() map[string]int {
	stats := map[string]int{
		"total":     len(tm.tasks),
		"completed": 0,
		"pending":   0,
		"high":      0,
		"medium":    0,
		"low":       0,
		"overdue":   0,
	}

	now := time.Now()
	for _, task := range tm.tasks {
		if task.Completed {
			stats["completed"]++
		} else {
			stats["pending"]++
		}

		stats[task.Priority]++

		if task.DueDate != nil && task.DueDate.Before(now) && !task.Completed {
			stats["overdue"]++
		}
	}

	return stats
}

// CLI 命令行界面
type CLI struct {
	taskManager *TaskManager
	reader      *bufio.Reader
}

// NewCLI 创建CLI
func NewCLI(filename string) *CLI {
	return &CLI{
		taskManager: NewTaskManager(filename),
		reader:      bufio.NewReader(os.Stdin),
	}
}

// Run 运行CLI
func (cli *CLI) Run() {
	fmt.Println("📝 欢迎使用任务管理器!")
	fmt.Println("输入 'help' 查看可用命令")
	fmt.Println()

	for {
		fmt.Print("task> ")
		input, err := cli.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\n再见!")
				break
			}
			fmt.Printf("读取输入错误: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if input == "exit" || input == "quit" {
			fmt.Println("再见!")
			break
		}

		cli.handleCommand(input)
		fmt.Println()
	}
}

// handleCommand 处理命令
func (cli *CLI) handleCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "help", "h":
		cli.showHelp()
	case "add", "a":
		cli.addTask(args)
	case "list", "ls", "l":
		cli.listTasks(args)
	case "show", "s":
		cli.showTask(args)
	case "update", "u":
		cli.updateTask(args)
	case "complete", "done", "c":
		cli.completeTask(args)
	case "delete", "del", "d":
		cli.deleteTask(args)
	case "stats":
		cli.showStats()
	default:
		fmt.Printf("未知命令: %s\n", command)
		fmt.Println("输入 'help' 查看可用命令")
	}
}

// showHelp 显示帮助
func (cli *CLI) showHelp() {
	fmt.Println("可用命令:")
	fmt.Println("  help, h                    - 显示此帮助信息")
	fmt.Println("  add, a <title>             - 添加新任务")
	fmt.Println("  list, ls, l [filter]       - 列出任务 (filter: all, pending, completed, high, overdue)")
	fmt.Println("  show, s <id>               - 显示任务详情")
	fmt.Println("  update, u <id>             - 更新任务")
	fmt.Println("  complete, done, c <id>     - 完成任务")
	fmt.Println("  delete, del, d <id>        - 删除任务")
	fmt.Println("  stats                      - 显示统计信息")
	fmt.Println("  exit, quit                 - 退出程序")
}

// addTask 添加任务
func (cli *CLI) addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("请提供任务标题")
		return
	}

	title := strings.Join(args, " ")

	fmt.Print("描述 (可选): ")
	description, _ := cli.reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("优先级 (high/medium/low, 默认 medium): ")
	priority, _ := cli.reader.ReadString('\n')
	priority = strings.TrimSpace(priority)
	if priority == "" {
		priority = "medium"
	}

	fmt.Print("截止日期 (YYYY-MM-DD, 可选): ")
	dueDateStr, _ := cli.reader.ReadString('\n')
	dueDateStr = strings.TrimSpace(dueDateStr)

	var dueDate *time.Time
	if dueDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", dueDateStr); err == nil {
			dueDate = &parsed
		} else {
			fmt.Printf("日期格式错误: %v\n", err)
			return
		}
	}

	task, err := cli.taskManager.AddTask(title, description, priority, dueDate)
	if err != nil {
		fmt.Printf("添加任务失败: %v\n", err)
		return
	}

	fmt.Printf("✅ 任务已添加: #%d %s\n", task.ID, task.Title)
}

// listTasks 列出任务
func (cli *CLI) listTasks(args []string) {
	filter := "all"
	if len(args) > 0 {
		filter = args[0]
	}

	tasks := cli.taskManager.ListTasks(filter)

	if len(tasks) == 0 {
		fmt.Printf("没有找到任务 (过滤器: %s)\n", filter)
		return
	}

	fmt.Printf("📋 任务列表 (过滤器: %s)\n", filter)
	fmt.Println(strings.Repeat("-", 80))

	for _, task := range tasks {
		status := "⏳"
		if task.Completed {
			status = "✅"
		}

		priority := ""
		switch task.Priority {
		case "high":
			priority = "🔴"
		case "medium":
			priority = "🟡"
		case "low":
			priority = "🟢"
		}

		dueDateStr := ""
		if task.DueDate != nil {
			dueDateStr = fmt.Sprintf(" (截止: %s)", task.DueDate.Format("2006-01-02"))
			if task.DueDate.Before(time.Now()) && !task.Completed {
				dueDateStr += " ⚠️"
			}
		}

		fmt.Printf("%s %s #%d %s%s\n", status, priority, task.ID, task.Title, dueDateStr)
	}
}

// showTask 显示任务详情
func (cli *CLI) showTask(args []string) {
	if len(args) == 0 {
		fmt.Println("请提供任务ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("无效的任务ID: %s\n", args[0])
		return
	}

	task, err := cli.taskManager.GetTask(id)
	if err != nil {
		fmt.Printf("获取任务失败: %v\n", err)
		return
	}

	fmt.Printf("📋 任务详情\n")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("标题: %s\n", task.Title)
	fmt.Printf("描述: %s\n", task.Description)
	fmt.Printf("状态: %s\n", map[bool]string{true: "已完成", false: "待完成"}[task.Completed])
	fmt.Printf("优先级: %s\n", task.Priority)
	fmt.Printf("创建时间: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("更新时间: %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))

	if task.DueDate != nil {
		fmt.Printf("截止日期: %s\n", task.DueDate.Format("2006-01-02"))
		if task.DueDate.Before(time.Now()) && !task.Completed {
			fmt.Println("⚠️ 任务已过期")
		}
	}
}

// updateTask 更新任务
func (cli *CLI) updateTask(args []string) {
	if len(args) == 0 {
		fmt.Println("请提供任务ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("无效的任务ID: %s\n", args[0])
		return
	}

	task, err := cli.taskManager.GetTask(id)
	if err != nil {
		fmt.Printf("获取任务失败: %v\n", err)
		return
	}

	fmt.Printf("更新任务 #%d: %s\n", task.ID, task.Title)
	fmt.Println("留空保持原值")

	fmt.Printf("新标题 (当前: %s): ", task.Title)
	title, _ := cli.reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("新描述 (当前: %s): ", task.Description)
	description, _ := cli.reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Printf("新优先级 (当前: %s): ", task.Priority)
	priority, _ := cli.reader.ReadString('\n')
	priority = strings.TrimSpace(priority)

	dueDateStr := ""
	if task.DueDate != nil {
		dueDateStr = task.DueDate.Format("2006-01-02")
	}
	fmt.Printf("新截止日期 (当前: %s): ", dueDateStr)
	newDueDateStr, _ := cli.reader.ReadString('\n')
	newDueDateStr = strings.TrimSpace(newDueDateStr)

	var dueDate *time.Time
	if newDueDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", newDueDateStr); err == nil {
			dueDate = &parsed
		} else {
			fmt.Printf("日期格式错误: %v\n", err)
			return
		}
	}

	err = cli.taskManager.UpdateTask(id, title, description, priority, dueDate)
	if err != nil {
		fmt.Printf("更新任务失败: %v\n", err)
		return
	}

	fmt.Printf("✅ 任务已更新: #%d\n", id)
}

// completeTask 完成任务
func (cli *CLI) completeTask(args []string) {
	if len(args) == 0 {
		fmt.Println("请提供任务ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("无效的任务ID: %s\n", args[0])
		return
	}

	err = cli.taskManager.CompleteTask(id)
	if err != nil {
		fmt.Printf("完成任务失败: %v\n", err)
		return
	}

	fmt.Printf("✅ 任务已完成: #%d\n", id)
}

// deleteTask 删除任务
func (cli *CLI) deleteTask(args []string) {
	if len(args) == 0 {
		fmt.Println("请提供任务ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("无效的任务ID: %s\n", args[0])
		return
	}

	task, err := cli.taskManager.GetTask(id)
	if err != nil {
		fmt.Printf("获取任务失败: %v\n", err)
		return
	}

	fmt.Printf("确定要删除任务 #%d: %s? (y/N): ", task.ID, task.Title)
	confirm, _ := cli.reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	if confirm != "y" && confirm != "yes" {
		fmt.Println("取消删除")
		return
	}

	err = cli.taskManager.DeleteTask(id)
	if err != nil {
		fmt.Printf("删除任务失败: %v\n", err)
		return
	}

	fmt.Printf("🗑️ 任务已删除: #%d\n", id)
}

// showStats 显示统计信息
func (cli *CLI) showStats() {
	stats := cli.taskManager.GetStats()

	fmt.Println("📊 任务统计")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("总任务数: %d\n", stats["total"])
	fmt.Printf("已完成: %d\n", stats["completed"])
	fmt.Printf("待完成: %d\n", stats["pending"])
	fmt.Println()
	fmt.Printf("高优先级: %d\n", stats["high"])
	fmt.Printf("中优先级: %d\n", stats["medium"])
	fmt.Printf("低优先级: %d\n", stats["low"])
	fmt.Println()
	fmt.Printf("过期任务: %d\n", stats["overdue"])

	if stats["total"] > 0 {
		completionRate := float64(stats["completed"]) / float64(stats["total"]) * 100
		fmt.Printf("完成率: %.1f%%\n", completionRate)
	}
}

// CLIExamples CLI工具示例
func CLIExamples() {
	fmt.Println("=== CLI工具示例 ===")
	fmt.Println()
	fmt.Println("这是一个任务管理CLI工具的演示")
	fmt.Println("功能包括:")
	fmt.Println("- 添加、更新、删除任务")
	fmt.Println("- 任务优先级管理")
	fmt.Println("- 截止日期提醒")
	fmt.Println("- 任务过滤和排序")
	fmt.Println("- 统计信息")
	fmt.Println("- 数据持久化")
	fmt.Println()
	fmt.Println("要运行CLI工具，请使用:")
	fmt.Println("  cli := NewCLI(\"tasks.json\")")
	fmt.Println("  cli.Run()")
	fmt.Println()
	fmt.Println("或者在测试中查看具体的API使用方法")

	// 演示API使用
	fmt.Println("\n🔹 API使用演示:")

	// 创建临时任务管理器
	tm := NewTaskManager("demo_tasks.json")
	defer os.Remove("demo_tasks.json") // 清理演示文件

	// 添加任务
	task1, _ := tm.AddTask("学习Go语言", "完成Go语言基础教程", "high", nil)
	fmt.Printf("添加任务: #%d %s\n", task1.ID, task1.Title)

	dueDate := time.Now().AddDate(0, 0, 7) // 一周后
	task2, _ := tm.AddTask("写项目文档", "完成项目的README文档", "medium", &dueDate)
	fmt.Printf("添加任务: #%d %s (截止: %s)\n", task2.ID, task2.Title, dueDate.Format("2006-01-02"))

	// 列出任务
	fmt.Println("\n任务列表:")
	tasks := tm.ListTasks("all")
	for _, task := range tasks {
		status := "待完成"
		if task.Completed {
			status = "已完成"
		}
		fmt.Printf("  #%d %s [%s] (%s)\n", task.ID, task.Title, task.Priority, status)
	}

	// 完成任务
	tm.CompleteTask(task1.ID)
	fmt.Printf("\n完成任务: #%d\n", task1.ID)

	// 显示统计
	stats := tm.GetStats()
	fmt.Printf("\n统计信息: 总计 %d, 已完成 %d, 待完成 %d\n",
		stats["total"], stats["completed"], stats["pending"])
}
