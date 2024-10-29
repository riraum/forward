package geometry

import (
	"math"
	"testing"
)

func TestPointEq(t *testing.T) {
	tests := []struct {
		p, q Point
		want bool
	}{
		{Point{1, 1}, Point{1, 1}, true},
		{Point{1, 1}, Point{1 + epsilon/2, 1 + epsilon/2}, true},
		{Point{1, 1}, Point{1 + epsilon*2, 1 + epsilon*2}, false},
	}

	for _, test := range tests {
		got := test.p.Eq(test.q)
		if got != test.want {
			t.Errorf("PointEq(%v, %v) = %v, want %v", test.p, test.q, got, test.want)
		}
	}
}

func TestPointAdd(t *testing.T) {
	tests := []struct {
		p, q Point
		want Point
	}{
		{Point{1, 1}, Point{2, 3}, Point{3, 4}},
		{Point{1, 1}, Point{0, 0}, Point{1, 1}},
		{Point{1, 1}, Point{-1, -1}, Point{0, 0}},
	}

	for _, test := range tests {
		got := test.p.Add(test.q)
		if !got.Eq(test.want) {
			t.Errorf("PointAdd(%v, %v) = %v, want %v", test.p, test.q, got, test.want)
		}
	}
}

func TestPointSub(t *testing.T) {
	tests := []struct {
		p, q Point
		want Point
	}{
		{Point{1, 1}, Point{2, 3}, Point{-1, -2}},
		{Point{1, 1}, Point{0, 0}, Point{1, 1}},
		{Point{1, 1}, Point{-1, -1}, Point{2, 2}},
	}

	for _, test := range tests {
		got := test.p.Sub(test.q)
		if !got.Eq(test.want) {
			t.Errorf("PointSub(%v, %v) = %v, want %v", test.p, test.q, got, test.want)
		}
	}
}

func TestPointNorm(t *testing.T) {
	tests := []struct {
		p    Point
		want float64
	}{
		{Point{3, 4}, 5},
		{Point{1, 1}, math.Sqrt(2)},
		{Point{0, 0}, 0},
	}

	for _, test := range tests {
		got := test.p.Norm()
		if got != test.want {
			t.Errorf("PointNorm(%v) = %v, want %v", test.p, got, test.want)
		}
	}
}

func TestPointDistance(t *testing.T) {
	tests := []struct {
		p, q Point
		want float64
	}{
		{Point{0, 0}, Point{3, 4}, 5},
		{Point{1, 1}, Point{2, 2}, math.Sqrt(2)},
		{Point{1, 1}, Point{1, 1}, 0},
	}

	for _, test := range tests {
		got := test.p.Distance(test.q)
		if got != test.want {
			t.Errorf("PointDistance(%v, %v) = %v, want %v", test.p, test.q, got, test.want)
		}
	}
}
