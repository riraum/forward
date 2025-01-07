package game

import (
	"fmt"

	"github.com/riraum/forward/15_tictactoe/go/grid"
	"github.com/riraum/forward/15_tictactoe/go/io"
	"github.com/riraum/forward/15_tictactoe/go/random"
)

/*
Initialize empty grid
Output empty grid to help understand the concept?
Allow player to choose X or O?
Ask player to select a grid index
Add player marker to correct index of grid
Check if there is a winner, if yes, exit
Output grid to visualize choice?
Get freecells
Make random computer choice, based on freecells
Add computer marker to correct index of grid
Output grid to visualize result
Output free cells
Ask player to select a grid index
Check if there is a winner, if yes, exit, if no, continue from the start
*/

func Play() {
	// Initialize empty grid
	g := grid.Grid{
		Cells: [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "},
	}
	// Output empty grid to help understand the concept?
	fmt.Println(g.String())
	// Ask player to select a grid index
	move := io.Read("What's your choice? (Enter index 0-8 if not taken before)")
	fmt.Println("You chose", move)
	// Add player marker to correct index of grid
	g.Cells[move] = "X"
	fmt.Println(g.String())
	// Check if there is a winner
	fmt.Println(g.IsWin())
	// Get freecells
	freeCellSlice := []int{}
	freeCellSlice = append(freeCellSlice, g.FreeCells()...)
	// Make random computer choice, based on freecells
	random.Choose(freeCellSlice)
	// Output grid to visualize result
	fmt.Println(g.String())

}

func Debug() {
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
	g.Cells[3] = " "
	g.Cells[4] = "O"
	g.Cells[5] = "X"
	g.Cells[6] = " "
	g.Cells[7] = "X"
	g.Cells[8] = "X"

	fmt.Println(g.String())

	fmt.Println(g.IsWin())

	fmt.Println(g.FreeCells())

	//debug int slice
	intSlice := []int{1, 5, 6, 8}

	fmt.Println(random.Choose(intSlice))

	// promptText := "Hurr durr"

	// fmt.Println(io.Read(promptText))

	move := io.Read("What's your move?")
	fmt.Println("You chose", move)
}
