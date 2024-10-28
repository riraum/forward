package main

import (
	"fmt"
	"math"
)

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

func (a area) areaSize() float64 {
	return a.side * a.side
}

// TODO: Adapt your code from 02_boolean to check the triangle type. The
// function should return a string with the type of the triangle.
func (t triangle) triangleType() string {
	if t.sumLengthSmallerLongestSide() {
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

func (t triangle) longestSide() float64 {
	if t.side1 < t.side2 && t.side2 < t.side3 {
		return t.side3
	}
	if t.side1 < t.side3 && t.side3 < t.side2 {
		return t.side2
	}
	return t.side1
}

func (t triangle) sumTwoShortestSides() float64 {
	if t.side1 < t.side2 && t.side2 < t.side3 {
		return t.side1 + t.side2
	}
	if t.side1 < t.side3 && t.side3 < t.side2 {
		return t.side1 + t.side3
	}
	return t.side2 + t.side3
}

func (t triangle) sumLengthSmallerLongestSide() bool {
	sum := t.sumTwoShortestSides()
	biggest := t.longestSide()
	return sum <= biggest
}

// TODO: Create a function `compare` on the structure `triangle` that takes another
// triangle as an argument and returns:
// - -1 if the area of the first triangle is smaller than the area of the second
// - 0 if the area of the first triangle is equal to the area of the second
// - 1 if the area of the first triangle is greater than the area of the second
// Hint: func (t1 triangle) compare(t2 triangle) int { ... }

func (t triangle) semiPerimeter() float64 {
	return (t.side1 + t.side2 + t.side3) / 0.5
}

// calculate square root
// func sqrt(x float64) float64 {
// }
// squareRoot :=  math.Sqrt()

func (t triangle) calcTriangleArea() float64 {
	return math.Sqrt(t.semiPerimeter() * (t.semiPerimeter() * t.side1) * (t.semiPerimeter() * t.side2) * (t.semiPerimeter() * t.side3))
}

func (t1 triangle) compare(t2 triangle) int {
	calcTriangleAreat1 := t1.calcTriangleArea()
	calcTriangleAreat2 := t2.calcTriangleArea()
	if calcTriangleAreat1 < calcTriangleAreat2 {
		return -1
	}
	if calcTriangleAreat1 == calcTriangleAreat2 {
		return 0
	}
	if calcTriangleAreat1 > calcTriangleAreat2 {
		return 1
	}
	// debug, to remove
	return 6060
}

// TODO: Create a struct named `person` with fields `name`, `age`, `height` and
// `weight` with the types you think are appropriate.
type person struct {
	name   string
	age    int
	height float64
	weight float64
}

// TODO: Create a function 'greet' on the structure `person` that prints "Hello,
// <name>!" to the console.
func (p person) greet() string {
	formatString := fmt.Sprintf("Hello, %s!", p.name)
	return formatString
}

// TODO: Create a function 'isAdult' on the structure `person` that returns
// a boolean indicating if the person is over 18 years old.
func (p person) isAdult() bool {
	return p.age >= 18
}

// TODO: Create a function 'bmi' on the structure `person` that returns the
// body mass index of the person. The formula is weight / (height * height).
func (p person) bmi() float64 {
	return p.weight / (p.height * p.height) * 10000
}

// TODO: Create a function that takes a list of person and prints "Hi, my
// name is <name>, my BMI is <bmi> and I am <an adult|an infant>".
func (p person) advIntro() string {
	bmi := p.bmi()
	// adultOrChild := p.isAdult()
	// if adultOrChild {
	// 	return "adult"
	// }
	adultOrChild := p.adultOrChild()
	formatString := fmt.Sprintf("Hi, my name is s%, my BMI is %t and I am an %u", p.name, bmi, adultOrChild)
	return formatString
}

// adult or child to string helper function
func (p person) adultOrChild() string {
	if p.isAdult() {
		return "adult"
	} else {
		return "infant"
	}
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
	fmt.Println(a.areaSize())

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
	biggerTriangle := triangle{
		side1: 20,
		side2: 20,
		side3: 30,
	}
	smallerTriangle := triangle{
		side1: 5,
		side2: 10,
		side3: 5,
	}

	fmt.Println(smallerTriangle.compare(bigTriangle))
	fmt.Println(smallerTriangle.compare(smallerTriangle))
	fmt.Println(biggerTriangle.compare(smallerTriangle))

	// TODO: Create a struct named `person` with fields `name`, `age`, `height` and
	// `weight` with the types you think are appropriate.
	// Placed before main function

	// TODO: Create a function 'greet' on the structure `person` that prints "Hello,
	// <name>!" to the console.
	greetName := person{
		name: "Gopher",
	}
	fmt.Println(greetName.greet())

	// TODO: Create a function 'isAdult' on the structure `person` that returns
	// a boolean indicating if the person is over 18 years old.
	anna := person{
		age: 17,
	}
	fmt.Println(anna.isAdult())

	hanna := person{
		age: 20,
	}
	fmt.Println(hanna.isAdult())

	// TODO: Create a function 'bmi' on the structure `person` that returns the
	// body mass index of the person. The formula is weight / (height * height).
	charlie := person{
		height: 180,
		weight: 85,
	}
	fmt.Println(charlie.bmi())

	dora := person{
		height: 150,
		weight: 39,
	}
	fmt.Println(dora.bmi())

	// // TODO: Create an array of 5 `person` instances with random data in it.
	randomArray := []person{
		{
			name:   "John",
			age:    60,
			height: 177,
			weight: 87,
		},
		{
			name:   "Anna",
			age:    18,
			height: 190,
			weight: 78,
		},
		{
			name:   "Francoise",
			age:    20,
			height: 176,
			weight: 60,
		},
		{
			name:   "Francois",
			age:    31,
			height: 169,
			weight: 66,
		},
		{
			name:   "Amelie",
			age:    64,
			weight: 71,
			height: 184,
		},
	}
	// debug
	fmt.Println(randomArray[0].name)

	// TODO: Create a function that takes a list of person and prints "Hi, my
	// name is <name>, my BMI is <bmi> and I am <an adult|an infant>".
	// debug hard coded
	fmt.Println(randomArray[0].advIntro())

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
