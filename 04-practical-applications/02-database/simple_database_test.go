package database

import (
	"fmt"
	"testing"
	"time"
)

// 测试辅助函数
func setupSimpleTestDB(t testing.TB) *SimpleDatabaseManager {
	dm := NewSimpleDatabaseManager()
	return dm
}

func TestSimpleDatabaseManager_CreateUser(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("CreateValidUser", func(t *testing.T) {
		user := &SimpleUser{
			Name:  "测试用户",
			Email: "test@example.com",
			Age:   25,
		}

		err := dm.CreateUser(user)
		if err != nil {
			t.Errorf("创建用户失败: %v", err)
		}

		if user.ID == 0 {
			t.Error("用户ID应该被设置")
		}

		if user.CreatedAt.IsZero() {
			t.Error("CreatedAt应该被设置")
		}

		if user.UpdatedAt.IsZero() {
			t.Error("UpdatedAt应该被设置")
		}

		t.Log("创建用户测试通过")
	})
}

func TestSimpleDatabaseManager_GetUserByID(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("GetExistingUser", func(t *testing.T) {
		// 创建用户
		originalUser := &SimpleUser{
			Name:  "查找用户",
			Email: "find@example.com",
			Age:   28,
		}

		err := dm.CreateUser(originalUser)
		if err != nil {
			t.Fatalf("创建用户失败: %v", err)
		}

		// 查找用户
		foundUser, err := dm.GetUserByID(originalUser.ID)
		if err != nil {
			t.Errorf("查找用户失败: %v", err)
		}

		if foundUser.Name != originalUser.Name {
			t.Errorf("用户名不匹配: 期望 %s, 实际 %s", originalUser.Name, foundUser.Name)
		}

		if foundUser.Email != originalUser.Email {
			t.Errorf("邮箱不匹配: 期望 %s, 实际 %s", originalUser.Email, foundUser.Email)
		}

		t.Log("查找存在用户测试通过")
	})

	t.Run("GetNonExistentUser", func(t *testing.T) {
		_, err := dm.GetUserByID(9999)
		if err == nil {
			t.Error("查找不存在的用户应该返回错误")
		}

		t.Log("查找不存在用户测试通过")
	})
}

func TestSimpleDatabaseManager_GetAllUsers(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("GetAllUsers", func(t *testing.T) {
		// 获取初始用户数量（种子数据）
		initialUsers := dm.GetAllUsers()
		initialCount := len(initialUsers)

		// 创建新用户
		newUser := &SimpleUser{
			Name:  "新增用户",
			Email: "new@example.com",
			Age:   25,
		}

		err := dm.CreateUser(newUser)
		if err != nil {
			t.Fatalf("创建用户失败: %v", err)
		}

		// 再次获取用户列表
		allUsers := dm.GetAllUsers()

		if len(allUsers) != initialCount+1 {
			t.Errorf("用户数量不正确: 期望 %d, 实际 %d", initialCount+1, len(allUsers))
		}

		t.Log("获取所有用户测试通过")
	})
}

