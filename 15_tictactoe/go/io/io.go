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

// Logic to check if duplicate user input only
var takenCells []int

func Read(prompt string, freeCellSlice []int) int {
	var input string
	for {
		fmt.Printf("%s > ", prompt)
		fmt.Scan(&input)

		if valid(input, freeCellSlice) {
			validInput, _ := strconv.Atoi(input)
			// to remove
			takenCells = append(takenCells, validInput)
			return validInput
		}
	}
}

func valid(input string, freeCellSlice []int) bool {
	// debug
	fmt.Println("Debug freeSlice", freeCellSlice)
	intInput, err := strconv.Atoi(input)
	// var value int
	if err != nil {
		fmt.Print("Invalid input! Only single digit numbers allowed!\n")
		return false
	}

	for index, value := range freeCellSlice {
		// debug
		fmt.Println("Debug index", index)
		fmt.Println("Debug value", value)
		fmt.Println("Debug intInput", intInput)
		if intInput == value {
			// debug
			fmt.Printf("%v not taken\n", value)
			return true
		}
		// if intInput != value {
		// 	fmt.Printf("%v already taken!\n", value)
		// 	return false
		// }
	}
	// for value := range takenCells {
	// 	fmt.Println("Print value", value)
	// 	fmt.Println("Print intInput", intInput)
	// 	if value == intInput {
	// 		fmt.Println("Already taken")
	// 		return false
	// 	}
	// }

	// if intInput >= 0 && intInput <= 8 {
	// 	return true
	// }

	fmt.Println("Print invalid input, already taken", intInput)
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
