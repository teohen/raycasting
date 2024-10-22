package player

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	ray "github.com/teohen/FPV/ray"
	vector2 "github.com/teohen/FPV/vector"
)

var CONE_COLOR = color.RGBA{200, 200, 200, 100}

const RAYS_COUNT = 300
const CONE_ANGLE = math.Pi / 2

type Player struct {
	Size   vector2.Vector2
	Pos    vector2.Vector2
	Dir    float64
	Color  color.Color
	Sprite *canvas.Circle
	Rays   []*ray.Ray
}

func (p *Player) debug() {
	fmt.Println("Size :", p.Size)
	fmt.Println("Pos :", p.Pos)
	fmt.Println("Dir :", p.Dir)
	fmt.Println("Color :", p.Color)
}

func NewPlayer(size, pos vector2.Vector2, dir float64) Player {

	p := Player{
		Size:  size,
		Pos:   pos,
		Dir:   dir,
		Color: color.RGBA{255, 0, 0, 255},
	}

	p.Sprite = canvas.NewCircle(color.RGBA{255, 0, 0, 255})

	v := vector2.Vector2{X: 0, Y: 0}

	rayDir := dir - (CONE_ANGLE / 2)

	for i := 0; i < RAYS_COUNT; i++ {
		ray := ray.NewRay(v, v, rayDir, 1)
		rayDir = rayDir - CONE_ANGLE/(RAYS_COUNT-1)
		p.Rays = append(p.Rays, &ray)
	}

	return p
}

func (p *Player) GetSprites() []fyne.CanvasObject {
	var sprites = make([]fyne.CanvasObject, 0)
	sprites = append(sprites, p.Sprite)
	for _, ray := range p.Rays {
		sprites = append(sprites, ray.GetSprites())
	}
	return sprites
}

func (p *Player) Rotate(angle float64) {
	p.Dir += angle
}

func (p *Player) Move(num float64) {
	v := vector2.Vector2{}
	v = v.FromAngle(p.Dir)
	v = v.Multiply(num)
	p.Pos.X += v.X
	p.Pos.Y += v.Y

}

func renderCircle(c *canvas.Circle, size, from, to vector2.Vector2, cColor color.Color) {
	c.Resize(fyne.NewSize(size.To32().X, size.To32().Y))
	c.Position1 = fyne.Position(from.To32())
	c.Position2 = fyne.Position(to.To32())
	c.FillColor = cColor
}

func getCenterPoint(c *canvas.Circle) vector2.Vector2 {
	cSize := c.Position2.Subtract(c.Position1)
	return vector2.Vector2{X: float64(c.Position1.X + (cSize.X / 2)), Y: float64(c.Position1.Y + (cSize.Y)/2)}
}

func findPointInCircle(c *canvas.Circle, dir float64, scale float64) vector2.Vector2 {
	centerPoint := getCenterPoint(c)
	hfx := (centerPoint.X - float64(c.Position1.X)) * (centerPoint.X - float64(c.Position1.X))
	hfy := (centerPoint.Y - float64(c.Position1.Y)) * (centerPoint.Y - float64(c.Position1.Y))

	radius := math.Sqrt(hfx + hfy)

	p2x := radius * math.Cos(dir) * scale
	p2y := radius * math.Sin(dir) * scale

	return vector2.Vector2{
		X: (centerPoint.X + p2x),
		Y: (centerPoint.Y + p2y),
	}
}

func calcRayDir(pDir float64, count int) vector2.Vector2 {

	finalAngle := pDir - (CONE_ANGLE / 2) + (CONE_ANGLE/(RAYS_COUNT-1))*float64(count)
	v := vector2.Vector2{}
	return v.FromAngle(finalAngle)
}

func (p *Player) renderRays() {
	from := getCenterPoint(p.Sprite)
	for i, ray := range p.Rays {
		ray.SetDir(calcRayDir(p.Dir, i))
		ray.SetFrom(from)
		ray.Cast()
	}
}

func (p *Player) Render() {
	renderCircle(p.Sprite, p.Size, p.Pos, p.Pos.Sum(p.Size), p.Color)
	p.renderRays()
}

func (p *Player) Refresh() {
	p.Render()
	p.Sprite.Refresh()
}

func (p *Player) GetRays() []*ray.Ray {
	return p.Rays
}
