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

	validWords, err := NewList("word_list/word_list")
	if err != nil {
		log.Fatal(err)
	}

	chosenWord := validWords.Random()
	// debug
	fmt.Println("chosenWord:", string(chosenWord))
	fmt.Printf("Enter 5 letter word. You have 5 tries\n")

	for i := 1; i < 6; i++ {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if i == 5 {
			fmt.Printf("%v/5 Did you really try 5 times with words that don't even have 5 characters?Goodbye!\n", i)
			break
		}

		if len(string(input)) != 5 {
			fmt.Printf("%v/5 5 characters required. Try again!11!!\n", i)
			continue
		}

		fmt.Printf("%v/5 %v\n", i, formatInput(input))

		getResult := checkChar(chosenWord, input)

		if bytes.Equal(input, chosenWord) {
			fmt.Print("Chosen word, yay!\n")
			break
		}

		if validWords.Contains(input) {
			fmt.Println("   ", strings.Join(coloredResult(getResult)[:], ""))

			fmt.Print("Valid, but not chosen word, try again!\n")
			continue
		}

		fmt.Print("Invalid word, try again!\n")
	}
}

func NewList(path string) (List, error) {
	rawWordList, err := os.ReadFile(path)
	if err != nil {
		return List{}, err
	}
	wordList := bytes.Split((rawWordList), []byte("\n"))

	return List{
		words: wordList,
	}, err
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
		if bytes.Equal(word, value) {
			return true
		}
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
