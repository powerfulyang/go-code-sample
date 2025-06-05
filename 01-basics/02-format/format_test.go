package format

import (
	"strconv"
	"testing"
	"time"
)

func TestAdvancedFormats(t *testing.T) {
	t.Run("NumberFormats", func(t *testing.T) {
		num := 1234
		t.Logf("%%d 十进制: %d", num)
		t.Logf("%%o 八进制: %o", num)
		t.Logf("%%x 十六进制(小写): %x", num)
		t.Logf("%%X 十六进制(大写): %X", num)
		t.Logf("%%b 二进制: %b", num)

		// 十进制转换为二进制
		const decimal = 42
		binary := strconv.FormatInt(decimal, 2)
		t.Logf("Decimal %d in binary is %s", decimal, binary)
	})

	t.Run("FloatFormats", func(t *testing.T) {
		num := 12345.6789
		t.Logf("%%f 浮点数: %f", num)
		t.Logf("%%.2f 保留2位小数: %.2f", num)
		t.Logf("%%e 科学记数法: %e", num)
		t.Logf("%%g 自动选择格式: %g", num)

		pi := 3.14159265359
		t.Logf("%%f 浮点数: %f", pi)
		t.Logf("%%.2f 保留2位小数: %.2f", pi)
		t.Logf("%%e 科学记数法: %e", pi)
		t.Logf("%%g 自动选择格式: %g", pi)
	})

	t.Run("StringAndCharFormats", func(t *testing.T) {
		char := 'A'
		str := "Hello 世界"
		t.Logf("%%c 字符: %c", char)
		t.Logf("%%q 带引号的字符: %q", char)
		t.Logf("%%s 字符串: %s", str)
		t.Logf("%%q 带引号的字符串: %q", str)
		t.Logf("%%x 字符串的十六进制: %x", str)
	})

	t.Run("WidthAndAlignment", func(t *testing.T) {
		t.Logf("|%%5d| 右对齐，宽度5: |%5d|", 42)
		t.Logf("|%%-5d| 左对齐，宽度5: |%-5d|", 42)
		t.Logf("|%%05d| 零填充，宽度5: |%05d|", 42)
		t.Logf("|%%10s| 右对齐，宽度10: |%10s|", "hello")
		t.Logf("|%%-10s| 左对齐，宽度10: |%-10s|", "hello")
	})
}

func TestPracticalExamples(t *testing.T) {
	t.Run("LogInfo", func(t *testing.T) {
		logInfo("INFO", "测试消息", true)
		logInfo("ERROR", "错误消息", false)
	})

	t.Run("DisplayUserProfile", func(t *testing.T) {
		displayUserProfile("张三", 25, true, 1234.56)
		displayUserProfile("李四", 30, false, 567.89)
	})

	t.Run("ReportError", func(t *testing.T) {
		reportError(404, "页面未找到", false)
		reportError(500, "服务器内部错误", true)
	})

	t.Run("ShowProgress", func(t *testing.T) {
		showProgress(50, 100, "下载文件")
		showProgress(100, 100, "处理完成")
	})
}

func TestTimeFormatting(t *testing.T) {
	now := time.Now()
	t.Run("TimeFormats", func(t *testing.T) {
		t.Logf("当前时间 %%v: %v", now)
		t.Logf("当前时间 %%s: %s", now.Format("2006-01-02 15:04:05"))
		t.Logf("日期 %%s: %s", now.Format("2006年01月02日"))
		t.Logf("时间 %%s: %s", now.Format("15:04:05"))
		t.Logf("星期 %%s: %s", now.Format("Monday"))
	})
}

func TestConditionalFormatting(t *testing.T) {
	t.Run("ScoreGrading", func(t *testing.T) {
		scores := []int{95, 87, 72, 65, 45}
		names := []string{"张三", "李四", "王五", "赵六", "钱七"}

		t.Logf("%-8s %-6s %-8s", "姓名", "分数", "等级")
		t.Log("------------------------")

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

			t.Logf("%-8s %-6d %-8s (通过: %t)",
				names[i], score, grade, pass)
		}
	})
}
