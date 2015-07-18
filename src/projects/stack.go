package main

import "fmt"

func main(){
	stack := new(Stack)
	
	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	
	fmt.Println("The last value pushed was ", stack.Top())
	
	for !stack.IsEmpty(){
		fmt.Println("Popping value ", stack.Pop())
	}
}