package game

import (
	"github.com/Pungyeon/gobaduk/board"
	"github.com/Pungyeon/gobaduk/player"
)

// Game is an object containing the rules of Go
type Game struct {
	player player.Player
	board  *board.Board
}

// New returns a new game of go
func New(size int) *Game {
	return &Game{
		player: player.WHITE,
		board:  board.New(size),
	}
}

// Move will place a move on the go board
func (g *Game) Move() *Game {
	g.changePlayer()
	return g
}

func (g *Game) changePlayer() {
	g.player = player.Opposite(g.player)
}
