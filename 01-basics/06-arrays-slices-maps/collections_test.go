package collections

import (
	"reflect"
	"sort"
	"testing"
)

func TestArrays(t *testing.T) {
	t.Run("ArrayDeclaration", func(t *testing.T) {
		// 测试数组声明和初始化
		var numbers [5]int
		scores := [3]int{95, 87, 92}
		fruits := [...]string{"苹果", "香蕉", "橙子"}

		t.Logf("零值数组: %v", numbers)
		t.Logf("分数数组: %v", scores)
		t.Logf("水果数组: %v", fruits)

		// 验证数组长度
		if len(numbers) != 5 {
			t.Errorf("numbers 长度应该是 5, 实际 %d", len(numbers))
		}
		if len(fruits) != 3 {
			t.Errorf("fruits 长度应该是 3, 实际 %d", len(fruits))
		}

		// 验证零值
		for i, v := range numbers {
			if v != 0 {
				t.Errorf("索引 %d 应该是 0, 实际 %d", i, v)
			}
		}
	})

	t.Run("ArrayAccess", func(t *testing.T) {
		numbers := [5]int{10, 20, 30, 40, 50}

		// 测试访问
		if numbers[0] != 10 {
			t.Errorf("第一个元素应该是 10, 实际 %d", numbers[0])
		}
		if numbers[4] != 50 {
			t.Errorf("最后一个元素应该是 50, 实际 %d", numbers[4])
		}

		// 测试修改
		numbers[2] = 100
		if numbers[2] != 100 {
			t.Errorf("修改后应该是 100, 实际 %d", numbers[2])
		}

		t.Logf("修改后的数组: %v", numbers)
	})

	t.Run("ArrayComparison", func(t *testing.T) {
		arr1 := [3]int{1, 2, 3}
		arr2 := [3]int{1, 2, 3}
		arr3 := [3]int{1, 2, 4}

		if arr1 != arr2 {
			t.Error("相同内容的数组应该相等")
		}
		if arr1 == arr3 {
			t.Error("不同内容的数组不应该相等")
		}

		t.Logf("arr1 == arr2: %t", arr1 == arr2)
		t.Logf("arr1 == arr3: %t", arr1 == arr3)
	})

	t.Run("ArrayIteration", func(t *testing.T) {
		scores := [4]int{95, 87, 92, 78}
		sum := 0

		// 使用 range 遍历
		for i, score := range scores {
			sum += score
			t.Logf("索引 %d: 分数 %d", i, score)
		}

		expectedSum := 352
		if sum != expectedSum {
			t.Errorf("总分应该是 %d, 实际 %d", expectedSum, sum)
		}
	})
}

