package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 2, 3)
	slice[0] = 1
	slice[1] = 2
	fmt.Println("slice len=", len(slice), "cap=", cap(slice), "array=", slice)

	d := slice[:cap(slice)]
	fmt.Println(d)

	newSlice := make([]int, 2)
	newSlice[0] = 101
	newSlice[1] = 102
	fmt.Println("newSlice len=", len(newSlice), "cap=", cap(newSlice), "array=", newSlice)

	// append takes a list so we convert the slice/array to element list
	slice = append(slice, newSlice...)

	fmt.Println("appended len=", len(slice), "cap=", cap(slice), "array=", slice)

	arr := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("arr %p %v\n", &arr, arr)

	s := arr[:]
	fmt.Printf("slice %p %v\n", s, s)

	s = append(s, []int{100, 200, 300}...)
	s = append(s, 99, 98, 97)
	fmt.Printf("appended slice %p %v\n", s, s)
}
