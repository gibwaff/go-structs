package main

import "fmt"

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		pt := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[pt] {
				pt = j
			}
		}
		for ; pt > 0; pt-- {
			arr[pt-1], arr[pt] = arr[pt], arr[pt-1]
		}
	}
	for pt := len(arr) - 1; pt > 0; pt-- {
		arr[pt-1], arr[pt] = arr[pt], arr[pt-1]
	}

	return arr
}

func main() {
	MyArray := []int{10, 2, 4, 15, 3, 0, 1}
	fmt.Println(SelectionSort(MyArray))
}
