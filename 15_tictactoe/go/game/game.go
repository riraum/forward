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
	g := grid.EmptyGrid()
	for i := 0; ; i++ {
		// Output empty grid to help understand the concept
		// Ask player to select a grid index
		fmt.Println(g.String())
		freeCellSlice := g.FreeCells()
		fmt.Println("Free cells:", freeCellSlice)
		// User choice
		if i%2 == 0 {
			PlayerMove := io.Read("What's your choice? (Enter index 0-8 if not taken before)", freeCellSlice)
			fmt.Println("You chose", PlayerMove)
			// Add player marker to correct index of grid
			g.Cells[PlayerMove] = "X"
		} else {
			// Make random computer choice, based on freecells
			computerMove := random.Choose(freeCellSlice)
			fmt.Println("Computer chose", computerMove)
			// Add computer marker to correct index of grid
			g.Cells[computerMove] = "O"
		}
		//  Break condition in case of winner
		checkWin, _ := g.IsWin()
		if checkWin {
			fmt.Println(g.IsWin())
			break
		}
		// Output grid to visualize result
		// freeCellSlice = g.FreeCells()
		// fmt.Println("Free cells:", freeCellSlice)
		// fmt.Println(g.String())
		if len(freeCellSlice) == 0 {
			break
		}
	}
}
