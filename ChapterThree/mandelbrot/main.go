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
			img.Set(px, py, mandelbrotFunc(z))
		}
	}
	// png.Encode(os.Stdout, img)
	file, err := os.Create("mandelbrot2.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func mandelbrotFunc(z complex128) color.Color {
	const iterations = 20

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// return color.Gray{255 - contrast*n}
			return colorize(n)
		}
	}
	return color.Black
}

func colorize(n uint8) color.Color {
	const contrast = 15
	return color.RGBA{255 - contrast*n, 128 - contrast*n, 255, 255}
}
