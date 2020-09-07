package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	var corners = [][]float64{}
	var maxY, minY float64 = math.Inf(-1), math.Inf(0)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
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

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for _, corner := range corners {
		y := (corner[1] + corner[3] + corner[5] + corner[7]) * .25
		norm := (y - minY) / (maxY - minY)

		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: rgba(%v,%v,%v,0.5); stroke-width: 0.3' />\n",
			corner[0], corner[1], corner[2], corner[3], corner[4], corner[5], corner[6], corner[7], 255 * (1 - norm), 0, 255 * norm)
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) || math.IsInf(z, -1) || math.IsNaN(z) {
		return 0, 0, fmt.Errorf("Invalid float64 value (x, y, z): (%v, %v, %v)", x, y, z)
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
