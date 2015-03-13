package main

import (
	"fmt"
)

func calculateSum(iteration int) {
	sum := 0
	for i := 0; i < iteration; i++ {
		sum += i
	}
	fmt.Printf("%d total\n", sum)
}

func main() {
	calculateSum(10)
	calculateSum(20)
	calculateSum(30)
}
