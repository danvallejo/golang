package main

import "fmt"

func multArray(x [4]int) {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] * 2
	}
}

func multSlice(x []int) {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] * 2
	}
}

func main() {
	arr := [4]int{1, 2, 3, 4}
	multArray(arr)
	fmt.Printf("%p", &arr)
	fmt.Println(arr)

	s := arr[0:len(arr)]
	multSlice(s)
	fmt.Printf("%p", s)
	fmt.Println(s)

	fmt.Println(arr)

}
