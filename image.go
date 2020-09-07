package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png" // permet automatiquement le décodage des formats png

	// _ "image/gif"
	// _"image/jpeg"

	"os"
)

type point struct {
	x int32
	y int32
}

func getAllPixelsFromFile(filename string) ([]*color.RGBA, []*point) {
	fmt.Println("filename", filename)
	file, err := os.Open(filename)

	if err != nil {
		panic("Erreur: le fichier ne peut pas être ouvert")

	}

	defer file.Close()

	img, _, _ := image.Decode(file)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels []*color.RGBA
	var points []*point

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			col := toRGBA(img.At(x, y))
			pixels = append(pixels, &col)
			points = append(points, &point{int32(x), int32(y)})
		}
	}
	return pixels, points
}

func getBlackPixelsFromFile(filename string) (image.Image, []*point) {
	file, err := os.Open(filename)

	if err != nil {
		panic("Erreur: le fichier ne peut pas être ouvert")

	}

	defer file.Close()

	img, _, _ := image.Decode(file)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	black := color.RGBA{0, 0, 0, 255}

	var pixels []*point

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if toRGBA(img.At(x, y)) == black {
				pixels = append(pixels, &point{int32(x), int32(y)})
			}
		}
	}
	return img, pixels
}

func toRGBA(col color.Color) color.RGBA {
	R, G, B, A := col.RGBA()
	return color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)}
}
