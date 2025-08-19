package main

import (
	"fmt"
	"sync"
)

type Bankaccount struct {
	balance int
	mu      sync.Mutex
}

func (a *Bankaccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amount
	fmt.Println("Deposited:", amount, "Balance:", a.balance)
}
func (a *Bankaccount) Withdraw(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance >= amount {
		a.balance -= amount
		fmt.Println("Withdraw:", amount, "Balance:", a.balance)
	} else {
		fmt.Println("Insufficient Balance:", a.balance)
	}
}
func main() {
	account := &Bankaccount{balance: 100}
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(50)
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Withdraw(30)
		}()
	}
	wg.Wait()
	fmt.Println("Final Balance:", account.balance)
}
