package main

import "fmt"

//import "sync"

var name string

//var lock sync.Locker

func setName() {

	name = "Mike"
	//	lock.Unlock()
}

func main() {
	//	lock = new(sync.Mutex)

	//	lock.Lock()
	go setName()

	//	lock.Lock()
	fmt.Println(name)
}
