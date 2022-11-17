package main

import "fmt"

func main() {
	// var chicken map[int]string
	// chicken = map[int]string{
	// 	1: "Arfan",
	// 	2: "Hidayatullah",
	// }

	// chicken[1] = "arfan"
	// chicken[2] = "Hidayatullah"

	// fmt.Println("januari", chicken[1], chicken[2])
	// fmt.Println("mei", chicken[2])

	// var chicken = map[string]int{"januari": 50, "februari": 40}
	// var value, isExist = chicken["januari"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("item is not exist")
	// }

	//slice map
	// var chickens = []map[string]string{
	// 	{"name": "chicken blue", "gender": "male"},
	// 	{"name": "chicken red", "gender": "male"},
	// 	{"name": "chicken yellow", "gender": "female"},
	// }

	// for _, chicken := range chickens {
	// 	fmt.Println(chicken["name"], chicken["gender"])
	// }

	var datas = []map[string]string{
		{"name": "Arfan Hidayatullah"},
		{"level": "Advance"},
	}

	for _, data := range datas {
		fmt.Println("data : ", data["name"], data["level"])
	}

}
