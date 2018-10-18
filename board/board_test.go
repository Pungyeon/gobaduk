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

	if b.Get(3, 2).group.liberties != 4 {
		t.Error("surrounding stone, was not granted back liberties")
	}
}

/*
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 2 0 0 0 0 0 0
0 2 ! 2 0 0 0 0 0 // ! should be illegal
0 0 2 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
*/

func TestIllegalMoveWithoutLiberties(t *testing.T) {
	b := New(9)
	b.Put(2, 3, 2)
	b.Put(2, 4, 3)
	b.Put(2, 3, 4)
	b.Put(2, 2, 3)

	if b.Put(1, 3, 3) == nil {
		t.Error("able to place move, without liberties, that doesn't capture any stones")
	}
}

/*
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 2 1 0 0 0 0 0
0 2 ! 2 1 0 0 0 0 // ! should be legal
0 0 2 1 0 0 0 0 0
0 0 0 0 0 0 0 0 0
*/

func TestKO(t *testing.T) {
	b := New(9)
	b.Put(2, 3, 2)
	b.Put(2, 4, 3)
	b.Put(2, 3, 4)
	b.Put(2, 2, 3)

	b.Put(1, 4, 2)
	b.Put(1, 5, 3)
	b.Put(1, 4, 4)

	if err := b.Put(1, 3, 3); err != nil {
		t.Error(err)
	}

	if b.Get(3, 3).group.liberties != 1 {
		t.Error("stone initialising ko, does not gain liberty:", b.Get(3, 3).group.liberties)
	}

	if b.Get(4, 3).player != player.NONE {
		t.Error("4,3 stone not captured:", b.Get(4, 3).player)
		b.Print()
	}

	if b.Put(2, 4, 3) == nil || b.Get(4, 3).player != player.BLACK {
		b.Print()
		t.Error(b.Get(3, 3))
		t.Error("allowed to place illegal move, in ko race")
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

/*
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
0 0 1 1 0 0 0 0 0
0 0 1 ! 0 0 0 0 0 // ! is an eventual 2 move
0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0
*/
func TestOverlapLiberty(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)
	b.Put(1, 3, 4)
	b.Put(1, 4, 4)
	expected := 8 // technically, this should be 7, but actually, not taking overlap into account, makes the program more simple

	if b.Get(3, 3).group.liberties != expected {
		t.Errorf("wrong liberty count for group: expected: %d, actual: %d", expected, b.Get(3, 3).group.liberties)
	}

	b.Put(2, 4, 3)
	expected = 6 // because of overlap, this 1 stone will take away 2 liberties
	if b.Get(3, 3).group.liberties != expected {
		t.Errorf("wrong liberty count for group: expected: %d, actual: %d", expected, b.Get(3, 3).group.liberties)
	}
}

/*
func TestPrint(t *testing.T) {
	b := New(9)
	b.Put(1, 3, 3)
	b.Put(1, 3, 4)
	b.Put(1, 4, 4)

	b.Print()

	t.Error()
}
*/
