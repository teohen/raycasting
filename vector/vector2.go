package vector2

import (
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

type Vector232 struct {
	X float32
	Y float32
}

func (v2 *Vector2) Multiply(num float64) *Vector2 {
	v := Vector2{v2.X * num, v2.Y * num}
	return &v
}

func (v2 *Vector2) SubtracBy(num float64) *Vector2 {
	v := Vector2{v2.X - num, v2.Y - num}
	return &v
}

func (v2 *Vector2) SumBy(num float64) *Vector2 {
	v := Vector2{v2.X + num, v2.Y + num}
	return &v
}

func (v2 *Vector2) To32() Vector232 {
	return Vector232{
		X: float32(v2.X),
		Y: float32(v2.Y),
	}
}

func (v2 *Vector2) Sum(v *Vector2) *Vector2 {
	vec := Vector2{
		X: v2.X + v.X,
		Y: v2.Y + v.Y,
	}

	return &vec
}
func FromAngle(angle float64) *Vector2 {
	v := Vector2{X: math.Cos(angle), Y: math.Sin(angle)}
	return &v
}

func (v *Vector2) Dot(v2 *Vector2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vector2) Sub(v2 *Vector2) *Vector2 {
	vec := Vector2{X: v.X - v2.X, Y: v.Y - v2.Y}
	return &vec
}
func LineToLineIntersection(ls, lf, p, dir Vector2) *Vector2 {
	x1 := ls.X
	y1 := ls.Y

	x2 := lf.X
	y2 := lf.Y
	x3 := p.X
	y3 := p.Y

	x4 := p.X + dir.X
	y4 := p.Y + dir.Y

	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if den == 0 {
		return nil
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / den
	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / den

	if t >= 0 && t <= 1 && u > 0 {
		v := Vector2{
			X: x1 + t*(x2-x1),
			Y: y1 + t*(y2-y1),
		}
		return &v
	}
	return nil
}

func FindPointInCircle(center, pos Vector2, dir float64, scale float64) *Vector2 {
	hfx := (center.X - pos.X) * (center.X - pos.X)
	hfy := (center.Y - pos.Y) * (center.Y - pos.Y)

	radius := math.Sqrt(hfx + hfy)

	p2x := radius * math.Cos(dir) * scale
	p2y := radius * math.Sin(dir) * scale

	return &Vector2{
		X: (center.X + p2x),
		Y: (center.Y + p2y),
	}
}
