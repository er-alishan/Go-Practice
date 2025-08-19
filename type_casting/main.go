package main

import "fmt"

func main() {

	// float to int
	var floatValue float64 = 10.9

	var intValue int = int(floatValue)

	fmt.Printf("floatValue:%v\n", floatValue)
	fmt.Printf("intValue:%d", intValue)

	//int to float
	var intValue2 int = 100
	var floatValue2 float64 = float64(intValue2)

	fmt.Printf("\nintValue2:%d\n", intValue2)
	fmt.Printf("floatValue2:%.1f", floatValue2)
	fmt.Printf("\nType of floatValue2:%T", floatValue2)

	//Add int and float values in go type casting
	var num1 int = 10
	var num2 float64 = 10.5
	sum1 := num2 + float64(num1)
	sum2 := num1 + int(num2)
	fmt.Printf("\nSum is:%.1f", sum1)
	fmt.Printf("\nSum is:%d", sum2)
}
