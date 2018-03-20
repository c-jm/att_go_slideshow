/*

 Filename: scanner.go
 Authors: Kyle A. Kreutzer, Colin J. Mills
 Date: March 20th 2018

 */


package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// An example of defer and scanner
	FILE_NAME := "test.txt"
	handle, err := os.Open(FILE_NAME)
	defer handle.Close()

	if err != nil {
		fmt.Printf("File: %s doesn't exist\n", FILE_NAME)
	} else {
		scanner := bufio.NewScanner(handle)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
