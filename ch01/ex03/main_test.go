package main_test

import (
	"strings"
	"testing"
)

func Benchmark_LongStringConcat(b *testing.B) {
	var slice = []string{
		strings.Repeat("1", 1000000),
		strings.Repeat("2", 1000000),
		strings.Repeat("3", 1000000),
		strings.Repeat("4", 1000000),
		strings.Repeat("5", 1000000),
	}

	b.ResetTimer()

	b.Run("assignment operator", func(b *testing.B) {
		str := ""
		for _, s := range slice {
			str += " " + s
		}
	})

	b.Run("strings.Join", func(b *testing.B) {
		_ = strings.Join(slice, " ")
	})
}
