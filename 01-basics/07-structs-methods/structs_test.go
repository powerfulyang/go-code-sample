package structs

import (
	"math"
	"testing"
	"time"
)

func TestBasicStructs(t *testing.T) {
	t.Run("StructCreation", func(t *testing.T) {
		// 不同方式创建结构体
		person1 := Person{
			Name: "张三",
			Age:  25,
			City: "北京",
		}

		person2 := Person{"李四", 30, "上海"}

		var person3 Person
		person3.Name = "王五"
		person3.Age = 28
		person3.City = "广州"

		t.Logf("person1: %+v", person1)
		t.Logf("person2: %+v", person2)
		t.Logf("person3: %+v", person3)

		// 验证字段值
		if person1.Name != "张三" || person1.Age != 25 {
			t.Error("person1 字段值不正确")
		}
		if person2.Name != "李四" || person2.Age != 30 {
			t.Error("person2 字段值不正确")
		}
		if person3.Name != "王五" || person3.Age != 28 {
			t.Error("person3 字段值不正确")
		}
	})

	t.Run("ConstructorFunctions", func(t *testing.T) {
		// 使用构造函数
		person := NewPerson("赵六", 35, "深圳")

		t.Logf("构造函数创建的person: %+v", *person)

		if person.Name != "赵六" || person.Age != 35 || person.City != "深圳" {
			t.Error("构造函数创建的结构体字段值不正确")
		}
	})

	t.Run("ZeroValues", func(t *testing.T) {
		// 零值结构体
		var person Person

		t.Logf("零值结构体: %+v", person)

		if person.Name != "" || person.Age != 0 || person.City != "" {
			t.Error("零值结构体的字段应该是对应类型的零值")
		}
	})
}

func TestNestedStructs(t *testing.T) {
	t.Run("EmbeddedStructs", func(t *testing.T) {
		// 嵌套结构体
		employee := Employee{
			Person: Person{
				Name: "钱七",
				Age:  32,
				City: "杭州",
			},
			ID:     1001,
			Salary: 8500.0,
			Address: Address{
				Street:   "西湖路123号",
				City:     "杭州",
				Province: "浙江",
				ZipCode:  "310000",
			},
		}

		t.Logf("员工信息: %+v", employee)

		// 测试嵌入字段的访问
		if employee.Name != "钱七" {
			t.Error("应该能直接访问嵌入结构体的字段")
		}
		if employee.Person.Name != "钱七" {
			t.Error("应该能通过嵌入字段名访问")
		}

		// 测试命名字段的访问
		if employee.Address.Street != "西湖路123号" {
			t.Error("命名字段访问失败")
		}

		t.Logf("员工姓名: %s", employee.Name)
		t.Logf("员工地址: %s, %s", employee.Address.Street, employee.Address.City)
	})

	t.Run("StructTags", func(t *testing.T) {
		// 带标签的结构体
		user := User{
			ID:       1,
			Username: "testuser",
			Email:    "test@example.com",
			IsActive: true,
		}

		t.Logf("用户信息: %+v", user)

		// 验证字段值
		if user.ID != 1 || user.Username != "testuser" {
			t.Error("用户字段值不正确")
		}
		if !user.IsActive {
			t.Error("用户应该是活跃状态")
		}
	})
}

