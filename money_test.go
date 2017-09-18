package cash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	m, err := New("$", 999)
	r.NoError(err)
	r.Equal("$9.99", m.String())
}

func Test_Unknown_Currency(t *testing.T) {
	r := require.New(t)
	_, err := New("XX", 999)
	r.Error(err)
}

func Test_String(t *testing.T) {
	r := require.New(t)

	table := []struct {
		Val int64
		Exp string
	}{
		{999, "$9.99"},
		{100, "$1.00"},
		{1, "$.01"},
		{123456789, "$1,234,567.89"},
		{100042, "$1,000.42"},
		{1000042, "$10,000.42"},
	}

	for _, tt := range table {
		m, err := New("$", tt.Val)
		r.NoError(err)
		r.Equal(tt.Exp, m.String())
	}
}

func Test_Amount(t *testing.T) {
	r := require.New(t)

	table := []struct {
		A int64
		N Numerator
		D Denominator
	}{
		{59, 0, 59},
		{100, 1, 0},
		{12345, 123, 45},
	}

	for _, tt := range table {
		m, err := New("$", tt.A)
		r.NoError(err)
		r.Equal(tt.N, m.Numerator)
		r.Equal(tt.D, m.Denominator)
		r.Equal(tt.A, m.Amount())
	}
}
