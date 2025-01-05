package grid

import (
	"fmt"
)

// Define column and row
type Grid struct {
	Cells [9]int
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
	if g.Cells[0] == 1 && g.Cells[1] == 1 && g.Cells[2] == 1 ||
		// 2nd row top to bottom horizontally
		g.Cells[3] == 1 && g.Cells[4] == 1 && g.Cells[5] == 1 ||
		// 3rd row top to bottom horizontally
		g.Cells[6] == 1 && g.Cells[7] == 1 && g.Cells[8] == 1 ||
		// 1st row top to bottom diagonally left to right
		g.Cells[0] == 1 && g.Cells[4] == 1 && g.Cells[8] == 1 ||
		// 2nd row top to bottom diagonally right to left
		g.Cells[2] == 1 && g.Cells[4] == 1 && g.Cells[6] == 1 {
		return true, "X"
	}

	// 1st row top to bottom vertically
	if g.Cells[0] == 1 && g.Cells[3] == 1 && g.Cells[6] == 1 ||
		// 2nd row top to bottom vertically
		g.Cells[1] == 1 && g.Cells[4] == 1 && g.Cells[7] == 1 ||
		// 3rd row top to bottom vertically
		g.Cells[2] == 1 && g.Cells[5] == 1 && g.Cells[8] == 1 {
		return true, "X"
	}

	// 1st row top to bottom horizontally
	if g.Cells[0] == 2 && g.Cells[1] == 2 && g.Cells[2] == 2 ||
		// 2nd row top to bottom horizontally
		g.Cells[3] == 2 && g.Cells[4] == 2 && g.Cells[5] == 2 ||
		// 3rd row top to bottom horizontally
		g.Cells[6] == 2 && g.Cells[7] == 2 && g.Cells[8] == 2 ||
		// 1st row top to bottom diagonally left to right
		g.Cells[0] == 2 && g.Cells[4] == 2 && g.Cells[8] == 2 ||
		// 2nd row top to bottom diagonally right to left
		g.Cells[2] == 2 && g.Cells[4] == 2 && g.Cells[6] == 2 {
		return true, "O"
	}

	// 1st row top to bottom vertically
	if g.Cells[0] == 2 && g.Cells[3] == 2 && g.Cells[6] == 2 ||
		// 2nd row top to bottom vertically
		g.Cells[1] == 2 && g.Cells[4] == 2 && g.Cells[7] == 2 ||
		// 3rd row top to bottom vertically
		g.Cells[2] == 2 && g.Cells[5] == 2 && g.Cells[8] == 2 {
		return true, "O"
	}
	return false, "Debug"
}

func (g Grid) FreeCells() []int {
	var resultIndexArray []int

	for index, value := range g.Cells {
		// for i := 0; i <= 8; i++ {
		if value == 0 {
			resultIndexArray = append(resultIndexArray, index)
		}
	}
	return resultIndexArray
}
