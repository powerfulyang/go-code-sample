package packages

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// PackageInfo 包信息
type PackageInfo struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Files       []string `json:"files"`
	Imports     []string `json:"imports"`
	Exports     []string `json:"exports"`
	Description string   `json:"description"`
}

// ModuleInfo 模块信息
type ModuleInfo struct {
	Name         string                 `json:"name"`
	Version      string                 `json:"version"`
	GoVersion    string                 `json:"go_version"`
	Packages     map[string]PackageInfo `json:"packages"`
	Dependencies []string               `json:"dependencies"`
}

// PackageManager 包管理器
type PackageManager struct {
	rootPath string
	modules  map[string]*ModuleInfo
}

// NewPackageManager 创建包管理器
func NewPackageManager(rootPath string) *PackageManager {
	return &PackageManager{
		rootPath: rootPath,
		modules:  make(map[string]*ModuleInfo),
	}
}

// AnalyzePackage 分析包
func (pm *PackageManager) AnalyzePackage(packagePath string) (*PackageInfo, error) {
	info := &PackageInfo{
		Path:    packagePath,
		Files:   make([]string, 0),
		Imports: make([]string, 0),
		Exports: make([]string, 0),
	}

	// 检查路径是否存在
	if _, err := os.Stat(packagePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("包路径不存在: %s", packagePath)
	}

	// 遍历Go文件
	err := filepath.Walk(packagePath, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			return pm.analyzeGoFile(path, info)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("分析包失败: %v", err)
	}

	return info, nil
}

// analyzeGoFile 分析Go文件
func (pm *PackageManager) analyzeGoFile(filePath string, pkgInfo *PackageInfo) error {
	// 解析Go文件
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("解析文件失败 %s: %v", filePath, err)
	}

	// 获取包名
	if pkgInfo.Name == "" {
		pkgInfo.Name = node.Name.Name
	}

	// 添加文件
	pkgInfo.Files = append(pkgInfo.Files, filepath.Base(filePath))

	// 分析导入
	for _, imp := range node.Imports {
		importPath := strings.Trim(imp.Path.Value, `"`)
		if !contains(pkgInfo.Imports, importPath) {
			pkgInfo.Imports = append(pkgInfo.Imports, importPath)
		}
	}

	// 分析导出的标识符
	for _, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name.IsExported() {
				export := fmt.Sprintf("func %s", d.Name.Name)
				if !contains(pkgInfo.Exports, export) {
					pkgInfo.Exports = append(pkgInfo.Exports, export)
				}
			}
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if s.Name.IsExported() {
						export := fmt.Sprintf("type %s", s.Name.Name)
						if !contains(pkgInfo.Exports, export) {
							pkgInfo.Exports = append(pkgInfo.Exports, export)
						}
					}
				case *ast.ValueSpec:
					for _, name := range s.Names {
						if name.IsExported() {
							export := fmt.Sprintf("var %s", name.Name)
							if !contains(pkgInfo.Exports, export) {
								pkgInfo.Exports = append(pkgInfo.Exports, export)
							}
						}
					}
				}
			}
		}
	}

	return nil
}

// contains 检查切片是否包含元素
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// PackageVisibility 包可见性示例
type PackageVisibility struct {
	// 导出的字段（首字母大写）
	PublicField   string
	ExportedValue int

	// 未导出的字段（首字母小写）
	privateField  string
	internalValue int
}

// NewPackageVisibility 构造函数（导出）
func NewPackageVisibility(public string, exported int) *PackageVisibility {
	return &PackageVisibility{
		PublicField:   public,
		ExportedValue: exported,
		privateField:  "internal",
		internalValue: 42,
	}
}

// GetPrivateField 获取私有字段（导出方法）
func (pv *PackageVisibility) GetPrivateField() string {
	return pv.privateField
}

// setInternalValue 设置内部值（未导出方法）
func (pv *PackageVisibility) setInternalValue(value int) {
	pv.internalValue = value
}

// UpdateInternal 更新内部值（导出方法调用未导出方法）
func (pv *PackageVisibility) UpdateInternal(value int) {
	pv.setInternalValue(value)
}

// 包级别的变量和常量
var (
	// 导出的包变量
	DefaultTimeout = 30
	MaxRetries     = 3

	// 未导出的包变量
	internalConfig = "default"
	debugMode      = false
)

// 导出的包常量
const (
	Version     = "1.0.0"
	Author      = "Go Developer"
	MaxFileSize = 1024 * 1024 // 1MB
)

// 未导出的包常量
const (
	bufferSize    = 4096
	retryInterval = 100
)

// 导出的包函数
func GetVersion() string {
	return Version
}

func IsDebugMode() bool {
	return debugMode
}

// 未导出的包函数
func validateInput(input string) bool {
	return len(input) > 0
}

func processData(data []byte) []byte {
	// 简单的数据处理
	return data
}

