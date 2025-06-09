package webapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestInMemoryUserRepository(t *testing.T) {
	repo := NewInMemoryUserRepository()

	t.Run("GetAll", func(t *testing.T) {
		users := repo.GetAll()
		if len(users) != 3 { // 种子数据有3个用户
			t.Errorf("Expected 3 users, got %d", len(users))
		}
		t.Log("GetAll测试通过")
	})

	t.Run("Create", func(t *testing.T) {
		user := &User{
			Name:  "测试用户",
			Email: "test@example.com",
			Age:   25,
		}

		err := repo.Create(user)
		if err != nil {
			t.Errorf("Create failed: %v", err)
		}

		if user.ID == 0 {
			t.Error("User ID should be set after creation")
		}

		if user.CreatedAt.IsZero() {
			t.Error("CreatedAt should be set after creation")
		}

		t.Log("Create测试通过")
	})

	t.Run("GetByID", func(t *testing.T) {
		// 创建用户
		user := &User{Name: "查找用户", Email: "find@example.com", Age: 30}
		repo.Create(user)

		// 查找用户
		found, err := repo.GetByID(user.ID)
		if err != nil {
			t.Errorf("GetByID failed: %v", err)
		}

		if found.Name != user.Name {
			t.Errorf("Expected name %s, got %s", user.Name, found.Name)
		}

		// 查找不存在的用户
		_, err = repo.GetByID(9999)
		if err == nil {
			t.Error("GetByID should return error for non-existent user")
		}

		t.Log("GetByID测试通过")
	})

	t.Run("Update", func(t *testing.T) {
		// 创建用户
		user := &User{Name: "原始用户", Email: "original@example.com", Age: 25}
		repo.Create(user)
		originalCreatedAt := user.CreatedAt

		// 更新用户
		time.Sleep(time.Millisecond) // 确保时间不同
		user.Name = "更新用户"
		user.Age = 30
		err := repo.Update(user)
		if err != nil {
			t.Errorf("Update failed: %v", err)
		}

		// 验证更新
		updated, _ := repo.GetByID(user.ID)
		if updated.Name != "更新用户" {
			t.Errorf("Expected updated name, got %s", updated.Name)
		}

		if updated.CreatedAt != originalCreatedAt {
			t.Error("CreatedAt should not change during update")
		}

		if updated.UpdatedAt.Equal(originalCreatedAt) {
			t.Error("UpdatedAt should change during update")
		}

		t.Log("Update测试通过")
	})

	t.Run("Delete", func(t *testing.T) {
		// 创建用户
		user := &User{Name: "删除用户", Email: "delete@example.com", Age: 25}
		repo.Create(user)

		// 删除用户
		err := repo.Delete(user.ID)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
		}

		// 验证删除
		_, err = repo.GetByID(user.ID)
		if err == nil {
			t.Error("User should not exist after deletion")
		}

		// 删除不存在的用户
		err = repo.Delete(9999)
		if err == nil {
			t.Error("Delete should return error for non-existent user")
		}

		t.Log("Delete测试通过")
	})
}

func TestUserService(t *testing.T) {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)

	t.Run("ValidateUser", func(t *testing.T) {
		// 有效用户
		validUser := &User{
			Name:  "有效用户",
			Email: "valid@example.com",
			Age:   25,
		}

		err := service.CreateUser(validUser)
		if err != nil {
			t.Errorf("Valid user should be created: %v", err)
		}

		// 无效用户 - 空名称
		invalidUser1 := &User{
			Name:  "",
			Email: "test@example.com",
			Age:   25,
		}

		err = service.CreateUser(invalidUser1)
		if err == nil {
			t.Error("Should return error for empty name")
		}

		// 无效用户 - 空邮箱
		invalidUser2 := &User{
			Name:  "测试",
			Email: "",
			Age:   25,
		}

		err = service.CreateUser(invalidUser2)
		if err == nil {
			t.Error("Should return error for empty email")
		}

		// 无效用户 - 无效年龄
		invalidUser3 := &User{
			Name:  "测试",
			Email: "test@example.com",
			Age:   -1,
		}

		err = service.CreateUser(invalidUser3)
		if err == nil {
			t.Error("Should return error for invalid age")
		}

		// 无效用户 - 无效邮箱格式
		invalidUser4 := &User{
			Name:  "测试",
			Email: "invalid-email",
			Age:   25,
		}

		err = service.CreateUser(invalidUser4)
		if err == nil {
			t.Error("Should return error for invalid email format")
		}

		t.Log("ValidateUser测试通过")
	})
}

