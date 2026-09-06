package main

import "fmt"

func main() {
	var userMap map[int]string = map[int]string{
		1: "哈哈",
		3: "嘿嘿",
		4: "",
	}

	fmt.Println(userMap)
	fmt.Println(userMap[3])
	fmt.Printf("%#v\n", userMap[4])

	// 如果我不知道值本身说空字符串还是没拿到咋办
	// 我们可以自行写一个判断

	value, okk := userMap[3]

	if okk {
		fmt.Printf("键存在，值是：%#v\n", value, okk)
	} else {
		fmt.Println("键不存在")
	}

	// to replace a value
	userMap[1] = "哈哈哈"
	fmt.Println(userMap)

	delete(userMap, 3)
	fmt.Println(userMap)

	// 两种初始化map的方式
	var map1 = make(map[int]string)
	map1[3] = "篮球"

	var map2 = map[string]string{}
	map2["西班牙"] = "冠军"

}
