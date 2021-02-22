package game

import (
	"github.com/split-cube-studios/ardent/engine"
)

type Tile interface {
	Pos() (int, int)
	TranslateTo(int, int)
	TranslateBy(int, int)
	Discard()
}

type tile struct {
	X int
	Y int
	Symbol string
	Img engine.Image
}

func (t *tile) Pos() (int, int) {
	return t.X, t.Y
}

func (t *tile) TranslateTo(x, y int) {
	t.Img.Translate(float64(x), float64(y))
}

func (t *tile) TranslateBy(dx, dy int) {
	t.TranslateTo(t.X+dx, t.Y+dy)
}

type Player struct {
	Tile
	*tile
}

type Wall struct {
	Tile
	*tile
}