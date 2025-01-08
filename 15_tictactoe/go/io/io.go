package io

import "fmt"

/*
Input validation:
Create loop
Break loop if input is an integer, else keep looping
*/
func Read(prompt string) int {
	var input int
	for i := 0; i <= 3; i++ {

		fmt.Printf("%s > ", prompt)
		value, err := fmt.Scan(&input)
		if err != nil {
			continue
		}
		if value == 0 || value == 1 || value == 2 || value == 3 || value == 4 || value == 5 || value == 6 || value == 7 || value == 8 {
			break
		} else {
			break
		}
	}
	return input
}
