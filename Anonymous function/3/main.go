// with return values
package main

import "fmt"

func main() {
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println(sum(5, 3)) // 8
}
