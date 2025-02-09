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
	// debug
	// fmt.Println("chosenWord:", string(chosenWord))
	fmt.Printf("Enter 5 letter word\n")

	for i := 0; ; i++ {
		// input validation
		_, err := fmt.Scan()
		if err != nil {
			fmt.Println("Character lenght has to be 5!")
		}
		fmt.Printf(">")
		fmt.Scan(&input)
		fmt.Println(formatInput(input))

		getResult := checkChar(chosenWord, input)
		fmt.Println(strings.Join(coloredResult(getResult)[:], ""))

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
	return string(bytes.Join(bytes.Split(input, []byte("")), []byte(" ")))
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
		return bytes.Equal(word, value)
	}
	return false
}

func (l List) Random() []byte {
	return l.words[rand.IntN(len(l.words))]
}

func checkChar(chosenWord, input []byte) []int {
	var containedSlice []int
	// 1 = contained but incorrect loc
	// 2 = contained in correct loc
	// 0 not contained
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
