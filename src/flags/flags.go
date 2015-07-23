package main

import (
	"flag"
	"fmt"
)

func main() {
	// max is a pointer
	max := flag.Int("max", 6, "(default) max iterations")

	flag.Parse()

	for i, v := range flag.Args() {
		fmt.Println(i, v)
	}

	fmt.Println(*max)
}
