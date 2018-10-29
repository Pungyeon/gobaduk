package move

import (
	"github.com/Pungyeon/gobaduk/dlgo/point"
)

type Move struct {
	Point  point.Point
	Play   bool
	Pass   bool
	Resign bool
}

func Play(p point.Point) Move {
	return Move{
		Point: p,
	}
}

func Pass() Move {
	return Move{
		Pass: true,
	}
}

func Resign() Move {
	return Move{
		Resign: true,
	}
}
