package main

import (
	"os"
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
	testDir := t.TempDir()
	tests := []struct {
	}
	os.WriteFile("testDir", []byte("abc"), 0666)
}
