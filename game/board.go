package game

import "log"

type Board struct {
	Rows [][]*Tile
}

func (g *SlideyBois) NewBoard(level *Level) *Board {
	log.Println("Setting up board")

	rowCount := len(level.Rows)
	colCount := len(level.Rows[0])

	log.Printf("Board dimensions: %d, %d\n", rowCount, colCount)

	board := make([][]*Tile, rowCount)

	var startX, startY int

	for i := 0; i < rowCount; i++ {
		row := make([]*Tile, colCount)
		y := tileHeight/2 + tileHeight*i

		for j := 0; j < colCount; j++ {
			x := tileWidth/2 + tileWidth*j

			switch level.Rows[i][j] {
			case PlayerSymbol:
			case StartSymbol:
				row[j] = g.NewStartTile(x, y)
				startX, startY = x, y
			case FloorSymbol:
				row[j] = g.NewFloorTile(x, y)
			case WallSymbol:
				row[j] = g.NewWallTile(x, y)
			case GoalSymbol:
				row[j] = g.NewGoalTile(x, y)
			default:
				panic("found unexpected symbol attempting to convert level to board")
			}

		}
		board[i] = row
	}
	g.Player = g.NewPlayerTile(startX, startY)

	log.Printf("Board generated:\n%v\n", board)

	return &Board{board}
}