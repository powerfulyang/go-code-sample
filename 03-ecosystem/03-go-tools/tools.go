package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// GoToolchain Go工具链管理器
type GoToolchain struct {
	goPath    string
	goRoot    string
	goVersion string
	workDir   string
}

// NewGoToolchain 创建Go工具链管理器
func NewGoToolchain() *GoToolchain {
	return &GoToolchain{
		goPath:    os.Getenv("GOPATH"),
		goRoot:    os.Getenv("GOROOT"),
		goVersion: runtime.Version(),
		workDir:   ".",
	}
}

// SetWorkDir 设置工作目录
func (gt *GoToolchain) SetWorkDir(dir string) {
	gt.workDir = dir
}

// GetInfo 获取Go环境信息
func (gt *GoToolchain) GetInfo() map[string]string {
	info := map[string]string{
		"GOPATH":    gt.goPath,
		"GOROOT":    gt.goRoot,
		"GOVERSION": gt.goVersion,
		"GOOS":      runtime.GOOS,
		"GOARCH":    runtime.GOARCH,
		"WORKDIR":   gt.workDir,
	}

	// 获取更多环境信息
	if output, err := gt.runCommand("go", "env"); err == nil {
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "=") {
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.Trim(strings.TrimSpace(parts[1]), `"`)
					info[key] = value
				}
			}
		}
	}

	return info
}

// runCommand 执行命令
func (gt *GoToolchain) runCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = gt.workDir

	output, err := cmd.CombinedOutput()
	return string(output), err
}

// Build 构建项目
func (gt *GoToolchain) Build(packagePath string, outputPath string) error {
	args := []string{"build"}
	if outputPath != "" {
		args = append(args, "-o", outputPath)
	}
	args = append(args, packagePath)

	_, err := gt.runCommand("go", args...)
	return err
}

// Test 运行测试
func (gt *GoToolchain) Test(packagePath string, verbose bool) (string, error) {
	args := []string{"test"}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, packagePath)

	return gt.runCommand("go", args...)
}

// Benchmark 运行基准测试
func (gt *GoToolchain) Benchmark(packagePath string) (string, error) {
	args := []string{"test", "-bench=.", packagePath}
	return gt.runCommand("go", args...)
}

// Format 格式化代码
func (gt *GoToolchain) Format(filePath string) error {
	_, err := gt.runCommand("gofmt", "-w", filePath)
	return err
}

// Vet 代码检查
func (gt *GoToolchain) Vet(packagePath string) (string, error) {
	return gt.runCommand("go", "vet", packagePath)
}

// ModInit 初始化模块
func (gt *GoToolchain) ModInit(moduleName string) error {
	_, err := gt.runCommand("go", "mod", "init", moduleName)
	return err
}

// ModTidy 整理模块依赖
func (gt *GoToolchain) ModTidy() error {
	_, err := gt.runCommand("go", "mod", "tidy")
	return err
}

// ModDownload 下载依赖
func (gt *GoToolchain) ModDownload() error {
	_, err := gt.runCommand("go", "mod", "download")
	return err
}

// Get 获取包
func (gt *GoToolchain) Get(packagePath string) error {
	_, err := gt.runCommand("go", "get", packagePath)
	return err
}

// Install 安装包
func (gt *GoToolchain) Install(packagePath string) error {
	_, err := gt.runCommand("go", "install", packagePath)
	return err
}

// Clean 清理构建缓存
func (gt *GoToolchain) Clean() error {
	_, err := gt.runCommand("go", "clean", "-cache")
	return err
}

// Doc 查看文档
func (gt *GoToolchain) Doc(packagePath string) (string, error) {
	return gt.runCommand("go", "doc", packagePath)
}

// List 列出包
func (gt *GoToolchain) List(pattern string) (string, error) {
	return gt.runCommand("go", "list", pattern)
}

// Version 获取Go版本
func (gt *GoToolchain) Version() (string, error) {
	return gt.runCommand("go", "version")
}

// ProjectManager 项目管理器
type ProjectManager struct {
	toolchain   *GoToolchain
	projectPath string
}

// NewProjectManager 创建项目管理器
func NewProjectManager(projectPath string) *ProjectManager {
	return &ProjectManager{
		toolchain:   NewGoToolchain(),
		projectPath: projectPath,
	}
}

