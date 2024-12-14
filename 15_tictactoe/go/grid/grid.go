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

	return fmt.Sprintf(
		"%v\n%v\n%v\n", g.Cells, g.Cells, g.Cells,
		// g.Cells...
	)
}

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
