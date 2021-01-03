package intmap

func (m *IntMap) SymmetricDifference(t *IntMap) {
	r := make(map[int]struct{})

	for v := range m.m {
		_, k := t.m[v]
		if !k {
			r[v] = struct{}{}
		}
	}

	for v := range t.m {
		_, k := m.m[v]
		if !k {
			r[v] = struct{}{}
		}
	}

	m.m = r
}
