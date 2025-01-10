package io

import "fmt"

/*
Input validation:
Create loop
Break loop if input is an integer, else keep looping
Break loop if input is on a free cell, else keep looping
*/

/*
initialize function that can be used outside of this package
set the input value the variable name `prompt` of the type string
set the return value to type `int`
*/
func Read(prompt string) int {
	// initialize variable `input` of type `int`
	var input int
	// start `for` loop, counting from 0, until <=3
	for i := 0; i <= 3; i++ {
		// print the variable value of `prompt`
		fmt.Printf("%s > ", prompt)
		// set variable, discarding the value, only focusing on a potential error of the input given by the user, stored in the variable input
		_, err := fmt.Scan(&input)
		// if the result of the function call to `valid` returns false
		if !valid(input) {
			// print variable `err`, the error message of the that resulted from the user input
			fmt.Println(err)
			// break the loop
			break
			// if that is not the case, the function call to `valid` returns true, continue the loop
		} else {
			continue
		}
	}
	// in case the loop was not broken before, return the input
	return input
}

/*
Initialize the function `valid`, only usable in this package
Name the value that is passed the variable `input` of type `int`
Set the return value to type `bool`
*/
func valid(input int) bool {
	// var result bool
	// if err != nil {
	// 	return true
	// }
	// if value == 00 {
	// 	return true
	// }
	// if the `input` is >= 0 AND <=8 return false
	if input >= 0 && input <= 8 {
		return false
	}
	// if that's not the case, also return false
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
