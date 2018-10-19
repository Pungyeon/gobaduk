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
		player: player.BLACK,
		board:  board.New(size),
	}
}

// Move will place a move on the go board
func (g *Game) Move(x, y int) (*Game, error) {
	if err := g.board.Put(g.player, x, y); err != nil {
		return nil, err
	}
	g.changePlayer()
	return g, nil
}

func (g *Game) unsafeMove(x, y int) *Game {
	g.board.Put(g.player, x, y)
	g.changePlayer()
	return g
}

func (g *Game) changePlayer() {
	g.player = player.Opposite(g.player)
}

func (g *Game) SGF() string {
	return g.board.SGF.String()
}
