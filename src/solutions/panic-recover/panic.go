package main

import "fmt"

func errorCheck(){
	str := recover()
	if str!= nil{
		fmt.Println(str)
	}
}

func Calc(f float64) float64 {
	defer errorCheck()
	if f == 0.0{
		panic("division by zero")
	}
	return 1/ f;
}

func main(){
	f := 0.0
	fmt.Println(Calc(f))
}