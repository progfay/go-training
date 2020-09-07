package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

// const (
//
// )

// var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", diagramHandler)

	fmt.Println("Example: http://localhost:8000/?cells=100&width=500&height=500&xyrange=30&angle=15&xyscale=10&zscale=150&max-r=255&max-g=182&max-b=38&min-r=0&min-g=165&min-b=173")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func diagramHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}

	var (
		width, height int     = 600, 320
		cells         int     = 100
		xyrange       float64 = 30.0
		angle         float64 = math.Pi / 6
	)

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

	if len(r.Form["cells"]) > 0 && r.Form["cells"][0] != "" {
		c, err := strconv.Atoi(r.Form["cells"][0])
		if err == nil {
			cells = c
		}
	}

	if len(r.Form["xyrange"]) > 0 && r.Form["xyrange"][0] != "" {
		r, err := strconv.ParseFloat(r.Form["xyrange"][0], 64)
		if err == nil {
			xyrange = r
		}
	}

	if len(r.Form["angle"]) > 0 && r.Form["angle"][0] != "" {
		a, err := strconv.Atoi(r.Form["angle"][0])
		if err == nil {
			angle = float64(a) / 180 * math.Pi
		}
	}

	var (
		xyscale float64 = float64(width) / 2 / xyrange
		zscale  float64 = float64(height) * 0.4
	)

	if len(r.Form["xyscale"]) > 0 && r.Form["xyscale"][0] != "" {
		r, err := strconv.ParseFloat(r.Form["xyscale"][0], 64)
		if err == nil {
			xyscale = r
		}
	}

	if len(r.Form["zscale"]) > 0 && r.Form["zscale"][0] != "" {
		r, err := strconv.ParseFloat(r.Form["zscale"][0], 64)
		if err == nil {
			zscale = r
		}
	}

	corners, maxY, minY := calculateDiagramCorners(cells, width, height, angle, xyrange, xyscale, zscale)

	var (
		maxR, maxG, maxB int = 255, 0, 0
		minR, minG, minB int = 0, 0, 255
	)

	if len(r.Form["max-r"]) > 0 && r.Form["max-r"][0] != "" {
		r, err := strconv.Atoi(r.Form["max-r"][0])
		if err == nil {
			maxR = r
		}
	}

	if len(r.Form["max-g"]) > 0 && r.Form["max-g"][0] != "" {
		g, err := strconv.Atoi(r.Form["max-g"][0])
		if err == nil {
			maxG = g
		}
	}

	if len(r.Form["max-b"]) > 0 && r.Form["max-b"][0] != "" {
		b, err := strconv.Atoi(r.Form["max-b"][0])
		if err == nil {
			maxB = b
		}
	}

	if len(r.Form["min-r"]) > 0 && r.Form["min-r"][0] != "" {
		r, err := strconv.Atoi(r.Form["min-r"][0])
		if err == nil {
			minR = r
		}
	}

	if len(r.Form["min-g"]) > 0 && r.Form["min-g"][0] != "" {
		g, err := strconv.Atoi(r.Form["min-g"][0])
		if err == nil {
			minG = g
		}
	}

	if len(r.Form["min-b"]) > 0 && r.Form["min-b"][0] != "" {
		b, err := strconv.Atoi(r.Form["min-b"][0])
		if err == nil {
			minB = b
		}
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	writeSvg(w, corners, width, height, maxY, minY, maxR, maxG, maxB, minR, minG, minB)
}

func writeSvg(w io.Writer, corners [][]float64, width, height int, maxY, minY float64, maxR, maxG, maxB, minR, minG, minB int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for _, corner := range corners {
		y := (corner[1] + corner[3] + corner[5] + corner[7]) * .25
		norm := (y - minY) / (maxY - minY)

		fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: rgba(%v,%v,%v,0.5); stroke-width: 0.3' />\n",
			corner[0], corner[1], corner[2], corner[3], corner[4], corner[5], corner[6], corner[7], lerp(minR, maxR, norm), lerp(minG, maxG, norm), lerp(minB, maxB, norm))
	}

	fmt.Fprintf(w, "</svg>")
}

func lerp(max, min int, norm float64) float64 {
	return float64(max-min)*norm + float64(min)
}

func calculateDiagramCorners(cells, width, height int, angle, xyrange, xyscale, zscale float64) (corners [][]float64, maxY, minY float64) {
	maxY, minY = math.Inf(-1), math.Inf(0)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j, cells, width, height, angle, xyrange, xyscale, zscale)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j, cells, width, height, angle, xyrange, xyscale, zscale)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1, cells, width, height, angle, xyrange, xyscale, zscale)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1, cells, width, height, angle, xyrange, xyscale, zscale)
			if err != nil {
				continue
			}

			y := (ay + by + cy + dy) * .25
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}

			corners = append(corners, []float64{ax, ay, bx, by, cx, cy, dx, dy})
		}
	}

	return corners, maxY, minY
}

func corner(i, j, cells, width, height int, angle, xyrange, xyscale, zscale float64) (float64, float64, error) {
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) || math.IsInf(z, -1) || math.IsNaN(z) {
		return 0, 0, fmt.Errorf("Invalid float64 value (x, y, z): (%v, %v, %v)", x, y, z)
	}

	sx := float64(width/2) + (x-y)*math.Cos(angle)*xyscale
	sy := float64(height/2) + (x+y)*math.Sin(angle)*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
