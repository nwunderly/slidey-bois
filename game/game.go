package game

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
	"log"
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

	Player *Tile
	Board *Board
}

func New() *SlideyBois {
	log.Println("Initializing")

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

func (g *SlideyBois) Tick() {
	//log.Println("Tick")
}

func (g *SlideyBois) Layout(w, h int) (int, int) {
	return width, height
}

func (g *SlideyBois) Setup(game engine.Game) {
	log.Println("Setting up")

	g.Game = game
	g.Renderer = game.NewRenderer()
	g.Game.AddRenderer(g.Renderer)
	//level := GenerateLevel(8, 6)
	level := DevTestLevel
	g.Board = g.NewBoard(level)
}
