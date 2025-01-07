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
Output grid to visualize choice?
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

	for i := 0; i <= len(g.Cells); i++ {
		if i == len(g.Cells) {
			fmt.Println("Grid full, no win!")
			break
		}
		// Output empty grid to help understand the concept?
		fmt.Println(g.String())
		// Ask player to select a grid index
		PlayerMove := io.Read("What's your choice? (Enter index 0-8 if not taken before)")
		fmt.Println("You chose", PlayerMove)
		// Add player marker to correct index of grid
		g.Cells[PlayerMove] = "X"
		// Output grid to visualize choice
		fmt.Println(g.String())
		// Check if there is a winner
		fmt.Println(g.IsWin())
		if g.IsWin() == true, "X won" || g.IsWin() == true, "O won" {
			break
		}
		// Get freecells
		freeCellSlice := []int{}
		freeCellSlice = append(freeCellSlice, g.FreeCells()...)
		// Make random computer choice, based on freecells
		computerMove := random.Choose(freeCellSlice)
		// Add computer marker to correct index of grid
		g.Cells[computerMove] = "O"
		// Output grid to visualize result
		fmt.Println(g.String())
		// Check if there is a winner, if yes, exit, if no, continue from the start
		fmt.Println(g.IsWin())
	}
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
