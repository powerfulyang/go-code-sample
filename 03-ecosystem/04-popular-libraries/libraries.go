package libraries

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 注意：这个模块展示了流行Go库的使用模式
// 在实际项目中，你需要使用 go get 安装这些库

// LogrusLikeLogger 模拟Logrus风格的日志记录器
type LogrusLikeLogger struct {
	level  string
	fields map[string]interface{}
}

// NewLogger 创建新的日志记录器
func NewLogger() *LogrusLikeLogger {
	return &LogrusLikeLogger{
		level:  "info",
		fields: make(map[string]interface{}),
	}
}

// WithField 添加字段
func (l *LogrusLikeLogger) WithField(key string, value interface{}) *LogrusLikeLogger {
	newLogger := &LogrusLikeLogger{
		level:  l.level,
		fields: make(map[string]interface{}),
	}

	// 复制现有字段
	for k, v := range l.fields {
		newLogger.fields[k] = v
	}

	// 添加新字段
	newLogger.fields[key] = value
	return newLogger
}

// WithFields 添加多个字段
func (l *LogrusLikeLogger) WithFields(fields map[string]interface{}) *LogrusLikeLogger {
	newLogger := &LogrusLikeLogger{
		level:  l.level,
		fields: make(map[string]interface{}),
	}

	// 复制现有字段
	for k, v := range l.fields {
		newLogger.fields[k] = v
	}

	// 添加新字段
	for k, v := range fields {
		newLogger.fields[k] = v
	}

	return newLogger
}

// Info 记录信息日志
func (l *LogrusLikeLogger) Info(msg string) {
	l.log("INFO", msg)
}

// Error 记录错误日志
func (l *LogrusLikeLogger) Error(msg string) {
	l.log("ERROR", msg)
}

// Warn 记录警告日志
func (l *LogrusLikeLogger) Warn(msg string) {
	l.log("WARN", msg)
}

// Debug 记录调试日志
func (l *LogrusLikeLogger) Debug(msg string) {
	l.log("DEBUG", msg)
}

// log 内部日志记录方法
func (l *LogrusLikeLogger) log(level, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	var fieldsStr string
	if len(l.fields) > 0 {
		fieldsJSON, _ := json.Marshal(l.fields)
		fieldsStr = fmt.Sprintf(" fields=%s", fieldsJSON)
	}

	fmt.Printf("[%s] %s: %s%s\n", timestamp, level, msg, fieldsStr)
}

// ViperLikeConfig 模拟Viper风格的配置管理器
type ViperLikeConfig struct {
	config map[string]interface{}
}

// NewConfig 创建新的配置管理器
func NewConfig() *ViperLikeConfig {
	return &ViperLikeConfig{
		config: make(map[string]interface{}),
	}
}

// SetDefault 设置默认值
func (c *ViperLikeConfig) SetDefault(key string, value interface{}) {
	if _, exists := c.config[key]; !exists {
		c.config[key] = value
	}
}

// Set 设置配置值
func (c *ViperLikeConfig) Set(key string, value interface{}) {
	c.config[key] = value
}

// Get 获取配置值
func (c *ViperLikeConfig) Get(key string) interface{} {
	return c.config[key]
}

// GetString 获取字符串配置值
func (c *ViperLikeConfig) GetString(key string) string {
	if value, ok := c.config[key]; ok {
		if str, ok := value.(string); ok {
			return str
		}
		return fmt.Sprintf("%v", value)
	}
	return ""
}

// GetInt 获取整数配置值
func (c *ViperLikeConfig) GetInt(key string) int {
	if value, ok := c.config[key]; ok {
		switch v := value.(type) {
		case int:
			return v
		case string:
			if i, err := strconv.Atoi(v); err == nil {
				return i
			}
		}
	}
	return 0
}

// GetBool 获取布尔配置值
func (c *ViperLikeConfig) GetBool(key string) bool {
	if value, ok := c.config[key]; ok {
		switch v := value.(type) {
		case bool:
			return v
		case string:
			return v == "true" || v == "1" || v == "yes"
		}
	}
	return false
}

// ReadInConfig 从文件读取配置
func (c *ViperLikeConfig) ReadInConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	for k, v := range config {
		c.config[k] = v
	}

	return nil
}

