package main

import (
	"fmt"
)

func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	// developer := map[string]string{"name": "Alishan"} //it is map
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	// for idx, val := range developer {
	// 	fmt.Printf("%v\t%v\n", idx, val)
	// }   //it is map
}
