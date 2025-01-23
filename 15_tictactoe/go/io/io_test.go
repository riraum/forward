package io

import (
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		input      string
		validSlice []int
		want       bool
	}{
		{"3", []int{3}, true},
		{"3", []int{4}, false},
		{"99", []int{}, false},
		{"", []int{}, false},
		{"abc", []int{}, false},
	}

	for _, test := range tests {
		got := valid(test.input, test.validSlice)
		if got != test.want {
			t.Errorf("valid = %v, want %v", test.input, test.want)
		}
	}
}