// WriteConfig 写入配置到文件
func (c *ViperLikeConfig) WriteConfig(filename string) error {
	data, err := json.MarshalIndent(c.config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	return os.WriteFile(filename, data, 0644)
}

// CobraLikeCommand 模拟Cobra风格的命令行工具
type CobraLikeCommand struct {
	Name        string
	Description string
	Flags       map[string]*Flag
	SubCommands map[string]*CobraLikeCommand
	RunFunc     func(cmd *CobraLikeCommand, args []string) error
}

// Flag 命令行标志
type Flag struct {
	Name        string
	Shorthand   string
	Description string
	Value       interface{}
	Required    bool
}

// NewCommand 创建新命令
func NewCommand(name, description string) *CobraLikeCommand {
	return &CobraLikeCommand{
		Name:        name,
		Description: description,
		Flags:       make(map[string]*Flag),
		SubCommands: make(map[string]*CobraLikeCommand),
	}
}

// AddCommand 添加子命令
func (c *CobraLikeCommand) AddCommand(subCmd *CobraLikeCommand) {
	c.SubCommands[subCmd.Name] = subCmd
}

// StringFlag 添加字符串标志
func (c *CobraLikeCommand) StringFlag(name, shorthand, defaultValue, description string) {
	c.Flags[name] = &Flag{
		Name:        name,
		Shorthand:   shorthand,
		Description: description,
		Value:       defaultValue,
		Required:    false,
	}
}

// IntFlag 添加整数标志
func (c *CobraLikeCommand) IntFlag(name, shorthand string, defaultValue int, description string) {
	c.Flags[name] = &Flag{
		Name:        name,
		Shorthand:   shorthand,
		Description: description,
		Value:       defaultValue,
		Required:    false,
	}
}

// BoolFlag 添加布尔标志
func (c *CobraLikeCommand) BoolFlag(name, shorthand string, defaultValue bool, description string) {
	c.Flags[name] = &Flag{
		Name:        name,
		Shorthand:   shorthand,
		Description: description,
		Value:       defaultValue,
		Required:    false,
	}
}

// Execute 执行命令
func (c *CobraLikeCommand) Execute(args []string) error {
	if len(args) == 0 {
		return c.showHelp()
	}

	// 检查是否是子命令
	if subCmd, exists := c.SubCommands[args[0]]; exists {
		return subCmd.Execute(args[1:])
	}

	// 解析标志
	parsedArgs, err := c.parseFlags(args)
	if err != nil {
		return err
	}

	// 执行命令
	if c.RunFunc != nil {
		return c.RunFunc(c, parsedArgs)
	}

	return c.showHelp()
}

// parseFlags 解析命令行标志
func (c *CobraLikeCommand) parseFlags(args []string) ([]string, error) {
	var parsedArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--") {
			// 长标志
			flagName := strings.TrimPrefix(arg, "--")
			if flag, exists := c.Flags[flagName]; exists {
				if i+1 < len(args) {
					flag.Value = args[i+1]
					i++ // 跳过值
				}
			}
		} else if strings.HasPrefix(arg, "-") {
			// 短标志
			shorthand := strings.TrimPrefix(arg, "-")
			for _, flag := range c.Flags {
				if flag.Shorthand == shorthand {
					if i+1 < len(args) {
						flag.Value = args[i+1]
						i++ // 跳过值
					}
					break
				}
			}
		} else {
			parsedArgs = append(parsedArgs, arg)
		}
	}

	return parsedArgs, nil
}

// showHelp 显示帮助信息
func (c *CobraLikeCommand) showHelp() error {
	fmt.Printf("Usage: %s [flags] [args]\n\n", c.Name)
	fmt.Printf("Description: %s\n\n", c.Description)

	if len(c.Flags) > 0 {
		fmt.Println("Flags:")
		for _, flag := range c.Flags {
			shorthand := ""
			if flag.Shorthand != "" {
				shorthand = fmt.Sprintf(", -%s", flag.Shorthand)
			}
			fmt.Printf("  --%s%s\t%s (default: %v)\n",
				flag.Name, shorthand, flag.Description, flag.Value)
		}
		fmt.Println()
	}

	if len(c.SubCommands) > 0 {
		fmt.Println("Available Commands:")
		for _, subCmd := range c.SubCommands {
			fmt.Printf("  %s\t%s\n", subCmd.Name, subCmd.Description)
		}
		fmt.Println()
	}

	return nil
}

// GetStringFlag 获取字符串标志值
func (c *CobraLikeCommand) GetStringFlag(name string) string {
	if flag, exists := c.Flags[name]; exists {
		if str, ok := flag.Value.(string); ok {
			return str
		}
	}
	return ""
}

// GetIntFlag 获取整数标志值
func (c *CobraLikeCommand) GetIntFlag(name string) int {
	if flag, exists := c.Flags[name]; exists {
		if i, ok := flag.Value.(int); ok {
			return i
		}
	}
	return 0
}

