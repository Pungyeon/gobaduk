package board

import (
	"github.com/Pungyeon/gobaduk/player"
)

type Stone struct {
	player player.Player
	group  *Group
	x      int
	y      int
}

func NewStone(p player.Player, x, y int) Stone {
	return Stone{
		player: p,
		group:  nil,
		x:      x,
		y:      y,
	}
}

type Group struct {
	liberties int
	id        int
	stones    []Stone
}

func NewGroup(id int) Group {
	return Group{
		liberties: 0,
		id:        id,
		stones:    []Stone{},
	}
}

func (g *Group) Add(stone Stone) {
	g.stones = append(g.stones, stone)
}

func (g *Group) Merge(mergeGroup *Group) {
	g.liberties += mergeGroup.liberties - 1
	g.stones = append(g.stones, mergeGroup.stones...)
	for _, stone := range g.stones {
		stone.group = g
	}
}
