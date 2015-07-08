package main

import "fmt"

func main() {
	fmt.Println("=", 4+2-1*4/2%3)

	fmt.Println("=", 4 + 2 - ((1*4) / (2%3)))
}