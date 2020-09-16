package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, supersampling(img))
}

func supersampling(origin *image.RGBA) *image.RGBA {
	bounds := origin.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	img := image.NewRGBA(image.Rect(0, 0, width-1, height-1))

	for x := 0; x < width-1; x++ {
		for y := 0; y < height-1; y++ {
			rgba1 := origin.RGBAAt(x+0, y+0)
			rgba2 := origin.RGBAAt(x+1, y+0)
			rgba3 := origin.RGBAAt(x+0, y+1)
			rgba4 := origin.RGBAAt(x+1, y+1)
			rgba := color.RGBA{
				R: uint8((uint64(rgba1.R) + uint64(rgba2.R) + uint64(rgba3.R) + uint64(rgba4.R)) / 4),
				G: uint8((uint64(rgba1.G) + uint64(rgba2.G) + uint64(rgba3.G) + uint64(rgba4.G)) / 4),
				B: uint8((uint64(rgba1.B) + uint64(rgba2.B) + uint64(rgba3.B) + uint64(rgba4.B)) / 4),
				A: uint8((uint64(rgba1.A) + uint64(rgba2.A) + uint64(rgba3.A) + uint64(rgba4.A)) / 4),
			}
			img.Set(x, y, rgba)
		}
	}

	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{255 - contrast*n, 255 - contrast*n, contrast * n}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
