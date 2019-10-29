package ptrs

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, 10)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(10)
		if err == nil {
			t.Errorf("wanted an error but nothin was thrown")
		}
		assertBalance(t, wallet, 0)
	})
	t.Run("deposit and withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		err := wallet.Withdraw(10)
		if err != nil {
			t.Errorf("Should have enough money")
		}
		assertBalance(t, wallet, 0)
	})

}
