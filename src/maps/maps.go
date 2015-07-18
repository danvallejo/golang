package main

import (
	"fmt"
)

func main() {
	var toys map[string]int

	fmt.Println(toys)

	toys["rattle"] = 2

	fmt.Println(toys["rattle"])
}
