package intset

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] &= word
		} else {
			s.words = append(s.words, word)
		}
	}
}
