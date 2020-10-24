// Server3 has a handler that echoes the HTTP request
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler which handles all requests beginning with /
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}

// Handler echoes the HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	/*
		The long form of the above if condition is:-

		err := r.ParseForm()
		if err != nil {
			log.Print(err)
		}

		Go allows simple statements to precede if statement, in this case the scope of err got reduced which is a good practice
	*/

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
