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

// func (slice []string) checkValid(input) bool {
// 	for value := range slice {
// 		if input ==  value {
// 			return true
// 		}
// 	}
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
		fmt.Printf("Enter 5 digit word\n>")
		fmt.Scan(&input)
		// check if word on list
		if isValid(input, words) {
			break
		}

		// for _, value := range words {
		// 	// bytes.Equal(value, input)
		// 	// if valid(input, word_list.wordSlice) {
		// 	if bytes.Equal(value, input) {
		// 		fmt.Print("Correct word!\n")
		// 		// if valid(input, word_list.wordSlice) {
		// 		break
		// 	}
		// debug
		// fmt.Print("Not correct word!\n")
		// range loop
		// break
	}
	// input loop
	// break
}

// end main
// }

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
	fmt.Print("Invalid word!\n")
	return false
}

//
// 	for _, value := range formattedData {
// 		if bytes.Equal(value, input) {
// 			// if valid(input, word_list.wordSlice) {
// 			break
// 		} else {
// 			fmt.Print("Not correct word, try again!\n")
// 			// break
// 		}
// 		// range loop
// 		fmt.Print("Correct word!\n")
// 		break
// 	}
// 	// input loop
// }

// debug
// 	scanner := bufio.NewScanner(strings.NewReader(string(data)))
// 	scanner.Split(bufio.ScanWords)
// 	count := 0
// 	for scanner.Scan() {
// 		count++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(os.Stderr, "reading input:", err)
// 	}
// 	fmt.Printf("%d\n", count)
// }

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
