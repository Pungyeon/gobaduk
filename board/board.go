package board

import (
	"errors"
	"fmt"

	"github.com/Pungyeon/gobaduk/player"
)

type Board struct {
	Size              int
	grid              [][]Stone
	groups            map[int]Group
	nextID            int
	mergeGroups       []int
	subtractLibGroups map[int]bool
}

func New(size int) *Board {
	return &Board{
		Size:   size,
		grid:   gridInit(size),
		groups: map[int]Group{},
	}
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
	if b.Get(x, y).player != player.NONE {
		return errors.New("stone already on specified coordinates. stone cannot be placed")
	}
	stone := NewStone(playerColor, b.getNextID())

	b.mergeGroups = make([]int, 0)
	b.subtractLibGroups = map[int]bool{}

	if y < b.Size {
		b.CheckNeighbours(b.Get(x, y+1), &stone)
	}
	if y > 1 {
		b.CheckNeighbours(b.Get(x, y-1), &stone)
	}
	if x < b.Size {
		b.CheckNeighbours(b.Get(x+1, y), &stone)
	}
	if x > 1 {
		b.CheckNeighbours(b.Get(x-1, y), &stone)
	}

	currentGroup := NewGroup(stone)
	if len(b.mergeGroups) == 0 {
		b.groups[currentGroup.id] = currentGroup
	}

	for _, id := range b.mergeGroups {
		b.Merge(currentGroup, b.groups[id])
	}

	for key, value := range b.subtractLibGroups {
		tmp := b.groups[key]
		tmp.liberties--
		b.groups[key] = tmp
		fmt.Printf("key: %d, value: %v\n", key, value)
		fmt.Println("liberties", b.groups[key].liberties)
	}

	_x, _y := b.translate(x, y)
	b.grid[_y][_x] = stone
	return nil
}

func (b *Board) Merge(group Group, mergeGroup Group) {
	if group.id != mergeGroup.id {
		b.groups[group.id] = Group{
			liberties: group.liberties + mergeGroup.liberties - 1,
			id:        group.id,
			stones:    append(group.stones, mergeGroup.stones...),
		}
	}
}

func (b *Board) CheckNeighbours(neighbour Stone, stone *Stone) {
	if neighbour.player == player.NONE {
		stone.liberties++
	}
	if neighbour.player == stone.player {
		b.mergeGroups = append(b.mergeGroups, neighbour.groupID)
	}
	if neighbour.player == player.Opposite(stone.player) {
		b.subtractLibGroups[neighbour.groupID] = true
	}
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
