package game

import "github.com/Pungyeon/gobaduk/board"

type player int

var (
	none  player
	white player = 1
	black player = 2
)

// Game is an object containing the rules of Go
type Game struct {
	player
	board *board.Board
}

// New returns a new game of go
func New(size int) *Game {
	return &Game{
		player: white,
		board:  board.New(size),
	}
}

// Move will place a move on the go board
func (g *Game) Move() *Game {
	g.changePlayer()
	return g
}

func (g *Game) changePlayer() {
	if g.player == white {
		g.player = black
	} else {
		g.player = white
	}
}
