package main

import "fmt"

// Speaker is implemented by any type that has a Speak method
// Speak takes no parameters and returns no values
type Speaker interface {
	Speak()
}

// empty interface
// What is the purpose of the empty interface?
type any interface {
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// Dog implements Speaker
func (d Dog) Speak() {
	// 有些我们要做的事情放这里
	fmt.Println(d.Name + "is speaking")
}

func (c Cat) Speak() {
	fmt.Println(c.Name + "在喵喵叫！(=^･ｪ･^=)\")")

}

func makeItSpeak(s Speaker) {
	fmt.Println("准备让它说话")
	s.Speak()
}

func main() {
	var dog Dog = Dog{Name: "小蓝"}
	var cat Cat = Cat{Name: "小猫"}
	// equals  dog := Dog{Name: "小蓝" }

	makeItSpeak(dog)
	makeItSpeak(cat)

}
