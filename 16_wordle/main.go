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
	// chosenWordStrSlice := strings.Split(chosenWordStr, " ")
	// debug
	fmt.Println("chosenWord:", string(chosenWordStr))

	for i := 0; ; i++ {
		fmt.Printf("Enter 5 letter word\n>")
		fmt.Scan(&input)

		// check characters
		// fmt.Println(chosenWordStrSlice)
		inputStr := string(input)
		inputStrSlice := strings.Split(inputStr, "")
		inputStrRdy := strings.Join(inputStrSlice[:], " ")
		fmt.Println(inputStrRdy)

		checkResult := checkChar(chosenWord, input)
		// fmt.Println(checkResult)
		coloredOutput := coloredResult(checkResult)
		coloredOutputStr := strings.Join(coloredOutput[:], "")
		fmt.Println(coloredOutputStr)

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

func coloredResult(checkResult []int) []string {
	var result []string
	for _, value := range checkResult {
		if value == 2 {
			result = append(result, "🟩")
		} else if value == 1 {
			result = append(result, "🟨")
		} else {
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
func checkChar(chosenWordSlice, inputSlice []byte) []int {
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
