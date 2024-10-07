package main

import "fmt"

/*

Go has boolean.

Boolean values are either `true` or `false`.
Incidentally, you cannot name a variable true or false in Go.

E.g.:
    var true bool // This will not work.
    var awesome bool = true
    var notAwesome bool = false


Alongside with the boolean values, Go also has the following operators:

- `&&`: logical and
- `||`: logical or
- `!`: logical not

Those operators work on boolean values exclusively.

E.g.:
    1 || 2        // This will not run, 1 and 2 are not boolean values
    !"true"       // This will not run, "true" is not a boolean value
    true && false // false
    true || false // true
    !true         // false

Go also has the following comparison operators:

- `==`: equal
- `!=`: not equal

E.g.:
    1 == 1 // true
    1 != 1 // false

Note: The comparison only works on similar type.

E.g.:
    1 == "1" // This will not run, 1 is an integer and "1" is a string


Note: more operators will be introduced in the next sections.

Lastly, Go has control flow with `if` statements.
This is very similar to other programming languages, except that Go does not
require parentheses around the condition.

E.g.:
    if CONDITION0 {
        // Execute this block if CONDITION0 is true
    } else if CONDITION1 {
        // Execute this block if CONDITION1 is true
    } else {
        // Execute this block if CONDITION0 and CONDITION1 are false
    }

    if 1 == 2 {
        fmt.Println("1 is equal to 2")
    } else {
        fmt.Println("1 is not equal to 2")
    }


*/

func main() {
	// Note how the two boolean values are printed.
	fmt.Println(true, false)

	// TODO: Define a variable named `awesome` and set it to `true`.
	var awesome bool = true

	// TODO: Change the value of `awesome` to `false`.
	// Print the value of `awesome`.
	awesome = false
	fmt.Println(awesome)

	// TODO: Define two variables named `a` and `b` and set them to `true` and
	// `false` respectively.
	// Print the result of `a && b` and `a || b`, and `!a`.
	// Change the value of `a` and `b`, and run the program again.
	// Observe the results.
	var a bool = true
	var b bool = false
	fmt.Println(a && b, a || b, !a)
	a = false
	b = true
	fmt.Println(a && b, a || b, !a)

	// TODO: Define a variable named `x` and set it to `1`.
	// Define a variable named `y` and set it to `2`.
	// Create an if statement that checks if `x` is equal to `y`.
	// If it is, print "x is equal to y".
	// If it is not, print "x is not equal to y".
	// Change the value of `y` to `1` and run the program again.
	// Observe the results.
	var x int = 1
	var y int = 2
	if x == y {
		fmt.Println("x is equal to y")
	} else {
		fmt.Println("x is not equal to y")
	}
	y = 1
	if x == y {
		fmt.Println("x is equal to y")
	} else {
		fmt.Println("x is not equal to y")
	}

	// TODO: Define a variable named `firstName` and set it to "John".
	// Define a variable named `lastName` and set it to "Doe".
	// Create an if statement that checks if `firstName` is equal to "John" and
	// `lastName` is equal to "Doe".
	// If it is, print "You are John Doe".
	// If it is not, but that `firstName` is equal to "John", print "You are
	// another John".
	// If it is not, but that `lastName` is equal to "Doe", print "You are
	// another Doe".
	// If it is not, print "You are someone else".
	// Change the value of `firstName` and `lastName` to different values and
	// run the program again to show all cases.
	var firstName string = "John"
	var lastName string = "Doe"
	if firstName == "John" && lastName == "Doe" {
		fmt.Println("You are John Doe")
	} else if firstName == "John" && lastName != "Doe" {
		fmt.Println("You are another John")
	} else if firstName != "John" && lastName == "Doe" {
		fmt.Println("You are another Doe")
	} else {
		fmt.Println("You are someone else")
	}

	// Hint: there are 2 ways to do the previous exercise:
	// 1. Use a single level of if-else statements
	// 2. Use nested if-else statements
	// TODO: Try both ways.
	if firstName == "John" {
		if lastName == "Doe" {
			fmt.Println("You are John Doe")
		} else {
			fmt.Println("You are another John")
		}
	}
	if firstName != "John" {
		if lastName == "Doe" {
			fmt.Println("You are another Doe")
		} else {
			fmt.Println("You are someone else")
	}
	}
}
