package game

import "testing"

func TestWhitePlaysFirst(t *testing.T) {
	game := New(9)
	if game.player != white {
		t.Errorf("game initialised with wrong player: %d", game.player)
	}
}

func TestBlackPlaysSecond(t *testing.T) {
	game := New(9)
	game.Move()

	if game.player != black {
		t.Errorf("player is not black after first move: %d", game.player)
	}
}

func TestWhitePlaysThird(t *testing.T) {
	game := New(9).Move().Move()

	if game.player != white {
		t.Errorf("player is not white after third move: %d", game.player)
	}
}

func TestBoardInitialisation(t *testing.T) {
	game := New(9)
	if game.board.Size != 9 {
		t.Error("board initialised to the wrong size")
	}
}
