package main

import (
	"fmt"
	"time"
)

func calculateSum(set int, iteration int) {
	time.Sleep(1 * time.Second)
	sum := 0
	for i := 0; i < iteration; i++ {
		sum += i
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Printf("%d - %d total\n", set, sum)
}

func main() {
	go calculateSum(1, 230)
	go calculateSum(2, 210)
	go calculateSum(3, 240)

	time.Sleep(1220 * time.Millisecond)
}
