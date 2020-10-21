package intset

func (s *IntSet) AddAll(entries ...uint) {
	for _, entry := range entries {
		s.Add(entry)
	}
}
