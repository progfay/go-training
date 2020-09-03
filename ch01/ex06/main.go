package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/progfay/colorcontrast"
)

var palette = []color.Color{}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	for r := 0x0F; r <= 0xFF; r += 0x0F {
		for g := 0x0F; g <= 0xFF; g += 0x0F {
			for b := 0x0F; b <= 0xFF; b += 0x0F {
				c := color.RGBA{uint8(r), uint8(g), uint8(b), 0xFF}
				contrast := colorcontrast.CalcContrastRatio(c, color.Black)
				if 8 < contrast && contrast < 10 {
					palette = append(palette, c)
				}
			}
		}
	}

	rand.Shuffle(len(palette), func(i, j int) { palette[i], palette[j] = palette[j], palette[i] })
	palette = append([]color.Color{color.Black}, palette[:0xFE]...)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {

		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			index := rand.Intn(len(palette)-1) + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(index))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
