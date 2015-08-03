package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}

func createTextFile(person Person) {
	file, err := os.Create("person.dat")
	if err != nil {
		return
	}
	defer file.Close()

	jsonBytes, _ := json.Marshal(person)

	file.Write(jsonBytes)
}

func readTextFile() {
	file, err := os.OpenFile("person.dat", os.O_RDONLY, 0)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	var b []byte = make([]byte, fileInfo.Size())
	_, err = file.Read(b)
	if err != nil {
		return
	}

	var person Person
	json.Unmarshal(b, &person)

	fmt.Println(person)
}

func main() {
	var person = Person{Name: "dave", Age: 22, Phone: "111-222-3333"}
	createTextFile(person)
	readTextFile()
}
