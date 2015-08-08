package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.cstructor.com/",
	}

	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)

		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			// Fetch the URL.
			resp, err := http.Get(url)
			fmt.Printf("Got %s\n", url)
			fmt.Println(resp, err)
		}(url)
	}

	// Wait for all HTTP fetches to complete.
	fmt.Println("Waiting...")
	wg.Wait()

	fmt.Println("Done.")
}
