package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go calculateSumConcurrent(c, 332)
	go calculateSumConcurrent(c, 331)
	go calculateSumConcurrent(c, 330)

	ReadResults(c)
}

func ReadResults(c <-chan string) {
	tc := time.After(30 * time.Microsecond)
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
