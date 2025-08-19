package main

import "fmt"

func main() {
	// Declare a map with string keys and int values
	myMap := make(map[string]int)

	// Add some key-value pairs
	myMap["apple"] = 5
	myMap["banana"] = 10

	// Retrieve a value
	fmt.Println(myMap["apple"]) // Output: 5

	// Delete a key-value pair
	delete(myMap, "banana")

	// Print the remaining key-value pair
	fmt.Println(myMap) // Output: map[apple:5]
}
