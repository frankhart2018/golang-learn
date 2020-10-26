package tempconv

// CToF converts a celsius temperature to fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a fahrenheit temperature to celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CtoK converts a celsius temperature to kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// KToC converts a kelvin temperature to celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// FToK converts a fahrenheit temperature to kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(FToC(f) + 273.15)
}

// KToF converts a kelvin temperature to fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(CToF(KToC(k)))
}
