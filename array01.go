package main

import "fmt"

func main() {
	// 1. An array has a fixed length, and uninitialized elements use zero values.
	var scores [3]int
	scores[0] = 90
	scores[1] = 85
	scores[2] = 95
	fmt.Println("数组：", scores)
	fmt.Println("数组长度：", len(scores))
	fmt.Println("第一个元素：", scores[0])

	var players = [5]string{"科比", "kd", "梅西", "罗德里", "德布劳内"}
	fmt.Println(players)

	// 2. Use [...] to let Go count the array length.
	//cities := [...]string{"乌鲁木齐", "西安", "上海"}
	cities := [...]string{"北京", "上海", "广州"}
	fmt.Println("城市数组：", cities)

	// 3. range returns the index and value of each element.
	for index, city := range cities {
		fmt.Printf("cities[%d] = %s\n", index, city)
	}

	// 4. A slice has no fixed length in its type and can grow with append.
	fruits := []string{"苹果", "香蕉"}
	fruits = append(fruits, "葡萄")
	fmt.Println("切片：", fruits)
	fmt.Printf("切片长度：%d，容量：%d\n", len(fruits), cap(fruits))

	// 5. A slice can be created from part of an array: [start:end).
	numbers := [5]int{10, 20, 30, 40, 50}
	part := numbers[3:5]
	fmt.Println("原数组：", numbers)
	fmt.Println("切片 numbers[3:5]：", part)

	// The slice and array share underlying data here.
	part[1] = 200
	fmt.Println("修改切片后，原数组：", numbers)

	// 切片其实还是开拓了内存空间，只是，切片这个变量本身，存的是对自己这个切片的len,cap
	// 以及对应原数组的存储指针
	// 我们可以通过使用make方法单独创建一个切片，go就会为这个切片准备一个新的对应的底层数组
	// 但这个由make创建的变量本质仍然是切片

	// 我们通常只能通过 slice 操作底层数据。因此，与其说底层数组是“主驾驶”，更准确地说，它是“真正存放数据的仓库”，切片是访问仓库某一部分的窗口。

}
