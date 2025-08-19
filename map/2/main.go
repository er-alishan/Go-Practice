package main

import "fmt"

func main() {
	var a = make(map[string]int)
	a["Alishan"] = 18
	a["Imran"] = 40
	a["Farhan"] = 30
	for k, v := range a {
		fmt.Printf("%v: %v\n", k, v)
	}

	//update value of map
	a["Alishan"] = 20
	fmt.Println(a["Alishan"])

	//delete value of map
	delete(a, "Alishan")
	fmt.Println(a["Alishan"])
}
