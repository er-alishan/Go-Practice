package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	rwmu    sync.RWMutex
}

func (a *BankAccount) Deposit(amount int) {
	a.rwmu.Lock() // write lock
	defer a.rwmu.Unlock()
	a.balance += amount
	fmt.Println("Deposited:", amount, "Balance:", a.balance)
}

func (a *BankAccount) GetBalance() int {
	a.rwmu.RLock() // read lock
	defer a.rwmu.RUnlock()
	return a.balance
}

func main() {
	account := &BankAccount{balance: 100}
	var wg sync.WaitGroup

	// Writer goroutines (deposit)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(50)
		}()
	}

	// Reader goroutines (get balance)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Reader %d sees balance: %d\n", id, account.GetBalance())
			time.Sleep(500 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("Final Balance:", account.GetBalance())
}
