package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type People []Person

func (this People) Len() int {
	return len(this)
}
func (this People) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this People) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Zach", 9},
		{"Abe", 10},
	}
	sort.Sort(People(kids))
	fmt.Println(kids)
}
