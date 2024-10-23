package minimap

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/teohen/FPV/global"
	"github.com/teohen/FPV/player"
	vector2 "github.com/teohen/FPV/vector"
)

var opacity = uint8(40)
var WALL_COLOR = color.RGBA{120, 120, 120, opacity}
var EMPTY_COLOR = color.RGBA{18, 18, 18, opacity}
var BORDER_COLOR = color.RGBA{80, 80, 80, opacity}

const CELL_W = 40
const CELL_H = 40

type MiniMap struct {
	World   [][]uint8
	Player  *player.Player
	sprites []fyne.CanvasObject
}

func NewMiniMap(world [][]uint8) *MiniMap {
	player := player.NewPlayer(
		vector2.Vector2{X: 20, Y: 20},
		vector2.Vector2{X: global.GetPadX(40), Y: global.GetPadY(100)},
		(math.Pi)/2,
	)
	minimap := MiniMap{World: world, Player: &player}
	return &minimap
}

func (mm *MiniMap) Draw() {
	mm.drawCells()
	mm.drawPlayer()
}

func (mm *MiniMap) drawPlayer() {
	for _, canvasObj := range mm.Player.GetSprites() {
		mm.AddSprites(canvasObj)
	}
	mm.Player.Render()
}

func (mm *MiniMap) drawCells() {
	cell := vector2.Vector2{X: CELL_W, Y: CELL_H}
	for i := 0; i < len(mm.World); i += 1 {
		for j := 0; j < len(mm.World[i]); j += 1 {
			var cellColor color.RGBA
			if mm.World[j][i] != 0 {
				cellColor = WALL_COLOR
			} else {
				cellColor = EMPTY_COLOR
			}
			cellPos := vector2.Vector2{
				X: global.GetPadX(cell.X * float64(i)),
				Y: global.GetPadY(cell.Y * float64(j)),
			}
			mm.drawRect(cell, cellPos, cellColor, BORDER_COLOR, 2)
		}
	}
}
func (mm *MiniMap) drawRect(size, pos vector2.Vector2, rectColor, borderColor color.Color, borderWidth float64) {
	if borderWidth > 0 {
		mm.drawRect(size, pos, borderColor, borderColor, 0)
		size = size.SubtracBy(borderWidth)
		pos = pos.SumBy(borderWidth / 2)
	}
	rect := canvas.NewRectangle(rectColor)

	rect.Resize(fyne.NewSize(float32(size.X), float32(size.Y)))
	rect.Move(fyne.NewPos(float32(pos.X), float32(pos.Y)))
	mm.AddSprites(rect)
}

func (mm *MiniMap) AddSprites(obj fyne.CanvasObject) {
	mm.sprites = append(mm.sprites, obj)
}

func (mm *MiniMap) GetSprites() []fyne.CanvasObject {
	return mm.sprites
}
