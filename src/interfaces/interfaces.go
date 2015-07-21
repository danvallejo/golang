package main

import (
	"fmt"
)

const (
	pi = float64(3.14)
)

type Shape interface {
	area() float64
}

type Circle struct {
	x      int
	y      int
	radius int
}

func (c *Circle) area() float64 {
	return pi * float64(c.radius*c.radius)
}

// Stringer interface
func (c Circle) String() string {
	return fmt.Sprintf("{x=%d, y=%d, radius=%d}", c.x, c.y, c.radius)
}

func calculateArea(shapes ...Shape) float64 {
	sum := float64(0)

	for _, v := range shapes {
		sum += v.area()
	}

	return sum
}

func main() {
	circle := Circle{x: 1, y: 2, radius: 2}

	fmt.Println(circle, circle.area(), calculateArea(&circle))
}