// 导出的函数调用未导出的函数
func ProcessInput(input string) ([]byte, error) {
	if !validateInput(input) {
		return nil, fmt.Errorf("无效输入")
	}

	data := []byte(input)
	return processData(data), nil
}

// 包初始化函数
func init() {
	fmt.Println("包 packages 正在初始化...")

	// 可以在这里进行包级别的初始化
	if os.Getenv("DEBUG") == "true" {
		debugMode = true
	}

	// 设置默认配置
	if config := os.Getenv("CONFIG"); config != "" {
		internalConfig = config
	}
}

// 类型别名示例
type (
	// 导出的类型别名
	UserID   int64
	Username string

	// 未导出的类型别名
	sessionID string
	timestamp int64
)

// 接口定义示例
type Processor interface {
	Process(data []byte) ([]byte, error)
	Validate(input string) bool
}

// 接口实现示例
type DefaultProcessor struct {
	config string
}

func (dp *DefaultProcessor) Process(data []byte) ([]byte, error) {
	// 实现处理逻辑
	return processData(data), nil
}

func (dp *DefaultProcessor) Validate(input string) bool {
	return validateInput(input)
}

// 工厂函数
func NewProcessor(config string) Processor {
	return &DefaultProcessor{config: config}
}

// PackageExamples 包和模块示例
func PackageExamples() {
	fmt.Println("📦 Go语言包和模块系统 - 代码组织的艺术")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🎯 学习目标: 掌握Go语言的模块化设计思想")
	fmt.Println()
	fmt.Println("📚 本节内容:")
	fmt.Println("   • 包的可见性规则 (public/private)")
	fmt.Println("   • 包级别变量和常量")
	fmt.Println("   • 类型别名和接口设计")
	fmt.Println("   • init函数的执行机制")
	fmt.Println("   • 包的最佳实践")
	fmt.Println()

	// 包可见性示例
	fmt.Println("🔹 1. 包可见性规则 (Go的访问控制)")
	fmt.Println("💡 核心概念: 首字母大写=公开(exported)，小写=私有(unexported)")

	pv := NewPackageVisibility("公开数据", 100)
	fmt.Printf("公开字段: %s\n", pv.PublicField)
	fmt.Printf("导出值: %d\n", pv.ExportedValue)
	fmt.Printf("私有字段(通过方法访问): %s\n", pv.GetPrivateField())

	// 更新内部值
	pv.UpdateInternal(200)
	fmt.Println("内部值已更新")

	// 包级别变量和常量
	fmt.Println("\n🔹 2. 包级别变量和常量 (全局状态管理)")
	fmt.Println("💡 最佳实践: 使用常量配置，变量保存状态")
	fmt.Printf("   📋 应用版本: %s\n", Version)
	fmt.Printf("   👤 开发作者: %s\n", Author)
	fmt.Printf("   ⏱️  默认超时: %d秒\n", DefaultTimeout)
	fmt.Printf("   🔄 最大重试: %d次\n", MaxRetries)
	fmt.Printf("   🐛 调试模式: %t\n", IsDebugMode())

	// 类型别名使用
	fmt.Println("\n🔹 类型别名示例:")
	var userID UserID = 12345
	var username Username = "gopher"

	fmt.Printf("用户ID: %d\n", userID)
	fmt.Printf("用户名: %s\n", username)

	// 接口使用
	fmt.Println("\n🔹 接口使用示例:")
	processor := NewProcessor("default")

	testInput := "Hello, Go Packages!"
	if processor.Validate(testInput) {
		result, err := processor.Process([]byte(testInput))
		if err != nil {
			fmt.Printf("处理错误: %v\n", err)
		} else {
			fmt.Printf("处理结果: %s\n", result)
		}
	}

	// 包函数使用
	fmt.Println("\n🔹 包函数示例:")
	input := "测试数据"
	processed, err := ProcessInput(input)
	if err != nil {
		fmt.Printf("处理失败: %v\n", err)
	} else {
		fmt.Printf("处理成功: %s\n", processed)
	}

	fmt.Println("\n🎉 包和模块系统学习完成！")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("🎓 您已经掌握了:")
	fmt.Println("   ✅ Go语言的可见性规则")
	fmt.Println("   ✅ 包级别变量和常量的使用")
	fmt.Println("   ✅ 类型别名和接口设计")
	fmt.Println("   ✅ init函数的执行机制")
	fmt.Println()
	fmt.Println("🚀 下一步建议:")
	fmt.Println("   📖 学习 go mod 命令管理依赖")
	fmt.Println("   🔨 实践创建自己的包")
	fmt.Println("   📚 阅读标准库源码学习设计")
	fmt.Println("   💼 在实际项目中应用模块化设计")
	fmt.Println()
	fmt.Println("💡 记住: 好的包设计是Go程序的基础！")
}
