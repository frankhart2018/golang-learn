// Package tempconv performs Celsius and Fahrenheit temperature computations
package tempconv

type Celsius float64
type Fahrenheit float64 // The type declarations should begin with uppercase character

// Even though both (Celsius and Fahrenheit) have the same underlying type but they don't have the same type and hence
// cannot be compared or combined in arithmetic expressions, comparing them will generate Compiler Error.

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
