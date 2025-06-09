package security

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"strings"
)

// EncryptionManager 加密管理器
type EncryptionManager struct {
	aesKey []byte
}

// NewEncryptionManager 创建加密管理器
func NewEncryptionManager(key string) *EncryptionManager {
	// 使用SHA256生成32字节的AES密钥
	hash := sha256.Sum256([]byte(key))
	return &EncryptionManager{
		aesKey: hash[:],
	}
}

// AESEncrypt AES加密
func (e *EncryptionManager) AESEncrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(e.aesKey)
	if err != nil {
		return "", fmt.Errorf("创建AES密码器失败: %v", err)
	}

	// 使用GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("创建GCM失败: %v", err)
	}

	// 生成随机nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("生成nonce失败: %v", err)
	}

	// 加密
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Base64编码
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AESDecrypt AES解密
func (e *EncryptionManager) AESDecrypt(ciphertext string) (string, error) {
	// Base64解码
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("Base64解码失败: %v", err)
	}

	block, err := aes.NewCipher(e.aesKey)
	if err != nil {
		return "", fmt.Errorf("创建AES密码器失败: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("创建GCM失败: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("密文太短")
	}

	nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", fmt.Errorf("解密失败: %v", err)
	}

	return string(plaintext), nil
}

// RSAKeyPair RSA密钥对
type RSAKeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// GenerateRSAKeyPair 生成RSA密钥对
func GenerateRSAKeyPair(bits int) (*RSAKeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, fmt.Errorf("生成RSA密钥失败: %v", err)
	}

	return &RSAKeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}, nil
}

// RSAEncrypt RSA加密
func (kp *RSAKeyPair) RSAEncrypt(plaintext string) (string, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, kp.PublicKey, []byte(plaintext))
	if err != nil {
		return "", fmt.Errorf("RSA加密失败: %v", err)
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// RSADecrypt RSA解密
func (kp *RSAKeyPair) RSADecrypt(ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("Base64解码失败: %v", err)
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, kp.PrivateKey, data)
	if err != nil {
		return "", fmt.Errorf("RSA解密失败: %v", err)
	}

	return string(plaintext), nil
}

// RSASign RSA签名
func (kp *RSAKeyPair) RSASign(message string) (string, error) {
	hash := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, kp.PrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", fmt.Errorf("RSA签名失败: %v", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// RSAVerify RSA验证签名
func (kp *RSAKeyPair) RSAVerify(message, signature string) error {
	sig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("签名解码失败: %v", err)
	}

	hash := sha256.Sum256([]byte(message))
	return rsa.VerifyPKCS1v15(kp.PublicKey, crypto.SHA256, hash[:], sig)
}

// ExportPrivateKeyPEM 导出私钥为PEM格式
func (kp *RSAKeyPair) ExportPrivateKeyPEM() (string, error) {
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(kp.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("序列化私钥失败: %v", err)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return string(privateKeyPEM), nil
}

// ExportPublicKeyPEM 导出公钥为PEM格式
func (kp *RSAKeyPair) ExportPublicKeyPEM() (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(kp.PublicKey)
	if err != nil {
		return "", fmt.Errorf("序列化公钥失败: %v", err)
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return string(publicKeyPEM), nil
}

// HashManager 哈希管理器
type HashManager struct{}

// NewHashManager 创建哈希管理器
func NewHashManager() *HashManager {
	return &HashManager{}
}

// MD5Hash MD5哈希
func (h *HashManager) MD5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SHA256Hash SHA256哈希
func (h *HashManager) SHA256Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HMACSHA256 HMAC-SHA256
func (h *HashManager) HMACSHA256(data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// SecureRandomString 生成安全随机字符串
func (h *HashManager) SecureRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("生成随机字符串失败: %v", err)
	}

	// 转换为可打印字符
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}

	return string(bytes), nil
}

// PasswordManager 密码管理器
type PasswordManager struct {
	hashManager *HashManager
}

// NewPasswordManager 创建密码管理器
func NewPasswordManager() *PasswordManager {
	return &PasswordManager{
		hashManager: NewHashManager(),
	}
}

// HashPassword 哈希密码（简化版，实际应使用bcrypt）
func (p *PasswordManager) HashPassword(password string) (string, error) {
	// 生成盐值
	salt, err := p.hashManager.SecureRandomString(16)
	if err != nil {
		return "", err
	}

	// 组合密码和盐值
	saltedPassword := password + salt

	// 多次哈希增强安全性
	hash := saltedPassword
	for i := 0; i < 10000; i++ {
		hash = p.hashManager.SHA256Hash(hash)
	}

	// 返回盐值+哈希值
	return salt + ":" + hash, nil
}

// VerifyPassword 验证密码
func (p *PasswordManager) VerifyPassword(password, hashedPassword string) bool {
	parts := strings.Split(hashedPassword, ":")
	if len(parts) != 2 {
		return false
	}

	salt := parts[0]
	expectedHash := parts[1]

	// 使用相同的方法计算哈希
	saltedPassword := password + salt
	hash := saltedPassword
	for i := 0; i < 10000; i++ {
		hash = p.hashManager.SHA256Hash(hash)
	}

	return hash == expectedHash
}

