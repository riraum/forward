package geometry

import (
	"fmt"
	"math"
	"testing"
)

// func TestPointAdd(t *testing.T) {
// 	tests := []struct {
// 		p, q Point
// 		want Point
// 	}{
// 		{Point{1, 1}, Point{2, 3}, Point{3, 4}},
// 		{Point{1, 1}, Point{0, 0}, Point{1, 1}},
// 		{Point{1, 1}, Point{-1, -1}, Point{0, 0}},
// 	}

//		for _, test := range tests {
//			got := test.p.Add(test.q)
//			if !got.Eq(test.want) {
//				t.Errorf("PointAdd(%v, %v) = %v, want %v", test.p, test.q, got, test.want)
//			}
//		}
//	}
func TestAddXY(t *testing.T) {
	tests := []struct {
		ta    Triangle
		want1 float64
		want2 float64
		want3 float64
		// want []float64
	}{
		{ta: Triangle{
			A: Point{X: 0, Y: 0},
			B: Point{X: 0, Y: 4},
			C: Point{X: 3, Y: 0},
		},
			want1: 0,
			want2: 4,
			want3: 3,
		},
		{ta: Triangle{
			A: Point{X: 0, Y: 1},
			B: Point{X: 0, Y: 4},
			C: Point{X: 4, Y: 0},
		},
			want1: 1,
			want2: 4,
			want3: 4},
	}

	// []float64{0, 4, 3},
	for _, test := range tests {
		got1, got2, got3 := test.ta.AddXY()
		if got1 != test.want1 && got2 != test.want2 && got3 != test.want3 {
			t.Errorf("AddXY(%v) = %v, %v, %v, want %v, %v and %v", test.ta, got1, got2, got3, test.want1, test.want2, test.want3)
		}
	}
}

// Test syntax from gotests
// func TestTriangle_AddXY(t *testing.T) {
// 	type fields struct {
// 		A Point
// 		B Point
// 		C Point
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   float64
// 		want1  float64
// 		want2  float64
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tr := Triangle{
// 				A: tt.fields.A,
// 				B: tt.fields.B,
// 				C: tt.fields.C,
// 			}
// 			got, got1, got2 := tr.AddXY()
// 			if got != tt.want {
// 				t.Errorf("Triangle.AddXY() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("Triangle.AddXY() got1 = %v, want %v", got1, tt.want1)
// 			}
// 			if got2 != tt.want2 {
// 				t.Errorf("Triangle.AddXY() got2 = %v, want %v", got2, tt.want2)
// 			}
// 		})
// 	}
// }

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
	test := fmt.Sprintf("\nTriangle{\nA: %+v, \nB: %+v, \nC: %+v, \nArea: %+v, \nPerimeter: %+v,\n}", ta.A, ta.B, ta.C, ta.Area(), ta.Perimeter())
	// wantedResult := ta
	wantedResult := fmt.Sprintf(`Triangle{
	A: {X:0 Y:0},
	B: {X:0 Y:4},
	C: {X:3 Y:0},
	Area: 146.83323874382123,
	Perimeter: 7,
	}`)
	if test != wantedResult {
		t.Errorf("Triangle.String() = %v; want %v", test, wantedResult)
	}
}