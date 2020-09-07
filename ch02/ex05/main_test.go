package main_test

import (
	"testing"

	popcount "github.com/progfay/go-training/ch02/ex05"
)

var x uint64 = 0xDA18
var dst1, dst2 int

func Benchmark_PopCount(b *testing.B) {

	b.Run("with bit shift", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst1 += popcount.PopCountWithShift(x)
		}
	})

	b.Run("with x & (x - 1)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst2 += popcount.PopCountWithLoop(x)
		}
	})
}
