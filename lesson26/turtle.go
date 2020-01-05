package main

import "fmt"

type position struct {
	x int
	y int
}

func (p *position) moveRight() {
	p.x++
}

func (p *position) moveLeft() {
	p.x--
}

func (p *position) moveUp() {
	p.y++
}

func (p *position) moveDown() {
	p.y--
}

func main() {
	turtle := position{
		x: 10,
		y: 15,
	}

	turtle.moveRight()
	turtle.moveRight()
	turtle.moveUp()
	turtle.moveUp()
	turtle.moveDown()
	turtle.moveLeft()
	fmt.Println("Turtle is ", turtle, ".")
}
