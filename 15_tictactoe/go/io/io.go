package io

import "fmt"

func Read(prompt string) int {
	var input int
	fmt.Printf("%s > ", prompt)
	fmt.Scan(&input)
	return input
}
