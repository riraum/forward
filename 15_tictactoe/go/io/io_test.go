package io

import (
	"testing"
)

// func testRead(t *testing.T) {
// 	tests := []struct {
// 		input int
// 		want int
// 	}{
// 		{}
// 	}
// }

func TestValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"3", true},
		{"99", false},
		{"", false},
		{"abc", false},
	}

	for _, test := range tests {
		got := valid(test.input)
		if got != test.want {
			t.Errorf("valid = %v, want %v", test.input, test.want)
		}
	}
}
