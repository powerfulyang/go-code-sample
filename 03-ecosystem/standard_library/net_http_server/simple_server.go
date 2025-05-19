package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// User 用户结构体
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 模拟用户数据
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
}

func main() {
	fmt.Println("=== Go HTTP 服务器示例 ===")

	// 注册路由处理器
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/api/time", timeHandler)
	http.HandleFunc("/health", healthHandler)

	// 静态文件服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 启动服务器
	port := ":8080"
	fmt.Printf("服务器启动在端口 %s\n", port)
	fmt.Println("访问以下URL测试:")
	fmt.Println("  http://localhost:8080/")
	fmt.Println("  http://localhost:8080/hello")
	fmt.Println("  http://localhost:8080/users")
	fmt.Println("  http://localhost:8080/api/time")
	fmt.Println("  http://localhost:8080/health")

	// 启动HTTP服务器
	log.Fatal(http.ListenAndServe(port, nil))
}

// homeHandler 首页处理器
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go HTTP 服务器示例</title>
    <meta charset="UTF-8">
</head>
<body>
    <h1>欢迎来到 Go HTTP 服务器</h1>
    <h2>可用的端点:</h2>
    <ul>
        <li><a href="/hello">Hello 页面</a></li>
        <li><a href="/users">用户列表 (JSON)</a></li>
        <li><a href="/api/time">当前时间 (JSON)</a></li>
        <li><a href="/health">健康检查</a></li>
    </ul>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

// helloHandler Hello页面处理器
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 获取查询参数
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	// 根据请求方法处理
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s! 这是一个GET请求。\n", name)
		fmt.Fprintf(w, "请求时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	case http.MethodPost:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s! 这是一个POST请求。\n", name)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
	}
}

// usersHandler 用户列表处理器
func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// 将用户数据编码为JSON
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "编码JSON失败", http.StatusInternalServerError)
		return
	}
}

// userHandler 单个用户处理器
func userHandler(w http.ResponseWriter, r *http.Request) {
	// 从URL路径中提取用户ID
	path := r.URL.Path
	if len(path) < 7 { // "/user/" 的长度
		http.Error(w, "无效的用户ID", http.StatusBadRequest)
		return
	}

	userID := path[6:] // 提取 "/user/" 后面的部分

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "请求的用户ID: %s\n", userID)
	fmt.Fprintf(w, "请求方法: %s\n", r.Method)
	fmt.Fprintf(w, "User-Agent: %s\n", r.Header.Get("User-Agent"))
}

// timeHandler 时间API处理器
func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 创建时间响应
	timeResponse := map[string]interface{}{
		"current_time": time.Now().Format(time.RFC3339),
		"timestamp":    time.Now().Unix(),
		"timezone":     time.Now().Location().String(),
		"formatted":    time.Now().Format("2006年01月02日 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(timeResponse); err != nil {
		http.Error(w, "编码JSON失败", http.StatusInternalServerError)
		return
	}
}

// healthHandler 健康检查处理器
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"uptime":    "运行中",
		"version":   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(health); err != nil {
		http.Error(w, "编码JSON失败", http.StatusInternalServerError)
		return
	}
}
