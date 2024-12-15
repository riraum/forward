package game

import (
	"fmt"

	"github.com/riraum/forward/15_tictactoe/go/grid"
	"github.com/riraum/forward/15_tictactoe/go/io"
)

func Play() {
	// TODO
	g := grid.Grid{
		Cells: []int{3},
	}
	fmt.Println(g.String())

	move := io.Read("What's your move?")
	fmt.Println("You chose", move)
}
