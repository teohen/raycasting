package vector2

import "math"

type Vector2 struct {
	X float64
	Y float64
}

type Vector232 struct {
	X float32
	Y float32
}

func (v2 *Vector2) Multiply(num float64) Vector2 {
	return Vector2{v2.X * num, v2.Y * num}
}

func (v2 *Vector2) SubtracBy(num float64) Vector2 {
	return Vector2{v2.X - num, v2.Y - num}
}

func (v2 *Vector2) SumBy(num float64) Vector2 {
	return Vector2{v2.X + num, v2.Y + num}
}

func (v2 *Vector2) To32() Vector232 {
	return Vector232{
		X: float32(v2.X),
		Y: float32(v2.Y),
	}
}

func (v2 *Vector2) Sum(v Vector2) Vector2 {
	return Vector2{
		X: v2.X + v.X,
		Y: v2.Y + v.Y,
	}
}
func (v2 *Vector2) FromAngle(angle float64) Vector2 {
	return Vector2{X: math.Cos(angle), Y: math.Sin(angle)}
}

func (v *Vector2) Dot(v2 Vector2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{X: v.X - v2.X, Y: v.Y - v2.Y}
}
