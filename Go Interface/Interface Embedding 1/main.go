package main

import "fmt"

// Depositor interface (sirf deposit kar sakta hai)
type Depositor interface {
	Deposit(amount int)
}

// Withdrawer interface (sirf withdraw kar sakta hai)
type Withdrawer interface {
	Withdraw(amount int)
}

// AccountOperations embeds Depositor + Withdrawer
type AccountOperations interface {
	Depositor
	Withdrawer
}

// Struct: BankAccount
type BankAccount struct {
	holder  string
	balance int
}

// Deposit method
func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Printf("%s deposited %d. Current Balance: %d\n", b.holder, amount, b.balance)
}

// Withdraw method
func (b *BankAccount) Withdraw(amount int) {
	if amount <= b.balance {
		b.balance -= amount
		fmt.Printf("%s withdrew %d. Current Balance: %d\n", b.holder, amount, b.balance)
	} else {
		fmt.Printf("%s has insufficient balance! Current Balance: %d\n", b.holder, b.balance)
	}
}

// Function that accepts AccountOperations
func performOperations(acc AccountOperations) {
	acc.Deposit(1000)
	acc.Withdraw(500)
	acc.Withdraw(700)
}

func main() {
	account := &BankAccount{holder: "Md Alishan", balance: 500}
	performOperations(account)
}
