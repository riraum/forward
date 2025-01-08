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

	for i := 0; i <= 9; i++ {
		// Output empty grid to help understand the concept?
		// Ask player to select a grid index
		fmt.Println(g.String())
		PlayerMove := io.Read("What's your choice? (Enter index 0-8 if not taken before)")
		fmt.Println("You chose", PlayerMove)
		// Add player marker to correct index of grid
		g.Cells[PlayerMove] = "X"
		// Output grid to visualize choice
		// fmt.Println(g.String())
		// Check if there is a winner
		fmt.Println(g.IsWin())
		//  Break condition in case of winner
		playerWin, _ := g.IsWin()
		if playerWin {
			fmt.Println(g.String())
			break
		}
		// Get freecells
		// freeCellSlice := []int{}
		// freeCellSlice = append(freeCellSlice, g.FreeCells()...)
		freeCellSlice := g.FreeCells()
		fmt.Println("Free cells:", freeCellSlice)
		// Break condition in case of no free cell
		freeCells := g.FreeCells()
		if len(freeCells) == 0 {
			break
		}
		// Make random computer choice, based on freecells
		computerMove := random.Choose(freeCellSlice)
		fmt.Println("Computer chose", computerMove)
		// Add computer marker to correct index of grid
		g.Cells[computerMove] = "O"
		// Output grid to visualize result
		// fmt.Println(g.String())
		// Check if there is a winner, if yes, exit, if no, continue from the start
		fmt.Println(g.IsWin())
		// Break condition in case of winner
		computerWin, _ := g.IsWin()
		if computerWin {
			fmt.Println(g.String())
			break
		}
	}
}
