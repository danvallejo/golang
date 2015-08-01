package main

import "fmt"

func main() {
	months := map[int]string{
		1: "January",
		2: "February",
		3: "March",
	}

	monthNumber := 1

	name, valid := months[monthNumber]

	if valid {
		fmt.Printf("Month is %s\n", name)
	} else {
		fmt.Printf("Month %v not found\n", monthNumber)
	}
}
