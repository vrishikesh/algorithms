package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	var width, height = 200, 100

	topLeft := image.Point{0, 0}
	bottomRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	cyan := color.RGBA{100, 200, 200, 0xff}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case (x+y)%2 == 0:
				// case x < width/2 && y < height/2:
				img.Set(x, y, cyan)
			case (x+y)%2 == 1:
				// case x >= width/2 && y >= height/2:
				img.Set(x, y, color.White)
			default:
			}
		}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
