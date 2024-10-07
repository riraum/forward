package main

import "fmt"

/*

Go has strings.

Strings are a list of individual characters.

They are defined using double quotes or backticks.

E.g.:
    var greetings string = "Hello, World!"
    var greetings string = `Hello, World!`

Using backticks allows you to use multi-line strings.

E.g.:
    var greetings string = `Hello,
    World!`


To use double quotes inside a string defined with double quotes, you can use
backslashes to escape them. Or use backticks.

E.g.:
    var greetings string = "Hello, \"World!\""
    var greetings string = `Hello, "World!"`

There are many functions that can be used with strings.

For now, let's focus on:
- `len()`: returns the length of the string, as an integer.
- `+`: concatenates two strings, returning a new string.

*/

func main() {
	// Notice how you can print an empty string:
	fmt.Println("")

	// TODO: Define a string variable called `greetings` with the value
	// "Hello, World!"
	var greetings string = "Hello, World!"

	// TODO: Print the variable `greetings`.
	fmt.Println(greetings)

	// TODO: Print the length of the string in the variable `greetings`.
	// Hint: it should be 13.
	fmt.Println(len(greetings))

	// TODO: Define a new variable called `name` containing your name.
	// Add your name to the `greetings` variable and store the result in a new
	// variable called `message`.
	// Print the `message` variable.
	var name string = "riraum"
	var message string = greetings + name
	fmt.Println(message)

	// TODO: Define a new variable called `quotedName` containing your name in
	// double quotes.
	// Add your name to the `greetings` variable and store the result in a new
	// variable called `quotedMessage`.
	// Print the `quotedMessage` variable.
	var quotedName string = `"riraum"`
	var quotedMessage string = greetings + quotedName
	fmt.Println(quotedMessage)

	// TODO: Define a new variable called `repeatedLetter` containing the letter
	// "a".
	// Add the letter "a" to the variable `repeatedLetter` 5 times.
	// Print the `repeatedLetter` variable after each addition.
	var repeatedLetter string = "a"
	repeatedLetter += repeatedLetter
	fmt.Println(repeatedLetter)
	repeatedLetter += repeatedLetter
	fmt.Println(repeatedLetter)
	repeatedLetter += repeatedLetter
	fmt.Println(repeatedLetter)
	repeatedLetter += repeatedLetter
	fmt.Println(repeatedLetter)
	repeatedLetter += repeatedLetter
	fmt.Println(repeatedLetter)
}
