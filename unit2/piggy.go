package main

import (
	"fmt"
	"math/rand"
)

func main() {
	bank := 0.0
	for bank < 20.0 {
		i := rand.Intn(3)
		switch i {
		case 0:
			bank += 0.05
		case 1:
			bank += 0.10
		case 2:
			bank += 0.25
		}
		fmt.Printf("%5.2f\n", bank)
	}
}
