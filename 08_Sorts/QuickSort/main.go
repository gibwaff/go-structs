package main

import "fmt"

func median(arr []int, l, r int) int {
	mid := (l + r) / 2
	if arr[l] > arr[mid] {
		arr[l], arr[mid] = arr[mid], arr[l]
	}
	if arr[l] > arr[r] {
		arr[l], arr[r] = arr[r], arr[l]
	}
	if arr[mid] > arr[r] {
		arr[mid], arr[r] = arr[r], arr[mid]
	}
	return arr[mid]
}

func partition(arr *[]int, l, r int) int {
	a := *arr
	pivot := median(a, l, r)
	i, j := l, r
	for i <= j {
		for a[i] < pivot {
			i++
		}
		for a[j] > pivot {
			j--
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	return i
}

func qsort(arr *[]int, l, r int) {
	if l >= r {
		return
	}
	m := partition(arr, l, r)
	qsort(arr, l, m-1)
	qsort(arr, m, r)
}

func main() {
	MyArray := []int{4, 5, 5, 3, 2, 2, 1}
	qsort(&MyArray, 0, len(MyArray)-1)
	fmt.Println(MyArray)
}
