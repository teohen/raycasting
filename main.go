package main

import (
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/teohen/FPV/global"
	"github.com/teohen/FPV/minimap"
	"github.com/teohen/FPV/player"
	"github.com/teohen/FPV/scene"
)

type GameWindow struct {
	container fyne.Container
	player    player.Player
}

func main() {
	a := app.New()
	w := a.NewWindow("First Person Viewer")

	minimap := minimap.NewMiniMap(global.GetWorld())
	sc := scene.NewScene(minimap.Player)
	minimap.Render()

	cont := container.NewWithoutLayout()
	for _, sprite := range sc.GetSprites() {
		cont.Add(sprite)
	}
	for _, sprite := range minimap.GetSprites() {
		cont.Add(sprite)
	}

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case "W":
			minimap.Player.Move(5)
		case "S":
			minimap.Player.Move(-5)
		case "D":
			minimap.Player.Rotate(math.Pi * 0.01)
		case "A":
			minimap.Player.Rotate(-math.Pi * 0.01)
		case "Q":
			w.Close()
		}
	})

	go func() {
		for range time.Tick((time.Millisecond * 33 * 1) / 2) {
			go minimap.Player.Refresh()
			sc.RenderScene()
		}
	}()

	w.SetContent(cont)
	w.Resize(fyne.NewSize(global.WINDOW_WIDTH, global.WINDOW_HEIGHT))
	w.ShowAndRun()
}
