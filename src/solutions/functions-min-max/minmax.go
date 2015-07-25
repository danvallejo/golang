package main

import "fmt"

func calculateMinMax(values []int) (min, max int) {
	min = values[0]
	max = values[0]

	for _, v := range values {
		if v < min{
			min = v
		}
		if v > max{
			max =v
		}
	}

	return
}

func calculateMinMaxList(values ...int) (min int, max int) {
	min = values[0]
	max = values[0]

	for _, v := range values {
		if v < min{
			min = v
		}
		if v > max{
			max =v
		}
	}

	return
}

func main(){
	numbers := []int{5, -2, 100, 12, 4, -5}

    min, max := calculateMinMax(numbers)

	fmt.Println("arr min=", min, "max=", max)

    min, max = calculateMinMaxList(5, -2, 101, 12, 4, -10)

	fmt.Println("... min=", min, "max=", max)
}