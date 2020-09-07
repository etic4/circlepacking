package main

import (
	"image"
	"image/color"
	"os"

	qt "github.com/machinbrol/quadtree"
	vec "github.com/machinbrol/vecmaths"
)

func getWidthHeight(filename string) (int, int) {
	f, err := os.Open(filename)

	if err != nil {
		panic("Peut pas ouvrir le fichier")
	}

	defer f.Close()

	img, _, _ := image.Decode(f)

	bounds := img.Bounds()

	return bounds.Max.X, bounds.Max.Y
}

func getQtree() *qt.Quadtree {
	c := vec.Vec2{X: float64(screenWidth / 2), Y: float64(screenHeight / 2)}
	rect := qt.NewRectangleCentered(c, float64(screenWidth/2), float64(screenHeight/2))
	return qt.NewQuadtree(rect, 4, nil)
}

//enlève de liste
func pop(points []*point, i int) ([]*point, *point) {
	pt := points[i]
	points[i] = points[len(points)-1]
	points[len(points)-1] = new(point)
	points = points[:len(points)-1]

	return points, pt
}

func popColor(colors []*color.RGBA, i int) ([]*color.RGBA, *color.RGBA) {
	col := colors[i]
	colors[i] = colors[len(colors)-1]
	colors[len(colors)-1] = new(color.RGBA)
	colors = colors[:len(colors)-1]

	return colors, col
}
