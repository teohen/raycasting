package player

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/teohen/FPV/draw"
	"github.com/teohen/FPV/global"
	"github.com/teohen/FPV/ray"
	vector2 "github.com/teohen/FPV/vector"
)

var CIRCLE_COLOR = color.RGBA{255, 0, 0, 255}
var VISION_LINE_COLOR = color.RGBA{0, 255, 0, 255}

type PlayerSprite struct {
	Circle *canvas.Circle
	tLine  *canvas.Line
	rLine  *canvas.Line
	lLine  *canvas.Line
}

type Player struct {
	Size        vector2.Vector2
	Pos         vector2.Vector2
	Dir         float64
	Sprite      *PlayerSprite
	Rays        []*ray.Ray
	CenterPoint vector2.Vector2
}

func NewPlayer(size, pos vector2.Vector2, dir float64) Player {
	p := Player{
		Size: size,
		Pos:  pos,
		Dir:  dir,
	}

	p.updateCenterPoint()

	pSprite := PlayerSprite{
		Circle: canvas.NewCircle(CIRCLE_COLOR),
		rLine:  canvas.NewLine(VISION_LINE_COLOR),
		lLine:  canvas.NewLine(VISION_LINE_COLOR),
		tLine:  canvas.NewLine(VISION_LINE_COLOR),
	}

	p.Sprite = &pSprite

	p.renderCircle()
	p.renderVision()

	for i := 0; i < global.RAYS_COUNT; i++ {
		dir := calcRayDir(p.Dir, i)
		ray := ray.NewRay(p.CenterPoint, *p.CenterPoint.Sum(dir.Multiply(25)), *dir, 1)
		ray.Render()
		p.Rays = append(p.Rays, &ray)
	}

	return p
}

func (p *Player) Print() {
	fmt.Println(p.Pos)
}

func (p *Player) GetSprites() []fyne.CanvasObject {
	var sprites = make([]fyne.CanvasObject, 0)
	sprites = append(sprites, p.Sprite.Circle)
	sprites = append(sprites, p.Sprite.lLine)
	sprites = append(sprites, p.Sprite.rLine)
	sprites = append(sprites, p.Sprite.tLine)
	for _, ray := range p.Rays {
		sprites = append(sprites, ray.GetSprite())
	}
	return sprites
}

func (p *Player) Rotate(angle float64) {
	p.Dir += angle
}

func (p *Player) Move(num float64) {
	v := vector2.FromAngle(p.Dir).Multiply(num)
	p.Pos.X += v.X
	p.Pos.Y += v.Y
}
func (p *Player) updateCenterPoint() {
	p.CenterPoint.X = p.Pos.X + (p.Size.X / 2)
	p.CenterPoint.Y = p.Pos.Y + (p.Size.Y / 2)
}

func (p *Player) renderCircle() {
	draw.Circle(p.Sprite.Circle, p.Size, p.Pos, p.Sprite.Circle.FillColor)
}

func calcRayDir(pDir float64, count int) *vector2.Vector2 {
	finalAngle := pDir - (global.CONE_ANGLE / 2) + (global.CONE_ANGLE/(global.RAYS_COUNT-1))*float64(count)
	v := vector2.FromAngle(finalAngle)
	return v
}

func (p *Player) renderRays() {
	for i, ray := range p.Rays {
		ray.SetFrom(&p.CenterPoint)
		ray.SetDir(calcRayDir(p.Dir, i))
		ray.Cast()
	}
}

func (p *Player) renderVision() {
	rLineTo := *vector2.FindPointInCircle(p.CenterPoint, p.Pos, (p.Dir - (global.CONE_ANGLE / 2)), 3)
	lLineTo := *vector2.FindPointInCircle(p.CenterPoint, p.Pos, (p.Dir + (global.CONE_ANGLE / 2)), 3)

	draw.Line(p.Sprite.lLine, p.CenterPoint, lLineTo, p.Sprite.lLine.StrokeColor, 1)
	draw.Line(p.Sprite.rLine, p.CenterPoint, rLineTo, p.Sprite.rLine.StrokeColor, 1)
	draw.Line(p.Sprite.tLine, rLineTo, lLineTo, p.Sprite.tLine.StrokeColor, 1)
}

func (p *Player) Render() {
	p.updateCenterPoint()
	p.renderCircle()
	p.renderVision()
	p.renderRays()
}

func (p *Player) Refresh() {
	p.Render()
	p.Sprite.Circle.Refresh()
}

func (p *Player) GetRays() []*ray.Ray {
	return p.Rays
}
