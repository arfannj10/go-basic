package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		message <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("Jason Bourne")
}
