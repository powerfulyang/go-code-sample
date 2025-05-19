package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Person 人员结构体
type Person struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Address  Address   `json:"address"`
	Hobbies  []string  `json:"hobbies"`
	Active   bool      `json:"active"`
}

// Address 地址结构体
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}

// Product 产品结构体（演示不同的JSON标签）
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"` // 空值时省略
	InStock     bool    `json:"in_stock"`
	Category    string  `json:"-"` // 忽略此字段
	Internal    string  `json:"internal,omitempty"`
}

func main() {
	fmt.Println("=== Go JSON 处理示例 ===")

	// 基本序列化
	fmt.Println("\n--- 基本序列化 (Marshal) ---")
	basicMarshalDemo()

	// 基本反序列化
	fmt.Println("\n--- 基本反序列化 (Unmarshal) ---")
	basicUnmarshalDemo()

	// 结构体序列化
	fmt.Println("\n--- 结构体序列化 ---")
	structMarshalDemo()

	// 结构体反序列化
	fmt.Println("\n--- 结构体反序列化 ---")
	structUnmarshalDemo()

	// JSON标签示例
	fmt.Println("\n--- JSON 标签示例 ---")
	jsonTagsDemo()

	// 嵌套结构体
	fmt.Println("\n--- 嵌套结构体 ---")
	nestedStructDemo()

	// 处理未知结构的JSON
	fmt.Println("\n--- 处理未知结构的JSON ---")
	unknownStructDemo()

	// 自定义JSON序列化
	fmt.Println("\n--- 自定义JSON序列化 ---")
	customMarshalDemo()
}

// 基本序列化示例
func basicMarshalDemo() {
	// 基本类型序列化
	name := "Alice"
	age := 25
	active := true
	scores := []int{85, 92, 78, 96}

	nameJSON, _ := json.Marshal(name)
	ageJSON, _ := json.Marshal(age)
	activeJSON, _ := json.Marshal(active)
	scoresJSON, _ := json.Marshal(scores)

	fmt.Printf("字符串: %s -> %s\n", name, nameJSON)
	fmt.Printf("整数: %d -> %s\n", age, ageJSON)
	fmt.Printf("布尔值: %t -> %s\n", active, activeJSON)
	fmt.Printf("切片: %v -> %s\n", scores, scoresJSON)

	// Map序列化
	data := map[string]interface{}{
		"name":   "Bob",
		"age":    30,
		"active": true,
		"scores": []int{88, 91, 85},
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("序列化错误: %v", err)
		return
	}
	fmt.Printf("Map: %v -> %s\n", data, dataJSON)
}

// 基本反序列化示例
func basicUnmarshalDemo() {
	// 基本类型反序列化
	nameJSON := `"Charlie"`
	ageJSON := `35`
	activeJSON := `false`
	scoresJSON := `[90, 85, 88, 92]`

	var name string
	var age int
	var active bool
	var scores []int

	json.Unmarshal([]byte(nameJSON), &name)
	json.Unmarshal([]byte(ageJSON), &age)
	json.Unmarshal([]byte(activeJSON), &active)
	json.Unmarshal([]byte(scoresJSON), &scores)

	fmt.Printf("字符串: %s -> %s\n", nameJSON, name)
	fmt.Printf("整数: %s -> %d\n", ageJSON, age)
	fmt.Printf("布尔值: %s -> %t\n", activeJSON, active)
	fmt.Printf("切片: %s -> %v\n", scoresJSON, scores)

	// Map反序列化
	dataJSON := `{"name":"David","age":28,"active":true,"scores":[87,93,89]}`
	var data map[string]interface{}

	err := json.Unmarshal([]byte(dataJSON), &data)
	if err != nil {
		log.Printf("反序列化错误: %v", err)
		return
	}
	fmt.Printf("Map: %s -> %v\n", dataJSON, data)
}

// 结构体序列化示例
func structMarshalDemo() {
	person := Person{
		Name:     "Alice Johnson",
		Age:      28,
		Email:    "alice@example.com",
		Birthday: time.Date(1995, 5, 15, 0, 0, 0, 0, time.UTC),
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
			ZipCode: "10001",
		},
		Hobbies: []string{"reading", "swimming", "coding"},
		Active:  true,
	}

	// 紧凑格式
	compactJSON, err := json.Marshal(person)
	if err != nil {
		log.Printf("序列化错误: %v", err)
		return
	}
	fmt.Printf("紧凑JSON: %s\n", compactJSON)

	// 格式化输出
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Printf("序列化错误: %v", err)
		return
	}
	fmt.Printf("格式化JSON:\n%s\n", prettyJSON)
}

