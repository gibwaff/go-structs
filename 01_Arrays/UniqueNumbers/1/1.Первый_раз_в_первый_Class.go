package main

import (
	"fmt"
)

func findMaxi(word []rune, n int) (foundMaxi []rune) {

	var prevMax rune = 2147483647
	for i := 0; i < n; i++ {
		var theMaxNow rune = -1
		for _, letter := range word {
			if letter > theMaxNow && letter < prevMax {
				theMaxNow = letter
			}
		}
		if theMaxNow != -1 {
			foundMaxi = append(foundMaxi, theMaxNow)
			prevMax = theMaxNow
		} else {
			break
		}
	}

	return
}

func main() {
	var word string = "АААФФФФФФФЖЫЫЫЫБЫРВАААААЛГГГХЫХЫБЛИА"
	var wordArray, wordArraySorted []rune
	wordArray = []rune(word)

	wordArraySorted = findMaxi(wordArray, len(wordArray))

	var learnedLetters string = string(wordArraySorted)
	fmt.Println(learnedLetters)
}
