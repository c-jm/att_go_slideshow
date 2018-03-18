package main

import (
	"fmt"
	"time"
)

func main() {
	SayWhatTimeItIs()
	r := SqrRange(10)

	for i := 0; i < len(r); i++ {
		x := r[i]
		fmt.Println(x)
	}
}

// This switch example was taken from a tour of go:
// https://tour.golang.org/flowcontrol/11
func SayWhatTimeItIs() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func SqrRange(r int) []int {

	var result []int

	// This is a slice, its  basically a "look" into an array. You can use slices to build dynamic arrays.

	result = make([]int, r)

	for i := 0; i < r; i++ {
		x := i + 1
		result[i] = x * x
	}

	return result
}
