package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Dog) Woof() {
	fmt.Println("Woof")
}

type Android struct {
	Person
	Dog
	Model string
}

func main() {
	a := new(Android)
	a.Person.Name = "dave"
	a.Person.Talk() // has-a

	a.Talk() //is-a
	a.Woof() //is-a
}
