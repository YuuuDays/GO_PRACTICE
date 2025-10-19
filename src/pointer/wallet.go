package pointer

type Wallet struct {
	money int
}

func (w Wallet) Deposit(x int) {
	w.money = x
}

func (w Wallet) Balance() int {
	return w.money
}
