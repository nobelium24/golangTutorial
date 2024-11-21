package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	angle   = math.Pi / 6
	cells   = 100
	xyrange = 30
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/displaySVG", displaySVG)
	http.ListenAndServe(":8080", nil)
}

func displaySVG(out http.ResponseWriter, in *http.Request) {
	widthStr := in.URL.Query().Get("width")
	heightStr := in.URL.Query().Get("height")

	width, err := strconv.Atoi(widthStr)
	if err != nil {
		http.Error(out, "Invalid width parameter", http.StatusBadRequest)
		return
	}
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		http.Error(out, "Invalid height parameter", http.StatusBadRequest)
		return
	}

	xyscale := float64(width) / 2 / xyrange // Pixels per x or y unit
	zscale := float64(height) * 0.4         // Pixels per z unit

	out.Header().Set("Content-Type", "image/svg+xml")
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: red; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, width, height, xyscale, zscale)
			bx, by, bz := corner(i, j, width, height, xyscale, zscale)
			cx, cy, cz := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy, dz := corner(i+1, j+1, width, height, xyscale, zscale)
			if isFinite(ax, ay, bx, by, cx, cy, dx, dy) {
				color := getColor((az + bz + cz + dz) / 4)
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j, width, height int, xyscale, zscale float64) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func isFinite(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func getColor(z float64) string {
	if z > 0 {
		return "blue"
	}
	return "red"
}
