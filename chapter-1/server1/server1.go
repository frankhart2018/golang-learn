// Server1 is a minimal echo server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler which handles all requests beginning with /
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r
func handler(w http.ResponseWriter, r *http.Request) { // r here is a struct of type http.Request
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
