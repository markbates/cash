package cash

import (
	"database/sql/driver"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var symX = regexp.MustCompile("(\\D+)")

func (m *Money) Scan(src interface{}) error {
	s, ok := src.(string)
	if !ok {
		return errors.New("cannot scan non-string type")
	}

	s = strings.TrimSuffix(s, ".")
	mx := symX.FindString(s)
	if _, err := FindCurrency(mx); err != nil {
		return err
	}

	s = strings.TrimPrefix(s, mx)

	s = symX.ReplaceAllString(s, "")

	if s == "" || s == "." {
		s = "0"
	}

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

func (m Money) Value() (driver.Value, error) {
	return m.String(), nil
}
