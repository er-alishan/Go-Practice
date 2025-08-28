package main

import (
	"fmt"
	"time"
)

func worker(id int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(time.Duration(id) * 200 * time.Millisecond)
			ch <- fmt.Sprintf("Worker %d -> task %d", id, i)
		}
		close(ch)
	}()
	return ch
}

// Fan-in: merge two channels into one
func fanIn(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for v := range ch1 {
			out <- v
		}
	}()
	go func() {
		for v := range ch2 {
			out <- v
		}
	}()
	return out
}

func main() {
	c1 := worker(1)
	c2 := worker(2)

	merged := fanIn(c1, c2)

	// Read from merged channel
	for i := 0; i < 6; i++ {
		fmt.Println(<-merged)
	}
}
