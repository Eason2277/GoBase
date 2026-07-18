package main

import "fmt"

var age = 12

func main() {
	// 1. Full declaration: var name type = initial value.
	var name string = "小明"
	var age int = 18
	fmt.Println("姓名：", name)
	fmt.Println("年龄：", age)

	// 2. Type inference: Go infers the type from the initial value.
	var city = "乌鲁木齐"
	fmt.Printf("城市：%s，类型：%T\n", city, city)

	// 3. Short declaration: := declares and initializes a variable inside a function.
	job := "前端工程师"
	fmt.Printf("职业：%s，类型：%T\n", job, job)

	// 4. A variable can be reassigned, but its type cannot change.
	age = 19
	fmt.Println("修改后的年龄：", age)

	// 5. A variable without an initial value receives its type's zero value.
	var score int
	var message string
	var passed bool
	fmt.Printf("零值：score=%d, message=%q, passed=%t\n", score, message, passed)

	// 6. Multiple variables can be declared together.
	var width, height int = 1920, 1080
	fmt.Printf("分辨率：%d x %d\n", width, height)
}
