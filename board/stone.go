package board

type Stone struct {
	liberties int
	player    int
	groupID   int
}

func NewStone(player int, id int) Stone {
	return Stone{
		liberties: 0,
		player:    player,
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
