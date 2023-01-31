package main

import (
	"fmt"
)

type Shape interface {
	getArea() float64
}

type Rectangle struct {
	sideLength float64
}

type Triangle struct {
	height float64
	base   float64
}

func (r Rectangle) getArea() float64 {
	return r.sideLength * r.sideLength
}

func (t Triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	var rectangle = Rectangle{sideLength: 20}
	var triangle = Triangle{height: 10, base: 10}

	printArea(rectangle)
	printArea(triangle)
}
