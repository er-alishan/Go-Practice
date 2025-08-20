package main

import "fmt"

// Define interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct: Circle
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// Struct: Rectangle
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Function that accepts interface
func printShape(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 4, height: 6}

	printShape(c) // Circle implements Shape
	printShape(r) // Rectangle implements Shape
}
