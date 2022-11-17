package main

import (
	"fmt"
	"go-basic/library"
)

func main() {
	// var s1 = library.Student{"ethan", 21}
	// fmt.Println("name ", s1.Name)
	// fmt.Println("grade ", s1.Grade)

	fmt.Printf("Name : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)

}
