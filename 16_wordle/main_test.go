package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	word1 := []byte("brown")
	word2 := []byte("aaaa")
	word3 := []byte("possa")

	words := [][]byte{word1, word3}

	tests := []struct {
		word []byte
		want bool
	}{
		{word1, true},
		{word2, false},
	}
	for _, test := range tests {
		got := isValid(test.word, words)

		if got != test.want {
			t.Errorf("valid = %v, want %v", test.word, test.want)
		}
	}
}

func TestIsChosen(t *testing.T) {
	word1 := []byte("brown")
	word2 := []byte("aaaa")

	chosenWord := word1

	tests := []struct {
		word []byte
		want bool
	}{
		{word1, true},
		{word2, false},
	}
	for _, test := range tests {
		got := isChosen(test.word, chosenWord)

		if got != test.want {
			t.Errorf("chosen = %v, want %v", test.word, test.want)
		}
	}
}
