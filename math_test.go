package cash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Add(t *testing.T) {
	r := require.New(t)

	a, err := New("$", 999)
	r.NoError(err)

	b, err := New("$", 199)
	r.NoError(err)

	c, err := a.Add(b)
	r.NoError(err)
	r.Equal("$11.98", c.String())
}

func Test_Sub(t *testing.T) {
	r := require.New(t)

	a, err := New("$", 999)
	r.NoError(err)

	b, err := New("$", 199)
	r.NoError(err)

	c, err := a.Sub(b)
	r.NoError(err)
	r.Equal("$8.00", c.String())

	d, err := c.Sub(a)
	r.NoError(err)
	r.Equal("$-1.99", d.String())
}

func Test_Multiply(t *testing.T) {
	r := require.New(t)

	a, err := New("$", 999)
	r.NoError(err)

	c, err := a.Multiply(5)
	r.NoError(err)
	r.Equal("$49.95", c.String())
}

func Test_Divide(t *testing.T) {
	r := require.New(t)

	a, err := New("$", 999)
	r.NoError(err)

	c, err := a.Divide(4)
	r.NoError(err)
	r.Equal("$2.49", c.String())
}
