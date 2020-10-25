// Boiling prints the boiling point of water
package main

import "fmt"

const boilingF = 212.0 // package level declaration

func main() {
	// Both f and c are local to main() function
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %g°F or %g°C\n", f, c)
}
