package shape

import (
	"fmt"
)

const (
	pi = float64(3.14)
	E = 2.7
)

type Shape interface {
	Area() float64
}

type Circle struct {
	X      int
	Y      int
	Radius int
}

func (c *Circle) Area() float64 {
	return pi * float64(c.Radius*c.Radius)
}

func (c Circle) String() string {
	return fmt.Sprintf("{x=%d, y=%d, radius=%d}", c.X, c.Y, c.Radius)
}
