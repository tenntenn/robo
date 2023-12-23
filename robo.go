package robo

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

//go:embed img/robo.png
var roboImg []byte

type Robo struct {
	game          *Game
	x0, y0        float64
	x, y          float64
	img           *ebiten.Image
	frames        int
	framesPerStep int
	step          int
	steps         []step
}

type step struct{ x, y float64 }

func NewRobo(game *Game, framesPerStep int) (*Robo, error) {
	r := bytes.NewReader(roboImg)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	eimg := ebiten.NewImageFromImage(img)
	w, h := eimg.Size()

	x0 := float64(game.width-w) / 2
	y0 := float64(game.height-h) / 2

	robo := &Robo{
		game:          game,
		x0:            x0,
		y0:            y0,
		img:           eimg,
		framesPerStep: framesPerStep,
	}

	robo.init()

	return robo, nil
}

func (r *Robo) init() {
	r.frames = 0
	r.x = r.x0
	r.y = r.y0
	r.step = 0
}

func (r *Robo) Draw(screen *ebiten.Image) {
	if r.game.started {
		r.doStep()
	}
	var op ebiten.DrawImageOptions
	op.GeoM.Translate(r.x, r.y)
	screen.DrawImage(r.img, &op)
}

func (r *Robo) doStep() {

	if r.finished() {
		return
	}

	r.frames++
	if r.frames%r.framesPerStep != 0 {
		return
	}

	s := r.steps[r.step]
	r.move(s.x, s.y)
	r.step++
}

func (r *Robo) finished() bool {
	return r.step >= len(r.steps)
}

func (r *Robo) MoveUp() {
	_, h := r.img.Size()
	r.steps = append(r.steps, step{0, -float64(h)})
}

func (r *Robo) U() {
	r.MoveUp()
}

func (r *Robo) MoveDown() {
	_, h := r.img.Size()
	r.steps = append(r.steps, step{0, float64(h)})
}

func (r *Robo) D() {
	r.MoveDown()
}

func (r *Robo) MoveRight() {
	w, _ := r.img.Size()
	r.steps = append(r.steps, step{float64(w), 0})
}

func (r *Robo) R() {
	r.MoveRight()
}

func (r *Robo) MoveLeft() {
	w, _ := r.img.Size()
	r.steps = append(r.steps, step{-float64(w), 0})
}

func (r *Robo) L() {
	r.MoveLeft()
}

func (r *Robo) Sleep(steps int) {
	for i := 0; i < steps; i++ {
		r.steps = append(r.steps, step{0, 0})
	}
}

func (r *Robo) S() {
	r.Sleep(1)
}

func (r *Robo) move(dx, dy float64) {
	r.x += dx
	switch {
	case r.x > float64(r.game.width):
		r.x = 0
	case r.x < 0:
		r.x = float64(r.game.width)
	}

	r.y += dy
	switch {
	case r.y > float64(r.game.height):
		r.y = 0
	case r.y < 0:
		r.y = float64(r.game.height)
	}
}
