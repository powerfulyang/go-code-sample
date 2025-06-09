package security

import (
	"strings"
	"testing"
	"time"
)

func TestJWTManager(t *testing.T) {
	jwtManager := NewJWTManager("test-secret-key", "test-issuer", 1*time.Hour)

	t.Run("GenerateAndValidateToken", func(t *testing.T) {
		subject := "testuser"
		claims := map[string]interface{}{
			"user_id": 123,
			"role":    "admin",
		}

		// 生成令牌
		token, err := jwtManager.GenerateToken(subject, claims)
		if err != nil {
			t.Fatalf("生成令牌失败: %v", err)
		}

		if token == "" {
			t.Error("生成的令牌不应该为空")
		}

		// 验证令牌
		payload, err := jwtManager.ValidateToken(token)
		if err != nil {
			t.Fatalf("验证令牌失败: %v", err)
		}

		if payload.Subject != subject {
			t.Errorf("主题不匹配: 期望 %s, 实际 %s", subject, payload.Subject)
		}

		if payload.Issuer != "test-issuer" {
			t.Errorf("签发者不匹配: 期望 test-issuer, 实际 %s", payload.Issuer)
		}

		if payload.Claims["user_id"] != float64(123) { // JSON数字解析为float64
			t.Errorf("用户ID不匹配: 期望 123, 实际 %v", payload.Claims["user_id"])
		}

		t.Log("JWT生成和验证测试通过")
	})

	t.Run("InvalidToken", func(t *testing.T) {
		invalidToken := "invalid.token.here"
		_, err := jwtManager.ValidateToken(invalidToken)
		if err == nil {
			t.Error("无效令牌应该验证失败")
		}

		t.Log("无效令牌测试通过")
	})

	t.Run("ExpiredToken", func(t *testing.T) {
		// 创建一个已过期的JWT管理器
		expiredManager := NewJWTManager("test-secret-key", "test-issuer", -1*time.Hour)
		
		token, err := expiredManager.GenerateToken("testuser", nil)
		if err != nil {
			t.Fatalf("生成过期令牌失败: %v", err)
		}

		// 验证过期令牌
		_, err = jwtManager.ValidateToken(token)
		if err == nil {
			t.Error("过期令牌应该验证失败")
		}

		if !strings.Contains(err.Error(), "过期") {
			t.Errorf("错误信息应该包含'过期': %v", err)
		}

		t.Log("过期令牌测试通过")
	})

	t.Run("RefreshToken", func(t *testing.T) {
		// 生成原始令牌
		originalToken, err := jwtManager.GenerateToken("testuser", map[string]interface{}{"role": "user"})
		if err != nil {
			t.Fatalf("生成原始令牌失败: %v", err)
		}

		// 刷新令牌
		newToken, err := jwtManager.RefreshToken(originalToken)
		if err != nil {
			t.Fatalf("刷新令牌失败: %v", err)
		}

		if newToken == originalToken {
			t.Error("刷新后的令牌应该与原始令牌不同")
		}

		// 验证新令牌
		payload, err := jwtManager.ValidateToken(newToken)
		if err != nil {
			t.Fatalf("验证刷新令牌失败: %v", err)
		}

		if payload.Subject != "testuser" {
			t.Errorf("刷新令牌主题不正确: %s", payload.Subject)
		}

		t.Log("刷新令牌测试通过")
	})
}

