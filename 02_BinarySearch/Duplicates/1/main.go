package main

import "fmt"

func main() {
	var textLength = 252
	var text = make([]rune, textLength)
	text_str := "HAello, Wrld!"
	for i, char := range text_str {
		text[i] = char
	}

	insertCharacterAt := func(position int, c rune) {
		for i := len(text) - 1; i > position; i-- {
			text[i] = text[i-1]
		}
		text[position] = c
	}
	insertCharacterAt(9, 'o')

	backspace := func(position int) {
		for i := position + 1; i < len(text); i++ {
			text[i-1] = text[i]
		}
	}
	backspace(1)

	fmt.Println(string(text))
}
