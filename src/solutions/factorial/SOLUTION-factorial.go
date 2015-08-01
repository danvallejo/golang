package main

import "fmt"

func main() {
	fmt.Print("Enter a number:")

	var number int
	fmt.Scanf("%v\n", &number)

	factorial := 1

	for index := 1; index <= number; index++ {
		factorial *= index
		fmt.Printf("%v! = %v\n", index, factorial)
	}
}
