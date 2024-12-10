package main

import "fmt"

func handleHash(value int) (hash int) {
	if value > 0 {
		hash = value
	} else {
		hash = 100 - value
	}

	return
}

func add(inputHashMap []int, value int) []int {
	hash := handleHash(value)
	index := hash % 100
	inputHashMap[index] = value
	return inputHashMap
}

func isHere(inputHashMap []int, value int) bool {
	hash := handleHash(value)
	index := hash % 100
	if inputHashMap[index] != 0 {
		return true
	} else {
		return false
	}
}

func isThereTwoNumbers(numbers []int, X int) bool {
	var hashNumbers [100]int
	for _, el := range numbers {
		hashNumbers = [100]int(add(hashNumbers[:], el))
		if isHere(hashNumbers[:], X-el) {
			return true
		}
	}
	return false
}

func main() {
	var numbers = []int{2, 7, -7, 8, 7, 10}
	fmt.Println(isThereTwoNumbers(numbers, 11))
}
