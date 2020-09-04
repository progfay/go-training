package main_test

import (
	"testing"

	popcount "github.com/progfay/go-training/ch02/ex05"
)

func Benchmark_PopCount(b *testing.B) {
	var x uint64 = 0xDA18

	b.Run("with bit shift", func(b *testing.B) {
		for i := 0; i < 1000000; i++ {
			popcount.PopCountWithShift(x)
		}
	})

	b.Run("x & (x - 1)", func(b *testing.B) {
		for i := 0; i < 1000000; i++ {
			popcount.PopCountWithLoop(x)
		}
	})
}
