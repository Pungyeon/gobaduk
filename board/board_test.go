package board

import "testing"

func TestPlaceStone(t *testing.T) {
	b := New(9)
	b.Put(1, 1, 1)

	if b.grid[8][0] != 1 {
		t.Error("stone was placed incorrectly")
	}
}

func TestGetStone(t *testing.T) {
	b := New(9)
	b.Put(1, 1, 1)

	if b.Get(1, 1) != 1 {
		t.Error("stone was retrieved incorrectly")
	}
}
