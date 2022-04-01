// Buffered channels are used to perform asynchronous communication within the goroutines.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	// initializing a WaitGroup
	var wg sync.WaitGroup

	// adding 3 counts/buffer to the WaitGroup
	wg.Add(3)

	fmt.Println("Start Goroutines")
	go responseSize("https://www.golangprograms.com", &wg)
	go responseSize("https://stackoverflow.com", &wg)
	go responseSize("https://coderwall.com", &wg)

	// wait for goroutines to finish
	wg.Wait()
	fmt.Println("Terminating the main program")

	// Output:
	// Start Goroutines
	// Step1:  https://coderwall.com
	// Step1:  https://stackoverflow.com
	// Step1:  https://www.golangprograms.com
	// Step2:  https://coderwall.com
	// Step2:  https://stackoverflow.com
	// Step3:  189752
	// Step2:  https://www.golangprograms.com
	// Step3:  31672
	// Step3:  188813
	// Terminating the main program
}

// just prints the response size of the body returned
func responseSize(url string, wg *sync.WaitGroup) {
	// schedule the Done() call when the goroutine is finished
	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step3: ", len(body))
}
