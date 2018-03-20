/*

 Filename: loops.go
 Authors: Kyle A. Kreutzer, Colin J. Mills
 Date: March 20th 2018

 */


package main

import (
	"fmt"
)

func main() {

	fmt.Println("EVERY NUMBER\n\n")

	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}

	fmt.Println("ONLY THE EVENS: \n\n")
	for i := 0; i < 10; i++ {

		if i % 2 == 0 {
			fmt.Println(i)
		}
        
	}
}
