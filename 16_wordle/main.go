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

	validWords, _ := NewList("word_list/word_list")

	chosenWord := validWords.Random()
	fmt.Println("chosenWord:", string(chosenWord))

	for i := 0; ; i++ {
		fmt.Printf("Enter 5 letter word\n>")
		fmt.Scan(&input)
		formatInput(input)

		inputOutput := bytes.Join(bytes.Split(input, []byte("")), []byte(" "))
		fmt.Println(string(inputOutput))

		getResult := checkChar(chosenWord, input)
		coloredOutput := strings.Join(coloredResult(getResult)[:], "")
		fmt.Println(coloredOutput)

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

func formatInput(input []byte) string {
	inputOutput := bytes.Join(bytes.Split(input, []byte("")), []byte(" "))
	return string(inputOutput)
}

func coloredResult(checkResult []int) []string {
	var result []string
	for _, value := range checkResult {
		switch value {
		case 2:
			result = append(result, "🟩")
		case 1:
			result = append(result, "🟨")
		default:
			result = append(result, "🟥")
		}
	}
	return result

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
func checkChar(chosenWord, input []byte) []int {
	var containedSlice []int
	// 1 = contained but incorrect loc
	// 2 = contained in correct loc
	// 0 not contained

	// correct word check
	if slices.Equal(chosenWord, input) {
		return []int{2, 2, 2, 2, 2}
	}

	for index, value := range input {
		if value == chosenWord[index] {
			containedSlice = append(containedSlice, 2)
		} else if slices.Contains(chosenWord, input[index]) {
			containedSlice = append(containedSlice, 1)
		} else {
			containedSlice = append(containedSlice, 0)
		}
	}

	return containedSlice
}
