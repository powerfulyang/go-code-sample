package stdlib

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestStringOperations(t *testing.T) {
	t.Run("BasicStringOperations", func(t *testing.T) {
		text := "  Hello, Go World!  "

		// 测试去除空格
		trimmed := strings.TrimSpace(text)
		expected := "Hello, Go World!"
		if trimmed != expected {
			t.Errorf("TrimSpace failed: got %q, want %q", trimmed, expected)
		}

		// 测试大小写转换
		upper := strings.ToUpper(trimmed)
		if upper != "HELLO, GO WORLD!" {
			t.Errorf("ToUpper failed: got %q", upper)
		}

		lower := strings.ToLower(trimmed)
		if lower != "hello, go world!" {
			t.Errorf("ToLower failed: got %q", lower)
		}

		// 测试包含检查
		if !strings.Contains(trimmed, "Go") {
			t.Error("Should contain 'Go'")
		}

		// 测试前缀和后缀
		if !strings.HasPrefix(trimmed, "Hello") {
			t.Error("Should start with 'Hello'")
		}

		if !strings.HasSuffix(trimmed, "!") {
			t.Error("Should end with '!'")
		}

		t.Log("字符串基本操作测试通过")
	})

	t.Run("StringSplitAndJoin", func(t *testing.T) {
		text := "apple,banana,cherry"
		parts := strings.Split(text, ",")

		expected := []string{"apple", "banana", "cherry"}
		if len(parts) != len(expected) {
			t.Errorf("Split length mismatch: got %d, want %d", len(parts), len(expected))
		}

		for i, part := range parts {
			if part != expected[i] {
				t.Errorf("Split part %d: got %q, want %q", i, part, expected[i])
			}
		}

		joined := strings.Join(parts, "-")
		expectedJoined := "apple-banana-cherry"
		if joined != expectedJoined {
			t.Errorf("Join failed: got %q, want %q", joined, expectedJoined)
		}

		t.Log("字符串分割和连接测试通过")
	})
}

func TestRegexOperations(t *testing.T) {
	t.Run("EmailExtraction", func(t *testing.T) {
		text := "联系我们: test@example.com 或 admin@company.org"
		emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
		emails := emailRegex.FindAllString(text, -1)

		expected := []string{"test@example.com", "admin@company.org"}
		if len(emails) != len(expected) {
			t.Errorf("Email extraction failed: got %d emails, want %d", len(emails), len(expected))
		}

		for i, email := range emails {
			if email != expected[i] {
				t.Errorf("Email %d: got %q, want %q", i, email, expected[i])
			}
		}

		t.Log("邮箱提取测试通过")
	})

	t.Run("PhoneNumberExtraction", func(t *testing.T) {
		text := "请拨打 13812345678 或 15987654321"
		phoneRegex := regexp.MustCompile(`1[3-9]\d{9}`)
		phones := phoneRegex.FindAllString(text, -1)

		expected := []string{"13812345678", "15987654321"}
		if len(phones) != len(expected) {
			t.Errorf("Phone extraction failed: got %d phones, want %d", len(phones), len(expected))
		}

		for i, phone := range phones {
			if phone != expected[i] {
				t.Errorf("Phone %d: got %q, want %q", i, phone, expected[i])
			}
		}

		t.Log("手机号提取测试通过")
	})
}

func TestTimeOperations(t *testing.T) {
	t.Run("TimeFormatting", func(t *testing.T) {
		// 使用固定时间进行测试
		testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)

		// 测试格式化
		dateStr := testTime.Format("2006-01-02")
		if dateStr != "2023-12-25" {
			t.Errorf("Date format failed: got %q, want %q", dateStr, "2023-12-25")
		}

		timeStr := testTime.Format("15:04:05")
		if timeStr != "15:30:45" {
			t.Errorf("Time format failed: got %q, want %q", timeStr, "15:30:45")
		}

		t.Log("时间格式化测试通过")
	})

	t.Run("TimeParsing", func(t *testing.T) {
		timeStr := "2023-12-25 15:30:45"
		parsed, err := time.Parse("2006-01-02 15:04:05", timeStr)
		if err != nil {
			t.Errorf("Time parsing failed: %v", err)
		}

		expected := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
		if !parsed.Equal(expected) {
			t.Errorf("Parsed time mismatch: got %v, want %v", parsed, expected)
		}

		t.Log("时间解析测试通过")
	})

	t.Run("TimeCalculation", func(t *testing.T) {
		base := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)

		// 添加一天
		tomorrow := base.AddDate(0, 0, 1)
		expected := time.Date(2023, 12, 26, 0, 0, 0, 0, time.UTC)
		if !tomorrow.Equal(expected) {
			t.Errorf("AddDate failed: got %v, want %v", tomorrow, expected)
		}

		// 添加小时
		later := base.Add(2 * time.Hour)
		expected = time.Date(2023, 12, 25, 2, 0, 0, 0, time.UTC)
		if !later.Equal(expected) {
			t.Errorf("Add failed: got %v, want %v", later, expected)
		}

		t.Log("时间计算测试通过")
	})
}