// CreateProject 创建新项目
func (pm *ProjectManager) CreateProject(moduleName string) error {
	// 创建项目目录
	if err := os.MkdirAll(pm.projectPath, 0755); err != nil {
		return fmt.Errorf("创建项目目录失败: %v", err)
	}

	// 设置工作目录
	pm.toolchain.SetWorkDir(pm.projectPath)

	// 初始化模块
	if err := pm.toolchain.ModInit(moduleName); err != nil {
		return fmt.Errorf("初始化模块失败: %v", err)
	}

	// 创建基本文件结构
	if err := pm.createBasicStructure(); err != nil {
		return fmt.Errorf("创建基本结构失败: %v", err)
	}

	return nil
}

// createBasicStructure 创建基本项目结构
func (pm *ProjectManager) createBasicStructure() error {
	// 创建main.go
	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
`
	if err := pm.writeFile("main.go", mainContent); err != nil {
		return err
	}

	// 创建README.md
	readmeContent := `# Go Project

This is a Go project created with Go toolchain.

## Usage

` + "```bash" + `
go run main.go
` + "```" + `

## Build

` + "```bash" + `
go build
` + "```" + `

## Test

` + "```bash" + `
go test ./...
` + "```" + `
`
	if err := pm.writeFile("README.md", readmeContent); err != nil {
		return err
	}

	// 创建.gitignore
	gitignoreContent := `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work
`
	if err := pm.writeFile(".gitignore", gitignoreContent); err != nil {
		return err
	}

	return nil
}

// writeFile 写入文件
func (pm *ProjectManager) writeFile(filename, content string) error {
	filePath := filepath.Join(pm.projectPath, filename)
	return os.WriteFile(filePath, []byte(content), 0644)
}

// BuildProject 构建项目
func (pm *ProjectManager) BuildProject() error {
	pm.toolchain.SetWorkDir(pm.projectPath)
	return pm.toolchain.Build(".", "")
}

// TestProject 测试项目
func (pm *ProjectManager) TestProject() (string, error) {
	pm.toolchain.SetWorkDir(pm.projectPath)
	return pm.toolchain.Test("./...", true)
}

// FormatProject 格式化项目代码
func (pm *ProjectManager) FormatProject() error {
	pm.toolchain.SetWorkDir(pm.projectPath)

	// 遍历所有Go文件
	return filepath.Walk(pm.projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".go") {
			return pm.toolchain.Format(path)
		}

		return nil
	})
}

// DevelopmentWorkflow 开发工作流示例
type DevelopmentWorkflow struct {
	toolchain   *GoToolchain
	projectPath string
}

// NewDevelopmentWorkflow 创建开发工作流
func NewDevelopmentWorkflow(projectPath string) *DevelopmentWorkflow {
	return &DevelopmentWorkflow{
		toolchain:   NewGoToolchain(),
		projectPath: projectPath,
	}
}

// RunWorkflow 运行完整的开发工作流
func (dw *DevelopmentWorkflow) RunWorkflow() error {
	dw.toolchain.SetWorkDir(dw.projectPath)

	fmt.Println("🚀 开始开发工作流...")

	// 1. 格式化代码
	fmt.Println("📝 格式化代码...")
	if err := dw.formatCode(); err != nil {
		return fmt.Errorf("格式化失败: %v", err)
	}

	// 2. 代码检查
	fmt.Println("🔍 代码检查...")
	if err := dw.vetCode(); err != nil {
		fmt.Printf("⚠️  代码检查警告: %v\n", err)
	}

	// 3. 运行测试
	fmt.Println("🧪 运行测试...")
	if err := dw.runTests(); err != nil {
		return fmt.Errorf("测试失败: %v", err)
	}

	// 4. 构建项目
	fmt.Println("🔨 构建项目...")
	if err := dw.buildProject(); err != nil {
		return fmt.Errorf("构建失败: %v", err)
	}

	// 5. 整理依赖
	fmt.Println("📦 整理依赖...")
	if err := dw.tidyDependencies(); err != nil {
		return fmt.Errorf("整理依赖失败: %v", err)
	}

	fmt.Println("✅ 开发工作流完成!")
	return nil
}

// formatCode 格式化代码
func (dw *DevelopmentWorkflow) formatCode() error {
	return filepath.Walk(dw.projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".go") {
			return dw.toolchain.Format(path)
		}

		return nil
	})
}

// vetCode 代码检查
func (dw *DevelopmentWorkflow) vetCode() error {
	_, err := dw.toolchain.Vet("./...")
	return err
}

// runTests 运行测试
func (dw *DevelopmentWorkflow) runTests() error {
	_, err := dw.toolchain.Test("./...", false)
	return err
}

// buildProject 构建项目
func (dw *DevelopmentWorkflow) buildProject() error {
	return dw.toolchain.Build(".", "")
}

// tidyDependencies 整理依赖
func (dw *DevelopmentWorkflow) tidyDependencies() error {
	return dw.toolchain.ModTidy()
}

// GoToolsExamples Go工具链示例
func GoToolsExamples() {
	fmt.Println("🔨 Go工具链 - 高效开发的利器")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🎯 学习目标: 掌握Go开发工具链的使用")
	fmt.Println()
	fmt.Println("🛠️ 核心工具:")
	fmt.Println("   • go build/run - 编译和运行")
	fmt.Println("   • go test/bench - 测试和基准")
	fmt.Println("   • go mod - 模块依赖管理")
	fmt.Println("   • go fmt/vet - 代码格式化和检查")
	fmt.Println("   • go doc - 文档生成和查看")
	fmt.Println()
	fmt.Println("💼 开发流程: 编码 → 格式化 → 测试 → 构建 → 部署")
	fmt.Println()

	// 创建工具链管理器
	toolchain := NewGoToolchain()

	// 显示Go环境信息
	fmt.Println("\n🔹 Go环境信息:")
	info := toolchain.GetInfo()
	for key, value := range info {
		if value != "" && len(key) <= 10 { // 只显示主要信息
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	// 获取Go版本
	fmt.Println("\n🔹 Go版本信息:")
	if version, err := toolchain.Version(); err == nil {
		fmt.Printf("  %s\n", strings.TrimSpace(version))
	}

	// 项目管理示例
	fmt.Println("\n🔹 项目管理示例:")

	// 创建临时项目目录
	tempDir := filepath.Join(os.TempDir(), "go-tools-example", fmt.Sprintf("%d", time.Now().Unix()))
	defer os.RemoveAll(tempDir) // 清理

	pm := NewProjectManager(tempDir)

	fmt.Printf("  创建项目目录: %s\n", tempDir)
	if err := pm.CreateProject("example.com/go-tools-demo"); err != nil {
		fmt.Printf("  ❌ 创建项目失败: %v\n", err)
	} else {
		fmt.Println("  ✅ 项目创建成功")

		// 构建项目
		fmt.Println("  🔨 构建项目...")
		if err := pm.BuildProject(); err != nil {
			fmt.Printf("  ❌ 构建失败: %v\n", err)
		} else {
			fmt.Println("  ✅ 构建成功")
		}

		// 格式化代码
		fmt.Println("  📝 格式化代码...")
		if err := pm.FormatProject(); err != nil {
			fmt.Printf("  ❌ 格式化失败: %v\n", err)
		} else {
			fmt.Println("  ✅ 格式化完成")
		}
	}

	// 开发工作流示例
	fmt.Println("\n🔹 开发工作流示例:")
	workflow := NewDevelopmentWorkflow(tempDir)

	if err := workflow.RunWorkflow(); err != nil {
		fmt.Printf("  ❌ 工作流失败: %v\n", err)
	}

	fmt.Println("\n🎉 Go工具链学习完成！")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("🎓 您已经掌握了:")
	fmt.Println("   ✅ Go环境信息查看")
	fmt.Println("   ✅ 项目创建和管理")
	fmt.Println("   ✅ 代码格式化和检查")
	fmt.Println("   ✅ 完整的开发工作流")
	fmt.Println()
	fmt.Println("🔧 常用命令速查:")
	fmt.Println("   📖 帮助: go help [command]")
	fmt.Println("   🌍 环境: go env")
	fmt.Println("   📦 模块: go mod init/tidy/download")
	fmt.Println("   🔨 构建: go build/install")
	fmt.Println("   🧪 测试: go test -v ./...")
	fmt.Println("   📊 基准: go test -bench=.")
	fmt.Println("   📝 格式: go fmt ./...")
	fmt.Println("   🔍 检查: go vet ./...")
	fmt.Println()
	fmt.Println("💡 效率提升技巧:")
	fmt.Println("   • 使用IDE集成工具链")
	fmt.Println("   • 配置git hooks自动格式化")
	fmt.Println("   • 编写Makefile简化命令")
	fmt.Println("   • 使用go generate自动生成代码")
	fmt.Println()
	fmt.Println("🚀 工具链是Go开发效率的基石！")
}
