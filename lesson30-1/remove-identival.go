package main

import "fmt"

func sourceGophoer(downstream chan string) {
	for _, v := range []string{"a", "a", "b", "b", "b", "c", "c", "c", "c"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	previous := ""
	for item := range upstream {
		if item != previous {
			downstream <- item
			previous = item
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
