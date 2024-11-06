package geometry

import "math"

type Triangle struct {
	A Point
	B Point
	C Point
}

func (t Triangle) semiPerimeter() {
	return (t.A + t.B + t.B) / 0.5
}

func (t Triangle) Area() float64 {
	semiPerimeter := t.semiPerimeter()
	return math.Sqrt(semiPerimeter * (semiPerimeter * t.A) * (semiPerimeter * t.B) * (semiPerimeter * t.C))
}

func (t Triangle) Perimeter() {
	return (t.A + t.B + t.C)
}

func (t Triangle) String() {
	// Get coordinates
	A := t.A
	B := t.B
	C := t.C
	// Get area of triangle
	// Get area of triangle
	Area := t.Area()
	// Get Perimeter of triangle
	Perimeter := t.Perimeter()
	// Return data in req formatting
	return fmt.Sprintf("Triangle{\n A: %s,\n B: %s,\n C: %s,\n Area: %s,\n Perimeter: %s,\n,}", A, B, C, Area, Perimeter)
}
