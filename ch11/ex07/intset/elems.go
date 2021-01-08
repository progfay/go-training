package intset

func (s *IntSet) Elems() []int {
	elems := []int{}

	for i, word := range s.words {
		bit := 0
		for word > 0 {
			if word&1 != 0 {
				elems = append(elems, i*64+bit)
			}
			bit++
			word >>= 1
		}
	}

	return elems
}
