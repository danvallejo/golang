package main

import "fmt"

func main() {
	var month int

	fmt.Print("Enter month (1-12):")

	fmt.Scanf("%v\n", &month)

	switch month {
	case 1:
		fmt.Println("January has 31 days")
	case 2:
		fmt.Println("February has 28 days")
	default:
		fmt.Println("Invalid month")
	}
}
