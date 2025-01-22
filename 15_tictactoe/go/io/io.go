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
// var takenCells []int

func Read(prompt string, freeCellSlice []int) int {
	var input string
	for {
		fmt.Printf("%s > ", prompt)
		fmt.Scan(&input)

		if valid(input, freeCellSlice) {
			validInput, _ := strconv.Atoi(input)
			// to remove
			// takenCells = append(takenCells, validInput)
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
	}

	fmt.Println("Print invalid input, already taken", intInput)
	return false

}
