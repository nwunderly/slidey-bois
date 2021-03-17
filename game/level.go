package game

import (
	"fmt"
	"log"
	"math/rand"
)

const (
	PlayerSymbol = "o"
	FloorSymbol  = "+"
	WallSymbol   = "X"
	GoalSymbol   = "="
	StartSymbol  = "|"
	FILLER       = "-"
)

var DevTestLevel = &Level{
	Rows: [][]string{
		{"X", "X", "X", "X", "X"},
		{"X", "+", "+", "+", "|"},
		{"X", "+", "+", "+", "X"},
		{"X", "+", "+", "+", "="},
		{"X", "X", "X", "X", "X"},
	},
}

func EmptyLevel(size int) [][]string {
	level := make([][]string, size)

	for i := 0; i < size; i++ {
		row := make([]string, size)
		for j := 0; j < size; j++ {
			row[j] = FloorSymbol
		}
		level[i] = row
	}

	return level
}

type Level struct {
	Rows [][]string
}

func GenerateLevel(size, moves int) *Level {
	level := &Level{}

	err := level.Generate(size, moves)

	for err != nil {
		log.Printf("ERROR: %s\nRECALCULATING\n", err)
		err = level.Generate(size, moves)
	}
	return level
}

func (level *Level) Print() {
	log.Print("\n")
	for _, row := range level.Rows {
		for _, col := range row {
			log.Print(" " + col)
		}
		log.Print("\n")
	}
}

func (level *Level) FillEdges() {
	size := len(level.Rows)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				level.Rows[i][j] = WallSymbol
			}
		}
	}
}

func (level *Level) RandomEdge(s string) (int, int) {
	size := len(level.Rows)
	var row, col int

	row = rand.Intn(size)
	if row == 0 || row == size-1 {
		col = rand.Intn(size-2) + 1
	} else {
		col = rand.Intn(2) * (size - 1)
	}

	level.Rows[row][col] = s
	return row, col
}

func (level *Level) ReplaceFillerSpaces() {
	size := len(level.Rows)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if level.Rows[i][j] == FILLER {
				level.Rows[i][j] = FloorSymbol
			}
		}
	}
}

func (level *Level) Generate(size int, moves int) error {
	level.Rows = EmptyLevel(size)
	level.FillEdges()

	row, col := level.RandomEdge(StartSymbol)

	solution := make([]string, moves)
	var move, last string
	var vx, vy int

	for i := 0; i < moves; i++ {
		options := make([]string, 0)

		if last != "DOWN" && row != 0 && level.Rows[row-1][col] != WallSymbol {
			options = append(options, "UP")
		}
		if last != "UP" && row != size-1 && level.Rows[row+1][col] != WallSymbol {
			options = append(options, "DOWN")
		}
		if last != "RIGHT" && col != 0 && level.Rows[row][col-1] != WallSymbol {
			options = append(options, "LEFT")
		}
		if last != "LEFT" && col != size-1 && level.Rows[row][col+1] != WallSymbol {
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

		maxDist := 0
		r, c := row, col
		for {
			r += vy
			c += vx
			if r <= 0 || r >= size-1 || c <= 0 || c >= size-1 || level.Rows[r][c] == WallSymbol {
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

		if s := level.Rows[r][c]; s == FILLER {
			return fmt.Errorf("error moving in space (%d, %d): filled with [%s]", r, c, s)
		}

		if i != moves-1 {
			level.Rows[r][c] = WallSymbol
		} else {
			level.Rows[r][c] = GoalSymbol
		}

		for j := 1; j <= dist; j++ {
			level.Rows[row+(j*vy)][col+(j*vx)] = FILLER
		}

		row += vy * (dist)
		col += vx * (dist)

		log.Println(move, maxDist, dist, []int{r, c}, []int{row, col})

		last = move
	}

	log.Println(level.Rows)
	level.ReplaceFillerSpaces()

	return nil
}
