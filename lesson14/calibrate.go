package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0.0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5.0
	sensor := calibrate(realSensor, offset)
	fmt.Println(sensor())
	offset = 10.0
	fmt.Println(sensor())

	fSensor := calibrate(fakeSensor, offset)
	for i := 0; i < 10; i++ {
		fmt.Println(fSensor())
	}

}
