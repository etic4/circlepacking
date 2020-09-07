package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	qt "github.com/machinbrol/quadtree"
	vec "github.com/machinbrol/vecmaths"
)

const (
	// screenWidth  int   = 1200
	// screenHeight int   = 1000
	pad int32 = 0
)

var (
	fps                int32
	nbrFrames          int32
	maxFrames          int32
	screenWidth        int
	screenHeight       int
	circList           []*circle
	img                image.Image
	colors             []*color.RGBA
	points             []*point
	filename           string
	maxIntialCercleRay int
	growSpeed          float64
	nbCerclesParFrame  int
	maxCreationTries   int
	qtree              *qt.Quadtree
	nbrComp            int
	pause              bool
	stop               bool
)

func atMouseCoords() {
	p := rl.GetMousePosition()

	pos := vec.Vec2{X: float64(p.X), Y: float64(p.Y)}

	circ := &circle{}
	qtLst := []*qt.Quadtree{}

	i := len(circList) - 1
	for i > 0 && pos.Distance(circList[i].center) > circList[i].r {
		i--
	}

	if i < len(circList) {
		circ = circList[i]
		qtLst = qtree.GetQuadtreesFor(circ.center)
	}

	circ.draw(rl.Red)
	for _, qtree := range qtLst {
		// qtree.Draw()
		qtree.DrawOne()
	}
}

func init() {
	circList = []*circle{}

	// img, pList = getBlackPixelsFromFile("fleurs.png")
	// colors, points = getAllPixelsFromFile(filename)

}

func update() {
	nbrFrames++
	rl.BeginDrawing()

	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		pause = !pause
		fmt.Println("Toggle Pause: ", pause)
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if pause {
			rl.ClearBackground(rl.Blank)
			for _, c := range circList {
				c.draw(rl.White)
			}
			atMouseCoords()
		}
	}

	if !stop && !pause && nbrFrames < maxFrames {
		rl.ClearBackground(rl.Blank)

		qtree = getQtree()
		// qtree.Clear()

		for _, c := range circList {
			qtree.Insert(c)
		}

		nbrComp = 0

		for _, circ := range circList {
			circ.grow(growSpeed)
			circ.keepGrowing()
			circ.draw(rl.White)
		}

		for i := 0; i < nbCerclesParFrame; i++ {
			circ := newCircle(float64(rand.Intn(maxIntialCercleRay + 1)))
			circList = append(circList, circ)
			circ.draw(rl.White)
		}
		fmt.Println("circList", len(circList))
		fmt.Println("qtree.Size()", qtree.Size())
		fmt.Println("nbrComp", nbrComp)
		if stop {
			fmt.Println("Stopped")
		}
	}

	rl.EndDrawing()
}

func main() {
	fps = int32(60)
	filename = "fleurs.png"
	screenWidth, screenHeight = getWidthHeight(filename)
	maxIntialCercleRay = 1
	growSpeed = 0.1
	nbCerclesParFrame = 100
	maxFrames = fps * 60 * 5
	maxCreationTries = 1000

	colors, points = getAllPixelsFromFile(filename)

	qtree = getQtree()

	rl.InitWindow(int32(screenWidth), int32(screenHeight), "circle packing")

	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		update()

	}

	rl.CloseWindow()

}
