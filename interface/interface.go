package main

import "fmt"

type HelloWorld struct {
}

func (hw *HelloWorld) String() string {
	return "hello"
}

func main() {
	fmt.Printf("%s\n", new(HelloWorld))
}
