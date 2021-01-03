package intmap

import "sort"

func (m *IntMap) Elems() []int {
	r := []int{}

	for k := range m.m {
		r = append(r, k)
	}

	sort.Ints(r)
	return r
}
