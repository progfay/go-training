package intset

import (
	origin "github.com/progfay/go-training/ch06/ex01/intset"
)

type IntSet struct {
	origin.IntSet
}

func (s *IntSet) AddAll(entries ...int) {
	for _, entry := range entries {
		s.Add(entry)
	}
}
