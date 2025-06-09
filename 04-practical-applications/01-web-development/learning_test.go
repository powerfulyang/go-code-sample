package webdev

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// 🎓 学习导向的测试 - 通过测试学习Go Web开发

// TestLearnHTTPBasics 学习HTTP基础
func TestLearnHTTPBasics(t *testing.T) {
	t.Log("🎯 学习目标: 理解Go HTTP编程的基础概念")
	t.Log("📚 本测试将教您: HTTP处理器、路由、中间件")
	
	t.Run("学习基础HTTP处理器", func(t *testing.T) {
		t.Log("📖 知识点: HTTP处理器是处理HTTP请求的函数")
		
		// 🔍 探索: 创建简单的HTTP处理器
		helloHandler := func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "World"
			}
			
			response := fmt.Sprintf("Hello, %s!", name)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}
		
		// 使用httptest测试处理器
		t.Log("🔍 HTTP处理器测试:")
		
		// 测试默认情况
		req1 := httptest.NewRequest("GET", "/hello", nil)
		w1 := httptest.NewRecorder()
		helloHandler(w1, req1)
		
		resp1 := w1.Result()
		body1, _ := io.ReadAll(resp1.Body)
		
		t.Logf("   请求: GET /hello")
		t.Logf("   状态码: %d", resp1.StatusCode)
		t.Logf("   响应体: %s", string(body1))
		t.Logf("   Content-Type: %s", resp1.Header.Get("Content-Type"))
		
		// 测试带参数的情况
		req2 := httptest.NewRequest("GET", "/hello?name=Go", nil)
		w2 := httptest.NewRecorder()
		helloHandler(w2, req2)
		
		resp2 := w2.Result()
		body2, _ := io.ReadAll(resp2.Body)
		
		t.Logf("   请求: GET /hello?name=Go")
		t.Logf("   响应体: %s", string(body2))
		
		// ✅ 验证HTTP处理器
		if resp1.StatusCode != http.StatusOK {
			t.Errorf("❌ 状态码错误: 期望200，得到%d", resp1.StatusCode)
		}
		if string(body1) != "Hello, World!" {
			t.Errorf("❌ 响应体错误: 期望'Hello, World!'，得到'%s'", string(body1))
		}
		if string(body2) != "Hello, Go!" {
			t.Errorf("❌ 响应体错误: 期望'Hello, Go!'，得到'%s'", string(body2))
		}
		
		t.Log("✅ 很好！您理解了HTTP处理器的基本使用")
		
		// 💡 学习提示
		t.Log("💡 处理器签名: func(http.ResponseWriter, *http.Request)")
		t.Log("💡 响应写入: 先设置头部，再写入状态码和响应体")
		t.Log("💡 测试工具: httptest包提供了测试HTTP的工具")
	})
	
	t.Run("学习JSON API处理", func(t *testing.T) {
		t.Log("📖 知识点: 现代Web应用通常使用JSON进行数据交换")
		
		// 🔍 探索: JSON API处理器
		type User struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		
		type Response struct {
			Success bool        `json:"success"`
			Data    interface{} `json:"data,omitempty"`
			Error   string      `json:"error,omitempty"`
		}
		
		// 模拟用户数据
		users := []User{
			{ID: 1, Name: "张三", Age: 25},
			{ID: 2, Name: "李四", Age: 30},
		}
		
		// GET /users - 获取用户列表
		getUsersHandler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			
			response := Response{
				Success: true,
				Data:    users,
			}
			
			json.NewEncoder(w).Encode(response)
		}
		
		// POST /users - 创建用户
		createUserHandler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response := Response{
					Success: false,
					Error:   "Invalid JSON",
				}
				json.NewEncoder(w).Encode(response)
				return
			}
			
			// 分配新ID
			user.ID = len(users) + 1
			users = append(users, user)
			
			w.WriteHeader(http.StatusCreated)
			response := Response{
				Success: true,
				Data:    user,
			}
			json.NewEncoder(w).Encode(response)
		}
		
		// 测试GET请求
		t.Log("🔍 JSON API测试:")
		
		req1 := httptest.NewRequest("GET", "/users", nil)
		w1 := httptest.NewRecorder()
		getUsersHandler(w1, req1)
		
		resp1 := w1.Result()
		var getResponse Response
		json.NewDecoder(resp1.Body).Decode(&getResponse)
		
		t.Logf("   GET /users:")
		t.Logf("   状态码: %d", resp1.StatusCode)
		t.Logf("   成功: %t", getResponse.Success)
		t.Logf("   用户数量: %d", len(getResponse.Data.([]interface{})))
		
		// 测试POST请求
		newUser := User{Name: "王五", Age: 28}
		jsonData, _ := json.Marshal(newUser)
		
		req2 := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
		req2.Header.Set("Content-Type", "application/json")
		w2 := httptest.NewRecorder()
		createUserHandler(w2, req2)
		
		resp2 := w2.Result()
		var postResponse Response
		json.NewDecoder(resp2.Body).Decode(&postResponse)
		
		t.Logf("   POST /users:")
		t.Logf("   状态码: %d", resp2.StatusCode)
		t.Logf("   成功: %t", postResponse.Success)
		
		// ✅ 验证JSON API
		if resp1.StatusCode != http.StatusOK {
			t.Errorf("❌ GET状态码错误: 期望200，得到%d", resp1.StatusCode)
		}
		if !getResponse.Success {
			t.Error("❌ GET响应应该成功")
		}
		if resp2.StatusCode != http.StatusCreated {
			t.Errorf("❌ POST状态码错误: 期望201，得到%d", resp2.StatusCode)
		}
		if !postResponse.Success {
			t.Error("❌ POST响应应该成功")
		}
		
		t.Log("✅ 很好！您理解了JSON API的处理")
		
		// 💡 学习提示
		t.Log("💡 JSON编解码: 使用json.Encoder/Decoder处理JSON")
		t.Log("💡 Content-Type: 设置正确的响应头")
		t.Log("💡 状态码: 使用合适的HTTP状态码")
	})
}

