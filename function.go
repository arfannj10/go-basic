package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var names = []string{"john", "wick"}
	// printMessage("halo", names)

	// rand.Seed(time.Now().Unix())
	// var randomValue int

	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number : ", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number : ", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("random number : ", randomValue)

	divideNumber(22, 0)
}

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot be divided by %d\n", m, n)
		return
	}

	res := m / n
	fmt.Printf("%d / %d = %d \n", m, n, res)
}
