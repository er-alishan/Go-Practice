package main

import "fmt"

// Interface
type Account interface {
	GetName() string
}

// Struct
type BankAccount struct {
	holder string
}

func (b BankAccount) GetName() string {
	return b.holder
}

func main() {
	acc := Account(BankAccount{"Md Alishan"})

	// ✅ Type assertion: get the concrete type (BankAccount)
	ba, ok := acc.(BankAccount)
	if ok {
		fmt.Println("Type assertion successful!")
		fmt.Println("Account holder:", ba.holder) // Access struct field
	} else {
		fmt.Println("Type assertion failed")
	}
}
