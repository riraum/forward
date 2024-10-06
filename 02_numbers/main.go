package main

import "fmt"

/*

Go has numbers.

You can define them with a couple of different types, the main ones being:
- `int` is a whole number, e.g. 42
- `float64` is a number with a decimal point, e.g. 3.14
    Don't worry about what `64` means for now, just remember that if you want
    number with a decimal point, you should use `float64`.

Numbers support a range of operations, the main ones being:
- `+` for addition
- `-` for subtraction
- `*` for multiplication
- `/` for division
- `%` for remainder (or modulus)

E.g.
    var x int = 10
    var y int
    x = x + 1 // x is now 11
    y = x - 1 // y is now 10
    y = y * 2 // y is now 20
    ...

Note: In this example, `x` and `y` must be of the same type.
    See a TODO later in this file.

*/

func main() {
	// TODO: See how the code doesn't run because the types are mismatched.
	// Fix the following code:
	var a int = 1
	var b string = "3" // wrong type!
	var c int = a + b
	fmt.Println(c) // want to print 4

	// TODO: See how the code doesn't run because the types are mismatched.
	// Fix the following code:
	var d int = 1
	var e int = 2.5
	fmt.Println(d + e) // we want to print 3.5

	// TODO: Define a variable `x` of type `int` and set it to 10

	// TODO: Define a variable `y` of type `int` and set it to 20

	// TODO: Define a variable `z` with the sum of `x` and `y`, what should be
	// its type?
	// Print the value of `z`.

	// TODO: Define 2 variables of type `int` and compute their sum, difference,
	// product, quotient, and remainder. Print all the results.
	// Hint: you can store the results in variables or print them directly.

	// TODO: Define a variable of type `int` and increment it by 1.
	// Then decrement it by 1.
	// Print the value of the variable after each operation.

	// TODO: Define a variable with a value of 1.
	// Increase its value by 0.5, 3 times.
	// Print the value of the variable after each operation.
}
