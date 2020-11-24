package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

const (
	iterations             = 100
	contrast               = 15
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

type point struct {
	x, y int
	c    color.Color
}

func main() {
	wg := sync.WaitGroup{}
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		points := make(chan point, width)
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			wg.Add(1)
			go func(px, py int, x, y float64) {
				defer wg.Done()
				z := complex(x, y)
				c := Complex128Mandelbrot(z)
				points <- point{px, py, c}
			}(px, py, x, y)
		}
		wg.Wait()
		close(points)
		for point := range points {
			img.Set(point.x, point.y, point.c)
		}
	}

	png.Encode(os.Stdout, img)
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
