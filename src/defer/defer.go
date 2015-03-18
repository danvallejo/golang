package main

import (
	"fmt"
)

func helloWorld() {
	fmt.Println("Hello World")
}

func printNumbers() {
	fmt.Println("printNumbers:start")
	defer helloWorld()
	fmt.Println("printNumbers:end")
}

func main() {
	fmt.Println("main:start")
	printNumbers()
	fmt.Println("main:end")
}
