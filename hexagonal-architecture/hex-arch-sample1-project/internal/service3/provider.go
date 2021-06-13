package service3

import (
	"context"
	"fmt"
)

// Provider ...
type Provider struct {
	r Repository
}

// NewProvider ...
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

// OpenAccount ...
func (p *Provider) OpenAccount(ctx context.Context, initialAmount int) (Account, error) {
	if initialAmount <= 0 {
		return Account{}, fmt.Errorf("provider: initial ammount must be greater than 0")
	}

	account, err := p.r.OpenAccount(ctx, initialAmount)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

// Transfer ...
func (p *Provider) Transfer(ctx context.Context, amount int, fromID, toID int64) (from, to Account, err error) {
	if fromID == toID {
		return Account{}, Account{}, fmt.Errorf("provider: cannot transfer money to oneself")
	}

	type Accounts struct {
		from Account
		to   Account
	}

	txFn := func(ctx context.Context) (interface{}, error) {
		from, to, err := p.r.GetAccountsForTransfer(ctx, fromID, toID)
		if err != nil {
			return Accounts{}, err
		}

		if !from.IsSufficient(amount) {
			return Account{}, fmt.Errorf("provider: balance is not sufficient - accountID: %d", from.ID)
		}

		from.Transfer(amount, &to)

		from, err = p.r.UpdateBalance(ctx, from)
		if err != nil {
			return Accounts{}, err
		}

		to, err = p.r.UpdateBalance(ctx, to)
		if err != nil {
			return Accounts{}, err
		}

		return Accounts{from: from, to: to}, nil
	}

	v, err := p.r.RunInTransaction(ctx, txFn)
	if err != nil {
		return Account{}, Account{}, err
	}

	val, ok := v.(Accounts)
	if !ok {
		return Account{}, Account{}, fmt.Errorf("provider: an error occurs - transfer")
	}

	return val.from, val.to, nil
}
