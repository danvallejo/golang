package main

import (
	"fmt"
)

var globalInteger int

func main() {
	var str string = "Hello"
	str += " World"
	fmt.Println(str)

	pi := 3.14
	fmt.Println(pi)

	var e float32
	e = 2.7
	fmt.Println(e)

	fmt.Println(globalInteger)
}
