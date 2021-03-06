package main

import "fmt"
import "container/list"

func main() {
	var x list.List

	x.PushBack(5)
	x.PushFront(1)
	x.PushBack(3)
	x.PushBack(4)
	x.PushFront(6)

	for element := x.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
