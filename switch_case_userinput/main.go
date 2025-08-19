package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter x: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	x, _ := strconv.Atoi(in.Text())

	switch x {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("unknown")
	}
}
