package main

import (
	"fmt"
)

func main() {
	var words = [][]rune{
		{'б', 'а', 'к', 'а', 'ш'},
		{'к', 'а', 'т', 'e', 'н', 'к'},
		{'г', 'у', 'с', 'ъ', 'к'},
		{'б', 'а', 'к', 'а', 'ш'},
		{'а', 'к', 'а', 'ш', 'б'},
		{'ш', 'б', 'а', 'к', 'а'},
	}

	permutateWords := func(words [][]rune) {
		for i := 0; i < len(words); i++ {
			for k := i + 1; k < len(words); k++ {
				if string(words[i]) == string(words[k]) {
					c := words[k][0]
					for let := 1; let < len(words[k]); let++ {
						words[k][let-1] = words[k][let]
					}
					words[k][len(words[k])-1] = c
					break
				}
			}
		}
	}

	permutateWords(words)
	for _, el := range words {
		fmt.Println(string(el))
	}
}
