package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func readTextFile(fileName string) (string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		return "", errors.New("Can't find file:" + fileName)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	var b []byte = make([]byte, fileInfo.Size())
	file.Read(b)
	return string(b), nil
}

func simpleHtml(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	fileName := req.RequestURI[1:] + ".html"
	fmt.Println(fileName)

	contents, _ := readTextFile(fileName)

	io.WriteString(res, contents)
}

func main() {
	fmt.Println(readTextFile("hello.html"))

	http.HandleFunc("/hello", simpleHtml)
	http.HandleFunc("/bye", simpleHtml)

	fmt.Println(http.ListenAndServe(":8080", nil))
}
