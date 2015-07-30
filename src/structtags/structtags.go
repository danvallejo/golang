package main

import (
	"fmt"
	"reflect"
)

func main() {

	type Animal struct {
		Name string `species:"gopher" color:"blue"`
	}

	animal := Animal{}

	st := reflect.TypeOf(animal)
	fmt.Println("st=",st)

	field := st.Field(0)
	fmt.Println("field=", field)
	fmt.Println("field.Tag=",field.Tag)

	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}
