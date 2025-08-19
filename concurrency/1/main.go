package main

import (
	"fmt"
	"sync"
	"time"
)

func greet(pharse string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello!", pharse)
}
func Slowgreet(pharse string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", pharse)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go greet("Nice to meet you!..", &wg)
	wg.Add(1)
	go greet("How are you?", &wg)
	wg.Add(1)
	go Slowgreet("How...are....you...?", &wg)
	wg.Add(1)
	go greet("I hope you are liking the course...", &wg)
	wg.Wait()
	fmt.Println("All goroutines finished executing.")
}
