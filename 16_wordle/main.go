package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func valid(input string, wordSlice []string) bool {
	if slices.Contains(wordSlice, input) {
		return true
	}
	return false
}

func main() {
	// var input string
	// var wl []string
	// wl = word_list.wordSlice

	// Read word list
	data, err := os.ReadFile("word_list/word_list")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
	// debug
	fmt.Print(data)

	// debug
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}

// for i := 0; ; i++ {
// 	fmt.Printf("Enter 5 digit word\n>")
// 	fmt.Scan(&input)

// for value, index := range wordSlice {
// 	if input == value {
// 		// if valid(input, word_list.wordSlice) {
// 		fmt.Print("Correct word!\n")
// 		break
// 	} else {
// 		fmt.Print("Not correct word, try again!\n")
// 	}
// }

// func (slice []string) checkValid(input) bool {
// 	for value := range slice {
// 		if input ==  value {
// 			return true
// 		}
// 	}
// }
// }
// }
