// Modify fetchall so that the output is printed to a file
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // Create a channel of strings
	for num, url := range os.Args[1:] {
		go fetch(url, ch, num) // Start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // Receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, num int) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // Send to channel ch
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		ch <- fmt.Sprint(err)
	}

	err = ioutil.WriteFile(strconv.Itoa(num)+".txt", bodyBytes, 0644) // ioutil.Discard discards the output
	resp.Body.Close()                                                 // Don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}

// A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.
