package grid

import (
	"fmt"
)

// Define column and row
type Grid struct {
	Cells [9]string
}

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

// Create variable for empty cell, filled with X and with O
// Return row of cells based on slice length, target index of slice and assign one of the 3 cell types
// Return column of cells based and do the same as with the row

// 	return fmt.Sprintf(
// 		"%s", cell, strings.Repeat(cell, count),
// 		// g.Cells...
// 	)
// }

// func (g Grid) String() string {
//   return fmt.Sprintf(
//     "<TODO>",
//     // g.Cells...
//   )
// }

// g = Grid struct
// .Cells = defines how many cells there are in this struct

// +---+---+---+
// | X |   |   |
// +---+---+---+
// |   | O | O |
// +---+---+---+
// | X |   |   |
// +---+---+---+

func (g Grid) IsWin() (bool, string) {
	// 1st row top to bottom horizontally
	if g.Cells[0] == "X" && g.Cells[1] == "X" && g.Cells[2] == "X" ||
		// 2nd row top to bottom horizontally
		g.Cells[3] == "X" && g.Cells[4] == "X" && g.Cells[5] == "X" ||
		// 3rd row top to bottom horizontally
		g.Cells[6] == "X" && g.Cells[7] == "X" && g.Cells[8] == "X" ||
		// 1st row top to bottom diagonally left to right
		g.Cells[0] == "X" && g.Cells[4] == "X" && g.Cells[8] == "X" ||
		// 2nd row top to bottom diagonally right to left
		g.Cells[2] == "X" && g.Cells[4] == "X" && g.Cells[6] == "X" {
		return true, "X"
	}

	// 1st row top to bottom vertically
	if g.Cells[0] == "X" && g.Cells[3] == "X" && g.Cells[6] == "X" ||
		// 2nd row top to bottom vertically
		g.Cells[1] == "X" && g.Cells[4] == "X" && g.Cells[7] == "X" ||
		// 3rd row top to bottom vertically
		g.Cells[2] == "X" && g.Cells[5] == "X" && g.Cells[8] == "X" {
		return true, "X"
	}

	// 1st row top to bottom horizontally
	if g.Cells[0] == "O" && g.Cells[1] == "O" && g.Cells[2] == "O" ||
		// 2nd row top to bottom horizontally
		g.Cells[3] == "O" && g.Cells[4] == "O" && g.Cells[5] == "O" ||
		// 3rd row top to bottom horizontally
		g.Cells[6] == "O" && g.Cells[7] == "O" && g.Cells[8] == "O" ||
		// 1st row top to bottom diagonally left to right
		g.Cells[0] == "O" && g.Cells[4] == "O" && g.Cells[8] == "O" ||
		// 2nd row top to bottom diagonally right to left
		g.Cells[2] == "O" && g.Cells[4] == "O" && g.Cells[6] == "O" {
		return true, "O"
	}

	// 1st row top to bottom vertically
	if g.Cells[0] == "O" && g.Cells[3] == "O" && g.Cells[6] == "O" ||
		// 2nd row top to bottom vertically
		g.Cells[1] == "O" && g.Cells[4] == "O" && g.Cells[7] == "O" ||
		// 3rd row top to bottom vertically
		g.Cells[2] == "O" && g.Cells[5] == "O" && g.Cells[8] == "O" {
		return true, "O"
	}
	return false, "Debug"
}

func (g Grid) FreeCells() []int {
	var resultIndexArray []int

	for index, value := range g.Cells {
		// for i := 0; i <= 8; i++ {
		if value == "" {
			resultIndexArray = append(resultIndexArray, index)
		}
	}
	return resultIndexArray
}
