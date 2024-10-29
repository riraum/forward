package geometry

import (
	"testing"
)

func TestSegmentLength(t *testing.T) {
	tests := []struct {
		s    Segment
		want float64
	}{
		{Segment{Point{0, 0}, Point{1, 1}}, 1.4142135623730951},
		{Segment{Point{0, 0}, Point{0, 0}}, 0},
		{Segment{Point{0, 0}, Point{1, 0}}, 1},
		{Segment{Point{0, 0}, Point{0, 1}}, 1},
	}

	for _, test := range tests {
		got := test.s.Length()
		if got != test.want {
			t.Errorf("SegmentLength(%v) = %v, want %v", test.s, got, test.want)
		}
	}
}

func TestSegmentEq(t *testing.T) {
	tests := []struct {
		s, t Segment
		want bool
	}{
		{Segment{Point{0, 0}, Point{1, 1}}, Segment{Point{0, 0}, Point{1, 1}}, true},
		{Segment{Point{0, 0}, Point{1, 1}}, Segment{Point{1, 1}, Point{0, 0}}, true},
		{Segment{Point{0, 0}, Point{1, 1}}, Segment{Point{0, 0}, Point{1, 0}}, false},
		{Segment{Point{0, 0}, Point{1, 1}}, Segment{Point{0, 0}, Point{0, 1}}, false},
	}

	for _, test := range tests {
		got := test.s.Eq(test.t)
		if got != test.want {
			t.Errorf("SegmentEq(%v, %v) = %v, want %v", test.s, test.t, got, test.want)
		}
	}
}

func TestSegmentContains(t *testing.T) {
	tests := []struct {
		s    Segment
		p    Point
		want bool
	}{
		{Segment{Point{0, 0}, Point{1, 1}}, Point{0.5, 0.5}, true},
		{Segment{Point{0, 0}, Point{1, 1}}, Point{1, 1}, true},
		{Segment{Point{0, 0}, Point{1, 1}}, Point{0, 0}, true},
		{Segment{Point{0, 0}, Point{1, 1}}, Point{0, 1}, false},
		{Segment{Point{0, 0}, Point{1, 1}}, Point{1, 0}, false},
		{Segment{Point{0, 0}, Point{1, 1}}, Point{2, 2}, false},
	}

	for _, test := range tests {
		got := test.s.Contains(test.p)
		if got != test.want {
			t.Errorf("SegmentContains(%v, %v) = %v, want %v", test.s, test.p, got, test.want)
		}
	}
}
