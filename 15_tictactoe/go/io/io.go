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

func Read(prompt string) int {
	var input string
	for {
		fmt.Printf("%s > ", prompt)
		fmt.Scan(&input)

		if valid(input) {
			validInput, _ := strconv.Atoi(input)
			return validInput
		}
	}
}

func valid(input string) bool {
	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("Invalid input! Only single digit numbers allowed!\n")
		return false
	}
	if intInput >= 0 && intInput <= 8 {
		return true
	}

	fmt.Println("Print invalid input", intInput)
	return false

}
