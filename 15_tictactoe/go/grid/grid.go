package grid

import (
	"fmt"
	"strconv"
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
	var resultIndexArray []string
	for index, value := range g.Cells {
		// for i := 0; i <= 8; i++ {
		index := strconv.Itoa(index)
		if value == " " {
			resultIndexArray = append(resultIndexArray, index)
		} else {
			resultIndexArray = append(resultIndexArray, value)
		}
	}
	return fmt.Sprintf(`
+---+---+---+
| %+v | %+v | %+v |
+---+---+---+
| %+v | %+v | %+v |
+---+---+---+
| %+v | %+v | %+v |
+---+---+---+`, resultIndexArray[0], resultIndexArray[1], resultIndexArray[2], resultIndexArray[3], resultIndexArray[4], resultIndexArray[5], resultIndexArray[6], resultIndexArray[7], resultIndexArray[8],
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

	if g.Cells[x] != g.Cells[y] || g.Cells[x] != g.Cells[z] {
		return false, ""
	}

	if g.Cells[x] == "O" {
		return true, computerWin
	}

	return false, ""
}

func (g Grid) IsWin() (bool, string) {
	// check horizontal
	if win, winner := g.checkLine(0, 1, 2); win {
		return win, winner
	}
	if win, winner := g.checkLine(3, 4, 5); win {
		return win, winner
	}
	if win, winner := g.checkLine(6, 7, 8); win {
		return win, winner
	}
	//  check vertical
	if win, winner := g.checkLine(0, 3, 6); win {
		return win, winner
	}
	if win, winner := g.checkLine(1, 4, 7); win {
		return win, winner
	}
	if win, winner := g.checkLine(2, 5, 8); win {
		return win, winner
	}
	// check diagonally
	if win, winner := g.checkLine(0, 4, 8); win {
		return win, winner
	}
	if win, winner := g.checkLine(2, 4, 6); win {
		return win, winner
	}
	return false, "No winner"
}

func (g Grid) FreeCells() []int {
	var resultIndexArray []int

	for index, value := range g.Cells {
		if value == " " {
			resultIndexArray = append(resultIndexArray, index)
		}
	}
	return resultIndexArray
}
