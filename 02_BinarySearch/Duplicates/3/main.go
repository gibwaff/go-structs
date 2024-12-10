package main

import (
	"fmt"
	"math/rand"
)

func myIsSorted(array []int) (sorted bool) {
	sorted = true
	for i := 1; i < len(array); i++ {
		if array[i-1] > array[i] {
			sorted = false
			break
		}
	}
	return sorted
}

func main() {
	var myArray = []int{15, 22, 10, 5, 1, 8, 100}

	randomSort := func(array []int) {
		for flag := myIsSorted(array); !flag; flag = myIsSorted(array) {
			num1 := rand.Intn(len(array))
			num2 := rand.Intn(len(array))
			for num1 == num2 {
				num2 = rand.Intn(len(array))
			}
			array[num1], array[num2] = array[num2], array[num1]
		}
	}

	randomSort(myArray)
	fmt.Println(myArray)
}
