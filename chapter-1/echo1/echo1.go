// echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// := is assignment with type inference kind of
	// There is only postfix increment and decrement
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*
	a traditional "while" loop, since there is no initialization or updation then we don't have to give ; (semicolons)
	for condition {
	// ...
	}

	a traditional infinite loop
	for {
	// ...
	}
*/
