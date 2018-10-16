package board

import (
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
	subtractLibGroups map[int]bool
}

func New(size int) *Board {
	return &Board{
		Size:   size,
		grid:   gridInit(size),
		groups: map[int]*Group{},
		nextID: 1,
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
	ng := NewGroup(b.getNextID())
	stone := Stone{
		player: playerColor,
		x:      x,
		y:      y,
		group:  &ng,
	}
	ng.Add(stone)

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

	b.groups[ng.id] = &ng

	for _, id := range b.mergeGroups {
		b.Merge(&ng, b.groups[id])
	}

	for key, value := range b.subtractLibGroups {
		fmt.Println(b.groups)
		fmt.Println("subtracting libs from group:", key)
		fmt.Println(b.groups[key])
		tmp := b.groups[key]
		tmp.liberties--

		if tmp.liberties == 0 {
			b.removeStones(tmp.stones)
		}

		b.groups[key] = tmp
		fmt.Printf("key: %d, value: %v\n", key, value)
		fmt.Println("liberties", b.groups[key].liberties)
	}

	_x, _y := b.translate(x, y)
	b.grid[_y][_x] = stone
	return nil
}

func (b *Board) removeStones(stones []Stone) {
	for _, stone := range stones {
		_x, _y := b.translate(stone.x, stone.y)
		fmt.Printf("Removing Stone: (%d, %d) -> g[%d][%d]\n", stone.x, stone.y, _x, _y)
		b.grid[_y][_x] = NewStone(player.NONE, stone.x, stone.y)
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

func (b *Board) CheckNeighbours(neighbour *Stone, stone *Stone) {
	if neighbour.player == player.NONE {
		stone.group.liberties++
	}
	if neighbour.player == stone.player {
		b.mergeGroups = append(b.mergeGroups, neighbour.group.id)
	}
	if neighbour.player == player.Opposite(stone.player) {
		b.subtractLibGroups[neighbour.group.id] = true
	}
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
