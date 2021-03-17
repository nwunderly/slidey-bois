package game

import (
	"github.com/split-cube-studios/ardent/engine"
	"image"
	"image/color"
	"log"
)

var (
	// Blue
	//PlayerColor = color.RGBA{G: 255, B: 255, A: 255}
	PlayerColor = color.RGBA{G: 255, B: 255, A: 100}

	// Gray
	WallColor = color.Gray16{Y: 0xfff}

	// Gold
	GoalColor = color.RGBA{R: 255, G: 255, B: 97, A: 255}

	// White
	FloorColor = color.White

	// Green
	StartColor = color.RGBA{R: 151, G: 231, B: 89, A: 255}
)

type Tile struct {
	X int
	Y int
	symbol string
	Img engine.Image
}

func (t *Tile) Pos() (int, int) {
	return t.X, t.Y
}

func (t *Tile) Symbol() string {
	return t.Symbol()
}

func (t *Tile) TranslateTo(x, y int) {
	t.Img.Translate(float64(x), float64(y))
}

func (t *Tile) TranslateBy(dx, dy int) {
	t.TranslateTo(t.X+dx, t.Y+dy)
}

func (t *Tile) Discard() {
	t.Img.Dispose()
}

func (g *SlideyBois) NewTile(x, y int, s string, c color.Color) *Tile {
	log.Printf("Making tile %s at (%d, %d)\n", s, x, y)
	img := image.NewNRGBA(image.Rect(0, 0, tileWidth, tileHeight))
	for _x := img.Bounds().Min.X; _x < img.Bounds().Max.X; _x++ {
		for _y := img.Bounds().Min.Y; _y < img.Bounds().Max.Y; _y++ {
			img.Set(_x, _y, c)
		}
	}
	tile := &Tile{
		X: x,
		Y: y,
		symbol: s,
		Img: g.Game.NewImageFromImage(img),
	}
	tile.Img.Translate(float64(x), float64(y))
	g.Renderer.AddImage(tile.Img)
	return tile
}

// TODO: maybe make player tile smaller than map tiles?
func (g *SlideyBois) NewPlayerTile(x, y int) *Tile {
	return g.NewTile(x, y, PlayerSymbol, PlayerColor)
	//s := PlayerSymbol
	//c := PlayerColor
	//
	//img := image.NewNRGBA(image.Rect(0, 0, tileWidth/2, tileHeight/2))
	//for _x := img.Bounds().Min.X; _x < img.Bounds().Max.X; _x++ {
	//	for _y := img.Bounds().Min.Y; _y < img.Bounds().Max.Y; _y++ {
	//		img.Set(_x, _y, c)
	//	}
	//}
	//tile := &Tile{
	//	X: x,
	//	Y: y,
	//	symbol: s,
	//	Img: g.Game.NewImageFromImage(img),
	//}
	//tile.Img.Translate(float64(x), float64(y))
	//g.Renderer.AddImage(tile.Img)
	//return tile
}

func (g *SlideyBois) NewWallTile(x, y int) *Tile {
	return g.NewTile(x, y, WallSymbol, WallColor)
}

func (g *SlideyBois) NewGoalTile(x, y int) *Tile {
	return g.NewTile(x, y, GoalSymbol, GoalColor)
}

func (g *SlideyBois) NewFloorTile(x, y int) *Tile {
	return g.NewTile(x, y, FloorSymbol, FloorColor)
}

func (g *SlideyBois) NewStartTile(x, y int) *Tile {
	return g.NewTile(x, y, StartSymbol, StartColor)
}


func (g *SlideyBois) DrawBorder() {
	border := image.NewNRGBA(image.Rect(0, 0, width, height))
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			if x < tileWidth || x > width-tileWidth || y < tileHeight || y > height-tileHeight {
				border.Set(x, y, color.White)
			}
		}
	}

	g.Renderer.AddImage(g.Game.NewImageFromImage(border))
}
