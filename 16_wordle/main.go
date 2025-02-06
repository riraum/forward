package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
)

type List struct {
	words [][]byte
}

func main() {
	var input []byte

	validWords, _ := NewList("word_list/word_list")
	// debug
	fmt.Println(validWords)

	chosenWord := validWords.Random()
	// debug
	fmt.Println(string(chosenWord))

	for i := 0; ; i++ {
		fmt.Printf("Enter 5 letter word\n>")
		fmt.Scan(&input)

		if bytes.Equal(input, chosenWord) {
			fmt.Print("Chosen word, yay!\n")
			break
		}

		if validWords.Contains(input) {
			fmt.Print("Valid word, but not chosen word, try again!\n")
			continue
		}

		fmt.Print("Invalid word, try again!\n")
	}
}

func NewList(path string) (List, error) {
	rawWordList, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
		return List{}, err
	}

	wordList := bytes.Split([]byte(rawWordList), []byte("\n"))

	validWords := List{
		words: wordList,
	}
	return validWords, nil
}

func (l List) Contains(word []byte) bool {
	for _, value := range l.words {
		if bytes.Equal(word, value) {
			return true
		}
	}
	return false
}

func (l List) Random() []byte {
	return l.words[rand.IntN(len(l.words))]
}
