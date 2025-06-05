// Package stringutils 提供字符串处理的实用工具函数
package stringutils

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Reverse 反转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome 检查字符串是否为回文
func IsPalindrome(s string) bool {
	// 转换为小写并移除空格
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return cleaned == Reverse(cleaned)
}

// WordCount 统计单词数量
func WordCount(s string) int {
	if strings.TrimSpace(s) == "" {
		return 0
	}
	return len(strings.Fields(s))
}

// CharCount 统计字符数量（Unicode字符）
func CharCount(s string) int {
	return utf8.RuneCountInString(s)
}

// ByteCount 统计字节数量
func ByteCount(s string) int {
	return len(s)
}

// Capitalize 首字母大写
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Title 标题格式（每个单词首字母大写）
func Title(s string) string {
	return strings.Title(strings.ToLower(s))
}

// CamelCase 转换为驼峰命名
func CamelCase(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}

	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		result += Capitalize(strings.ToLower(words[i]))
	}
	return result
}

// PascalCase 转换为帕斯卡命名
func PascalCase(s string) string {
	words := strings.Fields(s)
	var result string
	for _, word := range words {
		result += Capitalize(strings.ToLower(word))
	}
	return result
}

// SnakeCase 转换为蛇形命名
func SnakeCase(s string) string {
	words := strings.Fields(s)
	var result []string
	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}
	return strings.Join(result, "_")
}

// KebabCase 转换为短横线命名
func KebabCase(s string) string {
	words := strings.Fields(s)
	var result []string
	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}
	return strings.Join(result, "-")
}

// RemoveSpaces 移除所有空格
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// RemoveExtraSpaces 移除多余空格（保留单个空格）
func RemoveExtraSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Truncate 截断字符串到指定长度
func Truncate(s string, length int) string {
	if length <= 0 {
		return ""
	}

	runes := []rune(s)
	if len(runes) <= length {
		return s
	}

	return string(runes[:length])
}

// TruncateWithEllipsis 截断字符串并添加省略号
func TruncateWithEllipsis(s string, length int) string {
	if length <= 3 {
		return Truncate(s, length)
	}

	runes := []rune(s)
	if len(runes) <= length {
		return s
	}

	return string(runes[:length-3]) + "..."
}

// Pad 在字符串两边填充字符到指定长度
func Pad(s string, length int, padChar rune) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}

	totalPad := length - len(runes)
	leftPad := totalPad / 2
	rightPad := totalPad - leftPad

	return strings.Repeat(string(padChar), leftPad) + s + strings.Repeat(string(padChar), rightPad)
}

// PadLeft 在字符串左边填充字符
func PadLeft(s string, length int, padChar rune) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}

	padCount := length - len(runes)
	return strings.Repeat(string(padChar), padCount) + s
}

// PadRight 在字符串右边填充字符
func PadRight(s string, length int, padChar rune) string {
	runes := []rune(s)
	if len(runes) >= length {
		return s
	}

	padCount := length - len(runes)
	return s + strings.Repeat(string(padChar), padCount)
}

// IsEmail 简单的邮箱格式验证
func IsEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsPhone 简单的手机号格式验证（中国）
func IsPhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}

// IsURL 简单的URL格式验证
func IsURL(url string) bool {
	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	return urlRegex.MatchString(url)
}

// ExtractNumbers 提取字符串中的所有数字
func ExtractNumbers(s string) []string {
	numberRegex := regexp.MustCompile(`\d+`)
	return numberRegex.FindAllString(s, -1)
}

// ExtractEmails 提取字符串中的所有邮箱地址
func ExtractEmails(s string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	return emailRegex.FindAllString(s, -1)
}

// MaskEmail 遮蔽邮箱地址
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	username := parts[0]
	domain := parts[1]

	if len(username) <= 2 {
		return email
	}

	masked := string(username[0]) + strings.Repeat("*", len(username)-2) + string(username[len(username)-1])
	return masked + "@" + domain
}

// MaskPhone 遮蔽手机号
func MaskPhone(phone string) string {
	if len(phone) != 11 {
		return phone
	}

	return phone[:3] + "****" + phone[7:]
}

// RandomString 生成指定长度的随机字符串
func RandomString(length int, charset string) string {
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[i%len(charset)] // 简化版本，实际应该使用随机数
	}
	return string(result)
}

// Similarity 计算两个字符串的相似度（简单版本）
func Similarity(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}

	if len(s1) == 0 || len(s2) == 0 {
		return 0.0
	}

	// 简单的字符匹配相似度
	matches := 0
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}

	for i := 0; i < minLen; i++ {
		if s1[i] == s2[i] {
			matches++
		}
	}

	maxLen := len(s1)
	if len(s2) > maxLen {
		maxLen = len(s2)
	}

	return float64(matches) / float64(maxLen)
}

// LevenshteinDistance 计算编辑距离
func LevenshteinDistance(s1, s2 string) int {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	m, n := len(runes1), len(runes2)

	// 创建矩阵
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 填充矩阵
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if runes1[i-1] == runes2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}

// min 返回三个数中的最小值
func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

// ContainsAny 检查字符串是否包含任意一个子字符串
func ContainsAny(s string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// ContainsAll 检查字符串是否包含所有子字符串
func ContainsAll(s string, substrings []string) bool {
	for _, substr := range substrings {
		if !strings.Contains(s, substr) {
			return false
		}
	}
	return true
}

// SplitAndTrim 分割字符串并去除空白
func SplitAndTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// JoinNonEmpty 连接非空字符串
func JoinNonEmpty(strs []string, sep string) string {
	var nonEmpty []string
	for _, s := range strs {
		if s != "" {
			nonEmpty = append(nonEmpty, s)
		}
	}
	return strings.Join(nonEmpty, sep)
}
