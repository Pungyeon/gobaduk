package board

import (
	"testing"

	"github.com/Pungyeon/gobaduk/player"
)

func TestPlaceStone(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)

	if b.grid[6][2].player != 1 {
		t.Error("stone was placed incorrectly")
	}
}

func TestGetStone(t *testing.T) {
	b := New(9)
	b.grid[6][2].player = 1

	if b.Get(3, 3).player != 1 {
		t.Error("stone was retrieved incorrectly")
	}
}

func TestCannotPlaceStoneIfSpaceIsOccupied(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)
	err := b.Put(1, 3, 3)

	if err == nil {
		t.Error("can place move on already placed stone.")
	}
}

func TestLibertyCount(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)
	expected := 4

	if b.Get(3, 3).group.liberties != expected {
		t.Errorf("wrong liberty count: expected: %d, actual: %d", expected, b.Get(3, 3).group.liberties)
	}
}

func TestLibertyGroupCount(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)
	b.Put(1, 3, 4)
	expected := 6

	if b.Get(3, 3).group.liberties != expected {
		t.Errorf("wrong liberty count: expected: %d, actual: %d", expected, b.Get(3, 3).group.liberties)
	}
}

func TestNextID(t *testing.T) {
	b := New(9)
	d := b.getNextID()

	if d != 1 {
		t.Errorf("intial next id should equal zero, but actual: %d", d)
	}

	if b.nextID != 2 {
		t.Errorf("next id should equal 1, but actual: %d", b.nextID)
	}

}

/*
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 2 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0
0 0 1 0 0 0 0 0 0 // 1 should be captured
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
*/

func TestGroupLiberties(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)
	b.Put(1, 3, 4)
	b.Put(2, 3, 5)
	expected := 5

	if b.Get(3, 3).group.liberties != expected {
		t.Errorf("wrong liberty count for group: expected: %d, actual: %d", expected, b.Get(3, 3).group.liberties)
	}
}

func TestPlaceStoneEdge(t *testing.T) {
	tt := []struct {
		name string
		x    int
		y    int
	}{
		{"bottom left", 1, 1},
	}

	for _, tc := range tt {
		b := New(9)
		if err := b.Put(1, tc.x, tc.y); err != nil {
			t.Errorf("%s: %v", tc.name, err)
		}
	}
}

/*
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 2 0 0 0 0 0 0
0 2 1 2 0 0 0 0 0 // 1 should be captured
0 0 2 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
*/

func TestCaptureStone(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)
	b.Put(2, 3, 2)
	b.Put(2, 4, 3)
	b.Put(2, 3, 4)
	b.Put(2, 2, 3)

	if b.Get(3, 3).player != 0 {
		t.Error("stone was not captured:", b.Get(3, 3))
	}
}

/*
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 2 2 0 0 0 0 0
0 2 1 1 2 0 0 0 0 // 1 should be captured
0 0 2 2 0 0 0 0 0
0 0 0 0 0 0 0 0 0
*/

func TestCaptureStones(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)
	b.Put(1, 4, 3)

	b.Put(2, 3, 2)
	b.Put(2, 4, 2)
	b.Put(2, 5, 3)
	b.Put(2, 4, 4)
	b.Put(2, 3, 4)
	b.Put(2, 2, 3)

	if b.Get(3, 3).player != player.NONE {
		t.Errorf("(%d != %d): stone was not captured: %v", b.Get(3, 3).player, player.NONE, b.Get(3, 3))
	}

	if b.Get(4, 3).player != player.NONE {
		t.Errorf("(%d != %d): stone was not captured: %v", b.Get(4, 3).player, player.NONE, b.Get(3, 4))
	}
}
