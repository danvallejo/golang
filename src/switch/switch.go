package main

import (
	"fmt"
)

func main() {
	i := 1

	switch i {
	case 0:
		{
			fmt.Println("zero")
		}
	case 1:
		fmt.Println(1)
	case 2:
	case 3:
		fmt.Println("2 or 3")
	default:
		fmt.Println("unknown")
	}

	switch {
	case i >= 0 && i <= 1:
		fmt.Println("0 or 1")
	case i >= 2 && i <= 10:
		fmt.Println("2..10")
	}
}
