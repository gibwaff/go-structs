package main

import "fmt"

func RecursionEnum(first, last int, arr, print_arr []int) {
	if first == last {
		fmt.Println(print_arr)
		return
	}
	for current := 0; current < len(arr); current++ {
		print_arr[first] = arr[current]
		RecursionEnum(first+1, last, arr, print_arr)
	}
}

func Enumerate(arr []int) {
	full_arr := make([]int, len(arr))
	RecursionEnum(0, len(arr), arr, full_arr)
}

func main() {
	FichNums := []int{1, 2, 3}

	Enumerate(FichNums)
}
