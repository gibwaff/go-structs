package main

import "fmt"

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for pt, j := i-1, i; pt > -1; pt-- {
			if arr[pt] > arr[j] {
				arr[pt], arr[j] = arr[j], arr[pt]
				j--
			} else {
				break
			}
		}
	}
	return arr
}

func main() {
	MyArray := []int{3, 7, 1, 1, 0, 2, 4}
	fmt.Println(InsertionSort(MyArray))
}
