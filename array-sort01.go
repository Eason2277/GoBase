package main

import (
	"fmt"
	"slices"
)

func main() {
	// 1. Sort a slice in ascending order.
	numbers := []int{40, 10, 30, 20}
	fmt.Println("Slice before sorting:", numbers)

	slices.Sort(numbers)
	fmt.Println("Slice in ascending order:", numbers)

	slices.Reverse(numbers)
	fmt.Println("Slice in descending order:", numbers)

	// 2. Convert the whole array to a slice view, then sort it.
	scores := [5]int{88, 95, 72, 100, 85}
	fmt.Println("scores before sorting:", scores)

	slices.Sort(scores[:])
	fmt.Println("scores in ascending order:", scores)

	// 3. Sort first, then reverse for descending order.
	slices.Reverse(scores[:])
	fmt.Println("scores in descending order:", scores)

	// 4. Strings can also be sorted.
	players := []string{"Messi", "Kobe", "Rodri", "Durant"}
	slices.Sort(players)
	fmt.Println("Players in ascending order:", players)

	slices.Reverse(players)
	fmt.Println("Players in descending order:", players)
}
