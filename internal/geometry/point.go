package geometry

import "math"

type Point struct {
	X, Y float64
}

func (p Point) Eq(q Point) bool {
	return float64Eq(p.X, q.X) && float64Eq(p.Y, q.Y)
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Norm() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) Distance(q Point) float64 {
	return p.Sub(q).Norm()
}
