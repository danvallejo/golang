package main

import "fmt"

const factor float64 = 5.0 / 9.0

var celsius float64

func main() {
	fahr := 78.0
	
	celsius = (fahr - 32) * factor
	
	fmt.Printf("%T %v", celsius, celsius)
}
