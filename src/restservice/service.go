package main

import "github.com/go-martini/martini"

func get() string {
	return "Hello world!"
}

func notfound() string {
	return "Sorry"
}

// before building
// execute: go get github.com/go-martini/martini
func main() {
	m := martini.Classic()

	m.Get("/", get)
	m.NotFound(notfound)

	m.Run()
}
