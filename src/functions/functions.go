package main

import (
	"fmt"
)

func total(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func calc(values []int) (int, int) {
	sum := total(values)
	average := sum / len(values)
	return sum, average
}

func add(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func makeEvenGenerator() func() int {
	v := 0
	return func() int {
		v += 2
		return v
	}
}

func main() {
	values := []int{10, 20, 30}

	fmt.Println(total(values))

	fmt.Println(calc(values))

	fmt.Println(add(10, 20, 30, 40))

	subtract := func(x, y int) int {
		return x - y
	}
	fmt.Println(subtract(5, 2))

	x := 11
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
