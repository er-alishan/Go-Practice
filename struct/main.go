package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
	}
	var p1 person
	p1 = person{
		name: "Md Alishan",
		age:  18,
	}
	p2 := person{
		name: "Md Imran",
		age:  40,
	}
	p3 := person{
		name: "Md Farhan",
		age:  30,
	}
	p4 := person{
		name: "Md Fakhre",
		age:  30,
	}
	fmt.Println(p1, p2, p3, p4)
	//<----------------------------------------------------->
	type rectangle struct {
		lenght  int
		breadth int
	}

	var r1 rectangle
	r1 = rectangle{
		lenght:  10,
		breadth: 20,
	}
	fmt.Println("lenght:", r1.lenght)
	fmt.Println("breadth:", r1.breadth)

	area := r1.lenght * r1.breadth
	fmt.Println("area:", area)
}
