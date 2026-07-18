package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "hi")

	// %T means the type
	fmt.Printf("%v %T\n", "hi", "你好")

	// %d means decimal
	fmt.Printf("%d\n", 6)

	// %.2f formats a floating-point number to two decimal places
	fmt.Printf("%.2f\n", 3.8888888)

	var number = fmt.Sprintf("%.2f", 3.1415928873232)
	fmt.Printf(number)

	fmt.Print("\n请输入你的名字")
	var name string
	fmt.Scan(&name)
	fmt.Printf("%v\n\n", name)

	fmt.Printf("请输入你的年龄")
	var age int
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)

}
