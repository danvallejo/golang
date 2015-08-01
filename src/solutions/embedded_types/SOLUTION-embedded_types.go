package main

import "fmt"

type Shape struct {
	Name string
}

type Circle struct {
	Shape
	Radius int
}

func (this Shape) Print(){
	fmt.Println(this.Name)
}

func (this Circle) xPrint() {
	fmt.Println(this.Name, this.Radius)
}

func main() {
	circle := Circle{Shape: Shape{Name: "circle"}, Radius: 22}

	circle.Print()
	
}













