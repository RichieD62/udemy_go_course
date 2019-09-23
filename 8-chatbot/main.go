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
	tri := triangle{
		height: 3.2,
		base:   5.3,
	}
	sqr := square{
		sideLength: 3.6,
	}

	printArea(tri)
	printArea(sqr)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (squ square) getArea() float64 {
	area := squ.sideLength * squ.sideLength
	return area
}
