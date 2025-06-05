package format

import (
	"fmt"
	"time"
)

// Person 自定义类型，实现 String() 方法
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s(%d岁)", p.Name, p.Age)
}

func demonstrateAdvancedFormats() {
	fmt.Println("==== 更多格式化占位符 ====")

	// 数字格式化的变体
	num := 12345
	fmt.Printf("%%d 十进制: %d\n", num)
	fmt.Printf("%%o 八进制: %o\n", num)
	fmt.Printf("%%x 十六进制(小写): %x\n", num)
	fmt.Printf("%%X 十六进制(大写): %X\n", num)
	fmt.Printf("%%b 二进制: %b\n", num)

	// 浮点数格式化
	pi := 3.14159265359
	fmt.Printf("%%f 浮点数: %f\n", pi)
	fmt.Printf("%%.2f 保留2位小数: %.2f\n", pi)
	fmt.Printf("%%e 科学记数法: %e\n", pi)
	fmt.Printf("%%g 自动选择格式: %g\n", pi)

	// 字符和字符串
	char := 'A'
	fmt.Printf("%%c 字符: %c\n", char)
	fmt.Printf("%%q 带引号的字符: %q\n", char)

	str := "Hello 世界"
	fmt.Printf("%%s 字符串: %s\n", str)
	fmt.Printf("%%q 带引号的字符串: %q\n", str)
	fmt.Printf("%%x 字符串的十六进制: %x\n", str)

	// 指针
	ptr := &num
	fmt.Printf("%%p 指针地址: %p\n", ptr)

	// 类型信息
	fmt.Printf("%%T 类型: %T\n", num)
	fmt.Printf("%%T 类型: %T\n", str)
	fmt.Printf("%%T 类型: %T\n", true)

	// 通用格式
	fmt.Printf("%%v 默认格式: %v\n", num)
	fmt.Printf("%%+v 包含字段名: %+v\n", Person{Name: "张三", Age: 25})
	fmt.Printf("%%#v Go语法表示: %#v\n", Person{Name: "张三", Age: 25})

	// 宽度和对齐
	fmt.Println("\n==== 宽度和对齐 ====")
	fmt.Printf("|%%5d| 右对齐，宽度5: |%5d|\n", 42)
	fmt.Printf("|%%-5d| 左对齐，宽度5: |%-5d|\n", 42)
	fmt.Printf("|%%05d| 零填充，宽度5: |%05d|\n", 42)

	fmt.Printf("|%%10s| 右对齐，宽度10: |%10s|\n", "hello")
	fmt.Printf("|%%-10s| 左对齐，宽度10: |%-10s|\n", "hello")

	// 精度控制
	fmt.Printf("%%.2s 字符串精度: %.2s\n", "hello")
	fmt.Printf("%%8.2f 宽度8，精度2: %8.2f\n", pi)
}

func demonstrateConditionalFormatting() {
	fmt.Println("\n==== 条件格式化 ====")

	scores := []int{95, 87, 72, 65, 45}
	names := []string{"张三", "李四", "王五", "赵六", "钱七"}

	fmt.Printf("%-8s %-6s %-8s\n", "姓名", "分数", "等级")
	fmt.Println("------------------------")

	for i, score := range scores {
		var grade string
		var pass bool

		switch {
		case score >= 90:
			grade = "优秀"
			pass = true
		case score >= 80:
			grade = "良好"
			pass = true
		case score >= 70:
			grade = "中等"
			pass = true
		case score >= 60:
			grade = "及格"
			pass = true
		default:
			grade = "不及格"
			pass = false
		}

		fmt.Printf("%-8s %-6d %-8s (通过: %t)\n",
			names[i], score, grade, pass)
	}
}

func demonstrateTimeFormatting() {
	fmt.Println("\n==== 时间格式化 ====")

	now := time.Now()

	fmt.Printf("当前时间 %%v: %v\n", now)
	fmt.Printf("当前时间 %%s: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("日期 %%s: %s\n", now.Format("2006年01月02日"))
	fmt.Printf("时间 %%s: %s\n", now.Format("15:04:05"))
	fmt.Printf("星期 %%s: %s\n", now.Format("Monday"))
}
