package popcount_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/progfay/go-training/ch11/ex06/popcount"
)

var dst1, dst2, dst3 int

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Benchmark_PopCount(b *testing.B) {
	b.Run("with bit shift", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x := rand.Uint64()
			dst1 += popcount.PopCountWithShift(x)
		}
	})

	b.Run("check each bit digits", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x := rand.Uint64()
			dst2 += popcount.PopCountWithLoop(x)
		}
	})

	b.Run("with x & (x - 1)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x := rand.Uint64()
			dst3 += popcount.PopCountWithClear(x)
		}
	})
}
