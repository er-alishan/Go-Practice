// Go program to illustrate the
// use of function
package main

import "fmt"

// area() is used to find the
// area of the rectangle
// area() function two parameters,
// i.e, length and width
func area(length, width int) int {

	Ar := length * width
	return Ar
}

// Main function
func main() {
	rect := area(10, 5)

	// Display the area of the rectangle
	// with method calling
	fmt.Printf("Area of rectangle is : %d", rect)
}
