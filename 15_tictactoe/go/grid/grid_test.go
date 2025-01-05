package grid

import (
	"slices"
	"testing"
)

func TestString(t *testing.T) {
	testGrid := Grid{
		Cells: [9]int{1, 1, 1, 0, 2, 1, 0, 1, 1},
	}
	test := testGrid.String()
	wantedResult := `
+---+---+---+
| 1 | 1 | 1 |
+---+---+---+
| 0 | 2 | 1 |
+---+---+---+
| 0 | 1 | 1 |
+---+---+---+`
	if test != wantedResult {

		t.Errorf("testGrid.String() = %v; want %v", test, wantedResult)
	}
}

func TestIsWin(t *testing.T) {
	tests := []struct {
		g     Grid
		want1 bool
		want2 string
	}{
		{
			g: Grid{
				Cells: [9]int{1, 1, 1, 0, 2, 1, 0, 1, 1},
			},
			want1: true,
			want2: "X",
		},
		{
			g: Grid{
				Cells: [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want1: false,
			want2: ""},
	}

	for _, test := range tests {
		got1, got2 := test.g.IsWin()
		if got1 != test.want1 && got2 != test.want2 {
			t.Errorf("IsWin(%v) = %v, %v, want %v and %v", test.g, got1, got2, test.want1, test.want2)
		}
	}
}

func TestFreeCells(t *testing.T) {
	// tests := []struct {
	// 	g    Grid
	// 	want []int
	// }{
	// 	{
	// 		g: Grid{
	// 			Cells: [9]int{1, 1, 1, 0, 2, 1, 0, 1, 1},
	// 		},
	// 		want: []int{3, 6},
	// 	},
	// 	{
	// 		g: Grid{
	// 			Cells: [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		},
	// 		want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
	// 	},
	// }

	// testGrid struct to test multiple cases
	testGrid := []struct {
		Cells []int
	}{
		{
			Cells: []int{1, 1, 1, 0, 2, 1, 0, 1, 1},
		},
		{
			Cells: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// Single testGrid
	// testGridU := Grid{
	// 	Cells: [9]int{
	// 		1, 1, 1, 0, 2, 1, 0, 1, 1},
	// }
	// apply function to testGrid
	// testGridF := testGridU.FreeCells()

	// test := testGrids.Cells.FreeCells()
	// var resultIndexArray []int
	// for index, value := range testGrids.Cells {
	// 	if value == 0 {
	// 		resultIndexArray = append(resultIndexArray, index)
	// 	}
	// }

	// wantedResult := []int{3, 6}

	// wantedResults has multiple results that can confirm correct functions in struct
	wantedResult := []struct {
		Cells []int
	}{
		{
			Cells: []int{3, 6},
		},
		{
			Cells: []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// wantedGridResult := Grid{}

	// wantedResult2 := []int{3, 6}

	for _, value := range testGrid {
		cells := value.Cells
	}

	// comparison of slices
	if !slices.Equal(testGrid, wantedResult) {
		t.Errorf("testGridF = %v; want %v", testGrid, wantedResult)
	}
}

// test := tests.IsWin()
// for _, test := range tests {
// 	got := test.g.FreeCells()
// 	if got != test.want {
// 	}
// }

// for _, test := range tests {
// 	got := test.g.FreeCells()
// 	if got != test.want {}
// }

// testGrid := Grid{
// 	Cells: [9]int{1, 1, 1, 0, 2, 1, 0, 1, 1},
// }
// testBool, testValue := testGrid.IsWin()
// wantedResult := true, "X"
// if testBool != wantedResult && testValue != wantedResult {
// }
