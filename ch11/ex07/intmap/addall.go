package intmap

func (m *IntMap) AddAll(entries ...int) {
	for _, entry := range entries {
		m.Add(entry)
	}
}
