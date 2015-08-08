package main

import "fmt"
import "sync"
import "time"

var wg sync.WaitGroup

func calculate(ch chan<- string) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		ch <- fmt.Sprintf("%v", i)
		fmt.Println("Generating", i)
	}
}

func print(ch <-chan string) {
	for {
		select {
		case result := <-ch:
			fmt.Println(result)
			time.Sleep(500 * time.Millisecond)
			wg.Done()
		}
	}
}

func main() {
	ch := make(chan string, 5)

	go calculate(ch)
	go print(ch)

	// give the go routines a chance to fill in the waitgroup
	time.Sleep(1 * time.Second)
	wg.Wait()
}
