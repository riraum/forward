package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestIsValid(t *testing.T) {
	data, err := os.ReadFile("word_list/word_list")
	if err != nil {
		log.Fatal(err)
	}
	words := bytes.Split([]byte(data), []byte("\n"))

	word1Str := "brown"
	word1 := []byte(word1Str)
	word2Str := "aaaa"
	word2 := []byte(word2Str)

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
	word1Str := "brown"
	word1 := []byte(word1Str)
	word2Str := "aaaa"
	word2 := []byte(word2Str)

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
