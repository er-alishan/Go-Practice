// package main

// import "fmt"

// func main() {

// 	fmt.Println(test())

// }

// func test() int {

// 	result := 1

// 	defer func() { result++ }()

// 	return result

// }

package main

import "fmt"

func main() {

	ch := make(chan int, 4)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)
	}()
	for num := range ch {
		fmt.Println(num)

	}

}

// arr:=[]int{1,2,4,5,6,3,8,4}, target:=8

// [1,2,5]
// [1,4,3]
// [1,3,4]
// [2,6]
// [4,4]
// [5,3]
// [8]
