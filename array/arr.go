package main

import (
	"fmt"
)

func main() {
	var arr1 = [4]int{0, 1, 2, 3}
	var arr2 = [...]string{"Alishan", "Imran", "Farhan", "Fakhre"}
	arr3 := [5]float64{0.1, 0.2, 0.3, 0.4, 0.5}
	arr4 := [...]string{"Sohail", "Adil", "Afzal", "Aftab", "Tauqir"}

	arr1[0] = 1

	fmt.Println(arr1)
	fmt.Println(arr2[1])
	fmt.Println(arr3)
	fmt.Println(len(arr4))
}