func TestAuthService(t *testing.T) {
	jwtManager := NewJWTManager("test-secret", "test-app", 1*time.Hour)
	authService := NewAuthService(jwtManager)

	t.Run("RegisterUser", func(t *testing.T) {
		user, err := authService.RegisterUser("newuser", "new@example.com", "password123", []string{"user"})
		if err != nil {
			t.Fatalf("注册用户失败: %v", err)
		}

		if user.Username != "newuser" {
			t.Errorf("用户名不正确: %s", user.Username)
		}

		if user.Email != "new@example.com" {
			t.Errorf("邮箱不正确: %s", user.Email)
		}

		if !user.Active {
			t.Error("新用户应该是激活状态")
		}

		t.Log("用户注册测试通过")
	})

	t.Run("DuplicateUser", func(t *testing.T) {
		_, err := authService.RegisterUser("admin", "admin2@example.com", "password", []string{"user"})
		if err == nil {
			t.Error("重复用户名应该注册失败")
		}

		t.Log("重复用户测试通过")
	})

	t.Run("Login", func(t *testing.T) {
		token, user, err := authService.Login("admin", "admin123")
		if err != nil {
			t.Fatalf("登录失败: %v", err)
		}

		if token == "" {
			t.Error("登录应该返回令牌")
		}

		if user.Username != "admin" {
			t.Errorf("登录用户名不正确: %s", user.Username)
		}

		t.Log("用户登录测试通过")
	})

	t.Run("InvalidLogin", func(t *testing.T) {
		_, _, err := authService.Login("admin", "wrongpassword")
		if err == nil {
			t.Error("错误密码应该登录失败")
		}

		_, _, err = authService.Login("nonexistent", "password")
		if err == nil {
			t.Error("不存在的用户应该登录失败")
		}

		t.Log("无效登录测试通过")
	})

	t.Run("ValidateToken", func(t *testing.T) {
		token, _, err := authService.Login("admin", "admin123")
		if err != nil {
			t.Fatalf("登录失败: %v", err)
		}

		user, err := authService.ValidateToken(token)
		if err != nil {
			t.Fatalf("令牌验证失败: %v", err)
		}

		if user.Username != "admin" {
			t.Errorf("验证用户名不正确: %s", user.Username)
		}

		t.Log("令牌验证测试通过")
	})

	t.Run("HasRole", func(t *testing.T) {
		user, err := authService.GetUserInfo("admin")
		if err != nil {
			t.Fatalf("获取用户信息失败: %v", err)
		}

		if !authService.HasRole(user, "admin") {
			t.Error("admin用户应该有admin角色")
		}

		if !authService.HasRole(user, "user") {
			t.Error("admin用户应该有user角色")
		}

		if authService.HasRole(user, "nonexistent") {
			t.Error("admin用户不应该有不存在的角色")
		}

		t.Log("角色检查测试通过")
	})

	t.Run("ChangePassword", func(t *testing.T) {
		err := authService.ChangePassword("admin", "admin123", "newpassword123")
		if err != nil {
			t.Fatalf("修改密码失败: %v", err)
		}

		// 用新密码登录
		_, _, err = authService.Login("admin", "newpassword123")
		if err != nil {
			t.Fatalf("新密码登录失败: %v", err)
		}

		// 用旧密码登录应该失败
		_, _, err = authService.Login("admin", "admin123")
		if err == nil {
			t.Error("旧密码应该登录失败")
		}

		t.Log("修改密码测试通过")
	})
}

func TestEncryptionManager(t *testing.T) {
	encManager := NewEncryptionManager("test-key")

	t.Run("AESEncryptDecrypt", func(t *testing.T) {
		plaintext := "这是需要加密的测试数据"

		// 加密
		encrypted, err := encManager.AESEncrypt(plaintext)
		if err != nil {
			t.Fatalf("AES加密失败: %v", err)
		}

		if encrypted == "" {
			t.Error("加密结果不应该为空")
		}

		if encrypted == plaintext {
			t.Error("加密结果不应该等于原文")
		}

		// 解密
		decrypted, err := encManager.AESDecrypt(encrypted)
		if err != nil {
			t.Fatalf("AES解密失败: %v", err)
		}

		if decrypted != plaintext {
			t.Errorf("解密结果不正确: 期望 %s, 实际 %s", plaintext, decrypted)
		}

		t.Log("AES加密解密测试通过")
	})

	t.Run("InvalidDecryption", func(t *testing.T) {
		_, err := encManager.AESDecrypt("invalid-ciphertext")
		if err == nil {
			t.Error("无效密文应该解密失败")
		}

		t.Log("无效解密测试通过")
	})
}

func TestRSAKeyPair(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(1024) // 使用较小的密钥长度加快测试
	if err != nil {
		t.Fatalf("生成RSA密钥对失败: %v", err)
	}

	t.Run("RSAEncryptDecrypt", func(t *testing.T) {
		plaintext := "RSA加密测试消息"

		// 加密
		encrypted, err := keyPair.RSAEncrypt(plaintext)
		if err != nil {
			t.Fatalf("RSA加密失败: %v", err)
		}

		if encrypted == "" {
			t.Error("RSA加密结果不应该为空")
		}

		// 解密
		decrypted, err := keyPair.RSADecrypt(encrypted)
		if err != nil {
			t.Fatalf("RSA解密失败: %v", err)
		}

		if decrypted != plaintext {
			t.Errorf("RSA解密结果不正确: 期望 %s, 实际 %s", plaintext, decrypted)
		}

		t.Log("RSA加密解密测试通过")
	})

	t.Run("RSASignVerify", func(t *testing.T) {
		message := "需要签名的消息"

		// 签名
		signature, err := keyPair.RSASign(message)
		if err != nil {
			t.Fatalf("RSA签名失败: %v", err)
		}

		if signature == "" {
			t.Error("RSA签名不应该为空")
		}

		// 验证签名
		err = keyPair.RSAVerify(message, signature)
		if err != nil {
			t.Fatalf("RSA签名验证失败: %v", err)
		}

		// 验证错误消息的签名应该失败
		err = keyPair.RSAVerify("错误的消息", signature)
		if err == nil {
			t.Error("错误消息的签名验证应该失败")
		}

		t.Log("RSA签名验证测试通过")
	})

	t.Run("ExportKeys", func(t *testing.T) {
		// 导出私钥
		privateKeyPEM, err := keyPair.ExportPrivateKeyPEM()
		if err != nil {
			t.Fatalf("导出私钥失败: %v", err)
		}

		if !strings.Contains(privateKeyPEM, "PRIVATE KEY") {
			t.Error("私钥PEM格式不正确")
		}

		// 导出公钥
		publicKeyPEM, err := keyPair.ExportPublicKeyPEM()
		if err != nil {
			t.Fatalf("导出公钥失败: %v", err)
		}

		if !strings.Contains(publicKeyPEM, "PUBLIC KEY") {
			t.Error("公钥PEM格式不正确")
		}

		t.Log("密钥导出测试通过")
	})
}