// TestLearnMiddleware 学习中间件
func TestLearnMiddleware(t *testing.T) {
	t.Log("🎯 学习目标: 掌握HTTP中间件的概念和实现")
	t.Log("📚 本测试将教您: 中间件模式、日志记录、认证授权")
	
	t.Run("学习基础中间件", func(t *testing.T) {
		t.Log("📖 知识点: 中间件是包装HTTP处理器的函数")
		
		// 🔍 探索: 日志中间件
		var logBuffer bytes.Buffer
		
		loggingMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()
				
				// 包装ResponseWriter以捕获状态码
				wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
				
				// 调用下一个处理器
				next(wrapped, r)
				
				// 记录日志
				duration := time.Since(start)
				logEntry := fmt.Sprintf("%s %s %d %v\n", 
					r.Method, r.URL.Path, wrapped.statusCode, duration)
				logBuffer.WriteString(logEntry)
			}
		}
		
		// 基础处理器
		helloHandler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, Middleware!"))
		}
		
		// 应用中间件
		wrappedHandler := loggingMiddleware(helloHandler)
		
		// 测试中间件
		t.Log("🔍 中间件测试:")
		
		req := httptest.NewRequest("GET", "/hello", nil)
		w := httptest.NewRecorder()
		wrappedHandler(w, req)
		
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		logOutput := logBuffer.String()
		
		t.Logf("   请求: GET /hello")
		t.Logf("   响应: %s", string(body))
		t.Logf("   日志: %s", strings.TrimSpace(logOutput))
		
		// ✅ 验证中间件
		if resp.StatusCode != http.StatusOK {
			t.Errorf("❌ 状态码错误: 期望200，得到%d", resp.StatusCode)
		}
		if !strings.Contains(logOutput, "GET /hello 200") {
			t.Errorf("❌ 日志格式错误: %s", logOutput)
		}
		
		t.Log("✅ 很好！您理解了中间件的基本概念")
		
		// 💡 学习提示
		t.Log("💡 中间件模式: 函数返回函数的高阶函数")
		t.Log("💡 链式调用: 多个中间件可以链式组合")
		t.Log("💡 横切关注点: 日志、认证、CORS等")
	})
	
	t.Run("学习认证中间件", func(t *testing.T) {
		t.Log("📖 知识点: 认证中间件用于保护需要授权的资源")
		
		// 🔍 探索: 简单的Token认证中间件
		authMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				token := r.Header.Get("Authorization")
				
				// 简单的token验证
				if token != "Bearer valid-token" {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
					return
				}
				
				// 验证通过，继续处理
				next(w, r)
			}
		}
		
		// 受保护的处理器
		protectedHandler := func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Protected Resource"))
		}
		
		// 应用认证中间件
		wrappedHandler := authMiddleware(protectedHandler)
		
		// 测试无token的情况
		t.Log("🔍 认证中间件测试:")
		
		req1 := httptest.NewRequest("GET", "/protected", nil)
		w1 := httptest.NewRecorder()
		wrappedHandler(w1, req1)
		
		resp1 := w1.Result()
		body1, _ := io.ReadAll(resp1.Body)
		
		t.Logf("   无token请求:")
		t.Logf("   状态码: %d", resp1.StatusCode)
		t.Logf("   响应: %s", string(body1))
		
		// 测试有效token的情况
		req2 := httptest.NewRequest("GET", "/protected", nil)
		req2.Header.Set("Authorization", "Bearer valid-token")
		w2 := httptest.NewRecorder()
		wrappedHandler(w2, req2)
		
		resp2 := w2.Result()
		body2, _ := io.ReadAll(resp2.Body)
		
		t.Logf("   有效token请求:")
		t.Logf("   状态码: %d", resp2.StatusCode)
		t.Logf("   响应: %s", string(body2))
		
		// ✅ 验证认证中间件
		if resp1.StatusCode != http.StatusUnauthorized {
			t.Errorf("❌ 无token状态码错误: 期望401，得到%d", resp1.StatusCode)
		}
		if resp2.StatusCode != http.StatusOK {
			t.Errorf("❌ 有效token状态码错误: 期望200，得到%d", resp2.StatusCode)
		}
		if string(body2) != "Protected Resource" {
			t.Errorf("❌ 受保护资源响应错误: %s", string(body2))
		}
		
		t.Log("✅ 很好！您理解了认证中间件")
		
		// 💡 学习提示
		t.Log("💡 早期返回: 认证失败时直接返回，不调用下一个处理器")
		t.Log("💡 状态码: 使用合适的HTTP状态码表示认证状态")
		t.Log("💡 安全性: 实际应用中应使用更安全的认证机制")
	})
}

