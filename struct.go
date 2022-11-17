package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	// name  string
	grade int
	person
}

func main() {
	// var s1 student
	// s1.name = "john wick"
	// s1.grade = 2

	// fmt.Println("name : ", s1.name)
	// fmt.Println("class : ", s1.grade)

	// var s1 = student{}
	// s1.name = "wick"
	// s1.grade = 2

	// var s2 = student{"ethan", 2}

	// var s3 = student{name: "jason"}

	// fmt.Println("student1 : ", s1.name)
	// fmt.Println("student2 : ", s2.name)
	// fmt.Println("student3 : ", s3.name)

	// var s1 = student{name: "wick", grade: 2}

	// var s2 student = s1
	// fmt.Println("student 1, name : ", s1.name)
	// fmt.Println("student 4, name : ", s2.name)

	// s2.name = "ethan"
	// fmt.Println("student 1, name : ", s1.name)
	// fmt.Println("student 4, name : ", s2.name)

	// var s1 = student{}
	// s1.name = "wick"
	// s1.age = 21
	// s1.grade = 2

	// fmt.Println("name : ", s1.name)
	// fmt.Println("age : ", s1.age)
	// fmt.Println("age : ", s1.person.age)
	// fmt.Println("grade : ", s1.grade)

	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 23},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

}
