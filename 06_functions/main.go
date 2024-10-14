package main

import "fmt"

/*
Go has functions.


A function is a block of code that performs a specific task. It is a
self-contained block of code that can be reused whenever needed.

Defining functions that logically divide the code into smaller parts
makes it easier to read and maintain large codebases.

Unlike JS or Ruby, Go is a statically typed language. This means that
you have to define the type of the function parameters and the return
type of the function.

Syntax:
    func function_name(parameter_name type) return_type {
        // code...
        return value
    }

The `return` keyword is used to return a value from the function.

If the function expects a return value, then the return type should be
specified in the function signature and the keyword `return` should be
used to return the value.

The parameters are optional. If the function does not expect any
parameters, they can be ommited.

You can define multiple parameters by separating them with a comma.

E.g.:
    func add(a int, b int) int {
        let sum int = a + b
        return sum
    }
    Note: The parameters and return types are all defined and strictly respected.
    Go will fail to run if it's not the cases.

Calling a function is like many other languages. You just need to write
the function name and pass the required parameters.

E.g.:
    add(2, 3)


Important: You need to understand the differecence between printing a value
and returning a value.

- Printing a value
    This displays the value on the console.
    This does nothing else about the value.
    E.g. fmt.Println("Hello")

- Returning a value
    This makes a specific value available to the caller of the function.
    In the same way that you can pass a value to a function via a parameter,
    you can also get a value back from a function with a return.
    E.g. return "Hello"

When the exercise asks you to return a value, you should use the `return`.
When the exercise asks you to print a value, you should use `fmt.Println`.

It is more than appreciated to use the `fmt.Println` to debug your code. Don't
hesitate to use it.

Note: functions can be defined in any order in the file.
*/

func main() {
	sayHello("John")
	fmt.Println(diff(1, 9))
	fmt.Println(diff(9, 1))
	fmt.Println(diff(4, -10))
	printInput("안녕", 3)
	printInput("go", 7)
	printInput("ça va", 4)
	fmt.Println(concatString("CatDog", 5))
	fmt.Println(concatString("Rice", 5))
	fmt.Println(concatString("김치", 6))
	// concatString("CatDog", 5)
	// concatString("Rice", 5)
	// concatString("김치", 6)
	printToInt(4)
	printToInt(9)
	printToInt(6)
	fmt.Println(returnStrings("Anna", "Hanna"))
	length("Gopher")
	length("Go rocks")
	lengthAdvanced("A")
	lengthAdvanced("")
	lengthAdvanced("Luna")
	fmt.Println(doingMath(3, 4))
	fmt.Println(doingMath(5, 10))
}

// For each exercise, call the functions a couple of times with different
// values and print the results.

// TODO: Write a function that takes two integers as parameters and returns
// the difference of the two integers.
func diff(num1 int, num2 int) int {
	return num2 - num1
}

// TODO: Write a function that takes a string and an integer as parameters.
// The function should print the string as many times as the integer.
func printInput(str string, times int) {
	var count = 1
	for {
		if count > times {
			break
		}
		fmt.Println(str)
		count++
	}
}

// TODO: Write a function that takes a string and an integer as parameters.
// The function should concatenate (i.e. join) the string with itself as
// many times as the integer and return the result.
// Call the function a couple of times with different values and print the
// results.
func concatString(str string, times int) string {
	// fmt.Println("pre loop", str)
	// fmt.Println("pre loop time", times)
	var result string = ""
	for {
		if times < 0 {
			break
		}

		result += str
		times = times - 1
		// fmt.Println("times", times)
	}
	// fmt.Println("str", str)
	return result
}

// TODO: Write a function that takes a single integer as a parameter and
// prints all numbers from 1 to that number.
func printToInt(parInt int) {
	var count = 1
	for {
		if count > parInt {
			break
		}
		fmt.Println(count)
		count++
	}
}

// TODO: Write a function that takes 2 strings as parameters and returns
// the sentence "Hello <name1>, and <name2>."
// Prints the result.
func returnStrings(str1 string, str2 string) string {
	var result = "Hello " + str1 + ", and " + str2 + "."
	return result
}

// TODO: Write a function that takes a string as a parameter and prints:
// "The string '<the string>' has <number of characters> characters."
func length(str string) {
	var length int = len(str)
	fmt.Println("The string", str, "has", length, "characters.")
}

// TODO: Update the previous function to handle the cases where the string
// is empty or has a single character.
func lengthAdvanced(str string) {
	var length = len(str)
	// TODO Put generic print text in variable
	if str == "" {
		fmt.Println("Empty string.")
	} else if length == 1 {
		fmt.Println("The string", str, "has", length, "character.")
	} else {
		fmt.Println("The string", str, "has", length, "characters.")
	}
}

// TODO: Write two functions mul and add that take two integers as
// parameters and return the result of the multiplication and addition
// respectively.
// Using those functions, write a function that takes two integers and
// returns the result of the following formula: a * b + a
// Print the result.
func mul(num1 int, num2 int) int {
	var prod int = num1 * num2
	return prod
}

func add(num1 int, num2 int) int {
	var sum int = num1 + num2
	return sum
}

func doingMath(num1 int, num2 int) int {
	var mulValue = mul(num1, num2)
	var addValue int = add(mulValue, num1)
	return addValue
}

func sayHello(name string) {
	fmt.Println("Hello " + name)
}
