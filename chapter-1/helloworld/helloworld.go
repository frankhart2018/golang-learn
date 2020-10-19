package main

import "fmt"

func main() { // This { should be in the same line, if it is in next line then the compiler throws error
	fmt.Println("Hello, World!")
}

// gofmt can be used to format go programs properly
// goimports can be used to remove unnecessary imports
