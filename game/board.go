package game

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

type Board struct {
	Rows [][]string
}

func NewBoard(size, moves int) *Board {
	board := &Board{}

	err := board.Generate(size, moves)

	for err != nil {
		fmt.Printf("ERROR: %s\nRECALCULATING\n", err)
		err = board.Generate(size, moves)
	}
	return board
}

func (board *Board) Show() {
	fmt.Print("\n")
	for _, row := range board.Rows {
		for _, col := range row {
			fmt.Print(" " + col)
		}
		fmt.Print("\n")
	}
}

func (board *Board) FillEdges() {
	size := len(board.Rows)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				board.Rows[i][j] = WALL
			}
		}
	}
}

func (board *Board) RandomEdge(s string) (int, int) {
	size := len(board.Rows)
	var row, col int

	row = rand.Intn(size)
	if row == 0 || row == size-1 {
		col = rand.Intn(size-2) + 1
	} else {
		col = rand.Intn(2) * (size - 1)
	}

	board.Rows[row][col] = s
	return row, col
}

func (board *Board) ReplaceFillerSpaces() {
	size := len(board.Rows)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board.Rows[i][j] == FILLER {
				board.Rows[i][j] = EMPTY
			}
		}
	}
}

func (board *Board) Generate(size int, moves int) error {
	board.Rows = EmptyBoard(size)
	board.FillEdges()

	row, col := board.RandomEdge(PLAYER)

	solution := make([]string, moves)
	var move, last string
	var vx, vy int

	for i := 0; i < moves; i++ {
		options := make([]string, 0)

		if last != "DOWN" && row != 0 && board.Rows[row-1][col] != WALL {
			options = append(options, "UP")
		}
		if last != "UP" && row != size-1 && board.Rows[row+1][col] != WALL {
			options = append(options, "DOWN")
		}
		if last != "RIGHT" && col != 0 && board.Rows[row][col-1] != WALL {
			options = append(options, "LEFT")
		}
		if last != "LEFT" && col != size-1 && board.Rows[row][col+1] != WALL {
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
			if r <= 0 || r >= size-1 || c <= 0 || c >= size-1 || board.Rows[r][c] == WALL {
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

		if s := board.Rows[r][c]; s == FILLER {
			return fmt.Errorf("error moving in space (%d, %d): filled with [%s]", r, c, s)
		}

		if i != moves-1 {
			board.Rows[r][c] = WALL
		} else {
			board.Rows[r][c] = GOAL
		}

		for j := 1; j <= dist; j++ {
			board.Rows[row+(j*vy)][col+(j*vx)] = FILLER
		}

		row += vy * (dist)
		col += vx * (dist)

		fmt.Println(move, maxDist, dist, []int{r, c}, []int{row, col})

		last = move
	}

	board.Show()
	board.ReplaceFillerSpaces()

	return nil
}
