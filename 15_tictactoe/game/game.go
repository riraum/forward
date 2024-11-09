package game

import (
	"fmt"

	"github.com/riraum/forward/15_tictactoe/grid"
	"github.com/riraum/forward/15_tictactoe/io"
)

func Play() {
	// TODO
	g := grid.Grid{}
	fmt.Println(g.String())

	move := io.Read("What's your move?")
	fmt.Println("You chose", move)
}
