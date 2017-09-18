package cash

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var symX = regexp.MustCompile("(\\D+)")

func (m *Money) Scan(src interface{}) error {
	s, ok := src.(string)
	if !ok {
		return errors.New("can't scan non-string type")
	}

	mx := symX.FindString(s)
	if mx == "" {
		return errors.Errorf("couldn't find currency symbol: %s", s)
	}

	s = strings.TrimPrefix(s, mx)

	s = symX.ReplaceAllString(s, "")

	i, err := strconv.Atoi(s)
	if err != nil {
		return errors.WithStack(err)
	}

	mm, err := New(mx, int64(i))
	if err != nil {
		return errors.WithStack(err)
	}

	(*m) = mm

	return nil
}

func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("%d%s", m.amount, m.Currency.Symbol), nil
}
