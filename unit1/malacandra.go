package main

import "fmt"

func main() {
	const distance = 56000000 //km
	const etaDay = 28         //day
	var speed = distance / (etaDay * 24)
	fmt.Println(speed, "km/h")
}
