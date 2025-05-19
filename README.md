# Golang 学习资源库

这个仓库包含了 Golang 编程语言的示例代码集合，从基础语法到高级特性，以及常用的生态系统工具和库。

## 目录结构

### 1. 基础语法 (`01-basics/`)
- 变量与常量
- 数据类型
- 控制流
- 函数
- 数组、切片和映射
- 结构体与方法
- 指针
- 错误处理

### 2. 高级特性 (`02-advanced-features/`)
- 接口
- 协程 (Goroutines)
- 通道 (Channels)
- 并发模式
- 泛型 (Go 1.18+)
- 反射
- 测试

### 3. 生态系统 (`03-ecosystem/`)
- 标准库
  - HTTP 服务器
  - JSON 处理
  - 文件 I/O
  - Context 包
  - 时间处理
  - 执行外部命令
- 常用第三方库
  - Gin Web 框架
  - GORM ORM 库
  - Cobra CLI 工具
  - Zap 日志库

## 如何使用

每个目录中的示例都是独立的，可以单独运行。大多数示例都包含了 `main.go` 文件，可以通过以下命令直接运行：

```bash
cd <目录路径>
go run main.go
```

对于一些特定的示例，可能需要安装依赖：

```bash
go mod tidy
```

## 贡献

欢迎贡献更多的示例和改进现有的示例。请提交 Pull Request 或创建 Issue。

## 许可证

MIT 