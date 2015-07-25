package main

import "fmt"

func addArray(x [4]int) {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] * 2
	}
}

func addSlice(x []int) {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] * 2
	}
}

func main() {
	arr := [4]int{1, 2, 3, 4}
	addArray(arr)
	fmt.Printf("%p", &arr)
	fmt.Println(arr)

	s := arr[0:len(arr)]
	addSlice(s)
	fmt.Printf("%p", s)
	fmt.Println(s)

	fmt.Println(arr)

}
