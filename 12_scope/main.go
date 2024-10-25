package main

import "fmt"

/*

Go has scope.

Scope is the region of code where a variable is accessible.

- A variable declared inside a function is only accessible inside that function.
- A variable declared inside an if statement is only accessible inside that if
statement.
- A variable declared inside a for loop is only accessible inside that for loop.
- A parameter is only accessible inside the function it is declared in.

Some variables and most functions are declared at the package level.
It means that they are available to all the code in the package.

For now, "a package" means the current file. But we'll get to packages later.

A variable declared at the package level is called a "global variable".
They have their usage, but it's best to avoid them when possible.

To properly scope a variable, you should always declare it as close as possible
to where it is used.

When a variable is declared with the same name as a global variable, it
"shadows" the global variable. It means that the global variable is not
accessible anymore.

Global variables in Go must be declared with the `var` keyword.
They can be instantiated with a value or not.

*/

var globalA = 42

func main() {
	fmt.Println(globalA)

	localA := 21
	fmt.Println(localA)

	printSomething()
	globalShadowing()
	paramShadowing(41)
	ifScope()
}

func printSomething() {
	// TODO: fix this code to print the value of the localA variable from the
	// main function.
	fmt.Println(localA)
}

func globalShadowing() {
	fmt.Println(globalA)

	// TODO: shadow the variable globalA by declaring a new variable with the
	// same name.
	fmt.Println(globalA)
}

func paramShadowing(paramA int) {
	fmt.Println(paramA)
	// TODO: fix this code to update the value of the paramA variable.
	paramA := 42
	fmt.Println(paramA)
}

func ifScope() {
	a := 3

	if a%2 == 0 {
		// nextInt is only accessible inside the if statement.
		nextInt := a + 1
		fmt.Println(nextInt, " is not odd")
	} else {
		fmt.Println(a, " is not odd")
	}

	// TODO: fix this code to print the value of the nextInt variable if it is
	// defined.
	// Hint: you need an extra variable.
	fmt.Println(nextInt)
}

func forScope() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// See how the code complains when you try to print the value of the i
	// variable here.
	// TODO: fix the code to print the last value of the i variable here.
	fmt.Println(i)
}
