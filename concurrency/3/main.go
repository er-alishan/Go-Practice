package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from Goroutine!")
}
func sendData(ch chan string) {
	ch <- "Hello from Channel!"
}

func main() {
	ch := make(chan string) // Channel for concurrency
	go hello()              // Runs concurrently
	go sendData(ch)
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second) // Wait for goroutine to finish
}
