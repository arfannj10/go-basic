package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// var secret interface{}

	// secret = "ethan hunt"
	// fmt.Println(secret)

	// secret = []string{"apple", "manggo", "banana"}
	// fmt.Println(secret)

	// secret = 12.4
	// fmt.Println(secret)

	// var data map[string]interface{}

	// data = map[string]interface{}{
	// 	"name":      "ethan hunt",
	// 	"grade":     2,
	// 	"breakfast": []string{"apple", "manggo", "banana"},
	// }

	// fmt.Println("nama : ", data["name"], "grade : ", data["grade"], "fruit : ", data["breakfast"])

	//casting var interface kosong
	// var secrete interface{}

	// secrete = 2
	// var number = secrete.(int) * 10
	// fmt.Println(secrete, "multiplied by 10 us : ", number)

	// secrete = []string{"apple", "manggo", "banana"}
	// var gruits = strings.Join(secrete.([]string), ", ")
	// fmt.Println(gruits, "is my favorite fruits")

	//casting var interface kosong ke pointer objek
	// var secret interface{} = &person{name: "wick", age: 27}
	// var name = secret.(*person).name
	// fmt.Println(name)

	//kombinasi slice, map, dan interface
	var person = []map[string]interface{}{
		{"name": "wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 29},
		[]string{"nanas", "sabreng", "otok"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}
