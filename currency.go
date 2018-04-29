package cash

import (
	"sync"

	"github.com/pkg/errors"
)

var ErrCurrencyNotFound = errors.New("currency not found")
var curmoot = &sync.Mutex{}

var currencies = map[string]Currency{
	"$": USD,
	"€": EUR,
	"£": GBP,
}

func AddCurrency(c Currency) {
	curmoot.Lock()
	defer curmoot.Unlock()
	currencies[c.Symbol] = c
}

func FindCurrency(symbol string) (Currency, error) {
	curmoot.Lock()
	defer curmoot.Unlock()
	c, ok := currencies[symbol]
	if !ok {
		return c, errors.Wrapf(ErrCurrencyNotFound, "%s", symbol)
	}
	return c, nil
}

type Formatter func(Money) string

type Currency struct {
	Name   string
	Symbol string
	Format Formatter
}
