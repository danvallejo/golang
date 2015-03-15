package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 2, 3)
	slice[0] = 1
	slice[1] = 2
	fmt.Println("len=", len(slice), "cap=", cap(slice), "array=", slice)

	d := slice[:cap(slice)]
	fmt.Println(d)

	newSlice := make([]int, 4)
	newSlice[0] = 101
	//newSlice[1] = 102
	fmt.Println("len=", len(newSlice), "cap=", cap(newSlice), "array=", newSlice)

	slice = append(slice, newSlice...)

	fmt.Println("len=", len(slice), "cap=", cap(slice), "array=", slice)
}
