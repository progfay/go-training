package intset

func (s *IntSet) Elems() []uint {
	elems := []uint{}

	for i, word := range s.words {
		bit := 0
		for word > 0 {
			if word&1 != 0 {
				elems = append(elems, uint(i*bitLength+bit))
			}
			bit++
			word >>= 1
		}
	}

	return elems
}
