// // Program to run two goroutines concurrently

// package main

// import (
// 	"fmt"
// 	"time"
// )

// // create a function
// func display(message string) {

// 	// infinite for loop
// 	for {
// 		fmt.Println(message)

// 		// to sleep the process for 1 second
// 		time.Sleep(time.Second * 1)
// 		return
// 	}
// }

// func main() {

// 	// call function with goroutine
// 	go display("Process 1")

// 	// call function without goroutine
// 	display("Process 2")

// }

//<------------------------------------Other Program------------------------------>

// Program to illustrate multiple goroutines
package main

import (
	"fmt"
	"time"
)

func display(message string) {

	fmt.Println(message)

}

func main() {

	// run two different goroutine
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")
	display("Process 4")
	// to sleep main goroutine for 1 sec
	time.Sleep(time.Second * 1)
}
