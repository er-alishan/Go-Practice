package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s:%v\n", url, err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body from %s:%v\n", url, err)
		return
	}
	fmt.Printf("Fetched %s: (%d byte) in %v\n", url, len(body), time.Since(start))
}
func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://qlearning.in",
		"https://golang.org",
	}
	wg.Add(len(urls))
	for _, url := range urls {
		go fetchURL(url, &wg)
	}
	wg.Wait()
	fmt.Println("All fetch operaton completed!!")
}
