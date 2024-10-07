// This is a comment.

/*
   This is a multi-line
   comment.
*/

package main // IGNORE ME
import "fmt" // IGNORE ME

// The function main is what is called at the start of any program.
// Don't worry about the syntax yet, we'll get to that later.
// Just know that this is the entry point of the program.
func main() {
	// TODO: This line is wrong, remove it.

	// TODO: This line is badly formatted, run `go fmt` to fix it.
	fmt.Println("Hello, World!")

	// fmt.Println is the generic print function in Go.
	// It comes from the standard package "fmt" that was imported earlier.
	// Don't worry about imports for now.

	// It is similar to `console.log` in JavaScript and take any number of
	// arguments, with any number of types.

	// Whenever you are asked to "print" something, you should use fmt.Println.

	// E.g. (don't worry about the syntax yet):
	fmt.Println("Hello", "World!", 42, 3.14, true, false, []int{1, 2, 3})

	// TODO: Print "Hello, mom!"
	fmt.Println("Hello, mom!")
	// TODO: Print a couple of numbers
	fmt.Println(333, 3.1773, 1295912)
}
