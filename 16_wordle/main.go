package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

// func valid(input []byte) bool {
// 	if bytes.Contains(input []byte) {
// 		return true
// 	}
// 	return false
// }

func main() {
	var input []byte
	// var wl []string
	// wl = word_list.wordSlice
	// Read word list
	data, err := os.ReadFile("word_list/word_list")
	if err != nil {
		log.Fatal(err)
	}
	// os.Stdout.Write(data)
	// debug
	// fmt.Print(data)
	// separate bytes
	words := bytes.Split([]byte(data), []byte("\n"))
	// debug
	// fmt.Printf("%q\n", formattedData)
	// fmt.Printf("%q\n", bytes.Split([]byte(data), []byte("\n")))
	// fmt.Println(bytes.Contains([]byte("brown"), formattedData))
	for i := 0; ; i++ {
		fmt.Printf("Enter 5 letter word\n>")
		fmt.Scan(&input)
		// check if word on list
		if isValid(input, words) {
			break
		}
	}
}

func isValid(word []byte, validWords [][]byte) bool {
	// if slices.Contains(validWords, word) {
	// 	return true
	// }
	for _, value := range validWords {
		// bytes.Equal(word, input)
		// if valid(input, word_list.wordSlice) {
		if bytes.Equal(word, value) {
			fmt.Print("Valid word!\n")
			// if valid(input, word_list.wordSlice) {
			return true
		}
	}
	// debug
	fmt.Print("Invalid word, try again!\n")
	return false
}