func TestSimpleDatabaseManager_UpdateUser(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("UpdateExistingUser", func(t *testing.T) {
		// 创建用户
		user := &SimpleUser{
			Name:  "原始用户",
			Email: "original@example.com",
			Age:   25,
		}

		err := dm.CreateUser(user)
		if err != nil {
			t.Fatalf("创建用户失败: %v", err)
		}

		originalCreatedAt := user.CreatedAt
		time.Sleep(time.Millisecond) // 确保时间不同

		// 更新用户
		user.Name = "更新用户"
		user.Age = 30

		err = dm.UpdateUser(user)
		if err != nil {
			t.Errorf("更新用户失败: %v", err)
		}

		// 验证更新
		updatedUser, err := dm.GetUserByID(user.ID)
		if err != nil {
			t.Errorf("获取更新后用户失败: %v", err)
		}

		if updatedUser.Name != "更新用户" {
			t.Errorf("用户名未更新: 期望 %s, 实际 %s", "更新用户", updatedUser.Name)
		}

		if updatedUser.Age != 30 {
			t.Errorf("年龄未更新: 期望 %d, 实际 %d", 30, updatedUser.Age)
		}

		if updatedUser.CreatedAt != originalCreatedAt {
			t.Error("CreatedAt不应该改变")
		}

		if updatedUser.UpdatedAt.Equal(originalCreatedAt) {
			t.Error("UpdatedAt应该改变")
		}

		t.Log("更新用户测试通过")
	})

	t.Run("UpdateNonExistentUser", func(t *testing.T) {
		user := &SimpleUser{
			ID:    9999,
			Name:  "不存在用户",
			Email: "nonexistent@example.com",
			Age:   25,
		}

		err := dm.UpdateUser(user)
		if err == nil {
			t.Error("更新不存在的用户应该返回错误")
		}

		t.Log("更新不存在用户测试通过")
	})
}

func TestSimpleDatabaseManager_DeleteUser(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("DeleteExistingUser", func(t *testing.T) {
		// 创建用户
		user := &SimpleUser{
			Name:  "待删除用户",
			Email: "delete@example.com",
			Age:   25,
		}

		err := dm.CreateUser(user)
		if err != nil {
			t.Fatalf("创建用户失败: %v", err)
		}

		// 删除用户
		err = dm.DeleteUser(user.ID)
		if err != nil {
			t.Errorf("删除用户失败: %v", err)
		}

		// 验证删除
		_, err = dm.GetUserByID(user.ID)
		if err == nil {
			t.Error("删除后用户仍然存在")
		}

		t.Log("删除用户测试通过")
	})

	t.Run("DeleteNonExistentUser", func(t *testing.T) {
		err := dm.DeleteUser(9999)
		if err == nil {
			t.Error("删除不存在的用户应该返回错误")
		}

		t.Log("删除不存在用户测试通过")
	})
}

func TestSimpleDatabaseManager_CategoryOperations(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("CreateAndGetCategories", func(t *testing.T) {
		// 获取初始分类数量
		initialCategories := dm.GetAllCategories()
		initialCount := len(initialCategories)

		// 创建新分类
		newCategory := &SimpleCategory{
			Name:        "测试分类",
			Description: "这是一个测试分类",
		}

		err := dm.CreateCategory(newCategory)
		if err != nil {
			t.Errorf("创建分类失败: %v", err)
		}

		if newCategory.ID == 0 {
			t.Error("分类ID应该被设置")
		}

		// 获取所有分类
		allCategories := dm.GetAllCategories()

		if len(allCategories) != initialCount+1 {
			t.Errorf("分类数量不正确: 期望 %d, 实际 %d", initialCount+1, len(allCategories))
		}

		t.Log("分类操作测试通过")
	})
}

func TestSimpleDatabaseManager_ProductOperations(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("CreateAndGetProducts", func(t *testing.T) {
		// 创建产品
		product := &SimpleProduct{
			Name:        "测试产品",
			Description: "这是一个测试产品",
			Price:       99.99,
			CategoryID:  1, // 使用种子数据中的分类
			Stock:       50,
		}

		err := dm.CreateProduct(product)
		if err != nil {
			t.Errorf("创建产品失败: %v", err)
		}

		if product.ID == 0 {
			t.Error("产品ID应该被设置")
		}

		// 根据分类获取产品
		products := dm.GetProductsByCategory(1)

		// 检查是否包含我们创建的产品
		found := false
		for _, p := range products {
			if p.ID == product.ID {
				found = true
				break
			}
		}

		if !found {
			t.Error("创建的产品未在分类中找到")
		}

		t.Log("产品操作测试通过")
	})
}

