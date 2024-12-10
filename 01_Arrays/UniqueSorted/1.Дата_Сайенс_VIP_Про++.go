package main

import (
	"fmt"
)

func sortArr(arr []int) (sortedArr []int) {
	for len(arr) != 0 {
		var currentMin, minIndex int = 18446744073709551, -1
		for index, el := range arr {
			if el < currentMin {
				currentMin = el
				minIndex = index
			}
		}
		sortedArr = append(sortedArr, currentMin)
		arr = removeElem(arr, minIndex)
	}

	return
}

func removeElem(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	var phoneArray = []int{
		+79000000000, +79000000003,
		+79000000001, +79000000002,
		+79000000003, +79000000004,
		+79000000003, +79000000000,
		+79000000002, +79000000003,
	}

	sortedPhoneArray := append(sortArr(phoneArray), -1)
	currentPhoneNumber, numberCounter := sortedPhoneArray[0], -1

	for _, number := range sortedPhoneArray {
		numberCounter++
		if currentPhoneNumber != number {
			fmt.Printf("+%d - поступило заявок: %d\n", currentPhoneNumber, numberCounter)
			numberCounter = 0
		}
		currentPhoneNumber = number
	}
}
