package intset

import (
	"bytes"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

type IntSet struct {
	words []uint64
}

func New() *IntSet {
	s := &IntSet{}
	s.words = make([]uint64, 0)
	return s
}

func (s *IntSet) Has(x int) bool {
	if x < 0 {
		return false
	}
	word, bit := x>>6, x&0x3F
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	if x < 0 {
		panic("add negative integer to IntSet")
	}
	word, bit := x>>6, x&0x3F
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
		count += int(pc[word])
	}
	return count
}

func (s *IntSet) Remove(x int) {
	if x < 0 {
		return
	}
	word, bit := x>>6, x&0x3F
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
	ret.words = make([]uint64, len(s.words))
	copy(ret.words, s.words)
	return ret
}
