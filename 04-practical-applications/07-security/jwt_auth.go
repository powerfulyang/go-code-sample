package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// JWT结构体
type JWT struct {
	Header    JWTHeader  `json:"header"`
	Payload   JWTPayload `json:"payload"`
	Signature string     `json:"signature"`
}

// JWT头部
type JWTHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

// JWT载荷
type JWTPayload struct {
	Subject   string                 `json:"sub"`              // 主题
	Issuer    string                 `json:"iss"`              // 签发者
	Audience  string                 `json:"aud"`              // 受众
	ExpiresAt int64                  `json:"exp"`              // 过期时间
	NotBefore int64                  `json:"nbf"`              // 生效时间
	IssuedAt  int64                  `json:"iat"`              // 签发时间
	JwtID     string                 `json:"jti"`              // JWT ID
	Claims    map[string]interface{} `json:"claims,omitempty"` // 自定义声明
}

// JWTManager JWT管理器
type JWTManager struct {
	secretKey []byte
	issuer    string
	expiry    time.Duration
}

// NewJWTManager 创建JWT管理器
func NewJWTManager(secretKey, issuer string, expiry time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: []byte(secretKey),
		issuer:    issuer,
		expiry:    expiry,
	}
}

// GenerateToken 生成JWT令牌
func (j *JWTManager) GenerateToken(subject string, claims map[string]interface{}) (string, error) {
	now := time.Now()

	// 创建头部
	header := JWTHeader{
		Algorithm: "HS256",
		Type:      "JWT",
	}

	// 创建载荷
	payload := JWTPayload{
		Subject:   subject,
		Issuer:    j.issuer,
		ExpiresAt: now.Add(j.expiry).Unix(),
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
		JwtID:     generateJTI(),
		Claims:    claims,
	}

	// 编码头部和载荷
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("编码头部失败: %v", err)
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("编码载荷失败: %v", err)
	}

	// Base64编码
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerBytes)
	payloadEncoded := base64.RawURLEncoding.EncodeToString(payloadBytes)

	// 创建签名
	message := headerEncoded + "." + payloadEncoded
	signature := j.createSignature(message)

	// 组合JWT
	token := message + "." + signature

	return token, nil
}

// ValidateToken 验证JWT令牌
func (j *JWTManager) ValidateToken(tokenString string) (*JWTPayload, error) {
	// 分割JWT
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("无效的JWT格式")
	}

	headerEncoded := parts[0]
	payloadEncoded := parts[1]
	signature := parts[2]

	// 验证签名
	message := headerEncoded + "." + payloadEncoded
	expectedSignature := j.createSignature(message)

	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return nil, fmt.Errorf("JWT签名验证失败")
	}

	// 解码载荷
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadEncoded)
	if err != nil {
		return nil, fmt.Errorf("解码载荷失败: %v", err)
	}

	var payload JWTPayload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("解析载荷失败: %v", err)
	}

	// 验证时间
	now := time.Now().Unix()

	if payload.ExpiresAt < now {
		return nil, fmt.Errorf("JWT已过期")
	}

	if payload.NotBefore > now {
		return nil, fmt.Errorf("JWT尚未生效")
	}

	return &payload, nil
}

// RefreshToken 刷新令牌
func (j *JWTManager) RefreshToken(tokenString string) (string, error) {
	payload, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("无法刷新无效令牌: %v", err)
	}

	// 等待一毫秒确保时间戳不同
	time.Sleep(time.Millisecond)

	// 生成新令牌
	return j.GenerateToken(payload.Subject, payload.Claims)
}

// createSignature 创建签名
func (j *JWTManager) createSignature(message string) string {
	h := hmac.New(sha256.New, j.secretKey)
	h.Write([]byte(message))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// generateJTI 生成JWT ID
func generateJTI() string {
	return fmt.Sprintf("%d_%d", time.Now().UnixNano(), time.Now().Unix())
}

// User 用户结构体
type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // 不序列化密码
	Roles    []string  `json:"roles"`
	Active   bool      `json:"active"`
	Created  time.Time `json:"created"`
}

