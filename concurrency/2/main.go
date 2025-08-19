package main

import (
	"fmt"
	"time"
)

// Dummy function jo driver data fetch karega
func findDrivers(area string, ch chan string) {
	time.Sleep(time.Second) // Simulating API call delay
	ch <- fmt.Sprintf("Drivers found in %s area", area)
}

func main() {
	areas := []string{"Downtown", "Airport", "Suburbs"}
	ch := make(chan string) // Channel for concurrency

	for _, area := range areas {
		go findDrivers(area, ch) // Goroutines se multiple requests parallelly execute hongi
	}

	for i := 0; i < len(areas); i++ {
		fmt.Println(<-ch) // Responses handle karna
	}
}
