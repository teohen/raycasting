package global

import (
	vector2 "github.com/teohen/FPV/vector"
)

type Global struct {
	World [][]uint8
}

const WINDOW_HEIGHT = 800
const WINDOW_WIDTH = 800

const WALL_HEIGHT = WINDOW_HEIGHT * 0.9
const WALL_HEIGHT_MARGIN = WINDOW_HEIGHT * 0.3

const CELL_W = 40
const CELL_H = 40

var PADDING_X = float64(0)
var PADDING_Y = float64(0)

var global Global

func NewGlobal() Global {
	global = Global{World: [][]uint8{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}}

	return global
}

func GetGlobal() Global {
	if global.World == nil {
		global = NewGlobal()
	}
	return global
}

func GetWorldCellContent(x, y int) uint8 {
	//TODO REMOVE HARDCODED
	if x > 9 || y > 9 || x < 0 || y < 0 {
		return 1
	}
	return global.World[x][y]
}

func InsideGame(p vector2.Vector2) bool {
	if p.X == 0 || p.X == float64(len(global.World)*CELL_W) {
		return false
	}

	if p.Y == 0 || p.Y == float64(len(global.World)*CELL_H) {
		return false
	}

	return true
}

func GetPadX(x float64) float64 {
	return x + PADDING_X
}

func GetPadY(y float64) float64 {
	return y + PADDING_Y
}
