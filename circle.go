package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	qt "github.com/machinbrol/quadtree"
	vec "github.com/machinbrol/vecmath"
)

type circle struct {
	x      int32
	y      int32
	center vec.Vec2
	r      float64
	ok     bool // continue à grandir
	color  rl.Color
}

func newCircle(r float64) *circle {
	circ := &circle{}
	circ.r = r
	circ.ok = false

	i := 0
	for i < maxCreationTries && len(points) > 0 && !circ.ok {
		idx := rand.Intn(len(points))

		pt := &point{}
		col := &color.RGBA{}

		points, pt = pop(points, idx)

		circ.x = pt.x
		circ.y = pt.y
		circ.center = vec.Vec2{X: float64(pt.x), Y: float64(pt.y)}

		qtree.Insert(circ)

		collide, _ := circ.collide()
		circ.ok = !collide
		if circ.ok {
			colors, col = popColor(colors, idx)
			circ.color = rl.NewColor(col.R, col.G, col.B, col.A)
		} else {
			qtree.Remove(circ)
		}
		i++
	}
	if i > maxCreationTries || len(points) == 0 {
		stop = true
		fmt.Println("i = ", i, "len(points)", len(points))
	}
	return circ
}

func (c *circle) Center() vec.Vec2 {
	return c.center
}

func (c *circle) Width() float64 {
	return c.r
}

func (c *circle) Height() float64 {
	return c.r
}

func (c *circle) Intersect(other qt.Centered) bool {
	return c.r+other.Width() > other.Center().Sub(c.Center()).Length()
}

func (c *circle) grow(g float64) {
	if c.ok {
		c.r += g
	}
}

func (c *circle) keepGrowing() {
	if c.ok {
		collision, _ := c.collide()
		c.ok = !c.edges() && !collision // arrête de grandir si touche côtés
	}
}

func (c *circle) edges() bool {
	return c.x+int32(c.r)+pad > int32(screenWidth) ||
		c.x-int32(c.r)-pad < 0 ||
		c.y+int32(c.r)+pad > int32(screenHeight) ||
		c.y-int32(c.r)-pad < 0
}

func (c *circle) distance(other *circle) float64 {
	dx := other.x - c.x
	dy := other.y - c.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

//other retourne true s'il y a collision, false sinon
func (c *circle) collide() (bool, []qt.Centered) {
	rect := qt.NewRectangleCentered(c.center, c.r*2, c.r*2)

	res := qtree.QueryRange(rect)

	//ajout à nbrComp
	nbrComp += len(res)

	i := 0
	for _, r := range res {
		if c != r && c.r+r.(*circle).r+float64(pad) > c.distance(r.(*circle)) {
			return true, res
		}
		i++
	}
	return false, res
}

func (c *circle) draw(clr rl.Color) {
	rl.DrawCircle(c.x, c.y, float32(c.r), c.color)
	// rl.DrawCircleLines(c.x, c.y, float32(c.r), clr)
}
