package main

import (
	"fmt"
)

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["K"] = "kryptonite"

	fmt.Println(len(elements), elements)

	delete(elements, "K")

	name, valid := elements["K"]
	fmt.Println(name, valid)

	shapes := map[string]string{
		"Circle":    "circ",
		"Rectangle": "rect",
	}
	fmt.Println(shapes)

	for k, v := range shapes {
		fmt.Println("key:", k, "value:", v)
	}
}
