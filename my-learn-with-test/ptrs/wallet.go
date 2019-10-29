package ptrs

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance = w.balance + amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	w.balance -= amount
	return nil
}
