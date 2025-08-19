package main

import "fmt"

func main() {
	// Declare a variable and assign it a value
	num := 10

	// Create a pointer to the variable
	ptr := &num

	// Access the value through the pointer
	fmt.Println(*ptr) // Output: 10

	// Modify the value through the pointer
	*ptr = 20

	// Access the modified value through the variable
	fmt.Println(num) // Output: 20
}

//<-----------Another Programm ------------->
// package main

// import "fmt"

// func main() {
// 	num := 10
// 	ptr := &num // ptr is a pointer to num

// 	fmt.Println(ptr)  // prints the memory address of num
// 	fmt.Println(*ptr) // prints the value of num, which is 10
// }

//<-----------Another Programm ------------->
// package main

// import "fmt"

// func main() {
//     var a int = 42
//     var p *int = &a // p is a pointer to the memory address of a

//     fmt.Println("Value of a:", a)       // Output: Value of a: 42
//     fmt.Println("Address of a:", &a)    // Output: Address of a: 0xc0000140a0 (example address)
//     fmt.Println("Value of p:", p)       // Output: Value of p: 0xc0000140a0 (same address as &a)
//     fmt.Println("Value at p:", *p)      // Output: Value at p: 42

//     *p = 21 // Change the value at the memory address p points to
//     fmt.Println("New value of a:", a)   // Output: New value of a: 21
// }
