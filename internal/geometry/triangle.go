package geometry

import "math"

// import "encoding/json"

type Triangle struct {
	A Point
	B Point
	C Point
}

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

func (t Triangle) Area() float64 {
	semiPerimeter := t.semiPerimeter()
	return math.Sqrt(semiPerimeter * (semiPerimeter * t.addAXY()) * (semiPerimeter * t.addBXY()) * (semiPerimeter * t.addCXY()))
}

func (t Triangle) Perimeter() float64 {
	return (t.addAXY() + t.addBXY() + t.addCXY())
}

// func (t Triangle) String() {
// 	// Get coordinates
// 	A := t.addAXY()
// 	B := t.addBXY()
// 	C := t.addCXY()
// 	// Get area of triangle
// 	Area := t.Area()
// 	// Get Perimeter of triangle
// 	Perimeter := t.Perimeter()
// 	// Return data in req formatting
// 	return json.Marshal("Triangle{A: %s, B: %s, C: %s, Area: %s, Perimeter: %s,", A, B, C, Area, Perimeter,})
