package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	http.HandleFunc("/", lissajousHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	const (
		res    = 0.001
		size   = 100
		cycles = 5
	)
	var (
		nframes int = 64
		delay   int = 8
	)

	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}

	if len(r.Form["nframes"]) > 0 && r.Form["nframes"][0] != "" {
		f, err := strconv.Atoi(r.Form["nframes"][0])
		if err == nil {
			nframes = f
		}
	}

	if len(r.Form["delay"]) > 0 && r.Form["delay"][0] != "" {
		d, err := strconv.Atoi(r.Form["delay"][0])
		if err == nil {
			delay = d
		}
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
