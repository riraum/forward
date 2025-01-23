package grid

import (
	"fmt"
)

// Define column and row
type Grid struct {
	Cells [9]string
}

// create empty grid
func EmptyGrid() Grid {
	return Grid{
		Cells: [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "},
	}
}

// print grid
func (g Grid) String() string {
	return fmt.Sprintf(`
+---+---+---+
| %+v | %+v | %+v |
+---+---+---+
| %+v | %+v | %+v |
+---+---+---+
| %+v | %+v | %+v |
+---+---+---+`, g.Cells[0], g.Cells[1], g.Cells[2], g.Cells[3], g.Cells[4], g.Cells[5], g.Cells[6], g.Cells[7], g.Cells[8],
	)
}

func (g Grid) checkLine(x, y, z int) (bool, string) {
	computerWin := "Computer wins!"
	humanWin := "Human wins!"

	if g.Cells[x] != g.Cells[y] || g.Cells[x] != g.Cells[z] {
		return false, ""
	}

	if g.Cells[x] == "X" {
		return true, humanWin
	}

	return true, computerWin
}

// Create variable for empty cell, filled with X and with O
// Return row of cells based on slice length, target index of slice and assign one of the 3 cell types
// Return column of cells based and do the same as with the row
func (g Grid) IsWin() (bool, string) {
	// check horizontal
	if win, winner := g.checkLine(0, 1, 2); win {
		return win, winner
	}
	//  check vertical
	if win, winner := g.checkLine(0, 3, 6); win {
		return win, winner
	}
	// check diagonally
	if win, winner := g.checkLine(0, 4, 8); win {
		return win, winner
	}
	return false, "No winner"
}

func (g Grid) FreeCells() []int {
	var resultIndexArray []int

	for index, value := range g.Cells {
		// for i := 0; i <= 8; i++ {
		if value == " " {
			resultIndexArray = append(resultIndexArray, index)
		}
	}
	return resultIndexArray
}