func TestSlices(t *testing.T) {
	t.Run("SliceCreation", func(t *testing.T) {
		// 不同方式创建切片
		var slice1 []int
		slice2 := []int{1, 2, 3, 4, 5}
		slice3 := make([]int, 3)
		slice4 := make([]int, 3, 5)

		t.Logf("零值切片: %v (长度: %d, 容量: %d)", slice1, len(slice1), cap(slice1))
		t.Logf("字面量切片: %v (长度: %d, 容量: %d)", slice2, len(slice2), cap(slice2))
		t.Logf("make切片1: %v (长度: %d, 容量: %d)", slice3, len(slice3), cap(slice3))
		t.Logf("make切片2: %v (长度: %d, 容量: %d)", slice4, len(slice4), cap(slice4))

		// 验证属性
		if slice1 != nil {
			t.Error("零值切片应该是 nil")
		}
		if len(slice2) != 5 || cap(slice2) != 5 {
			t.Error("字面量切片长度和容量应该都是 5")
		}
		if len(slice4) != 3 || cap(slice4) != 5 {
			t.Error("make 切片长度应该是 3, 容量应该是 5")
		}
	})

	t.Run("SliceOperations", func(t *testing.T) {
		original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		// 切片操作
		sub1 := original[2:5]
		sub2 := original[:3]
		sub3 := original[7:]
		sub4 := original[:]

		t.Logf("original[2:5]: %v", sub1)
		t.Logf("original[:3]: %v", sub2)
		t.Logf("original[7:]: %v", sub3)
		t.Logf("original[:]: %v", sub4)

		// 验证切片内容
		expected1 := []int{2, 3, 4}
		expected2 := []int{0, 1, 2}
		expected3 := []int{7, 8, 9}

		if !reflect.DeepEqual(sub1, expected1) {
			t.Errorf("sub1 应该是 %v, 实际 %v", expected1, sub1)
		}
		if !reflect.DeepEqual(sub2, expected2) {
			t.Errorf("sub2 应该是 %v, 实际 %v", expected2, sub2)
		}
		if !reflect.DeepEqual(sub3, expected3) {
			t.Errorf("sub3 应该是 %v, 实际 %v", expected3, sub3)
		}
	})

	t.Run("SliceAppend", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		t.Logf("原始: %v", numbers)

		// 添加单个元素
		numbers = append(numbers, 4)
		t.Logf("添加一个元素: %v", numbers)

		// 添加多个元素
		numbers = append(numbers, 5, 6, 7)
		t.Logf("添加多个元素: %v", numbers)

		// 添加另一个切片
		more := []int{8, 9, 10}
		numbers = append(numbers, more...)
		t.Logf("添加另一个切片: %v", numbers)

		// 验证最终结果
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		if !reflect.DeepEqual(numbers, expected) {
			t.Errorf("最终结果应该是 %v, 实际 %v", expected, numbers)
		}
	})

	t.Run("SliceCopy", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		copied := make([]int, len(original))
		n := copy(copied, original)

		t.Logf("原始切片: %v", original)
		t.Logf("复制切片: %v", copied)
		t.Logf("复制了 %d 个元素", n)

		// 验证复制结果
		if !reflect.DeepEqual(original, copied) {
			t.Error("复制的切片应该与原始切片相同")
		}

		// 修改复制的切片不应影响原始切片
		copied[0] = 100
		if original[0] == 100 {
			t.Error("修改复制的切片不应影响原始切片")
		}

		t.Logf("修改后 - 原始: %v, 复制: %v", original, copied)
	})

	t.Run("SliceDelete", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		t.Logf("删除前: %v", slice)

		// 删除索引为 2 的元素
		index := 2
		slice = append(slice[:index], slice[index+1:]...)
		t.Logf("删除索引 %d 后: %v", index, slice)

		expected := []int{1, 2, 4, 5}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("删除后应该是 %v, 实际 %v", expected, slice)
		}
	})

	t.Run("SliceInsert", func(t *testing.T) {
		slice := []int{1, 2, 4, 5}
		t.Logf("插入前: %v", slice)

		// 在索引 2 处插入 3
		index := 2
		value := 3
		slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
		t.Logf("在索引 %d 插入 %d 后: %v", index, value, slice)

		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("插入后应该是 %v, 实际 %v", expected, slice)
		}
	})

	t.Run("SliceSort", func(t *testing.T) {
		numbers := []int{64, 34, 25, 12, 22, 11, 90}
		original := make([]int, len(numbers))
		copy(original, numbers)

		t.Logf("排序前: %v", numbers)
		sort.Ints(numbers)
		t.Logf("排序后: %v", numbers)

		// 验证排序结果
		for i := 1; i < len(numbers); i++ {
			if numbers[i] < numbers[i-1] {
				t.Error("排序后应该是升序")
				break
			}
		}

		// 字符串排序
		words := []string{"banana", "apple", "cherry", "date"}
		t.Logf("字符串排序前: %v", words)
		sort.Strings(words)
		t.Logf("字符串排序后: %v", words)

		expectedWords := []string{"apple", "banana", "cherry", "date"}
		if !reflect.DeepEqual(words, expectedWords) {
			t.Errorf("字符串排序结果应该是 %v, 实际 %v", expectedWords, words)
		}
	})
}

