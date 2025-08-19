// In Go, an anonymous function is simply a function without a name.
package main

import "fmt"

func main() {
	// Anonymous function assigned to a variable
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}

	greet("Go Developer") // Calling it through variable
}