func TestMethods(t *testing.T) {
	t.Run("ValueReceiverMethods", func(t *testing.T) {
		// 值接收者方法
		rect := Rectangle{Width: 5.0, Height: 3.0}

		area := rect.Area()
		perimeter := rect.Perimeter()

		t.Logf("矩形: 宽=%.1f, 高=%.1f", rect.Width, rect.Height)
		t.Logf("面积: %.2f", area)
		t.Logf("周长: %.2f", perimeter)

		// 验证计算结果
		expectedArea := 15.0
		expectedPerimeter := 16.0

		if area != expectedArea {
			t.Errorf("面积计算错误: 期望 %.2f, 实际 %.2f", expectedArea, area)
		}
		if perimeter != expectedPerimeter {
			t.Errorf("周长计算错误: 期望 %.2f, 实际 %.2f", expectedPerimeter, perimeter)
		}
	})

	t.Run("PointerReceiverMethods", func(t *testing.T) {
		// 指针接收者方法
		rect := NewRectangle(4.0, 6.0)
		originalArea := rect.Area()

		t.Logf("原始矩形: 宽=%.1f, 高=%.1f, 面积=%.2f", rect.Width, rect.Height, originalArea)

		// 缩放
		rect.Scale(2.0)
		newArea := rect.Area()

		t.Logf("缩放后矩形: 宽=%.1f, 高=%.1f, 面积=%.2f", rect.Width, rect.Height, newArea)

		// 验证缩放结果
		if rect.Width != 8.0 || rect.Height != 12.0 {
			t.Error("缩放后尺寸不正确")
		}
		if newArea != originalArea*4 { // 面积应该是原来的4倍
			t.Errorf("缩放后面积不正确: 期望 %.2f, 实际 %.2f", originalArea*4, newArea)
		}

		// 设置新尺寸
		rect.SetDimensions(10.0, 8.0)
		t.Logf("设置新尺寸后: 宽=%.1f, 高=%.1f", rect.Width, rect.Height)

		if rect.Width != 10.0 || rect.Height != 8.0 {
			t.Error("设置新尺寸失败")
		}
	})

	t.Run("CircleMethods", func(t *testing.T) {
		// 圆形方法测试
		circle := Circle{Radius: 5.0}

		area := circle.Area()
		circumference := circle.Circumference()
		diameter := circle.Diameter()

		t.Logf("圆形: 半径=%.1f", circle.Radius)
		t.Logf("面积: %.2f", area)
		t.Logf("周长: %.2f", circumference)
		t.Logf("直径: %.2f", diameter)

		// 验证计算结果
		expectedArea := math.Pi * 25             // π * r²
		expectedCircumference := 2 * math.Pi * 5 // 2πr
		expectedDiameter := 10.0

		if math.Abs(area-expectedArea) > 0.01 {
			t.Errorf("圆形面积计算错误: 期望 %.2f, 实际 %.2f", expectedArea, area)
		}
		if math.Abs(circumference-expectedCircumference) > 0.01 {
			t.Errorf("圆形周长计算错误: 期望 %.2f, 实际 %.2f", expectedCircumference, circumference)
		}
		if diameter != expectedDiameter {
			t.Errorf("圆形直径计算错误: 期望 %.2f, 实际 %.2f", expectedDiameter, diameter)
		}
	})
}

func TestBankAccount(t *testing.T) {
	t.Run("AccountOperations", func(t *testing.T) {
		// 银行账户操作测试
		account := NewBankAccount("123456789", "张三", 1000.0)

		t.Logf("初始状态: %s", account.GetAccountInfo())

		// 测试初始余额
		if account.GetBalance() != 1000.0 {
			t.Errorf("初始余额错误: 期望 1000.0, 实际 %.2f", account.GetBalance())
		}

		// 测试存款
		err := account.Deposit(500.0)
		if err != nil {
			t.Errorf("存款失败: %v", err)
		}

		t.Logf("存款后: %s", account.GetAccountInfo())

		if account.GetBalance() != 1500.0 {
			t.Errorf("存款后余额错误: 期望 1500.0, 实际 %.2f", account.GetBalance())
		}

		// 测试取款
		err = account.Withdraw(200.0)
		if err != nil {
			t.Errorf("取款失败: %v", err)
		}

		t.Logf("取款后: %s", account.GetAccountInfo())

		if account.GetBalance() != 1300.0 {
			t.Errorf("取款后余额错误: 期望 1300.0, 实际 %.2f", account.GetBalance())
		}
	})

	t.Run("AccountErrorHandling", func(t *testing.T) {
		account := NewBankAccount("987654321", "李四", 100.0)

		// 测试无效存款
		err := account.Deposit(-50.0)
		if err == nil {
			t.Error("负数存款应该返回错误")
		}
		t.Logf("负数存款错误: %v", err)

		err = account.Deposit(0)
		if err == nil {
			t.Error("零存款应该返回错误")
		}

		// 测试无效取款
		err = account.Withdraw(-30.0)
		if err == nil {
			t.Error("负数取款应该返回错误")
		}
		t.Logf("负数取款错误: %v", err)

		// 测试余额不足
		err = account.Withdraw(200.0)
		if err == nil {
			t.Error("余额不足应该返回错误")
		}
		t.Logf("余额不足错误: %v", err)

		// 余额应该没有变化
		if account.GetBalance() != 100.0 {
			t.Errorf("错误操作后余额应该保持不变: 期望 100.0, 实际 %.2f", account.GetBalance())
		}
	})
}

