package main

import (
	"fmt"
	"shape"
)

func calculateArea(shapes ...shape.Shape) float64 {
	sum := float64(0)

	for _, v := range shapes {
		sum += v.Area()
	}

	return sum
}

func main() {
	circle := shape.Circle{X: 1, Y: 2, Radius: 2}

	fmt.Println(circle, circle.Area(), calculateArea(&circle))
}
