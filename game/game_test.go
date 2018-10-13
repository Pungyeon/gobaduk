package game

import (
	"testing"

	"github.com/Pungyeon/gobaduk/player"
)

func TestWhitePlaysFirst(t *testing.T) {
	game := New(9)
	if game.player != player.WHITE {
		t.Errorf("game initialised with wrong player: %d", game.player)
	}
}

func TestBlackPlaysSecond(t *testing.T) {
	game := New(9)
	game.Move()

	if game.player != player.BLACK {
		t.Errorf("player is not black after first move: %d", game.player)
	}
}

func TestWhitePlaysThird(t *testing.T) {
	game := New(9).Move().Move()

	if game.player != player.WHITE {
		t.Errorf("player is not white after third move: %d", game.player)
	}
}

func TestBoardInitialisation(t *testing.T) {
	game := New(9)
	if game.board.Size != 9 {
		t.Error("board initialised to the wrong size")
	}
}
