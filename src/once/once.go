package main

import "fmt"
import "sync"

var name string
var once sync.Once

func setName() {
	name = "Mike"
	fmt.Println("Setting Mike")
}

func main() {

	for i := 0; i < 5; i++ {
		once.Do(setName)

		fmt.Println(name)
	}
}