// 结构体反序列化示例
func structUnmarshalDemo() {
	jsonData := `{
		"name": "Bob Smith",
		"age": 32,
		"email": "bob@example.com",
		"birthday": "1991-08-20T00:00:00Z",
		"address": {
			"street": "456 Oak Ave",
			"city": "Los Angeles",
			"country": "USA",
			"zip_code": "90210"
		},
		"hobbies": ["gaming", "cooking", "traveling"],
		"active": true
	}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Printf("反序列化错误: %v", err)
		return
	}

	fmt.Printf("反序列化结果:\n")
	fmt.Printf("姓名: %s\n", person.Name)
	fmt.Printf("年龄: %d\n", person.Age)
	fmt.Printf("邮箱: %s\n", person.Email)
	fmt.Printf("生日: %s\n", person.Birthday.Format("2006-01-02"))
	fmt.Printf("地址: %s, %s, %s\n", person.Address.Street, person.Address.City, person.Address.Country)
	fmt.Printf("爱好: %v\n", person.Hobbies)
	fmt.Printf("活跃: %t\n", person.Active)
}

// JSON标签示例
func jsonTagsDemo() {
	product := Product{
		ID:          1,
		Name:        "Laptop",
		Price:       999.99,
		Description: "", // 空值，会被省略
		InStock:     true,
		Category:    "Electronics", // 会被忽略
		Internal:    "",            // 空值，会被省略
	}

	productJSON, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Printf("序列化错误: %v", err)
		return
	}
	fmt.Printf("产品JSON (注意省略和忽略的字段):\n%s\n", productJSON)

	// 带描述的产品
	product2 := Product{
		ID:          2,
		Name:        "Mouse",
		Price:       29.99,
		Description: "Wireless optical mouse",
		InStock:     false,
		Category:    "Accessories",
		Internal:    "internal data",
	}

	product2JSON, err := json.MarshalIndent(product2, "", "  ")
	if err != nil {
		log.Printf("序列化错误: %v", err)
		return
	}
	fmt.Printf("产品2 JSON:\n%s\n", product2JSON)
}

// 嵌套结构体示例
func nestedStructDemo() {
	// 复杂嵌套结构
	data := map[string]interface{}{
		"users": []Person{
			{
				Name:  "User1",
				Age:   25,
				Email: "user1@example.com",
				Address: Address{
					Street:  "Street 1",
					City:    "City 1",
					Country: "Country 1",
				},
				Hobbies: []string{"hobby1", "hobby2"},
				Active:  true,
			},
			{
				Name:  "User2",
				Age:   30,
				Email: "user2@example.com",
				Address: Address{
					Street:  "Street 2",
					City:    "City 2",
					Country: "Country 2",
				},
				Hobbies: []string{"hobby3", "hobby4"},
				Active:  false,
			},
		},
		"total": 2,
		"page":  1,
	}

	nestedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("序列化错误: %v", err)
		return
	}
	fmt.Printf("嵌套结构JSON:\n%s\n", nestedJSON)
}

// 处理未知结构的JSON
func unknownStructDemo() {
	jsonData := `{
		"name": "Unknown Structure",
		"data": {
			"numbers": [1, 2, 3, 4, 5],
			"nested": {
				"key1": "value1",
				"key2": 42,
				"key3": true
			}
		},
		"timestamp": "2023-01-01T00:00:00Z"
	}`

	// 使用interface{}处理未知结构
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		log.Printf("反序列化错误: %v", err)
		return
	}

	fmt.Printf("未知结构解析结果:\n")
	for key, value := range result {
		fmt.Printf("  %s: %v (类型: %T)\n", key, value, value)
	}

	// 访问嵌套数据
	if data, ok := result["data"].(map[string]interface{}); ok {
		if nested, ok := data["nested"].(map[string]interface{}); ok {
			fmt.Printf("嵌套数据:\n")
			for k, v := range nested {
				fmt.Printf("  %s: %v\n", k, v)
			}
		}
	}
}

// 自定义时间类型，演示自定义序列化
type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time.Format("2006-01-02 15:04:05"))
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}

// 自定义JSON序列化示例
func customMarshalDemo() {
	type Event struct {
		Name string     `json:"name"`
		Time CustomTime `json:"time"`
	}

	event := Event{
		Name: "Meeting",
		Time: CustomTime{time.Now()},
	}

	eventJSON, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		log.Printf("序列化错误: %v", err)
		return
	}
	fmt.Printf("自定义时间格式:\n%s\n", eventJSON)

	// 反序列化
	var newEvent Event
	err = json.Unmarshal(eventJSON, &newEvent)
	if err != nil {
		log.Printf("反序列化错误: %v", err)
		return
	}
	fmt.Printf("反序列化后的时间: %s\n", newEvent.Time.Format("2006-01-02 15:04:05"))
}
