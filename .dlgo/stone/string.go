package stone

import (
	"errors"

	"github.com/Pungyeon/gobaduk/dlgo/player"
)

type String struct {
	color     player.Player
	stones    []int
	liberties []int
}

func (this *String) RemoveLiberty(point int) {
	this.liberties = append(this.liberties[:point], this.liberties[point+1:]...)
}

func (this *String) AddLiberty(point int) {
	this.liberties = append(this.liberties, point)
}

func (this *String) MergeWith(gs String) (String, error) {
	if this.color != gs.color {
		return errors.New("cannot merge two Strings of different color #racism")
	}

	return String{
		color:     this.color,
		stones:    append(this.stones, gs.stones...),
		liberties: []int{}, // no idea how this is done in golang
	}
}

func (this *String) Liberties() int {
	return len(this.liberties)
}

func (this *String) Equal(gs String) bool {
	// no idea how this is done in golang
}
