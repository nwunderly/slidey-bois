package misc

import (
	"fmt"
	"math/rand"
)

const (
	PLAYER = "o"
	EMPTY  = "+"
	WALL   = "X"
	GOAL   = "="
	FILLER = "-"
)

func EmptyBoard(size int) [][]string {
	board := make([][]string, size)

	for i := 0; i < size; i++ {
		row := make([]string, size)
		for j := 0; j < size; j++ {
			row[j] = EMPTY
		}
		board[i] = row
	}

	return board
}

type Game struct {
	Board [][]string
}

func NewGame(size, moves int) *Game {
	game := &Game{}

	err := game.GenerateBoard(size, moves)

	for err != nil {
		fmt.Printf("ERROR: %s\nRECALCULATING\n", err)
		err = game.GenerateBoard(size, moves)
	}
	return game
}

func (game *Game) Show() {
	fmt.Print("\n")
	for _, row := range game.Board {
		for _, col := range row {
			fmt.Print(" " + col)
		}
		fmt.Print("\n")
	}
}

func (game *Game) FillEdges() {
	size := len(game.Board)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				game.Board[i][j] = WALL
			}
		}
	}
}

func (game *Game) RandomEdge(s string) (int, int) {
	size := len(game.Board)
	var row, col int

	row = rand.Intn(size)
	if row == 0 || row == size-1 {
		col = rand.Intn(size-2) + 1
	} else {
		col = rand.Intn(2) * (size - 1)
	}

	game.Board[row][col] = s
	return row, col
}

func (game *Game) ReplaceFillerSpaces() {
	size := len(game.Board)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if game.Board[i][j] == FILLER {
				game.Board[i][j] = EMPTY
			}
		}
	}
}

func (game *Game) GenerateBoard(size, moves int) error {
	game.Board = EmptyBoard(size)
	game.FillEdges()

	row, col := game.RandomEdge(PLAYER)

	solution := make([]string, moves)
	var move, last string
	var vx, vy int

	for i := 0; i < moves; i++ {
		options := make([]string, 0)

		if last != "DOWN" && row != 0 && game.Board[row-1][col] != WALL {
			options = append(options, "UP")
		}
		if last != "UP" && row != size-1 && game.Board[row+1][col] != WALL {
			options = append(options, "DOWN")
		}
		if last != "RIGHT" && col != 0 && game.Board[row][col-1] != WALL {
			options = append(options, "LEFT")
		}
		if last != "LEFT" && col != size-1 && game.Board[row][col+1] != WALL {
			options = append(options, "RIGHT")
		}

		if l := len(options); l == 0 {
			return fmt.Errorf("no move options at (%d, %d)", row, col)
		}

		move = options[rand.Intn(len(options))]

		switch move {
		case "UP":
			vx, vy = 0, -1
		case "DOWN":
			vx, vy = 0, 1
		case "LEFT":
			vx, vy = -1, 0
		case "RIGHT":
			vx, vy = 1, 0
		}

		solution[i] = move

		var maxDist int
		r, c := row, col
		for {
			r += vy
			c += vx
			if r <= 0 || r >= size-1 || c <= 0 || c >= size-1 || game.Board[r][c] == WALL {
				break
			}
			maxDist++
		}

		var dist int
		if i != moves-1 {
			dist = rand.Intn(maxDist) + 1
		} else {
			dist = maxDist
		}
		r = row + vy*(dist+1)
		c = col + vx*(dist+1)

		if s := game.Board[r][c]; s == FILLER {
			return fmt.Errorf("error moving in space (%d, %d): filled with [%s]", r, c, s)
		}

		if i != moves-1 {
			game.Board[r][c] = WALL
		} else {
			game.Board[r][c] = GOAL
		}

		for j := 1; j <= dist; j++ {
			game.Board[row+(j*vy)][col+(j*vx)] = FILLER
		}

		row += vy * (dist)
		col += vx * (dist)

		fmt.Println(move, maxDist, dist, []int{r, c}, []int{row, col})

		last = move
	}

	game.Show()
	game.ReplaceFillerSpaces()

	return nil
}