func TestStudent(t *testing.T) {
	t.Run("StudentScoreManagement", func(t *testing.T) {
		student := NewStudent(1001, "李明")

		t.Logf("学生: %s (ID: %d)", student.Name, student.ID)

		// 测试空分数
		if student.Average() != 0 {
			t.Error("空分数列表的平均分应该是0")
		}
		if student.HighestScore() != 0 {
			t.Error("空分数列表的最高分应该是0")
		}
		if student.LowestScore() != 0 {
			t.Error("空分数列表的最低分应该是0")
		}

		// 添加分数
		scores := []int{85, 92, 78, 95, 88}
		for _, score := range scores {
			student.AddScore(score)
		}

		t.Logf("分数: %v", student.Scores)

		// 验证分数数量
		if len(student.Scores) != 5 {
			t.Errorf("分数数量错误: 期望 5, 实际 %d", len(student.Scores))
		}

		// 验证平均分
		expectedAverage := 87.6 // (85+92+78+95+88)/5
		average := student.Average()
		if math.Abs(average-expectedAverage) > 0.1 {
			t.Errorf("平均分计算错误: 期望 %.2f, 实际 %.2f", expectedAverage, average)
		}

		// 验证最高分和最低分
		highest := student.HighestScore()
		lowest := student.LowestScore()

		if highest != 95 {
			t.Errorf("最高分错误: 期望 95, 实际 %d", highest)
		}
		if lowest != 78 {
			t.Errorf("最低分错误: 期望 78, 实际 %d", lowest)
		}

		t.Logf("平均分: %.2f", average)
		t.Logf("最高分: %d", highest)
		t.Logf("最低分: %d", lowest)

		// 添加更多分数
		student.AddScore(90)
		student.AddScore(87)

		t.Logf("更新后分数: %v", student.Scores)
		t.Logf("更新后平均分: %.2f", student.Average())

		if len(student.Scores) != 7 {
			t.Errorf("更新后分数数量错误: 期望 7, 实际 %d", len(student.Scores))
		}
	})
}

