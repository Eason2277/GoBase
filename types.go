package main

import "fmt"

func main() {

	var number1 uint8 = 252
	fmt.Println(number1)

	//byte is an alias for unit8
	var char1 byte = 'a'
	fmt.Println(char1)

	// rune is an alias for unit32
	var char2 rune = '足'
	// unicode
	fmt.Println(char2)

	fmt.Printf("%c %c\n", char1, char2)

	var name string = "西班牙加油"
	fmt.Println(name)

	// escape sequence that represents a newline
	//\n  newline      换行
	//\t  tab          制表符
	//\"  double quote 双引号
	//\\  backslash    反斜杠

	// boolean type
	var hasPermission bool = true
	fmt.Println(hasPermission)

	// zero-value problem
	//int      → 0
	//float64  → 0.0
	//bool     → false
	//string   → ""
	//pointer  → nil
	//slice    → nil
	//map      → nil

	var number2 int32
	fmt.Println(number2)
}
