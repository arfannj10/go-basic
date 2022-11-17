package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumreference = calculated(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumreference)
}

func calculated(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	//hitung keliling
	var circumreference = math.Pi * d

	//kembalikan 2 nilai
	return area, circumreference
}
