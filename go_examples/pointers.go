
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


