package main

import "fmt"

/*

Go has arrays.

In Go, and many other languages, arrays have a fixed type.
It means that you cannot combine different types in the same array, à la Ruby or
JS.

In order to define an array, you need to specify the type:

    var myArray []int

You can also specify the size of the array:

    var myArray [5]int

To add elements in the array, you can use the `append` function:

    myArray = append(myArray, 1)

It's important to note that `append` returns a new array, so you need to assign
it to the original array.

To get an element in the array, you can use the index with an integer:

    myArray[0]

To set an element in the array:

    myArray[0] = 2

*/

func main() {
	// TODO: Fix the errors
	var a []int
	a = append(a, 0)
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a)
	fmt.Println(a[0])

	// TODO: Create an array with 5 integers: 1, 2, 3, 4, 5
	// Print the array.
	var fiveInt []int
	fiveInt = append(fiveInt, 1, 2, 3, 4, 5)
	fmt.Println(fiveInt)

	// TODO: Use a loop to insert 10 integers in an array.
	// Print the array.
	var tenInt []int
	var count int = 1
	var times int = 10
	for {
		if count > times {
			break
		}
		tenInt = append(tenInt, count)
		count++
	}
	fmt.Println(tenInt)

	// TODO: Using the array from the previous step, print the first and the
	// last elements
	fmt.Println(tenInt[0])
	fmt.Println(tenInt[len(tenInt)-1])

	// TODO: Using the array from the previous step, create a new array
	// containing the first 5 elements multiplied by 2.
	// Hint: Use a loop.
	// Hint: The new array should contain [2, 4, 6, 8, 10].
	// Print the new array.

	// TODO: Create an array with 10 different strings in it.
	// Using this array, create a new array containing the length of each
	// string.
	// E.g. ["hello", "!"] => [5, 1]
	// Hint: Use a loop and the `len` function.
	// Print the new array.

	// TODO: Create an array with 10 integers.
	// Create a new array containing strings with repeated "x" characters.
	// The number of "x" characters should be the value of the integer in the
	// original array.
	// E.g.: [1, 2, 3] => ["x", "xx", "xxx"]
	// Print the new array.

	// TODO: Create 2 arrays with 10 integers.
	// Create 3 more arrays:
	// - One with the sum of the elements of the 2 arrays.
	// - One with the difference of the elements of the 2 arrays.
	// - One with the multiplication of the elements of the 2 arrays.
	// E.g. [1, 2, 3] and [4, 5, 6] => [5, 7, 9], [-3, -3, -3], [4, 10, 18]
	// Print the 3 new arrays.
}
