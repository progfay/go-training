package intset

func (s *IntSet) AddAll(entries ...int) {
	for _, entry := range entries {
		s.Add(entry)
	}
}
