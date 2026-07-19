package main

import "fmt"

// make and append
func main() {
	// make([]T, length, capacity) creates a slice and prepares its underlying array.

	numbers := make([]int, 2, 4)
	numbers[0] = 10
	numbers[1] = 20

	numbers = append(numbers, 88)

	fmt.Println("先用make创建的numbers这个切片", numbers)

	var length int = len(numbers)
	fmt.Print(length, cap(numbers))

}