func TestHashManager(t *testing.T) {
	hashManager := NewHashManager()

	t.Run("HashFunctions", func(t *testing.T) {
		data := "测试数据"

		md5Hash := hashManager.MD5Hash(data)
		if len(md5Hash) != 32 {
			t.Errorf("MD5哈希长度不正确: %d", len(md5Hash))
		}

		sha256Hash := hashManager.SHA256Hash(data)
		if len(sha256Hash) != 64 {
			t.Errorf("SHA256哈希长度不正确: %d", len(sha256Hash))
		}

		hmacHash := hashManager.HMACSHA256(data, "secret")
		if len(hmacHash) != 64 {
			t.Errorf("HMAC-SHA256哈希长度不正确: %d", len(hmacHash))
		}

		// 相同数据应该产生相同哈希
		if hashManager.MD5Hash(data) != md5Hash {
			t.Error("相同数据的MD5哈希应该相同")
		}

		t.Log("哈希函数测试通过")
	})

	t.Run("SecureRandomString", func(t *testing.T) {
		length := 16
		randomStr, err := hashManager.SecureRandomString(length)
		if err != nil {
			t.Fatalf("生成随机字符串失败: %v", err)
		}

		if len(randomStr) != length {
			t.Errorf("随机字符串长度不正确: 期望 %d, 实际 %d", length, len(randomStr))
		}

		// 生成两个随机字符串应该不同
		randomStr2, err := hashManager.SecureRandomString(length)
		if err != nil {
			t.Fatalf("生成第二个随机字符串失败: %v", err)
		}

		if randomStr == randomStr2 {
			t.Error("两个随机字符串不应该相同")
		}

		t.Log("安全随机字符串测试通过")
	})
}

func TestPasswordManager(t *testing.T) {
	passwordManager := NewPasswordManager()

	t.Run("HashAndVerifyPassword", func(t *testing.T) {
		password := "testpassword123"

		// 哈希密码
		hashedPassword, err := passwordManager.HashPassword(password)
		if err != nil {
			t.Fatalf("密码哈希失败: %v", err)
		}

		if hashedPassword == password {
			t.Error("哈希密码不应该等于原密码")
		}

		// 验证正确密码
		if !passwordManager.VerifyPassword(password, hashedPassword) {
			t.Error("正确密码验证失败")
		}

		// 验证错误密码
		if passwordManager.VerifyPassword("wrongpassword", hashedPassword) {
			t.Error("错误密码验证应该失败")
		}

		t.Log("密码哈希和验证测试通过")
	})

	t.Run("GenerateSecurePassword", func(t *testing.T) {
		length := 12
		password, err := passwordManager.GenerateSecurePassword(length)
		if err != nil {
			t.Fatalf("生成安全密码失败: %v", err)
		}

		if len(password) != length {
			t.Errorf("生成密码长度不正确: 期望 %d, 实际 %d", length, len(password))
		}

		// 检查是否包含不同类型的字符
		hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
		hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		hasDigit := strings.ContainsAny(password, "0123456789")
		hasSpecial := strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?")

		if !hasLower || !hasUpper || !hasDigit || !hasSpecial {
			t.Errorf("生成的密码应该包含各种字符类型: %s", password)
		}

		t.Log("安全密码生成测试通过")
	})
}

// 基准测试
func BenchmarkJWTGenerate(b *testing.B) {
	jwtManager := NewJWTManager("benchmark-key", "benchmark-app", 1*time.Hour)
	claims := map[string]interface{}{"user_id": 123, "role": "user"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jwtManager.GenerateToken("testuser", claims)
	}
}

func BenchmarkJWTValidate(b *testing.B) {
	jwtManager := NewJWTManager("benchmark-key", "benchmark-app", 1*time.Hour)
	token, _ := jwtManager.GenerateToken("testuser", map[string]interface{}{"user_id": 123})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jwtManager.ValidateToken(token)
	}
}

func BenchmarkAESEncrypt(b *testing.B) {
	encManager := NewEncryptionManager("benchmark-key")
	plaintext := "这是用于基准测试的数据"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encManager.AESEncrypt(plaintext)
	}
}

func BenchmarkPasswordHash(b *testing.B) {
	passwordManager := NewPasswordManager()
	password := "benchmarkpassword123"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = passwordManager.HashPassword(password)
	}
}
