// strings
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// true
		strings.Contains("tester", "es"),

		// 2
		strings.Count("tester", "t"),

		// true
		strings.HasPrefix("tester", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("tester", "e"),

		// "john-madison"
		strings.Join([]string{"john", "madison"}, "-"),

		// == "daddaddaddaddad"
		strings.Repeat("dad", 5),

		// "mavim hodden"
		strings.Replace("david hodden", "d", "m", 2),

		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),

		// "tester"
		strings.ToLower("TESTER"),

		// "TESTER"
		strings.ToUpper("tester"),
	)
}
