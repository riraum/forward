package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	var input []byte

	rawWordList, err := os.ReadFile("word_list/word_list")
	if err != nil {
		log.Fatal(err)
	}

	validWords := bytes.Split([]byte(rawWordList), []byte("\n"))
	chosenWordStr := "brown"
	chosenWord := []byte(chosenWordStr)

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
