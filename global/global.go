package global

import (
	"math"

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

const RAYS_COUNT = 100
const CONE_ANGLE = math.Pi / 2

var PLAYER_START_POS = vector2.Vector2{X: 40, Y: 100}
var PLAYER_SIZE = vector2.Vector2{X: 20, Y: 20}
var PLAYER_START_DIR = (math.Pi) * 2

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

func GetWorld() [][]uint8 {
	if global.World == nil {
		global = NewGlobal()
	}
	return global.World
}

func GetWorldCellContent(x, y int) uint8 {
	max := len(global.World) - 1
	if x > max || y > max || x < 0 || y < 0 {
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

func GetXY(pos, origin vector2.Vector2) (int, int) {
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
