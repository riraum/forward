package game

import (
	"fmt"

	"github.com/riraum/forward/15_tictactoe/go/grid"
	"github.com/riraum/forward/15_tictactoe/go/io"
)

func Play() {
	// TODO
	// g := grid.Grid{}
	// fmt.Println(g.String())

	g := grid.Grid{
		Cells: [9]int{}}

	// gameGrid := Grid{
	// Cells: [9]int{}}
	// X := 1
	// O := 2
	// Hardcoded values for debugging
	g.Cells[0] = 1
	g.Cells[1] = 1
	g.Cells[2] = 1
	// g.Cells[3] = " "
	g.Cells[4] = 2
	g.Cells[5] = 1
	// g.Cells[6] = " "
	g.Cells[7] = 1
	g.Cells[8] = 1

	fmt.Println(g.String())

	fmt.Println(g.IsWin())

	fmt.Println(g.FreeCells())

	move := io.Read("What's your move?")
	fmt.Println("You chose", move)
}