func TestUserHandler(t *testing.T) {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	t.Run("GetUsers", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		handler.GetUsers(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var users []User
		err := json.NewDecoder(w.Body).Decode(&users)
		if err != nil {
			t.Errorf("Failed to decode response: %v", err)
		}

		if len(users) != 3 { // 种子数据
			t.Errorf("Expected 3 users, got %d", len(users))
		}

		t.Log("GetUsers测试通过")
	})

	t.Run("GetUser", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		w := httptest.NewRecorder()

		handler.GetUser(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var user User
		err := json.NewDecoder(w.Body).Decode(&user)
		if err != nil {
			t.Errorf("Failed to decode response: %v", err)
		}

		if user.ID != 1 {
			t.Errorf("Expected user ID 1, got %d", user.ID)
		}

		t.Log("GetUser测试通过")
	})

	t.Run("GetUserNotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/9999", nil)
		w := httptest.NewRecorder()

		handler.GetUser(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", w.Code)
		}

		t.Log("GetUserNotFound测试通过")
	})

	t.Run("CreateUser", func(t *testing.T) {
		user := User{
			Name:  "新用户",
			Email: "new@example.com",
			Age:   25,
		}

		jsonData, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.CreateUser(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status 201, got %d", w.Code)
		}

		var createdUser User
		err := json.NewDecoder(w.Body).Decode(&createdUser)
		if err != nil {
			t.Errorf("Failed to decode response: %v", err)
		}

		if createdUser.ID == 0 {
			t.Error("Created user should have an ID")
		}

		if createdUser.Name != user.Name {
			t.Errorf("Expected name %s, got %s", user.Name, createdUser.Name)
		}

		t.Log("CreateUser测试通过")
	})

	t.Run("CreateUserInvalidData", func(t *testing.T) {
		user := User{
			Name:  "", // 无效：空名称
			Email: "test@example.com",
			Age:   25,
		}

		jsonData, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.CreateUser(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		t.Log("CreateUserInvalidData测试通过")
	})

	t.Run("UpdateUser", func(t *testing.T) {
		// 首先创建一个用户
		user := User{
			Name:  "更新前",
			Email: "before@example.com",
			Age:   25,
		}

		jsonData, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler.CreateUser(w, req)

		var createdUser User
		json.NewDecoder(w.Body).Decode(&createdUser)

		// 更新用户
		updatedUser := User{
			Name:  "更新后",
			Email: "after@example.com",
			Age:   30,
		}

		jsonData, _ = json.Marshal(updatedUser)
		req = httptest.NewRequest(http.MethodPut, "/users/"+fmt.Sprintf("%d", createdUser.ID), bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w = httptest.NewRecorder()

		handler.UpdateUser(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var result User
		err := json.NewDecoder(w.Body).Decode(&result)
		if err != nil {
			t.Errorf("Failed to decode response: %v", err)
		}

		if result.Name != "更新后" {
			t.Errorf("Expected updated name, got %s", result.Name)
		}

		t.Log("UpdateUser测试通过")
	})

	t.Run("DeleteUser", func(t *testing.T) {
		// 首先创建一个用户
		user := User{
			Name:  "待删除",
			Email: "delete@example.com",
			Age:   25,
		}

		jsonData, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler.CreateUser(w, req)

		var createdUser User
		json.NewDecoder(w.Body).Decode(&createdUser)

		// 删除用户
		req = httptest.NewRequest(http.MethodDelete, "/users/"+fmt.Sprintf("%d", createdUser.ID), nil)
		w = httptest.NewRecorder()

		handler.DeleteUser(w, req)

		if w.Code != http.StatusNoContent {
			t.Errorf("Expected status 204, got %d", w.Code)
		}

		// 验证用户已删除
		req = httptest.NewRequest(http.MethodGet, "/users/"+fmt.Sprintf("%d", createdUser.ID), nil)
		w = httptest.NewRecorder()
		handler.GetUser(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status 404 after deletion, got %d", w.Code)
		}

		t.Log("DeleteUser测试通过")
	})

	t.Run("MethodNotAllowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/users", nil)
		w := httptest.NewRecorder()

		handler.GetUsers(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status 405, got %d", w.Code)
		}

		t.Log("MethodNotAllowed测试通过")
	})
}

