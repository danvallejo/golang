package main

import "fmt"

var total int

func calculateSum(max int) {
	for i := 0; i < max; i++ {
		total += i
	}
}

func main() {
	calculateSum(10)

	fmt.Println("Total ", total)
}
