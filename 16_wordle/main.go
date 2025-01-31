package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

type List struct {
	words [][]byte
}

func main() {
	var input []byte

	rawWordList, err := os.ReadFile("word_list/word_list")
	if err != nil {
		log.Fatal(err)
	}

	wordList := bytes.Split([]byte(rawWordList), []byte("\n"))

	validWords := List{
		words: wordList,
	}

	chosenWord := validWords.Random()
	// debug
	fmt.Println(string(chosenWord))

	for i := 0; ; i++ {
		fmt.Printf("Enter 5 letter word\n>")
		fmt.Scan(&input)

		inputStr := string(input)
		inputSlice := strings.Split(inputStr, "")

		fmt.Println(checkChar(inputSlice))

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

/*
	Check every character of `input` and return character(s) that are in `chosenWord`

	Convert at some point between byte and string type to return human readable output

	- Loop through each character
	- Add character that is contained to new slice
	- Return slice of contained characters
*/

func checkChar(input []string) []string {
	var containedSlice []string

	for _, value := range input {
		if slices.Contains(input, value) {
			containedSlice = append(containedSlice, value)
		}
	}
	return containedSlice
}
