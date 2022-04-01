package main

import (
	"fmt"
	"net/http"
)

func main() {
	// creating a channel to share string type data
	myChanel := make(chan string)

	// creating a slice of links
	links := []string{
		"http://abc.com",
		"http://pqr.com",
		"http://xyz.com",
	}
	// fetching data from each link
	for _, link := range links {
		go makeHttpCall(link, myChanel)
	}
	// listening for three messages coming from the chanel
	for i := 0; i < 3; i++ {
		fmt.Printf("Link %v is up \n", <-myChanel)
	}

	// Link http://abc.com is up
	// Link http://xyz.com is up
	// Link http://pqr.com is up
}

func makeHttpCall(link string, myChanel chan string) {
	_, err := http.Get(link)
	if err == nil {
		// sending the link name to the chanel 👈
		myChanel <- link
	}
}
