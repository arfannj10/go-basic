package main

import (
	"fmt"
)

func main() {
	// var a int = 4
	// var b *int = &a

	// fmt.Println("number a (value) : ", a)
	// fmt.Println("number b (address) : ", b)

	// fmt.Println("number (value) : ", *b)
	// fmt.Println("number (address) : ", b)

	// fmt.Println("")

	// a = 5

	// fmt.Println("a (value) : ", a)
	// fmt.Println("a (address) : ", &a)
	// fmt.Println("number (value) : ", *b)
	// fmt.Println("number (address) : ", b)

	var number = "arfan"
	fmt.Println("before : ", number)

	change(&number, "joko")
	fmt.Println("after : ", number)

}

func change(original *string, value string) {
	*original = value
}
