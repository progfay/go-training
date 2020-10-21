package intset

import (
	"bytes"
	"fmt"
)

const bitLength = 32 << (^uint(0) >> 63)

var pc [256]byte

func init() {
	fmt.Println(bitLength)
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x uint) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x uint) bool {
	if x < 0 {
		return false
	}
	word, bit := x/bitLength, uint(x%bitLength)
	return word < uint(len(s.words)) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x uint) {
	word, bit := x/bitLength, uint(x%bitLength)
	for word >= uint(len(s.words)) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitLength; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitLength*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popCount(word)
	}
	return count
}

func (s *IntSet) Remove(x uint) {
	word, bit := x/bitLength, uint(x%bitLength)
	for word >= uint(len(s.words)) {
		s.words = append(s.words, 0)
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	ret := &IntSet{}
	if s.words == nil {
		return ret
	}
	ret.words = make([]uint, len(s.words))
	copy(ret.words, s.words)
	return ret
}
