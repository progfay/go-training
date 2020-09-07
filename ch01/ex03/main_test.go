package main_test

import (
	"strings"
	"testing"
)

var slice = []string{
	strings.Repeat("1", 1000000),
	strings.Repeat("2", 1000000),
	strings.Repeat("3", 1000000),
	strings.Repeat("4", 1000000),
	strings.Repeat("5", 1000000),
}

var dst1, dst2 string

func Benchmark_LongStringConcat(b *testing.B) {
	b.Run("assignment operator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, s := range slice {
				dst1 += " " + s
			}
		}
	})

	b.Run("strings.Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst2 = strings.Join(slice, " ")
		}
	})
}
