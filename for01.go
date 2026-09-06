package main

import "fmt"

func main() {
	// 业务场景：算购物车里 5 件商品的总价（简化：每件价格就是它的编号 x 10 元
	//total := 0
	//for i := 1; i <= 5; i++ {
	//	total += i * 10
	//}
	//
	//fmt.Println(total)

	// 业务场景：秒杀扣库存。初始库存是57件，每件订单扣3件，一直扣到不够扣为止，看看能完成几单
	stock := 58
	order := 0

	for stock >= 3 {
		// 就是拿订单数作为牵引头子，术语即循环控制变量
		stock = stock - 3
		order++
		fmt.Printf("完成了第 %d 单，剩余库存 %d\n", order, stock)
	}

}
