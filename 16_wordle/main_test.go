package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	// data, err := os.ReadFile("word_list/word_list")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	word1 := []byte("brown")
	word2 := []byte("aaaa")
	word3 := []byte("possa")

	words := [][]byte{word1, word3}

	tests := []struct {
		word  []byte
		words [][]byte
		want  bool
	}{
		{word1, words, true},
		{word2, words, false},
	}
	for _, test := range tests {
		got := isValid(test.word, test.words)

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
		word       []byte
		chosenWord []byte
		want       bool
	}{
		{word1, chosenWord, true},
		{word2, chosenWord, false},
	}
	for _, test := range tests {
		got := isChosen(test.word, chosenWord)

		if got != test.want {
			t.Errorf("chosen = %v, want %v", test.word, test.want)
		}
	}
}
