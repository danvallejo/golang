package main

import (
	"fmt"
)

func printNumbers() {
	fmt.Println("printNumbers:start")
	panic("Error occured")
	fmt.Println("printNumbers:end")
}

func errorHandler() {
	str := recover()
	if str != nil {
		fmt.Println(str)
	}
}

func main() {
	defer errorHandler()
	fmt.Println("main:start")
	printNumbers()
	fmt.Println("main:end")
}