// GetBoolFlag 获取布尔标志值
func (c *CobraLikeCommand) GetBoolFlag(name string) bool {
	if flag, exists := c.Flags[name]; exists {
		if b, ok := flag.Value.(bool); ok {
			return b
		}
	}
	return false
}

// TestifyLikeAssert 模拟Testify风格的断言
type TestifyLikeAssert struct {
	t TestingT
}

// TestingT 测试接口
type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
}

// NewAssert 创建断言实例
func NewAssert(t TestingT) *TestifyLikeAssert {
	return &TestifyLikeAssert{t: t}
}

// Equal 断言相等
func (a *TestifyLikeAssert) Equal(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if expected != actual {
		msg := "断言失败: 期望值与实际值不相等"
		if len(msgAndArgs) > 0 {
			msg = fmt.Sprintf("%v", msgAndArgs[0])
		}
		a.t.Errorf("%s\n期望: %v\n实际: %v", msg, expected, actual)
		return false
	}
	return true
}

// NotEqual 断言不相等
func (a *TestifyLikeAssert) NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if expected == actual {
		msg := "断言失败: 期望值与实际值相等"
		if len(msgAndArgs) > 0 {
			msg = fmt.Sprintf("%v", msgAndArgs[0])
		}
		a.t.Errorf("%s\n不期望: %v\n实际: %v", msg, expected, actual)
		return false
	}
	return true
}

// Nil 断言为nil
func (a *TestifyLikeAssert) Nil(object interface{}, msgAndArgs ...interface{}) bool {
	if object != nil {
		msg := "断言失败: 期望为nil"
		if len(msgAndArgs) > 0 {
			msg = fmt.Sprintf("%v", msgAndArgs[0])
		}
		a.t.Errorf("%s\n实际: %v", msg, object)
		return false
	}
	return true
}

// NotNil 断言不为nil
func (a *TestifyLikeAssert) NotNil(object interface{}, msgAndArgs ...interface{}) bool {
	if object == nil {
		msg := "断言失败: 期望不为nil"
		if len(msgAndArgs) > 0 {
			msg = fmt.Sprintf("%v", msgAndArgs[0])
		}
		a.t.Errorf("%s", msg)
		return false
	}
	return true
}

// True 断言为true
func (a *TestifyLikeAssert) True(value bool, msgAndArgs ...interface{}) bool {
	if !value {
		msg := "断言失败: 期望为true"
		if len(msgAndArgs) > 0 {
			msg = fmt.Sprintf("%v", msgAndArgs[0])
		}
		a.t.Errorf("%s", msg)
		return false
	}
	return true
}

// False 断言为false
func (a *TestifyLikeAssert) False(value bool, msgAndArgs ...interface{}) bool {
	if value {
		msg := "断言失败: 期望为false"
		if len(msgAndArgs) > 0 {
			msg = fmt.Sprintf("%v", msgAndArgs[0])
		}
		a.t.Errorf("%s", msg)
		return false
	}
	return true
}

