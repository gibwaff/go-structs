package main

import "fmt"

func isUnique(arr [][]int) (report int) {
	var arrOne []int
	for _, client := range arr {
		arrOne = append(arrOne, client...)
	}
	sortedArrOne := sortArray(arrOne)
	currentPay, report := 0, -1

	for _, pay := range sortedArrOne {
		if pay == currentPay {
			report = pay

			return
		}
		currentPay = pay
	}

	return
}

func sortArray(arr []int) (sortedArr []int) {
	for len(arr) != 0 {
		var currentMax, maxIndex = 0, -1
		for index, el := range arr {
			if el >= currentMax {
				currentMax = el
				maxIndex = index
			}
		}
		sortedArr = append(sortedArr, currentMax)
		arr = removeElem(arr, maxIndex)
	}

	return
}

func removeElem(arr []int, index int) []int {

	return append(arr[:index], arr[index+1:]...)
}

func main() {
	var report = [][]int{
		{12391203, 3828382, 334934939},
		{45345345, 5341312, 55345345},
		{334934939, 1234122, 657657},
	}

	paymentReport := isUnique(report)
	fmt.Println(paymentReport)
}
