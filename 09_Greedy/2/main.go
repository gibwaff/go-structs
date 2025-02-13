package main

import "fmt"

func sort_meets(arr [][]int) [][]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j][1] > arr[j+1][1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func del(arr [][]int, i int) [][]int {
	return append(arr[:i], arr[i+1:]...)
}

func findMinimumManagers(meets [][]int) int {
	sorted_meets := sort_meets(meets)
	n := 0

	for len(sorted_meets) > 0 {
		n++
		sorted_meets = del(sorted_meets, 0)
		for i := 0; i < len(sorted_meets); {
			if sorted_meets[i][0] >= sorted_meets[0][1] {
				sorted_meets = del(sorted_meets, i)
			} else {
				i++
			}
		}
	}
	return n
}

func main() {
	meets := [][]int{{8, 15}, {10, 13}, {7, 9}, {7, 14}, {11, 12}, {8, 10}, {8, 9}}
	//[[7 9] [8 9] [8 10] [11 12] [10 13] [7 14] [8 15]] sorted

	managers := findMinimumManagers(meets)
	fmt.Println(managers)
}
