package board

import (
	"errors"

	"github.com/Pungyeon/gobaduk/dlgo/player"
	"github.com/Pungyeon/gobaduk/dlgo/point"
)

type Board struct {
	size int
	grid [][]GoString
}

func New(size int) Board {
	bg := make([][]GoString, size)
	for i := range bg {
		bg[i] = make([]GoString, size)
	}
	return Board{size, bg}
}

func (this *Board) PlaceStone(color player.Player, pt point.Point) error {
	if this.isOnGrid(point) != true {
		return errors.New("specified point is not on board grid")
	}

	if this.get(point) != nil {
		return errors.New("specified point, already has a stone placed on it")
	}

	neighbourAllies = []int{}
	neighbourEnemies = []int{}
	liberties := []int{}

	for _, neighbour := point.Neighbours() {

	}

}

func (this *Board) isOnGrid(pt point.Point) bool {
	return 1 <= pt.Row && pt.Row < this.size && 1 <= pt.Col && pt.Col < this.size
}

func (this *Board) get(pt point.Point) player.Player {
	return this.grid[pt.Row][pt.Col].color
}

func (this *Board) getGoString(pt point.Point) stone.String {
	return this.grid[pt.Row][pt.Col]
}
