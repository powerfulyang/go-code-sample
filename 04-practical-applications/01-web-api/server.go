package webapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// User 用户模型
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository 用户仓库接口
type UserRepository interface {
	GetAll() []User
	GetByID(id int) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}

// InMemoryUserRepository 内存用户仓库实现
type InMemoryUserRepository struct {
	users  map[int]*User
	nextID int
	mutex  sync.RWMutex
}

// NewInMemoryUserRepository 创建内存用户仓库
func NewInMemoryUserRepository() *InMemoryUserRepository {
	repo := &InMemoryUserRepository{
		users:  make(map[int]*User),
		nextID: 1,
	}
	
	// 添加一些示例数据
	repo.seedData()
	return repo
}

func (r *InMemoryUserRepository) seedData() {
	users := []*User{
		{Name: "张三", Email: "zhangsan@example.com", Age: 25},
		{Name: "李四", Email: "lisi@example.com", Age: 30},
		{Name: "王五", Email: "wangwu@example.com", Age: 28},
	}
	
	for _, user := range users {
		r.Create(user)
	}
}

func (r *InMemoryUserRepository) GetAll() []User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	users := make([]User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users
}

func (r *InMemoryUserRepository) GetByID(id int) (*User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("用户不存在: %d", id)
	}
	
	// 返回副本
	userCopy := *user
	return &userCopy, nil
}

func (r *InMemoryUserRepository) Create(user *User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	user.ID = r.nextID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	
	r.users[user.ID] = user
	r.nextID++
	
	return nil
}

func (r *InMemoryUserRepository) Update(user *User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	existing, exists := r.users[user.ID]
	if !exists {
		return fmt.Errorf("用户不存在: %d", user.ID)
	}
	
	// 保留创建时间，更新其他字段
	user.CreatedAt = existing.CreatedAt
	user.UpdatedAt = time.Now()
	
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.users[id]; !exists {
		return fmt.Errorf("用户不存在: %d", id)
	}
	
	delete(r.users, id)
	return nil
}

// UserService 用户服务
type UserService struct {
	repo UserRepository
}

// NewUserService 创建用户服务
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() []User {
	return s.repo.GetAll()
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(user *User) error {
	// 验证用户数据
	if err := s.validateUser(user); err != nil {
		return err
	}
	
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(user *User) error {
	// 验证用户数据
	if err := s.validateUser(user); err != nil {
		return err
	}
	
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func (s *UserService) validateUser(user *User) error {
	if user.Name == "" {
		return fmt.Errorf("用户名不能为空")
	}
	if user.Email == "" {
		return fmt.Errorf("邮箱不能为空")
	}
	if user.Age < 0 || user.Age > 150 {
		return fmt.Errorf("年龄必须在0-150之间")
	}
	if !strings.Contains(user.Email, "@") {
		return fmt.Errorf("邮箱格式不正确")
	}
	return nil
}

// UserHandler HTTP处理器
type UserHandler struct {
	service *UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUsers 获取所有用户
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}
	
	users := h.service.GetAllUsers()
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "编码错误", http.StatusInternalServerError)
		return
	}
}

// GetUser 获取单个用户
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}
	
	// 从URL路径中提取ID
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "无效的用户ID", http.StatusBadRequest)
		return
	}
	
	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "编码错误", http.StatusInternalServerError)
		return
	}
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}
	
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "无效的JSON数据", http.StatusBadRequest)
		return
	}
	
	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "编码错误", http.StatusInternalServerError)
		return
	}
}

// UpdateUser 更新用户
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}
	
	// 从URL路径中提取ID
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "无效的用户ID", http.StatusBadRequest)
		return
	}
	
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "无效的JSON数据", http.StatusBadRequest)
		return
	}
	
	user.ID = id
	if err := h.service.UpdateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "编码错误", http.StatusInternalServerError)
		return
	}
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}
	
	// 从URL路径中提取ID
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "无效的用户ID", http.StatusBadRequest)
		return
	}
	
	if err := h.service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

// Server Web服务器
type Server struct {
	handler *UserHandler
	port    string
}

// NewServer 创建新服务器
func NewServer(handler *UserHandler, port string) *Server {
	return &Server{
		handler: handler,
		port:    port,
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	mux := http.NewServeMux()
	
	// 注册路由
	mux.HandleFunc("/users", s.usersHandler)
	mux.HandleFunc("/users/", s.userHandler)
	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/", s.rootHandler)
	
	// 添加中间件
	handler := s.loggingMiddleware(s.corsMiddleware(mux))
	
	fmt.Printf("🚀 服务器启动在端口 %s\n", s.port)
	fmt.Println("API端点:")
	fmt.Println("  GET    /users      - 获取所有用户")
	fmt.Println("  POST   /users      - 创建用户")
	fmt.Println("  GET    /users/{id} - 获取单个用户")
	fmt.Println("  PUT    /users/{id} - 更新用户")
	fmt.Println("  DELETE /users/{id} - 删除用户")
	fmt.Println("  GET    /health     - 健康检查")
	
	return http.ListenAndServe(":"+s.port, handler)
}

func (s *Server) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handler.GetUsers(w, r)
	case http.MethodPost:
		s.handler.CreateUser(w, r)
	default:
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
	}
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handler.GetUser(w, r)
	case http.MethodPut:
		s.handler.UpdateUser(w, r)
	case http.MethodDelete:
		s.handler.DeleteUser(w, r)
	default:
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}
	
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now(),
		"service":   "user-api",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	response := map[string]interface{}{
		"message": "欢迎使用用户API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"users":  "/users",
			"health": "/health",
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 中间件

// loggingMiddleware 日志中间件
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// 创建响应写入器包装器来捕获状态码
		wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		next.ServeHTTP(wrapper, r)
		
		duration := time.Since(start)
		log.Printf("%s %s %d %v", r.Method, r.URL.Path, wrapper.statusCode, duration)
	})
}

// corsMiddleware CORS中间件
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// responseWriter 响应写入器包装器
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// WebAPIExamples Web API示例
func WebAPIExamples() {
	fmt.Println("=== Web API 示例 ===")
	
	// 创建依赖
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)
	server := NewServer(handler, "8080")
	
	fmt.Println("启动Web API服务器...")
	fmt.Println("你可以使用以下命令测试API:")
	fmt.Println()
	fmt.Println("# 获取所有用户")
	fmt.Println("curl http://localhost:8080/users")
	fmt.Println()
	fmt.Println("# 获取单个用户")
	fmt.Println("curl http://localhost:8080/users/1")
	fmt.Println()
	fmt.Println("# 创建用户")
	fmt.Println(`curl -X POST http://localhost:8080/users \`)
	fmt.Println(`  -H "Content-Type: application/json" \`)
	fmt.Println(`  -d '{"name":"新用户","email":"new@example.com","age":25}'`)
	fmt.Println()
	fmt.Println("# 更新用户")
	fmt.Println(`curl -X PUT http://localhost:8080/users/1 \`)
	fmt.Println(`  -H "Content-Type: application/json" \`)
	fmt.Println(`  -d '{"name":"更新用户","email":"updated@example.com","age":30}'`)
	fmt.Println()
	fmt.Println("# 删除用户")
	fmt.Println("curl -X DELETE http://localhost:8080/users/1")
	fmt.Println()
	fmt.Println("# 健康检查")
	fmt.Println("curl http://localhost:8080/health")
	
	// 启动服务器（这会阻塞）
	if err := server.Start(); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
