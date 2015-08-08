package main

import "fmt"
import "_SOLUTIONS_/accounts"

func main() {
	accounts := accounts.Accounts{
		{"Savings", 100},
		{"Checking", 500},
	}

	fmt.Println(accounts)
}
