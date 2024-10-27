package scene

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"github.com/teohen/FPV/draw"
	"github.com/teohen/FPV/global"
	vector2 "github.com/teohen/FPV/vector"
)

type SceneLine struct {
	from   vector2.Vector2
	to     vector2.Vector2
	color  color.Color
	width  int
	sprite *canvas.Line
}

func NewSceneLine(from, to vector2.Vector2, color color.Color, w int) *SceneLine {
	sprite := canvas.NewLine(color)
	sprite.StrokeWidth = float32(w)
	return &SceneLine{
		from:   from,
		to:     to,
		color:  color,
		width:  w,
		sprite: canvas.NewLine(color),
	}
}

func (sl *SceneLine) render() {
	draw.Line(sl.sprite, sl.from, sl.to, sl.color, sl.width)
}

func (sl *SceneLine) GetSprite() *canvas.Line {
	return sl.sprite
}

func (sl *SceneLine) setColor(distance float64) {
	var r, g, b, a uint8
	r, g, b, a = 0, 0, 0, 0
	if distance > 0 {
		r = uint8(30 + distance*0.3)
		g = uint8(30 + distance*0.3)
		b = uint8(30 + distance*0.3)
		a = uint8(255)
	}
	sl.color = color.RGBA{r, g, b, a}
}

func (sl *SceneLine) setFrom(lineLength float64, i, lineW int) {
	sl.from.X = global.PADDING_X + float64(i*lineW)
	sl.from.Y = global.PADDING_Y + (global.WALL_HEIGHT-lineLength)*0.5
}

func (sl *SceneLine) setTo(lineLength float64, i, lineW int) {
	sl.to.X = global.PADDING_X + float64(i*lineW)
	sl.to.Y = global.PADDING_Y + lineLength
}
func (sl *SceneLine) setWidth(w int) {
	sl.width = w
}
