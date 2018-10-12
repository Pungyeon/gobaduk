package board

import (
	"errors"
	"fmt"
)

type Board struct {
	Size   int
	grid   [][]Stone
	groups map[int]Group
	nextID int
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

func (b *Board) Put(player int, x, y int) error {
	if b.Get(x, y).player != 0 {
		return errors.New("stone already on specified coordinates. stone cannot be placed")
	}
	stone := NewStone(player, b.getNextID())

	groupsToMerge := make([]int, 0)

	if y < b.Size {
		groupsToMerge = b.CheckNeighbours(b.Get(x, y+1), &stone, groupsToMerge)
	}
	if y > 1 {
		groupsToMerge = b.CheckNeighbours(b.Get(x, y-1), &stone, groupsToMerge)
	}
	if x < b.Size {
		groupsToMerge = b.CheckNeighbours(b.Get(x+1, y), &stone, groupsToMerge)
	}
	if x > 1 {
		groupsToMerge = b.CheckNeighbours(b.Get(x-1, y), &stone, groupsToMerge)
	}

	currentGroup := NewGroup(stone)

	for _, id := range groupsToMerge {
		b.Merge(currentGroup, b.groups[id])
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

func (b *Board) CheckNeighbours(neighbour Stone, stone *Stone, groups []int) []int {
	if neighbour.player == 0 {
		stone.liberties++
	}
	if neighbour.player == stone.player {
		groups = append(groups, neighbour.groupID)
	}
	return groups
}

func (b *Board) Get(x, y int) Stone {
	_x, _y := b.translate(x, y)
	return b.grid[_y][_x]
}

func (b *Board) translate(x, y int) (int, int) {
	_y := b.Size - y
	_x := x - 1
	fmt.Printf("x: %d, y: %d, _x: %d, _y: %d\n", x, y, _x, _y)
	return _x, _y
}
