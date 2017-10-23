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
		{"$", "$0.00"},
		{"$.", "$0.00"},
	}

	for _, tt := range table {
		err := m.Scan(tt.Val)
		r.NoError(err)
		r.Equal(tt.Exp, m.String())
	}
}

func Test_Scan_Errors(t *testing.T) {
	r := require.New(t)

	m := Money{}

	table := []struct {
		Val string
		Exp string
	}{
		{"asdfasdf", "could not find currency for"},
		{"%100", "could not find currency for"},
		{"", "could not find currency for"},
	}

	for _, tt := range table {
		err := m.Scan(tt.Val)
		r.Error(err)
		r.Contains(err.Error(), tt.Exp)
	}
}

func Test_Value(t *testing.T) {
	r := require.New(t)

	table := []struct {
		M   Money
		Exp string
	}{
		{Money{Currency: USD, Numerator: 1, Denominator: 0}, "$1.00"},
		{Money{Currency: USD, Numerator: 123, Denominator: 45}, "$123.45"},
	}

	for _, tt := range table {
		v, err := tt.M.Value()
		r.NoError(err)
		r.Equal(tt.Exp, v.(string))
	}
}
