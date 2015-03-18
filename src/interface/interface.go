package main

import "fmt"

type HelloWorld struct {
}

func (*HelloWorld) String() string {
	return "hello"
}

func main() {
	fmt.Println(new(HelloWorld))
}
