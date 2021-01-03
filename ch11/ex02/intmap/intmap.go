package intmap

import (
	"bytes"
	"fmt"
	"sort"
)

type IntMap struct {
	m map[int]struct{}
}

func New() *IntMap {
	s := &IntMap{}
	s.m = make(map[int]struct{}, 0)
	return s
}

func (m *IntMap) Has(x int) bool {
	if x < 0 {
		return false
	}
	_, ok := m.m[x]
	return ok
}

func (m *IntMap) Add(x int) {
	if x < 0 {
		panic("add negative integer to IntMap")
	}

	m.m[x] = struct{}{}
}

func (m *IntMap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	keys := make([]int, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", m.m[k])
	}

	buf.WriteByte('}')
	return buf.String()
}

func (m *IntMap) Len() int {
	return len(m.m)
}

func (m *IntMap) Remove(x int) {
	delete(m.m, x)
}

func (m *IntMap) Clear() {
	m.m = make(map[int]struct{}, 0)
}

func (m *IntMap) Copy() IntMap {
	ret := New()
	for k, v := range m.m {
		ret.m[k] = v
	}

	return *ret
}
