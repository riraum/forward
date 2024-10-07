package main

import "fmt"

/*
TEST
Go has loops.

Loops are one of the main building blocks of any programming language.

They are defined by the 3 elements:
- Initial state
- Break condition
- Increment

There are a couple of different ways to write a loop in Go. But we'll focus for
now on the simples one: the for loop.

E.g.:

    var i int = 0 // initial state
    for { // code that loops
        if i >= 10 { // break condition
            break
        }

        fmt.Println(i)

        i++ // increment
    }

This loop will print the numbers from 0 to 10.

Note how the `break` statement is used to stop the loop when the condition is
matched.

It is crucial to have a good understanding of initial state, break condition and
increment to write a loop that does what you want it to do.

There are other ways to write loops in Go, but we will use this one for now, as
it's also the most explicit.

*/

func main() {
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i += 1
	}

	// TODO: Write a loop that prints the numbers from 0 to 100, inclusive.

	// TODO: Write a loop that prints the numbers from 100 to 0, inclusive.

	// TODO: Write a loop that print the numbers from 0 to 100, inclusive, but
	// only the even numbers.

	// TODO: Write a loop that prints the first 100 odd numbers.

	// TODO: Define a variable and write a loop that calculates the sum of the
	// numbers from 0 to 100, inclusive. Print the result.

	// TODO: Write a fizzbuzz.

	// TODO: Write a loop that prints the first 100 Fibonacci numbers.
}
