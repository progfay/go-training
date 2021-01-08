package intset

import (
	"bytes"
	"fmt"
)

var pc [256]int

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + (i & 1)
	}
}

type IntSet struct {
	words []uint16
}

func New() *IntSet {
	s := &IntSet{}
	s.words = make([]uint16, 0)
	return s
}

func (s *IntSet) Has(x int) bool {
	if x < 0 {
		return false
	}
	word, bit := x>>16, x&0xFFFF
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	if x < 0 {
		panic("add negative integer to IntSet")
	}
	word, bit := x>>16, x&0xFFFF
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		l, r := word>>8, word&0xFFFF
		count += pc[l] + pc[r]
	}
	return count
}

func (s *IntSet) Remove(x int) {
	if x < 0 {
		return
	}
	word, bit := x>>16, x&0xFFFF
	for word >= len(s.words) {
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
	ret.words = make([]uint16, len(s.words))
	copy(ret.words, s.words)
	return ret
}
