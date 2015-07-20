package main

import (
	"fmt"
)

func main() {
	var toys map[string]int = make(map[string]int)

	fmt.Println(toys)

	toys["rattle"] = 2
	toys["top"] = 4

	fmt.Println(toys["rattle"])

	fmt.Println(toys)
}
