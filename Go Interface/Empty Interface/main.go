package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Value:%v, Types:%T\n", i, i)
}
func main() {
	describe(25)
	describe("Golang")
	describe(3.14)
}
