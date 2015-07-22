// strings
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// true
		strings.Contains("tester", "es"))

	fmt.Println(
		// 2
		strings.Count("tester", "t"))

	fmt.Println(
		// true
		strings.HasPrefix("tester", "te"))

	fmt.Println(
		// true
		strings.HasSuffix("test", "st"))

	fmt.Println(
		// 1
		strings.Index("tester", "e"))

	fmt.Println(
		// "john-madison"
		strings.Join([]string{"john", "madison"}, "-"))

	fmt.Println(
		// == "daddaddaddaddad"
		strings.Repeat("dad", 5))

	fmt.Println(
		// "mavim hodden"
		strings.Replace("david hodden", "d", "m", 2))

	fmt.Println(
		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"))

	fmt.Println(
		// "tester"
		strings.ToLower("TESTER"))

	fmt.Println(
		// "TESTER"
		strings.ToUpper("tester"))
}
