package io

import (
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{3, true},
		{99, false},
	}

	for _, test := range tests {
		got := valid(test.input)
		if got != test.want {
			t.Errorf("valid = %v, want %v", test.input, test.want)
		}
	}
}
