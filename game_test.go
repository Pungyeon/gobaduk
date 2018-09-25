package game

import "testing"

func TestWhitePlaysFirst(t *testing.T) {
	game := New()
	if game.player != white {
		t.Errorf("game initialised with wrong player: %d", game.player)
	}
}

func TestBlackPlaysSecond(t *testing.T) {
	game := New()
	game.Move()

	if game.player != black {
		t.Errorf("player is not black after first move: %d", game.player)
	}
}

func TestWhitePlaysThird(t *testing.T) {
	game := New()
	game.Move()
	game.Move()

	if game.player != white {
		t.Errorf("player is not white after third move: %d", game.player)
	}
}
