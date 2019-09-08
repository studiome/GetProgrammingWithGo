package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var answer = 50
	for {
		guess := rand.Intn(100) + 1
		if answer == guess {
			fmt.Printf("Bingo, answer is %v.\n", guess)
			break
		} else if answer < guess {
			fmt.Println("No, Larger.")
		} else {
			fmt.Println("No, Smaller.")
		}
	}
}