func TestMaps(t *testing.T) {
	t.Run("MapCreation", func(t *testing.T) {
		// 不同方式创建映射
		var map1 map[string]int
		map2 := make(map[string]int)
		map3 := map[string]int{
			"张三": 95,
			"李四": 87,
			"王五": 92,
		}

		t.Logf("零值映射: %v", map1)
		t.Logf("make映射: %v", map2)
		t.Logf("字面量映射: %v", map3)

		// 验证属性
		if map1 != nil {
			t.Error("零值映射应该是 nil")
		}
		if len(map2) != 0 {
			t.Error("新创建的映射长度应该是 0")
		}
		if len(map3) != 3 {
			t.Error("字面量映射长度应该是 3")
		}
	})

	t.Run("MapOperations", func(t *testing.T) {
		scores := make(map[string]int)

		// 添加元素
		scores["Go"] = 100
		scores["Python"] = 95
		scores["Java"] = 90

		t.Logf("添加元素后: %v", scores)

		// 访问元素
		goScore := scores["Go"]
		if goScore != 100 {
			t.Errorf("Go 的分数应该是 100, 实际 %d", goScore)
		}

		// 检查键是否存在
		if score, exists := scores["C++"]; exists {
			t.Errorf("C++ 不应该存在, 但得到分数 %d", score)
		} else {
			t.Log("C++ 不在映射中（正确）")
		}

		// 修改元素
		scores["Python"] = 98
		if scores["Python"] != 98 {
			t.Errorf("Python 分数应该被修改为 98, 实际 %d", scores["Python"])
		}

		// 删除元素
		delete(scores, "Java")
		if _, exists := scores["Java"]; exists {
			t.Error("Java 应该被删除")
		}

		t.Logf("删除 Java 后: %v", scores)
	})

	t.Run("MapIteration", func(t *testing.T) {
		scores := map[string]int{
			"张三": 95,
			"李四": 87,
			"王五": 92,
		}

		// 遍历映射
		var keys []string
		var values []int
		total := 0

		for name, score := range scores {
			keys = append(keys, name)
			values = append(values, score)
			total += score
			t.Logf("学生: %s, 分数: %d", name, score)
		}

		t.Logf("所有键: %v", keys)
		t.Logf("所有值: %v", values)
		t.Logf("总分: %d", total)

		// 验证结果
		if len(keys) != 3 || len(values) != 3 {
			t.Error("键和值的数量应该都是 3")
		}
		if total != 274 { // 95 + 87 + 92
			t.Errorf("总分应该是 274, 实际 %d", total)
		}
	})

	t.Run("MapZeroValue", func(t *testing.T) {
		// 测试映射的零值行为
		counts := make(map[string]int)

		// 访问不存在的键返回零值
		count := counts["不存在"]
		if count != 0 {
			t.Errorf("不存在的键应该返回零值 0, 实际 %d", count)
		}

		// 利用零值特性进行计数
		words := []string{"go", "is", "great", "go", "is", "simple"}
		for _, word := range words {
			counts[word]++
		}

		t.Logf("单词计数: %v", counts)

		expectedCounts := map[string]int{
			"go":     2,
			"is":     2,
			"great":  1,
			"simple": 1,
		}

		for word, expectedCount := range expectedCounts {
			if counts[word] != expectedCount {
				t.Errorf("单词 '%s' 计数应该是 %d, 实际 %d", word, expectedCount, counts[word])
			}
		}
	})
}

