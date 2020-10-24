// Measure time difference between echo1, echo2, and echo3
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start)
	fmt.Println("Time for echo1 = ", elapsed)

	s, sep = "", ""
	start = time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg // New string is created everytime
		sep = " "
	}
	fmt.Println(s)
	elapsed = time.Since(start)
	fmt.Println("Time for echo2 = ", elapsed)

	s, sep = "", ""
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed = time.Since(start)
	fmt.Println("Time for echo3 = ", elapsed)
}
