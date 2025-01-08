package io

import "fmt"

/*
Input validation:
Create loop
Break loop if input is an integer, else keep looping
*/
func Read(prompt string) int {
	var input int
	fmt.Printf("%s > ", prompt)
	fmt.Scan(&input)
	return input
}
