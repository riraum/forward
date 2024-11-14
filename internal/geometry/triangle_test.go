package geometry

import (
	"fmt"
	"math"
	"testing"
)

func TestAddXY(t *testing.T) {
	tests := []struct {
		ta    Triangle
		want1 float64
		want2 float64
		want3 float64
	}{
		{
			ta: Triangle{
				A: Point{X: 0, Y: 0},
				B: Point{X: 0, Y: 4},
				C: Point{X: 3, Y: 0},
			},
			want1: 0,
			want2: 4,
			want3: 3,
		},
		{
			ta: Triangle{
				A: Point{X: 0, Y: 1},
				B: Point{X: 0, Y: 4},
				C: Point{X: 4, Y: 0},
			},
			want1: 1,
			want2: 4,
			want3: 4},
	}

	for _, test := range tests {
		got1, got2, got3 := test.ta.AddXY()

		if got1 != test.want1 || got2 != test.want2 || got3 != test.want3 {
			t.Errorf("AddXY(%v) = %v, %v, %v, want %v, %v and %v", test.ta, got1, got2, got3, test.want1, test.want2, test.want3)
		}
	}
}

func TestSemiPerimeter(t *testing.T) {
	ta := Triangle{
		A: Point{X: 0, Y: 0},
		B: Point{X: 0, Y: 4},
		C: Point{X: 3, Y: 0},
	}

	sideA, sideB, sideC := ta.AddXY()
	test := (sideA + sideB + sideC) / 0.5
	wantedResult := 14
	if test != float64(wantedResult) {

		t.Errorf("Triangle.SemiPerimeter() = %v; want %v", test, wantedResult)
	}
}

func TestArea(t *testing.T) {
	ta := Triangle{
		A: Point{X: 0, Y: 0},
		B: Point{X: 0, Y: 4},
		C: Point{X: 3, Y: 0},
	}

	sideA, sideB, sideC := ta.AddXY()
	semiPerimeter := ta.semiPerimeter()
	test := math.Sqrt(semiPerimeter * (semiPerimeter - sideA) * (semiPerimeter - sideB) * (semiPerimeter - sideC))
	wantedResult := 146.83323874382123
	if test != wantedResult {

		t.Errorf("Triangle.Area() = %v; want %v", test, wantedResult)
	}
}

func TestPerimeter(t *testing.T) {
	ta := Triangle{
		A: Point{X: 0, Y: 0},
		B: Point{X: 0, Y: 4},
		C: Point{X: 3, Y: 0},
	}
	sideA, sideB, sideC := ta.AddXY()
	test := sideA + sideB + sideC
	wantedResult := 7
	if test != float64(wantedResult) {
		t.Errorf("Triangle.Perimeter() =  %v; want %v", test, wantedResult)
	}
}

func TestString(t *testing.T) {
	ta := Triangle{
		A: Point{X: 0, Y: 0},
		B: Point{X: 0, Y: 4},
		C: Point{X: 3, Y: 0},
	}
	test := fmt.Sprintf("%v", ta)
	wantedResult := fmt.Sprintf("\nTriangle{\nA: %+v, \nB: %+v, \nC: %+v, \nArea: %+v, \nPerimeter: %+v,\n}", ta.A, ta.B, ta.C, ta.Area(), ta.Perimeter())
	if test != wantedResult {
		t.Errorf("Triangle.String() = %v; want %v", test, wantedResult)
	}
}
