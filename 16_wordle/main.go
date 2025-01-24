package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Printf("Enter 5 digit word %s \n>", input)
	fmt.Scan(&input)

	if input == "Test" {
		fmt.Print("Correct!\n")
	} else {
		fmt.Print("Not correct!\n")
	}
}
