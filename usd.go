package cash

import "fmt"

var USD = Currency{
	Name:   "USD",
	Symbol: "$",
	Format: func(m Money) string {
		return fmt.Sprintf("%s%s.%s", "$", m.Numerator.Format(","), m.Denominator)
	},
}
