package main

import "fmt"

type Account struct {
	Name    string
	Balance int
}

type Accounts []Account

func main() {
	accounts := Accounts{
		{"Savings", 100},
		{"Checking", 500},
	}

	fmt.Println(accounts)
}
