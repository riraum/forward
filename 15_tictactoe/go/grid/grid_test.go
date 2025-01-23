package grid

import (
	"slices"
	"testing"
)

func TestString(t *testing.T) {
	testGrid := Grid{
		Cells: [9]string{"X", "X", "X", " ", "O", "X", " ", "X", "X"},
	}
	test := testGrid.String()
	wantedResult := `
+---+---+---+
| X | X | X |
+---+---+---+
| 3 | O | X |
+---+---+---+
| 6 | X | X |
+---+---+---+`
	if test != wantedResult {

		t.Errorf("testGrid.String() = %v; want %v", test, wantedResult)
	}
}

func TestIsWin(t *testing.T) {
	computerWin := "Computer wins!"
	humanWin := "Human wins!"
	tests := []struct {
		g          Grid
		wantIsWin  bool
		wantWinner string
	}{
		{
			g: Grid{
				Cells: [9]string{"X", "X", "X", "", "O", "X", "", "X", "X"},
			},
			wantIsWin:  true,
			wantWinner: humanWin,
		},
		{
			g: Grid{
				Cells: [9]string{},
			},
			wantIsWin:  false,
			wantWinner: computerWin,
		},
		{
			g: Grid{
				Cells: [9]string{"O", " ", "X", "O", " ", "X", "O", "X", " "},
			},
			wantIsWin:  true,
			wantWinner: computerWin,
		},
		{
			g: Grid{
				Cells: [9]string{"X", " ", "O", " ", "X", " ", "O", " ", "X"},
			},
			wantIsWin:  true,
			wantWinner: humanWin,
		},
		{
			g: Grid{
				Cells: [9]string{"X", " ", " ", " ", " ", " ", " ", " ", " "},
			},
			wantIsWin:  false,
			wantWinner: "No winner",
		},
	}

	for _, test := range tests {
		got1, got2 := test.g.IsWin()
		if got1 != test.wantIsWin && got2 != test.wantWinner {
			t.Errorf("IsWin(%v) = %v, %v, want %v and %v", test.g, got1, got2, test.wantIsWin, test.wantWinner)
		}
	}
}

func TestFreeCells(t *testing.T) {
	testGrid := []struct {
		g    Grid
		want []int
	}{
		{
			g: Grid{
				Cells: [9]string{"X", "X", "X", " ", "O", "X", " ", "X", "X"},
			},
			want: []int{3, 6},
		},
		{
			g: Grid{
				Cells: [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, value := range testGrid {
		got := value.g.FreeCells()
		if !slices.Equal(got, value.want) {
			t.Errorf("testGrid = %vl want %v", got, value.want)
		}
	}
}