// PopularLibrariesExamples 流行库示例
func PopularLibrariesExamples() {
	fmt.Println("⭐ Go生态系统 - 流行库使用指南")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🎯 学习目标: 了解Go社区最受欢迎的库和使用模式")
	fmt.Println()
	fmt.Println("📚 本节涵盖的库类型:")
	fmt.Println("   🪵 日志库: Logrus风格的结构化日志")
	fmt.Println("   ⚙️  配置库: Viper风格的配置管理")
	fmt.Println("   💻 CLI库: Cobra风格的命令行工具")
	fmt.Println("   🧪 测试库: Testify风格的断言框架")
	fmt.Println()
	fmt.Println("💡 学习方式: 通过模拟实现理解设计思路")
	fmt.Println("🔗 实际使用: 学会后可直接使用真实库")
	fmt.Println()

	// Logrus风格日志示例
	fmt.Println("🔹 1. 结构化日志 (模拟Logrus)")
	fmt.Println("💡 核心理念: 结构化数据 + 链式调用 + 多级别日志")

	logger := NewLogger()
	logger.Info("应用程序启动")

	logger.WithField("user_id", 12345).
		WithField("action", "login").
		Info("用户登录")

	logger.WithFields(map[string]interface{}{
		"module": "database",
		"query":  "SELECT * FROM users",
		"time":   "150ms",
	}).Info("数据库查询完成")

	logger.WithField("error", "connection timeout").
		Error("数据库连接失败")

	// Viper风格配置示例
	fmt.Println("\n🔹 配置管理示例 (Viper风格):")

	config := NewConfig()

	// 设置默认值
	config.SetDefault("server.port", 8080)
	config.SetDefault("server.host", "localhost")
	config.SetDefault("database.driver", "mysql")
	config.SetDefault("debug", false)

	// 设置配置值
	config.Set("app.name", "Go示例应用")
	config.Set("app.version", "1.0.0")
	config.Set("server.port", 9000)
	config.Set("debug", true)

	// 读取配置值
	fmt.Printf("  应用名称: %s\n", config.GetString("app.name"))
	fmt.Printf("  应用版本: %s\n", config.GetString("app.version"))
	fmt.Printf("  服务器端口: %d\n", config.GetInt("server.port"))
	fmt.Printf("  服务器主机: %s\n", config.GetString("server.host"))
	fmt.Printf("  调试模式: %t\n", config.GetBool("debug"))

	// 保存配置到文件
	configFile := filepath.Join(os.TempDir(), "app_config.json")
	if err := config.WriteConfig(configFile); err != nil {
		fmt.Printf("  保存配置失败: %v\n", err)
	} else {
		fmt.Printf("  配置已保存到: %s\n", configFile)

		// 从文件读取配置
		newConfig := NewConfig()
		if err := newConfig.ReadInConfig(configFile); err != nil {
			fmt.Printf("  读取配置失败: %v\n", err)
		} else {
			fmt.Printf("  从文件读取的应用名称: %s\n", newConfig.GetString("app.name"))
		}

		// 清理
		os.Remove(configFile)
	}

	// Cobra风格CLI示例
	fmt.Println("\n🔹 命令行工具示例 (Cobra风格):")

	rootCmd := NewCommand("myapp", "一个示例CLI应用程序")

	// 添加标志
	rootCmd.StringFlag("config", "c", "config.json", "配置文件路径")
	rootCmd.BoolFlag("verbose", "v", false, "详细输出")
	rootCmd.IntFlag("port", "p", 8080, "服务器端口")

	// 添加子命令
	serveCmd := NewCommand("serve", "启动HTTP服务器")
	serveCmd.RunFunc = func(cmd *CobraLikeCommand, args []string) error {
		port := cmd.GetIntFlag("port")
		verbose := cmd.GetBoolFlag("verbose")

		fmt.Printf("  启动服务器在端口 %d\n", port)
		if verbose {
			fmt.Println("  详细模式已启用")
		}
		return nil
	}

	versionCmd := NewCommand("version", "显示版本信息")
	versionCmd.RunFunc = func(cmd *CobraLikeCommand, args []string) error {
		fmt.Println("  myapp version 1.0.0")
		return nil
	}

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)

	// 模拟命令行参数
	fmt.Println("  执行: myapp serve --port 9000 --verbose")
	if err := rootCmd.Execute([]string{"serve", "--port", "9000", "--verbose"}); err != nil {
		log.Printf("命令执行失败: %v", err)
	}

	fmt.Println("\n  执行: myapp version")
	if err := rootCmd.Execute([]string{"version"}); err != nil {
		log.Printf("命令执行失败: %v", err)
	}

	fmt.Println("\n🎉 Go生态系统学习完成！")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("🎓 您已经了解了:")
	fmt.Println("   ✅ 结构化日志的设计模式")
	fmt.Println("   ✅ 配置管理的最佳实践")
	fmt.Println("   ✅ CLI工具的开发框架")
	fmt.Println("   ✅ 测试断言的使用方法")
	fmt.Println()
	fmt.Println("📦 真实库安装命令:")
	fmt.Println("   🪵 日志: go get github.com/sirupsen/logrus")
	fmt.Println("   ⚙️  配置: go get github.com/spf13/viper")
	fmt.Println("   💻 CLI: go get github.com/spf13/cobra")
	fmt.Println("   🧪 测试: go get github.com/stretchr/testify")
	fmt.Println()
	fmt.Println("🌟 更多优秀库推荐:")
	fmt.Println("   • Gin/Echo - Web框架")
	fmt.Println("   • GORM - ORM框架")
	fmt.Println("   • Redis - 缓存客户端")
	fmt.Println("   • Prometheus - 监控指标")
	fmt.Println()
	fmt.Println("💡 选择库的原则:")
	fmt.Println("   \"选择活跃维护、文档完善、社区认可的库\"")
	fmt.Println("   \"理解设计思想比记住API更重要\"")
	fmt.Println()
	fmt.Println("🚀 下一步: 在实际项目中应用这些库！")
}
