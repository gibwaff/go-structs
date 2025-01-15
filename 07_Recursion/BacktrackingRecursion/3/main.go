package main

import "fmt"

func ButtonInit(num rune) []string {
	switch num {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	case '0':
		return []string{" "}
	}
	return []string{""}
}

func MessagesRecursion(first, last int, input_arr []rune, output string) {
	if first == last {
		fmt.Println(output)
		return
	}

	possible_letters := ButtonInit(input_arr[first])
	for current := 0; current < len(possible_letters); current++ {
		new_output := output + possible_letters[current]
		MessagesRecursion(first+1, last, input_arr, new_output)
	}
}

func PossibleMessages(message string) {
	message_rune := []rune(message)
	MessagesRecursion(0, len(message_rune), message_rune, "")
}

func main() {
	PossibleMessages("43")
}
