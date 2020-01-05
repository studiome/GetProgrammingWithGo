package main

import (
	"fmt"
	"strings"
)

func sourceGophoer(downstream chan string) {
	for _, v := range []string{"I am Gopher.", "You are Gopher, too"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		for _, word := range strings.Fields(item) {
			downstream <- word
		}
	}
	close(downstream)

}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGophoer(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
