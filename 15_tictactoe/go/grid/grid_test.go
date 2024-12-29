package grid

import "testing"

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

// testGrid := Grid{
// 	Cells: [9]int{1, 1, 1, 0, 2, 1, 0, 1, 1},
// }
// testBool, testValue := testGrid.IsWin()
// wantedResult := true, "X"
// if testBool != wantedResult && testValue != wantedResult {

// }
