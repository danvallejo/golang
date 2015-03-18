package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 100; i++ {
		if i < 98 {
			continue
		}
		if i == 99 {
			break
		}
		fmt.Println(i)
	}
}
