package account

import "sync"

// Define the Account type here.

type Account struct {
	mu      sync.Mutex
	Opened  bool
	balance int64
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		Opened:  true,
		balance: amount,
	}
}

func (a *Account) Balance() (int64, bool) {
	if a.Opened {
		return a.balance, true
	}
	return 0, false
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.Opened || (a.balance+amount < 0) {
		return 0, false
	}
	a.balance = a.balance + amount

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.Opened {

		a.Opened = false
		returnedPayout := a.balance
		a.balance = 0

		return returnedPayout, true

	}
	return 0, false
}
