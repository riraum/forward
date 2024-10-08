package main

import "fmt"

/*

More boolean exercises.

Also, note that you can use parentheses to group expressions, just like in math.

E.g.:
    false && (false || true)  // false
    (false && false) || true  // true

*/

func main() {
	fmt.Println(false && (false || true))
	fmt.Println((false && false) || true)

	// TODO: Write a function that prints a truth table for all the following
	// Operations:
	//  - AND
	//  - OR
	//  - NAND (NOT AND)
	//  - NOR  (NOT OR)
	//  - XOR  (EXCLUSIVE OR)
	// E.g.: logicalTable(true, true) would print:
	// AND : true
	// OR  : true
	// NAND: false
	// ...
	// Call this function for all 4 combinations of true and false.

	// TODO: Write a function that takes a string and an integer and performs
	// the following checks:
	// - If the string is empty, return "empty"
	// - If the string is longer than the integer
	//   - If the integer is even, return "longer and even"
	//   - If the integer is odd, return "longer and odd"
	// - If the string is shorter than the integer, return "shorter"

	// TODO: Write a function that takes 5 integers (grades) and returns the
	// following:
	// - If any of the grades is less than 0 or greater than 100, return
	// "invalid grade".
	// - If the average of the grades is greater than or equal to 90
	//   - If no grade is below 80, return "A+"
	//   - If no grade is below 70, return "A"
	//   - If any grade is below 70, return "A-"
	// - If the average of the grades is greater than or equal to 80, return
	// "B"
	//   - If no grade is below 80, return "B+"
	//   - If no grade is below 70, return "B"
	//   - If any grade is below 60, return "B-"
	// - If the average of the grades is greater than or equal to 70, return
	// "C"
	//   - If no grade is below 70, return "C+"
	//   - If no grade is below 60, return "C"
	//   - If any grade is below 60, return "C-"

	// TODO: Write a function that takes 3 numbers, representing the sides of a
	// triangle, and returns the type of the triangle:
	// - If any of the sides is less than or equal to 0, return "invalid"
	// - If the sum of the lengths of the two shortest sides is less than or
	// equal to the length of the longest side, return "invalid"
	// - If all three sides are equal, return "equilateral"
	// - If two sides are equal, return "isosceles"
	// - If no sides are equal, return "scalene"
}
