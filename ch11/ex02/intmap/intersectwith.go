package intmap

func (m *IntMap) IntersectWith(t *IntMap) {
	r := make(map[int]struct{})

	for v := range m.m {
		_, k := t.m[v]
		if k {
			r[v] = struct{}{}
		}
	}

	m.m = r
}
