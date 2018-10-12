package board

import "testing"

func TestPlaceStone(t *testing.T) {
	b := New(9)
	b.Put(1, 1, 1)

	if b.grid[8][0].player != 1 {
		t.Error("stone was placed incorrectly")
	}
}

func TestGetStone(t *testing.T) {
	b := New(9)
	b.Put(1, 1, 1)

	if b.Get(1, 1).player != 1 {
		t.Error("stone was retrieved incorrectly")
	}
}

func TestCannotPlaceStoneIfSpaceIsOccupied(t *testing.T) {
	b := New(9)
	b.Put(1, 1, 1)
	err := b.Put(1, 1, 1)

	if err == nil {
		t.Error("can place move on already placed stone.")
	}
}

func TestLibertyCount(t *testing.T) {
	b := New(9)
	b.Put(1, 1, 1)
	expected := 4

	if b.Get(1, 1).liberties != expected {
		t.Errorf("wrong liberty count: expected: %d, actual: %d", expected, b.Get(1, 1))
	}
}

func TestGroupLiberties(t *testing.T) {
	b := New(9)
	b.Put(1, 1, 1)
	b.Put(1, 1, 2)
	expected := 6

	if b.Get(1, 2).liberties != expected {
		t.Errorf("wrong liberty count for group: expected: %d, actual: %d", expected, b.Get(1, 2))
	}
}
