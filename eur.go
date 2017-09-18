package cash

import "fmt"

var EUR = Currency{
	Name:   "EUR",
	Symbol: "€",
	Format: func(m Money) string {
		return fmt.Sprintf("%s%s,%s", "€", m.Numerator.Format("."), m.Denominator)
	},
}
