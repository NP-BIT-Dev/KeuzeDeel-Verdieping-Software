package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Your original Pic function remains unchanged
func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		img[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			img[y][x] = uint8(x ^ y)
		}
	}
	return img
}

func main() {
	// 1. Define the size of the image (Go Tour uses 256x256)
	dx, dy := 256, 256
	
	// 2. Get your 2D slice of data
	grid := Pic(dx, dy)

	// 3. Create a new blank Go image
	img := image.NewRGBA(image.Rect(0, 0, dx, dy))

	// 4. Map your numbers to pixels
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := grid[y][x]
			// We map the number to the Blue channel to match the Go Tour's bluescale
			img.Set(x, y, color.RGBA{R: 0, G: 0, B: v, A: 255})
		}
	}

	// 5. Create a file and save the PNG
	file, err := os.Create("fractal.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}