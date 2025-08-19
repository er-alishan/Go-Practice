// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type BankAccount struct {
// 	balance int
// 	rwmu    sync.RWMutex
// }

// func (a *BankAccount) Deposit(amount int, wg *sync.WaitGroup, user string) {
// 	defer wg.Done()

// 	a.rwmu.Lock() // Writer lock (exclusive)
// 	fmt.Println(user, "depositing", amount)
// 	a.balance += amount
// 	a.rwmu.Unlock() // Unlock writer
// }

// func (a *BankAccount) GetBalance() int {
// 	a.rwmu.RLock() // Reader lock (multiple readers allowed)
// 	defer a.rwmu.RUnlock()
// 	return a.balance
// }

// func main() {
// 	var wg sync.WaitGroup
// 	account := &BankAccount{}

// 	// 4 users depositing
// 	users := []struct {
// 		name   string
// 		amount int
// 	}{
// 		{"User1", 100},
// 		{"User2", 200},
// 		{"User3", 300},
// 		{"User4", 400},
// 	}

// 	for _, u := range users {
// 		wg.Add(1)
// 		go account.Deposit(u.amount, &wg, u.name)
// 	}

// 	wg.Wait()
// 	fmt.Println("Final Balance:", account.GetBalance())
// }

package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	rwmu    sync.RWMutex
}

func (a *BankAccount) Deposit(amount int, wg *sync.WaitGroup, user string) {
	defer wg.Done()
	a.rwmu.Lock()
	a.balance += amount
	fmt.Println(user, "depositing", amount)
	a.rwmu.Unlock()
}
func (a *BankAccount) GetBalance() int {
	a.rwmu.RLock()
	defer a.rwmu.RUnlock()
	return a.balance
}
func main() {
	var wg sync.WaitGroup
	account := &BankAccount{}

	users := []struct {
		name   string
		amount int
	}{
		{"User1", 100},
		{"User2", 150},
		{"User3", 200},
		{"User4", 300},
	}

	for _, u := range users {
		wg.Add(1)
		go account.Deposit(u.amount, &wg, u.name)
	}
	wg.Wait()
	fmt.Println("Final Balance:", account.GetBalance())
}
