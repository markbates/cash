package cash

import (
	"sync"

	"github.com/pkg/errors"
)

var curmoot = &sync.Mutex{}

var currencies = map[string]Currency{
	"$": USD,
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
		return c, errors.Errorf("could not find currency for %s", symbol)
	}
	return c, nil
}

type Formatter func(Money) string

type Currency struct {
	Name   string
	Symbol string
	Format Formatter
}
