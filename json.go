package cash

func (m Money) MarshalJSON() ([]byte, error) {
	s := m.String()
	return []byte(s), nil
}

func (m *Money) UnmarshalJSON(data []byte) error {
	return m.Scan(string(data))
}
