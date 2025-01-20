package io

import (
	"fmt"
	"strconv"
)

/*
Input validation:
Create loop
Break loop if input is an integer, else keep looping
Break loop if input is on a free cell, else keep looping
*/

var takenCells []int

func Read(prompt string) int {
	var input string
	for {
		fmt.Printf("%s > ", prompt)
		fmt.Scan(&input)

		if valid(input) {
			validInput, _ := strconv.Atoi(input)
			takenCells = append(takenCells, validInput)
			return validInput
		}
	}
}

func valid(input string) bool {
	intInput, err := strconv.Atoi(input)
	// var value int
	if err != nil {
		fmt.Print("Invalid input! Only single digit numbers allowed!\n")
		return false
	}

	for value := range input {
		if value == intInput {
			fmt.Println("Already taken")
		}
	}

	if intInput >= 0 && intInput <= 8 {
		return true
	}

	fmt.Println("Print invalid input", intInput)
	return false

}

/*
Check user input and return invalid if cell index is already chosen

Existing code, `grid.FreeCells`
Used in game.go with `freeCellSlice` to print the free cells

Find a way to loop? the user user input through the `freeCellSlice` and return the correct boolean

Can't import and use freeCellSlice from `game.go` due to import cycle.

Can import other package `grid.go`
*/

// func freeCell(input int) bool {
// 		g := grid.Grid{
// 		Cells: [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "},
// 	}
// 	freeCellSlice := g.FreeCells()
// 	for index, value := range intInput {
// 		if intInput == freeCellSlice {
// 			fmt.Println("Freecell error")
// 			break
// 		}
// 	}

// }
