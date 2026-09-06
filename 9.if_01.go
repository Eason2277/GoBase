package main

import "fmt"

func main() {
	var stock int
	fmt.Println("请输入库存数量")
	fmt.Scanln(&stock)

	//if stock > 0 && stock <= 4 {
	//	fmt.Println("可以下单，但库存比较少了！")
	//	return
	//}
	//
	//if stock == 0 {
	//	fmt.Println("无货物")
	//	return
	//}
	//
	//if stock >= 5 {
	//	fmt.Println("挺多的")
	//	return
	//}
	//
	//if stock < 0 {
	//	fmt.Println("货物数不能为负数")
	//	return
	//}

	if stock > 0 && stock < 1000 {
		fmt.Println("可以下单，但库存比较少了！")
	} else if stock >= 1000 {
		fmt.Println("可以下单，并且库存非常充足")
	} else if stock < 0 {
		fmt.Println("货物数不能为负数")

	}
}
