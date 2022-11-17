package main

import (
	"fmt"
	"reflect"
)

type murid struct {
	Name  string
	Grade int
}

func main() {
	// var number = 23
	// var reflectiveValue = reflect.ValueOf(number)
	// var nilai = reflectiveValue.Interface().(int)

	// fmt.Println("tipe variable : ", reflectiveValue.Type())
	// fmt.Println("tipe variable : ", reflectiveValue.Interface())
	// fmt.Println("tipe variable : ", nilai)

	// if reflectiveValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variable : ", reflectiveValue.Int())
	// }

	var s1 = &murid{Name: "wick", Grade: 2}
	fmt.Println("name : ", s1.Name)
	// s1.getPropertyInfo()

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wi"),
	})

	fmt.Println("nama : ", s1.Name)
}

func (s *murid) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("name : ", reflectType.Field(i).Name)
		fmt.Println("tipe data : ", reflectType.Field(i).Type)
		fmt.Println("nilai : ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *murid) SetName(name string) {
	s.Name = name
}
