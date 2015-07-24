package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

type Contact struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:phone_number`
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

func getElement(id int, list list.List) *Contact {
	for element := contacts.Front(); element != nil; element = element.Next() {
		contact := element.Value.(Contact)
		if contact.Id == id {
			return &contact
		}
	}

	return nil
}

func getContacts() string {
	arr := listToContactSlice(contacts)

	jsonBytes, _ := json.Marshal(arr)

	fmt.Println(arr)
	fmt.Println(string(jsonBytes))

	return string(jsonBytes)
}

func getSpecific(parms martini.Params) (int, string) {
	id, _ := strconv.Atoi(parms["id"])

	contact := getElement(id, contacts)
	if contact == nil {
		return http.StatusNotFound, ""
	}

	jsonBytes, _ := json.Marshal(contact)

	fmt.Println(*contact)
	fmt.Println(string(jsonBytes))

	return http.StatusOK, string(jsonBytes)
}

func postContact(w http.ResponseWriter, r *http.Request) (int, string) {
	contact := getPostContact(r)
	contacts.PushBack(*contact)
	return http.StatusCreated, "created"
}

func getPostContact(r *http.Request) *Contact {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	phoneNumber := r.FormValue("phone_number")
	return &Contact{Id: id, Name: name, PhoneNumber: phoneNumber}
}

func notfound() string {
	return "Sorry"
}

// before building
// execute: go get github.com/go-martini/martini
func main() {
	m := martini.Classic()

	contacts.PushBack(Contact{Id: 2, Name: "mike", PhoneNumber: "111"})
	contacts.PushBack(Contact{Id: 1, Name: "dan", PhoneNumber: "222"})

	m.NotFound(notfound)
	m.Get("/contacts", getContacts)
	m.Get("/contacts/:id", getSpecific)
	m.Post("/contacts", postContact)

	m.Run()
}
