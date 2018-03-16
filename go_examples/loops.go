package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("EVERY NUMBER\n\n")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("ONLY THE EVENS: \n\n")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
