package main

import "fmt"

/*
- ID：商品编号，用 int
- Name：商品名称，用 string
- PriceFen：价格，单位是分，用 int
*/
type Product struct {
	ID       int
	Name     string
	PriceFen int
}

func getProduct(products []Product, index int) (Product, bool) {
	if index < 0 || index >= len(products) {
		return Product{}, false
	}

	return products[index], true
}

func insertProduct(products []Product, index int, newProduct Product) ([]Product, bool) {
	if index < 0 || index > len(products) {
		return products, false
	}

	// 扩展slice的长度
	products = append(products, Product{})

	// copy(目标位置, 来源位置)
	// 把插入位置以及后面的元素，整体向右移动一个位置，为新元素腾出空间
	copy(products[index+1:], products[index:])

	products[index] = newProduct
	return products, true

}

func main() {

	product := Product{
		ID:       1001,
		Name:     "小岛经济学",
		PriceFen: 200,
	}

	secondProduct := Product{
		ID:       1002,
		Name:     "麦哲伦与大航海时代",
		PriceFen: 220,
	}

	products := []Product{product, secondProduct}

	// 做插入这个动作
	newProducts := Product{
		ID:       1003,
		Name:     "算法",
		PriceFen: 300,
	}

	products, insertOK := insertProduct(products, 2, newProducts)

	if insertOK {
		fmt.Println("插入成功", products)
	} else {
		fmt.Println("插入失败")
	}

	fmt.Println(products)

	fmt.Println(products[1])
	fmt.Println(products[1].Name)

	result, ok := getProduct(products, 2)

	if ok {
		fmt.Println("找到商品：", result.Name)
	} else {
		fmt.Println("商品不存在")
	}

}
