package main

import (
	"log"
	"os"
	"reflect"
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
	// tmpDir := t.TempDir()
	err := os.WriteFile("t.TempDir(dir/list)", []byte("brown"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	got, _ := NewList("dir/list")

	want := List{
		words: [][]byte{
			[]byte("brown"),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but got: %v", want, got)
	}
}
