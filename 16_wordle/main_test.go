package main

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
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

	list := filepath.Join(tmpDir, "list")

	err := os.WriteFile(list, []byte("brown"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	got, _ := NewList(list)

	want := List{
		words: [][]byte{
			[]byte("brown"),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but got: %v", want, got)
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
