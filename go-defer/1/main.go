package main

import (
	"fmt"
)

func main() {
	//defer work on Stack(LIFO)
	defer funOne()
	defer funTwo()

	defer func(Ananomous string) {
		fmt.Println("Ananomous function")
	}("")

	secondLine := "Hlo Alishan"
	fmt.Println("Hello World Brother")
	fmt.Println(secondLine)
}

func funOne() {
	fmt.Println("funOne")
}

func funTwo() {
	fmt.Println("funTwo")
}
