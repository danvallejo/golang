package main

import "fmt"

func main() {
	numbers := []int{5, -2, 100, 12, 4, -5}

	min := numbers[0]
	max := min

	for _, number := range numbers {

		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	fmt.Printf("Min = %v\n", min)
	fmt.Printf("Max = %v\n", max)
}
