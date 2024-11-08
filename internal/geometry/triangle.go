package geometry

import (
	"fmt"
	"math"
)

// TODO: In this file, create a new type called 'Triangle' with
// fields 'A', 'B', and 'C' of type 'Point'.

type Triangle struct {
	A Point
	B Point
	C Point
}

// TODO: Add a method to the 'Triangle' type called 'Area' that returns the
// area of the triangle.

// Create function that adds X, Y values of A, B, C of Triangle, in order to be able to use it for the coming functions and calculations of Triangle values
func (t Triangle) addAXY() float64 {
	return t.A.X + t.A.Y
}

func (t Triangle) addBXY() float64 {
	return t.B.X + t.B.Y
}

func (t Triangle) addCXY() float64 {
	return t.C.X + t.C.Y
}

func (t Triangle) semiPerimeter() float64 {
	// Draft
	// return (t.A + t.B + t.B) / 0.5
	//  fmt.Println( Add(t.A))
	fullA := t.addAXY()
	fullB := t.addBXY()
	fullC := t.addCXY()
	return (fullA + fullB + fullC) / 0.5
}

// TO CHECK
func (t Triangle) Area() float64 {
	fullA := t.addAXY()
	fullB := t.addBXY()
	fullC := t.addCXY()
	semiPerimeter := t.semiPerimeter()
	return math.Sqrt(semiPerimeter * (semiPerimeter * fullA) * (semiPerimeter * fullB) * (semiPerimeter * fullC))
}

// TODO: Add a method to the 'Triangle' type called 'Perimeter' that returns
// the perimeter of the triangle.

func (t Triangle) Perimeter() float64 {
	return (t.addAXY() + t.addBXY() + t.addCXY())
}

// TODO: Add a method to the 'Triangle' type called 'String' that returns
// a string representation of the triangle in the following format:
// "Triangle{
//   A: {X: 0, Y: 0},
//   B: {X: 0, Y: 4},
//   C: {X: 3, Y: 0},
//   Area: 6.0,
//   Perimeter: 12.0,
// }"

func (t Triangle) String() string {
	// Get coordinates
	// a := t.addAXY()
	// b := t.addBXY()
	// c := t.addCXY()
	// Get area of triangle
	// area := t.Area()
	// Get Perimeter of triangle
	// perimeter := t.Perimeter()
	// Return data in req formatting
	// return json.Marshal("Triangle{A: %s, B: %s, C: %s, Area: %s, Perimeter: %s,", t.A, t.B, t.C, t.Area(), t.Perimeter()})
	// out, err := json.Marshal(t)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(out))
	return fmt.Sprintf("Triangle{\nA: %+v, \nB: %+v, \nC: %+v, \nArea: %+v, \nPerimeter: %+v,", t.A, t.B, t.C, t.Area(), t.Perimeter())
}
