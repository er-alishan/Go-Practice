package main

import (
	"fmt"
)

type Account struct {
	name    string
	balance int
}

// Pointer Receiver (changes reflect hote hain)
func (a *Account) deposit(amount int) {
	a.balance += amount
	fmt.Printf("Deposit Amount:%d\n", amount)
	fmt.Println("Deposit Successfull And Current Balance:", a.balance)
}
func main() {
	acc := Account{
		name:    "Alishan",
		balance: 1000,
	}
	fmt.Printf("Account Holder Name:%s\nCurrent Balance:%d\n", acc.name, acc.balance)

	acc.deposit(500) // balance update ho jayega
}
