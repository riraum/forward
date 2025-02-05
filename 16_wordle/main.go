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
	chosenWordStr := string(chosenWord)
	chosenWordSlice := strings.Split(chosenWordStr, "")
	// debug
	// fmt.Println(chosenWordSlice)
	fmt.Println("chosenWord:", string(chosenWordStr))

	for i := 0; ; i++ {
		fmt.Printf("Enter 5 letter word\n>")
		fmt.Scan(&input)

		// check characters
		inputStr := string(input)
		inputSlice := strings.Split(inputStr, "")
		// fmt.Println(inputSlice)
		// fmt.Println(inputSlice[0])
		// debug
		// fmt.Println(slices.Contains(inputSlice, "e"))
		// fmt.Println(slices.Contains(inputSlice, "t"))
		fmt.Println(inputSlice)
		// fmt.Println(checkChar(chosenWordSlice, inputSlice))

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

# Convert at some point between byte and string type to return human readable output

- Loop through each character
- Add character that is contained to new slice
- Return slice of contained characters
*/
func checkChar(chosenWordSlice []string, inputSlice []string) []int {
	var containedSlice []int
	// 1 = contained but incorrect loc
	// 2 = contained in correct loc
	// 0 not contained

	// correct word check
	if slices.Equal(chosenWordSlice, inputSlice) {
		return []int{2, 2, 2, 2, 2}
	}

	for index, value := range inputSlice {
		if value == chosenWordSlice[index] {
			containedSlice = append(containedSlice, 2)
		} else if slices.Contains(chosenWordSlice, inputSlice[index]) {
			containedSlice = append(containedSlice, 1)
		} else {
			containedSlice = append(containedSlice, 0)
		}
	}

	return containedSlice
}
