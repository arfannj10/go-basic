package main

import (
	"fmt"
)

func main() {
	var point = 20

	// switch point {
	// case 8, 10, 80:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("try again")

	// 	}
	// }

	switch {
	case point == 100:
		fmt.Println("Perfect")
	case (point < 100) && (point > 30):
		fmt.Println("awesome")
		fallthrough
	case point < 50:
		fmt.Println("learn again")
	default:
		fmt.Println("aw")
	}
}
