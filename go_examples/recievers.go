/*

 Filename: recievers.go
 Authors: Kyle A. Kreutzer, Colin J. Mills
 Date: March 20th 2018

 */


package main
import (
	"fmt"
	"math"
)

type Vec2 struct  {
	x, y float32
}

func (v Vec2) Length() float64 {
	x2 := float64(v.x*v.x)
	y2 := float64(v.y*v.y)

	return math.Sqrt(x2 + y2)
}

func (v* Vec2) Add(toAdd Vec2) {
	v.x += toAdd.x
	v.y += toAdd.y
}

func main() {
	v := Vec2{3, 4}
    fmt.Printf("Length: %v\n\n", v.Length())

	v.Add(Vec2{3, 2})
	fmt.Printf("New Vector %v, %v\n\n", v.x, v.y)
}
