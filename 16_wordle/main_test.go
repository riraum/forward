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

	word1 := "brown"
	word1Arr := []byte(word1)
	word2 := "aaaa"
	word2Arr := []byte(word2)
	tests := []struct {
		wordArr []byte
		words   [][]byte
		want    bool
	}{
		{word1Arr, words, true},
		{word2Arr, words, false},
	}
	for _, test := range tests {
		got := isValid(test.wordArr, test.words)
		if got != test.want {
			t.Errorf("valid = %v, want %v", test.wordArr, test.want)
		}
	}
}
