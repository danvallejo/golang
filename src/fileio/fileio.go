package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func createTextFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// ⌘ = ec2 8 98
	ret, err := file.WriteString("Hello World ⌘")
	if err != nil {
		return
	}
	fmt.Println(ret, "bytes written")
}

func readTextFile() {
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	fmt.Println(fileInfo.Size(), "file size")

	var b []byte = make([]byte, fileInfo.Size())
	n, err := file.Read(b)
	fmt.Println(n, b, string(b))
}

func readFile() {
	bytes, _ := ioutil.ReadFile("test.txt")
	str := string(bytes)
	fmt.Println(str)
}

func readDirectory() {
	dir, _ := os.Open(".")
	defer dir.Close()

	fileInfos, _ := dir.Readdir(-1)
	for _, fi := range fileInfos {
		fmt.Println(fi)
	}
}

func readError() error {
	return errors.New("My new error")
}

func main() {
	fmt.Println(os.Getwd())

	fmt.Println(os.Getenv("WinDir"))

	createTextFile()

	readTextFile()

	readFile()

	readDirectory()

	fmt.Println(readError())
}
