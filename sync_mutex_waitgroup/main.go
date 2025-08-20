package main

import (
	"fmt"
	"sync"
)

type Account struct {
	name    string
	balance int
	rwmu    sync.RWMutex
}

// Read-only function
func (a *Account) Display() {
	a.rwmu.RLock()         // multiple readers allowed
	defer a.rwmu.RUnlock() // unlock after reading

	fmt.Printf("Account Holder: %s | Balance: %d\n", a.name, a.balance)
}

// Write: Withdraw
func (a *Account) Withdraw(amount int) {
	a.rwmu.Lock()         // exclusive lock for write
	defer a.rwmu.Unlock() // unlock after writing

	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdraw: %d | Balance: %d\n", amount, a.balance)
	} else {
		fmt.Println("Insufficient Balance:", a.balance)
	}
}

// Write: Deposit
func (a *Account) Deposit(amount int) {
	a.rwmu.Lock()
	defer a.rwmu.Unlock()

	a.balance += amount
	fmt.Printf("Deposit: %d | Balance: %d\n", amount, a.balance)
}

func main() {
	acc := &Account{name: "Md Alishan", balance: 2000}
	var wg sync.WaitGroup

	wg.Add(4)

	// Reader goroutine
	go func() {
		defer wg.Done()
		acc.Display()
	}()

	// Writer goroutines
	go func() {
		defer wg.Done()
		acc.Deposit(500)
	}()

	go func() {
		defer wg.Done()
		acc.Withdraw(1200)
	}()

	// Another reader
	go func() {
		defer wg.Done()
		acc.Display()
	}()

	wg.Wait()
	fmt.Println("Final Balance:", acc.balance)
}
