package main

import (
	"fmt"
	"time"
)

func calculateSum(iteration int) {
	sum := 0
	for i := 0; i < iteration; i++ {
		sum += i
	}
	fmt.Printf("%d total\n", sum)
}

func main() {
	//calculateSum(10)
	//calculateSum(20)
	//calculateSum(30)

	c := make(chan string)
	go calculateSumConcurrent(c, 332)
	go calculateSumConcurrent(c, 331)
	go calculateSumConcurrent(c, 330)

	tc := time.After(1 * time.Microsecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			fmt.Print(result)
		case <-tc:
			fmt.Print("Timed out\n")
			return
		}
	}
}

func calculateSumConcurrent(c chan<- string, iteration int) {
	sum := 0
	for i := 0; i < iteration; i++ {
		sum += i
	}
	time.Sleep(12)
	c <- fmt.Sprintf("%d total\n", sum)
}
