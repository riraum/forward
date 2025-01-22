package io

import (
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		input     string
		testSlice []int
		want      bool
	}{
		{"3", []int{3}, true},
		{"99", []int{}, false},
		{"", []int{}, false},
		{"abc", []int{}, false},
	}

	for _, test := range tests {
		got := valid(test.input, []int{3})
		if got != test.want {
			t.Errorf("valid = %v, want %v", test.input, test.want)
		}
	}
}
