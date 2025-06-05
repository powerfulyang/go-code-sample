package stdlib

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 字符串处理示例
func StringExamples() {
	fmt.Println("=== 字符串处理示例 ===")

	text := "  Hello, Go World!  "
	fmt.Printf("原始字符串: '%s'\n", text)

	// 基本操作
	fmt.Printf("去除空格: '%s'\n", strings.TrimSpace(text))
	fmt.Printf("转大写: '%s'\n", strings.ToUpper(text))
	fmt.Printf("转小写: '%s'\n", strings.ToLower(text))
	fmt.Printf("包含'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("以'Hello'开头: %t\n", strings.HasPrefix(strings.TrimSpace(text), "Hello"))
	fmt.Printf("以'!'结尾: %t\n", strings.HasSuffix(strings.TrimSpace(text), "!"))

	// 分割和连接
	words := strings.Fields(strings.TrimSpace(text))
	fmt.Printf("分割单词: %v\n", words)
	fmt.Printf("连接单词: '%s'\n", strings.Join(words, "-"))

	// 替换
	fmt.Printf("替换'Go'为'Golang': '%s'\n", strings.ReplaceAll(text, "Go", "Golang"))

	// 重复
	fmt.Printf("重复3次'Go': '%s'\n", strings.Repeat("Go", 3))
}

// 正则表达式示例
func RegexExamples() {
	fmt.Println("\n=== 正则表达式示例 ===")

	text := "联系我们: email@example.com 或拨打 13812345678，访问 https://example.com"
	fmt.Printf("原始文本: %s\n", text)

	// 邮箱匹配
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	emails := emailRegex.FindAllString(text, -1)
	fmt.Printf("提取的邮箱: %v\n", emails)

	// 手机号匹配
	phoneRegex := regexp.MustCompile(`1[3-9]\d{9}`)
	phones := phoneRegex.FindAllString(text, -1)
	fmt.Printf("提取的手机号: %v\n", phones)

	// URL匹配
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)
	urls := urlRegex.FindAllString(text, -1)
	fmt.Printf("提取的URL: %v\n", urls)

	// 替换敏感信息
	maskedText := phoneRegex.ReplaceAllStringFunc(text, func(phone string) string {
		return phone[:3] + "****" + phone[7:]
	})
	fmt.Printf("遮蔽手机号: %s\n", maskedText)
}

// 时间处理示例
func TimeExamples() {
	fmt.Println("\n=== 时间处理示例 ===")

	now := time.Now()
	fmt.Printf("当前时间: %s\n", now.Format("2006-01-02 15:04:05"))

	// 时间格式化
	fmt.Printf("日期: %s\n", now.Format("2006-01-02"))
	fmt.Printf("时间: %s\n", now.Format("15:04:05"))
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Unix时间戳: %d\n", now.Unix())

	// 时间解析
	timeStr := "2023-12-25 15:30:45"
	parsed, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err == nil {
		fmt.Printf("解析时间: %s\n", parsed.Format("2006年01月02日 15:04:05"))
	}

	// 时间计算
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("明天: %s\n", tomorrow.Format("2006-01-02"))

	nextWeek := now.Add(7 * 24 * time.Hour)
	fmt.Printf("下周: %s\n", nextWeek.Format("2006-01-02"))

	// 时间比较
	fmt.Printf("明天在今天之后: %t\n", tomorrow.After(now))
	fmt.Printf("时间差: %v\n", tomorrow.Sub(now))

	// 时区处理
	utc := now.UTC()
	fmt.Printf("UTC时间: %s\n", utc.Format("2006-01-02 15:04:05"))

	// 定时器
	fmt.Println("3秒后执行...")
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	fmt.Println("定时器触发!")
}

