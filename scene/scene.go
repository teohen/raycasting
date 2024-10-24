package scene

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/teohen/FPV/global"
	"github.com/teohen/FPV/player"
)

var PADDING_X = 180
var PADDING_Y = float64(120)

type Scence struct {
	player *player.Player
	lines  []*canvas.Line
}

func (sc *Scence) GetSprites() []fyne.CanvasObject {
	var sprites = make([]fyne.CanvasObject, 0)
	for _, line := range sc.lines {
		sprites = append(sprites, line)
	}
	return sprites
}

func NewScene(p *player.Player) Scence {
	scene := Scence{player: p}

	for range p.Rays {
		scene.lines = append(scene.lines, canvas.NewLine(color.White))
	}
	return scene
}

func (s *Scence) RenderScene() {
	lw := global.WINDOW_WIDTH / len(s.lines)
	for i, line := range s.lines {
		ray := s.player.Rays[i]

		rl := ray.CalcLength(s.player.Dir)
		length := global.WALL_HEIGHT
		if rl > 0 {
			length -= rl
		} else {
			length = rl
		}
		line.Position1.X = float32(int(PADDING_X) + (i * lw))
		line.Position1.Y = float32(PADDING_Y + (global.WALL_HEIGHT-length)*0.5)
		line.Position2.X = float32(int(PADDING_X) + (i * lw))
		line.Position2.Y = float32(PADDING_Y + length)
		var r, g, b, a uint8
		r, g, b, a = 0, 0, 0, 0
		if length > 0 {
			r = uint8(30 + length*0.3)
			g = uint8(30 + length*0.3)
			b = uint8(30 + length*0.3)
			a = uint8(255)
		}
		line.StrokeColor = color.RGBA{r, g, b, a}
		line.StrokeWidth = float32(lw)
	}

	s.Refresh()
}

func (s *Scence) Refresh() {
	for _, line := range s.lines {
		line.Refresh()
	}
}