// TestLearnHTTPClient 学习HTTP客户端
func TestLearnHTTPClient(t *testing.T) {
	t.Log("🎯 学习目标: 掌握Go HTTP客户端编程")
	t.Log("📚 本测试将教您: HTTP请求、响应处理、超时控制")
	
	t.Run("学习HTTP客户端基础", func(t *testing.T) {
		t.Log("📖 知识点: Go提供了强大的HTTP客户端功能")
		
		// 🔍 探索: 创建测试服务器
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/api/users":
				users := []map[string]interface{}{
					{"id": 1, "name": "张三"},
					{"id": 2, "name": "李四"},
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(users)
			case "/api/slow":
				time.Sleep(100 * time.Millisecond)
				w.Write([]byte("Slow response"))
			default:
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Not Found"))
			}
		}))
		defer server.Close()
		
		// 创建HTTP客户端
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		
		t.Log("🔍 HTTP客户端测试:")
		
		// GET请求
		resp, err := client.Get(server.URL + "/api/users")
		if err != nil {
			t.Fatalf("GET请求失败: %v", err)
		}
		defer resp.Body.Close()
		
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("读取响应体失败: %v", err)
		}
		
		t.Logf("   GET %s/api/users", server.URL)
		t.Logf("   状态码: %d", resp.StatusCode)
		t.Logf("   Content-Type: %s", resp.Header.Get("Content-Type"))
		t.Logf("   响应体: %s", string(body))
		
		// 解析JSON响应
		var users []map[string]interface{}
		if err := json.Unmarshal(body, &users); err != nil {
			t.Fatalf("JSON解析失败: %v", err)
		}
		
		t.Logf("   解析的用户数量: %d", len(users))
		
		// ✅ 验证HTTP客户端
		if resp.StatusCode != http.StatusOK {
			t.Errorf("❌ 状态码错误: 期望200，得到%d", resp.StatusCode)
		}
		if len(users) != 2 {
			t.Errorf("❌ 用户数量错误: 期望2，得到%d", len(users))
		}
		
		t.Log("✅ 很好！您理解了HTTP客户端的基本使用")
		
		// 💡 学习提示
		t.Log("💡 资源管理: 记得关闭响应体")
		t.Log("💡 错误处理: 检查网络错误和HTTP状态码")
		t.Log("💡 超时控制: 设置合理的超时时间")
	})
}

// responseWriter 包装ResponseWriter以捕获状态码
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// BenchmarkLearnHTTPPerformance 学习HTTP性能
func BenchmarkLearnHTTPPerformance(b *testing.B) {
	b.Log("🎯 学习目标: 了解HTTP处理的性能特征")
	
	// 简单处理器
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}
	
	b.Run("基础处理器", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			handler(w, req)
		}
	})
	
	// 带中间件的处理器
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 简单的中间件逻辑
			w.Header().Set("X-Middleware", "true")
			next(w, r)
		}
	}
	
	wrappedHandler := middleware(handler)
	
	b.Run("带中间件处理器", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			wrappedHandler(w, req)
		}
	})
}

// Example_learnBasicHTTP HTTP基础示例
func Example_learnBasicHTTP() {
	// 创建简单的HTTP处理器
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("name"))
	}
	
	// 测试处理器
	req := httptest.NewRequest("GET", "/?name=Go", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应: %s\n", string(body))
	
	// Output:
	// 状态码: 200
	// 响应: Hello, Go!
}
