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
	case 2: // this is an implied break
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4, 5 or 6")
	default:
		fmt.Println("unknown")
	}

	switch {
	case i >= 0 && i <= 1:
		fmt.Println("0 or 1")
	case i >= 2 && i <= 10:
		fmt.Println("2..10")
	}

	var t interface{}
	t = 4.5
	switch t.(type) {
	default:
		fmt.Printf("Unexpected type %T", t)
	case int:
		fmt.Printf("Integer %T", t)
	case float64:
		fmt.Printf("Float64 %T", t)
	}
}
