package main

import (
	"log"
	"os"
	"path/filepath"
	"slices"
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

func TestNewList(t *testing.T) {
	tmpDir := t.TempDir()

	listPath := filepath.Join(tmpDir, "listPath")

	err := os.WriteFile(listPath, []byte("brown\nrossa\ncuppy"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		want List
	}{
		{
			want: List{
				words: [][]byte{
					[]byte("brown"), []byte("rossa"), []byte("cuppy"),
				},
			},
		},
		// {
		// 	want: List{
		// 		words: [][]byte{
		// 			[]byte("rossa"), []byte("cuppy"),
		// 		},
		// 	},
		// },
	}

	got, _ := NewList(listPath)

	for _, test := range tests {
		if len(got.words) != len(test.want.words) {
			t.Errorf("want: %v\n but got: %v\n", len(test.want.words), len(got.words))
		}
	}
	// slices.Equal(got.words, want.words)
	// for index := range got.words {
	// 	want.words[index] != got.words[index]
	//{
	// 	t.Errorf("want: %s\n but got: %s\n", want.words[index], want.words[index]),
	// }
	// }

	for index, test := range tests {
		if string(got.words[index]) != string(test.want.words[index]) {
			t.Errorf("want: %s\n but got: %s\n", test.want.words[index], got.words[index])
		}
	}
}

func TestCheckChar(t *testing.T) {
	tests := []struct {
		input  []byte
		chosen []byte
		want   []int
	}{
		{input: []byte("aaaaa"),
			chosen: []byte("aaaab"),
			want:   []int{2, 2, 2, 2, 1},
		},
		{
			input:  []byte("acaab"),
			chosen: []byte("aaaba"),
			want:   []int{2, 0, 2, 1, 1},
		},
		{
			input:  []byte("xy"),
			chosen: []byte("ab"),
			want:   []int{0, 0},
		},
	}
	for _, test := range tests {
		got := checkChar(test.chosen, test.input)

		if !slices.Equal(got, test.want) {
			t.Errorf("checkCharPrecise = %v, want %v", got, test.want)
		}
	}

}
