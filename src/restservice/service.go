package main

import "fmt"
import "encoding/json"
import "container/list"
import "github.com/go-martini/martini"

type Contact struct {
	Name        string
	PhoneNumber string
}

type Contacts list.List

var contacts list.List

func listToContactSlice(list list.List) []Contact {
	arr := make([]Contact, contacts.Len())
	index := 0
	for element := contacts.Front(); element != nil; element = element.Next() {
		arr[index] = element.Value.(Contact)
		index++
	}

	return arr
}

func get() string {
	arr := listToContactSlice(contacts)

	jsonBytes, _ := json.Marshal(arr)

	fmt.Println(arr)
	fmt.Println(string(jsonBytes))

	return string(jsonBytes)
}

func notfound() string {
	return "Sorry"
}

// before building
// execute: go get github.com/go-martini/martini
func main() {
	m := martini.Classic()

	contacts.PushBack(Contact{Name: "mike", PhoneNumber: "111"})
	contacts.PushBack(Contact{Name: "dan", PhoneNumber: "222"})

	fmt.Println(contacts)

	m.Get("/contacts", get)
	m.NotFound(notfound)

	m.Run()
}
