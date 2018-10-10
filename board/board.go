package board

type Board struct {
	Size int
	grid [][]int
}

func New(size int) *Board {
	return &Board{
		Size: size,
		grid: gridInit(size),
	}
}

func gridInit(size int) [][]int {
	outerlayer := make([][]int, size)
	for i := range outerlayer {
		outerlayer[i] = make([]int, size)
	}
	return outerlayer
}

func (b *Board) Put(player int, x, y int) {
	_x, _y := b.translate(x, y)
	b.grid[_y][_x] = player
}

func (b *Board) Get(x, y int) int {
	_x, _y := b.translate(x, y)
	return b.grid[_y][_x]
}

func (b *Board) translate(x, y int) (int, int) {
	_y := b.Size - y
	_x := x - 1
	return _x, _y
}
