package player

type Player int

var (
	NONE  Player = 0
	WHITE Player = 1
	BLACK Player = 2
)

// Opposite will return the opposite player
func Opposite(p Player) Player {
	if p == WHITE {
		return BLACK
	}
	if p == BLACK {
		return WHITE
	}
	return NONE
}
