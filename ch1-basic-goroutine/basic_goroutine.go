//basic go routine
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Calling without goroutine")
	makeHttpCall("http://www.google.com")
	fmt.Println("Calling the goroutine")
	go makeHttpCall("http://www.google.com")

	// This prints:
	// Calling without goroutine
	// The link http://www.google.com is up
	// Calling the goroutine

	//the call with a goroutine doesnt print because the main gourotine (main function) finish first
}

func makeHttpCall(link string) {
	_, err := http.Get(link)
	if err == nil {
		fmt.Printf("The link %v is up \n", link)
	}
}
