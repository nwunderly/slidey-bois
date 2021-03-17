package old

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

const (
	w, h  = 854, 480

	blockSize = 10

	speed = 5 // ticks per movement
)

var (
	game      engine.Game
	renderer engine.Renderer

	playerImg image.Image
	obstacleImg image.Image

	player Tile
	obstacle Tile

	pxv = -5
	oxv = 5
)

func oldmain() {
	game = ardent.NewGame(
		"Slidey Bois",
		w, h,
		engine.FlagResizable,
		Tick,
		Layout,
	)

	// other stuff with Game
	setup()

	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Tick() {
	if player.X <= 0 || player.X >= w {
		pxv *= -1
	}
	if obstacle.X <= 0 || obstacle.X >= w {
		oxv *= -1
	}

	player.Translate(player.X + pxv, player.Y)
	obstacle.Translate(obstacle.X + oxv, obstacle.Y)
}

func Layout(ow, oh int) (int, int) {
	return w, h
}

func setup() {
	renderer = game.NewRenderer()
	game.AddRenderer(renderer)

	playerImg = createTile(color.White)
	obstacleImg = createTile(color.RGBA{G: 0xff, B: 0xaf, A: 0xff})

	player = Tile{
		0, 0, game.NewImageFromImage(playerImg),
	}
	obstacle = Tile{
		0, 0, game.NewImageFromImage(obstacleImg),
	}

	drawBorder()

	renderer.AddImage(player.img)
	renderer.AddImage(obstacle.img)

	player.Translate(450, 100)
	obstacle.Translate(50, 300)
}

type Tile struct {
	X, Y int
	img  engine.Image
}

func (t *Tile) Translate(x, y int) {
	t.X = x
	t.Y = y
	t.img.Translate(float64(t.X), float64(t.Y))
}

func drawBorder() {
	border := image.NewNRGBA(image.Rect(0, 0, w, h))
	for x := 0; x <= w; x++ {
		for y := 0; y <= h; y++ {
			if x < blockSize || x > w-blockSize || y < blockSize || y > h-blockSize {
				border.Set(x, y, color.White)
			}
		}
	}

	renderer.AddImage(game.NewImageFromImage(border))
}

func createTile(c color.Color) image.Image {
	tile := image.NewNRGBA(image.Rect(0, 0, blockSize, blockSize))
	fillImage(tile, c)
	return tile
}

func fillImage(image draw.Image, color color.Color) {
	for x := image.Bounds().Min.X; x < image.Bounds().Max.X; x++ {
		for y := image.Bounds().Min.Y; y < image.Bounds().Max.Y; y++ {
			image.Set(x, y, color)
		}
	}
}
