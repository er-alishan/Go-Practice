// working with channels as function parameters
// main() ──go──> worker(1)
//       └─go──> worker(2)
//       └─go──> worker(3)

// worker → sends msg → channel → main() receives

package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Millisecond * 500)
		ch <- fmt.Sprintf("Worker %d: task %d done", id, i)
	}
}
func main() {
	ch := make(chan string)

	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	for i := 1; i <= 9; i++ {
		fmt.Println(<-ch)
	}
}
