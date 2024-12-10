package main

import (
	"fmt"
	"math"
)

func findAverage(arr []int) (average float32) {
	var sum float32
	for _, el := range arr {
		sum += float32(el)
	}
	average = sum / float32(len(arr))
	return
}

func main() {
	var income = [][]int{
		{95, 67, 13, 55, 44, 11, 10},
		{7, 190, 4, 44, 11, 1, 99},
		{0, 5, -1, 500, 14, 90, 1},
	}
	var bestCustomerIndex = 0
	var bestIncome = float32(math.MinInt)
	for index, customerIncome := range income {
		if bestIncome < findAverage(customerIncome) {
			bestCustomerIndex = index
		}
	}
	fmt.Println(bestCustomerIndex)
}
