package board

import (
	"github.com/Pungyeon/gobaduk/player"
)

var (
	color = []string{"N", "W", "B"}
	abc   = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s"}
)

type Stone struct {
	player player.Player
	group  *Group
	x      int
	y      int
	_x     int
	_y     int
}

func NewStone(p player.Player, x, y, _x, _y int) Stone {
	return Stone{
		player: p,
		group:  nil,
		x:      x,
		y:      y,
		_x:     _x,
		_y:     _y,
	}
}

func (s *Stone) Player() string {
	return color[s.player]
}

func (s *Stone) X() string {
	return abc[s._x]
}

func (s *Stone) Y() string {
	return abc[s._y]
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
