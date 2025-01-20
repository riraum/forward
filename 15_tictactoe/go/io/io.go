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

// func Read(prompt string) int {
// 	var input int
// 	for i := 0; i <= 999; i++ {

// 		fmt.Printf("%s > ", prompt)
// 		_, err := fmt.Scan(&input)
// 		if valid(input) {
// 			fmt.Printf("Valid input\n")
// 			return input
// 		}

// 		if !valid(input) {
// 			fmt.Printf("Invalid input!\n")
// 			return input
// 		} else {
// 			fmt.Println("Test scan error", err)
// 			break
// 		}
// 	}
// 	return 999
// }

func Read(prompt string) int {
	var input string
	var checkedInput int
	for {
		i := 0
		fmt.Printf("%s > ", prompt)
		_, err := fmt.Scan(&input)
		fmt.Println("Print input", input)
		// fmt.Println("Print value", value)
		if valid(input) {
			fmt.Printf("Valid input\n")
			i++
			// input, err = strconv.Atoi(input)
			return checkedInput
		}

		if !valid(input) {
			fmt.Printf("Invalid input!\n")
			i++
			break
		} else {
			fmt.Println("Test scan error", err)
			break
		}

	}
	return 999
}

func valid(input string) bool {
	// var result bool
	// if err != nil {
	// 	return true
	// }
	// if value == 00 {
	// 	return true
	// }
	// if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 || input == 5 || input == 6 || input == 7 || input == 8 {
	// 	return true
	// }
	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("Error with string to int conversion!\n")
		return false
	}
	if intInput >= 0 && intInput <= 8 {
		fmt.Println("Print valid input", input)
		return true
	}

	// if input/input != input {
	// 	fmt.Println("Print invalid input", input)
	// 	return false
	// }

	fmt.Println("Print invalid input", intInput)
	return false

}

// func Read(prompt string) int {
// 	var input int
// 	for i := 0; i <= 3; i++ {

// 		fmt.Printf("%s > ", prompt)
// 		value, err := fmt.Scan(&input)
// 		if err != nil {
// 			continue
// 		} else {
// 			fmt.Println(err)
// 			if value >= 0 && value <= 8 {
// 				break
// 			}
// 		}
// 	}
// 	return input
// }
