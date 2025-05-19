package datatypes

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// StringsExample 展示Go中的字符串类型和操作
func StringsExample() {
	fmt.Println("=== 字符串基础 ===")

	// 字符串声明和初始化
	var s1 string = "你好，Go语言！"
	s2 := "Hello, Go!"

	// 字符串长度
	fmt.Printf("s1 长度（字节数）: %d\n", len(s1))
	fmt.Printf("s1 长度（字符数）: %d\n", utf8.RuneCountInString(s1))
	fmt.Printf("s2 长度（字节数）: %d\n", len(s2))

	// 多行字符串
	multiLine := `这是一个
多行字符串
示例。
可以包含"引号"和其他特殊字符：!@#$%^&*()`

	fmt.Println("\n多行字符串示例:")
	fmt.Println(multiLine)

	// 字符串是不可变的
	//s1[0] = 'X' // 这行会导致编译错误

	// 字符串拼接
	fmt.Println("\n=== 字符串操作 ===")

	name := "张三"
	greeting := "您好，" + name + "！"
	fmt.Println("字符串拼接:", greeting)

	// 使用 fmt.Sprintf 格式化字符串
	formatted := fmt.Sprintf("用户: %s, 年龄: %d, 余额: %.2f", name, 30, 99.99)
	fmt.Println("格式化字符串:", formatted)

	// 字符串重复
	repeated := strings.Repeat("Go ", 3)
	fmt.Println("重复字符串:", repeated)

	// 字符串比较
	a := "apple"
	b := "banana"
	fmt.Printf("%s == %s: %t\n", a, b, a == b)
	fmt.Printf("%s != %s: %t\n", a, b, a != b)
	fmt.Printf("%s < %s: %t\n", a, b, a < b) // 字典序比较

	// 字符串查找和替换
	sentence := "Go语言是一种开源编程语言，Go语言的设计简洁高效"

	// 查找子串
	substring := "Go语言"
	pos := strings.Index(sentence, substring)
	fmt.Printf("在句子中查找 '%s': 位置为 %d\n", substring, pos)

	lastPos := strings.LastIndex(sentence, substring)
	fmt.Printf("在句子中查找 '%s' 的最后位置: %d\n", substring, lastPos)

	// 检查字符串是否包含子串
	contains := strings.Contains(sentence, "开源")
	fmt.Printf("句子包含 '开源': %t\n", contains)

	// 检查前缀和后缀
	hasPrefix := strings.HasPrefix(sentence, "Go")
	hasSuffix := strings.HasSuffix(sentence, "高效")
	fmt.Printf("句子以 'Go' 开始: %t\n", hasPrefix)
	fmt.Printf("句子以 '高效' 结束: %t\n", hasSuffix)

	// 替换
	replaced := strings.Replace(sentence, "Go语言", "Golang", 1) // 只替换第一次出现
	fmt.Println("替换一次:", replaced)

	replacedAll := strings.ReplaceAll(sentence, "Go语言", "Golang") // 替换所有出现
	fmt.Println("替换所有:", replacedAll)

	// 分割和连接
	text := "apple,banana,orange,grape"

	// 字符串分割
	parts := strings.Split(text, ",")
	fmt.Println("分割字符串:", parts)

	// 字符串连接
	joined := strings.Join(parts, " | ")
	fmt.Println("连接字符串:", joined)

	// 去除空白字符
	padded := "  \t Hello, World! \n  "

	trimmed := strings.TrimSpace(padded)
	fmt.Printf("原字符串: %q\n", padded)
	fmt.Printf("去除空白后: %q\n", trimmed)

	// 大小写转换
	original := "Hello, Go World!"

	lower := strings.ToLower(original)
	upper := strings.ToUpper(original)
	fmt.Printf("原字符串: %s\n", original)
	fmt.Printf("小写: %s\n", lower)
	fmt.Printf("大写: %s\n", upper)

	// 字符(Unicode码点)操作
	fmt.Println("\n=== 字符(Unicode码点)操作 ===")

	chinese := "你好，世界"

	// 遍历字符串中的字符(Unicode码点)
	fmt.Println("遍历字符串中的字符:")
	for i, char := range chinese {
		fmt.Printf("位置 %d: %c (Unicode: %U, UTF-8: %X)\n", i, char, char, char)
	}

	// 转换为字符切片
	runes := []rune(chinese)
	fmt.Printf("字符切片: %v, 长度: %d\n", runes, len(runes))

	// 将字符切片转回字符串
	modified := string(runes)
	fmt.Printf("转回字符串: %s\n", modified)

	// 转换为字节切片
	bytes := []byte(chinese)
	fmt.Printf("字节切片: %v, 长度: %d\n", bytes, len(bytes))

	// 将字节切片转回字符串
	fromBytes := string(bytes)
	fmt.Printf("从字节切片: %s\n", fromBytes)
}
