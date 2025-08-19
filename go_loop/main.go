package main

import "fmt"

func main() {
	person := [2]string{"Alishan", "Imran"}
	role := [2]string{"Student", "Engineer"}

	for i := 0; i < len(person); i++ {
		for j := 0; j < len(role); j++ {
			if person[i] == "Alishan" && role[j] == "Engineer" {
				continue
			}
			if person[i] == "Imran" && role[j] == "Student" {
				continue
			}
			fmt.Println(person[i], role[j])
		}
	}
}