// JSON处理示例
func JSONExamples() {
	fmt.Println("\n=== JSON处理示例 ===")

	// 定义结构体
	type Person struct {
		Name     string    `json:"name"`
		Age      int       `json:"age"`
		Email    string    `json:"email,omitempty"`
		IsActive bool      `json:"is_active"`
		Tags     []string  `json:"tags"`
		Created  time.Time `json:"created"`
	}

	// 创建对象
	person := Person{
		Name:     "张三",
		Age:      25,
		Email:    "zhangsan@example.com",
		IsActive: true,
		Tags:     []string{"developer", "golang"},
		Created:  time.Now(),
	}

	// 序列化为JSON
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Printf("JSON序列化失败: %v", err)
		return
	}
	fmt.Printf("JSON序列化:\n%s\n", jsonData)

	// 反序列化
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		log.Printf("JSON反序列化失败: %v", err)
		return
	}
	fmt.Printf("反序列化结果: %+v\n", newPerson)

	// 处理动态JSON
	jsonStr := `{"name":"李四","age":30,"skills":["Go","Python","JavaScript"]}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err == nil {
		fmt.Printf("动态JSON: %+v\n", data)
		if skills, ok := data["skills"].([]interface{}); ok {
			fmt.Printf("技能: %v\n", skills)
		}
	}
}

// 文件操作示例
func FileExamples() {
	fmt.Println("\n=== 文件操作示例 ===")

	filename := "test.txt"
	content := "Hello, Go!\n这是一个测试文件。\n"

	// 写入文件
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Printf("写入文件失败: %v", err)
		return
	}
	fmt.Printf("写入文件: %s\n", filename)

	// 读取文件
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("读取文件失败: %v", err)
		return
	}
	fmt.Printf("文件内容:\n%s", data)

	// 文件信息
	info, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("文件大小: %d 字节\n", info.Size())
		fmt.Printf("修改时间: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
		fmt.Printf("是否为目录: %t\n", info.IsDir())
	}

	// 逐行读取
	file, err := os.Open(filename)
	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		lineNum := 1
		fmt.Println("逐行读取:")
		for scanner.Scan() {
			fmt.Printf("第%d行: %s\n", lineNum, scanner.Text())
			lineNum++
		}
	}

	// 路径操作
	fmt.Printf("文件名: %s\n", filepath.Base(filename))
	fmt.Printf("目录: %s\n", filepath.Dir(filename))
	fmt.Printf("扩展名: %s\n", filepath.Ext(filename))

	// 清理
	os.Remove(filename)
	fmt.Printf("删除文件: %s\n", filename)
}

// 加密和编码示例
func CryptoExamples() {
	fmt.Println("\n=== 加密和编码示例 ===")

	data := "Hello, Go Crypto!"
	fmt.Printf("原始数据: %s\n", data)

	// Base64编码
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("Base64编码: %s\n", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err == nil {
		fmt.Printf("Base64解码: %s\n", decoded)
	}

	// MD5哈希
	md5Hash := md5.Sum([]byte(data))
	fmt.Printf("MD5哈希: %x\n", md5Hash)

	// SHA256哈希
	sha256Hash := sha256.Sum256([]byte(data))
	fmt.Printf("SHA256哈希: %x\n", sha256Hash)

	// 生成随机数据
	randomBytes := make([]byte, 16)
	_, err = rand.Read(randomBytes)
	if err == nil {
		fmt.Printf("随机数据: %x\n", randomBytes)
		fmt.Printf("随机数据(Base64): %s\n", base64.StdEncoding.EncodeToString(randomBytes))
	}
}

// HTTP客户端示例
func HTTPExamples() {
	fmt.Println("\n=== HTTP客户端示例 ===")

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// GET请求
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		log.Printf("GET请求失败: %v", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err == nil {
		var result map[string]interface{}
		if json.Unmarshal(body, &result) == nil {
			fmt.Printf("响应数据: %+v\n", result)
		}
	}

	// POST请求
	postData := map[string]string{
		"name":  "张三",
		"email": "zhangsan@example.com",
	}
	jsonData, _ := json.Marshal(postData)

	resp, err = client.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("POST请求失败: %v", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("POST状态码: %d\n", resp.StatusCode)
}

// URL处理示例
func URLExamples() {
	fmt.Println("\n=== URL处理示例 ===")

	rawURL := "https://example.com:8080/path/to/resource?name=张三&age=25&tags=go,web#section1"
	fmt.Printf("原始URL: %s\n", rawURL)

	// 解析URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Printf("URL解析失败: %v", err)
		return
	}

	fmt.Printf("协议: %s\n", parsedURL.Scheme)
	fmt.Printf("主机: %s\n", parsedURL.Host)
	fmt.Printf("端口: %s\n", parsedURL.Port())
	fmt.Printf("路径: %s\n", parsedURL.Path)
	fmt.Printf("查询参数: %s\n", parsedURL.RawQuery)
	fmt.Printf("片段: %s\n", parsedURL.Fragment)

	// 解析查询参数
	params := parsedURL.Query()
	fmt.Printf("name参数: %s\n", params.Get("name"))
	fmt.Printf("age参数: %s\n", params.Get("age"))
	fmt.Printf("tags参数: %v\n", params["tags"])

	// 构建URL
	newURL := &url.URL{
		Scheme: "https",
		Host:   "api.example.com",
		Path:   "/v1/users",
	}

	// 添加查询参数
	query := newURL.Query()
	query.Set("page", "1")
	query.Set("limit", "10")
	query.Add("filter", "active")
	query.Add("filter", "verified")
	newURL.RawQuery = query.Encode()

	fmt.Printf("构建的URL: %s\n", newURL.String())

	// URL编码
	encoded := url.QueryEscape("Hello, 世界!")
	fmt.Printf("URL编码: %s\n", encoded)

	decoded, _ := url.QueryUnescape(encoded)
	fmt.Printf("URL解码: %s\n", decoded)
}

// 排序示例
func SortExamples() {
	fmt.Println("\n=== 排序示例 ===")

	// 基本类型排序
	ints := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("原始整数: %v\n", ints)
	sort.Ints(ints)
	fmt.Printf("排序后: %v\n", ints)

	strings := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("原始字符串: %v\n", strings)
	sort.Strings(strings)
	fmt.Printf("排序后: %v\n", strings)

	// 自定义排序
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"张三", 25},
		{"李四", 30},
		{"王五", 20},
		{"赵六", 35},
	}

	fmt.Printf("原始人员: %+v\n", people)

	// 按年龄排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("按年龄排序: %+v\n", people)

	// 按姓名排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Printf("按姓名排序: %+v\n", people)

	// 检查是否已排序
	ages := []int{20, 25, 30, 35}
	fmt.Printf("年龄数组 %v 是否已排序: %t\n", ages, sort.IntsAreSorted(ages))

	// 二分查找
	target := 25
	index := sort.SearchInts(ages, target)
	fmt.Printf("在已排序数组中查找 %d: 索引 %d\n", target, index)
}

// 类型转换示例
func ConversionExamples() {
	fmt.Println("\n=== 类型转换示例 ===")

	// 字符串转数字
	strNum := "12345"
	if num, err := strconv.Atoi(strNum); err == nil {
		fmt.Printf("字符串 '%s' 转整数: %d\n", strNum, num)
	}

	strFloat := "3.14159"
	if f, err := strconv.ParseFloat(strFloat, 64); err == nil {
		fmt.Printf("字符串 '%s' 转浮点数: %.2f\n", strFloat, f)
	}

	// 数字转字符串
	intVal := 42
	fmt.Printf("整数 %d 转字符串: '%s'\n", intVal, strconv.Itoa(intVal))

	floatVal := 3.14159
	fmt.Printf("浮点数 %.5f 转字符串: '%s'\n", floatVal, strconv.FormatFloat(floatVal, 'f', 2, 64))

	// 布尔值转换
	boolStr := "true"
	if b, err := strconv.ParseBool(boolStr); err == nil {
		fmt.Printf("字符串 '%s' 转布尔值: %t\n", boolStr, b)
	}

	boolVal := true
	fmt.Printf("布尔值 %t 转字符串: '%s'\n", boolVal, strconv.FormatBool(boolVal))

	// 进制转换
	hexStr := "ff"
	if num, err := strconv.ParseInt(hexStr, 16, 64); err == nil {
		fmt.Printf("十六进制 '%s' 转十进制: %d\n", hexStr, num)
	}

	decNum := int64(255)
	fmt.Printf("十进制 %d 转十六进制: '%s'\n", decNum, strconv.FormatInt(decNum, 16))
}

// Context示例
func ContextExamples() {
	fmt.Println("\n=== Context示例 ===")

	// 带超时的Context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 模拟耗时操作
	done := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("操作在超时前完成")
	case <-ctx.Done():
		fmt.Printf("操作超时: %v\n", ctx.Err())
	}

	// 带值的Context
	type key string
	const userKey key = "user"

	ctx = context.WithValue(context.Background(), userKey, "张三")
	if user := ctx.Value(userKey); user != nil {
		fmt.Printf("从Context获取用户: %s\n", user)
	}

	// 可取消的Context
	ctx, cancel = context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		cancel() // 取消操作
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("操作完成")
	case <-ctx.Done():
		fmt.Printf("操作被取消: %v\n", ctx.Err())
	}
}

// 标准库示例主函数
func StandardLibraryExamples() {
	fmt.Println("🚀 Go标准库示例")
	fmt.Println("=" + strings.Repeat("=", 49))

	StringExamples()
	RegexExamples()
	TimeExamples()
	JSONExamples()
	FileExamples()
	CryptoExamples()
	HTTPExamples()
	URLExamples()
	SortExamples()
	ConversionExamples()
	ContextExamples()

	fmt.Println("\n✅ 标准库示例演示完成!")
}
