package main

import "fmt"

type Stack struct {
	stack   [10]int
	current int
}

func (this *Stack) Push(v int) {
	this.stack[this.current] = v
	this.current++
}

func (this *Stack) Top() int {
	return this.stack[this.current-1]
}

func (this *Stack) IsEmpty() bool {
	return this.current == 0
}

func (this *Stack) Pop() int {
	this.current--
	return this.stack[this.current]
}

func main() {
	stack := new(Stack)

	stack.Push(5)
	stack.Push(10)
	stack.Push(15)

	fmt.Println("The last value pushed was ", stack.Top())

	for !stack.IsEmpty() {
		fmt.Println("Popping value ", stack.Pop())
	}
}
