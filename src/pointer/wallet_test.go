package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		// depositの中のBitcoinは型変換であり関数ではない
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

// テストの重複を解消してみよう！
func TestWallet2(t *testing.T) {

	//前段階の関数
	checkMoney := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		// 前処理
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		//ここで共通処理
		checkMoney(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		// 前処理
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		checkMoney(t, wallet, Bitcoin(10))
	})
}
