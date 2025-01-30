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

	rawWordList, err := os.ReadFile("word_list/word_list")
	if err != nil {
		log.Fatal(err)
	}

	validWords := bytes.Split([]byte(rawWordList), []byte("\n"))
	// debug
	// fmt.Println(validWords)

	// chosenWord := randomWord(validWords)

	validWordsStruct := List{
		words: validWords,
	}
	// debug
	// fmt.Println(validWordsStruct)
	// fmt.Println(string(validWordsStruct.words[0]))

	chosenWord := randomWord(validWordsStruct.words)
	// debug
	// fmt.Println(string(chosenWord))

	for i := 0; ; i++ {
		fmt.Printf("Enter 5 letter word\n>")
		fmt.Scan(&input)
		if isChosen(input, chosenWord) {
			fmt.Print("Chosen word, yay!\n")
			break
		}
		if isValid(input, validWords) {
			fmt.Print("Valid word, but not chosen word, try again!\n")
			continue
		}
		fmt.Print("Invalid word, try again!\n")
	}
}

func (l List) IsValidList(word []byte) bool {
	for _, value := range l.words {
		if bytes.Equal(word, value) {
			return true
		}
	}
	return false
}

func isValid(word []byte, validWords [][]byte) bool {
	for _, value := range validWords {
		if bytes.Equal(word, value) {
			return true
		}
	}
	return false
}

func isChosen(word, chosenWord []byte) bool {
	return bytes.Equal(word, chosenWord)
}

func randomWord(validWords [][]byte) []byte {
	randomInt := rand.IntN(len(validWords))
	// debug
	fmt.Println(string(validWords[randomInt]))
	return validWords[randomInt]
}
