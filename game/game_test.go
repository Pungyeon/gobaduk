package game

import (
	"testing"

	"github.com/Pungyeon/gobaduk/player"
)

func TestWhitePlaysFirst(t *testing.T) {
	game := New(9)
	if game.player != player.BLACK {
		t.Errorf("game initialised with wrong player: %d", game.player)
	}
}

func TestBlackPlaysSecond(t *testing.T) {
	game := New(9)
	game.Move(3, 3)

	if game.player != player.WHITE {
		t.Errorf("player is not white after first move: %d", game.player)
	}
}

func TestWhitePlaysThird(t *testing.T) {
	game := New(9).unsafeMove(3, 4).unsafeMove(3, 3)

	if game.player != player.BLACK {
		t.Errorf("player is not white after third move: %d", game.player)
	}
}

func TestBoardInitialisation(t *testing.T) {
	game := New(9)
	if game.board.Size != 9 {
		t.Error("board initialised to the wrong size")
	}
}

func TestSGFCreation(t *testing.T) {
	game := New(19)

	game.Move(4, 4)
	game.Move(4, 5)
	game.Move(16, 17)
	game.Move(4, 3)
	game.Move(17, 17)
	game.Move(5, 4)
	game.Move(18, 17)
	game.Move(3, 4)

	expected := `(;B[dp];W[do];B[pc];W[dq];B[qc];W[ep];B[rc];W[cp])`

	if game.SGF() != expected {
		t.Error("wrong output from SGF generation")
		t.Error(game.SGF())
		game.board.Print()
	}
}
