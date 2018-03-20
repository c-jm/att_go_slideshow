/*

 Filename: goroutines.go
 Authors: Kyle A. Kreutzer, Colin J. Mills
 Date: March 20th 2018

 */

package main

import "time"
import "fmt"
import "net/http"

func makeRequest(url string) {
	_, err := http.Get(url)
	if err != nil {
        panic(err)
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
