package main

import (
	"fmt"
)

type Account struct {
	name    string
	balance int
}

func (a Account) Display() {
	fmt.Printf("Account Holder:%s\nBalance:%d\n", a.name, a.balance)
}

func (a *Account) Withdraw(amount int) {
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdraw Amount:%d\n", amount)
		fmt.Println("Withdraw Successfull And Current Balance:", a.balance)
	} else {
		fmt.Println("Insufficient Balance:", a.balance)
	}
}
func main() {
	acc := Account{
		name:    "Md Alishan",
		balance: 2000,
	}
	acc.Display()
	acc.Withdraw(1000)
}
