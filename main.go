package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Pungyeon/gobaduk/game"
)

func main() {
	game := game.New(19)

	for i := 0; i < 100; i++ {
		err := errors.New("the game has just started")
		for err != nil {
			x := rand.Intn(18) + 1
			y := rand.Intn(18) + 1
			fmt.Printf("x: %d, y: %d\n", x, y)
			_, err = game.Move(x, y)
		}
	}

	fmt.Println(game.SGF())
}
