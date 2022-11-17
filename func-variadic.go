package main

import (
	"fmt"
	"strings"
)

func main() {
	//cara 1
	// var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	// var msg = fmt.Sprintf("rata - rata : %.2f", avg)
	// fmt.Println(msg)

	//cara 2
	// var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	// var avg = calculate(numbers...)
	// var msg = fmt.Sprintf("Rata - rata : %.2f", avg)
	// fmt.Println(msg)

	hobbies := []string{"Eat", "Game"}
	yourHobbies("Arfan", hobbies...)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, My name is : %s\n", name)
	fmt.Printf("My Hobbies are : %s\n", hobbiesAsString)
}
