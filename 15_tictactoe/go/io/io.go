package io

import (
	"fmt"
	"slices"
	"strconv"
)

func Read(prompt string, freeCellSlice []int) int {
	var input string
	for {
		fmt.Printf("%s > ", prompt)
		fmt.Scan(&input)

		if valid(input, freeCellSlice) {
			validInput, _ := strconv.Atoi(input)
			return validInput
		}
	}
}

func valid(input string, freeCellSlice []int) bool {
	intInput, err := strconv.Atoi(input)

	if err != nil {
		fmt.Print("Invalid input! Only single digit numbers allowed!\n")
		return false
	}

	if slices.Contains(freeCellSlice, intInput) {
		return true
	}

	fmt.Printf("'%d' is already taken! Try another choice\n", intInput)
	return false
}