func TestJSONOperations(t *testing.T) {
	type Person struct {
		Name  string   `json:"name"`
		Age   int      `json:"age"`
		Email string   `json:"email,omitempty"`
		Tags  []string `json:"tags"`
	}

	t.Run("JSONMarshal", func(t *testing.T) {
		person := Person{
			Name:  "张三",
			Age:   25,
			Email: "zhangsan@example.com",
			Tags:  []string{"developer", "golang"},
		}

		jsonData, err := json.Marshal(person)
		if err != nil {
			t.Errorf("JSON marshal failed: %v", err)
		}

		// 验证JSON包含预期字段
		jsonStr := string(jsonData)
		if !strings.Contains(jsonStr, "张三") {
			t.Error("JSON should contain name")
		}
		if !strings.Contains(jsonStr, "25") {
			t.Error("JSON should contain age")
		}

		t.Log("JSON序列化测试通过")
	})

	t.Run("JSONUnmarshal", func(t *testing.T) {
		jsonStr := `{"name":"李四","age":30,"tags":["python","java"]}`

		var person Person
		err := json.Unmarshal([]byte(jsonStr), &person)
		if err != nil {
			t.Errorf("JSON unmarshal failed: %v", err)
		}

		if person.Name != "李四" {
			t.Errorf("Name mismatch: got %q, want %q", person.Name, "李四")
		}
		if person.Age != 30 {
			t.Errorf("Age mismatch: got %d, want %d", person.Age, 30)
		}
		if len(person.Tags) != 2 {
			t.Errorf("Tags length mismatch: got %d, want %d", len(person.Tags), 2)
		}

		t.Log("JSON反序列化测试通过")
	})
}

func TestFileOperations(t *testing.T) {
	t.Run("FileWriteAndRead", func(t *testing.T) {
		filename := "test_file.txt"
		content := "Hello, Go Testing!"

		// 写入文件
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Errorf("WriteFile failed: %v", err)
		}

		// 读取文件
		data, err := os.ReadFile(filename)
		if err != nil {
			t.Errorf("ReadFile failed: %v", err)
		}

		if string(data) != content {
			t.Errorf("File content mismatch: got %q, want %q", string(data), content)
		}

		// 清理
		os.Remove(filename)

		t.Log("文件读写测试通过")
	})

	t.Run("FileInfo", func(t *testing.T) {
		filename := "test_info.txt"
		content := "Test file info"

		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Errorf("WriteFile failed: %v", err)
		}

		info, err := os.Stat(filename)
		if err != nil {
			t.Errorf("Stat failed: %v", err)
		}

		if info.Size() != int64(len(content)) {
			t.Errorf("File size mismatch: got %d, want %d", info.Size(), len(content))
		}

		if info.IsDir() {
			t.Error("File should not be a directory")
		}

		// 清理
		os.Remove(filename)

		t.Log("文件信息测试通过")
	})
}

func TestCryptoOperations(t *testing.T) {
	t.Run("Base64Encoding", func(t *testing.T) {
		data := "Hello, Go Crypto!"

		encoded := base64.StdEncoding.EncodeToString([]byte(data))
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			t.Errorf("Base64 decode failed: %v", err)
		}

		if string(decoded) != data {
			t.Errorf("Base64 round-trip failed: got %q, want %q", string(decoded), data)
		}

		t.Log("Base64编码测试通过")
	})

	t.Run("HashFunctions", func(t *testing.T) {
		data := "Hello, Go Crypto!"

		// MD5哈希
		md5Hash := md5.Sum([]byte(data))
		if len(md5Hash) != 16 {
			t.Errorf("MD5 hash length should be 16, got %d", len(md5Hash))
		}

		// SHA256哈希
		sha256Hash := sha256.Sum256([]byte(data))
		if len(sha256Hash) != 32 {
			t.Errorf("SHA256 hash length should be 32, got %d", len(sha256Hash))
		}

		// 相同输入应该产生相同哈希
		md5Hash2 := md5.Sum([]byte(data))
		if md5Hash != md5Hash2 {
			t.Error("MD5 hash should be deterministic")
		}

		t.Log("哈希函数测试通过")
	})
}

func TestURLOperations(t *testing.T) {
	t.Run("URLParsing", func(t *testing.T) {
		rawURL := "https://example.com:8080/path/to/resource?name=test&age=25#section1"

		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			t.Errorf("URL parse failed: %v", err)
		}

		if parsedURL.Scheme != "https" {
			t.Errorf("Scheme mismatch: got %q, want %q", parsedURL.Scheme, "https")
		}

		if parsedURL.Host != "example.com:8080" {
			t.Errorf("Host mismatch: got %q, want %q", parsedURL.Host, "example.com:8080")
		}

		if parsedURL.Path != "/path/to/resource" {
			t.Errorf("Path mismatch: got %q, want %q", parsedURL.Path, "/path/to/resource")
		}

		if parsedURL.Fragment != "section1" {
			t.Errorf("Fragment mismatch: got %q, want %q", parsedURL.Fragment, "section1")
		}

		t.Log("URL解析测试通过")
	})

	t.Run("QueryParameters", func(t *testing.T) {
		rawURL := "https://example.com/search?q=golang&page=1&limit=10"

		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			t.Errorf("URL parse failed: %v", err)
		}

		params := parsedURL.Query()

		if params.Get("q") != "golang" {
			t.Errorf("Query param 'q' mismatch: got %q, want %q", params.Get("q"), "golang")
		}

		if params.Get("page") != "1" {
			t.Errorf("Query param 'page' mismatch: got %q, want %q", params.Get("page"), "1")
		}

		t.Log("查询参数测试通过")
	})
}

