package main

import "fmt"

// Reader interface
type Reader interface {
	Read() string
}

// Writer interface
type Writer interface {
	Write(data string)
}

// ReadWriter interface embeds Reader + Writer
type ReadWriter interface {
	Reader
	Writer
}

// Struct: File implements both Reader & Writer
type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

// Function that accepts ReadWriter
func process(rw ReadWriter) {
	rw.Write("Hello, Go Interface Embedding!")
	fmt.Println("Reading:", rw.Read())
}

func main() {
	f := &File{} // pointer receiver ke liye &File
	process(f)
}
