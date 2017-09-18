package cash

import "fmt"

var GBP = Currency{
	Name:   "GBP",
	Symbol: "£",
	Format: func(m Money) string {
		return fmt.Sprintf("%s%s.%s", "£", m.Numerator.Format(","), m.Denominator)
	},
}
