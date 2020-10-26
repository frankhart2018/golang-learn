package tempconv

// CToF converts a celsius temperature to fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a fahrenheit temperature to celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
