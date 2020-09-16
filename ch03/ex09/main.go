package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Example: http://localhost:8080/?xmin=0&ymin=0&xmax=1&ymax=1&width=1500&height=1500")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}

	var (
		xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
		width, height                  = 1024, 1024
	)

	if len(r.Form["xmin"]) > 0 && r.Form["xmin"][0] != "" {
		x, err := strconv.ParseFloat(r.Form["xmin"][0], 64)
		if err == nil {
			xmin = x
		}
	}

	if len(r.Form["ymin"]) > 0 && r.Form["ymin"][0] != "" {
		y, err := strconv.ParseFloat(r.Form["ymin"][0], 64)
		if err == nil {
			ymin = y
		}
	}

	if len(r.Form["xmax"]) > 0 && r.Form["xmax"][0] != "" {
		x, err := strconv.ParseFloat(r.Form["xmax"][0], 64)
		if err == nil {
			xmax = x
		}
	}

	if len(r.Form["ymax"]) > 0 && r.Form["ymax"][0] != "" {
		y, err := strconv.ParseFloat(r.Form["ymax"][0], 64)
		if err == nil {
			ymax = y
		}
	}

	if len(r.Form["width"]) > 0 && r.Form["width"][0] != "" {
		w, err := strconv.Atoi(r.Form["width"][0])
		if err == nil {
			width = w
		}
	}

	if len(r.Form["height"]) > 0 && r.Form["height"][0] != "" {
		h, err := strconv.Atoi(r.Form["height"][0])
		if err == nil {
			height = h
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*ymax - ymin + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*xmax - xmin + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
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
