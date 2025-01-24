package main

import (
	"fmt"
)

func main() {
	var input string

	for i := 0; ; i++ {
		fmt.Printf("Enter 5 digit word\n>")
		fmt.Scan(&input)

		if input == "Test" {
			fmt.Print("Correct!\n")
			break
		} else {
			fmt.Print("Not correct, try again!\n")
		}
	}
}