func TestSimpleDatabaseManager_GetProductsWithCategory(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("GetProductsWithCategory", func(t *testing.T) {
		products := dm.GetProductsWithCategory()

		if len(products) == 0 {
			t.Error("应该有产品数据")
		}

		// 检查第一个产品是否有分类名称
		if products[0].CategoryName == "" {
			t.Error("产品应该有分类名称")
		}

		t.Log("产品分类联合查询测试通过")
	})
}

func TestSimpleDatabaseManager_TransferStock(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("SuccessfulTransfer", func(t *testing.T) {
		// 获取初始库存
		products := dm.GetProductsByCategory(1)
		if len(products) < 2 {
			t.Skip("需要至少2个产品进行库存转移测试")
		}

		fromProduct := products[0]
		toProduct := products[1]

		initialFromStock := fromProduct.Stock
		initialToStock := toProduct.Stock
		transferQuantity := 5

		if initialFromStock < transferQuantity {
			t.Skip("源产品库存不足进行转移测试")
		}

		// 执行库存转移
		err := dm.TransferStock(fromProduct.ID, toProduct.ID, transferQuantity)
		if err != nil {
			t.Errorf("库存转移失败: %v", err)
		}

		// 验证库存变化
		updatedProducts := dm.GetProductsByCategory(1)

		var updatedFromProduct, updatedToProduct SimpleProduct
		for _, p := range updatedProducts {
			if p.ID == fromProduct.ID {
				updatedFromProduct = p
			}
			if p.ID == toProduct.ID {
				updatedToProduct = p
			}
		}

		expectedFromStock := initialFromStock - transferQuantity
		expectedToStock := initialToStock + transferQuantity

		if updatedFromProduct.Stock != expectedFromStock {
			t.Errorf("源产品库存不正确: 期望 %d, 实际 %d", expectedFromStock, updatedFromProduct.Stock)
		}

		if updatedToProduct.Stock != expectedToStock {
			t.Errorf("目标产品库存不正确: 期望 %d, 实际 %d", expectedToStock, updatedToProduct.Stock)
		}

		t.Log("库存转移测试通过")
	})

	t.Run("InsufficientStock", func(t *testing.T) {
		// 尝试转移超过库存的数量
		err := dm.TransferStock(1, 2, 99999)
		if err == nil {
			t.Error("库存不足时应该返回错误")
		}

		t.Log("库存不足测试通过")
	})
}

func TestSimpleDatabaseManager_GetCategoryStats(t *testing.T) {
	dm := setupSimpleTestDB(t)

	t.Run("GetCategoryStats", func(t *testing.T) {
		stats := dm.GetCategoryStats()

		if len(stats) == 0 {
			t.Error("应该有分类统计数据")
		}

		// 检查统计数据的合理性
		for _, stat := range stats {
			if stat.CategoryName == "" {
				t.Error("分类名称不应该为空")
			}

			if stat.ProductCount < 0 {
				t.Error("产品数量不应该为负数")
			}

			if stat.TotalStock < 0 {
				t.Error("总库存不应该为负数")
			}

			if stat.AvgPrice < 0 {
				t.Error("平均价格不应该为负数")
			}
		}

		t.Log("分类统计测试通过")
	})
}

// 基准测试
func BenchmarkSimpleDatabaseManager_CreateUser(b *testing.B) {
	dm := setupSimpleTestDB(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := &SimpleUser{
			Name:  fmt.Sprintf("基准用户%d", i),
			Email: fmt.Sprintf("bench%d@example.com", i),
			Age:   25,
		}
		dm.CreateUser(user)
	}
}

func BenchmarkSimpleDatabaseManager_GetUserByID(b *testing.B) {
	dm := setupSimpleTestDB(b)

	// 创建一个用户用于查询
	user := &SimpleUser{
		Name:  "基准查询用户",
		Email: "benchquery@example.com",
		Age:   25,
	}
	dm.CreateUser(user)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dm.GetUserByID(user.ID)
	}
}

func BenchmarkSimpleDatabaseManager_GetAllUsers(b *testing.B) {
	dm := setupSimpleTestDB(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dm.GetAllUsers()
	}
}
