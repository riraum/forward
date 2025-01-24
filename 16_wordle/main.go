package main

import (
	"fmt"
	"slices"
)

func valid(input string, wordSlice []string) bool {
	if slices.Contains(wordSlice, input) {
		return true
	}
	return false
}

func main() {
	var input string
	var wl []string
	wl = word_list.wordSlice

	for i := 0; ; i++ {
		fmt.Printf("Enter 5 digit word\n>")
		fmt.Scan(&input)

		// for value, index := range wordSlice {
		// 	if input == value {
		if valid(input, word_list.wordSlice) {
			fmt.Print("Correct word!\n")
			break
		} else {
			fmt.Print("Not correct word, try again!\n")
		}
	}

	// func (slice []string) checkValid(input) bool {
	// 	for value := range slice {
	// 		if input ==  value {
	// 			return true
	// 		}
	// 	}
	// }
}
