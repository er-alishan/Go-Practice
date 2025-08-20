package main

import (
	"fmt"
)

var i interface{} = "Golang"

func main() {
	s, ok := i.(string)
	if ok {
		fmt.Println("String Value:", s)
	} else {
		fmt.Println("Not a string")
	}
}
