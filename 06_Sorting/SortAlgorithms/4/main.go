package main

import "fmt"

func FindMax(arr []int) (maxi int) {
	maxi = 0
	for _, el := range arr {
		if el > maxi {
			maxi = el
		}
	}
	return
}

func CountingSort(arr []int) []int {
	maxi := FindMax(arr)
	counter := make([]int, maxi)
	new_arr := make([]int, len(arr))
	for i := range arr {
		counter[arr[i]-1]++
	}
	pt := 0
	for i := range counter {
		if counter[i] != 0 {
			for j := counter[i]; j > 0; j-- {
				new_arr[pt] = i + 1
				pt++
			}
		}
	}
	return new_arr
}

func main() {
	MatchesArray := []int{10, 10, 8, 9, 5, 1, 1, 2, 7, 6}
	SortedMatchesArray := CountingSort(MatchesArray)
	fmt.Println(SortedMatchesArray)
}
