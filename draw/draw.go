package draw

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	vector2 "github.com/teohen/FPV/vector"
)

func Circle(circle *canvas.Circle, size, pos vector2.Vector2, color color.Color) {
	circle.Resize(fyne.NewSize(size.To32().X, size.To32().Y))
	to := *pos.Sum(&size)
	circle.Position1 = fyne.Position(pos.To32())
	circle.Position2 = fyne.Position(to.To32())
	circle.FillColor = color
}

func Line(l *canvas.Line, from, to vector2.Vector2, color color.Color, width int) {
	l.Position1 = fyne.Position(from.To32())
	l.Position2 = fyne.Position(to.To32())
	l.StrokeColor = color
	l.StrokeWidth = float32(width)
}

func Rect(r *canvas.Rectangle, size, pos vector2.Vector2, rectColor, borderColor color.Color, borderWidth float64) {
	if borderWidth > 0 {
		Rect(r, size, pos, borderColor, borderColor, 0)
		size = *size.SubtracBy(borderWidth)
		pos = *pos.SumBy(borderWidth / 2)
	}

	r.Resize(fyne.NewSize(float32(size.X), float32(size.Y)))
	r.Move(fyne.NewPos(float32(pos.X), float32(pos.Y)))
}