func TestComplexDataStructures(t *testing.T) {
	t.Run("SliceOfSlices", func(t *testing.T) {
		// 二维切片
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		t.Log("二维切片:")
		sum := 0
		for i, row := range matrix {
			rowSum := 0
			for _, val := range row {
				rowSum += val
				sum += val
			}
			t.Logf("  行 %d: %v (和: %d)", i, row, rowSum)
		}

		t.Logf("矩阵总和: %d", sum)

		if sum != 45 { // 1+2+...+9 = 45
			t.Errorf("矩阵总和应该是 45, 实际 %d", sum)
		}
	})

	t.Run("SliceOfMaps", func(t *testing.T) {
		// 映射的切片
		students := []map[string]interface{}{
			{"name": "张三", "age": 20, "score": 95},
			{"name": "李四", "age": 21, "score": 87},
			{"name": "王五", "age": 19, "score": 92},
		}

		t.Log("学生信息:")
		totalScore := 0
		for i, student := range students {
			score := student["score"].(int)
			totalScore += score
			t.Logf("  学生 %d: %v", i+1, student)
		}

		averageScore := float64(totalScore) / float64(len(students))
		t.Logf("平均分: %.2f", averageScore)

		expectedAverage := 274.0 / 3.0 // (95+87+92)/3
		if averageScore != expectedAverage {
			t.Errorf("平均分应该是 %.2f, 实际 %.2f", expectedAverage, averageScore)
		}
	})

	t.Run("MapOfSlices", func(t *testing.T) {
		// 切片的映射
		grades := map[string][]int{
			"数学": {95, 87, 92, 78, 85},
			"英语": {88, 92, 85, 90, 87},
			"物理": {90, 85, 88, 92, 89},
		}

		t.Log("各科成绩:")
		subjectAverages := make(map[string]float64)

		for subject, scores := range grades {
			total := 0
			for _, score := range scores {
				total += score
			}
			average := float64(total) / float64(len(scores))
			subjectAverages[subject] = average
			t.Logf("  %s: %v (平均: %.2f)", subject, scores, average)
		}

		// 验证数学平均分
		mathAverage := subjectAverages["数学"]
		expectedMathAverage := 437.0 / 5.0 // (95+87+92+78+85)/5
		if mathAverage != expectedMathAverage {
			t.Errorf("数学平均分应该是 %.2f, 实际 %.2f", expectedMathAverage, mathAverage)
		}
	})

	t.Run("MapOfMaps", func(t *testing.T) {
		// 映射的映射
		company := map[string]map[string]interface{}{
			"技术部": {
				"人数":  10,
				"预算":  1000000,
				"负责人": "张经理",
			},
			"销售部": {
				"人数":  15,
				"预算":  800000,
				"负责人": "李经理",
			},
		}

		t.Log("公司部门信息:")
		totalPeople := 0
		totalBudget := 0

		for dept, info := range company {
			people := info["人数"].(int)
			budget := info["预算"].(int)
			manager := info["负责人"].(string)

			totalPeople += people
			totalBudget += budget

			t.Logf("  %s:", dept)
			t.Logf("    人数: %d", people)
			t.Logf("    预算: %d", budget)
			t.Logf("    负责人: %s", manager)
		}

		t.Logf("总人数: %d", totalPeople)
		t.Logf("总预算: %d", totalBudget)

		if totalPeople != 25 {
			t.Errorf("总人数应该是 25, 实际 %d", totalPeople)
		}
		if totalBudget != 1800000 {
			t.Errorf("总预算应该是 1800000, 实际 %d", totalBudget)
		}
	})
}

func TestPracticalExamples(t *testing.T) {
	t.Run("WordCount", func(t *testing.T) {
		text := "go is great go is simple go is powerful"
		words := []string{"go", "is", "great", "go", "is", "simple", "go", "is", "powerful"}
		wordCount := make(map[string]int)

		for _, word := range words {
			wordCount[word]++
		}

		t.Logf("文本: %s", text)
		t.Log("单词计数:")
		for word, count := range wordCount {
			t.Logf("  %s: %d", word, count)
		}

		// 验证计数
		expectedCounts := map[string]int{
			"go":       3,
			"is":       3,
			"great":    1,
			"simple":   1,
			"powerful": 1,
		}

		for word, expectedCount := range expectedCounts {
			if wordCount[word] != expectedCount {
				t.Errorf("单词 '%s' 计数应该是 %d, 实际 %d", word, expectedCount, wordCount[word])
			}
		}
	})

	t.Run("DataDeduplication", func(t *testing.T) {
		numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
		seen := make(map[int]bool)
		var unique []int

		for _, num := range numbers {
			if !seen[num] {
				seen[num] = true
				unique = append(unique, num)
			}
		}

		t.Logf("原始数据: %v", numbers)
		t.Logf("去重后: %v", unique)

		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(unique, expected) {
			t.Errorf("去重结果应该是 %v, 实际 %v", expected, unique)
		}
	})

	t.Run("Grouping", func(t *testing.T) {
		people := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
		groups := make(map[int][]string)

		for i, person := range people {
			groupIndex := i % 3 // 分成3组
			groups[groupIndex] = append(groups[groupIndex], person)
		}

		t.Log("分组结果:")
		for groupIndex, members := range groups {
			t.Logf("  组 %d: %v", groupIndex, members)
		}

		// 验证分组
		if len(groups) != 3 {
			t.Errorf("应该有 3 个组, 实际 %d 个", len(groups))
		}

		totalMembers := 0
		for _, members := range groups {
			totalMembers += len(members)
		}

		if totalMembers != len(people) {
			t.Errorf("总成员数应该是 %d, 实际 %d", len(people), totalMembers)
		}
	})
}

// 基准测试
func BenchmarkSliceAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var slice []int
		for j := 0; j < 1000; j++ {
			slice = append(slice, j)
		}
	}
}

func BenchmarkMapAccess(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[i%1000]
	}
}
