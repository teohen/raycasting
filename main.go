package main

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/teohen/FPV/global"
	"github.com/teohen/FPV/player"
	"github.com/teohen/FPV/ray"
	vector2 "github.com/teohen/FPV/vector"
)

const WINDOW_HEIGHT = 800
const WINDOW_WIDTH = 800
const WALL_HEIGHT = WINDOW_HEIGHT * 0.8
const WALL_HEIGHT_MARGIN = WINDOW_HEIGHT * 0.3

var WALL_COLOR = color.RGBA{128, 128, 128, 255}
var EMPTY_COLOR = color.RGBA{18, 18, 18, 255}
var BORDER_COLOR = color.RGBA{80, 80, 80, 255}

const CELL_W = 40
const CELL_H = 40

type MiniMap struct {
	container fyne.Container
	Walls     [][]uint8
}

type GameWindow struct {
	container fyne.Container
	player    player.Player
}

func (mm *MiniMap) NewMiniMap() {
	mm.container = *container.NewWithoutLayout()
}

func (mm *MiniMap) NewWalls(walls [][]uint8) {
	mm.Walls = walls
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
	mm.container.Add(rect)
}

func (mm *MiniMap) drawPlayer(p *player.Player) {
	for _, canvasObj := range p.GetSprites() {
		mm.container.Add(canvasObj)
	}
	p.Render()
}

func (mm *MiniMap) drawCells() {
	for i := 0; i < len(mm.Walls); i += 1 {
		for j := 0; j < len(mm.Walls[i]); j += 1 {
			var cellColor color.RGBA
			if mm.Walls[j][i] != 0 {
				cellColor = WALL_COLOR
			} else {
				cellColor = EMPTY_COLOR
			}

			cell := vector2.Vector2{X: CELL_W, Y: CELL_H}
			cellPos := vector2.Vector2{X: cell.X * float64(i), Y: cell.Y * float64(j)}
			mm.drawRect(cell, cellPos, cellColor, BORDER_COLOR, 2)
		}
	}
}

func renderColumns(lines []*canvas.Line, rays []*ray.Ray, p *player.Player) {

	lw := WINDOW_WIDTH / len(rays)

	for i, line := range lines {

		length := WALL_HEIGHT - rays[i].CalcLength(p.Dir)
		line.Position1.X = float32(450 + (i * lw))
		line.Position1.Y = float32(WALL_HEIGHT-length) * 0.5
		line.Position2.X = float32(450 + (i * lw))
		line.Position2.Y = float32(length)
		r := uint8(30 + length*0.25)
		g := uint8(30 + length*0.25)
		b := uint8(30 + length*0.25)
		line.StrokeColor = color.RGBA{r, g, b, 255}
		line.StrokeWidth = float32(lw)
		line.Refresh()
	}

}

func main() {
	a := app.New()
	w := a.NewWindow("First Person Viewer")

	minimap := MiniMap{}
	walls := global.GetGlobal().World
	player := player.NewPlayer(
		vector2.Vector2{X: 20, Y: 20},
		vector2.Vector2{X: 50, Y: 100},
		(math.Pi)/2,
	)

	minimap.NewMiniMap()
	minimap.NewWalls(walls)
	minimap.drawCells()
	minimap.drawPlayer(&player)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case "W":
			player.Move(5)
		case "S":
			player.Move(-5)
		case "D":
			player.Rotate(math.Pi * 0.01)
		case "A":
			player.Rotate(-math.Pi * 0.01)
		case "Q":
			w.Close()
		}
	})

	var lines []*canvas.Line

	for range player.Rays {
		line := canvas.NewLine(color.White)
		lines = append(lines, line)
		minimap.container.Add(line)
	}

	go func() {
		for range time.Tick(time.Millisecond * 33) {
			player.Refresh()
			renderColumns(lines, player.Rays, &player)
		}
	}()

	w.SetContent(&minimap.container)
	w.Resize(fyne.NewSize(WINDOW_WIDTH*2, WINDOW_HEIGHT))
	w.ShowAndRun()
}
