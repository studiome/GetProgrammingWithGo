package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg."
	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c -= 3
			if c < 'a' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
}
