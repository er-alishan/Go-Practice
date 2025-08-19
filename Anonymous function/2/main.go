package main

import "fmt"

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("Hello from anonymous function!")
}
