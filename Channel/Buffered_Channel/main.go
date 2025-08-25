package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Hello buffered 1"
	ch <- "Hello buffered 2"
	//ch <- "Hello buffered 3" // this is not send and error because the capacity of bufferd channel is 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
