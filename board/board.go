package board

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/Pungyeon/gobaduk/player"
)

type Board struct {
	Size              int
	grid              [][]Stone
	groups            map[int]*Group
	nextID            int
	mergeGroups       []int
	subtractLibGroups []int
	activeKO          Stone
	SGF               SGF
	AI                bool
}

func New(size int) *Board {
	return &Board{
		Size:           size,
		grid:           gridInit(size),
		groups:         map[int]*Group{},
		nextID:         1,
		activeKO:       NewStone(player.NONE, size*2, size*2, size*2, size*2),
		SGF:            NewSGF(),
		AI:             false,
}

func gridInit(size int) [][]Stone {
	outerlayer := make([][]Stone, size)
	for i := range outerlayer {
		outerlayer[i] = make([]Stone, size)
	}
	return outerlayer
}

func (b *Board) getNextID() int {
	val := b.nextID
	b.nextID++
	return val
}

func (b *Board) Put(playerColor player.Player, x, y int) error {
	check := b.Get(x, y)
	if check.player != player.NONE {
		return errors.New("stone already on specified coordinates. stone cannot be placed")
	}

	if check.x == b.activeKO.x && check.y == b.activeKO.y && playerColor == b.activeKO.player {
		return errors.New("cannot place stone, on active KO")
	}

	_x, _y := b.translate(x, y)
	ng := NewGroup(b.getNextID())
	stone := Stone{
		player: playerColor,
		x:      x,
		y:      y,
		_x:     _x,
		_y:     _y,
		group:  &ng,
	}
	ng.Add(stone)

	b.mergeGroups = make([]int, 0)
	b.subtractLibGroups = make([]int, 0)

	up := b.CheckNeighbours(x, y+1, &stone)
	down := b.CheckNeighbours(x, y-1, &stone)
	right := b.CheckNeighbours(x+1, y, &stone)
	left := b.CheckNeighbours(x-1, y, &stone)

	if up && down && right && left && b.AI {
		if b.checkEye(&stone) {
			return errors.New("should not place a stone in liberty")
		}
	}

	b.groups[ng.id] = &ng

	for _, id := range b.mergeGroups {
		b.Merge(&ng, b.groups[id])
	}

	stonesToRemove := make([]*Group, 0)
	for _, key := range b.subtractLibGroups {
		tmp := b.groups[key]
		tmp.liberties--

		if tmp.liberties == 0 {
			stonesToRemove = append(stonesToRemove, tmp)
		}
		b.groups[key] = tmp
	}
	if ng.liberties == 0 && len(stonesToRemove) == 0 {
		return errors.New("cannot place stone without liberties, that doesn't capture any stones")
	}

	b.grid[_y][_x] = stone
	b.SGF.Move(stone)

	for _, rmGroup := range stonesToRemove {
		b.removeStones(rmGroup.stones)
	}

	b.activeKO = NewStone(player.NONE, b.Size*2, b.Size*2, b.Size*2, b.Size*2)
	if b.grid[_y][_x].group.liberties == 1 &&
		len(b.grid[_y][_x].group.stones) == 1 &&
		len(stonesToRemove) == 1 {
		if len(stonesToRemove[0].stones) == 1 {
			// activate KO,
			// It feels like there is a simpler way of defining a KO...
			fmt.Println(stonesToRemove[0].stones[0])
			b.activeKO = stonesToRemove[0].stones[0]
		}
	}

	return nil
}

func (b *Board) removeStones(stones []Stone) {
	for _, stone := range stones {
		_x, _y := b.translate(stone.x, stone.y)
		b.grid[_y][_x] = NewStone(player.NONE, stone.x, stone.y, _x, _y)
		fmt.Println(b.grid[_y][_x])
		if stone.y < b.Size {
			b.addLibertyIfOppositePlayer(
				stone.player, b.Get(stone.x, stone.y+1),
			) // up
		}
		if stone.y > 1 {
			b.addLibertyIfOppositePlayer(
				stone.player, b.Get(stone.x, stone.y-1),
			) // down
		}
		if stone.x < b.Size {
			b.addLibertyIfOppositePlayer(
				stone.player, b.Get(stone.x+1, stone.y),
			) // right
		}
		if stone.x > 1 {
			b.addLibertyIfOppositePlayer(
				stone.player, b.Get(stone.x-1, stone.y),
			) // left
		}
	}
}

func (b *Board) addLibertyIfOppositePlayer(p player.Player, stone *Stone) {
	if stone.player == player.Opposite(p) {
		stone.group.liberties++
	}
}

func (b *Board) Merge(group *Group, mergeGroup *Group) {
	if group.id != mergeGroup.id {
		group.Merge(mergeGroup)
		for _, stone := range mergeGroup.stones {
			ptr := b.Get(stone.x, stone.y)
			ptr.group = group
		}
		delete(b.groups, mergeGroup.id)
	}
}

func (b *Board) CheckNeighbours(x, y int, stone *Stone) bool {
	fmt.Printf("x: %d, y: %d\n", x, y)
	if x < 1 || x > b.Size || y < 1 || y > b.Size {
		return true // check if out of bounds
	}
	neighbour := b.Get(x, y)

	if neighbour.player == stone.player {
		b.mergeGroups = append(b.mergeGroups, neighbour.group.id)
		return true
	}
	if neighbour.player == player.NONE {
		stone.group.liberties++
	}
	if neighbour.player == player.Opposite(stone.player) {
		b.subtractLibGroups = append(b.subtractLibGroups, neighbour.group.id)
	}
	return false
}

func (b *Board) checkEye(stone *Stone) bool {
	diagonalNeighbors := []*Stone{b.Get(stone.x+1, stone.y+1), // upper right
		b.Get(stone.x-1, stone.y+1), // upper left
		b.Get(stone.x+1, stone.y-1), // bottom right
		b.Get(stone.x-1, stone.y-1), // bottom left
	}

	total := 0
	for _, nb := range diagonalNeighbors {
		fmt.Println(nb)
		if nb.player == stone.player { // || nb.player == player.NONE {
			total++
		}
	}

	fmt.Println("total:", total)
	return total >= 3
}

func (b *Board) Get(x, y int) *Stone {
	_x, _y := b.translate(x, y)
	return &b.grid[_y][_x]
}

func (b *Board) translate(x, y int) (int, int) {
	_y := b.Size - y
	_x := x - 1
	return _x, _y
}

func (b *Board) Print() {
	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			fmt.Printf("%d ", b.grid[y][x].player)
		}
		fmt.Println()
	}
}

type SGF struct {
	sequence *bytes.Buffer
}

func NewSGF() SGF {
	return SGF{
		sequence: bytes.NewBufferString("("),
	}
}

func (s *SGF) Move(stone Stone) {
	s.sequence.WriteString(";")
	s.sequence.WriteString(stone.Player())
	s.sequence.WriteString("[")
	s.sequence.WriteString(stone.X())
	s.sequence.WriteString(stone.Y())
	s.sequence.WriteString("]")
}

func (s *SGF) String() string {
	tmp := s.sequence.String()
	return tmp + ")"
}
