package main

import "fmt"

/*
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
	// var i int = 0
	// for {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// TODO: Write a loop that prints the numbers from 0 to 100, inclusive.
	// var j int = 0
	// for {
	// 	if j > 100 {
	// 		break
	// 	}
	// 	fmt.Println(j)
	// 	j += 1
	// }

	// TODO: Write a loop that prints the numbers from 100 to 0, inclusive.
	// var k int = 100
	// for {
	// 	if k < 0 {
	// 		break
	// 	}
	// 	fmt.Println(k)
	// 	k -= 1
	// }

	// TODO: Write a loop that print the numbers from 0 to 100, inclusive, but
	// only the even numbers.
	// var l int = 0
	// for {
	// 	if l > 100{
	// 		break
	// 	}
	// 	if l % 2 == 0 {
	// 		fmt.Println(l)
	// 		}
	// 	l += 1
	// }

	// TODO: Write a loop that prints the first 100 odd numbers.
	// var m int = 0
	// for {
	// 	if m > 100 {
	// 		break
	// 	}
	// 	if m % 2 != 0 {
	// 		fmt.Println(m)
	// 	}
	// 	m += 1
	// }

	// TODO: Define a variable and write a loop that calculates the sum of the
	// numbers from 0 to 100, inclusive. Print the result.
	// var n int = 0
	// var sum int 
	// for {
	// 	if n > 100 {
	// 		break
	// 	}
	// 	sum += n
	// 	n += 1
	// }
	// fmt.Println(sum)

	// TODO: Write a fizzbuzz.
	// Fizz buzz is a group word game for children to teach them about division.
	//  Players take turns to count incrementally, replacing 
	// any number divisible by three with the word "fizz", and any number divisible by five
	//  with the word "buzz", and any number divisible 
	// by both three and five with the word "fizzbuzz". 
	// If none applies, return value.


	// TODO: Write a loop that prints the first 100 Fibonacci numbers.
}
