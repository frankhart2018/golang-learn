// Use io.Copy(dst, src) to copy output to os.Stdout without requiring a buffer to display the HTML to terminal
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) // resp is a structure

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :%v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
