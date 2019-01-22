package cash

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (m Money) MarshalJSON() ([]byte, error) {
	s := m.String()
	return json.Marshal(s)
}

func (m *Money) UnmarshalJSON(data []byte) error {
	d := string(data)
	if err := m.Scan(d); err != nil {
		if errors.Cause(err) == ErrCurrencyNotFound {
			return m.Scan("$" + d)
		}
	}
	return nil
}
