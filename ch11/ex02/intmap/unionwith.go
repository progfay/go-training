package intmap

func (m *IntMap) UnionWith(t *IntMap) {
	for k, v := range t.m {
		m.m[k] = v
	}
}
