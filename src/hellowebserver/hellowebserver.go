package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
	)
}

func bye(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
        <title>Goodbye</title>
    </head>
    <body>
        Goodbye
    </body>
</html>`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)

	fmt.Println(http.ListenAndServe(":8080", nil))
}
