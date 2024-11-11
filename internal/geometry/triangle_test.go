package geometry

import "testing"

func TestAddXY(t *testing.T) {
	tests := []struct {
		ta   Triangle
		want []float64
	}{
		Triangle{
			Point{0, 0},
			Point{0, 4},
			Point{3, 0},
			{0, 4, 3, float64},
	},
	for _, test := range tests {
		got := test.ta.AddXY()
		if got != test.want {
			t.Errorf("AddxXY(%v) = %v, want%v", test.ta, got, test.want)
		}
	}
}
