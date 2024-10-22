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
	// - "age" -> 25
	// - "height" -> 180
	// - "weight" -> 70
	// Print the map and its length.
	var person map[string]int = map[string]int{}
	person["age"] = 25
	person["height"] = 180
	person["weight"] = 70
	fmt.Println(person)
	fmt.Println(len(person))

	// TODO: Create a function that takes a map as an argument and creates a new
	// map with the following key-value pairs:
	// - "age" -> <existing age> + 2
	// - "height" -> <existing height> * 2
	// - "weight" -> <existing weight> * 2
	// - "tall" -> 1 if height > 200, 0 otherwise
	// Print the new map.

	createNewMap(person)

	// TODO: Create an array of maps called "people" with the following maps:
	// - { "height": 180, "weight": 70 }
	// - { "height": 160, "weight": 50 }
	// - { "height": 170, "weight": 80 }
	// - { "height": 150, "weight": 60 }
	// - { "height": 190, "weight": 90 }
	// Print the BMI of each person in the array.
	// Hint: Use a loop.

	// create empty array of maps
	var people []map[string]int = []map[string]int{}
	// fill array of maps
	people = append(people, map[string]int{"height": 180, "weight": 70})
	people = append(people, map[string]int{"height": 160, "weight": 50})
	people = append(people, map[string]int{"height": 170, "weight": 80})
	people = append(people, map[string]int{"height": 150, "weight": 60})
	people = append(people, map[string]int{"height": 190, "weight": 90})
	// print people[0]
	fmt.Println(people)

	var counter int = 0
	// // var calcPeople []map[string]int = people[counter]

	for {
		if counter >= len(people) {
			break
		}
		// print bmi by calling the calcBMI with an array value of people
		fmt.Println("BMI is", calcBMI(people[counter]))
		counter++
	}

	// For later exercises, use the same format as the previous one.

	// calculate average BMI
	fmt.Println("Avg BMI is", calcAvgBMI(people))

	// calculate highest BMI
	fmt.Println("Highest BMI is", highestBMI(people))

	// test helper int function
	fmt.Println(largestInt([]int{}))                      // 0
	fmt.Println(largestInt([]int{1, 2, 3}))               // 3
	fmt.Println(largestInt([]int{-1, -2, 0}))             // 0
	fmt.Println(largestInt([]int{100, 0, 3000, 0, -100})) // 3000

	// test helper string function
	fmt.Println(largestString([]string{"", ""}))           // ""
	fmt.Println(largestString([]string{"a", "bc", "def"})) // "def"
	// Close main function, to be able to declare another function
}

// create empty map
// create function
// create counter
// create loop
// create break condition if counter is > than the key-value pairs the map has
// operate on the input value of the keys and do math on it
// add the result to a new map
// add a new key value element and fill it depending on the value of the input height key
// print new map

var newMap map[string]int = map[string]int{}

func createNewMap(mapInput map[string]int) {
	var counter int = 0

	for {
		if counter > len(mapInput) {
			break
		}
		newMap["age"] = mapInput["age"] + 2
		newMap["height"] = mapInput["height"] * 2
		newMap["weight"] = mapInput["weight"] * 2
		if mapInput["height"] > 200 {
			newMap["tall"] = 1
		}
		newMap["tall"] = 0
		counter++
	}
	fmt.Println("createNewMap:", newMap)
}

// TODO: Create a function that takes a map as an argument and prints the
// BMI of the person. The formula for BMI is:
// BMI = weight / (height * height)
// If the map does not have the key "height", print "Height not found".
// If the map does not have the key "weight", print "Weight not found".
// Otherwise, print "<name>'s BMI is <bmi>".
// Hint: Use this function in later exercises.

func calcBMI(mapInput map[string]int) float64 {
	var value int
	var exists bool
	var bmi float64
	var heightConvert int = mapInput["height"]
	var heightConverted float64 = float64(heightConvert)
	var weightConvert int = mapInput["weight"]
	var weightConverted float64 = float64(weightConvert)

	// var result float64 = bmi

	value, exists = mapInput["height"]
	if !exists || value == 0 {
		fmt.Println("Height not found")
	}
	value, exists = mapInput["weight"]
	if !exists || value == 0 {
		fmt.Println("Weight not found")
	}
	// debug
	// fmt.Println("weight debug", mapInput["weight"])
	// fmt.Println("height debug", mapInput["height"])
	// debug
	// fmt.Println(heightAdjust)
	// fmt.Println(weightConverted)
	bmi = (weightConverted / (heightConverted * heightConverted))
	bmi = bmi * 10000
	// debug
	// fmt.Println(bmi)
	// TODO: check and fix syntax for name and bmi variable
	// fmt.Println("BMI is", bmi)
	return bmi

	// for {
	// 	var counter int = 0
	// 	if counter > len(mapInput) {
	// 		break
	// 	}
	// 	counter++
	// 	return bmi
	// }

}

