/*
package main

import (

	"fmt"
	"os"

)

	func main() {
		// Try to open a non-existent file
		file, err := os.Open("file.txt")
		if err != nil {
			// If there is an error, print it and exit
			fmt.Println("Error occurred:", err)
			return
		}
		defer file.Close()

		fmt.Println("File opened successfully")
	}
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s", r.URL.Path) // Log incoming request URL path
	err := processRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func processRequest(r *http.Request) error {
	if r.URL.Path != "/" {
		return &MyError{
			time.Now(),
			"Invalid URL path: " + r.URL.Path, // Include the invalid path in the error message
		}
	}
	return nil
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
