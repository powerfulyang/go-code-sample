# Golang Examples Repository

这是一个全面的 Go 语言学习资源库，包含从基础语法到高级特性以及常用生态系统的实用示例。

## 🎯 项目特点

- **系统性学习路径**：从基础到高级的完整学习路径
- **实用代码示例**：每个示例都可以独立运行并解决实际问题
- **完善的测试**：包含单元测试和基准测试示例
- **丰富的注释**：详细的中文注释，便于理解
- **标准项目结构**：遵循 Go 语言最佳实践

## 📁 目录结构

### 01-basics - 基础语法
- `01-variables-constants/` - 变量和常量
- `02-format/` - 格式化输出
- `03-data-types/` - 数据类型（数值、字符串、布尔）
- `04-control-flow/` - 控制流程（条件、循环、选择）
- `05-functions/` - 函数（基础、多返回值、可变参数、闭包）
- `06-arrays-slices-maps/` - 集合类型（数组、切片、映射）
- `07-structs-methods/` - 结构体和方法
- `08-pointers/` - 指针
- `09-error-handling/` - 错误处理
- `10-packages-modules/` - 包和模块系统

### 02-advanced-features - 高级特性
- `01-interfaces/` - 接口
- `02-concurrency/` - 并发编程 (Goroutines + Channels)
- `03-generics/` - 泛型 (Go 1.18+)
- `04-reflection/` - 反射
- `05-testing/` - 测试框架
- `06-performance/` - 性能优化

### 03-ecosystem - 生态系统
- `01-standard-library/` - 标准库示例
- `02-third-party-libraries/` - 第三方库和函数式编程
- `03-go-tools/` - Go工具链
- `04-popular-libraries/` - 流行库使用模式

### 03-practical-examples - 实际应用示例
- `01-package-management/` - 包管理

### 04-practical-applications - 实际应用开发
- `01-web-api/` - Web API开发
- `02-database/` - 数据库操作
- `03-cli-tool/` - CLI工具开发
- `04-network/` - 网络编程 (TCP/UDP/WebSocket)
- `07-security/` - 安全和认证 (JWT/加密)

## 🚀 快速开始

### 运行示例
```bash
# 运行主程序演示
go run main.go demo          # 基础演示
go run main.go interfaces    # 接口示例
go run main.go concurrency   # 并发编程示例
go run main.go generics      # 泛型示例
go run main.go reflection    # 反射示例
go run main.go testing       # 测试框架示例
go run main.go stdlib        # 标准库示例
go run main.go functional    # 函数式编程示例
go run main.go webapi        # Web API示例
go run main.go database      # 数据库操作示例
go run main.go cli           # CLI工具示例
go run main.go network       # 网络编程示例
go run main.go security      # 安全和认证示例
go run main.go packages      # 包和模块系统示例
go run main.go performance   # 性能优化示例
go run main.go tools         # Go工具链示例
go run main.go popular       # 流行库使用示例
go run main.go all           # 运行所有示例

# 运行所有测试
go test ./...

# 运行特定模块的测试
go test -v ./01-basics/03-data-types/...
go test -v ./02-advanced-features/03-generics/...
go test -v ./03-ecosystem/01-standard-library/...

# 运行基准测试
go test -bench=. ./01-basics/05-functions/...
go test -bench=. ./02-advanced-features/03-generics/...

# 查看测试覆盖率
go test -cover ./...
```

### 学习建议
1. **按顺序学习**：建议按照目录编号顺序学习
2. **动手实践**：每个示例都可以直接运行，建议修改代码观察结果
3. **阅读测试**：测试文件包含了丰富的使用示例
4. **运行基准测试**：了解性能特性

## 📚 学习路径

### 初学者路径
1. 01-basics/01-variables-constants - 了解变量和常量
2. 01-basics/02-format - 学习格式化输出
3. 01-basics/03-data-types - 掌握基本数据类型
4. 01-basics/04-control-flow - 学习控制流程
5. 01-basics/05-functions - 理解函数概念

### 进阶路径
1. 01-basics/06-arrays-slices-maps - 掌握集合类型
2. 01-basics/07-structs-methods - 学习面向对象
3. 01-basics/08-pointers - 理解指针概念
4. 02-advanced-features/01-interfaces - 学习接口

### 高级路径
1. 02-advanced-features/02-concurrency - 并发编程 (Goroutines + Channels)
2. 02-advanced-features/03-generics - 泛型编程 (Go 1.18+)
3. 02-advanced-features/04-reflection - 反射编程
4. 02-advanced-features/05-testing - 测试框架和技巧
5. 02-advanced-features/06-performance - 性能优化技巧
6. 01-basics/10-packages-modules - 包和模块系统
7. 03-ecosystem/01-standard-library - 标准库深入
8. 03-ecosystem/02-third-party-libraries - 函数式编程和工具库
9. 03-ecosystem/03-go-tools - Go工具链使用
10. 03-ecosystem/04-popular-libraries - 流行库使用模式

### 实战路径
1. 04-practical-applications/01-web-api - Web API开发
2. 04-practical-applications/02-database - 数据库操作和设计
3. 04-practical-applications/03-cli-tool - CLI工具开发
4. 04-practical-applications/04-network - 网络编程和通信
5. 04-practical-applications/07-security - 安全认证和加密

## 🚀 快速开始

### 运行基础示例
