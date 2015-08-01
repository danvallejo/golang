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

// Rename this method to see if Shape.Print is called
func (this Circle) Print() {
	fmt.Println(this.Name, this.Radius)
}

func main() {
	circle := Circle{Shape: Shape{Name: "circle"}, Radius: 22}

	circle.Print()
	
	circle.Shape.Print()	
}













