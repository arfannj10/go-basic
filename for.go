package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("angka 1 = ", i)
	// }

	// var a = 1
	// for a <= 5 {
	// 	fmt.Println("angka 2 = ", a)
	// 	a++
	// }

	// var b = 0
	// for b < 20 {
	// 	fmt.Println("angka 3 = ", b)
	// 	b++
	// }

	// for c := 0; c <= 10; c++ {

	// 	if c%2 == 1 {
	// 		continue
	// 	}

	// 	// if c > 8 {
	// 	// 	break
	// 	// }

	// 	fmt.Println(c)
	// }

	for d := 0; d < 5; d++ {
		for e := d; e < 5; e++ {
			fmt.Print(e, " ")
		}
		fmt.Println()
	}

	// outerloop:
	// 	for f := 0; f < 5; f++ {
	// 		for g := 0; g < 5; g++ {
	// 			if f == 3 {
	// 				break outerloop
	// 			}
	// 			fmt.Printf("matriks [", f, "][", g, "]", "\n")
	// 		}
	// 	}

}
