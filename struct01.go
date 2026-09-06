package main

import "fmt"

type Category struct {
	ID   int
	Name string
}

type Product struct {
	ID       int
	Name     string
	PriceFen int
	Stock    int
	Category Category
}

func main() {
	//productName := "Go 语言入门"
	//productPrice := 30

	product := Product{
		ID:       1001,
		Name:     "go语言入门",
		PriceFen: 3000,
		Stock:    10,
		Category: Category{
			ID:   10,
			Name: "编程图书",
		},
	}

	//fmt.Println("商品：", productName, "价格：", productPrice)
	fmt.Println(product)
	fmt.Println("商品名称：", product.Name)
	fmt.Println("商品库存：", product.Stock)
	fmt.Println("商品分类：", product.Category.Name)

}
