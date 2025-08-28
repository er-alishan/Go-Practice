package main

import "fmt"

func generator(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
func main() {
	ch := generator(1, 5)
	for val := range ch {
		fmt.Println("Received:", val)
	}
}
