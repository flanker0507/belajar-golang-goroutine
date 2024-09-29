package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter Ke: ", x)
}

type BankAccount struct {
	mutex   sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.mutex.Lock()
	account.Balance = amount + account.Balance
	account.mutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.mutex.Lock()
	balance := account.Balance
	account.mutex.Unlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance: ", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(Amount int) {
	user.Balance = user.Balance + Amount
}

func Transfer(User1 *UserBalance, User2 *UserBalance, Amount int) {
	User1.Lock()
	fmt.Println("Lock User 1 ", User1.Name)
	User1.Change(-Amount)

	time.Sleep(1 * time.Second)

	User2.Lock()
	fmt.Println("Lock User 2 ", User2.Name)
	User2.Change(Amount)

	time.Sleep(1 * time.Second)

	User1.Unlock()
	User2.Unlock()
}

func TestDeadLock(t *testing.T) {
	User1 := UserBalance{
		Name:    "Yuda",
		Balance: 1000000,
	}

	User2 := UserBalance{
		Name:    "Dilla",
		Balance: 1000000,
	}

	go Transfer(&User1, &User2, 100000)
	go Transfer(&User2, &User1, 200000)

	time.Sleep(2 * time.Second)

	fmt.Println("User ", User1.Name, " Balance ", User1.Balance)
	fmt.Println("User ", User2.Name, " Balance ", User2.Balance)
}
