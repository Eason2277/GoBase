package main

import (
	"errors"
	"fmt"
)

// 业务场景：订单结算
// 商城满99就可以免去运费，不满的话，加8元运费，现在又两个订单要算实付金额

// 判断的逻辑单独抽取出来
func finalPrice(price int, quantity int) int {
	total := price * quantity

	if total < 99 {
		total += 8
	}
	return total
}

// 业务升级了，之前上写死价格
func getPrice(name string) (int, error) {
	if name == "苹果" {
		return 5, nil
	}

	if name == "牛奶" {
		return 12, nil
	}

	return 0, errors.New("商品不存在")
}

func main() {
	// 我们以订单作为一个计算模块
	// 订单1
	//price1 := 55
	//quantity1 := 2
	//
	//fmt.Println("订单1实付金额：", finalPrice(price1, quantity1))

	price1, err := getPrice("牛奶")

	if err != nil {
		fmt.Println("获取价格失败：", err)
		return
	}

	fmt.Println("订单1实付金额：", finalPrice(price1, 5))

	// 订单2
	price2 := 66
	quantity2 := 1

	fmt.Println("订单2实付金额：", finalPrice(price2, quantity2))

}
