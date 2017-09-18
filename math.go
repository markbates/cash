package cash

import "github.com/pkg/errors"

func (a Money) Add(monies ...Money) (Money, error) {
	amount := a.Amount()
	for _, m := range monies {
		if a.Currency.Symbol != m.Currency.Symbol {
			return a, errors.Errorf("can not add two different currencies: %s, %s", a.Currency.Symbol, m.Currency.Symbol)
		}
		amount += m.Amount()
	}
	return New(a.Currency.Symbol, amount)
}

func (a Money) Sub(monies ...Money) (Money, error) {
	amount := a.Amount()
	for _, m := range monies {
		if a.Currency.Symbol != m.Currency.Symbol {
			return a, errors.Errorf("can not substract two different currencies: %s, %s", a.Currency.Symbol, m.Currency.Symbol)
		}
		amount -= m.Amount()
	}
	return New(a.Currency.Symbol, amount)
}

func (a Money) Multiply(m int64) (Money, error) {
	amount := a.Amount() * m
	return New(a.Currency.Symbol, amount)
}

func (a Money) Divide(m int64) (Money, error) {
	amount := a.Amount() / m
	return New(a.Currency.Symbol, amount)
}