// GenerateSecurePassword 生成安全密码
func (p *PasswordManager) GenerateSecurePassword(length int) (string, error) {
	if length < 8 {
		length = 8 // 最小长度
	}

	// 确保包含各种字符类型
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()_+-=[]{}|;:,.<>?"

	// 至少包含每种类型的一个字符
	password := ""
	password += string(lower[randomInt(len(lower))])
	password += string(upper[randomInt(len(upper))])
	password += string(digits[randomInt(len(digits))])
	password += string(special[randomInt(len(special))])

	// 填充剩余长度
	allChars := lower + upper + digits + special
	for len(password) < length {
		password += string(allChars[randomInt(len(allChars))])
	}

	// 打乱字符顺序
	return shuffleString(password), nil
}

// 辅助函数
func randomInt(max int) int {
	b := make([]byte, 1)
	rand.Read(b)
	return int(b[0]) % max
}

func shuffleString(s string) string {
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		j := randomInt(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// EncryptionExamples 加密示例
func EncryptionExamples() {
	fmt.Println("=== 加密和安全示例 ===")
	fmt.Println()

	// AES加密示例
	fmt.Println("🔹 AES对称加密演示:")
	encManager := NewEncryptionManager("my-secret-key")

	plaintext := "这是需要加密的敏感信息"
	fmt.Printf("原文: %s\n", plaintext)

	encrypted, err := encManager.AESEncrypt(plaintext)
	if err != nil {
		fmt.Printf("加密失败: %v\n", err)
	} else {
		fmt.Printf("密文: %s\n", encrypted)

		decrypted, err := encManager.AESDecrypt(encrypted)
		if err != nil {
			fmt.Printf("解密失败: %v\n", err)
		} else {
			fmt.Printf("解密: %s\n", decrypted)
		}
	}

	// RSA加密示例
	fmt.Println("\n🔹 RSA非对称加密演示:")
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		fmt.Printf("生成密钥对失败: %v\n", err)
		return
	}

	message := "RSA加密测试消息"
	fmt.Printf("原文: %s\n", message)

	rsaEncrypted, err := keyPair.RSAEncrypt(message)
	if err != nil {
		fmt.Printf("RSA加密失败: %v\n", err)
	} else {
		fmt.Printf("RSA密文: %s...\n", rsaEncrypted[:50])

		rsaDecrypted, err := keyPair.RSADecrypt(rsaEncrypted)
		if err != nil {
			fmt.Printf("RSA解密失败: %v\n", err)
		} else {
			fmt.Printf("RSA解密: %s\n", rsaDecrypted)
		}
	}

	// 数字签名示例
	fmt.Println("\n🔹 数字签名演示:")
	signMessage := "需要签名的重要文档"
	signature, err := keyPair.RSASign(signMessage)
	if err != nil {
		fmt.Printf("签名失败: %v\n", err)
	} else {
		fmt.Printf("消息: %s\n", signMessage)
		fmt.Printf("签名: %s...\n", signature[:50])

		err = keyPair.RSAVerify(signMessage, signature)
		if err != nil {
			fmt.Printf("签名验证失败: %v\n", err)
		} else {
			fmt.Println("✅ 签名验证成功")
		}
	}

	// 哈希示例
	fmt.Println("\n🔹 哈希函数演示:")
	hashManager := NewHashManager()
	data := "需要哈希的数据"

	md5Hash := hashManager.MD5Hash(data)
	sha256Hash := hashManager.SHA256Hash(data)
	hmacHash := hashManager.HMACSHA256(data, "secret-key")

	fmt.Printf("原始数据: %s\n", data)
	fmt.Printf("MD5: %s\n", md5Hash)
	fmt.Printf("SHA256: %s\n", sha256Hash)
	fmt.Printf("HMAC-SHA256: %s\n", hmacHash)

	// 密码管理示例
	fmt.Println("\n🔹 密码管理演示:")
	passwordManager := NewPasswordManager()

	password := "mypassword123"
	hashedPassword, err := passwordManager.HashPassword(password)
	if err != nil {
		fmt.Printf("密码哈希失败: %v\n", err)
	} else {
		fmt.Printf("原始密码: %s\n", password)
		fmt.Printf("哈希密码: %s\n", hashedPassword)

		// 验证密码
		if passwordManager.VerifyPassword(password, hashedPassword) {
			fmt.Println("✅ 密码验证成功")
		} else {
			fmt.Println("❌ 密码验证失败")
		}

		// 错误密码验证
		if passwordManager.VerifyPassword("wrongpassword", hashedPassword) {
			fmt.Println("❌ 错误密码验证成功（不应该发生）")
		} else {
			fmt.Println("✅ 错误密码验证失败（正确行为）")
		}
	}

	// 生成安全密码
	fmt.Println("\n🔹 安全密码生成演示:")
	securePassword, err := passwordManager.GenerateSecurePassword(12)
	if err != nil {
		fmt.Printf("生成安全密码失败: %v\n", err)
	} else {
		fmt.Printf("生成的安全密码: %s\n", securePassword)
	}

	// 随机字符串生成
	randomString, err := hashManager.SecureRandomString(32)
	if err != nil {
		fmt.Printf("生成随机字符串失败: %v\n", err)
	} else {
		fmt.Printf("随机字符串: %s\n", randomString)
	}

	fmt.Println("\n✅ 加密和安全示例演示完成!")
	fmt.Println("💡 提示: 生产环境中建议使用经过验证的加密库")
	fmt.Println("💡 提示: 密码存储应使用bcrypt或argon2")
	fmt.Println("💡 提示: 敏感数据传输应使用HTTPS")
}
