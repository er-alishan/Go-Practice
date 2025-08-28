package main

import "fmt"

type Server struct {
	requests chan string
	quit     chan bool
}

func (s *Server) Start() {
	go func() {
		for {
			select {
			case req := <-s.requests:
				fmt.Println("Processing:", req)
			case <-s.quit:
				fmt.Println("Server shutting down...")
				return
			}
		}
	}()
}

func main() {
	server := &Server{
		requests: make(chan string),
		quit:     make(chan bool),
	}

	server.Start()

	server.requests <- "Request 1"
	server.requests <- "Request 2"

	server.quit <- true // gracefully stop server
}
