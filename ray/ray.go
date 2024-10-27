package ray

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"github.com/teohen/FPV/draw"
	"github.com/teohen/FPV/global"
	vector2 "github.com/teohen/FPV/vector"
)

var RAY_COLOR = color.RGBA{0, 0, 255, 255}

type Ray struct {
	from   vector2.Vector2
	to     vector2.Vector2
	dir    vector2.Vector2
	len    float64
	sprite *canvas.Line
	color  color.Color
	right  Wall
	left   Wall
	top    Wall
	bottom Wall
}

func NewRay(from, to, dir vector2.Vector2, len float64) Ray {
	line := canvas.NewLine(RAY_COLOR)
	line.StrokeWidth = 1

	ray := Ray{
		from:   from,
		to:     to,
		dir:    dir,
		len:    len,
		sprite: line,
		color:  RAY_COLOR,
	}

	return ray
}

func (r *Ray) GetFrom() *vector2.Vector2 {
	return &r.from
}

func (r *Ray) SetFrom(from *vector2.Vector2) {
	r.from = *from
}

func (r *Ray) GetTo() *vector2.Vector2 {
	return &r.to
}

func (r *Ray) SetTo(to *vector2.Vector2) {
	r.to = *to
}

func (r *Ray) GetDir() *vector2.Vector2 {
	return &r.dir
}

func (r *Ray) SetDir(dir *vector2.Vector2) {
	r.dir = *dir
}

func (r *Ray) GetSprite() *canvas.Line {
	return r.sprite
}

func (r *Ray) Cast() {
	from := r.from
	to := r.getHittingPoint(r.from)
	for {
		to = r.getHittingPoint(from)
		from = *to

		x, y := global.GetXY(*to, r.from)
		if global.GetWorldCellContent(y, x) > 0 || !global.InsideGame(*to) {
			break
		}
	}
	if !global.InsideGame(*to) {
		to = &r.from
	}
	r.to = *to
	r.Render()
}

func (r *Ray) getHittingPoint(from vector2.Vector2) *vector2.Vector2 {
	v := vector2.Vector2{}
	for _, wall := range r.updateWalls(from, r.from).wallsInSlice() {
		p := vector2.LineToLineIntersection(wall.from, wall.to, from, r.dir)
		if p == nil {
			continue
		}
		v = *p
	}
	return &v

}

func (r *Ray) Render() {
	draw.Line(r.sprite, r.from, r.to, r.color, 1)
}

func (r *Ray) updateWalls(from, origin vector2.Vector2) *Ray {
	x, y := global.GetXY(from, origin)

	right := float64(global.CELL_W + x*global.CELL_W)
	left := float64(global.CELL_W + (x * global.CELL_W) - global.CELL_W)
	top := float64(global.CELL_H + (y * global.CELL_H) - global.CELL_H)
	bottom := float64(global.CELL_H + y*global.CELL_H)

	r.right = *NewWall(
		vector2.Vector2{X: right, Y: top},
		vector2.Vector2{X: right, Y: bottom},
	)

	r.left = *NewWall(
		vector2.Vector2{X: left, Y: top},
		vector2.Vector2{X: left, Y: bottom},
	)

	r.top = *NewWall(
		vector2.Vector2{X: left, Y: top},
		vector2.Vector2{X: right, Y: top},
	)

	r.bottom = *NewWall(
		vector2.Vector2{X: left, Y: bottom},
		vector2.Vector2{X: right, Y: bottom},
	)

	return r
}

func (r *Ray) wallsInSlice() [4]*Wall {
	return [4]*Wall{&r.right, &r.left, &r.top, &r.bottom}
}

func (r *Ray) CalcLength(dir float64) float64 {
	to := r.GetTo()
	d := vector2.FromAngle(dir)
	rFrom := r.GetFrom()
	v := to.Sub(rFrom).Dot(d)
	return v
}
