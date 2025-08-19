package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func calculateShapeProperties(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	rect := Rectangle{width: 5, height: 10}
	circle := Circle{radius: 3}

	calculateShapeProperties(rect)
	calculateShapeProperties(circle)
}

// In this example, we define two types Rectangle and Circle that implement the Shape interface. The Shape interface defines the Area() and Perimeter() methods.

// We then define a function calculateShapeProperties() that takes a Shape interface as an argument and prints the area and perimeter of the shape.

// In the main() function, we create instances of Rectangle and Circle and pass them to the calculateShapeProperties() function. The function can work with both types because they implement the Shape interface.

// When you run this program, it will output the area and perimeter for both the rectangle and the circle:
