package minimap

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/teohen/FPV/draw"
	"github.com/teohen/FPV/global"
	"github.com/teohen/FPV/player"
	vector2 "github.com/teohen/FPV/vector"
)

var opacity = uint8(40)
var WALL_COLOR = color.RGBA{120, 120, 120, opacity}
var EMPTY_COLOR = color.RGBA{18, 18, 18, opacity}
var BORDER_COLOR = color.RGBA{80, 80, 80, opacity}

var cell = vector2.Vector2{X: global.CELL_W, Y: global.CELL_H}

type MiniMap struct {
	World   [][]uint8
	Player  *player.Player
	sprites []fyne.CanvasObject
}

func NewMiniMap(world [][]uint8) *MiniMap {
	player := player.NewPlayer(
		global.PLAYER_SIZE,
		global.PLAYER_START_POS,
		global.PLAYER_START_DIR,
	)
	minimap := MiniMap{World: world, Player: &player}
	return &minimap
}

func (mm *MiniMap) Render() {
	mm.renderMap()
	mm.renderPlayer()
}

func (mm *MiniMap) renderPlayer() {
	mm.Player.Render()
	for _, canvasObj := range mm.Player.GetSprites() {
		mm.AddSprites(canvasObj)
	}
}

func (mm *MiniMap) renderMap() {
	for i := 0; i < len(mm.World); i += 1 {
		for j := 0; j < len(mm.World[i]); j += 1 {
			color := getColor(mm.World[j][i])
			pos := getPos(i, j)
			mm.renderRect(cell, pos, color, BORDER_COLOR, 0)
		}
	}
}

func (mm *MiniMap) renderRect(size, pos vector2.Vector2, rectColor, borderColor color.Color, borderWidth float64) {
	rect := canvas.NewRectangle(rectColor)
	draw.Rect(rect, size, pos, rectColor, borderColor, borderWidth)
	mm.AddSprites(rect)
}

func getColor(value uint8) color.Color {
	if value != 0 {
		return WALL_COLOR
	} else {
		return EMPTY_COLOR
	}
}

func getPos(idxI, idxJ int) vector2.Vector2 {
	return vector2.Vector2{
		X: global.GetPadX(cell.X * float64(idxI)),
		Y: global.GetPadY(cell.Y * float64(idxJ)),
	}
}
func (mm *MiniMap) AddSprites(obj fyne.CanvasObject) {
	mm.sprites = append(mm.sprites, obj)
}

func (mm *MiniMap) GetSprites() []fyne.CanvasObject {
	return mm.sprites
}
