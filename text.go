package cash

func (m Money) MarshalText() ([]byte, error) {
	s := m.String()
	return []byte(s), nil
}

func (m *Money) UnmarshalText(data []byte) error {
	return m.Scan(string(data))
}
