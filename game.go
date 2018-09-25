package game

type player int

var (
	none  player
	white player = 1
	black player = 2
)

// Game is an object containing the rules of Go
type Game struct {
	player
}

// New returns a new game of go
func New() *Game {
	return &Game{
		player: white,
	}
}

// Move will place a move on the go board
func (g *Game) Move() {
	g.changePlayer()
}

func (g *Game) changePlayer() {
	if g.player == white {
		g.player = black
	} else {
		g.player = white
	}
}
