package board

import "errors"

type Board struct {
	Size int
	grid [][]Stone
}

func New(size int) *Board {
	return &Board{
		Size: size,
		grid: gridInit(size),
	}
}

func gridInit(size int) [][]Stone {
	outerlayer := make([][]Stone, size)
	for i := range outerlayer {
		outerlayer[i] = make([]Stone, size)
	}
	return outerlayer
}

func (b *Board) Put(player int, x, y int) error {
	if b.Get(x, y).player != 0 {
		return errors.New("stone already on specified coordinates. stone cannot be placed")
	}
	_x, _y := b.translate(x, y)
	b.grid[_y][_x] = Stone{liberties: 4, player: player}
	return nil
}

func (b *Board) Get(x, y int) Stone {
	_x, _y := b.translate(x, y)
	return b.grid[_y][_x]
}

func (b *Board) translate(x, y int) (int, int) {
	_y := b.Size - y
	_x := x - 1
	return _x, _y
}
