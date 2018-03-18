/*
 * Inspired by a great example at:
 https://gobyexample.com/interfaces
*/

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	GetArea() float64
	GetPerimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64 
}

func (c Circle) GetArea() float64 {
	cRadiusSq := c.radius * c.radius
	return math.Pi * cRadiusSq
}

func (c Circle) GetPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) GetArea() float64 {
	return r.width * r.height
}

func (r Rectangle) GetPerimeter() float64  {
	w := 2 * r.width
	h := 2 * r.height
	return w + h
}

func MeasureShape(g geometry) {
	fmt.Println("\n\n")
	fmt.Printf("Measuring: %v\n", g)
	fmt.Printf("Area: %v\n", g.GetArea())
	fmt.Printf("Perimeter: %v\n", g.GetPerimeter())
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 3, height: 4}

	MeasureShape(r)
	MeasureShape(c)
}

