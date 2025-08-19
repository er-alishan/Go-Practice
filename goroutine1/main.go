package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	go sayHello() // Create a goroutine to execute the sayHello function

	fmt.Println("Main function")

	time.Sleep(1 * time.Second) // Wait for the goroutine to complete
}
