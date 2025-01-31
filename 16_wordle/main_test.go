package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		l    List
		want bool
	}{
		{
			l: List{
				words: [][]byte{
					[]byte("brown"),
				},
			},
			want: false,
		},
		{
			l: List{
				words: [][]byte{
					[]byte("aaaa"),
				},
			},
			want: false,
		},
	}
	for _, test := range tests {
		got1 := test.l.Contains([]byte("brown"))
		got2 := test.l.Contains([]byte("aaaa"))

		if got1 != test.want && got2 != test.want {
			t.Errorf("Contains = %v, want %v and want %v", got1, got2, test.want)
		}
	}
}

// func TestIsValid(t *testing.T) {
// 	words := [][]byte{[]byte("brown"), []byte("possa")}

// 	tests := []struct {
// 		word []byte
// 		want bool
// 	}{
// 		{[]byte("brown"), true},
// 		{[]byte("aaaa"), false},
// 	}
// 	for _, test := range tests {
// 		got := isValid(test.word, words)

// 		if got != test.want {
// 			t.Errorf("valid = %v, want %v", test.word, test.want)
// 		}
// 	}
// }

// func TestIsChosen(t *testing.T) {
// 	tests := []struct {
// 		word       []byte
// 		chosenWord []byte
// 		want       bool
// 	}{
// 		{word: []byte("brown"), chosenWord: []byte("brown"), want: true},
// 		{word: []byte("aaaa"), chosenWord: []byte("brown"), want: false},
// 	}
// 	for _, test := range tests {
// 		got := isChosen(test.word, []byte("brown"))

// 		if got != test.want {
// 			t.Errorf("chosen = %v, want %v", test.word, test.want)
// 		}
// 	}
// }
