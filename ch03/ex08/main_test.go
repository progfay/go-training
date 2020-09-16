package main_test

import (
	"image/color"
	"testing"

	main "github.com/progfay/go-training/ch03/ex08"
	. "github.com/progfay/go-training/ch03/ex08/complex"
)

var real float64 = 10.274
var imagine float64 = -423.15
var dst1, dst2, dst3, dst4 color.Color

func Benchmark_Mandelbrot(b *testing.B) {
	b.Run("complex64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var z complex64 = complex(float32(real), float32(imagine))
			dst1 = main.Complex64Mandelbrot(z)
		}
	})

	b.Run("complex128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var z complex128 = complex(real, imagine)
			dst2 = main.Complex128Mandelbrot(z)
		}
	})

	b.Run("big.Float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var z = NewBigFloatComplex(real, imagine)
			dst3 = main.BigFloatComplexMandelbrot(z)
		}
	})

	b.Run("big.Rat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z := NewBigRatComplex(real, imagine)
			dst4 = main.BigRatComplexMandelbrot(z)
		}
	})
}
