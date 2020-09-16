package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"

	. "github.com/progfay/go-training/ch03/ex08/complex"
)

const iterations = 100
const contrast = 15

var modes = []string{"complex64", "complex128", "big.Float", "big.Rat"}
var mu sync.Mutex

func supportMode(mode string) bool {
	for _, m := range modes {
		if mode == m {
			return true
		}
	}
	return false
}

func main() {
	mode := os.Args[1]
	if !supportMode(mode) {
		log.Panicf("Unsupported mode: %s", mode)
	}

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	wg := sync.WaitGroup{}
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			wg.Add(1)
			go func(px, py int, x, y float64) {
				defer wg.Done()

				switch mode {
				case "complex64":
					var z complex64 = complex(float32(x), float32(y))
					mu.Lock()
					img.Set(px, py, Complex64Mandelbrot(z))
					mu.Unlock()

				case "complex128":
					var z complex128 = complex(x, y)
					mu.Lock()
					img.Set(px, py, Complex128Mandelbrot(z))
					mu.Unlock()

				case "big.Float":
					z := NewBigFloatComplex(x, y)
					mu.Lock()
					img.Set(px, py, BigFloatComplexMandelbrot(z))
					mu.Unlock()

				case "big.Rat":
					z := NewBigRatComplex(x, y)
					mu.Lock()
					img.Set(px, py, BigRatComplexMandelbrot(z))
					mu.Unlock()
				}
			}(px, py, x, y)
		}
	}
	wg.Wait()
	png.Encode(os.Stdout, img)
}

func Complex64Mandelbrot(z complex64) color.Color {
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func Complex128Mandelbrot(z complex128) color.Color {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func BigFloatComplexMandelbrot(z *BigFloatComplex) color.Color {
	v := NewBigFloatComplex(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v = v.Mul(v).Add(z)
		if v.Abs() > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func BigRatComplexMandelbrot(z *BigRatComplex) color.Color {
	v := NewBigRatComplex(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v = v.Mul(v).Add(z)
		if v.Abs() > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
