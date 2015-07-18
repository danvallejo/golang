package main

import "fmt"

func main() {
	var month int

	fmt.Print("Enter month (1-12):")

	fmt.Scanf("%v\n", &month)

	switch month {
	case 1:
		fmt.Println("January")
	default:
		fmt.Println("Invalid month")
	}
}
