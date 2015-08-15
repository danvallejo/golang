package main

import "fmt"
import "container/list"

func Even(x list.List) {
	var element = x.Front()
	for element != nil {
		val := element.Value.(int)

		fmt.Println(element.Value.(int))

		if val%2 == 0 {
			fmt.Println("Deleting..")
			x.Remove(element)
			element = x.Front()
		} else {
			element = element.Next()
		}
	}
}

func main() {
	var x list.List

	x.PushBack(5)
	x.PushFront(4)
	x.PushFront(1)
	x.PushFront(6)
	x.PushBack(3)

	for element := x.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

	fmt.Println()

	Even(x)

	fmt.Println()

	for element := x.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
