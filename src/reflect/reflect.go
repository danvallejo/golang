package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func myPrint(args ...interface{}) {
	for _, arg := range args {
		str := "{unknown-type}"
		v := reflect.ValueOf(arg)
		switch v.Kind() {
		case reflect.String:
			str = v.String()
		case reflect.Int:
			str = strconv.FormatInt(v.Int(), 10)
		}
		fmt.Println(reflect.TypeOf(arg), "["+str+"]")
	}
}

func main() {
	myPrint("Hello World! ", 42, "\n", 'c')
}
