package io

import "fmt"

func Read(prompt string) string {
	var input string
	fmt.Scan(&input)
	fmt.Printf("%s > ", prompt)
	return input
}
