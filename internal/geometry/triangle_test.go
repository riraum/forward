package geometry

import "testing"

func TestAddXY(t *testing.T) {
	tests := []struct {
		ty   Triangle
		want float64
	}{
	Triangle{
		A: Point{X: 0, Y: 0},
		B: Point{X: 0, Y: 4},
		C: Point{X: 3, Y: 0},
		0, 4, 3,}
	}
	}