func TestSortOperations(t *testing.T) {
	t.Run("BasicSorting", func(t *testing.T) {
		ints := []int{64, 34, 25, 12, 22, 11, 90}
		originalLen := len(ints)

		sort.Ints(ints)

		// 检查长度没有改变
		if len(ints) != originalLen {
			t.Errorf("Sort changed slice length: got %d, want %d", len(ints), originalLen)
		}

		// 检查是否已排序
		if !sort.IntsAreSorted(ints) {
			t.Errorf("Slice should be sorted: %v", ints)
		}

		// 检查第一个和最后一个元素
		if ints[0] != 11 {
			t.Errorf("First element should be 11, got %d", ints[0])
		}
		if ints[len(ints)-1] != 90 {
			t.Errorf("Last element should be 90, got %d", ints[len(ints)-1])
		}

		t.Log("基本排序测试通过")
	})

	t.Run("CustomSorting", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{"张三", 25},
			{"李四", 30},
			{"王五", 20},
		}

		// 按年龄排序
		sort.Slice(people, func(i, j int) bool {
			return people[i].Age < people[j].Age
		})

		// 检查排序结果
		if people[0].Age != 20 {
			t.Errorf("First person age should be 20, got %d", people[0].Age)
		}
		if people[2].Age != 30 {
			t.Errorf("Last person age should be 30, got %d", people[2].Age)
		}

		t.Log("自定义排序测试通过")
	})
}

func TestConversionOperations(t *testing.T) {
	t.Run("StringToNumber", func(t *testing.T) {
		// 字符串转整数
		strNum := "12345"
		num, err := strconv.Atoi(strNum)
		if err != nil {
			t.Errorf("Atoi failed: %v", err)
		}
		if num != 12345 {
			t.Errorf("Atoi result mismatch: got %d, want %d", num, 12345)
		}

		// 字符串转浮点数
		strFloat := "3.14159"
		f, err := strconv.ParseFloat(strFloat, 64)
		if err != nil {
			t.Errorf("ParseFloat failed: %v", err)
		}
		if f < 3.14 || f > 3.15 {
			t.Errorf("ParseFloat result out of range: got %f", f)
		}

		t.Log("字符串转数字测试通过")
	})

	t.Run("NumberToString", func(t *testing.T) {
		// 整数转字符串
		intVal := 42
		strVal := strconv.Itoa(intVal)
		if strVal != "42" {
			t.Errorf("Itoa result mismatch: got %q, want %q", strVal, "42")
		}

		// 浮点数转字符串
		floatVal := 3.14159
		strFloat := strconv.FormatFloat(floatVal, 'f', 2, 64)
		if strFloat != "3.14" {
			t.Errorf("FormatFloat result mismatch: got %q, want %q", strFloat, "3.14")
		}

		t.Log("数字转字符串测试通过")
	})
}

func TestContextOperations(t *testing.T) {
	t.Run("ContextWithTimeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		select {
		case <-time.After(200 * time.Millisecond):
			t.Error("Context should have timed out")
		case <-ctx.Done():
			if ctx.Err() != context.DeadlineExceeded {
				t.Errorf("Context error should be DeadlineExceeded, got %v", ctx.Err())
			}
		}

		t.Log("Context超时测试通过")
	})

	t.Run("ContextWithValue", func(t *testing.T) {
		type key string
		const userKey key = "user"

		ctx := context.WithValue(context.Background(), userKey, "testuser")

		value := ctx.Value(userKey)
		if value != "testuser" {
			t.Errorf("Context value mismatch: got %v, want %v", value, "testuser")
		}

		// 测试不存在的键
		nonExistent := ctx.Value("nonexistent")
		if nonExistent != nil {
			t.Errorf("Non-existent key should return nil, got %v", nonExistent)
		}

		t.Log("Context值测试通过")
	})
}

// 基准测试
func BenchmarkStringOperations(b *testing.B) {
	text := "Hello, Go World!"

	b.Run("ToUpper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strings.ToUpper(text)
		}
	})

	b.Run("Contains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strings.Contains(text, "Go")
		}
	})
}

func BenchmarkJSONOperations(b *testing.B) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	person := Person{Name: "张三", Age: 25}
	jsonData, _ := json.Marshal(person)

	b.Run("Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = json.Marshal(person)
		}
	})

	b.Run("Unmarshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var p Person
			_ = json.Unmarshal(jsonData, &p)
		}
	})
}

func BenchmarkCryptoOperations(b *testing.B) {
	data := []byte("Hello, Go Crypto!")

	b.Run("MD5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = md5.Sum(data)
		}
	})

	b.Run("SHA256", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = sha256.Sum256(data)
		}
	})
}
