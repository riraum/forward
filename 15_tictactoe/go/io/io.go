package io

import "fmt"

/*
Input validation:
Create loop
Break loop if input is an integer, else keep looping
Break loop if input is on a free cell, else keep looping
*/

func Read(prompt string) int {
	var input int
	for i := 0; i <= 3; i++ {

		fmt.Printf("%s > ", prompt)
		value, err := fmt.Scan(&input)
		if valid(value, err) {
			continue
		} else {
			fmt.Println(err)
			break
		}
	}
	return input
}

func valid(value int, err error) bool {
	// var result bool
	if err != nil {
		return true
	} else {
		fmt.Println(err)
		if value >= 0 && value <= 8 {
			return false
		}
	}
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
