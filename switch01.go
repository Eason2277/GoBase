package main

import "fmt"

func main() {
	var status int
	fmt.Println("请输入状态码")
	fmt.Scanln(&status)

	switch status {
	case 1:
		fmt.Println("有待支付")

	case 2:
		fmt.Println("待发货")

	default:
		fmt.Println("未知状态")
	}
}
