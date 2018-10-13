package board

import "github.com/Pungyeon/gobaduk/player"

type Stone struct {
	liberties int
	player    player.Player
	groupID   int
}

func NewStone(p player.Player, id int) Stone {
	return Stone{
		liberties: 0,
		player:    p,
		groupID:   id,
	}
}

type Group struct {
	liberties int
	id        int
	stones    []Stone
}

func NewGroup(stone Stone) Group {
	return Group{
		liberties: stone.liberties,
		id:        stone.groupID,
		stones:    []Stone{stone},
	}
}
