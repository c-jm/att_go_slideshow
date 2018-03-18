

package main

import "time"
import "fmt"
import "net/http"

//import "github.com/veandco/go-sdl2/sdl"

func makeRequest(value string) {
	_, err := http.Get(value)
	if err != nil {
		fmt.Println(err)
	}
}

func makeRequestsAndTimeNormal() {
	timeBefore := time.Now()
	
	makeRequest("http://www.google.com")
	makeRequest("http://www.facebook.com")
	makeRequest("http://www.twitter.com")

	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timeBefore)
	fmt.Printf("From non go routine: %v\n", elapsed)
}

func makeRequestsWithGoRoutines() {
	timeBefore := time.Now()

	go makeRequest("http://www.google.com")
	go makeRequest("http://www.facebook.com")
	go makeRequest("http://www.twitter.com")

	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timeBefore)
	fmt.Printf("From go routine: %v\n", elapsed)
}

func main() {
	makeRequestsAndTimeNormal()
	makeRequestsWithGoRoutines()
}
