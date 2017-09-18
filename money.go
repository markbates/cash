package cash

import "github.com/pkg/errors"

type Money struct {
	Currency    Currency
	Numerator   Numerator
	Denominator Denominator
	amount      int64
}

func (m Money) String() string {
	return m.Currency.Format(m)
}

func New(symbol string, amount int64) (Money, error) {
	c, err := FindCurrency(symbol)
	if err != nil {
		return Money{}, errors.WithStack(err)
	}

	n := amount / 100
	d := amount % 100
	m := Money{
		Currency:    c,
		amount:      amount,
		Numerator:   Numerator(n),
		Denominator: Denominator(d),
	}

	return m, nil
}
