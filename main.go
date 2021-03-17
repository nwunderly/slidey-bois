package main

import (
	"github.com/nwunderly/slidey-bois/game"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	g := game.New()

	log.Println(g.Player.Pos())
	log.Println(g.Board.Rows[0][0].Pos())

	log.Fatal(g.Game.Run())
}
