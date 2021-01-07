package bench_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/progfay/go-training/ch11/ex07/intmap"
	"github.com/progfay/go-training/ch11/ex07/intset"
)

var (
	m        *intmap.IntMap
	s        *intset.IntSet
	rm       *intmap.IntMap
	rs       *intset.IntSet
	wordSize = []int{1000, 1000000}
	sliceLen = 10
	dst      interface{}
)

func init() {
	rand.Seed(time.Now().UnixNano())
	m, rm = intmap.New(), intmap.New()
	s, rs = intset.New(), intset.New()
}

func randomIntSlice(max int) []int {
	r := make([]int, sliceLen)
	for i := 0; i < sliceLen; i++ {
		r[i] = rand.Intn(max)
	}
	return r
}

func Benchmark_Add(b *testing.B) {
	for _, w := range wordSize {
		b.Run(fmt.Sprintf("type: intmap; word < %d", w), func(b *testing.B) {
			m.Clear()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r := rand.Intn(w)
				m.Add(r)
			}
		})

		b.Run(fmt.Sprintf("type: intset; word < %d", w), func(b *testing.B) {
			s.Clear()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r := rand.Intn(w)
				s.Add(r)
			}
		})
	}
}

func Benchmark_UnionWith(b *testing.B) {
	for _, w := range wordSize {
		b.Run(fmt.Sprintf("type: intmap; word < %d", w), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Clear()
				rm.Clear()
				m.AddAll(randomIntSlice(w)...)
				rm.AddAll(randomIntSlice(w)...)
				m.UnionWith(rm)
			}
		})

		b.Run(fmt.Sprintf("type: intset; word < %d", w), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Clear()
				rs.Clear()
				s.AddAll(randomIntSlice(w)...)
				rs.AddAll(randomIntSlice(w)...)
				s.UnionWith(rs)
			}
		})
	}
}

func Benchmark_DifferenceWith(b *testing.B) {
	for _, w := range wordSize {
		b.Run(fmt.Sprintf("type: intmap; word < %d", w), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Clear()
				rm.Clear()
				m.AddAll(randomIntSlice(w)...)
				rm.AddAll(randomIntSlice(w)...)
				m.DifferenceWith(rm)
			}
		})

		b.Run(fmt.Sprintf("type: intset; word < %d", w), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Clear()
				rs.Clear()
				s.AddAll(randomIntSlice(w)...)
				rs.AddAll(randomIntSlice(w)...)
				s.DifferenceWith(rs)
			}
		})
	}
}

func Benchmark_IntersectWith(b *testing.B) {
	for _, w := range wordSize {
		b.Run(fmt.Sprintf("type: intmap; word < %d", w), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Clear()
				rm.Clear()
				m.AddAll(randomIntSlice(w)...)
				rm.AddAll(randomIntSlice(w)...)
				m.IntersectWith(rm)
			}
		})

		b.Run(fmt.Sprintf("type: intset; word < %d", w), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Clear()
				rs.Clear()
				s.AddAll(randomIntSlice(w)...)
				rs.AddAll(randomIntSlice(w)...)
				s.IntersectWith(rs)
			}
		})
	}
}

func Benchmark_SymmetricDifference(b *testing.B) {
	for _, w := range wordSize {
		b.Run(fmt.Sprintf("type: intmap; word < %d", w), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Clear()
				rm.Clear()
				m.AddAll(randomIntSlice(w)...)
				rm.AddAll(randomIntSlice(w)...)
				m.SymmetricDifference(rm)
			}
		})

		b.Run(fmt.Sprintf("type: intset; word < %d", w), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Clear()
				rs.Clear()
				s.AddAll(randomIntSlice(w)...)
				rs.AddAll(randomIntSlice(w)...)
				s.SymmetricDifference(rs)
			}
		})
	}
}
