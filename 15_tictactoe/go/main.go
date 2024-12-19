package main

import "github.com/riraum/forward/15_tictactoe/go/game"

/*

Write a tictactoe game in Go. The game should be played on the command line and
should play against the computer.

Guidance:

- Look for existing structure to get you started.

- Create a package called `grid` which contains _at least_:
    - The array of cells
    - `function (g Grid) String() string {}` that returns the 2d visualization of the grid
    - `function (g Grid) IsWin() (bool, string) {}` that returns if the grid is a win and who won
    - `function (g Grid) FreeCells() []int {}` that returns the indexes of free cells

- Create tests for each methods in the packages `grid`

- Create a package called `random` which contains:
    - `function Choose(values []int) int {}` that takes a random int from the list of ints
    - See https://pkg.go.dev/math/rand

- Create a package called `io` which contains:
    - `function Read(prompt string) string {}` that asks for an input
    - See https://pkg.go.dev/fmt#Scan

- Create a package called `game` that uses all the previous packages to make the game work.

- You don't need to test the `random`, `io`, and `game` packages.

*/

/*
Create array of cells
    - Create struct(ure)
    - Create a 9 cell structure
    - Change some cells and fill them with X, or 0

Print grid as strings
 - Create 3x3 grid, with line breaks
 - Input grid value of empty, X or O

*/

func main() {
	game.Play()
	// type Grid struct {
	// 	Cells [9]int
	// }
	// gameGrid := Grid{
	// 	Cells: [9]int{}}
	// X := 1
	// O := 0
	// gameGrid.Cells[0] = X
	// gameGrid.Cells[1] = O
	// gameGrid.Cells[4] = O
	// gameGrid.Cells[7] = X
	// gameGrid.Cells[8] = X
	// fmt.Println(gameGrid.String())
	// func (g Grid) String() string {
	//     return fmt.Sprintf()
	// }
}
