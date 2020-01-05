package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil {
		return
	}
	if i != nil {
		c.leftHand = i
	}
}

func (c *character) give(to *character) {
	if c == nil {
		return
	}
	if to.leftHand != nil {
		return
	}
	if to != nil {
		to.leftHand = c.leftHand
		c.leftHand = nil
	}

}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v has nothing.", c.name)
	}
	return fmt.Sprintf("%v has %v.", c.name, c.leftHand.name)
}

func main() {
	arthur := &character{
		name: "Arthur",
	}
	knight := &character{
		name: "a kinght",
	}
	fmt.Println(arthur)
	fmt.Println(knight)

	arthur.pickup(&item{
		name: "Caliburn",
	})
	fmt.Println(arthur)
	arthur.give(knight)
	fmt.Println(arthur)
	fmt.Println(knight)

}
