package main

import "fmt"

/*

Go has variables with fixed types.
It means you cannot change the type of a variable after it has been declared.

Unlike JavaScript, Go is statically typed, which means that the type of a
variable is known at compile time. This is more restrictive but also ensures a
more robust codebase.

To declare a variable, you use the `var` keyword followed by the name of the
variable, then the type of the variable, and finally the value of the variable.

E.g.
    var age int = 42
    var greeting string = "hello"
    var awesome bool = true

Note: There are other ways to declare variables in Go, but we'll get to that later.
For now, each time you are asked to declare a variable, use this syntax.

You can also declare a variable without a value, in which case the variable will
be initialized with the default value of its type.

E.g.
    var age int // age is initialized to 0
    var greeting string // greeting is initialized to ""
    var awesome bool // awesome is initialized to false

There are other types, but we'll get to that later.

Notes:
- Names in Go are case-sensitive and follow the camelCase convention.
- Each defined variable must be used, otherwise the code will not run.
- You cannot declare a variable twice with the same name.
*/

func main() {
	var theAnswer int = 42
	fmt.Println(theAnswer)

	// TODO: See how the code fails to run because the variable `unused` is
	// declared but not used.
	// Fix the code by removing the unused variable.

	// TODO: See how the code fails fails to run because the variable
	// `mustBeUsed` is declared but not used.
	// Print the value of `mustBeUsed` to fix the error.
	var mustBeUsed int = 42
	fmt.Println(mustBeUsed)

	// TODO: See how the code fails to run because the variable `redeclared` is
	// declared twice.
	// Fix the error by removing the second declaration.
	var redeclared int = 42
	fmt.Println(redeclared)

	// TODO: Declare a variable `name` of type `string` and initialize it to
	// "Alice"
	var name string = "Alice"

	// TODO: Declare a variable `count` of type `int` and initialize it to 123
	var count int = 123

	// TODO: Declare a variable `isAwesome` of type `bool` and initialize it to
	// true
	var isAwesome bool = true

	// TODO: Declare a variable `otherName` of type `string` without value.
	var otherName string

	// TODO: Declare a variable `otherCount` of type `int` without value.
	var otherCount int

	// TODO: Declare a variable `otherIsAwesome` of type `bool` without value.
	var otherIsAwesome bool

	// TODO: Print all the variables. Note the default values of the variables that
	// were not initialized.
	// Hint, you can use multiple `fmt.Println` statements, or a single
	// statement. Try both ways.
	fmt.Println(name, count, isAwesome, otherName, otherCount, otherIsAwesome)
	fmt.Println(otherName)
	fmt.Println(otherCount)
	fmt.Println(otherIsAwesome)
}
