package ray

import vector2 "github.com/teohen/FPV/vector"

type Wall struct {
	from vector2.Vector2
	to   vector2.Vector2
}

func NewWall(to, from vector2.Vector2) *Wall {
	return &Wall{to: to, from: from}
}
