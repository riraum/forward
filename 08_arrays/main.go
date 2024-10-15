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

To get the length of an array, use `len`:

   var l int = len(myArray)

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

	// create an array with 10 integers, fill it
	// create emtpy array that will be filled with "x"
	// loop to the value of intArray and add x as many times as the value to lengthIntArray
	// print result
	var intArray []int
	intArray = append(intArray, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9)
	// debug
	fmt.Println("Print intArray:", intArray)

	// create new empty array of 10 strings
	var lengthIntArray []string
	// counter
	var arrayAssemblyCounter = 0

	// loop to create output array
	for {
		// if counter is > than the last/[10] array element break the loop
		if arrayAssemblyCounter > intArray[arrayAssemblyCounter]-1 {
			break
		}
		// append "x" to the array, by calling the function with a loop that returns the concat'd string
		lengthIntArray = append(lengthIntArray, concatVariable(intArray[arrayAssemblyCounter]))
		// debug
		fmt.Println("intArray element:", intArray[arrayAssemblyCounter])
		// increase counter to run with the next array element
		arrayAssemblyCounter++
	}
	// print result array, task req, DONE
	fmt.Println(lengthIntArray)

	// TODO: Create 2 arrays with 10 integers.
	// Create 3 more arrays:
	// - One with the sum of the elements of the 2 arrays.
	// - One with the difference of the elements of the 2 arrays.
	// - One with the multiplication of the elements of the 2 arrays.
	// E.g. [1, 2, 3] and [4, 5, 6] => [5, 7, 9], [-3, -3, -3], [4, 10, 18]
	//
	//	// Print the 3 new arrays.

	// create int array with 10 elements
	var anotherIntArray []int
	// fill it
	anotherIntArray = append(anotherIntArray, 1, 4, 5, 4, 6, 8, 5, 13, 5, 12)
	// create another
	var yetAnotherIntArray []int
	// fill it
	yetAnotherIntArray = append(yetAnotherIntArray, 23, 4, 1, 1, 3, 5, 6, 4, 6, 2)

	// create array for sums of elements of the 2 arrays
	var sumArray []int
	// create array for difference of elements of the 2 arrays
	var diffArray []int
	// // create array for multiplication of elements of the 2 arrays
	var multArray []int

	// create loop for element calculations
	// counter
	var calcCounter int = 0
	for {
		// TODO: replace hard coded value with last index of
		if calcCounter > 9 {
			break
		}
		// do math and append elements of arrays
		sumArray = append(sumArray, anotherIntArray[calcCounter]+yetAnotherIntArray[calcCounter])
		diffArray = append(diffArray, anotherIntArray[calcCounter]-yetAnotherIntArray[calcCounter])
		multArray = append(multArray, anotherIntArray[calcCounter]*yetAnotherIntArray[calcCounter])
		// increase counter
		calcCounter++
	}
	// print result sumArray
	fmt.Println(sumArray)
	// print result diffArray
	fmt.Println(diffArray)
	// print result multArray
	fmt.Println(multArray)

	// close main function, to be able to create function
}

// intArray function below

// concat "x" as many times as the respective element of intArray
func concatVariable(arrayElementValue int) string {
	// create empty variable for "x" string
	var arrayString string = ""
	// counter
	var concatCounter = 0

	for {
		// break if counter > respective element of intArray
		if concatCounter > arrayElementValue-1 {
			break
		}
		// concat "x"
		arrayString += "x"
		// increase counter
		concatCounter++

	}
	// debug
	fmt.Println(arrayString)
	// return concat'd string
	return arrayString
}
