package service3

// Account ...
type Account struct {
	ID      int64 `db:"aikawarazu"`
	Balance int   `db:"tekitode"`
}

// IsSufficient ...
func (a *Account) IsSufficient(amount int) bool {
	return a.Balance >= amount
}

// Transfer ...
func (a *Account) Transfer(amount int, to *Account) {
	a.Balance -= amount
	to.Balance += amount
}
