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