// AuthService 认证服务
type AuthService struct {
	jwtManager *JWTManager
	users      map[string]*User // 简单的内存存储
	userID     int
}

// NewAuthService 创建认证服务
func NewAuthService(jwtManager *JWTManager) *AuthService {
	service := &AuthService{
		jwtManager: jwtManager,
		users:      make(map[string]*User),
		userID:     1,
	}

	// 添加示例用户
	service.createSampleUsers()

	return service
}

// createSampleUsers 创建示例用户
func (a *AuthService) createSampleUsers() {
	users := []struct {
		username string
		email    string
		password string
		roles    []string
		active   bool
	}{
		{
			username: "admin",
			email:    "admin@example.com",
			password: "admin123",
			roles:    []string{"admin", "user"},
			active:   true,
		},
		{
			username: "user1",
			email:    "user1@example.com",
			password: "user123",
			roles:    []string{"user"},
			active:   true,
		},
		{
			username: "user2",
			email:    "user2@example.com",
			password: "user456",
			roles:    []string{"user"},
			active:   false,
		},
	}

	for _, userData := range users {
		user := &User{
			ID:       a.userID,
			Username: userData.username,
			Email:    userData.email,
			Password: hashPassword(userData.password),
			Roles:    userData.roles,
			Active:   userData.active,
			Created:  time.Now(),
		}

		a.users[userData.username] = user
		a.userID++
	}
}

// RegisterUser 注册用户
func (a *AuthService) RegisterUser(username, email, password string, roles []string) (*User, error) {
	// 检查用户是否已存在
	if _, exists := a.users[username]; exists {
		return nil, fmt.Errorf("用户名已存在: %s", username)
	}

	// 创建用户
	user := &User{
		ID:       a.userID,
		Username: username,
		Email:    email,
		Password: hashPassword(password),
		Roles:    roles,
		Active:   true,
		Created:  time.Now(),
	}

	a.users[username] = user
	a.userID++

	return user, nil
}

// Login 用户登录
func (a *AuthService) Login(username, password string) (string, *User, error) {
	// 查找用户
	user, exists := a.users[username]
	if !exists {
		return "", nil, fmt.Errorf("用户不存在: %s", username)
	}

	// 检查用户状态
	if !user.Active {
		return "", nil, fmt.Errorf("用户已被禁用: %s", username)
	}

	// 验证密码
	if !verifyPassword(password, user.Password) {
		return "", nil, fmt.Errorf("密码错误")
	}

	// 生成JWT令牌
	claims := map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
		"roles":   user.Roles,
	}

	token, err := a.jwtManager.GenerateToken(username, claims)
	if err != nil {
		return "", nil, fmt.Errorf("生成令牌失败: %v", err)
	}

	return token, user, nil
}

// ValidateToken 验证令牌
func (a *AuthService) ValidateToken(tokenString string) (*User, error) {
	payload, err := a.jwtManager.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// 查找用户
	user, exists := a.users[payload.Subject]
	if !exists {
		return nil, fmt.Errorf("用户不存在: %s", payload.Subject)
	}

	// 检查用户状态
	if !user.Active {
		return nil, fmt.Errorf("用户已被禁用: %s", payload.Subject)
	}

	return user, nil
}

// HasRole 检查用户是否有指定角色
func (a *AuthService) HasRole(user *User, role string) bool {
	for _, userRole := range user.Roles {
		if userRole == role {
			return true
		}
	}
	return false
}

// ChangePassword 修改密码
func (a *AuthService) ChangePassword(username, oldPassword, newPassword string) error {
	user, exists := a.users[username]
	if !exists {
		return fmt.Errorf("用户不存在: %s", username)
	}

	// 验证旧密码
	if !verifyPassword(oldPassword, user.Password) {
		return fmt.Errorf("旧密码错误")
	}

	// 更新密码
	user.Password = hashPassword(newPassword)

	return nil
}

