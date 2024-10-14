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

	// create copy of array
	// loop 5 times (first 4 elements needed)
	// multiply elements * 2
	// add elements to new array
	// print new array
	// var timesTwo []int = tenInt
	var timesTwo [5]int
	var count2 = 0

	for {
		if count2 > 4 {
			break
		}
		timesTwo[count2] = (tenInt[count2] * 2)
		count2++
	}

	fmt.Println(timesTwo)

	// TODO: Create an array with 10 different strings in it.
	// Using this array, create a new array containing the length of each
	// string.
	// E.g. ["hello", "!"] => [5, 1]
	// Hint: Use a loop and the `len` function.
	// Print the new array.

	var stringArray [10]string
	stringArray[0] = "So..."
	stringArray[1] = "how"
	stringArray[2] = "are"
	stringArray[3] = "you"
	stringArray[4] = "holding"
	stringArray[5] = "up..."
	stringArray[6] = "BECAUSE"
	stringArray[7] = "IM"
	stringArray[8] = "A"
	stringArray[9] = "POTATO"
	fmt.Println(stringArray)
	fmt.Println("len test=", len(stringArray[0]))

	var lengthStringArray [10]int
	var count3 = 0

	for {
		if count3 > 9 {
			break
		}
		lengthStringArray[count3] = len(stringArray[count3])
		count3++
	}

	fmt.Println(lengthStringArray)

	// TODO: Create an array with 10 integers.
	// Create a new array containing strings with repeated "x" characters.
	// The number of "x" characters should be the value of the integer in the
	// original array.
	// E.g.: [1, 2, 3] => ["x", "xx", "xxx"]
	// Print the new array.
	var intArray [10]int
	intArray[0] = 2
	intArray[1] = 5
	intArray[2] = 3
	intArray[3] = 3
	intArray[4] = 6
	intArray[5] = 7
	intArray[6] = 2
	intArray[7] = 4
	intArray[8] = 6
	intArray[9] = 4
	fmt.Println("Print intArray:", intArray)

	var lengthIntArray [10]int
	var count4 = 0

	for {
		if count4 > 9 {
			break
		}
		lengthIntArray[count4] = len(intArray[count4])
		count4++
	}

	fmt.Println(lengthIntArray)

	// TODO: Create 2 arrays with 10 integers.
	// Create 3 more arrays:
	// - One with the sum of the elements of the 2 arrays.
	// - One with the difference of the elements of the 2 arrays.
	// - One with the multiplication of the elements of the 2 arrays.
	// E.g. [1, 2, 3] and [4, 5, 6] => [5, 7, 9], [-3, -3, -3], [4, 10, 18]
	// Print the 3 new arrays.
}
