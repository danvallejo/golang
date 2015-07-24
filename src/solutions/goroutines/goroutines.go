package main

import "fmt"
import "sync"

var total int
var lock sync.Locker

func calculateSum(max int) {
	for i := 0; i < max; i++ {
		total += i
	}
	lock.Unlock()
}

func main() {
	lock = new(sync.Mutex)

	lock.Lock()
	go calculateSum(10)

	lock.Lock()
	fmt.Println("Total ", total)
}
