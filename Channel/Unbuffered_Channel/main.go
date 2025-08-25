package main

import "fmt"

func main() {
	ch := make(chan string) // unbuffered channel with no capacity

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}