func TestEvent(t *testing.T) {
	t.Run("EventTiming", func(t *testing.T) {
		// 创建一个过去的事件用于测试
		startTime := time.Now().Add(-time.Hour) // 1小时前开始
		event := Event{
			Name:      "Go语言培训",
			StartTime: startTime,
			Duration:  2 * time.Hour, // 持续2小时
		}

		t.Logf("事件: %s", event.Name)
		t.Logf("开始时间: %s", event.StartTime.Format("2006-01-02 15:04:05"))
		t.Logf("结束时间: %s", event.EndTime().Format("2006-01-02 15:04:05"))
		t.Logf("持续时间: %v", event.Duration)

		// 验证结束时间
		expectedEndTime := startTime.Add(2 * time.Hour)
		if !event.EndTime().Equal(expectedEndTime) {
			t.Error("结束时间计算错误")
		}

		// 测试当前状态
		now := time.Now()
		isOngoing := event.IsOngoing(now)
		timeRemaining := event.TimeRemaining(now)

		t.Logf("当前时间: %s", now.Format("2006-01-02 15:04:05"))
		t.Logf("是否正在进行: %t", isOngoing)
		t.Logf("剩余时间: %v", timeRemaining)

		// 由于事件1小时前开始，持续2小时，所以现在应该正在进行
		if !isOngoing {
			t.Error("事件应该正在进行")
		}
		if timeRemaining <= 0 {
			t.Error("应该还有剩余时间")
		}
	})

	t.Run("EventStates", func(t *testing.T) {
		now := time.Now()

		// 未来事件
		futureEvent := Event{
			Name:      "未来事件",
			StartTime: now.Add(time.Hour),
			Duration:  time.Hour,
		}

		// 过去事件
		pastEvent := Event{
			Name:      "过去事件",
			StartTime: now.Add(-2 * time.Hour),
			Duration:  time.Hour,
		}

		t.Logf("未来事件正在进行: %t", futureEvent.IsOngoing(now))
		t.Logf("过去事件正在进行: %t", pastEvent.IsOngoing(now))
		t.Logf("未来事件剩余时间: %v", futureEvent.TimeRemaining(now))
		t.Logf("过去事件剩余时间: %v", pastEvent.TimeRemaining(now))

		// 验证状态
		if futureEvent.IsOngoing(now) {
			t.Error("未来事件不应该正在进行")
		}
		if pastEvent.IsOngoing(now) {
			t.Error("过去事件不应该正在进行")
		}
		if pastEvent.TimeRemaining(now) != 0 {
			t.Error("过去事件剩余时间应该是0")
		}
	})
}

func TestStructComparison(t *testing.T) {
	t.Run("ComparableStructs", func(t *testing.T) {
		person1 := Person{Name: "张三", Age: 25, City: "北京"}
		person2 := Person{Name: "张三", Age: 25, City: "北京"}
		person3 := Person{Name: "李四", Age: 30, City: "上海"}

		t.Logf("person1: %+v", person1)
		t.Logf("person2: %+v", person2)
		t.Logf("person3: %+v", person3)

		// 测试结构体比较
		if person1 != person2 {
			t.Error("相同内容的结构体应该相等")
		}
		if person1 == person3 {
			t.Error("不同内容的结构体不应该相等")
		}

		t.Logf("person1 == person2: %t", person1 == person2)
		t.Logf("person1 == person3: %t", person1 == person3)
	})

	t.Run("NonComparableStructs", func(t *testing.T) {
		// 包含切片的结构体不能直接比较
		student1 := Student{ID: 1, Name: "张三", Scores: []int{85, 90}}
		student2 := Student{ID: 1, Name: "张三", Scores: []int{85, 90}}

		// 手动比较各个字段
		fieldsEqual := student1.ID == student2.ID &&
			student1.Name == student2.Name &&
			len(student1.Scores) == len(student2.Scores)

		if fieldsEqual {
			// 比较切片内容
			scoresEqual := true
			for i, score := range student1.Scores {
				if score != student2.Scores[i] {
					scoresEqual = false
					break
				}
			}
			fieldsEqual = scoresEqual
		}

		t.Logf("学生结构体字段相等: %t", fieldsEqual)

		if !fieldsEqual {
			t.Error("相同内容的学生结构体应该在字段级别相等")
		}
	})
}

// 基准测试
func BenchmarkRectangleArea(b *testing.B) {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	for i := 0; i < b.N; i++ {
		_ = rect.Area()
	}
}

func BenchmarkStudentAverage(b *testing.B) {
	student := NewStudent(1, "测试学生")
	for i := 0; i < 100; i++ {
		student.AddScore(85 + i%15)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = student.Average()
	}
}
