package robo

import (
	"fmt"
	"os"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	ExitOK    = 0
	ExitError = 1
)

var (
	ScreenWidth  = 512
	ScreenHeight = 512
	FramesPerStep = 30
)

func Main(run func(r *Robo)) int {
	if err := Run(run); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return ExitError
	}
	return ExitOK
}

func Run(run func(r *Robo)) error {
	game := &Game{
		width:  ScreenWidth,
		height: ScreenHeight,
	}

	robo, err := NewRobo(game, FramesPerStep)
	if err != nil {
		return err
	}

	run(robo)
	game.robo = robo

	ebiten.SetWindowSize(game.width, game.height)
	ebiten.SetWindowTitle("Robot")
	if err := ebiten.RunGame(game); err != nil {
		return err
	}

	return nil
}
