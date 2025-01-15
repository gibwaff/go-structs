package main

import "fmt"

func MergeTwoArrays(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	for i, j, pt := 0, 0, 0; i != len(a) || j != len(b); pt++ {
		if i == len(a) {
			res[pt] = b[j]
			j++
		} else if j == len(b) {
			res[pt] = a[i]
			i++
		} else if a[i] < b[j] {
			res[pt] = a[i]
			i++
		} else {
			res[pt] = b[j]
			j++
		}
	}
	return res
}

func MergeSort(arr []int) []int {
	mid := len(arr) / 2
	a := arr[:mid]
	b := arr[mid:]

	if len(a) == 1 && len(b) == 1 {
		return MergeTwoArrays(a, b)
	} else if len(a) == 1 {
		return MergeTwoArrays(a, MergeSort(b))
	} else if len(b) == 1 {
		return MergeTwoArrays(MergeSort(a), b)
	} else {
		return MergeTwoArrays(MergeSort(a), MergeSort(b))
	}
}

func main() {
	MyArray := []int{2, 1, 8, 7, 3, 5, 0}
	fmt.Println(MergeSort(MyArray))
}
