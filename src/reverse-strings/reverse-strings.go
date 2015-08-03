package main

import "fmt"

func ErrorCheck() {
	error := recover()
	fmt.Println("Error", error)
}

func ReverseStringsAttemptANSI() {
	fmt.Println("\n---ATTEMPT #1---\n")

	input := "The quick brown fox jumped over the lazy moon"

	for i, ch := range input {
		fmt.Println(i, ch, string(ch))
	}

	output := make([]byte, len(input))
	begin := 0
	end := len(input) - 1

	for begin < len(input) {
		output[begin] = input[end]
		begin++
		end--
	}

	fmt.Println("reverse=", string(output))
}

func ReverseStringsAttemptUNICODE() {
	fmt.Println("\n---ATTEMPT #2---\n")
	input := "The quick brown 狐 jumped over the lazy 犬"

	for i, ch := range input {
		fmt.Println(i, ch, string(ch))
	}

	output := make([]byte, len(input))
	begin := 0
	end := len(input) - 1

	for begin < len(input) {
		output[begin] = input[end]
		begin++
		end--
	}

	fmt.Println("reverse=", string(output))
}

func ReverseStringFixedUNICODE() {
	fmt.Println("\n---ATTEMPT #3---\n")
	input := "The quick brown 狐 jumped over the lazy 犬"

	// Get Unicode code points.
	output := []rune(input)
	n := len(output)

	for i, ch := range output {
		fmt.Println(i, ch, string(ch))
	}

	// Reverse
	for i := 0; i < n/2; i++ {
		output[i], output[n-1-i] = output[n-1-i], output[i]
	}
	// Convert back to UTF-8.
	fmt.Println("reverse=", string(output))
}

func main() {
	ReverseStringsAttemptANSI()
	ReverseStringsAttemptUNICODE()
	ReverseStringFixedUNICODE()
}
