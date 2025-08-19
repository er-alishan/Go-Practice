package main

import "fmt"

func addnumbers() {
	var number1 int
	var number2 int
	fmt.Print("Enter number1:")
	fmt.Scan(&number1)
	fmt.Print("Enter number2:")
	fmt.Scan(&number2)

	sum := number1 + number2

	fmt.Println("Sum:", sum)
}
func main() {
	addnumbers()
}
