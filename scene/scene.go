package scene

import (
	"image/color"

	"fyne.io/fyne/v2"
	"github.com/teohen/FPV/global"
	"github.com/teohen/FPV/player"
	vector2 "github.com/teohen/FPV/vector"
)

type Scence struct {
	player  *player.Player
	stripes []*SceneLine
}

func (sc *Scence) GetSprites() []fyne.CanvasObject {
	var sprites = make([]fyne.CanvasObject, 0)
	for _, line := range sc.stripes {
		sprites = append(sprites, line.GetSprite())
	}
	return sprites
}

func NewScene(p *player.Player) Scence {
	scene := Scence{player: p}
	v := vector2.Vector2{}

	for range p.Rays {
		scene.stripes = append(scene.stripes, NewSceneLine(v, v, color.White, 1))
	}

	return scene
}

func calcLength(rayLenght float64) float64 {
	if rayLenght > 0 {
		return global.WALL_HEIGHT - rayLenght
	}
	return rayLenght
}

func (s *Scence) RenderScene() {
	lw := global.WINDOW_WIDTH / len(s.stripes)
	for i, line := range s.stripes {
		ray := s.player.Rays[i]
		length := calcLength(ray.CalcLength(s.player.Dir))
		line.setFrom(length, i, lw)
		line.setTo(length, i, lw)
		line.setColor(length)
		line.setWidth(lw)
		line.render()
	}
	s.Refresh()
}

func (s *Scence) Refresh() {
	for _, line := range s.stripes {
		line.GetSprite().Refresh()
	}
}
