// echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg // New string is created everytime
		sep = " "
	}
	fmt.Println(s)
}

/*
	Types of function declarations:-

	1) s := "" - Short variable declaration, most compact, but it can be used only in function not for package-level variables.
	2) var s string - Relies on default initialization to zero value for strings which is "".
	3) var s = "" - This is rarely used, except when multiple declarations of variables are done.
	4) var s string = ""
*/
