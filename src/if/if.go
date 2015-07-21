package main

import (
	"fmt"
)

func main() {
	i := 3

	if i%2 == 0 {
		fmt.Println(i)
	}

	if t := 44; t%2 == 0 {
		fmt.Println(t)
	}

	i = 2
	if i == 2 {
		fmt.Println(2)
	} else if i%3 == 0 {
		fmt.Println("Divisible by 3")
	} else {
		fmt.Println("Nope")
	}
}
