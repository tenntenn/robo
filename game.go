package robo

import (
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	robo    *Robo
	width   int
	height  int
	started bool
}

func (g *Game) Update() error {	

	var keys []ebiten.Key
	keys = inpututil.AppendPressedKeys(keys)	

	if len(keys) == 1 && keys[0] == ebiten.KeySpace {
		if g.robo.finished() && g.started {
			g.robo.init()
			g.started = false
		} else {
			g.started = true
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	g.robo.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
