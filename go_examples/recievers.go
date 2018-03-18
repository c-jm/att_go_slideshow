
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

	v.Add(Vec2{3, 2})

	fmt.Printf("%v, %v\n", v.x, v.y)
}
