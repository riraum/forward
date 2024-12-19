package game

import (
	"fmt"

	"github.com/riraum/forward/15_tictactoe/go/grid"
	"github.com/riraum/forward/15_tictactoe/go/io"
)

func Play() {
	// TODO

	g := grid.Grid{
		Cells: [9]int{}}

	// gameGrid := Grid{
	// Cells: [9]int{}}
	X := 1
	O := 0
	g.Cells[0] = X
	g.Cells[1] = O
	g.Cells[4] = O
	g.Cells[7] = X
	g.Cells[8] = X

	fmt.Println(g.String())

	move := io.Read("What's your move?")
	fmt.Println("You chose", move)
}