func TestServer(t *testing.T) {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)
	server := NewServer(handler, "8080")

	t.Run("HealthCheck", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		w := httptest.NewRecorder()

		server.healthHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var response map[string]interface{}
		err := json.NewDecoder(w.Body).Decode(&response)
		if err != nil {
			t.Errorf("Failed to decode response: %v", err)
		}

		if response["status"] != "healthy" {
			t.Errorf("Expected status 'healthy', got %v", response["status"])
		}

		t.Log("HealthCheck测试通过")
	})

	t.Run("RootHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		server.rootHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var response map[string]interface{}
		err := json.NewDecoder(w.Body).Decode(&response)
		if err != nil {
			t.Errorf("Failed to decode response: %v", err)
		}

		if !strings.Contains(response["message"].(string), "欢迎") {
			t.Error("Response should contain welcome message")
		}

		t.Log("RootHandler测试通过")
	})

	t.Run("NotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
		w := httptest.NewRecorder()

		server.rootHandler(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", w.Code)
		}

		t.Log("NotFound测试通过")
	})
}

func TestMiddleware(t *testing.T) {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)
	server := NewServer(handler, "8080")

	t.Run("CORSMiddleware", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodOptions, "/users", nil)
		w := httptest.NewRecorder()

		// 创建一个简单的处理器来测试CORS中间件
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		corsHandler := server.corsMiddleware(testHandler)
		corsHandler.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		// 检查CORS头
		if w.Header().Get("Access-Control-Allow-Origin") != "*" {
			t.Error("CORS Allow-Origin header not set correctly")
		}

		if !strings.Contains(w.Header().Get("Access-Control-Allow-Methods"), "GET") {
			t.Error("CORS Allow-Methods header not set correctly")
		}

		t.Log("CORSMiddleware测试通过")
	})
}

// 集成测试
func TestIntegration(t *testing.T) {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	t.Run("CompleteUserLifecycle", func(t *testing.T) {
		// 1. 创建用户
		user := User{
			Name:  "集成测试用户",
			Email: "integration@example.com",
			Age:   25,
		}

		jsonData, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.CreateUser(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("Failed to create user: status %d", w.Code)
		}

		var createdUser User
		json.NewDecoder(w.Body).Decode(&createdUser)

		// 2. 获取用户
		req = httptest.NewRequest(http.MethodGet, "/users/"+fmt.Sprintf("%d", createdUser.ID), nil)
		w = httptest.NewRecorder()

		handler.GetUser(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Failed to get user: status %d", w.Code)
		}

		// 3. 更新用户
		createdUser.Name = "更新的集成测试用户"
		jsonData, _ = json.Marshal(createdUser)
		req = httptest.NewRequest(http.MethodPut, "/users/"+fmt.Sprintf("%d", createdUser.ID), bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w = httptest.NewRecorder()

		handler.UpdateUser(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Failed to update user: status %d", w.Code)
		}

		// 4. 删除用户
		req = httptest.NewRequest(http.MethodDelete, "/users/"+fmt.Sprintf("%d", createdUser.ID), nil)
		w = httptest.NewRecorder()

		handler.DeleteUser(w, req)

		if w.Code != http.StatusNoContent {
			t.Fatalf("Failed to delete user: status %d", w.Code)
		}

		// 5. 验证用户已删除
		req = httptest.NewRequest(http.MethodGet, "/users/"+fmt.Sprintf("%d", createdUser.ID), nil)
		w = httptest.NewRecorder()

		handler.GetUser(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("User should be deleted: status %d", w.Code)
		}

		t.Log("CompleteUserLifecycle测试通过")
	})
}

// 基准测试
func BenchmarkGetUsers(b *testing.B) {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler.GetUsers(w, req)
	}
}

func BenchmarkCreateUser(b *testing.B) {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	user := User{
		Name:  "基准测试用户",
		Email: "benchmark@example.com",
		Age:   25,
	}

	jsonData, _ := json.Marshal(user)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler.CreateUser(w, req)
	}
}
