package main

import "fmt"

func main() {

	//array biasa
	// fruit := [...]string{"apple", "banana", "melon", "orange"}
	// fmt.Println("fruit => ", len(fruit))
	// fmt.Println("fruit => ", fruit)

	//array multidimensi
	// numbers := [2][2]int{[2]int{3, 2}, [2]int{3, 4}}
	// numberss := [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers ", numbers)
	// fmt.Println("numberss ", numberss)

	//loop array
	// fruit := [...]string{"apple", "banana", "melon", "orange"}
	// for i := 0; i < len(fruit); i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Printf("element %d : %s \n", i, fruit[i])
	// }

	fruits := [...]string{"apple", "banana", "melon", "orange"}
	for i, fruit := range fruits {
		fmt.Printf("index %d : %s \n", i, fruit)
	}
}