// TODO: Create a function that takes an array of maps as an argument and
// prints the average BMI of the people in the array.

// create function
// accept array of maps as input
// loop through all arrays (e.g. "people" array of maps)
// calculate BMI (calcBMI function)
// add the sum of all array elements
// increase counter
// divide through the amount of array elements
// print result = avg BMI
func calcAvgBMI(arrayOfMapInput []map[string]int) float64 {
	var counter int
	var bmiSum float64
	var avgBMI float64
	var lenConvert int = len(arrayOfMapInput)
	var lenConverted float64 = float64(lenConvert)
	for {
		if counter >= len(arrayOfMapInput) {
			break
		}
		bmiSum += calcBMI(arrayOfMapInput[counter])
		// debug
		// fmt.Println(bmiSum)
		counter++
	}
	// debug
	// fmt.Println(lenConverted)
	avgBMI = bmiSum / lenConverted
	// debug
	// fmt.Println("avg debug print", avgBMI)
	return avgBMI
}

// TODO: Create a function that takes an array of maps as an argument and
// returns the map with the highest BMI.

// create function
// accept array of maps as input
// calculate the BMI with function calcBMI for all elements of the array
// compare all elements with each other to get highest value/BMI
// if one element is bigger than ALL of the others, return that element
// else continue comparing

// helper functions
// write a function that takes a list of int and return the biggest, or 0.
func largestInt(s []int) int {
	if len(s) == 0 {
		return 0
	}
	var largest = s[0]

	for {
		var counter int
		if counter > len(s) {
			break
		}
		if s[counter] > largest {
			largest = s[counter]
		}
		counter++
	}
	return largest
}

// helper functions
// write a function that takes a list of string and return the longest string, or an empty string.
func largestString(s []string) string {
	if len(s) == 0 {
		return ""
	}
	var largest = s[0]
	for {
		var counter int
		if counter > len(s) {
			break
		}
		if s[counter] > largest {
			largest = s[counter]
		}
		counter++
	}
	return largest
}

// main function
func highestBMI(arrayOfMapInput []map[string]int) map[string]int {
	// testing hard coded array element comparison with AND condition
	// if calcBMI(arrayOfMapInput[0]) > calcBMI(arrayOfMapInput[1]) && calcBMI(arrayOfMapInput[0]) > calcBMI(arrayOfMapInput[2]) && calcBMI(arrayOfMapInput[0]) > calcBMI(arrayOfMapInput[3]) && calcBMI(arrayOfMapInput[0]) > calcBMI(arrayOfMapInput[4]) {
	// 	return arrayOfMapInput[0]
	// }
	// var counter1 int = 0
	// var counter2 int = 0
	// for {
	// 	if counter1 >= len(arrayOfMapInput) {
	// 		break
	// 	}
	// 	if calcBMI(arrayOfMapInput[counter1]) > calcBMI(arrayOfMapInput[counter2]) {
	// 		counter1++
	// 	}
	// 	// if calcBMI(arrayOfMapInput[counter1])
	// }
	// test return
	// return arrayOfMapInput[0]
	// return arrayOfMapInput[counter1]

	// if len(arrayOfMapInput) == 0 {
	// 	TODO:
	// }
	var highestBMI = arrayOfMapInput[0]
	for {
		var counter int
		if counter > len(arrayOfMapInput) {
			break
		}
		if calcBMI((arrayOfMapInput[counter])) > calcBMI(highestBMI) {
			highestBMI = (arrayOfMapInput[counter])
		}
		counter++
	}
	return highestBMI
}

// if calcBMI()

// TODO: Create a function that takes an array of maps as an argument and
// add as a new key-value pair to each map the key "bmi" with the value
// being the BMI of the corresponding person.
// Hint: Use a loop.
// Hint: Create a new map for each person.
// Hint: Create a new array of maps.
// }
