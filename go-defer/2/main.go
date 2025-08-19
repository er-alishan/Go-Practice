package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.text")
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}
	defer file.Close() // Alls are complete then defer will execute and properly close
	file.WriteString("Ya, golang is awesome")
	fmt.Println("File opened successfully")
}
