package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	c := k.celsius()
	return c.fahrenheit()
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	c := f.celsius()
	return c.kelvin()

}

func main() {
	var c celsius = 27.0
	fmt.Println(c, "deg Celsius = ", c.kelvin(), "K = ", c.fahrenheit(), "deg Fahrenheit")
}
