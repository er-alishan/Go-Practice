package main

import "fmt"

type logger interface {
	Log(message string)
	Error(message string)
}

type fileLogger struct {
	filename string
}

func (f fileLogger) Log(message string) {
	fmt.Printf("Logging to file %s: %s\n", f.filename, message)
}
func (f fileLogger) Error(message string) {
	fmt.Printf("Logging to file %s: %s\n", f.filename, message)
}

type consoleLogger struct{} //empty struct

func (c consoleLogger) Log(message string) {
	fmt.Printf("Logging to console: %s\n", message)
}
func (c consoleLogger) Error(message string) {
	fmt.Printf("Logging to console: %s", message)
}

func main() {

	l := fileLogger{"log.txt"}
	l.Log("Hello")
	l.Error("Error")

	l2 := consoleLogger{}
	l2.Log("Hello")
	l2.Error("Error")
}
