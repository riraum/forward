package geometry

import "testing"

func TestFloat64Eq(t *testing.T) {
	tests := []struct {
		a, b float64
		want bool
	}{
		{1.0, 1.0, true},
		{1.0, 1.0 + epsilon/2, true},
		{1.0, 1.0 + epsilon*2, false},
	}

	for _, test := range tests {
		got := float64Eq(test.a, test.b)
		if got != test.want {
			t.Errorf("float64Eq(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}
}
