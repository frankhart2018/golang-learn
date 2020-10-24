// Fetchall fetches URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // Create a channel of strings
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // Receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // Send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // ioutil.Discard discards the output
	resp.Body.Close()                                 // Don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.
