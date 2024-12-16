package grid

import (
	"fmt"
)

// Define column and row
type Grid struct {
	Cells []int
}

// testGrid := Grid {
// 	Cells: []int{3,3},
// }

func (g Grid) String() string {
	1 := "X"
	0 := "O"
	// TODO
	// |||____
	// |__|__|__|
	// |__|__|__|
	// |__|__|__|
	// row := g.Cells * "__"
	// column := g.Cells
	// Cells := []string{}
	// result = ""
	// for _, value := range row {
	// 	result += "__"
	// }
	// return ""
	// 	"|__|",
	// cell := "|__|"
	// count := g.Cells
	return fmt.Sprintf(`\n
+---+---+---+\n
| %+v | %+v  | %+v |\n
+---+---+---+\n
| %+v  | %+v | %+v |\n
+---+---+---+\n
| %+v | %+v  | %+v |\n
+---+---+---+\n`, g.Cells[0], g.Cells[1], g.Cells[2], g.Cells[3], g.Cells[4], g.Cells[5], g.Cells[6], g.Cells[7], g.Cells[8],
	)

	// return fmt.Sprintf(
	// 	"%v %v %v\n %v %v %v\n %v %v %v\n", g.Cells[0], g.Cells[1], g.Cells[2], g.Cells[3], g.Cells[4], g.Cells[5], g.Cells[6], g.Cells[7], g.Cells[8],
	// 	// g.Cells...
	// )
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
