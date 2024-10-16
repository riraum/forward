package main

import "fmt"

/*

Go has maps.

Maps are a collection of key-value pairs.
In Python, they are called dictionaries.
In JavaScript, they are called objects.

They map from one type to another.

E.g.:
    - string to int: "age" -> 25
    - string to string: "name" -> "John"

To create a map, you need to specify both the key and value types.

E.g.:
    - map[string]int: key is a string, value is an int
    - map[string]string: key is a string, value is a string

    var m map[string]int = map[string]int{}

Like `int`, `string`, and `bool`, `map[string]int` is a type.

Like arrays, you can use the `len` function to get the number of key-value pairs
in a map.

Adding a key-value pair to a map is done by using the key in square brackets:

    m["age"] = 25

To delete a key-value pair from a map, use the `delete` function:

    delete(m, "age")

To check if a key exists in a map, you can use the following syntax:

To get the value of a key in a map, use the following syntax:

    var value int
    var exists bool
    value, exists = m["age"]

`exists` will be true if the key exists in the map, and false otherwise.

If `exists` is false, then `value` will be the zero value of the corresponding
type.

You can thus create an array of maps: with [] in front of the type.

    []map[string]int
    E.g. [ { "age": 25 }, { "age": 30 } ]

You can also create a map of maps:

    map[string]map[string]int
    E.g. { "person1": { "age": 25 }, "person2": { "age": 30 } }

*/

func main() {
	var human map[string]int = map[string]int{}
	human["age"] = 25
	human["height"] = 180
	fmt.Println(human)

	// TODO: Print the length of the map human
	fmt.Println(len(human))

	// TODO: Check if the key "age" exists in the map human
	// If so print "age exists in the map"
	// If not print "age does not exist in the map"
	// Do the same for the key "weight"
	var value int
	var exists bool
	value, exists = human["age"]
	if exists {
		fmt.Println("age exists in the map")
	}
	if value == 0 {
		fmt.Println("age does not exist in the map")
	}

	// TODO: Change the value of the key "age" to 30.
	// Print the map.
	human["age"] = 30
	fmt.Println(human)

	// TODO: Create a new map called "person" with the following key-value
	// pairs:
	// - "name" -> "John"
	// - "age" -> 25
	// - "height" -> 180
	// - "weight" -> 70
	// Print the map and its length.

	// TODO: Create a function that takes a map as an argument and creates a new
	// map with the following key-value pairs:
	// - "name" -> "<existing name> + number 2"
	// - "age" -> "<existing age> + number 2"
	// - "height" -> <existing height> * 2
	// - "weight" -> <existing weight> * 2
	// - "tall" -> true if height > 200, false otherwise
	// Print the new map.

	// TODO: Create a function that takes a map as an argument and prints the
	// BMI of the person. The formula for BMI is:
	// BMI = weight / (height * height)
	// If the map does not have the key "height", print "Height not found".
	// If the map does not have the key "weight", print "Weight not found".
	// Otherwise, print "<name>'s BMI is <bmi>".
	// Hint: Use this function in later exercises.

	// TODO: Create an array of maps called "people" with the following maps:
	// - { "name": "John", "age": 25, "height": 180, "weight": 70 }
	// - { "name": "Jane", "age": 30, "height": 160, "weight": 50 }
	// - { "name": "Jack", "age": 35, "height": 170, "weight": 80 }
	// - { "name": "Jill", "age": 40, "height": 150, "weight": 60 }
	// - { "name": "James", "age": 45, "height": 190, "weight": 90 }
	// Print the BMI of each person in the array.
	// Hint: Use a loop.

	// TODO: Create a function that takes an array of maps as an argument and
	// prints the average BMI of the people in the array.

	// TODO: Create a function that takes an array of maps as an argument and
	// returns the map with the highest BMI.

	// TODO: Create a function that takes an array of maps as an argument and
	// add as a new key-value pair to each map the key "bmi" with the value
	// being the BMI of the corresponding person.
	// Hint: Use a loop.
	// Hint: Create a new map for each person.
	// Hint: Create a new array of maps.
}