// DeactivateUser 禁用用户
func (a *AuthService) DeactivateUser(username string) error {
	user, exists := a.users[username]
	if !exists {
		return fmt.Errorf("用户不存在: %s", username)
	}

	user.Active = false
	return nil
}

// GetUserInfo 获取用户信息
func (a *AuthService) GetUserInfo(username string) (*User, error) {
	user, exists := a.users[username]
	if !exists {
		return nil, fmt.Errorf("用户不存在: %s", username)
	}

	// 返回用户副本（不包含密码）
	userCopy := *user
	userCopy.Password = ""

	return &userCopy, nil
}

// ListUsers 列出所有用户
func (a *AuthService) ListUsers() []*User {
	var users []*User
	for _, user := range a.users {
		userCopy := *user
		userCopy.Password = "" // 不返回密码
		users = append(users, &userCopy)
	}
	return users
}

// 简单的密码哈希（实际项目中应使用bcrypt）
func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password + "salt")) // 简单加盐
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 验证密码
func verifyPassword(password, hash string) bool {
	return hashPassword(password) == hash
}

// JWTExamples JWT认证示例
func JWTExamples() {
	fmt.Println("=== JWT认证示例 ===")
	fmt.Println()
	fmt.Println("JWT (JSON Web Token) 特点:")
	fmt.Println("- 无状态认证")
	fmt.Println("- 跨域支持")
	fmt.Println("- 自包含信息")
	fmt.Println("- 安全可靠")
	fmt.Println("- 易于扩展")
	fmt.Println()

	// 创建JWT管理器
	jwtManager := NewJWTManager("my-secret-key", "my-app", 24*time.Hour)

	// 创建认证服务
	authService := NewAuthService(jwtManager)

	fmt.Println("🔹 用户注册和登录演示:")

	// 注册新用户
	newUser, err := authService.RegisterUser("testuser", "test@example.com", "password123", []string{"user"})
	if err != nil {
		fmt.Printf("注册失败: %v\n", err)
	} else {
		fmt.Printf("✅ 用户注册成功: %s (ID: %d)\n", newUser.Username, newUser.ID)
	}

	// 用户登录
	token, user, err := authService.Login("admin", "admin123")
	if err != nil {
		fmt.Printf("登录失败: %v\n", err)
	} else {
		fmt.Printf("✅ 登录成功: %s\n", user.Username)
		fmt.Printf("JWT令牌: %s...\n", token[:50])
	}

	// 验证令牌
	fmt.Println("\n🔹 令牌验证演示:")
	validatedUser, err := authService.ValidateToken(token)
	if err != nil {
		fmt.Printf("令牌验证失败: %v\n", err)
	} else {
		fmt.Printf("✅ 令牌验证成功: %s\n", validatedUser.Username)
		fmt.Printf("用户角色: %v\n", validatedUser.Roles)
	}

	// 角色检查
	fmt.Println("\n🔹 权限检查演示:")
	if authService.HasRole(validatedUser, "admin") {
		fmt.Println("✅ 用户具有管理员权限")
	} else {
		fmt.Println("❌ 用户没有管理员权限")
	}

	// 列出所有用户
	fmt.Println("\n🔹 用户管理演示:")
	users := authService.ListUsers()
	fmt.Printf("系统用户数量: %d\n", len(users))
	for _, u := range users {
		status := "激活"
		if !u.Active {
			status = "禁用"
		}
		fmt.Printf("  - %s (%s) [%s] - 角色: %v\n",
			u.Username, u.Email, status, u.Roles)
	}

	fmt.Println("\n✅ JWT认证示例演示完成!")
	fmt.Println("💡 提示: 实际项目中应使用bcrypt进行密码哈希")
	fmt.Println("💡 提示: 建议使用HTTPS传输JWT令牌")
	fmt.Println("💡 提示: 可以集成OAuth2.0进行第三方登录")
}
