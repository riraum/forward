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
	var checkedInput int
	for {
		// i := 0
		fmt.Printf("%s > ", prompt)
		_, err := fmt.Scan(&input)
		// fmt.Println("Print input", input)
		// fmt.Println("Print value", value)
		if valid(input) {
			// fmt.Printf("Valid input\n")
			// i++
			checkedInput, err = strconv.Atoi(input)
			if err != nil {
				fmt.Printf("Error with string to int conversion!")
			}
			return checkedInput
		}

		if !valid(input) {
			// fmt.Printf("Invalid input!\n")
			// i++
		} else {
			fmt.Println("Test scan error", err)
			break
		}

	}
	return 999
}

func valid(input string) bool {
	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("Error with string to int conversion=Invalid input!\n")
		return false
	}
	if intInput >= 0 && intInput <= 8 {
		fmt.Println("Print valid input", input)
		return true
	}

	fmt.Println("Print invalid input", intInput)
	return false

}
