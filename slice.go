package main

import "fmt"

func main() {
	// fruits := []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(fruits[0])
	// newFruits := fruits[0:4]
	// fmt.Println(newFruits)

	// afrutis := fruits[0:2]
	// bfrutis := fruits[1:4]

	// aafrutis := fruits[0:2]
	// bafrutis := fruits[0:1]

	// fmt.Println(afrutis)
	// fmt.Println(bfrutis)
	// fmt.Println(aafrutis)
	// fmt.Println(bafrutis)

	// bafrutis[0] = "pinaple"
	// fmt.Println(afrutis)
	// fmt.Println(bfrutis)
	// fmt.Println(aafrutis)
	// fmt.Println(bafrutis)
	// fmt.Println(bafrutis)

	//append di slice
	// cfruits := append(fruits, "otok")
	// fmt.Println(fruits)
	// fmt.Println(cfruits)
	// fmt.Println(cap(fruits))
	// fmt.Println(len(fruits))

	// fmt.Println(len(afrutis))
	// fmt.Println(cap(afrutis))
	// fmt.Println(afrutis)

	// dfruits := append(afrutis, "sabreng")
	// fmt.Println(dfruits)

	//copy slice
	// dst := make([]string, 3)
	// dst := []string{"potato", "fries", "milk"}
	src := []string{"watermelon", "pineapple", "apple", "orange"}
	// n := copy(dst, src)

	s := src[0:3:3]
	fmt.Println(s)

	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)
}
