package game

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

const (
	width = 500
	height = 500
	tileWidth = 50
	tileHeight = 50
	tilesX = 10
	tilesY = 10
)

type SlideyBois struct {
	Game engine.Game
	Renderer engine.Renderer

	Player Player
	Board *Board
}

func New() *SlideyBois {
	g := &SlideyBois{}
	game := ardent.NewGame(
		"SlideyBois",
		width, height,
		engine.FlagResizable,
		g.Tick, g.Layout,
	)
	g.Setup(game)

	return g
}

func (g *SlideyBois) Tick() {}

func (g *SlideyBois) Layout(w, h int) (int, int) {
	return width, height
}

func (g *SlideyBois) Setup(game engine.Game) {
	g.Game = game
	g.Renderer = game.NewRenderer()
	g.Board = NewBoard(10, 10)
}

// TODO: implement tiles with Board struct

//func (g *SlideyBois) NewTile(size int, col color.Color, symbol string) {
//	t := tile{
//		0, 0,
//		symbol,
//		g.Game.NewImageFromImage(),
//	}
//}