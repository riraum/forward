package geometry

import "testing"

// func TestPointAdd(t *testing.T) {
// 	tests := []struct {
// 		p, q Point
// 		want Point
// 	}{
// 		{Point{1, 1}, Point{2, 3}, Point{3, 4}},
// 		{Point{1, 1}, Point{0, 0}, Point{1, 1}},
// 		{Point{1, 1}, Point{-1, -1}, Point{0, 0}},
// 	}

// 	for _, test := range tests {
// 		got := test.p.Add(test.q)
// 		if !got.Eq(test.want) {
// 			t.Errorf("PointAdd(%v, %v) = %v, want %v", test.p, test.q, got, test.want)
// 		}
// 	}
// }
func TestAddXY(t *testing.T) {
	tests := []struct {
		ta Triangle
		want1 float64
		want2 float64
		want3 float64
		// want []float64
	}{
		Triangle{
			Point{0, 0},
			Point{0, 4},
			Point{3, 0},
		}
	}
		// []float64{0, 4, 3},
	for _, test := range tests {
		got := test.ta.AddXY()
		if got != test.want {
			t.Errorf("AddXY(%v) = %v, want %v, %v and %v", test.ta, got, test.want1, test.want2, test.want3)
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
