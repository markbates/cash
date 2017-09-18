package cash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Numerator_Format(t *testing.T) {
	r := require.New(t)

	table := []struct {
		Num Numerator
		Exp string
	}{
		{123456789, "123,456,789"},
		{1234, "1,234"},
		{1, "1"},
		{10, "10"},
		{100, "100"},
		{1000, "1,000"},
		{10000, "10,000"},
	}

	for _, tt := range table {
		fmt.Println("----------")
		r.Equal(tt.Exp, tt.Num.Format(","))
	}
}
