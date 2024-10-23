package ray

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/teohen/FPV/global"
	vector2 "github.com/teohen/FPV/vector"
)

const CELL_W = 40
const CELL_H = 40

type Ray struct {
	from   vector2.Vector2
	to     vector2.Vector2
	dir    vector2.Vector2
	len    float64
	sprite *canvas.Line
}

type Wall struct {
	From vector2.Vector2
	To   vector2.Vector2
	Name string
}

func NewRay(from, to vector2.Vector2, dir float64, len float64) Ray {
	line := canvas.NewLine(color.RGBA{0, 0, 255, 0})
	line.StrokeWidth = 1

	v := vector2.Vector2{}
	ray := Ray{
		from:   from,
		to:     to,
		dir:    v.FromAngle(dir),
		len:    len,
		sprite: line,
	}

	return ray
}

func (r *Ray) GetFrom() vector2.Vector2 {
	return r.from
}

func (r *Ray) SetFrom(from vector2.Vector2) {
	r.from = from
}

func (r *Ray) GetTo() vector2.Vector2 {
	return r.to
}

func (r *Ray) SetTo(to vector2.Vector2) {
	r.to = to
}

func (r *Ray) GetDir() vector2.Vector2 {
	return r.dir
}

func (r *Ray) SetDir(dir vector2.Vector2) {
	r.dir = dir
}

func (r *Ray) GetSprite() *canvas.Line {
	return r.sprite
}

func (r *Ray) DetSprite(line *canvas.Line) {
	r.sprite = line
}

func (r *Ray) GetSprites() *canvas.Line {
	return r.sprite
}

func (r *Ray) Cast() {
	from := r.from
	to := getPoint(r.from, r.dir, r.from)
	for {
		to = getPoint(from, r.dir, r.from)
		from = to

		x, y := getCellIdxs(to, r.from)
		if global.GetWorldCellContent(y, x) > 0 || !global.InsideGame(to) {
			break
		}

	}

	if !global.InsideGame(to) {
		to = r.from
	}
	r.to = to

	r.Render()

}

func getPoint(from, dir, origin vector2.Vector2) vector2.Vector2 {
	walls := getWalls(from, origin)
	for _, wall := range walls {
		x1 := wall.From.X
		y1 := wall.From.Y
		x2 := wall.To.X
		y2 := wall.To.Y

		x3 := from.X
		y3 := from.Y

		x4 := from.X + dir.X
		y4 := from.Y + dir.Y

		den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
		if den == 0 {
			continue
		}

		t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / den
		u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / den

		if t >= 0 && t <= 1 && u > 0 {
			return vector2.Vector2{
				X: x1 + t*(x2-x1),
				Y: y1 + t*(y2-y1),
			}
		} else {
			continue
		}
	}

	return vector2.Vector2{}
}

func (r *Ray) Render() {
	r.sprite.Position1 = fyne.Position(r.from.To32())
	r.sprite.Position2 = fyne.Position(r.to.To32())
}

func getCellIdxs(pos, origin vector2.Vector2) (int, int) {
	x := math.Floor(pos.X / CELL_W)
	y := math.Floor(pos.Y / CELL_H)

	if math.Mod(pos.X, CELL_W) == 0 && pos.X < origin.X {
		x -= 1
	}
	if math.Mod(pos.Y, CELL_H) == 0 && pos.Y < origin.Y {
		y -= 1
	}

	return int(x), int(y)
}

func getWalls(from, origin vector2.Vector2) []Wall {
	x, y := getCellIdxs(from, origin)

	cellpadw := CELL_W
	cellpadh := CELL_H
	r := float64(cellpadw + x*CELL_W)
	l := float64(cellpadw + (x * CELL_W) - CELL_W)
	t := float64(cellpadh + (y * CELL_H) - CELL_H)
	b := float64(cellpadh + y*CELL_H)

	rLine := Wall{
		From: vector2.Vector2{X: r, Y: t},
		To:   vector2.Vector2{X: r, Y: b},
		Name: "right",
	}
	lLine := Wall{
		From: vector2.Vector2{X: l, Y: t},
		To:   vector2.Vector2{X: l, Y: b},
		Name: "left",
	}
	tLine := Wall{
		From: vector2.Vector2{X: l, Y: t},
		To:   vector2.Vector2{X: r, Y: t},
		Name: "top",
	}
	bLine := Wall{
		From: vector2.Vector2{X: l, Y: b},
		To:   vector2.Vector2{X: r, Y: b},
		Name: "bottom",
	}
	return []Wall{rLine, lLine, tLine, bLine}
}

func (r *Ray) CalcLength(dir float64) float64 {
	vec := vector2.Vector2{}
	to := r.GetTo()
	v := to.Sub(r.GetFrom())

	d := vec.FromAngle(dir)
	return v.Dot(d)
}
