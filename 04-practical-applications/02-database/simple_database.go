package database

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// SimpleUser 简化的用户模型
type SimpleUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SimpleProduct 简化的产品模型
type SimpleProduct struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  int       `json:"category_id"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SimpleCategory 简化的分类模型
type SimpleCategory struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SimpleDatabaseManager 简化的数据库管理器（内存实现）
type SimpleDatabaseManager struct {
	users      map[int]*SimpleUser
	categories map[int]*SimpleCategory
	products   map[int]*SimpleProduct
	nextUserID int
	nextCatID  int
	nextProdID int
	mutex      sync.RWMutex
}

// NewSimpleDatabaseManager 创建简化的数据库管理器
func NewSimpleDatabaseManager() *SimpleDatabaseManager {
	dm := &SimpleDatabaseManager{
		users:      make(map[int]*SimpleUser),
		categories: make(map[int]*SimpleCategory),
		products:   make(map[int]*SimpleProduct),
		nextUserID: 1,
		nextCatID:  1,
		nextProdID: 1,
	}
	
	// 初始化示例数据
	dm.seedData()
	
	return dm
}

// seedData 初始化示例数据
func (dm *SimpleDatabaseManager) seedData() {
	// 添加分类
	categories := []*SimpleCategory{
		{Name: "电子产品", Description: "各种电子设备和配件"},
		{Name: "服装", Description: "男女服装和配饰"},
		{Name: "图书", Description: "各类图书和杂志"},
	}
	
	for _, cat := range categories {
		dm.CreateCategory(cat)
	}
	
	// 添加用户
	users := []*SimpleUser{
		{Name: "张三", Email: "zhangsan@example.com", Age: 25},
		{Name: "李四", Email: "lisi@example.com", Age: 30},
		{Name: "王五", Email: "wangwu@example.com", Age: 28},
	}
	
	for _, user := range users {
		dm.CreateUser(user)
	}
	
	// 添加产品
	products := []*SimpleProduct{
		{Name: "iPhone 15", Description: "最新款苹果手机", Price: 5999.00, CategoryID: 1, Stock: 100},
		{Name: "MacBook Pro", Description: "专业级笔记本电脑", Price: 12999.00, CategoryID: 1, Stock: 50},
		{Name: "连衣裙", Description: "时尚女装连衣裙", Price: 299.00, CategoryID: 2, Stock: 200},
		{Name: "Go语言编程", Description: "Go语言学习指南", Price: 89.00, CategoryID: 3, Stock: 150},
	}
	
	for _, product := range products {
		dm.CreateProduct(product)
	}
}

// 用户相关操作

// CreateUser 创建用户
func (dm *SimpleDatabaseManager) CreateUser(user *SimpleUser) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	
	user.ID = dm.nextUserID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	
	dm.users[user.ID] = user
	dm.nextUserID++
	
	return nil
}

// GetUserByID 根据ID获取用户
func (dm *SimpleDatabaseManager) GetUserByID(id int) (*SimpleUser, error) {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	
	user, exists := dm.users[id]
	if !exists {
		return nil, fmt.Errorf("用户不存在: %d", id)
	}
	
	// 返回副本
	userCopy := *user
	return &userCopy, nil
}

// GetAllUsers 获取所有用户
func (dm *SimpleDatabaseManager) GetAllUsers() []SimpleUser {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	
	users := make([]SimpleUser, 0, len(dm.users))
	for _, user := range dm.users {
		users = append(users, *user)
	}
	
	// 按创建时间排序
	sort.Slice(users, func(i, j int) bool {
		return users[i].CreatedAt.After(users[j].CreatedAt)
	})
	
	return users
}

// UpdateUser 更新用户
func (dm *SimpleDatabaseManager) UpdateUser(user *SimpleUser) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	
	existing, exists := dm.users[user.ID]
	if !exists {
		return fmt.Errorf("用户不存在: %d", user.ID)
	}
	
	// 保留创建时间
	user.CreatedAt = existing.CreatedAt
	user.UpdatedAt = time.Now()
	
	dm.users[user.ID] = user
	return nil
}

// DeleteUser 删除用户
func (dm *SimpleDatabaseManager) DeleteUser(id int) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	
	if _, exists := dm.users[id]; !exists {
		return fmt.Errorf("用户不存在: %d", id)
	}
	
	delete(dm.users, id)
	return nil
}

// 分类相关操作

// CreateCategory 创建分类
func (dm *SimpleDatabaseManager) CreateCategory(category *SimpleCategory) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	
	category.ID = dm.nextCatID
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	
	dm.categories[category.ID] = category
	dm.nextCatID++
	
	return nil
}

// GetAllCategories 获取所有分类
func (dm *SimpleDatabaseManager) GetAllCategories() []SimpleCategory {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	
	categories := make([]SimpleCategory, 0, len(dm.categories))
	for _, category := range dm.categories {
		categories = append(categories, *category)
	}
	
	// 按名称排序
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Name < categories[j].Name
	})
	
	return categories
}

// 产品相关操作

// CreateProduct 创建产品
func (dm *SimpleDatabaseManager) CreateProduct(product *SimpleProduct) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	
	product.ID = dm.nextProdID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	
	dm.products[product.ID] = product
	dm.nextProdID++
	
	return nil
}

// GetProductsByCategory 根据分类获取产品
func (dm *SimpleDatabaseManager) GetProductsByCategory(categoryID int) []SimpleProduct {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	
	var products []SimpleProduct
	for _, product := range dm.products {
		if product.CategoryID == categoryID {
			products = append(products, *product)
		}
	}
	
	// 按名称排序
	sort.Slice(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})
	
	return products
}

// ProductWithCategory 产品和分类的联合查询结果
type ProductWithCategory struct {
	Product      SimpleProduct `json:"product"`
	CategoryName string        `json:"category_name"`
}

// GetProductsWithCategory 获取产品及其分类信息
func (dm *SimpleDatabaseManager) GetProductsWithCategory() []ProductWithCategory {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	
	var results []ProductWithCategory
	for _, product := range dm.products {
		categoryName := "未知分类"
		if category, exists := dm.categories[product.CategoryID]; exists {
			categoryName = category.Name
		}
		
		results = append(results, ProductWithCategory{
			Product:      *product,
			CategoryName: categoryName,
		})
	}
	
	// 按产品名称排序
	sort.Slice(results, func(i, j int) bool {
		return results[i].Product.Name < results[j].Product.Name
	})
	
	return results
}

// TransferStock 库存转移（事务示例）
func (dm *SimpleDatabaseManager) TransferStock(fromProductID, toProductID, quantity int) error {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	
	// 检查源产品
	fromProduct, exists := dm.products[fromProductID]
	if !exists {
		return fmt.Errorf("源产品不存在: %d", fromProductID)
	}
	
	// 检查目标产品
	toProduct, exists := dm.products[toProductID]
	if !exists {
		return fmt.Errorf("目标产品不存在: %d", toProductID)
	}
	
	// 检查库存
	if fromProduct.Stock < quantity {
		return fmt.Errorf("库存不足: 需要 %d, 可用 %d", quantity, fromProduct.Stock)
	}
	
	// 执行转移
	fromProduct.Stock -= quantity
	fromProduct.UpdatedAt = time.Now()
	
	toProduct.Stock += quantity
	toProduct.UpdatedAt = time.Now()
	
	return nil
}

// CategoryStats 分类统计信息
type CategoryStats struct {
	CategoryID   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
	ProductCount int     `json:"product_count"`
	TotalStock   int     `json:"total_stock"`
	AvgPrice     float64 `json:"avg_price"`
}

// GetCategoryStats 获取分类统计信息
func (dm *SimpleDatabaseManager) GetCategoryStats() []CategoryStats {
	dm.mutex.RLock()
	defer dm.mutex.RUnlock()
	
	var stats []CategoryStats
	
	for _, category := range dm.categories {
		stat := CategoryStats{
			CategoryID:   category.ID,
			CategoryName: category.Name,
			ProductCount: 0,
			TotalStock:   0,
			AvgPrice:     0,
		}
		
		var totalPrice float64
		for _, product := range dm.products {
			if product.CategoryID == category.ID {
				stat.ProductCount++
				stat.TotalStock += product.Stock
				totalPrice += product.Price
			}
		}
		
		if stat.ProductCount > 0 {
			stat.AvgPrice = totalPrice / float64(stat.ProductCount)
		}
		
		stats = append(stats, stat)
	}
	
	// 按分类名称排序
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].CategoryName < stats[j].CategoryName
	})
	
	return stats
}

// DatabaseExamples 数据库操作示例
func DatabaseExamples() {
	fmt.Println("=== 数据库操作示例 ===")
	
	// 创建数据库管理器
	dm := NewSimpleDatabaseManager()
	
	fmt.Println("✅ 内存数据库创建成功，示例数据已初始化")
	
	// 用户操作示例
	fmt.Println("\n🔹 用户操作示例")
	
	// 获取所有用户
	users := dm.GetAllUsers()
	fmt.Printf("用户总数: %d\n", len(users))
	for _, user := range users {
		fmt.Printf("  - %s (%s), 年龄: %d\n", user.Name, user.Email, user.Age)
	}
	
	// 创建新用户
	newUser := &SimpleUser{
		Name:  "新用户",
		Email: "newuser@example.com",
		Age:   26,
	}
	
	if err := dm.CreateUser(newUser); err != nil {
		fmt.Printf("创建用户失败: %v\n", err)
	} else {
		fmt.Printf("✅ 创建用户成功: ID=%d\n", newUser.ID)
	}
	
	// 分类和产品操作示例
	fmt.Println("\n🔹 分类和产品操作示例")
	
	// 获取所有分类
	categories := dm.GetAllCategories()
	fmt.Printf("分类总数: %d\n", len(categories))
	for _, category := range categories {
		fmt.Printf("  - %s: %s\n", category.Name, category.Description)
	}
	
	// 获取产品及分类信息
	fmt.Println("\n🔹 产品和分类联合查询")
	productsWithCategory := dm.GetProductsWithCategory()
	for _, pwc := range productsWithCategory {
		fmt.Printf("  - %s (%.2f元) - 分类: %s, 库存: %d\n",
			pwc.Product.Name, pwc.Product.Price, pwc.CategoryName, pwc.Product.Stock)
	}
	
	// 统计查询示例
	fmt.Println("\n🔹 分类统计信息")
	stats := dm.GetCategoryStats()
	for _, stat := range stats {
		fmt.Printf("  - %s: 产品数=%d, 总库存=%d, 平均价格=%.2f元\n",
			stat.CategoryName, stat.ProductCount, stat.TotalStock, stat.AvgPrice)
	}
	
	// 事务示例
	fmt.Println("\n🔹 事务操作示例")
	fmt.Println("尝试库存转移...")
	
	if err := dm.TransferStock(1, 2, 10); err != nil {
		fmt.Printf("库存转移失败: %v\n", err)
	} else {
		fmt.Println("✅ 库存转移成功")
		
		// 再次查看产品信息
		updatedProducts := dm.GetProductsWithCategory()
		for i, pwc := range updatedProducts {
			if i < 2 { // 只显示前两个产品
				fmt.Printf("  - %s: 库存=%d\n", pwc.Product.Name, pwc.Product.Stock)
			}
		}
	}
	
	fmt.Println("\n✅ 数据库操作示例演示完成!")
	fmt.Println("💡 提示: 这是一个内存数据库实现，用于演示数据库操作概念")
	fmt.Println("💡 在实际项目中，你可以使用真实的数据库如PostgreSQL、MySQL等")
}
