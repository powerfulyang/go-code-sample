package format

import (
	"fmt"
	"os"
)

// 日志输出函数
func logInfo(level string, message string, success bool) {
	fmt.Printf("[%s] %s - 成功: %t\n", level, message, success)
}

// 用户信息展示
func displayUserProfile(name string, age int, isVIP bool, balance float64) {
	fmt.Printf("用户档案:\n")
	fmt.Printf("  姓名: %s\n", name)
	fmt.Printf("  年龄: %d 岁\n", age)
	fmt.Printf("  VIP状态: %t\n", isVIP)
	fmt.Printf("  余额: %.2f 元\n", balance)

	if isVIP {
		fmt.Printf("  享受VIP特权: %t\n", true)
	}
}

// 错误报告
func reportError(errorCode int, errorMsg string, isCritical bool) {
	if isCritical {
		fmt.Fprintf(os.Stderr, "严重错误 [%d]: %s - 关键: %t\n",
			errorCode, errorMsg, isCritical)
	} else {
		fmt.Printf("警告 [%d]: %s - 关键: %t\n",
			errorCode, errorMsg, isCritical)
	}
}

// 进度显示
func showProgress(current int, total int, taskName string) {
	percentage := float64(current) / float64(total) * 100
	isComplete := current >= total

	fmt.Printf("任务 '%s': %d/%d (%.1f%%) - 完成: %t\n",
		taskName, current, total, percentage, isComplete)
}
