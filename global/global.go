package global

import (
	"image/color"

	vector2 "github.com/teohen/FPV/vector"
)

type Global struct {
	World [][]uint8
}

const WINDOW_HEIGHT = 800
const WINDOW_WIDTH = 800
const WALL_HEIGHT = WINDOW_HEIGHT * 0.8
const WALL_HEIGHT_MARGIN = WINDOW_HEIGHT * 0.3

var WALL_COLOR = color.RGBA{128, 128, 128, 255}
var EMPTY_COLOR = color.RGBA{18, 18, 18, 255}
var BORDER_COLOR = color.RGBA{80, 80, 80, 255}

const CELL_W = 40
const CELL_H = 40

var global Global

func NewGlobal() Global {
	global = Global{World: [][]uint8{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
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
	if x > len(global.World[0])-1 || y < 0 {
		return 1
	}
	return global.World[x][y]
}

func InsideGame(p vector2.Vector2) bool {
	if 0 <= p.X && p.X <= WINDOW_WIDTH && 0 <= p.Y && p.Y <= WINDOW_HEIGHT {
		return true
	}
	return false
}
