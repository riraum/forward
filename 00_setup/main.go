package main

import "fmt"

func main() {
	// This is a comment.

	/*
	   This is a multi-line
	   comment.
	*/

	fmt.Println("Hello, World!")

	// fmt.Println is the generic print function in Go.
	// It comes from the standard package "fmt".

	// It is similar to `console.log` in JavaScript and take any number of
	// arguments.

	// Whenever you are asked to "print" something, you should use fmt.Println.

	// E.g. (don't worry about the syntax yet):
	fmt.Println("Hello", "World!", 42, 3.14, true, false, []int{1, 2, 3})

	// TODO:
	// 1. Print "Hello, mom!"
	// 2. Print a couple of numbers
}
