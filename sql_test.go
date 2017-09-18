package cash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Scan(t *testing.T) {
	r := require.New(t)

	m := Money{}

	table := []struct {
		Val string
		Exp string
	}{
		{"$123456789", "$1,234,567.89"},
		{"$100", "$1.00"},
		{"$1.00", "$1.00"},
		{"$1,000.42", "$1,000.42"},
	}

	for _, tt := range table {
		err := m.Scan(tt.Val)
		r.NoError(err)
		r.Equal(tt.Exp, m.String())
	}
}

func Test_Value(t *testing.T) {
	r := require.New(t)

	table := []struct {
		M   Money
		Exp string
	}{
		{Money{Currency: USD, amount: 100}, "100$"},
		{Money{Currency: USD, amount: 12345}, "12345$"},
	}

	for _, tt := range table {
		v, err := tt.M.Value()
		r.NoError(err)
		r.Equal(tt.Exp, v.(string))
	}
}
