package main

import (
	"fmt"
)

func main() {
	var values [5]int

	values[0] = 0
	values[1] = 1
	values[2] = 2
	values[3] = 3
	values[4] = 4

	fmt.Println(values)

	t := [2]int{
		5,
		4,
	}
	fmt.Println(t)

	scores := []float64{88.1, 72.2, 93.4, 104.0, 65.8}
	fmt.Println(scores)

	total := 0.0
	for _, v := range scores {
		total += v
	}
	fmt.Println("Total=", total)
	fmt.Println("Average=", total/float64(len(scores)))
}
