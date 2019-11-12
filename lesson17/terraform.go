package main

import (
	"fmt"
)

type Planets []string

func (p Planets) terraform() Planets {
	newPlanets := make(Planets, len(p))
	for i := range p {
		newPlanets[i] = "New" + p[i]
	}
	return newPlanets
}

func main() {
	planets := []string{"Mars", "Uranus", "Neptune"}
	fmt.Println(planets)
	fmt.Println(Planets(planets).terraform())
	fmt.Println(planets)
}
