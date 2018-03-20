/*

 Filename: pointers.go
 Authors: Kyle A. Kreutzer, Colin J. Mills
 Date: March 20th 2018

 */


package main

import (
    "fmt"
)

func main() {
    theMeaningOfLife := 42

    mutableMeaning := &theMeaningOfLife

    fmt.Println(*mutableMeaning)

    *mutableMeaning = 100;

    fmt.Println(*mutableMeaning)

}


