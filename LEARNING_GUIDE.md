# Go语言学习指南

欢迎来到Go语言学习之旅！这个项目提供了从基础到高级的完整Go语言学习资源。

## 🎯 学习目标

通过这个项目，你将学会：
- Go语言的基础语法和概念
- 面向对象编程在Go中的实现
- 并发编程和goroutines
- 错误处理最佳实践
- 实际项目开发技巧

## 📚 学习路径

### 第一阶段：基础语法 (1-2周)

#### 1. 变量和常量 (`01-basics/01-variables-constants/`)
- **学习内容**：变量声明、常量定义、作用域
- **重点概念**：零值、类型推断、iota
- **实践**：运行测试 `go test -v ./01-basics/01-variables-constants/...`

#### 2. 格式化输出 (`01-basics/02-format/`)
- **学习内容**：fmt包的使用、格式化动词
- **重点概念**：Printf、Sprintf、格式化占位符
- **实践**：修改示例代码，尝试不同的格式化选项

#### 3. 数据类型 (`01-basics/03-data-types/`)
- **学习内容**：基本类型、复合类型、类型转换
- **重点概念**：数值类型、字符串、布尔值、类型别名
- **实践**：创建自己的类型别名和转换函数

#### 4. 控制流程 (`01-basics/04-control-flow/`)
- **学习内容**：条件语句、循环、分支
- **重点概念**：if/else、switch、for、range
- **实践**：实现简单的算法（排序、查找）

#### 5. 函数 (`01-basics/05-functions/`)
- **学习内容**：函数定义、参数、返回值
- **重点概念**：多返回值、可变参数、闭包、递归
- **实践**：编写工具函数库

### 第二阶段：数据结构 (1-2周)

#### 6. 数组、切片和映射 (`01-basics/06-arrays-slices-maps/`)
- **学习内容**：Go的核心数据结构
- **重点概念**：数组vs切片、映射操作、内存管理
- **实践**：实现数据处理算法

#### 7. 结构体和方法 (`01-basics/07-structs-methods/`)
- **学习内容**：自定义类型、方法定义
- **重点概念**：值接收者vs指针接收者、嵌入
- **实践**：设计一个简单的对象模型

#### 8. 指针 (`01-basics/08-pointers/`)
- **学习内容**：内存地址、指针操作
- **重点概念**：指针安全、内存分配
- **实践**：理解值传递vs引用传递

#### 9. 错误处理 (`01-basics/09-error-handling/`)
- **学习内容**：Go的错误处理哲学
- **重点概念**：error接口、错误包装、panic/recover
- **实践**：构建健壮的错误处理机制

## 🚀 快速开始

### 1. 环境准备
```bash
# 确保Go已安装
go version

# 克隆或下载项目
cd golang-examples

# 运行演示
go run main.go demo
```

### 2. 运行测试
```bash
# 运行所有测试
go test ./...

# 运行特定模块测试（详细输出）
go test -v ./01-basics/03-data-types/...

# 运行基准测试
go test -bench=. ./01-basics/05-functions/...

# 查看测试覆盖率
go test -cover ./01-basics/03-data-types/...
```

### 3. 学习建议

#### 每日学习计划
1. **阅读代码**：先阅读主要的.go文件，理解概念
2. **运行测试**：执行测试，观察输出结果
3. **修改代码**：尝试修改示例，观察变化
4. **编写代码**：基于学到的概念编写自己的代码
5. **解决问题**：尝试解决测试中的练习题

#### 学习技巧
- **动手实践**：不要只看代码，要亲自运行和修改
- **理解错误**：当代码出错时，仔细阅读错误信息
- **查阅文档**：使用 `go doc` 命令查看官方文档
- **写测试**：为自己的代码编写测试用例
- **代码审查**：对比自己的代码和示例代码

## 📖 学习资源

### 官方资源
- [Go官方文档](https://golang.org/doc/)
- [Go语言规范](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go.html)

### 推荐书籍
- 《Go语言圣经》
- 《Go语言实战》
- 《Go并发编程实战》

### 在线资源
- [Go Playground](https://play.golang.org/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)

## 🔧 开发工具

### 推荐IDE
- **VS Code** + Go扩展
- **GoLand** (JetBrains)
- **Vim/Neovim** + vim-go

### 有用的命令
```bash
# 格式化代码
go fmt ./...

# 检查代码
go vet ./...

# 下载依赖
go mod tidy

# 查看文档
go doc fmt.Printf

# 运行代码
go run main.go

# 构建可执行文件
go build
```

## 🎯 练习项目

完成基础学习后，尝试这些项目：

### 初级项目
1. **计算器**：实现基本的数学运算
2. **待办事项**：命令行TODO应用
3. **文件处理器**：读取和处理文本文件

### 中级项目
1. **HTTP服务器**：简单的Web API
2. **数据库应用**：CRUD操作
3. **并发下载器**：多线程文件下载

### 高级项目
1. **微服务**：分布式系统
2. **CLI工具**：命令行工具开发
3. **性能监控**：系统监控工具

## 🤝 学习社区

### 获取帮助
- [Go官方论坛](https://forum.golangbridge.org/)
- [Reddit r/golang](https://www.reddit.com/r/golang/)
- [Stack Overflow](https://stackoverflow.com/questions/tagged/go)

### 贡献代码
- 发现bug？提交issue
- 有改进建议？提交PR
- 想要新功能？讨论需求

## 📝 学习记录

建议创建学习日志，记录：
- 每日学习内容
- 遇到的问题和解决方案
- 代码片段和笔记
- 项目想法和实现

## 🎉 下一步

完成基础学习后，可以继续学习：
- 并发编程（goroutines, channels）
- 网络编程（HTTP, TCP/UDP）
- 数据库操作（SQL, NoSQL）
- 微服务架构
- 性能优化
- 部署和运维

祝你学习愉快！记住，编程最重要的是实践。多写代码，多解决问题，你会很快掌握Go语言的精髓。
