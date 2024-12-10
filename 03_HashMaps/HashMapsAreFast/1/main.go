package main

import "fmt"

func getIniqueNumbers(inputArray []int) (uniqueNumbersArray []int) {
	var hashNumbers [13]int
	for _, el := range inputArray {
		hashNumbers, success := add(hashNumbers[:], el)
		if success {
			uniqueNumbersArray = append(uniqueNumbersArray, get(hashNumbers, el))
		}
	}
	return
}

func hashInt(input int) (index int) {
	index = 1
	var numbers []int

	del := 1
	for input >= del {
		del *= 10
	}
	del /= 10
	for del != 0 {
		numbers = append(numbers, (input/del)%10)
		del /= 10
	}

	for _, el := range numbers {
		if el != 0 {
			index *= el
		}
	}

	return
}

func add(inputHashMap []int, value int) ([]int, bool) {
	success := false
	hash := hashInt(value)
	index := hash % 13
	if inputHashMap[index] == 0 {
		inputHashMap[index] = value
		success = true
	}
	return inputHashMap, success
}

func get(inputHashMap []int, key int) int {
	hash := hashInt(key)
	index := hash % 13
	value := inputHashMap[index]
	return value
}

func main() {
	var numbers = []int{2223, 102, 2223, 102, 4589, 6803, 102}

	sortedNumbers := getIniqueNumbers(numbers)

	fmt.Println(sortedNumbers)

}
