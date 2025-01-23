package game

import (
	"fmt"
	"time"

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
		fmt.Println("Available choices:", freeCellSlice)
		// User choice
		if i%2 == 0 {
			PlayerMove := io.Read("What's your choice?", freeCellSlice)
			fmt.Println("You chose", PlayerMove)
			// Add player marker to correct index of grid
			g.Cells[PlayerMove] = "X"
			// Slow down output
			time.Sleep(300 * (time.Millisecond))
		} else {
			// Make random computer choice, based on freecells
			computerMove := random.Choose(freeCellSlice)
			fmt.Println("Computer chose", computerMove)
			// Add computer marker to correct index of grid
			g.Cells[computerMove] = "O"
			// Slow down output
			time.Sleep(300 * (time.Millisecond))
		}
		//  Break condition in case of winner
		winTrue, winner := g.IsWin()
		if winTrue {
			fmt.Println(winner)
			fmt.Println(g.String())
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
