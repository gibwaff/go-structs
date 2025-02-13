package main

import "fmt"

func sortArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func findMaximumIncome(workRates []int, k int) (sum int) {
	workRates_sorted := sortArray(workRates)
	for i := 0; i < k; i++ {
		sum += workRates_sorted[i]
	}
	return
}

func main() {

	workRates := []int{60, 10, 20, 10, 30, 40, 50}
	freeTime := 3

	fmt.Println(findMaximumIncome(workRates, freeTime))
}
