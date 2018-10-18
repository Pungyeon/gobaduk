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
	subtractLibGroups []int
	activeKO          Stone
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
	check := b.Get(x, y)
	if check.player != player.NONE {
		return errors.New("stone already on specified coordinates. stone cannot be placed")
	}

	fmt.Printf("check: %v, ko: %v\n", check, b.activeKO)
	if check == &b.activeKO {
		return errors.New("cannot place stone, on active KO")
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
	b.subtractLibGroups = make([]int, 0)

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

	_x, _y := b.translate(x, y)
	b.grid[_y][_x] = stone

	for _, rmGroup := range stonesToRemove {
		b.removeStones(rmGroup.stones)
	}

	fmt.Printf("(%d, %d): libs: %d, stones: %d, removed_g: %d\n",
		x, y,
		b.grid[_y][_x].group.liberties,
		len(b.grid[_y][_x].group.stones),
		len(stonesToRemove))
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
		b.grid[_y][_x] = NewStone(player.NONE, stone.x, stone.y)
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

func (b *Board) CheckNeighbours(neighbour *Stone, stone *Stone) {
	if neighbour.player == player.NONE {
		stone.group.liberties++
	}
	if neighbour.player == stone.player {
		b.mergeGroups = append(b.mergeGroups, neighbour.group.id)
	}
	if neighbour.player == player.Opposite(stone.player) {
		b.subtractLibGroups = append(b.subtractLibGroups, neighbour.group.id)
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

func (b *Board) Print() {
	for y := 0; y < b.Size; y++ {
		for x := 0; x < b.Size; x++ {
			fmt.Printf("%d ", b.grid[y][x].player)
		}
		fmt.Println()
	}
}
