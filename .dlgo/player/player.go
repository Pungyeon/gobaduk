package player

type Player int

const (
	NONE  Player = 0
	BLACK Player = 1
	WHITE Player = 2
)

func Other(p Player) Player {
	if p == BLACK {
		return WHITE
	}
	return BLACK
}
