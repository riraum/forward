package main

import "fmt"

/*
Go has structures.

Structures are the closest thing to classes in Go.
It is slightly different, but for now we can describe it as a class.

A struct is a custom type.
A struct is a collection of fields.
The fields can have different types and must have different names.

E.g.

	type person struct {
	    name string
	    age  int
	    height int
	    weight int
	}

	This defines a new type called `person`.

	Note how this is better than a map, because the fields are typed and can
	have different types.

	But, you also can't add/remove fields from a struct, those are fixed.

Creating a struct is like creating a variable.
You can access the fields using the dot operator.

E.g.

	bob := person{} // with a zero value for each field
	bob.name = "Bob"
	bob.age = 20
	fmt.Println(bob.name) // Bob

	alice := person{
	    name: "Alice",
	    age:  30,
	} // with a mix of zero and custom values
	fmt.Println(alice.age) // 30

Structs can have functions associated with them.
It means that when defined, you can call a function directly on the struct
instance.

E.g.

	type square struct {
	    side float64
	}

	func (s square) area() float64 {
	    return s.side * s.side
	}

Here, look closely at the `(s square)` part.

This means that the `area` function is associated with the `square` struct.
Inside the `area` function, `s` is the `square` struct instance, you can access
its fields.

Note: You must define your struct _outside_ of your function in most cases.
For now, always define your struct outside of any function, but create the
instances inside the function.

E.g.

	type square struct { ... }

	func (s square) area() float64 { ... }

	func main() {
	    s := square{side: 10}
	    fmt.Println(s.area()) // 100
	}
*/

// TODO: Create a struct named `triangle` with fields `side1`, `side2` and
// `side3` of type `float64`.
type triangle struct {
	side1 float64
	side2 float64
	side3 float64
}

// TODO: Create a function on the structure `area` that returns a `float64`.
type area struct {
	side float64
}

func (a area) area() float64 {
	return a.side * a.side
}

// TODO: Adapt your code from 02_boolean to check the triangle type. The
// function should return a string with the type of the triangle.
func (t triangle) triangleType() string {
	if sumLengthSmallerLongestSide(t.side1, t.side2, t.side3) {
		return "invalid triangle values"
	}
	// any side <=0 return error
	if t.side1 <= 0 || t.side2 <= 0 || t.side3 <= 0 {
		return "invalid triangle values"
	}
	if t.side1 == t.side2 && t.side1 == t.side3 {
		return "equilateral"
	}
	if t.side1 == t.side2 || t.side1 == t.side3 ||
		t.side2 == t.side3 {
		return "isosceles"
	}
	if t.side1 != t.side2 || t.side1 != t.side3 ||
		t.side2 != t.side3 {
		return "scalene"
	}

	return "Error 404: Triangle logic not yet found"
}

func longestSide(side1, side2, side3 float64) float64 {
	if side1 < side2 && side2 < side3 {
		return side3
	}
	if side1 < side3 && side3 < side2 {
		return side2
	}
	return side1
}

func sumTwoShortestSides(side1, side2, side3 float64) float64 {
	if side1 < side2 && side2 < side3 {
		return side1 + side2
	}
	if side1 < side3 && side3 < side2 {
		return side1 + side3
	}
	return side2 + side3
}

func sumLengthSmallerLongestSide(side1, side2, side3 float64) bool {
	sum := sumTwoShortestSides(side1, side2, side3)
	biggest := longestSide(side1, side2, side3)
	return sum <= biggest
}

func main() {
	// TODO: Create a struct named `triangle` with fields `side1`, `side2` and
	// `side3` of type `float64`.
	bigTriangle := triangle{
		side1: 313.32,
		side2: 231,
		side3: 422.66,
	}

	fmt.Println(bigTriangle)

	// TODO: Create a function on the structure `area` that returns a `float64`.
	a := area{side: 5}
	fmt.Println(a.area())

	// TODO: Adapt your code from 02_boolean to check the triangle type. The
	// function should return a string with the type of the triangle.
	invalidLengths := triangle{
		side1: 4,
		side2: 6,
		side3: 40,
	}
	fmt.Println(invalidLengths.triangleType())

	invalidValues := triangle{
		side1: 0,
		side2: -10,
		side3: 41,
	}
	fmt.Println(invalidValues.triangleType())

	// all sides equal
	equilateral := triangle{
		side1: 10,
		side2: 10,
		side3: 10,
	}
	fmt.Println(equilateral.triangleType())

	// two sides equal
	isosceles := triangle{
		side1: 10,
		side2: 10,
		side3: 33,
	}
	fmt.Println(isosceles.triangleType())

	// no sides equal
	scalene := triangle{
		side1: 11,
		side2: 12,
		side3: 13,
	}
	fmt.Println(scalene.triangleType())

	// TODO: Create a function `compare` on the structure `triangle` that takes another
	// triangle as an argument and returns:
	// - -1 if the area of the first triangle is smaller than the area of the second
	// - 0 if the area of the first triangle is equal to the area of the second
	// - 1 if the area of the first triangle is greater than the area of the second
	// Hint: func (t1 triangle) compare(t2 triangle) int { ... }

	// TODO: Create a struct named `person` with fields `name`, `age`, `height` and
	// `weight` with the types you think are appropriate.
	// TODO: Create a function 'greet' on the structure `person` that prints "Hello,
	// <name>!" to the console.
	// TODO: Create a function 'isAdult' on the structure `person` that returns
	// a boolean indicating if the person is over 18 years old.
	// TODO: Create a function 'bmi' on the structure `person` that returns the
	// body mass index of the person. The formula is weight / (height * height).
	// TODO: Create an array of 5 `person` instances with random data in it.
	// TODO: Create a function that takes a list of person and prints "Hi, my
	// name is <name>, my BMI is <bmi> and I am <an adult|an infant>".
	// TODO: Print the name of the oldest person in the array.
	// Hint: create a function that takes a list of person and return the
	// oldest.
	// TODO: Print the name of the person with the highest BMI in the array.
	// Hint: create a function that takes a list of person and return the person
	// with the highest BMI.

	// TODO: Create a struct named "album" with the following fields:
	// - name: string
	// - year: int
	// TODO: Create a struct named "artist" with the following fields:
	// - name: string
	// - members: a list of strings
	// - albums: a map[int]album where (int) is the year of the album
	// TODO: Create a couple of albums and artists.
	// TODO: Create a function on the artist struct that prints the list of all
	// the albums of the artist.
	// TODO: Create a function that takes a list of artists and prints the first
	// and last album ever released.
}
