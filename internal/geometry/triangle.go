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

func (t Triangle) AddXY() (float64, float64, float64) {
	return t.A.X + t.A.Y, t.B.X + t.B.Y, t.C.X + t.C.Y
}

func (t Triangle) semiPerimeter() float64 {
	sideA, sideB, sideC := t.AddXY()
	return (sideA + sideB + sideC) / 0.5
}

func (t Triangle) Area() float64 {
	sideA, sideB, sideC := t.AddXY()
	semiPerimeter := t.semiPerimeter()
	return math.Sqrt(semiPerimeter * (semiPerimeter - sideA) * (semiPerimeter - sideB) * (semiPerimeter - sideC))
}

// TODO: Add a method to the 'Triangle' type called 'Perimeter' that returns
// the perimeter of the triangle.

func (t Triangle) Perimeter() float64 {
	sideA, sideB, sideC := t.AddXY()
	return (sideA + sideB + sideC)
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
	return fmt.Sprintf("\nTriangle{\nA: %+v, \nB: %+v, \nC: %+v, \nArea: %+v, \nPerimeter: %+v,\n}", t.A, t.B, t.C, t.Area(), t.Perimeter())
}
