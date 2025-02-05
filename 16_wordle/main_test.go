package main

import (
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

func TestCheckCharPrecise(t *testing.T) {
	tests := []struct {
		input  []string
		chosen []string
		want   []int
	}{
		{input: []string{"a", "a", "a", "a", "a"},
			chosen: []string{"a", "a", "a", "a", "b"},
			want:   []int{2, 2, 2, 2, 1},
		},
		{
			input:  []string{"a", "c", "a", "a", "b"},
			chosen: []string{"a", "a", "a", "b", "a"},
			want:   []int{2, 0, 2, 1, 1},
		},
	}
	for _, test := range tests {
		got := checkCharPrecise(test.chosen, test.input)

		if !slices.Equal(got, test.want) {
			t.Errorf("checkCharPrecise = %v, want %v", got, test.want)
		}
	}

}
