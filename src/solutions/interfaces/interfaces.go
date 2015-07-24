package main

import "fmt"
import "sort"

type Account struct {
	Name    string
	Balance int
}

type Accounts []Account

func (this Accounts) Len() int {
	return len(this)
}
func (this Accounts) Less(i, j int) bool {
	return this[i].Balance < this[j].Balance
}
func (this Accounts) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	accounts := Accounts{
		{"Savings", 100},
		{"Checking", 500},
	}

	sort.Sort(accounts) // sort by balance
	fmt.Println(accounts)
}
