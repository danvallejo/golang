package main

import "fmt"

func swap(x *int, y *int) {
	t := *x
	*x = *y
	*y = t
}

func swapper(x,y int) (int,int){
	return y,x
}

func main() {
	x := 1
	y := 100

	fmt.Println("x=", x, "y=", y)
	swap(&x, &y)
	fmt.Println("x=", x, "y=", y)

	z := 50
	x, y, z = z, x, y
	fmt.Println("x=", x, "y=", y, "z=", z)
}
