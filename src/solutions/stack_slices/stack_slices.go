package main

import "fmt"

type Stack struct {
	stack   []int
	current int
}

func (this *Stack) Push(v int) {
	this.stack = append(this.stack, v)
	fmt.Println("Push-cap=", cap(this.stack))

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
	v := this.stack[this.current]

    // Shrinking the stack by 1
	newStack := make([]int, cap(this.stack)-1)
	copy(newStack, this.stack)
	this.stack = newStack
	fmt.Println("Pop-cap=", cap(this.stack))

	return v
}

func main() {
	stack := new(Stack)

	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	stack.Push(20)
	stack.Push(25)

	fmt.Println("The last value pushed was ", stack.Top())

	for !stack.IsEmpty() {
		fmt.Println("Popping value ", stack.Pop())
	}
}
