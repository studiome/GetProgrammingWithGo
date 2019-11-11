package main

import (
	"fmt"
)

type Planets []string

func (p Planets) terraform() Planets {
	newPlanets := p
	for i := range p {
		newPlanets[i] = "New" + p[i]
	}
	return newPlanets
}

func main() {
	planets := []string{"Mars", "Uranus", "Neptune"}
	fmt.Println(Planets(planets).terraform())
}
