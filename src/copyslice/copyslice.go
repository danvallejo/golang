package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)

	fmt.Println(slice1)

	copy(slice2, slice1)

	fmt.Println("1st attempt", slice2)

	if cap(slice2) < len(slice1) {
		slice2 = make([]int, len(slice1))

		copy(slice2, slice1)

		fmt.Println("2nd attempt", slice2)
	}
}
