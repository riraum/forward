package io

import "fmt"

func Read(prompt string) string {
	var input string
	fmt.Printf("%s > ", prompt)
	fmt.Scan(&input)
	return input
}
