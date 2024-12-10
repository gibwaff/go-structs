package main

import (
	"fmt"
	"math"
)

func findMaxUnderBoundary(inputArray []int, topBoundary int) (currentMax int) {
	currentMax = math.MinInt32
	for _, el := range inputArray {
		if el < topBoundary {
			currentMax = max(currentMax, el)
		}
	}
	return
}

func findTopElements(inputArray []int, numberOfElements int) (topElements []int) {
	if numberOfElements > len(inputArray) {
		fmt.Println("Выход за пределы массива")
		return
	}
	previousMax := math.MaxInt32
	for i := 0; i < numberOfElements; i++ {
		currentMax := findMaxUnderBoundary(inputArray, previousMax)
		previousMax = currentMax
		topElements = append(topElements, currentMax)
	}
	return
}

func findBottomElements(inputArray []int, numberOfElements int) (bottomElements []int) {

	return
}

func main() {
	array := []int{10, 66, 4, 16, 99, 35, 11, 123}

	top5 := findTopElements(array, 5)
	fmt.Println(top5)

	top8 := findTopElements(array, 9)
	fmt.Println(top8)

}
