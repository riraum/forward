package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	words := [][]byte{[]byte("brown"), []byte("possa")}

	tests := []struct {
		word []byte
		want bool
	}{
		{[]byte("brown"), true},
		{[]byte("aaaa"), false},
	}
	for _, test := range tests {
		got := isValid(test.word, words)

		if got != test.want {
			t.Errorf("valid = %v, want %v", test.word, test.want)
		}
	}
}

func TestIsChosen(t *testing.T) {
	tests := []struct {
		word       []byte
		chosenWord []byte
		want       bool
	}{
		{word: []byte("brown"), chosenWord: []byte("brown"), want: true},
		{word: []byte("aaaa"), chosenWord: []byte("brown"), want: false},
	}
	for _, test := range tests {
		got := isChosen(test.word, []byte("brown"))

		if got != test.want {
			t.Errorf("chosen = %v, want %v", test.word, test.want)
		}
	}
}
