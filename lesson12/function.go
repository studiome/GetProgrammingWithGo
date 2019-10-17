package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0

}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func main() {
	fmt.Println(celsiusToFahrenheit(0.0))
	fmt.Println(kelvinToFahrenheit(0.0))
}
