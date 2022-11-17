package main

import "fmt"

func main() {
	var point = 8785.0

	// if point == 10 {
	// 	fmt.Println("lulus dengan nilai sempurna")
	// } else if point > 5 {
	// 	fmt.Println("lulus")
	// } else if point == 4 {
	// 	fmt.Println("Hampir lulus")
	// } else {
	// 	fmt.Println("Gak lulus")
	// }

	if percent := point / 10; percent >= 1000 {
		fmt.Println("%.1f %s perfect ! \n", percent, "%")
	} else if percent >= 70 {
		fmt.Println("%.1f %s good \n", percent, "%")
	} else {
		fmt.Println("%.1f %s not bad \n", percent, "%")
	}

}
