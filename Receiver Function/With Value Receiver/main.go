package main

import (
	"fmt"
)

type Rectangle struct {
	lenght int
	width  int
}

// Value Receiver (copy)
func (r Rectangle) Area() int {
	return r.lenght * r.width
}
func main() {
	rect := Rectangle{lenght: 5, width: 10}

	fmt.Println("Rectangle:", rect)

	fmt.Println("Area:", rect.Area())

}
