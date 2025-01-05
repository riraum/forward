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
		Cells: [9]string{}}

	// gameGrid := Grid{
	// Cells: [9]int{}}
	// X := "X"
	// O := "O"
	// Hardcoded values for debugging
	g.Cells[0] = "X"
	g.Cells[1] = "X"
	g.Cells[2] = "X"
	// g.Cells[3] = " "
	g.Cells[4] = "O"
	g.Cells[5] = "X"
	// g.Cells[6] = " "
	g.Cells[7] = "X"
	g.Cells[8] = "X"

	fmt.Println(g.String())

	fmt.Println(g.IsWin())

	fmt.Println(g.FreeCells())

	move := io.Read("What's your move?")
	fmt.Println("You chose", move)
}
