package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		height: 5,
		base:   2,
	}

	s := square{
		sideLength: 7,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return t.height * 0.5 * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
