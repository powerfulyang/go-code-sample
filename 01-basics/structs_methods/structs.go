package main

import "fmt"

// Person 定义一个人的结构体
type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address // 嵌套结构体
}

// Address 地址结构体
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

// Student 学生结构体，演示匿名字段
type Student struct {
	Person    // 匿名字段，嵌入Person
	StudentID string
	Grade     int
	Subjects  []string
}

// Book 书籍结构体，演示标签
type Book struct {
	Title     string  `json:"title" xml:"title"`
	Author    string  `json:"author" xml:"author"`
	ISBN      string  `json:"isbn" xml:"isbn"`
	Price     float64 `json:"price" xml:"price"`
	Published bool    `json:"published" xml:"published"`
}

func runStructExample() {
	fmt.Println("=== Go 结构体示例 ===")

	// 结构体声明和初始化
	fmt.Println("\n--- 结构体声明和初始化 ---")

	// 方式1：零值初始化
	var p1 Person
	fmt.Printf("零值Person: %+v\n", p1)

	// 方式2：字段赋值
	p1.Name = "Alice"
	p1.Age = 25
	p1.Email = "alice@example.com"
	fmt.Printf("赋值后Person: %+v\n", p1)

	// 方式3：结构体字面量（按字段顺序）
	p2 := Person{
		"Bob",
		30,
		"bob@example.com",
		Address{
			"123 Main St",
			"New York",
			"USA",
			"10001",
		},
	}
	fmt.Printf("字面量Person: %+v\n", p2)

	// 方式4：结构体字面量（指定字段名）
	p3 := Person{
		Name:  "Charlie",
		Age:   35,
		Email: "charlie@example.com",
		Address: Address{
			Street:  "456 Oak Ave",
			City:    "Los Angeles",
			Country: "USA",
			ZipCode: "90210",
		},
	}
	fmt.Printf("指定字段Person: %+v\n", p3)

	// 方式5：部分初始化
	p4 := Person{
		Name: "David",
		Age:  28,
		// Email和Address使用零值
	}
	fmt.Printf("部分初始化Person: %+v\n", p4)

	// 结构体字段访问
	fmt.Println("\n--- 结构体字段访问 ---")
	fmt.Printf("p3的姓名: %s\n", p3.Name)
	fmt.Printf("p3的年龄: %d\n", p3.Age)
	fmt.Printf("p3的城市: %s\n", p3.Address.City)

	// 修改字段
	p3.Age = 36
	p3.Address.ZipCode = "90211"
	fmt.Printf("修改后的p3: %+v\n", p3)

	// 结构体指针
	fmt.Println("\n--- 结构体指针 ---")
	p5 := &Person{
		Name:  "Eve",
		Age:   22,
		Email: "eve@example.com",
	}
	fmt.Printf("结构体指针: %+v\n", p5)
	fmt.Printf("通过指针访问姓名: %s\n", p5.Name)    // Go自动解引用
	fmt.Printf("通过指针访问姓名: %s\n", (*p5).Name) // 显式解引用

	// 修改指针指向的结构体
	p5.Age = 23
	fmt.Printf("修改后的结构体: %+v\n", p5)

	// 匿名结构体
	fmt.Println("\n--- 匿名结构体 ---")
	point := struct {
		X, Y int
		Name string
	}{
		X:    10,
		Y:    20,
		Name: "Point A",
	}
	fmt.Printf("匿名结构体: %+v\n", point)

	// 结构体比较
	fmt.Println("\n--- 结构体比较 ---")
	addr1 := Address{
		Street:  "123 Main St",
		City:    "New York",
		Country: "USA",
		ZipCode: "10001",
	}
	addr2 := Address{
		Street:  "123 Main St",
		City:    "New York",
		Country: "USA",
		ZipCode: "10001",
	}
	addr3 := Address{
		Street:  "456 Oak Ave",
		City:    "Los Angeles",
		Country: "USA",
		ZipCode: "90210",
	}

	fmt.Printf("addr1 == addr2: %t\n", addr1 == addr2)
	fmt.Printf("addr1 == addr3: %t\n", addr1 == addr3)

	// 嵌入结构体（匿名字段）
	fmt.Println("\n--- 嵌入结构体 ---")
	student := Student{
		Person: Person{
			Name:  "Frank",
			Age:   20,
			Email: "frank@university.edu",
			Address: Address{
				Street:  "789 College Rd",
				City:    "Boston",
				Country: "USA",
				ZipCode: "02101",
			},
		},
		StudentID: "S12345",
		Grade:     2,
		Subjects:  []string{"Math", "Physics", "Computer Science"},
	}

	fmt.Printf("学生信息: %+v\n", student)

	// 可以直接访问嵌入结构体的字段
	fmt.Printf("学生姓名: %s\n", student.Name) // 等同于 student.Person.Name
	fmt.Printf("学生年龄: %d\n", student.Age)
	fmt.Printf("学生ID: %s\n", student.StudentID)

	// 结构体切片
	fmt.Println("\n--- 结构体切片 ---")
	people := []Person{
		{Name: "Alice", Age: 25, Email: "alice@example.com"},
		{Name: "Bob", Age: 30, Email: "bob@example.com"},
		{Name: "Charlie", Age: 35, Email: "charlie@example.com"},
	}

	fmt.Println("人员列表:")
	for i, person := range people {
		fmt.Printf("  %d: %s (%d岁)\n", i+1, person.Name, person.Age)
	}

	// 结构体映射
	fmt.Println("\n--- 结构体映射 ---")
	employees := map[string]Person{
		"emp001": {Name: "Alice", Age: 25, Email: "alice@company.com"},
		"emp002": {Name: "Bob", Age: 30, Email: "bob@company.com"},
		"emp003": {Name: "Charlie", Age: 35, Email: "charlie@company.com"},
	}

	fmt.Println("员工信息:")
	for id, employee := range employees {
		fmt.Printf("  %s: %s (%d岁)\n", id, employee.Name, employee.Age)
	}

	// 结构体作为函数参数
	fmt.Println("\n--- 结构体作为函数参数 ---")
	p6 := Person{Name: "Grace", Age: 28, Email: "grace@example.com"}

	fmt.Printf("原始Person: %+v\n", p6)
	printPersonInfo(p6)

	// 值传递，不会修改原结构体
	updatePersonAge(p6, 29)
	fmt.Printf("值传递后Person: %+v\n", p6)

	// 指针传递，会修改原结构体
	updatePersonAgeByPointer(&p6, 29)
	fmt.Printf("指针传递后Person: %+v\n", p6)

	// 结构体标签示例
	fmt.Println("\n--- 结构体标签 ---")
	book := Book{
		Title:     "Go Programming",
		Author:    "John Doe",
		ISBN:      "978-0123456789",
		Price:     29.99,
		Published: true,
	}
	fmt.Printf("书籍信息: %+v\n", book)
	// 注意：标签主要用于序列化/反序列化，这里只是展示定义
}

// printPersonInfo 打印人员信息
func printPersonInfo(p Person) {
	fmt.Printf("姓名: %s, 年龄: %d, 邮箱: %s\n", p.Name, p.Age, p.Email)
}

// updatePersonAge 更新人员年龄（值传递）
func updatePersonAge(p Person, newAge int) {
	p.Age = newAge
	fmt.Printf("函数内修改后的年龄: %d\n", p.Age)
}

// updatePersonAgeByPointer 通过指针更新人员年龄
func updatePersonAgeByPointer(p *Person, newAge int) {
	p.Age = newAge
	fmt.Printf("通过指针修改后的年龄: %d\n", p.Age)
}
